package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	models "zendx.io/P2P-Drive/models"
)

var ipfs string = "https://ipfs.io/ipfs/"

// Utilizing Models
var request models.AddRequest
var response models.AddResponse

// -------------------------- Register User --------------------------\\

func Add(c *fiber.Ctx) error {

	// Getting form data from request
	owner := c.Query("Token")
	file, err := c.FormFile("File")
	if err != nil {
		os.Exit(0)
	}

	request.FileReceived = file
	request.Id = file.Filename
	//Utilizing IPFS node API
	url := "http://127.0.0.1:5001/api/v0/add"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	openFile, _ := file.Open()
	formfile, errFile1 := writer.CreateFormFile("file", file.Filename)

	//---------- Encryption here ----------

	_, errFile1 = io.Copy(formfile, openFile)
	if errFile1 != nil {
		fmt.Println("I/O error: ", errFile1)
		os.Exit(0)
	}
	err1 := writer.Close()
	if err1 != nil {
		fmt.Println("multipart/writer error: ", err1)
		os.Exit(0)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println("err")
		os.Exit(0)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error Setting Header")
		os.Exit(0)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error Reading Body")
		os.Exit(0)
	}

	// Unmarshalling JSON response into Addresponse struct
	json.Unmarshal(body, &response)

	//creating IPFS link for later use
	data := ipfs + response.Hash + "/"

	fmt.Println(data)
	fmt.Println(string(body))

	response.Link = data
	response.Owner = owner

	//Upload model to DB
	Connection().DBupload(response)

	jsonData, _ := json.Marshal(response)

	//returning JSON response through API
	return c.SendString(string(jsonData))
}
