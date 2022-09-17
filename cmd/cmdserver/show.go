package cmdserver

import (
	"context"
	"encoding/json"
	"github.com/dnsdao/dnsdao.resolver/cmd/pbs"
	"github.com/dnsdao/dnsdao.resolver/config"
)

func (cs *CmdSrv) ShowConfig(context.Context, *pbs.EmptyMessage) (*pbs.CommonResponse, error) {

	rc := config.GetRConf()

	j, _ := json.MarshalIndent(*rc, "\t", " ")

	return &pbs.CommonResponse{
		Msg: string(j),
	}, nil
}
