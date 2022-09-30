package main

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "udiddeploy",

	Short: "udid deploy command tools",

	Long: `usage description::TODO::`,

	Run: mainRun,
}

var dnsTopCmd = &cobra.Command{
	Use:   "top",
	Short: "deploy top level name contract",
	Long:  "deploy top level name contract",
	Run:   deployDnsTop,
}

var dnsSecondCmd = &cobra.Command{
	Use:   "second",
	Short: "deploy second level name contract",
	Long:  "deploy second level name contract",
	Run:   deployDnsSecond,
}

var dnsCostCmd = &cobra.Command{
	Use:   "cost",
	Short: "deploy cost contract",
	Long:  "deploy cost contract",
	Run:   deployDnsCost,
}

var dnsAccountantCmd = &cobra.Command{
	Use:   "accountant",
	Short: "deploy accountant contract",
	Long:  "deploy accountant contract",
	Run:   deployDnsAccountant,
}

var dnsTopSetContractCmd = &cobra.Command{
	Use:   "set-contract",
	Short: "dns top level name contract set dnssecond cost accountant contract address",
	Long:  "dns top level name contract set dnssecond cost accountant contract address",
	Run:   dnsTopSetContract,
}

var dnsTopOwnerCmd = &cobra.Command{
	Use:   "owner",
	Short: "dns top level name contract show owner",
	Long:  "dns top level name contract show owner",
	Run:   dnsTopShowOwner,
}

var dnsTopMintNameCmd = &cobra.Command{
	Use:   "mint",
	Short: "mint a top level name",
	Long:  "mint a top level name",
	Run:   dnsTopMintName,
}

var dnsTopMintNameBySigCmd = &cobra.Command{
	Use:   "mintbysig",
	Short: "mint a top level name by signature",
	Long:  "mint a top level name by signature",
	Run:   dnsTopMintNameBySig,
}

var dnsTopErc721AddrCmd = &cobra.Command{
	Use:   "erc721",
	Short: "show erc721 Addr",
	Long:  "show erc721 Addr",
	Run:   dnsTopErc721Addr,
}

var dnsTopNameInfoCmd = &cobra.Command{
	Use:   "name-info",
	Short: "show name info",
	Long:  "show name info",
	Run:   dnsTopNameInfo,
}

var dnsTopOpen2RegCmd = &cobra.Command{
	Use:   "open",
	Short: "open to register second level name",
	Long:  "open to register second level name",
	Run:   dnsTopOpen2Reg,
}

var dnsSecondSetContractCmd = &cobra.Command{
	Use:   "set-contract",
	Short: "dns second level name contract set top cost accountant contract address",
	Long:  "dns second level name contract set top cost accountant contract address",
	Run:   dnsSecondSetContract,
}

var dnsSecondMintNameCmd = &cobra.Command{
	Use:   "mint",
	Short: "mint second name",
	Long:  "mint second name",
	Run:   dnsMintSecondName,
}

var costAddPaymentCmd = &cobra.Command{
	Use:   "add-payment",
	Short: "add payment",
	Long:  "add payment",
	Run:   costAddPayment,
}

var costDisablePaymentCmd = &cobra.Command{
	Use:   "disable-payment",
	Short: "disable payment",
	Long:  "disable payment",
	Run:   costDisablePayment,
}

var costSetTopPriceCmd = &cobra.Command{
	Use:   "top-price",
	Short: "set top level name price",
	Long:  "set top level name price",
	Run:   costSetTopPrice,
}

var costSetSecondTaxCmd = &cobra.Command{
	Use:   "second-tax",
	Short: "set second level name tax",
	Long:  "set second level name tax",
	Run:   costSetSecondTax,
}
var costSetUsdtCmd = &cobra.Command{
	Use:   "usdt",
	Short: "set usdt address and decimal",
	Long:  "set usdt address and decimal",
	Run:   costSetUsdt,
}

var costSetSecondPriceCmd = &cobra.Command{
	Use:   "second-price",
	Short: "set second level name price",
	Long:  "set second level name price",
	Run:   costSetSecondPrice,
}

var costShowPriceCmd = &cobra.Command{
	Use:   "price",
	Short: "show price",
	Long:  "show price",
	Run:   costShowPrice,
}

var costShowTaxCmd = &cobra.Command{
	Use:   "tax",
	Short: "show tax",
	Long:  "show tax",
	Run:   costShowTax,
}

var makeSignCmd = &cobra.Command{
	Use:   "sign",
	Short: "make sign to register name",
	Long:  "make sign to register name",
	Run:   makeSign,
}
