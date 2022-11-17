package main

import (
	tls2 "crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/lizrice/secure-connections/utils"
)


type SuccessBuy struct { // 201 status
	Thanks	string
	Change	int
}

type FailBuy struct { // 400 status
	Error	string
}

type NoMoney struct { // 401 status
	Error	string
}

//Cool Eskimo: 10 cents
//Apricot Aardvark: 15 cents
//Natural Tiger: 17 cents
//Dazzling 	Elderberry: 21 cents
//Yellow Rambutan: 23 cents

type RequestBody struct {
	Money		int `json:"money"`
	CandyType	string `json:"candyType"`
	CandyCount	int	`json:"candyCount"`
}

var CandyTypes =  map[string]int{
	"CK": 10,
	"AA": 15,
	"NT": 17,
	"DE": 21,
	"YR": 23,
}

func	BuyCandy(w http.ResponseWriter, r *http.Request) {
	var Request RequestBody
	err := json.NewDecoder(r.Body).Decode(&Request)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(FailBuy{
			Error: "Wrong request",
		})
		return
	}

	CandyTypePrice := CandyTypes[Request.CandyType]
	Money := Request.Money
	CandyCount := Request.CandyCount

	switch {
	case CandyTypePrice == 0:
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(FailBuy{Error: fmt.Sprintf("No such candies")})
		return
	case Money < 0:
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(FailBuy{Error: fmt.Sprintf("You need some money")})
		return
	case CandyCount < 0:
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(FailBuy{Error: fmt.Sprintf("But one at least")})
		return
	}

	switch {
	case CandyTypePrice * CandyCount <= Money:
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(SuccessBuy{
			Thanks: "Thank you!",
			Change: Money - CandyTypePrice * CandyCount,
		})
		return
	case CandyTypePrice * CandyCount > Money:
		w.WriteHeader(402)
		json.NewEncoder(w).Encode(NoMoney{Error: fmt.Sprintf("You need %d more money",
			CandyTypePrice * CandyCount - Money)})
		return
	}
}

func GetServer() *http.Server{
	tls := &tls2.Config{

		GetCertificate: utils.CertReqFunc("cert.pem", "key.pem"),
		VerifyPeerCertificate: utils.CertificateChains,
	}

	server := &http.Server{
		Addr: ":8080",
		TLSConfig: tls,
	}

	return server
}

func main() {
	server := GetServer()
	http.HandleFunc("/buy_candy", BuyCandy)
	server.ListenAndServeTLS("", "")
}

