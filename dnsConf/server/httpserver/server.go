package httpserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/dnsConf/config"
	"github.com/dnsdao/dnsdao.resolver/dnsConf/dnsconf"
	"github.com/dnsdao/dnsdao.resolver/dnsConf/server/httpserver/api"
	"io/ioutil"
	"log"

	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	DNSTypes          = "/conf/types"
	SaveChanges       = "/conf/save"
	DnsMapping        = "/conf/mapping"
	DnsReverseMapping = "/conf/reversemapping"
	DnsAuthMapping    = "/conf/authmapping"
)

func Cors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                                                            // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
		w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    //设置为true，允许ajax异步请求带cookie信息
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             //允许请求方法
		w.Header().Set("content-type", "application/json;charset=UTF-8")                                              //返回数据格式是json
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		f(w, r)
	}
}

type WebProxyServer struct {
	listenAddr string
	quit       chan struct{}
	server     *http.Server
}

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexpHander struct {
	routes []*route
}

func (rh *RegexpHander) Handle(pattern string, handler http.Handler) {
	rh.routes = append(rh.routes, &route{pattern: regexp.MustCompilePOSIX(pattern), handler: handler})
}

func (rh *RegexpHander) HandleFunc(pattern string, handleFunc func(http.ResponseWriter, *http.Request)) {
	rh.routes = append(rh.routes, &route{pattern: regexp.MustCompilePOSIX(pattern), handler: http.HandlerFunc(handleFunc)})
}

func (rh *RegexpHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, item := range rh.routes {
		if item.pattern.MatchString(r.URL.Path) {
			item.handler.ServeHTTP(w, r)
			return
		}
	}
	// no pattern matched; send 404 response
	http.NotFound(w, r)
}

func NewWebServer(networkAddr string) *WebProxyServer {
	ws := WebProxyServer{
		listenAddr: networkAddr,
		quit:       make(chan struct{}, 8),
	}

	return ws.init()

}

func (ws *WebProxyServer) init() *WebProxyServer {
	rh := &RegexpHander{
		routes: make([]*route, 0),
	}

	rh.HandleFunc(DNSTypes, Cors(ws.dnsTypes))
	rh.HandleFunc(SaveChanges, Cors(ws.saveChanges))
	rh.HandleFunc(DnsMapping, Cors(ws.namemapping))
	rh.HandleFunc(DnsReverseMapping, Cors(ws.revertMapping))
	rh.HandleFunc(DnsAuthMapping, Cors(ws.authMapping))

	server := &http.Server{
		Handler: rh,
	}

	ws.server = server

	return ws
}

func (ws *WebProxyServer) namemapping(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}
	query := request.URL.Query()
	log.Println(request.URL.Path, query)
	var (
		v  []string
		ok bool
	)

	if v, ok = query["name"]; !ok {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
		return
	}

	name := v[0]

	var (
		conf *dnsconf.NameConfDb
		err  error
	)

	resp := &api.Response{
		ErrMsg: "success",
	}

	if conf, err = dnsconf.GetNameConfFromDB(name); err != nil {
		resp.Code = api.RespErr
		resp.ErrMsg = err.Error()
	} else {
		resp.Result = conf
	}

	j, _ := json.Marshal(resp)

	writer.WriteHeader(200)
	writer.Write(j)

}

func (ws *WebProxyServer) authMapping(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}

	req := &dnsconf.NameConfWithSig{}

	resp := api.Response{
		ErrMsg: "success",
	}

	if data, err := ioutil.ReadAll(request.Body); err != nil {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
	} else {
		if err = json.Unmarshal(data, req); err != nil {
			resp.Code = api.UnmarshalSigContentErr
			resp.ErrMsg = err.Error()
		}
	}
	t := time.Now().UnixMicro()
	if resp.Code == 0 {
		log.Println(t, request.URL.Path, "\r\n", req.String())
		if err := req.SaveAuthEthChanges(); err != nil {
			resp.Code = api.SaveErr
			resp.ErrMsg = err.Error()
		}
	}

	data, _ := json.Marshal(resp)
	log.Println(t, request.URL.Path, string(data))
	writer.WriteHeader(200)
	writer.Write(data)

}
func (ws *WebProxyServer) revertMapping(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}
	query := request.URL.Query()
	log.Println(request.URL.Path, query)
	var (
		vtype   []string
		vmapKey []string
		ok      bool
	)

	if vtype, ok = query["type"]; !ok {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
		return
	}
	if vmapKey, ok = query["key"]; !ok {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
		return
	}

	typ := vtype[0]
	key := vmapKey[0]

	fmt.Println(request.URL.Path, typ, key)
	resp := &api.Response{
		ErrMsg: "success",
	}

	n := dnsconf.GetRevertValue(typ, key)
	if n == nil {
		resp.Code = api.RespErr
		resp.ErrMsg = "not found"
	} else {
		resp.Result = n
	}

	j, _ := json.Marshal(resp)

	writer.WriteHeader(200)
	writer.Write(j)

}

func (ws *WebProxyServer) saveChanges(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}

	req := &dnsconf.NameConfWithSig{}

	resp := api.Response{
		ErrMsg: "success",
	}

	if data, err := ioutil.ReadAll(request.Body); err != nil {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
	} else {
		if err = json.Unmarshal(data, req); err != nil {
			resp.Code = api.UnmarshalSigContentErr
			resp.ErrMsg = err.Error()
		}
	}
	t := time.Now().UnixMicro()
	if resp.Code == 0 {
		log.Println(t, request.URL.Path, "\r\n", req.String())
		if err := req.SaveChanges(); err != nil {
			resp.Code = api.SaveErr
			resp.ErrMsg = err.Error()
		}
	}

	data, _ := json.Marshal(resp)
	log.Println(t, request.URL.Path, string(data))

	writer.WriteHeader(200)
	writer.Write(data)

}

func isRoot(name string) bool {
	arr := strings.Split(name, ".")
	if len(arr) > 1 {
		return false
	}
	return true
}

func (ws *WebProxyServer) dnsTypes(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}

	resp := api.Response{
		Code:   api.RespSuccess,
		ErrMsg: "success",
		Result: dnsconf.SupportConfItem(),
	}

	msg, _ := json.Marshal(resp)

	writer.WriteHeader(200)
	writer.Write(msg)
}

func splitPath(p string) (error, int32) {
	sarr := strings.Split(p, "/")

	if len(sarr) == 0 {
		return errors.New("path error"), -1
	}

	stokenId := sarr[len(sarr)-1]

	tokenId, err := strconv.Atoi(stokenId)
	if err != nil {
		return err, -1
	}
	return nil, int32(tokenId)
}

func (ws *WebProxyServer) Start() error {
	if l, err := net.Listen("tcp4", ws.listenAddr); err != nil {
		panic("start wss server failed")
	} else {
		return ws.server.Serve(l)
	}
}

func (ws *WebProxyServer) Shutdown() error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ws.server.Shutdown(ctx)
}

var webServer *WebProxyServer

func StartWebDaemon() {

	//c := config.GetRConf().WConf

	webServer = NewWebServer(config.GetDnsConfSetting().ServerListenAddr)

	fmt.Println("start proxy at ", webServer.listenAddr, "  ...")

	webServer.Start()
}

func StopWebDaemon() {
	webServer.Shutdown()

	fmt.Println("stop proxy ...")
}
