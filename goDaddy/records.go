package goDaddy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"net/http"
	"os"
)

type InputRecord struct {
	Data     string  `json:"data"`
	Name     string  `json:"name"`
	Port     int     `json:"port"`
	Priority int     `json:"priority"`
	Protocol string  `json:"protocol"`
	Service  string  `json:"service"`
	Ttl      int     `json:"ttl"`
	Type     string  `json:"type"`
	Weight   float64 `json:"weight"`
}

var baseUrl = "https://api.godaddy.com/v1"

type Record struct {
	Data string `json:"data"`
	Name string `json:"name"`
	Ttl  int    `json:"ttl"`
	Type string `json:"type"`
}

func GetEnv() (string, string, string) {
	accessKey := os.Getenv("GODADDY_ACCESS_KEY")
	secret := os.Getenv("GO_DADDY_SECRET")
	domain := os.Getenv("TEST_DOMAIN")
	return accessKey, secret, domain
}

func GetRecords(accessKey, secret, recordType, name, domain string) ([]Record, error) {
	if domain == "" {
		return nil, errors.New("please mention a correct domain")
	}
	var url = baseUrl + "/domains/" + domain + "/records"
	if recordType != "" {
		url = url + "/" + recordType
	}
	if name != "" {
		url = url + "/" + name
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("sso-key %s:%s", accessKey, secret))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var r []Record
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func AddRecord(accessKey, secret, domain, data, name, recordType string, weight float64, ttl, priority int) error {
	var url = baseUrl + "domains/" + domain + "/records"
	var ir = InputRecord{
		Data:     data,
		Name:     name,
		Port:     65535,
		Priority: priority,
		Protocol: "",
		Service:  "",
		Ttl:      ttl,
		Type:     recordType,
		Weight:   weight,
	}
	marshal, err := json.Marshal(ir)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(marshal))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("sso-key %s:%s", accessKey, secret))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.Status != http.StatusText(http.StatusOK) {
		return errors.New("some error occurred")
	}
	return nil

}

func DeleteRecord(accessKey, secret, recordType, name, domain string) error {
	if recordType == "" || domain == "" || name == "" {
		return errors.New("please enter a valid record type,name or domain")
	}
	var url = baseUrl + "/domains/" + domain + "/records/" + recordType + "/" + name
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("sso-key %s:%s", accessKey, secret))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return errors.New("some error occurred")
	}
	return nil

}
