package model

// Address is the Brazilian address type
type Address struct {
	Provider    string `json:"provider"`
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Cidade      string `json:"cidade"`
	Estado      string `json:"estado"`
	Ibge        string `json:"ibge"`
}
