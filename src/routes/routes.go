package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	cf "api_hubspot_go/src/config"
)

func GetClients() (response []map[string]interface{}) {
	req, _ := http.NewRequest("GET", "https://api.hubapi.com/crm/v3/objects/contacts?limit=30&archived=false", nil)

	token := cf.LoadConfig()
	bearer := fmt.Sprintf("Bearer %s", token)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseApi map[string]interface{}
	data := string([]byte(body))

	erro := json.Unmarshal([]byte(data), &responseApi)
	if erro != nil {
		fmt.Println("Error al deserializar la respuesta:", err)
		return
	}

	var userData []map[string]interface{}

	result := responseApi["results"].([]interface{})
	for _, elem := range result {
		item := elem.(map[string]interface{})
		properties := item["properties"].(map[string]interface{})
		userData = append(userData, properties)
	}

	return userData
}

func GetClientId(id string) (response map[string]interface{}) {
	url := fmt.Sprintf("https://api.hubapi.com/crm/v3/objects/contacts/%s?archived=false", id)
	req, _ := http.NewRequest("GET", url, nil)

	token := cf.LoadConfig()
	bearer := fmt.Sprintf("Bearer %s", token)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseApi map[string]interface{}
	data := string([]byte(body))

	erro := json.Unmarshal([]byte(data), &responseApi)
	if erro != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}

	properties := responseApi["properties"].(map[string]interface{})
	fmt.Println("Propiedades:", properties)

	return properties
}
