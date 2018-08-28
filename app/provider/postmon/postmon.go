package postmon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"cep-provider/app/api/statuscode"
	"cep-provider/app/model"
)

// Postmon : represents the Postmon provider
type Postmon struct {
}

// Name : TODO
func (provider Postmon) Name() string {
	return "Postmon"
}

// GetAddress : TODO
func (provider Postmon) GetAddress(cep string) (statusCode statuscode.StatusCode, address *model.Address) {

	// default values
	statusCodeReturn := statuscode.InternalServerError
	var addressReturn *model.Address

	url := fmt.Sprintf("http://api.postmon.com.br/v1/cep/%s", cep)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Postmon.GetAddress [#1] http.Get() error:", err)
		return statuscode.InternalServerError, addressReturn
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		// if status code is different of 200 meaning that `cep` is invalid so we'll return a `Unprocessable Entity`
		fmt.Println("Postmon.GetAddress [#2] status code:", resp.StatusCode)

		if resp.StatusCode == 404 {
			return statuscode.NoContent, addressReturn
		}
		return statuscode.UnprocessableEntity, addressReturn
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Postmon.GetAddress [#3] ReadAll error:", err)
		return statuscode.InternalServerError, addressReturn
	}

	// get postmon address from response content
	var postmonAddress *postmonAddress
	err = unmarshal(content, &postmonAddress)
	if err != nil {
		fmt.Println("Postmon.GetAddress unmarshal error:", err)
		return statuscode.InternalServerError, addressReturn
	} else if postmonAddress == nil {

		fmt.Println("Cep not found")
		return statuscode.NoContent, addressReturn
	}

	//map json to final string
	statusCodeReturn = statuscode.OK
	addressReturn = parseToAddress(postmonAddress)

	return statusCodeReturn, addressReturn
}

func unmarshal(data []byte, v interface{}) error {

	// first check if `cep` has an address associated
	var postmonError postmonError
	err := json.Unmarshal(data, &postmonError)
	if err != nil {
		return err
	}

	if postmonError.Error {
		return nil
	}

	// if no errors... unmarshal postmon address
	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}

func parseToAddress(postmonAddress *postmonAddress) *model.Address {

	var address model.Address

	address.Provider = "Postmon"
	address.Cep = postmonAddress.Cep
	address.Logradouro = postmonAddress.Logradouro
	address.Bairro = postmonAddress.Bairro
	address.Complemento = ""
	address.Cidade = postmonAddress.Cidade
	address.Estado = postmonAddress.Estado
	address.Ibge = postmonAddress.CidadeInfo.CodigoIBGE

	return &address
}
