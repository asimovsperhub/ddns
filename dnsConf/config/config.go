package config

import (
	"encoding/json"
	"github.com/kprc/nbsnetwork/tools"
	"os"
	"path"
)

type DnsSetting struct {
	TopContractAddr    string `json:"top_contract_addr"`
	SecondContractAddr string `json:"second_contract_addr"`
	RpcEndPoint        string `json:"rpc_end_point"`
	ServerListenAddr   string `json:"server_listen_addr"`
	DbPath             string `json:"db_path"`
}

var gDnsConfSetting *DnsSetting

const (
	DnsConfHome        = ".dnsConfHome"
	DnsConfSettingFile = ".setting"
	sockFile           = ".sock"
	dbPath             = "db"
)

func (ds *DnsSetting) Save() error {
	return nil
}

func dnsConfHome() string {
	if home, err := tools.Home(); err != nil {
		panic(err)
	} else {
		return path.Join(home, DnsConfHome)
	}
}

func dnsSettingFile() string {
	return path.Join(dnsConfHome(), DnsConfSettingFile)
}

func (ds *DnsSetting) save() error {
	if data, err := json.MarshalIndent(*ds, "\t", " "); err != nil {
		return err
	} else {
		return tools.Save2File(data, dnsSettingFile())
	}
}

func loadDnsConfSetting() *DnsSetting {
	h := dnsConfHome()
	if !tools.FileExists(h) {
		if err := os.Mkdir(h, 0755); err != nil {
			panic(err)
		}
	}

	if !tools.FileExists(dnsSettingFile()) {
		sett := &DnsSetting{
			TopContractAddr:    "0xc614058F8AF2440c1535D7B14B5BA992e23CCA81",
			SecondContractAddr: "0x896b01643E48Ca08cE30A6A9acB06C8119b158ae",
			RpcEndPoint:        "https://goerli.infura.io/v3/cbc6ef92e0714bf39f4397904f3f15d2",
			ServerListenAddr:   "0.0.0.0:28288",
			DbPath:             dbPath,
		}
		sett.save()
		return sett
	} else {
		if data, err := tools.OpenAndReadAll(dnsSettingFile()); err != nil {
			panic(err)
		} else {
			sett := &DnsSetting{}
			if err = json.Unmarshal(data, sett); err != nil {
				panic(err)
				return nil
			}
			return sett
		}
	}
}

func init() {
	gDnsConfSetting = loadDnsConfSetting()
}

func GetDnsConfSetting() *DnsSetting {
	return gDnsConfSetting
}

func CmdSockFile() string {
	return path.Join(dnsConfHome(), sockFile)
}

func (ds *DnsSetting) GetDbPath() string {
	return path.Join(dnsConfHome(), ds.DbPath)
}
