module github.com/dnsdao/dnsdao.resolver/contract/dns-deploy/udiddeploy/deployV22

go 1.16

require github.com/dnsdao/dnsdao.resolver v0.0.0-20220601030651-c698d986a8e3

require (
	github.com/dnsdao/dnsdao.resolver/dnsDaoExchange v0.0.0
	github.com/ethereum/go-ethereum v1.10.25
	github.com/kprc/nbsnetwork v0.0.0-20210413001421-be5dd1f08982
	github.com/spf13/cobra v1.2.1
)

replace github.com/dnsdao/dnsdao.resolver => ../../../../

replace github.com/dnsdao/dnsdao.resolver/contract/dnscontractv22 => ../../../../contract/dnscontractv22

replace github.com/dnsdao/dnsdao.resolver/dnsDaoExchange => ../../../../dnsDaoExchange
