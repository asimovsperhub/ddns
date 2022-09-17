package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/xuri/excelize/v2"
	"log"
	"math/big"
	"strings"
)

func printK(hi *big.Int) (int64, int64, int64, *big.Int) {
	a := hi.Mod(hi, big.NewInt(102400))

	aint := a.Int64()
	//fmt.Println(aint)
	//
	//fmt.Println(aint/256, aint%256)

	return aint, aint / 256, aint % 256, a.Exp(big.NewInt(2), big.NewInt(aint%256), nil)
}

func intigridAddr(addr string) string {
	idx := strings.IndexAny(addr, "0x")
	if idx == -1 {
		return ""
	}

	return addr[idx : 42+idx]
}

func main() {
	var bitSet [400]*big.Int
	for i := 0; i < 400; i++ {
		bitSet[i] = big.NewInt(0)
	}
	//read file
	f, err := excelize.OpenFile("all-uniq0829-uniq.xlsx")
	if err != nil {
		println(err.Error())
		return
	}
	//// 获取工作表中指定单元格的值
	//cell, err := f.GetCellValue("Sheet1", "A1")
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//println(cell)
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	cnt := 0
	//h := common.Hash{}

	for _, row := range rows {
		cnt++
		for _, colCell := range row {
			//fmt.Println(colCell)
			printcol(colCell, &bitSet)
			//cnt++
			//if cnt < 3960 {
			if colCell != intigridAddr(colCell) {
				fmt.Println(colCell, cnt)
			}

			//}
		}
		//println()
	}
	//fmt.Println(cnt)

	fmt.Println("===================", cnt)
	fmt.Print("[")
	for i := 0; i < 400; i++ {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print("" + bitSet[i].String() + "")
	}
	fmt.Print("]")
	fmt.Println("")
	//
	//fmt.Println("===================")
	//fmt.Print("[")
	//for i := 10; i < 20; i++ {
	//	if i > 0 {
	//		fmt.Print(",")
	//	}
	//	fmt.Print("" + bitSet[i].String() + "")
	//}
	//fmt.Print("]")
	//fmt.Println("")
	//fmt.Println("===================")
	//fmt.Print("[")
	//for i := 20; i < 30; i++ {
	//	if i > 0 {
	//		fmt.Print(",")
	//	}
	//	fmt.Print("" + bitSet[i].String() + "")
	//}
	//fmt.Print("]")
	//fmt.Println("")
	//
	//fmt.Println("===================")
	//fmt.Print("[")
	//for i := 30; i < 40; i++ {
	//	if i > 0 {
	//		fmt.Print(",")
	//	}
	//	fmt.Print("" + bitSet[i].String() + "")
	//}
	//fmt.Print("]")
	//fmt.Println("")

	//b1 := printcol("0xfdf9811f657f07EA91b9d92366A80357dB52aaBA", &bitSet)
	//calc bytes32
	//b2 := printcol("0xfD8295EE3cc3D7796feD3dC3E815d61e1593ee14", &bitSet)

	//z := big.NewInt(0)
	//z.Xor(bitSet[b1], bitSet[b2])

	//fmt.Println(z)

	//print
	account := common.HexToAddress("0x3D002404deee63697fBEf95657DcE57335BF561D")
	blockNumber := big.NewInt(15184782)
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/ced16671f5894b2796224e49062999ca")
	if err != nil {
		log.Fatal(err)
	}
	balance, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		println(err)
	}

	fmt.Println(balance.String())

	//fmt.Println(intigridAddr(" 0x677FA078f0E979C34aABf287E1b6DaAA9d545074  \t"))

}

func printcol(colCell string, pbitSet *[400]*big.Int) int64 {
	bitSet := *pbitSet
	//print(colCell, "\t")
	if colCell == "0x677FA078f0E979C34aABf287E1b6DaAA9d545074" {
		println(colCell + "----------------->")
	}
	//colCell11 := strings.Trim(colCell, " ")
	addr := common.HexToAddress(colCell)
	if strings.ToLower(colCell) == strings.ToLower("0x677FA078f0E979C34aABf287E1b6DaAA9d545074") {
		fmt.Println("addr===>", addr.String())
	}

	//fmt.Println(addr.String())
	h := crypto.Keccak256Hash(addr[:])
	hi := new(big.Int)
	hi.SetBytes(h[:])
	//fmt.Println(h.String())
	//fmt.Println(hi.String())
	_, b1, c, _ := printK(hi)
	//fmt.Println(a, b1, c)
	bitSet[b1].SetBit(bitSet[b1], int(c), 1)

	h256 := sha256.Sum256(addr[:])
	h2 := new(big.Int)
	h2.SetBytes(h256[:])
	//fmt.Println(h2.String())
	var b2 int64
	_, b2, c, _ = printK(h2)
	//fmt.Println(a, b2, c)
	bitSet[b2].SetBit(bitSet[b2], int(c), 1)

	//fmt.Println("0xfdf9811f657f07EA91b9d92366A80357dB52aaBA\t0x" + hex.EncodeToString(h256[:]))

	i := new(big.Int)
	i.SetBytes(addr[:])
	//i.SetString(colCell11, 16)
	//fmt.Println(i.String())
	var b3 int64
	_, b3, c, _ = printK(i)
	//fmt.Println(a, b3, c)
	bitSet[b3].SetBit(bitSet[b3], int(c), 1)

	//fmt.Println(bitSet[b1], bitSet[b2], bitSet[b3])

	return b1

}
