SHELL=PATH='$(PATH)' /bin/sh

.PHONY: all
all:
	abigen --abi dns-contract-abi/BasDomain.abi --pkg dnscontract --type BasDomain --out ./dnscontract/BasDomain.go
	abigen --abi dns-contract-abi/BasMail.abi --pkg dnscontract --type BasMail --out ./dnscontract/BasMail.go
	abigen --abi dns-contract-abi/BasMarket.abi --pkg dnscontract --type BasMarket --out ./dnscontract/BasMarket.go
	abigen --abi dns-contract-abi/BasMiner.abi --pkg dnscontract --type BasMiner --out ./dnscontract/BasMiner.go
	abigen --abi dns-contract-abi/BasOANN.abi --pkg dnscontract --type BasOANN --out ./dnscontract/BasOANN.go
	abigen --abi dns-contract-abi/BasAccountant.abi --pkg dnscontract --type BasAccountant --out ./dnscontract/BasAccountant.go
	abigen --abi dns-contract-abi/BasToken.abi --pkg dnscontract --type BasToken --out ./dnscontract/BasToken.go
	abigen --abi dns-contract-abi/BasView.abi --pkg dnscontract --type BasView --out ./dnscontract/BasView.go
	abigen --abi dns-contract-abi/FreeUDID.abi --pkg dnscontract --type FreeUDID --out ./dnscontract/FreeUDID.go
	abigen --abi dns-contract-abi/BasDomainConf.abi --pkg dnscontract --type BasDomainConf --out ./dnscontract/BasDomainConf.go
	abigen --abi dns-contract-abi/BasExpiredOwnership.abi --pkg dnscontract --type BasExpiredOwnership --out ./dnscontract/BasExpiredOwnership.go
	abigen --abi dns-contract-abi/BasHash.abi --pkg dnscontract --type BasHash --out ./dnscontract/BasHash.go
	abigen --abi dns-contract-abi/BasMailManager.abi --pkg dnscontract --type BasMailManager --out ./dnscontract/BasMailManager.go
	abigen --abi dns-contract-abi/BasRelations.abi --pkg dnscontract --type BasRelations --out ./dnscontract/BasRelations.go
	abigen --abi dns-contract-abi/BasRootDomain.abi --pkg dnscontract --type BasRootDomain --out ./dnscontract/BasRootDomain.go
	abigen --abi dns-contract-abi/BasSet.abi --pkg dnscontract --type BasSet --out ./dnscontract/BasSet.go
	abigen --abi dns-contract-abi/BasSubDomain.abi --pkg dnscontract --type BasSubDomain --out ./dnscontract/BasSubDomain.go
	abigen --abi dns-contract-abi/BasTradableOwnership.abi --pkg dnscontract --type BasTradableOwnership --out ./dnscontract/BasTradableOwnership.go
	abigen --abi dns-contract-abi/ContractReceiver.abi --pkg dnscontract --type ContractReceiver --out ./dnscontract/ContractReceiver.go
	abigen --abi dns-contract-abi/HasRelations.abi --pkg dnscontract --type HasRelations --out ./dnscontract/HasRelations.go
	abigen --abi dns-contract-abi/ManagedByDAO.abi --pkg dnscontract --type ManagedByDAO --out ./dnscontract/ManagedByDAO.go
	abigen --abi dns-contract-abi/WhiteListContract.abi --pkg dnscontract --type WhiteList --out ./dnscontract/WhiteList.go
	abigen --abi dns-contractv2-abi/ColdBoot.abi --pkg udidc --type ColdBoot --out ./dnscontractv2/ColdBoot.go
	abigen --abi dns-contractv2-abi/SimpleWhiteList.abi --pkg udidc --type SimpleWhiteList --out ./dnscontractv2/SimpleWhiteList.go
	abigen --abi dns-contractv2-abi/DaoTax.abi --pkg udidc --type DaoTax --out ./dnscontractv2/DaoTax.go
	abigen --abi dns-contractv2-abi/DnsAccountant.abi --pkg udidc --type DnsAccountant --out ./dnscontractv2/DnsAccountant.go
	abigen --abi dns-contractv2-abi/DnsConf.abi --pkg udidc --type DnsConf --out ./dnscontractv2/DnsConf.go
	abigen --abi dns-contractv2-abi/DnsName.abi --pkg udidc --type DnsName --out ./dnscontractv2/DnsName.go
	abigen --abi dns-contractv2-abi/DnsPayment.abi --pkg udidc --type DnsPayment --out ./dnscontractv2/DnsPayment.go
	abigen --abi dns-contractv2-abi/DnsPrice.abi --pkg udidc --type DnsPrice --out ./dnscontractv2/DnsPrice.go
	abigen --abi dns-contractv2-abi/DnsSubName.abi --pkg udidc --type DnsSubName --out ./dnscontractv2/DnsSubName.go

abi:
	solc --abi --overwrite dns-contract/BasDomain.sol -o ./dns-contract-abi
	solc --abi --overwrite dns-contract/BasMail.sol -o ./dns-contract-abi
	solc --abi --overwrite dns-contract/BasMarket.sol -o ./dns-contract-abi
	solc --abi --overwrite dns-contract/BasMiner.sol -o ./dns-contract-abi
	solc --abi --overwrite dns-contract/BasOANN.sol -o ./dns-contract-abi
	solc --abi --overwrite dns-contract/BasOwnership.sol -o ./dns-contract-abi
	solc --abi --overwrite dns-contract/BasToken.sol -o ./dns-contract-abi
	solc --abi --overwrite dns-contract/BasView.sol -o ./dns-contract-abi
	solc --abi --overwrite dns-contract/FreeUDID.sol -o ./dns-contract-abi
	solc --abi --overwrite dns-contract/WhiteList.sol -o ./dns-contract-abi
	solc --abi --overwrite dns-contractv2/bootProject/DnsBoot.sol -o ./dns-contractv2-abi --allow-paths *,
	solc --abi --overwrite dns-contractv2/dnsProject/DnsName.sol -o ./dns-contractv2-abi --allow-paths *,
	solc --abi --overwrite dns-contractv2/dnsProject/DnsSubName.sol -o ./dns-contractv2-abi --allow-paths *,
	#solc --abi --overwrite dns-contractv2/dnsProject/owners/DnsOwners.sol -o ./dns-contractv2-abi --allow-paths *,
	solc --abi --overwrite dns-contractv2/dnsProject/accountant/DnsAccountant.sol -o ./dns-contractv2-abi --allow-paths *,
	solc --abi --overwrite dns-contractv2/dnsProject/conf/DnsConf.sol -o ./dns-contractv2-abi --allow-paths *,
	solc --abi --overwrite dns-contractv2/dnsProject/payment/DnsPayment.sol -o ./dns-contractv2-abi --allow-paths *,
	solc --abi --overwrite dns-contractv2/dnsProject/price/DnsPrice.sol -o ./dns-contractv2-abi --allow-paths *,
	solc --abi --overwrite dns-contractv2/dnsProject/tax/DaoTax.sol -o ./dns-contractv2-abi --allow-paths *,

clean:
#	rm -f *.abi
	rm -f *.go
