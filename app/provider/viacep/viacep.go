package viacep

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"cep-provider/app/api/statuscode"
	"cep-provider/app/model"
)

// ViaCEP : represents the ViaCEP provider
type ViaCEP struct {
}

// Name : TODO
func (provider ViaCEP) Name() string {
	return "ViaCEP"
}

// GetAddress : TODO
func (provider ViaCEP) GetAddress(cep string) (statusCode statuscode.StatusCode, address *model.Address) {

	// default values
	statusCodeReturn := statuscode.InternalServerError
	var addressReturn *model.Address

	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ViaCep.GetAddress [#1] http.Get() error:", err)
		return statuscode.InternalServerError, addressReturn
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		// if status code is different of 200 meaning that `cep` is invalid so we'll return a `Unprocessable Entity`
		fmt.Println("ViaCep.GetAddress [#2] status code:", resp.StatusCode)
		return statuscode.UnprocessableEntity, addressReturn
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ViaCep.GetAddress [#3] ReadAll error:", err)
		return statuscode.InternalServerError, addressReturn
	}

	// get viaCep address from response content
	var viacepAddress *viacepAddress
	err = unmarshal(content, &viacepAddress)
	if err != nil {
		fmt.Println("ViaCep.GetAddress unmarshal error:", err)
		return statuscode.InternalServerError, addressReturn
	} else if viacepAddress == nil {

		fmt.Println("Cep not found")
		return statuscode.NoContent, addressReturn
	}

	//map json to final string
	statusCodeReturn = statuscode.OK
	addressReturn = parseToAddress(viacepAddress)

	return statusCodeReturn, addressReturn
}

func unmarshal(data []byte, v interface{}) error {

	// first check if `cep` has an address associated
	var viacepError viacepError
	err := json.Unmarshal(data, &viacepError)
	if err != nil {
		return err
	}

	if viacepError.Error {
		return nil
	}

	// if no errors... unmarshal viacep address
	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}

func parseToAddress(viaCepAddress *viacepAddress) *model.Address {

	var address model.Address

	address.Provider = "ViaCep"
	address.Cep = viaCepAddress.Cep
	address.Logradouro = viaCepAddress.Logradouro
	address.Bairro = viaCepAddress.Bairro
	address.Complemento = viaCepAddress.Complemento
	address.Cidade = viaCepAddress.Localidade
	address.Estado = viaCepAddress.Uf
	address.Ibge = viaCepAddress.Ibge

	return &address
}
