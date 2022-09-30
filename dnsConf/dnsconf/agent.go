package dnsconf

import (
	"errors"
	udidc22 "github.com/dnsdao/dnsdao.resolver/contract/dnscontractv22"
	"github.com/dnsdao/dnsdao.resolver/dnsConf/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"time"
)

func IsRootName(name string) bool {
	if len(name) > 64 || len(name) == 0 {
		return false
	}

	if name[0] == 45 || name[len(name)-1] == 45 {
		return false
	}
	//45 ==> -
	//46 ==> .
	for i := 0; i < len(name); i++ {
		if !((name[i] >= 48 && name[i] <= 57) || name[i] == 45 || (name[i] >= 97 && name[i] <= 122)) {
			return false
		}
	}

	return true
}

func IsNormalName(name string) bool {
	if len(name) > 192 || len(name) <= 2 {
		return false
	}

	if name[0] == 45 || name[0] == 46 {
		return false
	}

	if name[len(name)-1] == 45 || name[len(name)-1] == 46 {
		return false
	}

	//45 ==> -
	//46 ==> .
	var pre bool
	for i := 0; i < len(name); i++ {
		if !((name[i] >= 48 && name[i] <= 57) || name[i] == 46 || name[i] == 45 || (name[i] >= 97 && name[i] <= 122)) {
			return false
		}
		if name[i] == 46 && pre {
			return false
		} else if name[i] == 46 && !pre {
			pre = true
		} else {
			pre = false
		}
	}

	return true
}

func getHash(name string, arrdot []string) (father common.Hash, second common.Hash, nameHash common.Hash) {
	l := len(arrdot)
	father = crypto.Keccak256Hash([]byte(arrdot[l-1]))

	secondName := arrdot[l-2] + "." + arrdot[l-1]

	second = crypto.Keccak256Hash([]byte(secondName))

	nameHash = crypto.Keccak256Hash([]byte(name))

	return
}

func FindFromContract(name string) (expired bool, owner common.Address, err error) {
	//if top
	if IsRootName(name) {
		return findTop(crypto.Keccak256Hash([]byte(name)))
	}
	//if second
	//if other
	if !IsNormalName(name) {
		return false, [20]byte{}, err
	}
	arrDot := strings.Split(name, ".")
	if len(arrDot) < 2 {
		return false, [20]byte{}, errors.New("not a correct name")
	}
	if len(arrDot) == 2 {
		return findSecond(crypto.Keccak256Hash([]byte(arrDot[1])), crypto.Keccak256Hash([]byte(name)))
	} else {
		fh, sh, eh := getHash(name, arrDot)
		return findSub(fh, sh, eh)
	}

	return
}

func findTop(nameHash common.Hash) (expired bool, owner common.Address, err error) {
	var cli *ethclient.Client
	if cli, err = ethclient.Dial(config.GetDnsConfSetting().RpcEndPoint); err != nil {
		return false, [20]byte{}, err
	}
	defer cli.Close()

	var top *udidc22.DnsTopLevelName

	if top, err = udidc22.NewDnsTopLevelName(common.HexToAddress(config.GetDnsConfSetting().TopContractAddr), cli); err != nil {
		return false, [20]byte{}, err
	}

	var r struct {
		EntireName string
		ExpireTime uint32
		TokenId    *big.Int
	}
	if r, err = top.NameStore(nil, nameHash); err != nil {
		return false, [20]byte{}, err
	}

	var erc721Addr common.Address
	if erc721Addr, err = top.Erc721Store(nil, common.Hash{}); err != nil {
		return false, [20]byte{}, err
	}

	var erc721 *udidc22.DnsNameErc721
	if erc721, err = udidc22.NewDnsNameErc721(erc721Addr, cli); err != nil {
		return false, [20]byte{}, err
	}

	if owner, err = erc721.OwnerOf(nil, r.TokenId); err != nil {
		return false, [20]byte{}, err
	}

	if time.Now().Unix() > int64(r.ExpireTime) {
		expired = true
	}

	return
}

func findSecond(fatherHash common.Hash, nameHash common.Hash) (expired bool, owner common.Address, err error) {
	var cli *ethclient.Client
	if cli, err = ethclient.Dial(config.GetDnsConfSetting().RpcEndPoint); err != nil {
		return false, [20]byte{}, err
	}
	defer cli.Close()

	var second *udidc22.DnsSecondLevelName
	if second, err = udidc22.NewDnsSecondLevelName(common.HexToAddress(config.GetDnsConfSetting().SecondContractAddr), cli); err != nil {
		return false, [20]byte{}, err
	}

	var r struct {
		EntireName string
		ExpireTime uint32
		TokenId    *big.Int
	}

	if r, err = second.NameStore(nil, fatherHash, nameHash); err != nil {
		return false, [20]byte{}, err
	}

	var top *udidc22.DnsTopLevelName

	if top, err = udidc22.NewDnsTopLevelName(common.HexToAddress(config.GetDnsConfSetting().TopContractAddr), cli); err != nil {
		return false, [20]byte{}, err
	}
	var erc721Addr common.Address
	if erc721Addr, err = top.Erc721Store(nil, fatherHash); err != nil {
		return false, [20]byte{}, err
	}

	var erc721 *udidc22.DnsNameErc721
	if erc721, err = udidc22.NewDnsNameErc721(erc721Addr, cli); err != nil {
		return false, [20]byte{}, err
	}

	if owner, err = erc721.OwnerOf(nil, r.TokenId); err != nil {
		return false, [20]byte{}, err
	}

	if time.Now().Unix() > int64(r.ExpireTime) {
		expired = true
	}

	return
}

func findSub(fatherHash, secondHash, nameHash common.Hash) (expired bool, owner common.Address, err error) {
	var cli *ethclient.Client
	if cli, err = ethclient.Dial(config.GetDnsConfSetting().RpcEndPoint); err != nil {
		return false, [20]byte{}, err
	}
	defer cli.Close()

	var second *udidc22.DnsSecondLevelName
	if second, err = udidc22.NewDnsSecondLevelName(common.HexToAddress(config.GetDnsConfSetting().SecondContractAddr), cli); err != nil {
		return false, [20]byte{}, err
	}

	var subName string
	if subName, err = second.SubNameStore(nil, secondHash, nameHash); err != nil {
		return false, [20]byte{}, err
	}
	if subName == "" {
		return false, [20]byte{}, errors.New("not found sub name")
	}

	var r struct {
		EntireName string
		ExpireTime uint32
		TokenId    *big.Int
	}

	if r, err = second.NameStore(nil, fatherHash, secondHash); err != nil {
		return false, [20]byte{}, err
	}

	var top *udidc22.DnsTopLevelName

	if top, err = udidc22.NewDnsTopLevelName(common.HexToAddress(config.GetDnsConfSetting().TopContractAddr), cli); err != nil {
		return false, [20]byte{}, err
	}
	var erc721Addr common.Address
	if erc721Addr, err = top.Erc721Store(nil, fatherHash); err != nil {
		return false, [20]byte{}, err
	}

	var erc721 *udidc22.DnsNameErc721
	if erc721, err = udidc22.NewDnsNameErc721(erc721Addr, cli); err != nil {
		return false, [20]byte{}, err
	}

	if owner, err = erc721.OwnerOf(nil, r.TokenId); err != nil {
		return false, [20]byte{}, err
	}

	if time.Now().Unix() > int64(r.ExpireTime) {
		expired = true
	}

	return

}
