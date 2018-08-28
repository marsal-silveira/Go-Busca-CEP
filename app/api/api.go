package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"cep-provider/app/api/statuscode"
	"cep-provider/app/model"
	"cep-provider/app/provider"

	"github.com/gin-gonic/gin"
	// "github.com/subosito/gotenv"
)

// TODO : move to another place...
const notFoundMessage = `
404
¯\_(ツ)_/¯
`

// ConfigureServer TODO:
func ConfigureServer() *http.Server {

	// gotenv.Load(".env")

	if os.Getenv("GIN_MODE") == "test" {
		gin.SetMode(gin.TestMode)
	} else if os.Getenv("GIN_MODE") == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(gin.ErrorLogger())

	router.NoRoute(notFoundEndpoint)
	router.GET("/api/v1/cep/:cep", cepEndpoint)

	// port := os.Getenv("PORT")
	port := "8081"
	fmt.Println("starting server on", port)

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return server
}

// 404 error showing start page
func notFoundEndpoint(c *gin.Context) {
	c.String(404, notFoundMessage)
}

func cepEndpoint(c *gin.Context) {

	c.Header("Content-Type", "application/json; charset=utf-8")
	cep := beautify(c.Param("cep"))

	statusCode, address := provider.GetAddress(cep)
	fmt.Println("Status Code", statusCode)

	switch statusCode {
	case statuscode.OK:
		addressJSON, err := encode(address)
		if err != nil {
			fmt.Println("apiCepJSON error:", err)
			c.String(statuscode.ToInt(statuscode.InternalServerError), err.Error())
		} else {
			fmt.Println("apiCepJSON address", addressJSON)
			c.String(statuscode.ToInt(statusCode), addressJSON)
		}
	case statuscode.UnprocessableEntity:
		c.String(statuscode.ToInt(statusCode), "CEP `"+cep+"` is invalid")
	default:
		c.Status(statuscode.ToInt(statusCode))
	}
}

// encode Address to JSON string to be used on response
func encode(address *model.Address) (addressJSON string, err error) {

	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		fmt.Printf("RegExp error: %s", err)
		return "", err
	}

	JSONConvert := &model.Address{
		Provider:    address.Provider,
		Cep:         reg.ReplaceAllString(address.Cep, ""),
		Logradouro:  address.Logradouro,
		Bairro:      address.Bairro,
		Complemento: address.Complemento,
		Cidade:      address.Cidade,
		Estado:      address.Estado,
		Ibge:        address.Ibge,
	}

	conv, err := json.MarshalIndent(JSONConvert, "", "  ")
	if err != nil {
		fmt.Printf("apiWriteJSON error: %s", err)
		return "", err
	}

	return string(conv), err
}

// remeve mask from CEP
func beautify(cep string) string {

	result := strings.Trim(cep, " ")
	result = strings.Replace(result, ".", "", -1)
	result = strings.Replace(result, "-", "", -1)

	return result
}
