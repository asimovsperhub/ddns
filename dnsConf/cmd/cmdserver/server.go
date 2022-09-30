package cmdserver

import (
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/dnsConf/cmd/pbs"
	"github.com/dnsdao/dnsdao.resolver/dnsConf/config"
	"github.com/kprc/nbsnetwork/tools"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

var (
	gListener net.Listener
)

func StartCmdService() {

	if b := tools.FileExists(config.CmdSockFile()); b {
		os.RemoveAll(config.CmdSockFile())
	}
	var err error
	gListener, err = net.Listen("unix", config.CmdSockFile())
	if err != nil {
		panic(err)
	}

	cmdServer := grpc.NewServer()

	pbs.RegisterCmdServiceServer(cmdServer, cmdServerInstance)

	reflection.Register(cmdServer)

	if err = cmdServer.Serve(gListener); err != nil {
		fmt.Println(err)
	}
}

func StopCmdService() {
	gListener.Close()
}
