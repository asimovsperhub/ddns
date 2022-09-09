package config

import (
	"bytes"
	"encoding/json"
	"github.com/kprc/nbsnetwork/tools"
	"path"
	"sync"
)

const (
	resolveHome = ".resolve"
	confFile    = "conf.json"
	cmdSock     = ".cmd-sock"
)

var (
	resolveconf       *RConfig
	resolveconfOnceDo sync.Once
)

func defaultConf() *RConfig {
	return &RConfig{
		WConf: &WebSeverConf{
			ListenServer: ":30500",
		},
		Cconf: &ContractConf{
			ChainName:            "rinkeby",
			ChainID:              4,
			OANNContract:         "0xead7705317ece44137bca59080e9881dcfc2879f",
			BasSubContract:       "0xc1998718088d652bbe816525e3a87e73a6cbe662",
			BasRootContract:      "0x0acd28b5c1c65ee70ccb20efc33db1b2c32aa183",
			BasTradableOwnership: "0x39e981f08f3c917e49669649a61720d5cf555ddc",
			BasDomainConf:        "0x68d3aa73d9479cb3b162d279234fdbfd45a1f9b6",
			RpcEndPoint:          "https://rinkeby.infura.io/v3/",
			RpcMainEndPoint:      "https://mainnet.infura.io/v3/",
			WsRPCEndPoint:        "wss://ropsten.infura.io/ws/v3/",
			RpcEndPointID:        "ced16671f5894b2796224e49062999ca",
			DnsName:              "0xE468Ba38F6DaF40E07B4E830E099cF96e90Af8aD",
			DnsOwner:             "0xB865828659694D2C91baE6fB7DDa468b4C64541e",
			DnsPrice:             "0x3086F038757EB4a3185BfEb9690CC4e3103E6223",
			DnsConf:              "0x9218bc65C2Bfb0bc59dCB7DFC11b9CdD94e8eF55",
			DnsAccountant:        "0xb57b91567AC4bA767825755457334C091000Dfa4",
			ColdBoot:             "0x22871b977AAe43d44FE50dF03f632134c3e3e490",
		},
		SConf: &SysConf{
			DBPath: "ldb",
		},
	}
}

func (rc *RConfig) GetDBPath() string {
	return path.Join(ResolveHome(), rc.SConf.DBPath)
}

func (rc *RConfig) GetRPCEndPoint() string {
	return rc.Cconf.RpcEndPoint + rc.Cconf.RpcEndPointID
}
func (rc *RConfig) GetMainRPCEndPoint() string {
	return rc.Cconf.RpcMainEndPoint + rc.Cconf.RpcEndPointID
}

func (rc *RConfig) GetWsRPCENDPoint() string {
	return rc.Cconf.WsRPCEndPoint + rc.Cconf.RpcEndPointID
}

func GetRConf() *RConfig {
	if resolveconf == nil {
		resolveconfOnceDo.Do(func() {
			resolveconf = defaultConf()
			resolveconf.load()
			resolveconf.compareAndSave()
		})
	}

	return resolveconf
}

func InitConfig() {
	nc := defaultConf()
	nc.save()
}

func (nc *RConfig) compareAndSave() {
	data, err := json.Marshal(*nc)
	if err != nil {
		return
	}

	ncdefault := defaultConf()

	var dataDefault []byte
	dataDefault, err = json.Marshal(ncdefault)
	if err != nil {
		return
	}

	if cp := bytes.Compare(data, dataDefault); cp != 0 {
		nc.save()
		return
	}
	var dataload []byte

	ncload := &RConfig{}
	ncload.load()

	dataload, err = json.Marshal(ncload)
	if err != nil {
		return
	}

	if cp := bytes.Compare(dataload, dataDefault); cp != 0 {
		nc.save()
		return
	}

}

func ResolveHome() string {
	h, _ := tools.Home()

	return path.Join(h, resolveHome)
}

func ConfFile() string {
	return path.Join(ResolveHome(), confFile)
}

func CmdSockFile() string {
	return path.Join(ResolveHome(), cmdSock)
}

type WebSeverConf struct {
	ListenServer string `json:"listen_server"`
}

type ContractConf struct {
	ChainID              int64
	ChainName            string
	RpcEndPoint          string
	WsRPCEndPoint        string
	RpcEndPointID        string
	RpcMainEndPoint      string
	OANNContract         string
	BasRootContract      string
	BasTradableOwnership string
	BasDomainConf        string
	BasSubContract       string
	DnsName              string
	DnsOwner             string
	DnsPrice             string
	DnsConf              string
	DnsAccountant        string
	ColdBoot             string
}

type SysConf struct {
	DBPath string
}

type RConfig struct {
	WConf *WebSeverConf `json:"w_conf"`
	Cconf *ContractConf `json:"c_conf"`
	SConf *SysConf      `json:"s_conf"`
}

func (nc *RConfig) load() error {
	cfile := ConfFile()

	fdata, err := tools.OpenAndReadAll(cfile)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(fdata, nc); err != nil {
		return err
	}

	return nil
}

func (nc *RConfig) save() error {
	cfile := ConfFile()

	if j, err := json.MarshalIndent(*nc, "\t", " "); err != nil {
		return err
	} else {
		return tools.Save2File(j, cfile)
	}
}
