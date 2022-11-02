package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"zendx.io/P2P-Drive/models"
)

var getresp models.GetResponse

func Get(arg string) models.GetResponse {
	url := "http://127.0.0.1:5001/api/v0/block/get?arg="
	endpoint := url + arg
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, endpoint, nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	test:= string(body)

	fmt.Println(test)
	getresp.Content = test

	return getresp
}
