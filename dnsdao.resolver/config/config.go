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
			WsRPCEndPoint:        "wss://ropsten.infura.io/ws/v3/",
			RpcEndPointID:        "ced16671f5894b2796224e49062999ca",
			DnsName:              "0x3Bb534461b700A752aac8dCe9363E16910d5F66a",
			DnsOwner:             "0xbfdAF6EcD61f078D4f52b75679D89320c00d2368",
			DnsPrice:             "0x39f8817db77Fc26C0B2EE1C8f841868e811d4375",
			DnsConf:              "0xa8764D4109607955dA14bfEc62bbD9a76339db52",
			DnsAccountant:        "0xB9E7ECCfb7f48f9bCd10BC8F768ca3ab43E46E37",
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
