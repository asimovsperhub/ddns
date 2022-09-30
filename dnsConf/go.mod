module github.com/dnsdao/dnsdao.resolver/dnsConf

go 1.16

require github.com/dnsdao/dnsdao.resolver/contract/dnscontractv22 v0.0.0

require (
	github.com/btcsuite/goleveldb v1.0.0
	github.com/dnsdao/dnsdao.resolver/dnsDaoExchange v0.0.0
	github.com/ethereum/go-ethereum v1.10.25
	github.com/kprc/nbsnetwork v0.0.0-20210413001421-be5dd1f08982
	github.com/spf13/cobra v1.5.0
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.27.1
)

replace github.com/dnsdao/dnsdao.resolver/dnsDaoExchange => ../dnsDaoExchange

replace github.com/dnsdao/dnsdao.resolver/contract/dnscontractv22 => ../contract/dnscontractv22
