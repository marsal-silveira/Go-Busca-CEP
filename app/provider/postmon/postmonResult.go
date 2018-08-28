package postmon

// postmonAddressEstadoInfo : TODO
type postmonAddressEstadoInfo struct {
	AreaKm2    string `json:"area_km2"`
	CodigoIBGE string `json:"codigo_ibge"`
	Nome       string `json:"nome"`
}

// postmonAddressCidadeInfo : TODO
type postmonAddressCidadeInfo struct {
	AreaKm2    string `json:"area_km2"`
	CodigoIBGE string `json:"codigo_ibge"`
}

// postmonAddress : TODO
type postmonAddress struct {
	Cep        string                   `json:"cep"`
	Logradouro string                   `json:"logradouro"`
	Bairro     string                   `json:"bairro"`
	Cidade     string                   `json:"cidade"`
	CidadeInfo postmonAddressCidadeInfo `json:"cidade_info"`
	Estado     string                   `json:"estado"`
	EstadoInfo postmonAddressEstadoInfo `json:"estado_info"`
}

// postmonError : TODO
type postmonError struct {
	Error bool `json:"erro"`
}
