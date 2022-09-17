package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/xuri/excelize/v2"
	"strings"
)

func getCorrectName(p string) string {
	iben := 0
	var r []byte
	idx := 0
	for i := 0; i < len(p); i++ {
		if (p[i] > 97 && p[i] < 122) || p[i] == 45 {
			if iben == 0 {
				iben = i
			}
			r = append(r, p[i])
			idx = i
		}
	}

	if (idx - iben + 1) > len(r) {
		return ""
	}

	return string(r)
}

func main() {
	f, err := excelize.OpenFile("reserve.xlsx")
	if err != nil {
		println(err.Error())
		return
	}

	rows, err := f.GetRows("Sheet1")
	cnt := 0

	msg := "[\""

	msg1 := "[\""

	for _, row := range rows {
		cnt++
		//for _, colCell := range row {
		//fmt.Println()
		lows := getCorrectName(strings.ToLower(row[0]))
		if len(lows) == 0 {
			continue
		}

		if len(msg1) > 2 {
			msg1 += "\",\""
			msg1 += lows
		} else {
			msg1 += lows
		}
		//lows := strings.Trim(strings.ToLower(row[0]), " ")
		//fmt.Println(lows, crypto.Keccak256Hash([]byte(lows)))
		h := crypto.Keccak256Hash([]byte(lows))
		s := h.String()
		if len(msg) > 20 {
			msg += "\",\""
			msg += s
		} else {
			msg += s
		}
	}

	msg += "\"]"
	msg1 += "\"]"

	fmt.Println(msg)

	fmt.Println(cnt)
}
