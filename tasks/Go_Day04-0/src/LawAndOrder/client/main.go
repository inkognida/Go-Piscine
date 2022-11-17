package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/lizrice/secure-connections/utils"
	"io/ioutil"
	"net/http"
	"os"
)

func	GetClient() *http.Client{

	data, _ := ioutil.ReadFile("../ca/minica.pem")
	cp, _ := x509.SystemCertPool()
	cp.AppendCertsFromPEM(data)

	config := &tls.Config{
		RootCAs: cp,
		GetClientCertificate: utils.ClientCertReqFunc("", ""),
		VerifyPeerCertificate: utils.CertificateChains,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}
	return client
}

type RequestBody struct {
	Money		int `json:"money"`
	CandyType	string `json:"candyType"`
	CandyCount	int	`json:"candyCount"`
}

func main()  {
	client := GetClient()

	var Request RequestBody

	Request.Money = 20
	Request.CandyType = "AA"
	Request.CandyCount = 1

	jsonValue, _ := json.Marshal(Request)

	resp, err := client.Post("https://localhost:8080/buy_candy", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	defer resp.Body.Close()
}