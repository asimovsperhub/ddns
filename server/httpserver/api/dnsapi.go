package api

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/Agent/dns"
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
	"time"
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

type QuerySub struct {
	Name       string         `json:"name"`
	Owner      common.Address `json:"owner"`
	Expiration *big.Int       `json:"expiration"`
}
type QuerySubDomainByPage struct {
	TotalCount int         `json:"total_count"`
	Result     []*QuerySub `json:"result"`
}

type QueryName struct {
	TotalCount int      `json:"total_count"`
	Result     []string `json:"result"`
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
	//if pageNumber <= len(resArr) {
	//	if pageNumber+pageSize > len(resArr) {
	//		res = resArr[pageNumber:]
	//	} else {
	//		res = resArr[pageNumber : pageSize+pageNumber]
	//	}
	//} else {
	//	res = nil
	//}
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
			res, _ := json.Marshal(RootNameInfo.SubNameCount)
			msg = string(res)
		}
	} else {
		msg = fmt.Sprintf("%s:is not root", name)
	}

	writer.WriteHeader(200)
	writer.Write([]byte(msg))
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
			TotalCount := len(resArr)
			SubInfoL := []*QuerySub{}
			for _, sub := range subnameList {
				subnameHash := crypto.Keccak256([]byte(sub))
				subval, _ := db.GetSubName(hex.EncodeToString(subnameHash))
				json.Unmarshal([]byte(subval), SubNameInfo)
				SubInfoL = append(SubInfoL, &QuerySub{Name: SubNameInfo.Name, Owner: SubNameInfo.Owner, Expiration: SubNameInfo.ExpireTime})
			}

			res := &QuerySubDomainByPage{TotalCount: TotalCount, Result: SubInfoL}
			resbyte, _ := json.Marshal(res)
			msg = string(resbyte)
		}
	} else {
		msg = fmt.Sprintf("%s:is not root", rootname)
	}

	writer.WriteHeader(200)
	writer.Write([]byte(msg))

}

//func (a *Api) GetTotalByPrice(writer http.ResponseWriter, request *http.Request) {
//	if request.Method != "GET" {
//		writer.WriteHeader(500)
//		fmt.Fprintf(writer, "not a get request")
//		return
//	}
//
//	query := request.URL.Query()
//	log.Println("/dns/totalbyprice query", query)
//	var (
//		v  []string
//		ok bool
//	)
//
//	if v, ok = query["tldName"]; !ok {
//		writer.WriteHeader(200)
//		fmt.Fprintf(writer, "not a valid request")
//		return
//	}
//
//	name := v[0]
//
//	ir := isRoot(name)
//
//	nameHash := crypto.Keccak256([]byte(name))
//
//	db := ldb.GetLdb()
//
//	msg := "not found tldName"
//	countPrice := 0.0
//	if ir {
//		val, err := db.GetRootName(hex.EncodeToString(nameHash))
//		if err == nil {
//			RootNameInfo := new(dns.RootNameInfo)
//			json.Unmarshal([]byte(val), RootNameInfo)
//			//res, _ := json.Marshal(RootNameInfo.SubNameCount)
//			sn := RootNameInfo.SubName
//			pc := RootNameInfo.Price
//			for k, value := range sn {
//				if pc[k] != nil {
//					price, _ := strconv.Atoi(pc[k].String())
//					countPrice += (float64(price) / float64(1000000000000000000)) * float64(len(value))
//				} else {
//					price, _ := strconv.Atoi(pc["default"].String())
//					countPrice += (float64(price) / float64(1000000000000000000)) * float64(len(value))
//				}
//			}
//			res := &GetTotalByPrice{
//				TotalPrice: countPrice,
//			}
//			resbyte, _ := json.Marshal(res)
//			msg = string(resbyte)
//		}
//	} else {
//		msg = fmt.Sprint("%s:is not root", name)
//	}
//
//	writer.WriteHeader(200)
//	writer.Write([]byte(msg))
//}

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
		res := &QueryName{TotalCount: len(val), Result: data}
		resbyte, _ := json.Marshal(res)
		if err == nil {
			msg = string(resbyte)
		}
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

}

func (a *Api) GetEarningsByAddress(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}
	addr := ""
	query := request.URL.Query()
	log.Println("GetEarningsByAddress query ", query)
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
	val, err := db.GetRootEarnings(strings.ToLower(addr))
	if err == nil {
		cc := map[string]float64{}
		for k, v := range val.RootNameMap {
			cc[k] = v.Earnings
		}
		res := &GetEarningsByAddressRes{
			Code:    1,
			Message: "ok",
			Data: &GetEarningsByAddress{
				IncomeSummary: val.Earnings,
				Details:       cc,
			},
		}
		resbyte, errc := json.Marshal(res)
		if errc == nil {
			msg = string(resbyte)
		}
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

}

func (a *Api) GetEarningsDetailsByAddress(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}
	addr := ""
	query := request.URL.Query()
	log.Println("GetEarningsDetailsByAddress query ", query)
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
	val, err := db.GetRootEarnings(strings.ToLower(addr))
	// log.Println(val)
	if err == nil {
		res := &GetEarningsDetailsByAddress{Code: 1, Message: "ok", Data: val}
		resbyte, errc := json.Marshal(res)
		if errc == nil {
			msg = string(resbyte)
		}
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

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
		res := &GetCashDetailsByAddress{Code: 1, Message: "ok", Data: val}
		resbyte, errc := json.Marshal(res)
		if errc == nil {
			msg = string(resbyte)
		}
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

}

func (a *Api) GetQuerySignTldList(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}
	addr := ""
	query := request.URL.Query()
	log.Println("GetQuerySignTldList query ", query)
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
	cli, err := ethclient.Dial(config.GetRConf().GetRPCEndPoint())
	if err != nil {
		cli.Close()
		log.Println("NewRootClient EthClient conn err", err)
	}
	defer cli.Close()
	dnsaccountant, err := dns.NewDnsAccountantClient(cli)
	if err != nil {
		log.Println("NewDnsOwnerClient", err)
		return
	}
	Id_ := big.NewInt(int64(0))
	length, count, list, err := dnsaccountant.GetContractList(nil, common.HexToAddress(addr), Id_)
	if err != nil {
		log.Println("dnsaccountant.GetContractList", err)
	}
	log.Println("GetContractList ", length, count, list)
	data := []*Sign{}
	for _, contract := range list {
		if contract != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			multi, _ := dnsaccountant.MultiSignerStore(nil, contract)
			signers, err1 := dnsaccountant.GetSignerSetAddress(nil, contract)
			if err1 != nil {
				log.Println("dnsaccountant.GetSignerSetAddress err", err1)
			}
			countS := len(signers) / 20
			signL := []common.Address{}
			for i := 0; i < countS; i++ {
				a, b := i*20, (i+1)*20
				signL = append(signL, common.BytesToAddress(signers[a:b]))
			}
			nameHash, _ := db.GetContractAddr(contract.String())
			rootname, _ := db.GetRootName(nameHash)
			rootinfo := new(dns.RootNameInfo)
			err = json.Unmarshal([]byte(rootname), rootinfo)
			if err != nil {
				log.Println("json Unmarshal ", err)
				continue
			} else {
				var Earnings float64
				owner := rootinfo.Owner
				AddressEarnings, err2 := db.GetRootEarnings(strings.ToLower(owner.String()))
				if err2 != nil {
					log.Println("GetRootEarnings err", err2)
				} else {
					RootEarnings := AddressEarnings.RootNameMap[rootinfo.Name]
					Earnings = RootEarnings.Earnings
				}
				data = append(data, &Sign{Name: rootinfo.Name, Work: multi.Work, Erc721Addr: contract.String(), Erc20Addr: "0x0000000000000000000000000000000000000000",
					Income: Earnings, Singners: signL, Owner: rootinfo.Owner, TokenId: rootinfo.TokenId})
			}
		}
	}
	res := &QuerySignTldList{Code: 1, Message: "ok", Data: data}
	resbyte, errc := json.Marshal(res)
	if errc == nil {
		msg = string(resbyte)
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

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
		res := &QueryName{TotalCount: len(val), Result: data}
		resbyte, _ := json.Marshal(res)
		if err == nil {
			msg = string(resbyte)
		}
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

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

	notfoundRes := &AddrDomainsResolveRes{Code: 1, Message: "not found", TotalCount: 0, Data: []*AddrDomains{}}
	notfoundresbyte, _ := json.Marshal(notfoundRes)
	msg := string(notfoundresbyte)

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
		res := &AddrDomainsResolveRes{Code: 1, Message: "ok", TotalCount: len(val), Data: data}
		resbyte, _ := json.Marshal(res)
		if err == nil {
			msg = string(resbyte)
		}
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

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
		var data []*AddrTopListDataItems
		for _, name := range val {
			if strings.Contains(name, "rnH_1_") {
				rndata = append(rndata, name)
			}
		}
		rndatares := Paging(number, size, rndata)
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
		addrtoplistdata := &AddrTopListData{Total: len(rndata), PageNumber: number, PageSize: size, Items: data}
		addrtoplist := &Res{Code: 1, Message: "ok", Data: addrtoplistdata}
		resbyte, _ := json.Marshal(addrtoplist)
		if err == nil {
			msg = string(resbyte)
		}
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

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
		addrtoplistdata := &AddrTopListData{Total: len(rndata), PageNumber: number, PageSize: size, Items: data}
		addrtoplist := &Res{Code: 1, Message: "ok", Data: addrtoplistdata}
		resbyte, _ := json.Marshal(addrtoplist)
		if err == nil {
			msg = string(resbyte)
		}
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

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
				data = append(data, &AddrTopListDataItems{
					Name: rootinfo.Name, Erc721_Addr: rootinfo.Contract, TokenId: rootinfo.TokenId, ExpireTime: rootinfo.ExpireTime, OpenToReg: rootinfo.OpenToReg, Owner: rootinfo.Owner, PayTokens: []string{},
				})
			}
		}
		addrtoplistdata := &AddrTopListData{Total: len(contractL), PageNumber: number, PageSize: size, Items: data}
		addrtoplist := &Res{Code: 1, Message: "ok", Data: addrtoplistdata}
		resbyte, _ := json.Marshal(addrtoplist)
		if err == nil {
			msg = string(resbyte)
		}
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

}

func (a *Api) GetMyPassCardList(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}
	pageNumber := "0"
	pageSize := "2000"
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
		// 遍历data 查是否使用
		rootC, cli, _ := dns.NewRootClient()
		subC, cli, _ := dns.NewSubClient("0x111cFbc15FC56AD0aC8C7536BcA3347a558a0012")
		defer cli.Close()
		for _, tkid := range data {
			if tkid.RemainingTimes != 1 {
				u64tk, _ := strconv.ParseUint(tkid.TokenId.String(), 10, 32)
				isusedr, _ := rootC.PassCardUsed(nil, uint32(u64tk))
				isuseds, _ := subC.PassCardUsed(nil, uint32(u64tk))
				if isusedr {
					tkid.RemainingTimes = 1
					if addrpass[tkid.TokenId.String()] != "" {
						tklen := len(tkid.TokenId.String())
						tkid.RemainingTimes = 0
						tkid.ExtTokenId = "10000"[:5-tklen] + tkid.TokenId.String()
						u64tk2, _ := strconv.ParseUint(tkid.ExtTokenId, 10, 32)
						isusedr2, _ := rootC.PassCardUsed(nil, uint32(u64tk2))
						isuseds2, _ := subC.PassCardUsed(nil, uint32(u64tk2))
						if isusedr2 {
							tkid.RemainingTimes = 1
						}
						if isuseds2 {
							tkid.RemainingTimes = 1
						}

					}

				}
				if isuseds {
					tkid.RemainingTimes = 1
					if addrpass[tkid.TokenId.String()] != "" {
						tklen := len(tkid.TokenId.String())
						tkid.RemainingTimes = 0
						tkid.ExtTokenId = "10000"[:5-tklen] + tkid.TokenId.String()
						u64tk2, _ := strconv.ParseUint(tkid.ExtTokenId, 10, 32)
						isusedr2, _ := rootC.PassCardUsed(nil, uint32(u64tk2))
						isuseds2, _ := subC.PassCardUsed(nil, uint32(u64tk2))
						if isusedr2 {
							tkid.RemainingTimes = 1
						}
						if isuseds2 {
							tkid.RemainingTimes = 1
						}
					}
				}
			}
		}
		db.SaveNftPass(strings.ToLower(coinbase), data)
		addrtoplistdata := &AddrTopListData{Total: len(data), PageNumber: number, PageSize: size, Items: data}
		addrtoplist := &Res{Code: 1, Message: "ok", Data: addrtoplistdata}
		resbyte, _ := json.Marshal(addrtoplist)
		if err == nil {
			msg = string(resbyte)
			writer.WriteHeader(200)
			writer.Write([]byte(msg))
			return
		}
	} else {
		res_ := NotDataRes(msg)
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
	db := ldb.GetLdb()
	opensign, _ := db.GetKey("OpenSign")
	if opensign != "1" {
		res_ := NotDataRes("The signature mint service has not been started")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	msg := "not found"
	decoder := json.NewDecoder(request.Body)
	var reqparams map[string]string
	decoder.Decode(&reqparams)
	name := reqparams["domain_name"]
	if name == "" {
		res_ := NotDataRes("not name")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	year := reqparams["year"]
	if name == "" {
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
	// todo 取消白名单限制
	//for _, v := range domainslist {
	//	if name == v {
	//		res_ := NotDataRes("Domain name cannot be registered")
	//		writer.WriteHeader(200)
	//		writer.Write([]byte(res_))
	//		return
	//	}
	//}
	//rootC, cli, _ := dns.NewRootClient()
	//defer cli.Close()
	//u64tk, _ := strconv.ParseUint(token_id, 10, 32)
	//isusedr, _ := rootC.PassCardUsed(nil, uint32(u64tk))
	//subC, _, _ := dns.NewSubClient("0x111cFbc15FC56AD0aC8C7536BcA3347a558a0012")
	//isuseds, _ := subC.PassCardUsed(nil, uint32(u64tk))
	//if isusedr {
	//	res_ := NotDataRes(fmt.Sprintf("The cardid %s has already been used", token_id))
	//	writer.WriteHeader(200)
	//	writer.Write([]byte(res_))
	//	return
	//}
	//if isuseds {
	//	res_ := NotDataRes(fmt.Sprintf("The cardid %s has already been used", token_id))
	//	writer.WriteHeader(200)
	//	writer.Write([]byte(res_))
	//	return
	//}
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
	lock, _ := db.GetSignLock(name)
	if lock != nil {
		if lock.MsgSender != strings.ToLower(msg_sender) {
			if int32(time.Now().Unix())-lock.LockTime < 300 {
				res_ := NotDataRes(fmt.Sprintf("Domain %s being registered", name))
				writer.WriteHeader(200)
				writer.Write([]byte(res_))
				return
			}
		}
		if lock.TokenId != token_id {
			if int32(time.Now().Unix())-lock.LockTime < 300 {
				res_ := NotDataRes(fmt.Sprintf("Domain %s being registered", name))
				writer.WriteHeader(200)
				writer.Write([]byte(res_))
				return
			}
		}
	}
	cctoken := ""
	if token_id != "9999" {
		PassList, _ := db.GetNftPass(strings.ToLower(msg_sender))
		if PassList == nil {
			res_ := NotDataRes(fmt.Sprintf("%s There is no card", msg_sender))
			writer.WriteHeader(200)
			writer.Write([]byte(res_))
			return
		} else {
			in := false
			if len(token_id) == 5 {
				tk, _ := strconv.Atoi(token_id)
				cctoken, token_id = token_id, strconv.Itoa(tk-10000)
			}
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
	if cctoken != "" {
		token_id = cctoken
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
		// todo 加锁
		timeUnix := int32(time.Now().Unix())
		errsign := db.SaveSignLock(name, &ldb.SignLock{LockTime: timeUnix, Name: name, MsgSender: strings.ToLower(msg_sender), TokenId: token_id})
		if errsign != nil {
			log.Println("SaveSignLock ", name)
		}
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

}

// todo mint 开关暂留
func (a *Api) SwitchSignMint(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		res_ := NotDataRes("not a get request")
		writer.WriteHeader(500)
		writer.Write([]byte(res_))
		return
	}
	db := ldb.GetLdb()
	decoder := json.NewDecoder(request.Body)
	var reqparams map[string]string
	decoder.Decode(&reqparams)
	opensign := reqparams["opensign"]
	if opensign == "" {
		res_ := NotDataRes("not opensign")
		writer.WriteHeader(200)
		writer.Write([]byte(res_))
		return
	}
	log.Println("SwitchSignMint query", reqparams, len(reqparams))
	if opensign == "asimov" {
		err := db.SaveKey("OpenSign", "1")
		if err == nil {
			writer.WriteHeader(200)
			writer.Write([]byte("open success"))
			return
		} else {
			writer.WriteHeader(200)
			writer.Write([]byte("open save err"))
			return
		}
	} else {
		err := db.SaveKey("OpenSign", "0")
		if err == nil {
			writer.WriteHeader(200)
			writer.Write([]byte("shutdown success"))
			return
		} else {
			writer.WriteHeader(200)
			writer.Write([]byte("shutdown save err"))
			return
		}
	}

}
