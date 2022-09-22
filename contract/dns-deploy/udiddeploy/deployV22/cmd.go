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

var dnsSecondSetContractCmd = &cobra.Command{
	Use:   "set-contract",
	Short: "dns second level name contract set top cost accountant contract address",
	Long:  "dns second level name contract set top cost accountant contract address",
	Run:   dnsSecondSetContract,
}

var dnsTopOwnerCmd = &cobra.Command{
	Use:   "owner",
	Short: "dns top level name contract show owner",
	Long:  "dns top level name contract show owner",
	Run:   dnsTopShowOwner,
}

var costSetUsdtCmd = &cobra.Command{
	Use:   "usdt",
	Short: "set usdt address and decimal",
	Long:  "set usdt address and decimal",
	Run:   costSetUsdt,
}
