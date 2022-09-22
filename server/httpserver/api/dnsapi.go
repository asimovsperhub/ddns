package api

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/Agent/dns"
	"github.com/dnsdao/dnsdao.resolver/Agent/dnsv2"
	"github.com/dnsdao/dnsdao.resolver/config"
	"github.com/dnsdao/dnsdao.resolver/database/ldb"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"unsafe"
)

type Api struct {
}

func NewApi() *Api {
	return &Api{}
}
func isRoot(name string) bool {
	arr := strings.Split(name, ".")
	if len(arr) > 1 {
		return false
	}
	return true
}
func byte32(s []byte) (a *[32]byte) {
	if len(a) <= len(s) {
		a = (*[len(a)]byte)(unsafe.Pointer(&s[0]))
	}
	return a
}

type QuerySub struct {
	Name       string         `json:"name"`
	Owner      common.Address `json:"owner"`
	Expiration *big.Int       `json:"expiration"`
}
type QuerySubDomainByPage struct {
	Total  int         `json:"total"`
	Result []*QuerySub `json:"result"`
}

type QueryName struct {
	Total  int      `json:"total"`
	Result []string `json:"result"`
}
type GetTotalByPrice struct {
	TotalPrice float64 `json:"total_price"`
}

func Paging(pageNumber int, pageSize int, resArr []string) []string {
	var res []string
	// 不带分页则返回全部
	if pageNumber == 0 && pageSize == 0 {
		return resArr
	}
	start := pageNumber * pageSize
	end := (pageNumber + 1) * pageSize
	if start <= len(resArr) {
		if end > len(resArr) {
			res = resArr[start:]
		} else {
			res = resArr[start:end]
		}
	} else {
		res = nil
	}
	return res
}
func (a *Api) Blockchain(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}

	query := request.URL.Query()
	msg := "not found blockchain"
	var (
		v  []string
		ok bool
	)

	if v, ok = query["blockchain"]; !ok {
		res_ := NotDataRes("not blockchain ")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}

	blockchain := v[0]
	db := ldb.GetLdb()
	val, err := db.GetKey("BlockChain_" + blockchain)
	if err == nil {
		var respstr []byte
		data, err := db.GetKey(val)
		if err == nil {
			respstr, err = json.Marshal(data)
			if err == nil {
				msg = string(respstr)
			}
		}
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

}

func (a *Api) CName(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}

	query := request.URL.Query()
	msg := "not found cname"
	var (
		v  []string
		ok bool
	)

	if v, ok = query["cname"]; !ok {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not cname")
		return
	}

	cname := v[0]
	db := ldb.GetLdb()
	val, err := db.GetKey("CName_" + cname)
	if err == nil {
		var respstr []byte
		data, err := db.GetKey(val)
		if err == nil {
			respstr, err = json.Marshal(data)
			if err == nil {
				msg = string(respstr)
			}
		}
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

}

func (a *Api) Ipv6(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}

	query := request.URL.Query()
	msg := "not found ipv6"
	var (
		v  []string
		ok bool
	)

	if v, ok = query["ipv6"]; !ok {
		res_ := NotDataRes("not ipv6")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}

	ipv6 := v[0]
	db := ldb.GetLdb()
	val, err := db.GetKey("AAAA_" + ipv6)
	if err == nil {
		var respstr []byte
		data, err := db.GetKey(val)
		if err == nil {
			respstr, err = json.Marshal(data)
			if err == nil {
				msg = string(respstr)
			}
		}
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

}

func (a *Api) GetTotalBySubLen(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}

	query := request.URL.Query()
	log.Println("/dns/totalbysub query", query)
	var (
		v  []string
		ok bool
	)

	if v, ok = query["tldName"]; !ok {
		res_ := NotDataRes("not tldName")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}

	name := v[0]

	ir := isRoot(name)

	nameHash := crypto.Keccak256([]byte(name))

	db := ldb.GetLdb()

	msg := "not found tldName"

	if ir {
		val, err := db.GetRootName(hex.EncodeToString(nameHash))
		if err == nil {
			RootNameInfo := new(dns.RootNameInfo)
			json.Unmarshal([]byte(val), RootNameInfo)
			res := &Res{
				Code:    1,
				Message: "ok",
				Data:    RootNameInfo.SubNameCount,
			}
			resbyte, _ := json.Marshal(res)
			msg = string(resbyte)
			writer.WriteHeader(200)
			writer.Write([]byte(msg))
			return
		} else {
			res_ := NotDataRes(fmt.Sprintf("%s not data", name))
			writer.WriteHeader(200)
			writer.Write([]byte(res_))
			return
		}
	} else {
		res_ := NotDataRes(fmt.Sprintf("%s:is not root", name))
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
}

func (a *Api) QuerySubDomainByPage(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}

	query := request.URL.Query()
	log.Println("/dns/querysub query", query)
	var (
		v  []string
		ok bool
	)
	minlen := ""
	maxlen := ""
	pageNumber := ""
	pageSize := ""
	rootname := ""
	if v, ok = query["tldName"]; !ok {
		res_ := NotDataRes("not tldName")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	} else {
		rootname = v[0]
	}

	if v, ok = query["minlen"]; ok {
		minlen = v[0]
	}
	if v, ok = query["maxlen"]; ok {
		maxlen = v[0]
	}
	if v, ok = query["pageNumber"]; ok {
		pageNumber = v[0]
	}
	if v, ok = query["pageSize"]; ok {
		pageSize = v[0]
	}

	ir := isRoot(rootname)

	nameHash := crypto.Keccak256([]byte(rootname))

	db := ldb.GetLdb()

	msg := "not found tldName"

	if ir {
		resArr := []string{}
		subnameList := []string{}
		val, err := db.GetRootName(hex.EncodeToString(nameHash))
		if err == nil {
			min, _ := strconv.Atoi(minlen)
			max, _ := strconv.Atoi(maxlen)
			number, _ := strconv.Atoi(pageNumber)
			size, _ := strconv.Atoi(pageSize)
			RootNameInfo := new(dns.RootNameInfo)
			json.Unmarshal([]byte(val), RootNameInfo)
			MapSubName := RootNameInfo.SubName
			if min == max && min != 0 {
				for _, data := range MapSubName[minlen] {
					resArr = append(resArr, data)
				}
			}
			if min == 0 && max == 0 {
				for _, value := range MapSubName {
					for _, data := range value {
						resArr = append(resArr, data)
					}
				}
			}
			if min != max && min < max {
				for i := min; i <= max; i++ {
					for _, data := range MapSubName[fmt.Sprint(i)] {
						resArr = append(resArr, data)
					}
				}
			}
			SubNameInfo := new(dns.SubNameInfo)
			subnameList = Paging(number, size, resArr)
			SubInfoL := []*QuerySub{}
			for _, sub := range subnameList {
				subnameHash := crypto.Keccak256([]byte(sub))
				subval, _ := db.GetSubName(hex.EncodeToString(subnameHash))
				json.Unmarshal([]byte(subval), SubNameInfo)
				SubInfoL = append(SubInfoL, &QuerySub{Name: SubNameInfo.Name, Owner: SubNameInfo.Owner, Expiration: SubNameInfo.ExpireTime})
			}
			res := &Res{Code: 1, Message: "ok", Total: len(resArr), PageNumber: number, PageSize: size, Data: SubInfoL}
			resbyte, _ := json.Marshal(res)
			msg = string(resbyte)
			writer.WriteHeader(200)
			writer.Write([]byte(msg))
			return
		} else {
			res_ := NotDataRes(fmt.Sprintf("%s not data", rootname))
			writer.WriteHeader(200)
			writer.Write([]byte(res_))
			return
		}
	} else {
		res_ := NotDataRes(fmt.Sprintf("%s:is not root", rootname))
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}

}

func (a *Api) AddrResolve(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}
	pageNumber := ""
	pageSize := ""
	addr := ""
	query := request.URL.Query()
	log.Println("/dns/addr query ", query)
	msg := "not found addr"
	var (
		v  []string
		ok bool
	)
	if v, ok = query["pageNumber"]; ok {
		pageNumber = v[0]
	}
	if v, ok = query["pageSize"]; ok {
		pageSize = v[0]
	}

	if v, ok = query["addr"]; ok {
		addr = v[0]
	} else {
		res_ := NotDataRes("not addr")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	db := ldb.GetLdb()
	val, err := db.GetAddressList(strings.ToLower(addr))
	log.Println("AddrResolve GetAddressList ", val)
	if err == nil {
		number, _ := strconv.Atoi(pageNumber)
		size, _ := strconv.Atoi(pageSize)
		var data []string
		hasharr := Paging(number, size, val)
		for _, v := range hasharr {
			in, _ := db.GetKey(v)
			data = append(data, in)
		}
		res := &Res{Code: 1, Message: "ok", Total: len(val), PageNumber: number, PageSize: size, Data: data}
		resbyte, _ := json.Marshal(res)
		msg = string(resbyte)
		writer.WriteHeader(200)
		writer.Write([]byte(msg))
		return
	} else {
		res_ := NotDataRes(fmt.Sprintf("not found %s data", addr))
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
}

// 获取钱包地址收入
func (a *Api) GetEarningsByAddress(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}
	addr := ""
	pageNumber := "0"
	pageSize := "100"
	query := request.URL.Query()
	log.Println("GetCashDetailsByAddress query ", query)
	var (
		v  []string
		ok bool
	)
	if v, ok = query["coinbase"]; ok {
		addr = v[0]
	} else {
		res_ := NotDataRes("not coinbase")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	if v, ok = query["pageNumber"]; ok {
		pageNumber = v[0]
	}
	if v, ok = query["pageSize"]; ok {
		pageSize = v[0]
	}
	cli, err := ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		cli.Close()
		log.Println("NewRootClient EthClient conn err", err)
	}
	defer cli.Close()
	dnsaccountant, err := dnsv2.NewDnsAccountantClient(cli)
	if err != nil {
		log.Println("NewDnsOwnerClient", err)
		return
	}
	db := ldb.GetLdb()
	val, err := db.GetAddressList(strings.ToLower(addr))
	if err == nil {
		number, _ := strconv.Atoi(pageNumber)
		size, _ := strconv.Atoi(pageSize)
		var rndata []string
		for _, name := range val {
			if strings.Contains(name, "rnH_1_") {
				rndata = append(rndata, name)
			}
		}
		var income []*GetEarningsByAddressIncome
		var items []*GetEarningsByAddressItems
		work := false
		rndatares := Paging(number, size, rndata)
		for _, name := range rndatares {
			in, _ := db.GetKey(name)
			root := new(dns.RootNameInfo)
			errun := json.Unmarshal([]byte(in), &root)
			if errun != nil {
				log.Println("GetCashDetailsByAddress Json.Unmarshal err", errun)
			} else {
				// root
				if root.Contract != common.HexToAddress("0x0000000000000000000000000000000000000000") {
					usdt, _ := dnsaccountant.Get(nil, root.Contract, common.HexToAddress(config.GetRConf().Cconf.Usdt))
					eth, _ := dnsaccountant.Get(nil, root.Contract, common.HexToAddress("0x0000000000000000000000000000000000000000"))
					income = []*GetEarningsByAddressIncome{{Erc20Addr: common.HexToAddress(config.GetRConf().Cconf.Usdt), WithDrawWei: usdt},
						{WithDrawWei: eth}}
					multi, _ := dnsaccountant.MultiSignerStore(nil, root.Contract)
					work = multi.Work
				}
				items = append(items, &GetEarningsByAddressItems{Name: root.Name, Owner: root.Owner, Erc721Addr: root.Owner, TokenId: root.TokenId, Work: work, Income: income})
			}
		}
		res := &Res{
			Code:    1,
			Message: "ok",
			Data:    &GetEarningsByAddressRes{PageNumber: number, PageSize: size, Items: items},
		}
		resbyte, _ := json.Marshal(res)
		msg := string(resbyte)
		writer.WriteHeader(200)
		writer.Write([]byte(msg))
		return

	} else {
		res_ := NotDataRes(fmt.Sprintf("not found %s data", addr))
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
}

func (a *Api) GetCashDetailsByAddress(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}
	addr := ""
	query := request.URL.Query()
	log.Println("GetCashDetailsByAddress query ", query)
	msg := "not found addr"
	var (
		v  []string
		ok bool
	)
	if v, ok = query["addr"]; ok {
		addr = v[0]
	} else {
		res_ := NotDataRes("not addr")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	db := ldb.GetLdb()
	val, err := db.GetAddressCash(strings.ToLower(addr))
	if err == nil {
		res := &Res{
			Code:    1,
			Message: "ok",
			Data:    val,
		}
		resbyte, _ := json.Marshal(res)
		msg = string(resbyte)
		writer.WriteHeader(200)
		writer.Write([]byte(msg))
		return

	} else {
		res_ := NotDataRes(fmt.Sprintf("not found %s data", addr))
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
}

// 获取didname多签列表
func (a *Api) GetSignTldListByDidName(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}
	didName := ""
	//pageNumber := "0"
	//pageSize := "100"
	query := request.URL.Query()
	log.Println("GetQuerySignTldList query ", query)
	// msg := "not found addr"
	var (
		v  []string
		ok bool
	)
	if v, ok = query["didName"]; ok {
		didName = v[0]
	} else {
		res_ := NotDataRes("not didName")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	//if v, ok = query["pageNumber"]; ok {
	//	pageNumber = v[0]
	//}
	//if v, ok = query["pageSize"]; ok {
	//	pageSize = v[0]
	//}
	db := ldb.GetLdb()
	cli, err := ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		cli.Close()
		log.Println("NewRootClient EthClient conn err", err)
	}
	defer cli.Close()
	dnsaccountant, err := dnsv2.NewDnsAccountantClient(cli)
	if err != nil {
		log.Println("NewDnsOwnerClient", err)
		return
	}
	data := []*Sign{}
	rootnamehash := hex.EncodeToString(crypto.Keccak256([]byte(didName)))
	rootinfo, _ := db.GetRootName(rootnamehash)
	Rootname := new(dnsv2.RootNameInfo)
	err = json.Unmarshal([]byte(rootinfo), Rootname)
	if err != nil {
		log.Println("json Unmarshal ", err)
	} else {
		if Rootname.Contract != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			multi, _ := dnsaccountant.MultiSignerStore(nil, Rootname.Contract)
			// signers, err1 := dnsaccountant.GetSignerSetAddress(nil, Rootname.Contract)
			topnamehashbyte := byte32(crypto.Keccak256([]byte(Rootname.Name)))
			maxsig, _, signers, _ := dnsaccountant.GetTaskInfo(nil, Rootname.Contract, *topnamehashbyte)
			// 取多签列表地址 20个字节一个地址
			countS := len(signers) / 20
			var signL []*GetSignTldListByDidNameSignList
			// var  items []*GetSignTldListByDidNameItems
			for i := 0; i < countS; i++ {
				a, b := i*20, (i+1)*20
				// common.BytesToAddress(signers[a:b])
				signL = append(signL, &GetSignTldListByDidNameSignList{
					Signer: common.BytesToAddress(signers[a:b]),
				})
			}
			items := []*GetSignTldListByDidNameItems{{NameHash: Rootname.Hash, Work: multi.Work, Lock: multi.Lock, MaxSig: maxsig, SignerList: signL}}
			res := &Res{
				Code:    1,
				Message: "ok",
				Data:    &GetSignTldListByDidNameRes{Items: items},
			}
			resbyte, _ := json.Marshal(res)
			msg := string(resbyte)
			writer.WriteHeader(200)
			writer.Write([]byte(msg))
			return
		}
	}

	if len(data) > 0 {
		res := &Res{
			Code:    1,
			Message: "ok",
			Data:    data,
		}
		resbyte, _ := json.Marshal(res)
		msg := string(resbyte)
		writer.WriteHeader(200)
		writer.Write([]byte(msg))
		return
	} else {
		res_ := NotDataRes(fmt.Sprintf("not found %s data", didName))
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
}

func (a *Api) ConfResolve(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}
	pageNumber := ""
	pageSize := ""
	conftype := ""
	confval := ""
	query := request.URL.Query()
	log.Println("ConfResolve query ", query)
	msg := "not found conftype confval"
	var (
		v  []string
		ok bool
	)
	if v, ok = query["pageNumber"]; ok {
		pageNumber = v[0]
	}
	if v, ok = query["pageSize"]; ok {
		pageSize = v[0]
	}

	if v, ok = query["conftype"]; ok {
		conftype = v[0]
	} else {
		res_ := NotDataRes("not conftype")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	if v, ok = query["confval"]; ok {
		confval = v[0]
	} else {
		res_ := NotDataRes("not a get confval")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	db := ldb.GetLdb()
	val, err := db.GetConfNameHashList(strings.ToLower(conftype) + "_" + confval)
	log.Println("GetConfNameHashList", val)
	if err == nil {
		number, _ := strconv.Atoi(pageNumber)
		size, _ := strconv.Atoi(pageSize)
		var data []string
		hasharr := Paging(number, size, val)
		for _, v := range hasharr {
			in, _ := db.GetKey(v)
			data = append(data, in)
		}
		res := &Res{Code: 1, Message: "ok", Total: len(val), PageNumber: number, PageSize: size, Data: data}
		resbyte, _ := json.Marshal(res)
		msg = string(resbyte)
		writer.WriteHeader(200)
		writer.Write([]byte(msg))
		return
	} else {
		res_ := NotDataRes(fmt.Sprintf("not found %s data", confval))
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}

}

func (a *Api) AddrDomainsResolve(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}
	pageNumber := ""
	pageSize := ""
	addrtype := ""
	// addr := ""
	query := request.URL.Query()
	log.Println("AddrDomainsResolve query ", query)

	msg := string("not found data")

	var (
		v  []string
		ok bool
	)
	if v, ok = query["pageNumber"]; ok {
		pageNumber = v[0]
	}
	if v, ok = query["pageSize"]; ok {
		pageSize = v[0]
	}

	if v, ok = query["addrtype"]; ok {
		addrtype = v[0]
	} else {
		res_ := NotDataRes("not addrtype")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	//if v, ok = query["addr"]; ok {
	//	addr = v[0]
	//} else {
	//	writer.WriteHeader(200)
	//	fmt.Fprintf(writer, "not a valid request")
	//	return
	//}
	db := ldb.GetLdb()
	val, err := db.GetConfNameHashList(strings.ToLower(addrtype))
	log.Println("GetConfNameHashList", val)
	if err == nil {
		number, _ := strconv.Atoi(pageNumber)
		size, _ := strconv.Atoi(pageSize)
		var data []*AddrDomains
		hasharr := Paging(number, size, val)
		for _, v := range hasharr {
			in, _ := db.GetKey(v)
			if strings.HasPrefix(v, "rnH_1_") {
				root := new(dns.RootNameInfo)
				errun := json.Unmarshal([]byte(in), root)
				if errun != nil {
					log.Println("AddrDomainsResolve Json.Unmarshal err", errun)
				} else {
					data = append(data, &AddrDomains{
						Name: root.Name, ConfType: addrtype, ConfValue: root.Conf[addrtype],
					})
				}

			} else {
				sub := new(dns.SubNameInfo)
				errun := json.Unmarshal([]byte(in), sub)
				if errun != nil {
					log.Println("AddrDomainsResolve Json.Unmarshal err", errun)
				} else {
					data = append(data, &AddrDomains{
						Name: sub.Name, ConfType: addrtype, ConfValue: sub.Conf[addrtype],
					})
				}
			}
		}
		res := &Res{Code: 1, Message: "ok", Total: len(val), PageNumber: number, PageSize: size, Data: data}
		resbyte, _ := json.Marshal(res)
		msg = string(resbyte)
		writer.WriteHeader(200)
		writer.Write([]byte(msg))
		return
	} else {
		res_ := NotDataRes(fmt.Sprintf("not found %s data", addrtype))
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}

}

func (a *Api) AddrTopList(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}
	pageNumber := "0"
	pageSize := "10"
	addr := ""
	query := request.URL.Query()
	log.Println("AddrTopList query ", query)
	msg := "not found coinbase"
	var (
		v  []string
		ok bool
	)
	if v, ok = query["pageNumber"]; ok {
		pageNumber = v[0]
	}
	if v, ok = query["pageSize"]; ok {
		pageSize = v[0]
	}

	if v, ok = query["coinbase"]; ok {
		addr = v[0]
	} else {
		res_ := NotDataRes("not coinbase")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	db := ldb.GetLdb()
	val, err := db.GetAddressList(strings.ToLower(addr))
	log.Println("AddrTopList GetAddressList ", val)
	if err == nil {
		number, _ := strconv.Atoi(pageNumber)
		size, _ := strconv.Atoi(pageSize)
		var rndata []string
		for _, name := range val {
			if strings.Contains(name, "rnH_1_") {
				rndata = append(rndata, name)
			}
		}
		rndatares := Paging(number, size, rndata)
		var data []*AddrTopListDataItems
		for _, name := range rndatares {
			in, _ := db.GetKey(name)
			root := new(dns.RootNameInfo)
			errun := json.Unmarshal([]byte(in), &root)
			if errun != nil {
				log.Println("AddrTopList Json.Unmarshal err", errun)
			} else {
				data = append(data, &AddrTopListDataItems{
					Name: root.Name, Erc721_Addr: root.Contract, TokenId: root.TokenId, OpenToReg: root.OpenToReg, ExpireTime: root.ExpireTime, Owner: root.Owner, PayTokens: []string{},
				})
			}
		}
		res := &Res{Code: 1, Message: "ok", Total: len(rndata), PageNumber: number, PageSize: size, Data: data}
		resbyte, _ := json.Marshal(res)
		msg = string(resbyte)
		writer.WriteHeader(200)
		writer.Write([]byte(msg))
		return
	} else {
		res_ := NotDataRes(fmt.Sprintf("not found %s data", addr))
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}

}
func (a *Api) AddrSubList(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}
	pageNumber := "0"
	pageSize := "10"
	addr := ""
	query := request.URL.Query()
	log.Println("AddrSubList query ", query)
	msg := "not found coinbase"
	var (
		v  []string
		ok bool
	)
	if v, ok = query["pageNumber"]; ok {
		pageNumber = v[0]
	}
	if v, ok = query["pageSize"]; ok {
		pageSize = v[0]
	}

	if v, ok = query["coinbase"]; ok {
		addr = v[0]
	} else {
		res_ := NotDataRes("not coinbase")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	db := ldb.GetLdb()
	val, err := db.GetAddressList(strings.ToLower(addr))
	log.Println("AddrSubList GetAddressList ", val)
	if err == nil {
		number, _ := strconv.Atoi(pageNumber)
		size, _ := strconv.Atoi(pageSize)
		var rndata []string
		var data []*AddrTopListDataItems
		for _, name := range val {
			if strings.Contains(name, "snH_1_") {
				rndata = append(rndata, name)
			}
		}
		rndatares := Paging(number, size, rndata)
		for _, name := range rndatares {
			in, _ := db.GetKey(name)
			sub := new(dns.SubNameInfo)
			errun := json.Unmarshal([]byte(in), &sub)
			if errun != nil {
				log.Println("AddrSubList Json.Unmarshal err", errun)
			} else {
				data = append(data, &AddrTopListDataItems{
					Name: sub.Name, Erc721_Addr: sub.Contract, TokenId: sub.TokenId, ExpireTime: sub.ExpireTime, Owner: sub.Owner, PayTokens: []string{},
				})
			}
		}
		res := &Res{Code: 1, Message: "ok", Total: len(rndata), PageNumber: number, PageSize: size, Data: data}
		resbyte, _ := json.Marshal(res)
		msg = string(resbyte)
		writer.WriteHeader(200)
		writer.Write([]byte(msg))
	} else {
		res_ := NotDataRes(fmt.Sprintf("not found %s data", addr))
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}

}

func (a *Api) GetOpenRegister(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}
	pageNumber := "0"
	pageSize := "10"
	query := request.URL.Query()
	log.Println("GetOpenRegister query ", query)
	msg := "not found data"
	var (
		v  []string
		ok bool
	)
	if v, ok = query["pageNumber"]; ok {
		pageNumber = v[0]
	}
	if v, ok = query["pageSize"]; ok {
		pageSize = v[0]
	}
	db := ldb.GetLdb()
	contractL, err := db.GetContractList()
	if err == nil {
		number, _ := strconv.Atoi(pageNumber)
		size, _ := strconv.Atoi(pageSize)
		contractlist := Paging(number, size, contractL)
		var data []*AddrTopListDataItems
		for _, contract := range contractlist {
			nameHash, _ := db.GetContractAddr(contract)
			rootname, _ := db.GetRootName(nameHash)
			rootinfo := new(dns.RootNameInfo)
			err = json.Unmarshal([]byte(rootname), rootinfo)
			if err != nil {
				log.Println("GetOpenRegister Json.Unmarshal err", err)
			} else {
				hasprice := false
				if rootinfo.Price != nil {
					hasprice = true
				}
				data = append(data, &AddrTopListDataItems{
					Name: rootinfo.Name, Erc721_Addr: rootinfo.Contract, TokenId: rootinfo.TokenId, ExpireTime: rootinfo.ExpireTime, OpenToReg: rootinfo.OpenToReg, HasPrice: hasprice, Owner: rootinfo.Owner, PayTokens: []string{},
				})
			}
		}
		res := &Res{Code: 1, Message: "ok", Total: len(contractL), PageNumber: number, PageSize: size, Data: data}
		resbyte, _ := json.Marshal(res)
		msg = string(resbyte)
		writer.WriteHeader(200)
		writer.Write([]byte(msg))
	} else {
		res_ := NotDataRes(fmt.Sprintf("not found data"))
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}

}

func (a *Api) GetMyPassCardList(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}
	pageNumber := "0"
	pageSize := "10"
	coinbase := ""
	query := request.URL.Query()
	log.Println("GetMyPassCardList query ", query)
	msg := "not found coinbase"
	var (
		v  []string
		ok bool
	)
	if v, ok = query["pageNumber"]; ok {
		pageNumber = v[0]
	}
	if v, ok = query["pageSize"]; ok {
		pageSize = v[0]
	}

	if v, ok = query["coinbase"]; ok {
		coinbase = v[0]
	} else {
		res_ := NotDataRes("not coinbase")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	db := ldb.GetLdb()
	val, err := db.GetNftPass(strings.ToLower(coinbase))
	log.Println("GetMyPassCardList GetNftPass ", val)
	if err == nil {
		number, _ := strconv.Atoi(pageNumber)
		size, _ := strconv.Atoi(pageSize)
		//var data []*ldb.NftPass
		var data []*ldb.NftPass
		if number == 0 && size == 0 {
			data = val
		}
		if number <= len(val) {
			if number+size > len(val) {
				data = val[number:]
			} else {
				data = val[number : size+number]
			}
		} else {
			data = nil
		}
		res := &Res{Code: 1, Message: "ok", Total: len(val), PageNumber: number, PageSize: size, Data: data}
		resbyte, _ := json.Marshal(res)
		msg = string(resbyte)
		writer.WriteHeader(200)
		writer.Write([]byte(msg))
		return
	} else {
		res_ := NotDataRes(fmt.Sprintf("not found %s data", coinbase))
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}

}

func (a *Api) PostSignMint(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}
	msg := "not found"
	decoder := json.NewDecoder(request.Body)
	var reqparams map[string]string
	decoder.Decode(&reqparams)
	log.Println("PostSignMint query", reqparams)
	name := reqparams["domain_name"]
	if name == "" {
		res_ := NotDataRes("not name")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	db := ldb.GetLdb()
	year := reqparams["year"]
	if year == "" {
		res_ := NotDataRes("not year")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	erc_20_addr := reqparams["erc_20_addr"]
	if erc_20_addr == "" {
		res_ := NotDataRes("not erc_20_addr")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	price := reqparams["price"]
	if price == "" {
		res_ := NotDataRes("not price")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	msg_sender := reqparams["msg_sender"]
	if msg_sender == "" {
		res_ := NotDataRes("not msg_sender")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	token_id := reqparams["token_id"]
	if token_id == "" {
		res_ := NotDataRes("not token_id")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	log.Println("PostSignMint query", reqparams, len(reqparams))
	// 保存name对应的pass卡信息
	callparms, _ := db.GetSignMintCallParams(name)
	if callparms == nil {
		callparms = &ldb.SignMintCallParams{
			TokenId: token_id, DomainsName: name, MsgSender: msg_sender,
		}
		err := db.SaveSignMintCallParams(name, callparms)
		if err != nil {
			log.Println("SaveSignMintCallParams err", err)
		} else {
			log.Println("SaveSignMintCallParams", callparms)
			//CallNftPass(token_id, name)
		}
	}
	if token_id != "9999" {
		PassList, _ := db.GetNftPass(strings.ToLower(msg_sender))
		if PassList == nil {
			res_ := NotDataRes(fmt.Sprintf("%s There is no card", msg_sender))
			writer.WriteHeader(200)
			writer.Write([]byte(res_))
			return
		} else {
			in := false
			for _, pass := range PassList {
				if pass.TokenId.String() == token_id {
					in = true
					if pass.CardColor == 3 {
						if !strings.HasSuffix(name, ".did") {
							res_ := NotDataRes("green card can only be mint .did")
							writer.WriteHeader(200)
							writer.Write([]byte(res_))
							return
						}
					} else {
						if strings.Contains(name, ".") {
							res_ := NotDataRes("Color and gold cards only mint top name")
							writer.WriteHeader(200)
							writer.Write([]byte(res_))
							return
						}
					}
				}
			}
			if !in {
				res_ := NotDataRes(fmt.Sprintf("%s not tokenid %s card", msg_sender, token_id))
				writer.WriteHeader(200)
				writer.Write([]byte(res_))
				return
			}
		}
	}

	result := encodePacked(
		encodeBytesString(hex.EncodeToString([]byte(name))),
		encodeUint256(year)[32-1:],
		encodeBytesString(erc_20_addr[2:]),
		encodeUint256(price),
		encodeBytesString(msg_sender[2:]),
		encodeUint256(token_id)[32-4:],
	)
	hash := crypto.Keccak256Hash(result)
	log.Println("Hash", hex.EncodeToString(hash[:]))
	signer, err := Signer(hash[:])
	if err != nil {
		log.Println("Sign err", err)
	}
	if signer[len(signer)-1] <= 1 {
		signer[len(signer)-1] = signer[len(signer)-1] + 27
	}
	year_int, _ := strconv.Atoi(year)
	token_id_int, _ := strconv.Atoi(token_id)
	// db := ldb.GetLdb()
	params := &PostSignMintParams{DomainName: name, Year: int32(year_int), Erc20Addr: common.HexToAddress(erc_20_addr), Price: price,
		MsgSender: common.HexToAddress(msg_sender), TokenId: int32(token_id_int)}
	data := &PostSignMintData{Signature: hex.EncodeToString(signer), Params: params}
	res := &Res{
		Code:    1,
		Message: "ok",
		Data:    data,
	}
	resbyte, _ := json.Marshal(res)
	if err == nil {
		msg = string(resbyte)
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

}
