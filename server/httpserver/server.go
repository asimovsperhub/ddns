package httpserver

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/Agent/dns"
	"github.com/dnsdao/dnsdao.resolver/config"
	"github.com/dnsdao/dnsdao.resolver/database/ldb"
	"github.com/dnsdao/dnsdao.resolver/server/httpserver/api"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	DnsName                 = "/dns/name"
	DnsEthAddress           = "/dns/ethaddr"
	DnsIP4                  = "/dns/ip4"
	DnsBlockchainAddress    = "/dns/blockchainaddr"
	DnsCName                = "/dns/cname"
	DnsIP6                  = "/dns/ip6"
	DnsGetTotalBySubLen     = "/dns/did/top/get_second_counts"
	DnsQuerySubDomainByPage = "/dns/did/top/get_second_infos"
	DnsGetTotalByPrice      = "/dns/totalbyprice"
	DnsAddress              = "/dns/addr"
	GetEarningsByAddress    = "/dns/multisign/dids/get_todos"
	GetQuerySignTldList     = "/dns/multisign/tasks/get_list"
	GetCashDetailsByAddress = "/dns/getcashbyaddr"
	ConfResolve             = "/dns/confresolve"
	AddrDomainsResolve      = "/dns/resolve"
	DnsTokenId              = "/udid"
	AddrTopList             = "/dns/dids/top/get_list"
	AddrSubList             = "/dns/dids/second/get_list"
	GetOpenRegister         = "/dns/dids/top/get_open_dids"
	GetMyPassCardList       = "/dns/passcard/get_all"
	PostSignMint            = "/dns/passcard/signed_mint"
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
	for _, route := range rh.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
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
	wapi := api.NewApi()
	rh.HandleFunc(DnsName, Cors(ws.dnsNameResolve))
	rh.HandleFunc(DnsIP4, Cors(ws.ipv4Resolve))
	rh.HandleFunc(DnsBlockchainAddress, Cors(wapi.Blockchain))
	rh.HandleFunc(DnsCName, Cors(wapi.CName))
	rh.HandleFunc(DnsIP6, Cors(wapi.Ipv6))
	rh.HandleFunc(DnsGetTotalBySubLen, Cors(wapi.GetTotalBySubLen))
	rh.HandleFunc(DnsQuerySubDomainByPage, Cors(wapi.QuerySubDomainByPage))
	rh.HandleFunc(DnsAddress, Cors(wapi.AddrResolve))
	rh.HandleFunc(GetEarningsByAddress, Cors(wapi.GetEarningsByAddress))
	rh.HandleFunc(GetCashDetailsByAddress, Cors(wapi.GetCashDetailsByAddress))
	rh.HandleFunc(GetQuerySignTldList, Cors(wapi.GetSignTldListByDidName))
	rh.HandleFunc(ConfResolve, Cors(wapi.ConfResolve))
	rh.HandleFunc(AddrDomainsResolve, Cors(wapi.AddrDomainsResolve))
	rh.HandleFunc(DnsTokenId, Cors(wapi.DnsTokenId))
	rh.HandleFunc(AddrTopList, Cors(wapi.AddrTopList))
	rh.HandleFunc(AddrSubList, Cors(wapi.AddrSubList))
	rh.HandleFunc(GetOpenRegister, Cors(wapi.GetOpenRegister))
	rh.HandleFunc(GetMyPassCardList, Cors(wapi.GetMyPassCardList))
	rh.HandleFunc(PostSignMint, Cors(wapi.PostSignMint))
	server := &http.Server{
		Handler: rh,
	}

	ws.server = server

	return ws
}

func (ws *WebProxyServer) ipv4Resolve(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}

	query := request.URL.Query()
	msg := "not found ipv4"
	var (
		v  []string
		ok bool
	)

	if v, ok = query["ipv4"]; !ok {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
		return
	}

	ipv4 := v[0]
	db := ldb.GetLdb()
	key, err := db.GetTypeNameHash("A_" + ipv4)
	if err == nil {
		var respstr []byte
		val, _ := db.GetKey(key)
		respstr, err = json.Marshal(val)
		if err == nil {
			msg = string(respstr)
		}
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

}

func isRoot(name string) bool {
	arr := strings.Split(name, ".")
	if len(arr) > 1 {
		return false
	}
	return true
}

type SubNameResp struct {
	dns.SubName
	Root *dns.RootName `json:"root"`
}

func (ws *WebProxyServer) dnsNameResolve(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}

	query := request.URL.Query()
	log.Println("/dns/name query", query)
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

	ir := isRoot(name)

	nameHash := crypto.Keccak256([]byte(name))

	db := ldb.GetLdb()

	msg := "not found name"

	if ir {
		val, err := db.GetRootName(hex.EncodeToString(nameHash))
		if err == nil {
			msg = val
		}
	} else {
		val, err := db.GetSubName(hex.EncodeToString(nameHash))
		if err == nil {
			//resp := &SubNameResp{}
			//sub := &dns.SubName{}
			//err = json.Unmarshal([]byte(val), sub)
			//if err == nil {
			//	resp.SubName = *sub
			//	val, err = db.GetRootName(sub.RootHash)
			//	if err == nil {
			//		root := &dns.RootName{}
			//		err = json.Unmarshal([]byte(val), root)
			//		resp.Root = root
			//	}
			//}
			//var respstr []byte
			//respstr, err = json.Marshal(resp)
			//if err == nil {
			//	msg = string(respstr)
			//}
			msg = val
		}
	}

	writer.WriteHeader(200)
	writer.Write([]byte(msg))

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

func (ws *WebProxyServer) DnsTokenId(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}

	err, tokenId := splitPath(request.URL.Path)
	if err != nil {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "token Id error")
	}

	if tokenId < 0 || tokenId >= 9999 {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "token Id error")
	}

}

func (ws *WebProxyServer) ethAddrResolve(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}

	query := request.URL.Query()
	msg := "not found addr"
	var (
		v  []string
		ok bool
	)

	if v, ok = query["addr"]; !ok {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
		return
	}

	addr := v[0]
	db := ldb.GetLdb()
	val, err := db.GetOwnerNameHash(addr[2:])
	if err == nil {
		var respstr []byte
		var data [][]byte
		for _, v := range val {
			in, _ := db.GetKey(v)
			data = append(data, []byte(in))
		}
		respstr, err = json.Marshal(data)
		if err == nil {
			msg = string(respstr)
		}
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

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

	c := config.GetRConf().WConf

	webServer = NewWebServer(c.ListenServer)

	fmt.Println("start proxy at ", webServer.listenAddr, "  ...")

	webServer.Start()
}

func StopWebDaemon() {
	webServer.Shutdown()

	fmt.Println("stop proxy ...")
}
