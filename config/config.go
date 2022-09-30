package config

import (
	"bytes"
	"encoding/json"
	"github.com/kprc/nbsnetwork/tools"
	"math/rand"
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
			ListenServer: ":30518",
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
			RpcEndPointID: []string{"45c3b62303004dc08aeb7f1bf2a875b5", "135eed6e983843e3b90ae667a1d6a3b6", "5ac1f01fe23347fb9fd7ab0ceb2817cf",
				"ee1f72c75a3f46768cf0f6cea3830881"},
			DnsName:       "0x45419EcDBe2B0419372510d3da5149b88BDcBEe6",
			DnsOwner:      "0x20d86efD2CB75967Acfa24C9580298c0c3D0c684",
			DnsPrice:      "0x5Ae881f69847C47DAaaEdD5864FEef9aBAf763bA",
			DnsConf:       "0xfCaD0039d54E6034588bF28Bfa3d869b6FAf7baB",
			DnsAccountant: "0xE7Ef2FC9B298CAd19184613fc39F93a79e9ff25B",
			ColdBoot:      "0x22871b977AAe43d44FE50dF03f632134c3e3e490",
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
	return rc.Cconf.RpcEndPoint + rc.Cconf.RpcEndPointID[rand.Intn(len(rc.Cconf.RpcEndPointID))]
}
func (rc *RConfig) GetMainRPCEndPoint() string {
	return rc.Cconf.RpcMainEndPoint + rc.Cconf.RpcEndPointID[rand.Intn(len(rc.Cconf.RpcEndPointID))]
}

func (rc *RConfig) GetWsRPCENDPoint() string {
	return rc.Cconf.WsRPCEndPoint + rc.Cconf.RpcEndPointID[rand.Intn(len(rc.Cconf.RpcEndPointID))]
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
	RpcEndPointID        []string
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
