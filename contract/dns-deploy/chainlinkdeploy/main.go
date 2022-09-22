package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/config"
	"github.com/dnsdao/dnsdao.resolver/dnsDaoExchange/ethacct"
	"github.com/ethereum/go-ethereum/common"
	"github.com/kprc/nbsnetwork/tools"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path"
	"strconv"
	"syscall"
)

var rootCmd = &cobra.Command{
	Use: "fcl",

	Short: "fake chain link command tools",

	Long: `usage description::TODO::`,

	Run: mainRun,
}

var ethUsdtCmd = &cobra.Command{
	Use:   "eth-usdt",
	Short: "deploy eth-usdt contract",
	Long:  `deploy eth-usdt contract`,
	Run:   ethUsdtRun,
}

var fakesolCmd = &cobra.Command{
	Use:   "fake-sol",
	Short: "deploy fake-sol contract",
	Long:  `deploy fake-sol contract`,
	Run:   fakeSolRun,
}

var fakeUsdcCmd = &cobra.Command{
	Use:   "fake-usdc",
	Short: "deploy fake-usdc contract",
	Long:  `deploy fake-usdc contract`,
	Run:   fakeUsdcRun,
}

var fakeUsdt6Cmd = &cobra.Command{
	Use:   "fake-usdt6",
	Short: "deploy fake-usdt6 contract",
	Long:  `deploy fake-usdt6 contract`,
	Run:   fakeUsdt6Run,
}

var solUsdtCmd = &cobra.Command{
	Use:   "sol-usdt",
	Short: "deploy sol-usdt contract",
	Long:  `deploy sol-usdt contract`,
	Run:   solUsdtRun,
}

var usdcUsdtCmd = &cobra.Command{
	Use:   "usdc-usdt",
	Short: "deploy usdc-usdt contract",
	Long:  `deploy usdc-usdt contract`,
	Run:   usdcUsdtRun,
}

var showtCmd = &cobra.Command{
	Use:   "show",
	Short: "show contract",
	Long:  `show contract`,
	Run:   showRun,
}

var privCmd = &cobra.Command{
	Use:   "priv",
	Short: "show private key",
	Long:  `show private key`,
	Run:   privRun,
}

var param struct {
	version bool
}

func init() {

	flags := rootCmd.Flags()

	flags.BoolVarP(&param.version, "version",
		"v", false, "fcl -v")

	//rootCmd.AddCommand(cmdclient.ShowCmd)
	//rootCmd.AddCommand(cmd.InitCmd)
	rootCmd.AddCommand(ethUsdtCmd)
	rootCmd.AddCommand(fakesolCmd)
	rootCmd.AddCommand(fakeUsdcCmd)
	rootCmd.AddCommand(fakeUsdt6Cmd)
	rootCmd.AddCommand(solUsdtCmd)
	rootCmd.AddCommand(usdcUsdtCmd)
	rootCmd.AddCommand(showtCmd)
	rootCmd.AddCommand(privCmd)
}

type FakeChainLink struct {
	EthUsdt   common.Address `json:"eth_usdt"`
	FakeSol   common.Address `json:"fake_sol"`
	FakeUsdc  common.Address `json:"fake_usdc"`
	FakeUsdt6 common.Address `json:"fake_usdt_6"`
	SolUsdt   common.Address `json:"sol_usdt"`
	UsdcUsdt  common.Address `json:"usdc_usdt"`
}

func (fcl *FakeChainLink) String() string {
	//msg := ""
	//msg += fmt.Sprintf("eth-usdt:%s\r\n",fcl.EthUsdt.String())

	if data, err := json.MarshalIndent(*fcl, "\t", " "); err != nil {
		return err.Error()
	} else {
		return string(data)
	}

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

const (
	configFile = "./.fakechainlink"
	pidfile    = "./.chainlinkpid"
	Version    = "0.0.1"
)

func load() *FakeChainLink {
	fc := &FakeChainLink{}

	if tools.FileExists(configFile) {
		if data, err := tools.OpenAndReadAll(configFile); err != nil {
			panic(err)
		} else {
			if err = json.Unmarshal(data, fc); err != nil {
				panic(err)
			}
		}
	}
	return fc
}

func (fcl *FakeChainLink) Save() error {
	if data, err := json.Marshal(*fcl); err != nil {
		return err
	} else {
		return tools.Save2File(data, configFile)
	}
}

func loadacct() *ethacct.EthAccount {
	acct, err := ethacct.LoadAcct("test")
	if err != nil {
		fmt.Println(err.Error())
		acct, err = ethacct.NewEthAccount()
		if err != nil {
			panic(err)
		}
		acct.Save("test")
	}

	log.Println(acct.SAddr)

	return acct
}

func ethUsdtRun(_ *cobra.Command, _ []string) {
	fcl := load()
	var err error
	if fcl.EthUsdt, err = DeployEthUsdt(loadacct().PrivKey); err != nil {
		fmt.Println(err.Error())
		return
	}
	fcl.Save()
	fmt.Println(fcl.String())
}

func fakeSolRun(_ *cobra.Command, _ []string) {
	fcl := load()
	var err error
	if fcl.FakeSol, err = DeployFakeSol(loadacct().PrivKey); err != nil {
		fmt.Println(err.Error())
		return
	}
	fcl.Save()
	fmt.Println(fcl.String())
}

func fakeUsdcRun(_ *cobra.Command, _ []string) {
	fcl := load()
	var err error
	if fcl.FakeUsdc, err = DeployFakeUsdc(loadacct().PrivKey); err != nil {
		fmt.Println(err.Error())
		return
	}
	fcl.Save()
	fmt.Println(fcl.String())
}

func fakeUsdt6Run(_ *cobra.Command, _ []string) {
	fcl := load()
	var err error
	if fcl.FakeUsdt6, err = DeployFakeUsdt6(loadacct().PrivKey); err != nil {
		fmt.Println(err.Error())
		return
	}
	fcl.Save()
	fmt.Println(fcl.String())
}

func solUsdtRun(_ *cobra.Command, _ []string) {
	fcl := load()
	var err error
	if fcl.SolUsdt, err = DeploySolUsdt(loadacct().PrivKey); err != nil {
		fmt.Println(err.Error())
		return
	}
	fcl.Save()
	fmt.Println(fcl.String())
}

func usdcUsdtRun(_ *cobra.Command, _ []string) {
	fcl := load()
	var err error
	if fcl.UsdcUsdt, err = DeployUsdcUsdt(loadacct().PrivKey); err != nil {
		fmt.Println(err.Error())
		return
	}
	fcl.Save()
	fmt.Println(fcl.String())
}

func showRun(_ *cobra.Command, _ []string) {
	fcl := load()
	fmt.Println(fcl.String())
}

func privRun(_ *cobra.Command, _ []string) {
	acct := loadacct()

	fmt.Println(hex.EncodeToString(acct.PrivKey.D.Bytes()))
}

func mainRun(_ *cobra.Command, _ []string) {

	//if b := tools.FileExists(configFile); !b {
	//	fmt.Println("please initial fcl first!")
	//	return
	//}

	if param.version {
		fmt.Println(Version)
		return
	}

	//go server.GetServerInstance().StartDaemon()

	//go cmdserver.StartCmdService()
	//go dns.GetDNSAgent().DNSLoopEvent()
	//go httpserver.StartWebDaemon()

	//waitShutdownSignal()
}

func waitShutdownSignal() {

	pid := strconv.Itoa(os.Getpid())
	fmt.Printf("\n>>>>>>>>>>fcl start at pid(%s)<<<<<<<<<<\n", pid)

	pidfile := path.Join(config.ResolveHome(), pidfile)

	if err := ioutil.WriteFile(pidfile, []byte(pid), 0644); err != nil {
		fmt.Println("failed to write running pid", err)
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGUSR1,
		syscall.SIGUSR2)

	sig := <-sigCh

	//server.GetServerInstance().StopDaemon()

	//cmdserver.StopCmdService()
	//
	//ldb.GetLdb().CloseLdb()
	//dns.GetDNSAgent().Quit()
	//httpserver.StopWebDaemon()

	fmt.Printf("\n>>>>>>>>>>process finished(%s)<<<<<<<<<<\n", sig)
}
