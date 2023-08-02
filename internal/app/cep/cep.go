package cep

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ViaCEP = "Via CEP"
	ApiCEP = "Api CEP"
)

type Output struct {
	Provider string
	Response interface{}
}

func (o *Output) ToJSON() string {
	data, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return ""
	}
	return string(data)
}

type ViaCEPResponse struct {
	CEP         string `json:"cep"`
	Logradouro  string
	Complemento string
	Bairro      string
	Localidade  string
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

type ApiCEPResponse struct {
	Code       string
	State      string
	City       string
	District   string
	Address    string
	Status     string
	Ok         bool
	StatusText string
}

func buildURL(provider string, term string) (string, error) {
	if provider == ViaCEP {
		return fmt.Sprintf("https://viacep.com.br/ws/%s/json/", term), nil
	}
	if provider == ApiCEP {
		return fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", term), nil
	}
	return "", fmt.Errorf("invalid provider: %s", provider)
}

func Search(provider string, term string) (*Output, error) {
	url, err := buildURL(provider, term)
	if err != nil {
		return nil, err
	}
	var res *http.Response
	res, err = http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("searching for CEP %s from %s", term, provider)
	}
	var response interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &Output{
		Provider: provider,
		Response: response,
	}, nil
}
