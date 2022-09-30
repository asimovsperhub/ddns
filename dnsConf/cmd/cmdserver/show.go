package cmdserver

import (
	"context"
	"encoding/json"
	"github.com/dnsdao/dnsdao.resolver/dnsConf/cmd/pbs"
	"github.com/dnsdao/dnsdao.resolver/dnsConf/config"
)

func (cs *CmdSrv) ShowConfig(context.Context, *pbs.EmptyMessage) (*pbs.CommonResponse, error) {

	rc := config.GetDnsConfSetting()

	j, _ := json.MarshalIndent(*rc, "\t", " ")

	return &pbs.CommonResponse{
		Msg: string(j),
	}, nil
}
