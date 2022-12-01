package routes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	models "zendx.io/P2P-Drive/models"
)

var userRegister models.RegisterRequest

// -------------------------- Register User --------------------------\\

func UserRegister(c *fiber.Ctx) error {

	json.Unmarshal(c.Body(), &userRegister)

	//fmt.Print(uuid.New().String()[:31])

	//return c.JSON(user)

	Database := Connection()

	val := Database.DBemailCheck(userRegister.Email)

	if val == "Not Found" {
		Database.DBregister(&userRegister)
		return c.JSON(userRegister)

	} else {
		return c.SendString("Account with email exists")
	}
}

func encrypt(data []byte, password string) []byte {
	block, _ := aes.NewCipher([]byte(password))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}
