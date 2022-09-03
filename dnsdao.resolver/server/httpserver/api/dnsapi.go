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
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
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
	if pageNumber <= len(resArr) {
		if pageNumber+pageSize > len(resArr) {
			res = resArr[pageNumber:]
		} else {
			res = resArr[pageNumber : pageSize+pageNumber]
		}
	} else {
		res = nil
	}
	return res
}
func (a *Api) Blockchain(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}

	query := request.URL.Query()
	msg := "not found blockchain"
	var (
		v  []string
		ok bool
	)

	if v, ok = query["blockchain"]; !ok {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
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
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
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
		fmt.Fprintf(writer, "not a valid request")
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
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}

	query := request.URL.Query()
	msg := "not found ipv6"
	var (
		v  []string
		ok bool
	)

	if v, ok = query["ipv6"]; !ok {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
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
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}

	query := request.URL.Query()
	log.Println("/dns/totalbysub query", query)
	var (
		v  []string
		ok bool
	)

	if v, ok = query["tldName"]; !ok {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
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
		msg = fmt.Sprint("%s:is not root", name)
	}

	writer.WriteHeader(200)
	writer.Write([]byte(msg))
}

func (a *Api) QuerySubDomainByPage(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
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
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "tldName is not a valid request")
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
		msg = fmt.Sprint("%s:is not root", rootname)
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
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
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
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
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
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
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
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
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
			Code:    200,
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
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
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
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
		return
	}
	db := ldb.GetLdb()
	val, err := db.GetRootEarnings(strings.ToLower(addr))
	// log.Println(val)
	if err == nil {
		res := &GetEarningsDetailsByAddress{Code: 200, Message: "ok", Data: val}
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
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
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
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
		return
	}
	db := ldb.GetLdb()
	val, err := db.GetAddressCash(strings.ToLower(addr))
	if err == nil {
		res := &GetCashDetailsByAddress{Code: 200, Message: "ok", Data: val}
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
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
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
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
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
	res := &QuerySignTldList{Code: 200, Message: "ok", Data: data}
	resbyte, errc := json.Marshal(res)
	if errc == nil {
		msg = string(resbyte)
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

}

func (a *Api) ConfResolve(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
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
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
		return
	}
	if v, ok = query["confval"]; ok {
		confval = v[0]
	} else {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
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
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}
	pageNumber := ""
	pageSize := ""
	addrtype := ""
	// addr := ""
	query := request.URL.Query()
	log.Println("AddrDomainsResolve query ", query)

	notfoundRes := &AddrDomainsResolveRes{Code: 200, Message: "not found", TotalCount: 0, Data: []*AddrDomains{}}
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
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
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
		res := &AddrDomainsResolveRes{Code: 200, Message: "ok", TotalCount: len(val), Data: data}
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
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
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
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
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
		addrtoplist := &AddrTopListRes{Code: 1, Message: "ok", Data: addrtoplistdata}
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
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
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
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
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
		addrtoplist := &AddrTopListRes{Code: 1, Message: "ok", Data: addrtoplistdata}
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
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
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
		addrtoplist := &AddrTopListRes{Code: 1, Message: "ok", Data: addrtoplistdata}
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
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
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
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "not a valid request")
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
		addrtoplistdata := &AddrTopListData{Total: len(data), PageNumber: number, PageSize: size, Items: data}
		addrtoplist := &AddrTopListRes{Code: 1, Message: "ok", Data: addrtoplistdata}
		resbyte, _ := json.Marshal(addrtoplist)
		if err == nil {
			msg = string(resbyte)
		}
	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))

}

func (a *Api) PostSignMint(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}
	//coinbase : current connected address . [required]
	//erc721Addr: Pass card contract address [required]
	//tokenId : Pass card ID [required]
	//domainhash: the did name sha256 [required]
	//years: [required]
	//erc20Addr: 支付token address [required]
	s, _ := ioutil.ReadAll(request.Body)
	writer.WriteHeader(200)
	writer.Write(s)

}
