package main

import (
	"encoding/json"
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/dnsDaoExchange/ethacct"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/kprc/nbsnetwork/tools"
	"github.com/spf13/cobra"
	"log"
)

var ReceiptResult = [2]string{"Failed", "success"}

const contract_conf = "./contract.create"
const Version = "0.0.1"

var param struct {
	version bool
	force   bool
	addr    string
	decimal uint8
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
	costSetUsdtCmd.Flags().StringVarP(&param.addr, "address", "a", "", "usdt address")
	costSetUsdtCmd.Flags().Uint8VarP(&param.decimal, "decimal", "d", 6, "usdt decimal")

	rootCmd.AddCommand(dnsTopCmd)
	rootCmd.AddCommand(dnsSecondCmd)
	rootCmd.AddCommand(dnsCostCmd)
	rootCmd.AddCommand(dnsAccountantCmd)
	dnsTopCmd.AddCommand(dnsTopSetContractCmd)
	dnsTopCmd.AddCommand(dnsTopOwnerCmd)
	dnsCostCmd.AddCommand(costSetUsdtCmd)
	dnsSecondCmd.AddCommand(dnsSecondSetContractCmd)

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
	if r, err = c.SetDnsTopContract(loadAcct().PrivKey); err != nil {
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

	if addr, err := c.GetOwner(loadAcct().PrivKey); err != nil {
		fmt.Println("get owner address failed")
	} else {
		fmt.Println("top level name contract owner is: ", addr.String())
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

	fmt.Println("tcost set usdt result:", ReceiptResult[r.Status], "tx:", r.TxHash.String())

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
