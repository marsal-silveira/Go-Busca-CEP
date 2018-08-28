package provider

import (
	"fmt"
	"regexp"
	"strings"

	"cep-provider/app/api/statuscode"
	"cep-provider/app/model"
	"cep-provider/app/provider/postmon"
	"cep-provider/app/provider/viacep"
)

// TODO: Review this...
// This has to much reponsibilty. Break it into two or more components with more specific concern.
// e.g. Provider Factory, Executing Manager/Coordinator, Parallelism, etc.

// Provider : TODO
type Provider interface {
	Name() string
	GetAddress(cep string) (statusCode statuscode.StatusCode, address *model.Address)
}

// requestResult : TODO
type requestResult struct {
	statusCode statuscode.StatusCode
	address    *model.Address
}

// internalGetAddress : TODO
func internalGetAddress(provider Provider, cep string, result chan requestResult) {

	fmt.Printf("provider -> %s, cep -> %s\n", provider.Name(), cep)
	go func() {
		statusCode, address := provider.GetAddress(cep)
		result <- requestResult{statusCode, address}
	}()
}

// GetAddress : TODO
func GetAddress(cep string) (statusCode statuscode.StatusCode, address *model.Address) {

	// first validate `CEP` value before continue
	if !validate(cep) {
		return statuscode.UnprocessableEntity, nil
	}

	//
	result := make(chan requestResult)

	// get all available providers and try to get address using a list of provider starting from first and go on until find some Address...
	providers := getProviders()
	for _, provider := range providers {
		internalGetAddress(provider, cep, result)
	}

	var _result requestResult
	for index := 0; index < len(providers); index++ {
		fmt.Print("#", index)

		_result = <-result
		fmt.Println(". result ->", _result.statusCode)
		if _result.statusCode == statuscode.OK {
			break
		}
	}

	// loop:
	// 	for {
	// 		fmt.Println("for")
	// 		select {
	// 		case <-result:

	// 			_result = <-result
	// 			fmt.Println("result ->", _result.statusCode)
	// 			if _result.statusCode == statuscode.OK {
	// 				close(result)
	// 				break loop
	// 			}
	// 		}
	// 	}

	return _result.statusCode, _result.address
}

func getProviders() []Provider {

	providers := []Provider{
		viacep.ViaCEP{},
		postmon.Postmon{},
	}
	return providers
}

// validate func validate if input `cep` has only numbers with 8 of size
func validate(cep string) bool {

	const cepSize int = 8

	re := regexp.MustCompile("[0-9]+")
	numbers := strings.Join(re.FindAllString(cep, -1), "")

	originalSize := len(cep)
	onlyNumbersSize := len(numbers)

	// log
	// fmt.Println(cep + " -> " + strconv.Itoa(originalSize))
	// fmt.Println(numbers + " -> " + strconv.Itoa(onlyNumbersSize))

	return originalSize == onlyNumbersSize && originalSize == cepSize
}
