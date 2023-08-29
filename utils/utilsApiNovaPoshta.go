package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type NovaPoshtaRequest struct {
	ModelName        string                 `json:"modelName"`
	CalledMethod     string                 `json:"calledMethod"`
	MethodProperties map[string]interface{} `json:"methodProperties"`
	ApiKey           string                 `json:"apiKey"`
}

type NovaPoshtaResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func FindCityApiNovaPoshta(searchWord string) (interface{}, error) {
	requestBody := NovaPoshtaRequest{
		ModelName:    "Address",
		CalledMethod: "getCities",
		MethodProperties: map[string]interface{}{
			"FindByString": searchWord,
		},
		ApiKey: ApiTokenNovaPoshta,
	}

	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	url := "https://api.novaposhta.ua/v2.0/json/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestJSON))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var npResponse NovaPoshtaResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&npResponse)
	if err != nil && err != io.EOF {
		return nil, err
	}

	if npResponse.Success {
		return npResponse, nil
	} else {
		return nil, errors.New("not success")
	}
}

func GetWarehousesApiNovaPoshta(cityRef string) (interface{}, error) {
	requestBody := NovaPoshtaRequest{
		ModelName:    "Address",
		CalledMethod: "getWarehouses",
		MethodProperties: map[string]interface{}{
			"CityRef": cityRef,
		},
		ApiKey: ApiTokenNovaPoshta,
	}

	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	url := "https://api.novaposhta.ua/v2.0/json/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestJSON))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var npResponse NovaPoshtaResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&npResponse)
	if err != nil && err != io.EOF {
		return nil, err
	}

	if npResponse.Success {
		return npResponse, nil
	} else {
		return nil, errors.New("not success")
	}
}

func GetSettlementsApiNovaPoshta(cityRef string) (interface{}, error) {
	requestBody := NovaPoshtaRequest{
		ModelName:    "Address",
		CalledMethod: "getSettlements",
		MethodProperties: map[string]interface{}{
			"SettlementRef": cityRef,
		},
		ApiKey: ApiTokenNovaPoshta,
	}

	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	url := "https://api.novaposhta.ua/v2.0/json/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestJSON))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var npResponse NovaPoshtaResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&npResponse)
	if err != nil && err != io.EOF {
		return nil, err
	}

	if npResponse.Success {
		return npResponse, nil
	} else {
		return nil, errors.New("not success")
	}
}
