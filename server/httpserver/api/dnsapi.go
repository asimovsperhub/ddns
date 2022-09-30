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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
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

// 获取二级域名每个长度的数量
func (a *Api) GetTotalBySubLen(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}
	decoder := json.NewDecoder(request.Body)
	var reqparams map[string]string
	decoder.Decode(&reqparams)
	log.Println("/dns/totalbysub query query", reqparams)
	name := reqparams["tldName"]
	if name == "" {
		res_ := NotDataRes("not tldName ")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	rules := reqparams["rules"]
	if rules == "" {
		res_ := NotDataRes("not rules")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	var ru []map[string]string
	json.Unmarshal([]byte(rules), &ru)

	ir := isRoot(name)
	nameHash := crypto.Keccak256([]byte(name))

	db := ldb.GetLdb()

	msg := "not found tldName"

	if ir {
		val, err := db.GetRootName(hex.EncodeToString(nameHash))
		if err == nil {
			RootNameInfo := new(dns.RootNameInfo)
			json.Unmarshal([]byte(val), RootNameInfo)
			var totalbysb []*TotalBySubRules
			all := int64(0)
			for _, total := range RootNameInfo.SubNameCount {
				all += total
			}
			if ru == nil {
				log.Println("ru", ru)
				for rule, count := range RootNameInfo.SubNameCount {
					ruleint, _ := strconv.Atoi(rule)
					totalbysb = append(totalbysb, &TotalBySubRules{Rule: rule, Min: int64(ruleint), Max: int64(ruleint + 1), Total: count})
				}
				res := &Res{
					Code:    1,
					Message: "ok",
					Data:    &TotalBySubLenResData{All: all, Rules: totalbysb},
				}
				resbyte, _ := json.Marshal(res)
				msg = string(resbyte)
				writer.WriteHeader(200)
				writer.Write([]byte(msg))
				return
			} else {
				log.Println(RootNameInfo.SubNameCount)
				for _, rulemap := range ru {
					min, _ := strconv.Atoi(rulemap["min"])
					max, _ := strconv.Atoi(rulemap["max"])
					if max > 8 {
						total := int64(0)
						for rule, count := range RootNameInfo.SubNameCount {
							ruleint, _ := strconv.Atoi(rule)
							if ruleint >= min {
								total += count
							}

						}
						totalbysb = append(totalbysb, &TotalBySubRules{Rule: rulemap["rule"], Min: int64(min), Max: int64(max), Total: total})
					} else {
						total := int64(0)
						for rule, count := range RootNameInfo.SubNameCount {
							ruleint, _ := strconv.Atoi(rule)
							if ruleint >= min && ruleint < max {
								total += count
							}

						}
						totalbysb = append(totalbysb, &TotalBySubRules{Rule: rulemap["rule"], Min: int64(min), Max: int64(max), Total: total})
					}
				}
				res := &Res{
					Code:    1,
					Message: "ok",
					Data:    &TotalBySubLenResData{All: all, Rules: totalbysb},
				}
				resbyte, _ := json.Marshal(res)
				msg = string(resbyte)
				writer.WriteHeader(200)
				writer.Write([]byte(msg))
				return
			}
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

// 获取指定长度的二级域名信息
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
			subnameList = Paging(number, size, resArr)
			SubInfoL := []*QuerySub{}
			for _, sub := range subnameList {
				SubNameInfo := new(dns.SubNameInfo)
				subnameHash := crypto.Keccak256([]byte(sub))
				subval, _ := db.GetSubName(hex.EncodeToString(subnameHash))
				json.Unmarshal([]byte(subval), SubNameInfo)
				SubInfoL = append(SubInfoL, &QuerySub{Name: SubNameInfo.Name, Owner: SubNameInfo.Owner, Expiration: SubNameInfo.ExpireTime})
			}
			res := &Res{Code: 1, Message: "ok", Data: &ResData{Total: len(resArr), PageNumber: number, PageSize: size, Items: SubInfoL}}
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
		var data []interface{}
		hasharr := Paging(number, size, val)
		for _, v := range hasharr {
			in, _ := db.GetKey(v)
			r := new(dnsv2.RootNameInfo)
			n := new(dnsv2.SubNameInfo)
			rr := json.Unmarshal([]byte(in), r)
			if rr == nil {
				data = append(data, r)
			} else {
				nn := json.Unmarshal([]byte(in), n)
				if nn == nil {
					data = append(data, n)
				}
			}

		}
		res := &Res{Code: 1, Message: "ok", Data: &ResData{Total: len(val), PageNumber: number, PageSize: size, Items: data}}
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
	log.Println("GetEarningsByAddress query ", query)
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
		var signL []*GetSignTldListByDidNameSignList
		work := false
		rndatares := Paging(number, size, rndata)
		for _, name := range rndatares {
			in, _ := db.GetKey(name)
			root := new(dns.RootNameInfo)
			errun := json.Unmarshal([]byte(in), &root)
			if errun != nil {
				log.Println("GetEarningsByAddress Json.Unmarshal err", errun)
			} else {
				if root.Contract != common.HexToAddress("0x0000000000000000000000000000000000000000") {
					usdt, _ := dnsaccountant.Get(nil, root.Contract, common.HexToAddress(config.GetRConf().Cconf.Usdt))
					eth, _ := dnsaccountant.Get(nil, root.Contract, common.HexToAddress("0x0000000000000000000000000000000000000000"))
					usdc, _ := dnsaccountant.Get(nil, root.Contract, common.HexToAddress("0x571787b5E033bF8bb2A1e5930265828ef3fFca00"))
					sol, _ := dnsaccountant.Get(nil, root.Contract, common.HexToAddress("0x65eEdD3b0B0A7c5cfBe97b22bBF19Cd492f2237E"))
					income = []*GetEarningsByAddressIncome{{Erc20Addr: common.HexToAddress(config.GetRConf().Cconf.Usdt), WithDrawWei: usdt, Symbol: "USDT", Decimals: 6},
						{Erc20Addr: common.HexToAddress("0x0000000000000000000000000000000000000000"), WithDrawWei: eth, Symbol: "ETH", Decimals: 18},
						{Erc20Addr: common.HexToAddress("0x571787b5E033bF8bb2A1e5930265828ef3fFca00"), WithDrawWei: usdc, Symbol: "USDC", Decimals: 18},
						{Erc20Addr: common.HexToAddress("0x65eEdD3b0B0A7c5cfBe97b22bBF19Cd492f2237E"), WithDrawWei: sol, Symbol: "SOL", Decimals: 6},
					}
					multi, _ := dnsaccountant.MultiSignerStore(nil, root.Contract)
					work = multi.Work
					topnamehashbyte := byte32(crypto.Keccak256([]byte(root.Name)))
					_, _, signers, _ := dnsaccountant.GetTaskInfo(nil, root.Contract, *topnamehashbyte)
					//taskhash, _ := dnsaccountant.GetAllTask(nil, root.Contract)
					// 取多签列表地址 20个字节一个地址
					countS := len(signers) / 20
					// var  items []*GetSignTldListByDidNameItems
					for i := 0; i < countS; i++ {
						a, b := i*20, (i+1)*20
						// common.BytesToAddress(signers[a:b])
						signL = append(signL, &GetSignTldListByDidNameSignList{
							Signer: common.BytesToAddress(signers[a:b]),
						})
					}
				}
				items = append(items, &GetEarningsByAddressItems{Name: root.Name, Owner: root.Owner, Erc721Addr: root.Contract, TokenId: root.TokenId, Work: work, Income: income, Signers: signL})
			}
		}
		res := &Res{
			Code:    1,
			Message: "ok",
			Data:    &ResData{PageNumber: number, PageSize: size, Items: items},
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
			Data:    &ResData{Items: val},
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
	rootnamehash := hex.EncodeToString(crypto.Keccak256([]byte(didName)))
	rootinfo, err := db.GetRootName(rootnamehash)
	Rootname := new(dnsv2.RootNameInfo)
	json.Unmarshal([]byte(rootinfo), Rootname)
	if err != nil {
		log.Println("json Unmarshal ", err)
		res_ := NotDataRes(fmt.Sprintf("not found %s MultiSign data", didName))
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	} else {
		// 开启了注册
		if Rootname.Contract != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			multi, _ := dnsaccountant.MultiSignerStore(nil, Rootname.Contract)
			// signers, err1 := dnsaccountant.GetSignerSetAddress(nil, Rootname.Contract)
			topnamehashbyte := byte32(crypto.Keccak256([]byte(Rootname.Name)))
			maxsig, _, signers, _ := dnsaccountant.GetTaskInfo(nil, Rootname.Contract, *topnamehashbyte)
			taskhash, _ := dnsaccountant.GetAllTask(nil, Rootname.Contract)
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
			items := []*GetSignTldListByDidNameItems{{NameHash: Rootname.Hash, TaskHash: taskhash, Work: multi.Work, Lock: multi.Lock, MaxSig: maxsig, SignerList: signL}}
			res := &Res{
				Code:    1,
				Message: "ok",
				Data:    &ResData{Items: items},
			}

			resbyte, _ := json.Marshal(res)
			msg := string(resbyte)
			writer.WriteHeader(200)
			writer.Write([]byte(msg))
			return
		} else {
			res_ := NotDataRes(fmt.Sprintf("The domain %s is not registered", didName))
			writer.WriteHeader(200)
			writer.Write([]byte(res_))
			return
		}
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
		res := &Res{Code: 1, Message: "ok", Data: &ResData{Total: len(val), PageNumber: number, PageSize: size, Items: data}}
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
		res := &Res{Code: 1, Message: "ok", Data: &ResData{Total: len(val), PageNumber: number, PageSize: size, Items: data}}
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
			log.Println("in", in)
			root := new(dns.RootNameInfo)
			errun := json.Unmarshal([]byte(in), &root)
			if errun != nil {
				log.Println("AddrTopList Json.Unmarshal err", errun)
			} else {
				data = append(data, &AddrTopListDataItems{
					Name: root.Name, Erc721Addr: root.Contract, TokenId: root.TokenId, ExpireTime: root.ExpireTime, Owner: root.Owner,
				})
			}
		}
		res := &Res{Code: 1, Message: "ok", Data: &ResData{Total: len(rndata), PageNumber: number, PageSize: size, Items: data}}
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
					Name: sub.Name, Erc721Addr: sub.Contract, TokenId: sub.TokenId, ExpireTime: sub.ExpireTime, Owner: sub.Owner,
				})
			}
		}
		res := &Res{Code: 1, Message: "ok", Data: &ResData{Total: len(rndata), PageNumber: number, PageSize: size, Items: data}}
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

// 开启注册的所有域名
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
		cost, _, _ := dnsv2.NewCostClient()
		var data []*AddrTopListDataItems
		for _, contract := range contractlist {
			nameHash, _ := db.GetContractAddr(contract)
			rootname, _ := db.GetRootName(nameHash)
			rootinfo := new(dns.RootNameInfo)
			err = json.Unmarshal([]byte(rootname), rootinfo)
			if err != nil {
				log.Println("GetOpenRegister Json.Unmarshal err", err)
			} else {
				log.Println("GetOpenRegister rootinfo", rootinfo)
				hasprice := false
				topnamehashbyte := byte32(crypto.Keccak256([]byte(rootinfo.Name)))
				cc, _ := cost.GetAllSecondLevelNamePrice(nil, *topnamehashbyte)
				for _, price := range cc {
					if price.String() != "0" {
						hasprice = true
						break
					}
				}
				log.Println("rootinfo.Name", rootinfo.Name)
				log.Println("GetAllSecondLevelNamePrice", cc)
				if hasprice {
					data = append(data, &AddrTopListDataItems{
						Name: rootinfo.Name, Erc721Addr: rootinfo.Contract, TokenId: rootinfo.TokenId, ExpireTime: rootinfo.ExpireTime, Owner: rootinfo.Owner, UsdtPrices: cc,
					})
				}
			}
		}
		if data != nil {
			res := &Res{Code: 1, Message: "ok", Data: &ResData{Total: len(data), PageNumber: number, PageSize: size, Items: data}}
			resbyte, _ := json.Marshal(res)
			msg = string(resbyte)
			writer.WriteHeader(200)
			writer.Write([]byte(msg))
			return
		} else {
			res_ := NotDataRes(fmt.Sprintf("No data or set prices"))
			writer.WriteHeader(200)
			writer.Write([]byte(res_))
			return
		}

	} else {
		res_ := NotDataRes(fmt.Sprintf("not found data"))
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}

}

// 获取钱包地址pass卡
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
		for _, pass := range data {
			// 验证pass卡是否使用
			//"msg_sender": strings.ToLower(pass.Owner.String()),
			filter := bson.M{"tokenid": pass.TokenId.String()}
			opm := &options.FindOptions{Skip: func(i int64) *int64 { return &i }(0), Limit: func(i int64) *int64 { return &i }(1), Sort: bson.M{"_id": -1}}
			selectres, _ := ldb.SELECT("signcoll", filter, opm)
			if selectres != nil {
				if selectres[0]["status"] == "2" {
					pass.Status = 2
				}
			}
			// 验证lock tokenid 是否在使用中
			locktk, _ := db.GetSignLock("token_id_tk" + pass.TokenId.String())
			if locktk != nil {
				if int32(time.Now().Unix())-locktk.LockTime < 600 {
					pass.Status = 1
				}
			}
		}
		res := &Res{Code: 1, Message: "ok", Data: &ResData{Total: len(val), PageNumber: number, PageSize: size, Items: data}}
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

// v2签名接口
func (a *Api) PostSignMint(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}
	decoder := json.NewDecoder(request.Body)
	var reqparams map[string]string
	decoder.Decode(&reqparams)
	log.Println("PostSignMint query", reqparams)
	name := reqparams["domain_name"]
	if name == "" {
		res_ := NotDataRes("not domain_name ")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
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
	//if token_id == "" {
	//	res_ := NotDataRes("not token_id")
	//	writer.WriteHeader(200)
	//	writer.Write([]byte(res_))
	//	return
	//}
	nonce := reqparams["nonce"]
	if nonce == "" {
		res_ := NotDataRes("not nonce")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	//lastPriceId
	lastPriceId := reqparams["lastPriceId"]
	if lastPriceId == "" {
		res_ := NotDataRes("not lastPriceId")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	log.Println("PostSignMint query", reqparams, len(reqparams))
	db := ldb.GetLdb()
	nameHash := crypto.Keccak256([]byte(name))
	namer, _ := db.GetRootName(hex.EncodeToString(nameHash))
	if namer != "" {
		res_ := NotDataRes(fmt.Sprintf("The domain %s has been registered", name))
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	names, _ := db.GetSubName(hex.EncodeToString(nameHash))
	if names != "" {
		res_ := NotDataRes(fmt.Sprintf("The domain %s has been registered", name))
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	// todo  验证token_id
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

	// 验证lock name
	lock, _ := db.GetSignLock(name)
	if lock != nil {
		if lock.MsgSender != strings.ToLower(msg_sender) {
			if int32(time.Now().Unix())-lock.LockTime < 600 {
				res_ := NotDataRes(fmt.Sprintf("Domain %s being registered", name))
				writer.WriteHeader(200)
				writer.Write([]byte(res_))
				return
			}
		}
		if lock.TokenId != token_id {
			if int32(time.Now().Unix())-lock.LockTime < 600 {
				res_ := NotDataRes(fmt.Sprintf("Domain %s being registered", name))
				writer.WriteHeader(200)
				writer.Write([]byte(res_))
				return
			}
		}

	}
	// 查最新的pass卡id 数据
	op := &options.FindOptions{Skip: func(i int64) *int64 { return &i }(0), Limit: func(i int64) *int64 { return &i }(1), Sort: bson.M{"_id": -1}}
	// find pass卡
	selectpass, _ := ldb.SELECT("signcoll", bson.M{"tokenid": token_id}, op)
	log.Println("find pass ", selectpass)
	if selectpass != nil {
		if selectpass[0]["status"] == "2" {
			res_ := NotDataRes(fmt.Sprintf("The %s pass card has been used", token_id))
			writer.WriteHeader(200)
			writer.Write([]byte(res_))
			return
		} else {
			// 验证lock tokenid
			locktk, _ := db.GetSignLock("token_id_tk" + token_id)
			if locktk != nil {
				// 这个假数据 没查pass所属，有pass检查话 不用这个验证
				if locktk.MsgSender != strings.ToLower(msg_sender) {
					if int32(time.Now().Unix())-locktk.LockTime < 600 {
						res_ := NotDataRes(fmt.Sprintf("The card %s is registering a domain name", token_id))
						writer.WriteHeader(200)
						writer.Write([]byte(res_))
						return
					}
				}
				if locktk.Name != name {
					if int32(time.Now().Unix())-locktk.LockTime < 600 {
						res_ := NotDataRes(fmt.Sprintf("The card %s is registering a domain name", token_id))
						writer.WriteHeader(200)
						writer.Write([]byte(res_))
						return
					}
				}
			}
		}
	}
	// save sign info
	filter := bson.M{"name": name}
	selectres, _ := ldb.SELECT("signcoll", filter, op)
	if selectres == nil {
		ldb.INSERTOne("signcoll", bson.M{"name": name, "msg_sender": strings.ToLower(msg_sender), "tokenid": token_id})
	} else {
		if selectres[0]["msg_sender"] == strings.ToLower(msg_sender) && selectres[0]["tokenid"] == token_id {
		} else {
			// 更新sender和token签名加锁10分钟如果还没监听到注册该域名 则更改name的注册信息(ps注册过的name的tokenid数据会有has_registered标记不会走到这)
			ldb.UPDATEOne("signcoll", filter, bson.M{"msg_sender": strings.ToLower(msg_sender), "tokenid": token_id})
		}
	}
	/*
		keccak256(abi.encodePacked(entireName_,
		                                                    year_,
		                                                    erc20Addr_,
		                                                    lastPriceId,
		                                                    nonce,
		                                                    price_,
		                                                    msg.sender)
	*/
	result := encodePacked(
		encodeBytesString(hex.EncodeToString([]byte(name))),
		encodeUint256(year)[32-1:],
		encodeBytesString(erc_20_addr[2:]),
		//uint80
		encodeUint256(lastPriceId)[32-10:],
		encodeUint256(nonce),
		encodeUint256(price),
		encodeBytesString(msg_sender[2:]),
		//encodeUint256(token_id)[32-4:],
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
	params := &PostSignMintParams{DomainName: name, Year: int32(year_int), Erc20Addr: common.HexToAddress(erc_20_addr), Price: price,
		MsgSender: common.HexToAddress(msg_sender), TokenId: int32(token_id_int), Nonce: nonce, LastPriceId: lastPriceId}
	data := &PostSignMintData{Signature: hex.EncodeToString(signer), Params: params}
	res := &Res{
		Code:    1,
		Message: "ok",
		Data:    &ResData{Items: data},
	}
	resbyte, _ := json.Marshal(res)
	msg := string(resbyte)
	// 加锁
	timeUnix := int32(time.Now().Unix())
	errsign := db.SaveSignLock(name, &ldb.SignLock{LockTime: timeUnix, Name: name, MsgSender: strings.ToLower(msg_sender), TokenId: token_id})
	if errsign != nil {
		log.Println("SaveSignLock ", name)
	}
	errsigntk := db.SaveSignLock("token_id_tk"+token_id, &ldb.SignLock{LockTime: timeUnix, Name: name, MsgSender: strings.ToLower(msg_sender), TokenId: token_id})
	if errsigntk != nil {
		log.Println("SaveSignLock tokenid", name)
	}
	// 加锁更新tokenid status
	// ldb.UPDATEOne("signcoll", bson.M{"tokenid": token_id}, bson.M{"status": "1"})
	writer.WriteHeader(200)
	writer.Write([]byte(msg))
	return

}
