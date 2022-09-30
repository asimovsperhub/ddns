module github.com/dnsdao/dnsdao.resolver/dnsDaoExchange

go 1.16

require (
	github.com/dnsdao/dnsdao.resolver/contract/dnscontractv22 v0.0.0
	github.com/ethereum/go-ethereum v1.10.25
	github.com/kprc/nbsnetwork v0.0.0-20210413001421-be5dd1f08982
)

require github.com/kprc/nbsdht v0.0.0-20190523081012-710019eb9202 // indirect

replace github.com/dnsdao/dnsdao.resolver/contract/dnscontractv22 => ../contract/dnscontractv22
