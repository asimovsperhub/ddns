package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CallNftPass(tokenid, name string) {
	client := &http.Client{}
	//var data = strings.NewReader(fmt.Sprintf(`{"operationName":"getNamesFromSubgraph","variables":{"address":"%s"},"query":"query getNamesFromSubgraph($address: String!) {\n  domains(first: 1000, where: {resolvedAddress: $address}) {\n    name\n    __typename\n  }\n}\n"}`, addr))
	req, err := http.NewRequest("GET", "https://nftmetadata.udid.network/api/gettokenname", nil)
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("tokenid", tokenid)
	q.Add("name", name)
	req.URL.RawQuery = q.Encode()
	log.Println(req.URL.String())
	//req.Header.Set("authority", "api.thegraph.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("content-type", "application/json")
	//req.Header.Set("origin", "https://app.ens.domains")
	//req.Header.Set("referer", "https://app.ens.domains/")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	res := make(map[string]string)
	json.Unmarshal(bodyText, &res)
	log.Println("CallNftPass  :", res)
}
