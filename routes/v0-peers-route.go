package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
	"zendx.io/P2P-Drive/models"
)

var peers models.SwarmPeersResponse

// -------------------------- Get # of Peer Connections --------------------------\\

func Status(c *fiber.Ctx) error {

	url := "http://127.0.0.1:5001/api/v0/swarm/peers"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	json.Unmarshal(body, &peers)
	num := len(peers.Peers)

	return c.JSON(num)
}
