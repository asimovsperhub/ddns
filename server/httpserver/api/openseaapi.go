package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dnsdao/dnsdao.resolver/config"
	"github.com/dnsdao/dnsdao.resolver/database/ldb"
	"log"
	"net/http"
	"strings"
)

func splitPath(p string) (string, string, error) {
	sarr := strings.Split(p, "/")

	if len(sarr) == 0 {
		return "", "", errors.New("path error")
	}
	stokenId := sarr[len(sarr)-1]
	addr := sarr[len(sarr)-2]
	//tokenId, err := strconv.Atoi(stokenId)
	//if err != nil {
	//	return err, -1
	//}
	return addr, stokenId, nil
}

type TokenIDRes struct {
	Name          string `json:"name"`
	Image         string `json:"image"`
	External_Url  string `json:"external_url"`
	Animation_Url string `json:"animation_url"`
	Description   string `json:"description"`
	//Image_Data    string `json:"image_data"`
}

func (a *Api) DnsTokenId(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "not a get request")
		return
	}

	nametype := ""
	msg := "not found"
	addr, tokenid, err := splitPath(request.URL.Path)
	if err != nil {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "token Id error")
		return
	}
	fmt.Println(addr, tokenid)
	db := ldb.GetLdb()
	contractaddr := strings.ToLower(addr)
	tokenName, err := db.GetContractTokenIdName(contractaddr)
	if err != nil {
		writer.WriteHeader(500)
		fmt.Fprintf(writer, "contract addr not found")
		return
	}
	// root
	if contractaddr == strings.ToLower(config.GetRConf().Cconf.DnsName) {
		nametype = "root"
	} else {
		nametype = "sub"
	}
	if tokenName != nil {
		name := tokenName.TokenName[tokenid]
		if name != "" {
			res := new(TokenIDRes)
			res.Name = name
			res.Image = fmt.Sprintf("https://dnsdaonftpasscard.s3.ap-east-1.amazonaws.com/dnsname/%s_%s.png", contractaddr, tokenid)
			res.Description = fmt.Sprintf("Udid")
			respstr, errj := json.Marshal(res)
			if errj == nil {
				msg = string(respstr)
			}
			Ss3, _ := db.GetS3()
			if Ss3 == nil {
				saveimg := fmt.Sprintf("./opensea/%s_%s.png", contractaddr, tokenid)
				err1 := CreateImage(name, nametype, saveimg)
				if err1 != nil {
					log.Println("CreateImg err", err1)
				} else {
					raw, errrd := Read(saveimg)
					if errrd == nil {
						s3res, errs3 := UploadFileToS3(raw, fmt.Sprintf("%s_%s.png", contractaddr, tokenid))
						if errs3 != nil {
							log.Println("S3 Upload failed", errs3)
						} else {
							log.Println("s3 upload success", s3res)
							Ss3 = &ldb.S3{
								Upload: map[string]string{fmt.Sprintf("%s_%s", contractaddr, tokenid): "yes"},
							}
							db.SaveS3(Ss3)
						}
					}
				}

			}
			if Ss3.Upload[fmt.Sprintf("%s_%s", contractaddr, tokenid)] != "yes" {
				saveimg := fmt.Sprintf("./opensea/%s_%s.png", contractaddr, tokenid)
				err1 := CreateImage(name, nametype, saveimg)
				if err1 != nil {
					log.Println("CreateImg err", err1)
				}
				raw, errrd := Read(saveimg)
				if errrd == nil {
					s3res, errs3 := UploadFileToS3(raw, fmt.Sprintf("%s_%s.png", contractaddr, tokenid))
					if errs3 != nil {
						log.Println("S3 Upload failed", err)
					} else {
						log.Println("s3 upload success", s3res)
						Ss3.Upload[fmt.Sprintf("%s_%s", contractaddr, tokenid)] = "yes"
						db.SaveS3(Ss3)
					}
				}
			}
		}

	}
	writer.WriteHeader(200)
	writer.Write([]byte(msg))
}
