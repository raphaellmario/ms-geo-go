package tecban

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetAll(latitude string , longitude string) string {
	url := "https://api.homol.buscaatm.24horasdigital.com.br/v3/localizacao?latitude=" +latitude+ "&longitude=" +longitude+ "&acessivel=0&lista=1&limite=10"
	print(url)
	
	resp := Get(url);

	return resp
}

func Find(id string) string {
	url := "https://api.homol.buscaatm.24horasdigital.com.br/v3/localizacaopc?pc=" + id
	
	resp := Get(url);

	return resp
}

func Get(url string) string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("x-api-key", "XaVmpwFgQU7sG9KZFbm50agQzgI5zw4satFE1FVJ")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	sb := string(body)

	return sb
}