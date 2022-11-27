package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
	"os"
)

// ------------------------- Get File Status --------------------------\\

func Get(c *fiber.Ctx) error {
	//Fetching query params
	arg := c.Query("arg")

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
	test := string(body)

	fmt.Println(test)
	getresp := test

	jsonData, err := json.Marshal(getresp)
	if err != nil {
		panic(err)
	}
	return c.SendString(string(jsonData))
}
