package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/dnsDaoExchange/ethacct"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kprc/nbsnetwork/tools"
	"github.com/spf13/cobra"
	"log"
	"math"
	"math/big"
)

var ReceiptResult = [2]string{"Failed", "success"}

const contract_conf = "./contract.create"
const Version = "0.0.1"

var param struct {
	version   bool
	force     bool
	addr      string
	addr2     string
	decimal   uint8
	price     uint64
	name      string
	year      uint8
	hash      string
	lastPrice string
	nonce     string
	sigPrice  string
}

type CheckSumContract struct {
	LibSig          string `json:"lib_sig"`
	LibToolKit      string `json:"lib_tool_kit"`
	LibErc721       string `json:"lib_erc_721"`
	LibAddressArray string `json:"lib_address_array"`
	LibBytes32      string `json:"lib_bytes_32"`
	LibMultiSig     string `json:"lib_multi_sig"`
	LibAccountant   string `json:"lib_accountant"`
	TopName         string `json:"top_name"`
	Cost            string `json:"cost"`
	Accountant      string `json:"accountant"`
	SecondName      string `json:"second_name"`
	UsdtAddr        string `json:"usdt_addr"`
}

func checksumContract(c *Contract) *CheckSumContract {
	return &CheckSumContract{
		LibSig:          c.LibSig.String(),
		LibToolKit:      c.LibToolKit.String(),
		LibErc721:       c.LibErc721.String(),
		LibAddressArray: c.LibAddressArray.String(),
		LibBytes32:      c.LibBytes32.String(),
		LibMultiSig:     c.LibMultiSig.String(),
		LibAccountant:   c.LibAccountant.String(),
		TopName:         c.TopName.String(),
		Cost:            c.Cost.String(),
		Accountant:      c.Accountant.String(),
		SecondName:      c.SecondName.String(),
		UsdtAddr:        c.UsdtAddr.String(),
	}
}

func savefile(c *Contract) {

	if data, err := json.Marshal(*checksumContract(c)); err != nil {
		log.Println(err.Error())
	} else {
		tools.Save2File(data, contract_conf)
	}
}

func init() {

	flags := rootCmd.Flags()

	flags.BoolVarP(&param.version, "version",
		"v", false, "fcl -v")

	dnsTopCmd.Flags().BoolVarP(&param.force, "force", "f", false, "force to redeploy")
	dnsSecondCmd.Flags().BoolVarP(&param.force, "force", "f", false, "force to redeploy")
	dnsCostCmd.Flags().BoolVarP(&param.force, "force", "f", false, "force to redeploy")
	dnsAccountantCmd.Flags().BoolVarP(&param.force, "force", "f", false, "force to redeploy")
	dnsTopSetContractCmd.Flags().Uint8VarP(&param.decimal, "switch", "s", 0, "set switch")
	costSetUsdtCmd.Flags().StringVarP(&param.addr, "address", "a", "", "usdt address")
	costSetUsdtCmd.Flags().Uint8VarP(&param.decimal, "decimal", "d", 6, "usdt decimal")
	costAddPaymentCmd.Flags().StringVarP(&param.addr, "address", "a", "", "erc20 address")
	costAddPaymentCmd.Flags().StringVarP(&param.addr2, "chaddr", "c", "", "chain link contract address")
	costAddPaymentCmd.Flags().Uint8VarP(&param.decimal, "decimal", "d", 0, "erc20 decimal")
	costDisablePaymentCmd.Flags().StringVarP(&param.addr, "address", "a", "", "erc20 address")
	costDisablePaymentCmd.Flags().Uint8VarP(&param.decimal, "decimal", "d", 0, "erc20 decimal")
	costSetTopPriceCmd.Flags().Uint64VarP(&param.price, "price", "p", 0, "top level name price")
	costSetTopPriceCmd.Flags().Uint8VarP(&param.decimal, "decimal", "d", 0, "price of erc20 decimal")
	costSetSecondTaxCmd.Flags().Uint64VarP(&param.price, "tax", "p", 0, "second level name tax")
	costSetSecondTaxCmd.Flags().Uint8VarP(&param.decimal, "decimal", "d", 0, "tax of erc20 decimal")
	dnsTopMintNameCmd.Flags().StringVarP(&param.name, "name", "n", "", "entire name")
	dnsTopMintNameCmd.Flags().StringVarP(&param.addr, "address", "a", "", "erc20 address")
	dnsTopMintNameCmd.Flags().Uint8VarP(&param.year, "year", "y", 1, "year")
	dnsTopErc721AddrCmd.Flags().StringVarP(&param.hash, "hash", "c", "", "top level name hash")
	dnsTopErc721AddrCmd.Flags().StringVarP(&param.name, "name", "n", "", "top level name")
	dnsTopNameInfoCmd.Flags().StringVarP(&param.hash, "hash", "c", "", "top level name hash")
	dnsTopNameInfoCmd.Flags().StringVarP(&param.name, "name", "n", "", "top level name")
	dnsTopOpen2RegCmd.Flags().StringVarP(&param.hash, "hash", "c", "", "top level name hash")
	dnsTopOpen2RegCmd.Flags().StringVarP(&param.name, "name", "n", "", "top level name")
	costSetSecondPriceCmd.Flags().StringVarP(&param.hash, "hash", "c", "", "second level name hash")
	costSetSecondPriceCmd.Flags().StringVarP(&param.name, "name", "n", "", "second level name")
	costSetSecondPriceCmd.Flags().Uint64VarP(&param.price, "price", "p", 0, "second level name price")
	costSetSecondPriceCmd.Flags().Uint8VarP(&param.decimal, "decimal", "d", 0, "second of erc20 decimal")
	costShowPriceCmd.Flags().StringVarP(&param.hash, "hash", "c", "", "second level name hash")
	costShowPriceCmd.Flags().StringVarP(&param.name, "name", "n", "", "second level name")

	dnsSecondMintNameCmd.Flags().StringVarP(&param.name, "name", "n", "", "entire name")
	dnsSecondMintNameCmd.Flags().StringVarP(&param.addr, "address", "a", "", "erc20 address")
	dnsSecondMintNameCmd.Flags().Uint8VarP(&param.year, "year", "y", 1, "year")

	makeSignCmd.Flags().StringVarP(&param.name, "name", "n", "", "name to register")
	makeSignCmd.Flags().Uint8VarP(&param.year, "year", "y", 1, "register year")
	makeSignCmd.Flags().StringVarP(&param.addr, "erc20Addr", "e", "", "erc20 address")
	makeSignCmd.Flags().StringVarP(&param.addr2, "user", "u", "", "user address")
	makeSignCmd.Flags().StringVarP(&param.lastPrice, "lastPrice", "l", "", "last price id from chainlink")
	makeSignCmd.Flags().StringVarP(&param.nonce, "nonce", "o", "", "nonce")
	makeSignCmd.Flags().StringVarP(&param.sigPrice, "price", "p", "", "price")

	dnsTopMintNameBySigCmd.Flags().StringVarP(&param.name, "name", "n", "", "name to register")
	dnsTopMintNameBySigCmd.Flags().Uint8VarP(&param.year, "year", "y", 1, "register year")
	dnsTopMintNameBySigCmd.Flags().StringVarP(&param.addr, "erc20Addr", "e", "", "erc20 address")
	//dnsTopMintNameBySigCmd.Flags().StringVarP(&param.addr2, "user", "u", "", "user address")
	//dnsTopMintNameBySigCmd.Flags().StringVarP(&param.lastPrice, "lastPrice", "l", "", "last price id from chainlink")
	//dnsTopMintNameBySigCmd.Flags().StringVarP(&param.nonce, "nonce", "o", "", "nonce")
	dnsTopMintNameBySigCmd.Flags().StringVarP(&param.sigPrice, "price", "p", "", "price")

	rootCmd.AddCommand(dnsTopCmd)
	rootCmd.AddCommand(dnsSecondCmd)
	rootCmd.AddCommand(dnsCostCmd)
	rootCmd.AddCommand(dnsAccountantCmd)
	dnsTopCmd.AddCommand(dnsTopSetContractCmd)
	dnsTopCmd.AddCommand(dnsTopOwnerCmd)
	dnsTopCmd.AddCommand(dnsTopMintNameCmd)
	dnsTopCmd.AddCommand(dnsTopErc721AddrCmd)
	dnsTopCmd.AddCommand(dnsTopNameInfoCmd)
	dnsTopCmd.AddCommand(dnsTopOpen2RegCmd)

	dnsSecondCmd.AddCommand(dnsSecondSetContractCmd)
	dnsSecondCmd.AddCommand(dnsSecondMintNameCmd)

	dnsCostCmd.AddCommand(costAddPaymentCmd)
	dnsCostCmd.AddCommand(costDisablePaymentCmd)
	dnsCostCmd.AddCommand(costSetTopPriceCmd)
	dnsCostCmd.AddCommand(costSetSecondTaxCmd)
	dnsCostCmd.AddCommand(costSetUsdtCmd)
	dnsCostCmd.AddCommand(costSetSecondPriceCmd)
	dnsCostCmd.AddCommand(costShowPriceCmd)
	dnsCostCmd.AddCommand(costShowTaxCmd)
	rootCmd.AddCommand(makeSignCmd)

}

func mainRun(_ *cobra.Command, _ []string) {

	if param.version {
		fmt.Println(Version)
		return
	}

}

func deployDnsTop(_ *cobra.Command, _ []string) {
	c := loadConf()
	if param.force || AddrIsZero(c.TopName) {
		if err := c.DeployDnsTopLevelName(loadAcct().PrivKey); err != nil {
			log.Println(err.Error())
			savefile(c)
			return
		}
	}
	savefile(c)
	fmt.Println(c.String())
}

func deployDnsSecond(_ *cobra.Command, _ []string) {
	c := loadConf()
	if param.force || AddrIsZero(c.SecondName) {
		if err := c.DeploySecondContract(loadAcct().PrivKey); err != nil {
			log.Println(err.Error())
			savefile(c)
			return
		}
	}
	savefile(c)
	fmt.Println(c.String())
}

func deployDnsCost(_ *cobra.Command, _ []string) {
	c := loadConf()
	if param.force || AddrIsZero(c.Cost) {
		if err := c.DeployCost(loadAcct().PrivKey); err != nil {
			log.Println(err.Error())
			savefile(c)
			return
		}
	}
	savefile(c)
	fmt.Println(c.String())
}

func deployDnsAccountant(_ *cobra.Command, _ []string) {
	c := loadConf()
	if param.force || AddrIsZero(c.Accountant) {
		if err := c.DeployDnsAccountant(loadAcct().PrivKey); err != nil {
			log.Println(err.Error())
			savefile(c)
			return
		}
	}
	savefile(c)
	fmt.Println(c.String())
}

func dnsTopSetContract(_ *cobra.Command, _ []string) {
	c := loadConf()
	if AddrIsZero(c.TopName) || AddrIsZero(c.SecondName) || AddrIsZero(c.Cost) || AddrIsZero(c.Accountant) {
		fmt.Println("please deploy contract first")
		return
	}

	var (
		r   *types.Receipt
		err error
	)
	if r, err = c.SetDnsTopContract(loadAcct().PrivKey, param.decimal); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("top level set contract result:", ReceiptResult[r.Status], "tx:", r.TxHash.String())
}

func IsRootName(name string) bool {
	if len(name) > 64 || len(name) == 0 {
		return false
	}

	if name[0] == 45 || name[len(name)-1] == 45 {
		return false
	}
	//45 ==> -
	//46 ==> .
	for i := 0; i < len(name); i++ {
		if !((name[i] >= 48 && name[i] <= 57) || name[i] == 45 || (name[i] >= 97 && name[i] <= 122)) {
			return false
		}
	}

	return true
}

func IsNormalName(name string) bool {
	if len(name) > 192 || len(name) <= 2 {
		return false
	}

	if name[0] == 45 || name[0] == 46 {
		return false
	}

	if name[len(name)-1] == 45 || name[len(name)-1] == 46 {
		return false
	}

	//45 ==> -
	//46 ==> .
	var pre bool
	for i := 0; i < len(name); i++ {
		if !((name[i] >= 48 && name[i] <= 57) || name[i] == 46 || name[i] == 45 || (name[i] >= 97 && name[i] <= 122)) {
			return false
		}
		if name[i] == 46 && pre {
			return false
		} else if name[i] == 46 && !pre {
			pre = true
		} else {
			pre = false
		}
	}

	return true
}

func dnsTopMintName(_ *cobra.Command, _ []string) {
	c := loadConf()
	if AddrIsZero(c.TopName) {
		fmt.Println("please deploy contract first")
		return
	}

	if AddrIsZero(c.UsdtAddr) {
		fmt.Println("please set usdt address")
		return
	}

	if !IsRootName(param.name) {
		fmt.Println("not a valid top level name")
		return
	}

	if param.year <= 0 || param.year > 10 {
		fmt.Println("not a valid year")
		return
	}

	if b, err := c.NameIsExists(param.name); err != nil {
		fmt.Println(err)
		return
	} else if b {
		fmt.Println("name exists")
		return
	}

	var (
		r   *types.Receipt
		err error
	)
	if r, err = c.MintName(loadAcct().PrivKey, param.name, param.year, common.HexToAddress(param.addr)); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("top level set contract result:", ReceiptResult[r.Status], "tx:", r.TxHash.String())
}

func dnsTopMintNameBySig(_ *cobra.Command, _ []string) {
	c := loadConf()
	if AddrIsZero(c.TopName) {
		fmt.Println("please deploy contract first")
		return
	}

	if AddrIsZero(c.UsdtAddr) {
		fmt.Println("please set usdt address")
		return
	}

	if !IsRootName(param.name) {
		fmt.Println("not a valid top level name")
		return
	}

	if param.year <= 0 || param.year > 10 {
		fmt.Println("not a valid year")
		return
	}

	if b, err := c.NameIsExists(param.name); err != nil {
		fmt.Println(err)
		return
	} else if b {
		fmt.Println("name exists")
		return
	}

	acctsig, err := ethacct.LoadAcctFromFile("asimov")
	if err != nil {
		panic(err)
	}

	acctTx := loadAcct()

	if param.name == "" || param.year == 0 || param.addr == "" || param.sigPrice == "" {
		fmt.Println("parameter error")
		return
	}
	//addr2 == > user addr
	//lastprice from chainlink
	//nonce from ethereum

	erc20Addr := common.HexToAddress(param.addr)
	userAddr := acctTx.EAddr
	lastPrice, _ := (&big.Int{}).SetString(param.lastPrice, 10)
	nonce, _ := (&big.Int{}).SetString(param.nonce, 10)
	price, _ := (&big.Int{}).SetString(param.sigPrice, 10)

	data := pack(param.name, param.year, erc20Addr, lastPrice, nonce, price, userAddr)

	h1 := crypto.Keccak256(data)

	var sig []byte
	sig, err = crypto.Sign(h1[:], acctsig.PrivKey)
	if err != nil {
		panic(err)
	}

	if sig[len(sig)-1] <= 1 {
		sig[len(sig)-1] = sig[len(sig)-1] + 27
	}

	fmt.Println("name", param.name)
	fmt.Println("year", param.year)
	fmt.Println("erc20Addr", erc20Addr.String())
	fmt.Println("lastPrice", lastPrice.String())
	fmt.Println("nonce", nonce.String())
	fmt.Println("price", price.String())
	fmt.Println("userAddr", userAddr.String())

	fmt.Println("sig", hex.EncodeToString(sig))

	//var (
	//	r   *types.Receipt
	//	err error
	//)
	//if r, err = c.MintName(loadAcct().PrivKey, param.name, param.year, common.HexToAddress(param.addr)); err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	//fmt.Println("top level set contract result:", ReceiptResult[r.Status], "tx:", r.TxHash.String())
}

func dnsMintSecondName(_ *cobra.Command, _ []string) {
	c := loadConf()
	if AddrIsZero(c.SecondName) {
		fmt.Println("please deploy contract first")
		return
	}

	if AddrIsZero(c.UsdtAddr) {
		fmt.Println("please set usdt address")
		return
	}

	if !IsNormalName(param.name) {
		fmt.Println("not a valid second level name")
		return
	}

	if param.year <= 0 || param.year > 10 {
		fmt.Println("not a valid year")
		return
	}

	if tl, _, err := GetTLNameFromNormalName(param.name); err != nil {
		fmt.Println("not a correct second name")
		return
	} else {
		var b bool
		if b, err = c.NameIsExists(tl); err != nil {
			fmt.Println(err)
			return
		} else if !b {
			fmt.Println("top level name not exists")
			return
		}
		var addr common.Address
		if addr, err = c.GetErc721Addr(crypto.Keccak256Hash([]byte(tl))); err != nil {
			fmt.Println(err)
			return
		}
		if AddrIsZero(addr) {
			fmt.Println("top level name not opened to register", tl)
			return
		}
	}

	var (
		r   *types.Receipt
		err error
	)
	if r, err = c.MintSecondLevelName(loadAcct().PrivKey, param.name, param.year, common.HexToAddress(param.addr)); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("top level set contract result:", ReceiptResult[r.Status], "tx:", r.TxHash.String())
}

func dnsSecondSetContract(_ *cobra.Command, _ []string) {
	c := loadConf()
	if AddrIsZero(c.TopName) || AddrIsZero(c.SecondName) || AddrIsZero(c.Cost) || AddrIsZero(c.Accountant) {
		fmt.Println("please deploy contract first")
		return
	}

	var (
		r   *types.Receipt
		err error
	)
	if r, err = c.SetDnsSecondContract(loadAcct().PrivKey); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("second level set contract result:", ReceiptResult[r.Status], "tx:", r.TxHash.String())
}

func dnsTopShowOwner(_ *cobra.Command, _ []string) {
	c := loadConf()
	if AddrIsZero(c.TopName) {
		fmt.Println("please deploy top level name contract")
		return
	}

	if addr, err := c.GetOwner(); err != nil {
		fmt.Println("get owner address failed")
	} else {
		fmt.Println("top level name contract owner is: ", addr.String())
	}

}

func dnsTopErc721Addr(_ *cobra.Command, _ []string) {
	c := loadConf()
	if AddrIsZero(c.TopName) {
		fmt.Println("please deploy top level name contract")
		return
	}

	hash := common.Hash{}

	if param.hash != "" {
		if h, err := hex.DecodeString(param.hash); err != nil {
			fmt.Println(err)
			return
		} else {
			copy(hash[:], h)
		}
	} else if param.name != "" {
		hash = crypto.Keccak256Hash([]byte(param.name))
	}

	if addr, err := c.GetErc721Addr(hash); err != nil {
		fmt.Println("get erc721 address failed")
	} else {
		fmt.Println("Erc721 Address: ", addr.String())
	}

}

func dnsTopNameInfo(_ *cobra.Command, _ []string) {
	c := loadConf()
	if AddrIsZero(c.TopName) {
		fmt.Println("please deploy top level name contract")
		return
	}

	hash := common.Hash{}

	if param.hash != "" {
		if h, err := hex.DecodeString(param.hash); err != nil {
			fmt.Println(err)
			return
		} else {
			copy(hash[:], h)
		}
	} else if param.name != "" {
		hash = crypto.Keccak256Hash([]byte(param.name))
	}

	if hash == common.Hash(common.Hash{}) {
		fmt.Println("name not given")
		return
	}

	if name, expireTime, tokenId, err := c.GetDnsInfo(hash); err != nil {
		fmt.Println("get name info failed")
	} else {
		fmt.Println("name: ", name, "expire time:", expireTime, "token id:", tokenId.String())
	}
}

func dnsTopOpen2Reg(_ *cobra.Command, _ []string) {
	c := loadConf()
	if AddrIsZero(c.TopName) {
		fmt.Println("please deploy top level name contract")
		return
	}

	hash := common.Hash{}

	if param.hash != "" {
		if h, err := hex.DecodeString(param.hash); err != nil {
			fmt.Println(err)
			return
		} else {
			copy(hash[:], h)
		}
	} else if param.name != "" {
		hash = crypto.Keccak256Hash([]byte(param.name))
	}

	if hash == common.Hash(common.Hash{}) {
		fmt.Println("name not given")
		return
	}

	if r, err := c.Open2Reg(loadAcct().PrivKey, hash); err != nil {
		fmt.Println("open to register failed")
	} else {
		fmt.Println("open to register result:", ReceiptResult[r.Status], "tx:", r.TxHash.String())
	}
}

func costSetUsdt(_ *cobra.Command, _ []string) {
	c := loadConf()
	if AddrIsZero(c.Cost) {
		fmt.Println("please deploy cost contract first")
		return
	}

	if b := common.IsHexAddress(param.addr); !b {
		fmt.Println("please set correct usdt address")
		return
	}

	if param.decimal > 18 || param.decimal < 1 {
		fmt.Println("please set correct usdt decimal")
		return
	}

	var (
		r   *types.Receipt
		err error
	)
	if r, err = c.SetUsdt(loadAcct().PrivKey, common.HexToAddress(param.addr), param.decimal); err != nil {
		fmt.Println(err.Error())
		return
	}

	c.UsdtAddr = common.HexToAddress(param.addr)
	savefile(c)

	fmt.Println("tcost set usdt result:", ReceiptResult[r.Status], "tx:", r.TxHash.String())

}

func costShowPrice(_ *cobra.Command, _ []string) {
	c := loadConf()
	if AddrIsZero(c.Cost) {
		fmt.Println("please deploy cost contract first")
		return
	}

	hash := common.Hash{}

	if param.hash != "" {
		if h, err := hex.DecodeString(param.hash); err != nil {
			fmt.Println(err)
			return
		} else {
			copy(hash[:], h)
		}
	} else if param.name != "" {
		hash = crypto.Keccak256Hash([]byte(param.name))
	}

	if price, err := c.GetPriceConf(hash); err != nil {
		fmt.Println(err)
		return
	} else {
		for i := 0; i < len(price); i++ {
			fmt.Println(i+1, ":", price[i].String())
		}
	}

}

func costShowTax(_ *cobra.Command, _ []string) {
	c := loadConf()
	if AddrIsZero(c.Cost) {
		fmt.Println("please deploy cost contract first")
		return
	}

	if tax, err := c.GetTaxConf(); err != nil {
		fmt.Println(err)
		return
	} else {
		for i := 0; i < len(tax); i++ {
			fmt.Println(i+1, ":", tax[i].String())
		}
	}

}

func costAddPayment(_ *cobra.Command, _ []string) {
	c := loadConf()
	if AddrIsZero(c.Cost) {
		fmt.Println("please deploy cost contract first")
		return
	}

	if b := common.IsHexAddress(param.addr); !b {
		fmt.Println("please set correct erc20 address")
		return
	}

	if b := common.IsHexAddress(param.addr2); !b {
		fmt.Println("please set correct chain link address")
		return
	}

	if param.decimal > 18 || param.decimal < 1 {
		fmt.Println("please set correct erc20 decimal")
		return
	}

	var (
		r   *types.Receipt
		err error
	)
	if r, err = c.AddPayment(loadAcct().PrivKey,
		common.HexToAddress(param.addr),
		param.decimal, common.HexToAddress(param.addr2)); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("cost add payment result:", ReceiptResult[r.Status], "tx:", r.TxHash.String())

}

func costDisablePayment(_ *cobra.Command, _ []string) {
	c := loadConf()
	if AddrIsZero(c.Cost) {
		fmt.Println("please deploy cost contract first")
		return
	}

	if b := common.IsHexAddress(param.addr); !b {
		fmt.Println("please set correct erc20 address")
		return
	}

	if param.decimal > 18 || param.decimal < 1 {
		fmt.Println("please set correct erc20 decimal")
		return
	}

	var (
		r   *types.Receipt
		err error
	)
	if r, err = c.DisablePayment(loadAcct().PrivKey,
		common.HexToAddress(param.addr),
		param.decimal); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("cost disable payment result:", ReceiptResult[r.Status], "tx:", r.TxHash.String())

}

func costSetTopPrice(_ *cobra.Command, _ []string) {
	c := loadConf()
	if AddrIsZero(c.Cost) {
		fmt.Println("please deploy cost contract first")
		return
	}

	if param.decimal > 18 || param.decimal < 1 {
		fmt.Println("please set correct erc20 decimal")
		return
	}

	var ps [8]*big.Int

	p := int64(math.Pow(10, float64(param.decimal)))

	pb := big.NewInt(p)

	price := big.NewInt(int64(param.price))
	price = price.Mul(price, pb)

	for i := 0; i < 8; i++ {
		ps[i] = price
	}

	var (
		r   *types.Receipt
		err error
	)
	if r, err = c.SetTopLevelPrice(loadAcct().PrivKey, ps); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("cost set top level price result:", ReceiptResult[r.Status], "tx:", r.TxHash.String())

}

func costSetSecondTax(_ *cobra.Command, _ []string) {
	c := loadConf()
	if AddrIsZero(c.Cost) {
		fmt.Println("please deploy cost contract first")
		return
	}

	if param.decimal > 18 || param.decimal < 1 {
		fmt.Println("please set correct erc20 decimal")
		return
	}

	var ps [8]*big.Int

	p := int64(math.Pow(10, float64(param.decimal)))

	pb := big.NewInt(p)

	price := big.NewInt(int64(param.price))
	price = price.Mul(price, pb)

	for i := 0; i < 8; i++ {
		ps[i] = price
	}

	var (
		r   *types.Receipt
		err error
	)
	if r, err = c.SetSecondTax(loadAcct().PrivKey, ps); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("cost set second level tax result:", ReceiptResult[r.Status], "tx:", r.TxHash.String())

}

func costSetSecondPrice(_ *cobra.Command, _ []string) {
	c := loadConf()
	if AddrIsZero(c.Cost) {
		fmt.Println("please deploy cost contract first")
		return
	}
	hash := common.Hash{}

	if param.hash != "" {
		if h, err := hex.DecodeString(param.hash); err != nil {
			fmt.Println(err)
			return
		} else {
			copy(hash[:], h)
		}
	} else if param.name != "" {
		hash = crypto.Keccak256Hash([]byte(param.name))
	}

	if hash == common.Hash(common.Hash{}) {
		fmt.Println("name not given")
		return
	}

	if param.decimal > 18 || param.decimal < 1 {
		fmt.Println("please set correct erc20 decimal")
		return
	}

	var ps [8]*big.Int

	p := int64(math.Pow(10, float64(param.decimal)))

	pb := big.NewInt(p)

	price := big.NewInt(int64(param.price))
	price = price.Mul(price, pb)

	for i := 0; i < 8; i++ {
		ps[i] = price
	}

	var (
		r   *types.Receipt
		err error
	)
	if r, err = c.SetSecondPrice(loadAcct().PrivKey, hash, ps); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("cost set second level tax result:", ReceiptResult[r.Status], "tx:", r.TxHash.String())

}

func pack(name string, year uint8, erc20Addr common.Address, lastPriceId *big.Int,
	nonce *big.Int, price *big.Int, userAddr common.Address) []byte {
	var concatBytes []byte

	concatBytes = append(concatBytes, []byte(name)...)
	concatBytes = append(concatBytes, year)
	concatBytes = append(concatBytes, erc20Addr[:]...)

	var pid [10]byte
	lastPriceId.FillBytes(pid[:])
	concatBytes = append(concatBytes, pid[:]...)
	var pn [32]byte
	nonce.FillBytes(pn[:])
	concatBytes = append(concatBytes, pn[:]...)
	var pb [32]byte
	price.FillBytes(pb[:])

	concatBytes = append(concatBytes, pb[:]...)

	concatBytes = append(concatBytes, userAddr[:]...)

	return concatBytes

}

func makeSign(_ *cobra.Command, _ []string) {
	acct, err := ethacct.LoadAcctFromFile("asimov")
	if err != nil {
		panic(err)
	}

	if param.name == "" || param.year == 0 || param.addr == "" ||
		param.addr2 == "" || param.lastPrice == "" ||
		param.nonce == "" || param.sigPrice == "" {
		fmt.Println("parameter error")
		return
	}

	erc20Addr := common.HexToAddress(param.addr)
	userAddr := common.HexToAddress(param.addr2)
	lastPrice, _ := (&big.Int{}).SetString(param.lastPrice, 10)
	nonce, _ := (&big.Int{}).SetString(param.nonce, 10)
	price, _ := (&big.Int{}).SetString(param.sigPrice, 10)

	data := pack(param.name, param.year, erc20Addr, lastPrice, nonce, price, userAddr)

	h1 := crypto.Keccak256(data)

	var sig []byte
	sig, err = crypto.Sign(h1[:], acct.PrivKey)
	if err != nil {
		panic(err)
	}

	if sig[len(sig)-1] <= 1 {
		sig[len(sig)-1] = sig[len(sig)-1] + 27
	}

	fmt.Println("name", param.name)
	fmt.Println("year", param.year)
	fmt.Println("erc20Addr", erc20Addr.String())
	fmt.Println("lastPrice", lastPrice.String())
	fmt.Println("nonce", nonce.String())
	fmt.Println("price", price.String())
	fmt.Println("userAddr", userAddr.String())

	fmt.Println("sig", hex.EncodeToString(sig))

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func loadAcct() *ethacct.EthAccount {
	acct, err := ethacct.LoadAcct("test")
	if err != nil {
		fmt.Println(err.Error())
		acct, err = ethacct.NewEthAccount()
		if err != nil {
			panic(err)
		}
		acct.Save("test")
	}
	return acct
}

func loadConf() *Contract {
	c := &Contract{}
	if tools.FileExists(contract_conf) {
		var (
			data []byte
			err  error
		)
		if data, err = tools.OpenAndReadAll(contract_conf); err != nil {
			panic(err)
		}
		if err = json.Unmarshal(data, c); err != nil {
			panic(err)
		}
	}
	return c
}
