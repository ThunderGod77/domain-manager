package database

import (
	"encoding/json"
	"errors"
	"os"
)

type Domain struct {
	DomainName  string `json:"domainName"`
	Provider    string `json:"provider"`
	Description string `json:"description"`
}

var Providers = []string{"route53", "namecheap", "godaddy", "googleDomains"}

func verifyProvider(provider string) bool {
	for i := 0; i < len(Providers); i++ {
		if provider == Providers[i] {
			return true
		}
	}
	return false
}

func GetDomains() (map[string]Domain, error) {
	data, err := os.OpenFile("./domain.json", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	stat, err := data.Stat()
	if err != nil {
		return nil, err
	}
	if stat.Size() == 0 {
		return map[string]Domain{}, nil
	}

	var domains map[string]Domain
	err = json.NewDecoder(data).Decode(&domains)

	if err != nil {
		return nil, err
	}
	return domains, nil
}

func updateDb(domains map[string]Domain) error {
	marshal, err := json.Marshal(domains)
	if err != nil {
		return err
	}

	err = os.WriteFile("./domain.json", marshal, 0644)
	if err != nil {
		return err
	}
	return nil
}

func AddDomain(domainName, provider, description string) error {
	domains, err := GetDomains()
	if err != nil {
		return err
	}

	if !verifyProvider(provider) {
		return errors.New("provider not present")
	}
	_, pres := domains[domainName]
	if pres {
		return errors.New("domain is already exists - please use update command")
	}
	domains[domainName] = Domain{
		DomainName:  domainName,
		Provider:    provider,
		Description: description,
	}
	err = updateDb(domains)
	if err != nil {
		return err
	}

	return nil
}

func UpdateOrAddDomain(domainName, provider, description string) error {
	domains, err := GetDomains()
	if err != nil {
		return err
	}

	if !verifyProvider(provider) {
		return errors.New("provider not present")
	}
	domains[domainName] = Domain{
		DomainName:  domainName,
		Provider:    provider,
		Description: description,
	}
	err = updateDb(domains)
	if err != nil {
		return err
	}

	return nil
}

//providers
//name cheap -  api username and api key
//route 53 - access key and secret access key
//go daddy - production key and secret

type ProviderKeys struct {
	Provider  string `json:"provider"`
	AccessKey string `json:"accessKey"`
	Secret    string `json:"secret"`
}

func GetProviders() (map[string]ProviderKeys, error) {
	data, err := os.OpenFile("./provider.json", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	stat, err := data.Stat()
	if err != nil {
		return nil, err
	}
	if stat.Size() == 0 {
		return map[string]ProviderKeys{}, nil
	}

	var providers map[string]ProviderKeys
	err = json.NewDecoder(data).Decode(&providers)

	if err != nil {
		return nil, err
	}
	return providers, nil
}

func updateProviderKeys(providers map[string]ProviderKeys) error {
	marshal, err := json.Marshal(providers)
	if err != nil {
		return err
	}

	err = os.WriteFile("./provider.json", marshal, 0644)
	if err != nil {
		return err
	}
	return nil
}

func AddProvider(provider, accessKey, secret string) error {
	providers, err := GetProviders()
	if err != nil {
		return err
	}
	if !verifyProvider(provider) {
		return errors.New("not a valid provider")
	}
	providers[provider] = ProviderKeys{
		Provider:  provider,
		AccessKey: accessKey,
		Secret:    secret,
	}
	err = updateProviderKeys(providers)
	if err != nil {
		return err
	}
	return nil
}

func GetKey(provider string) (*ProviderKeys, error) {
	if !verifyProvider(provider) {
		return nil, errors.New("invalid provider")
	}
	providers, err := GetProviders()
	if err != nil {
		return nil, err
	}
	val, ok := providers[provider]
	if !ok {
		return nil, errors.New("provider keys not present")
	}
	return &ProviderKeys{
		Provider:  val.Provider,
		AccessKey: val.AccessKey,
		Secret:    val.Secret,
	}, nil
}

func GetCredentials(domain string) (*ProviderKeys, error) {
	domains, err := GetDomains()
	if err != nil {
		return nil, err
	}

	domainData, ok := domains[domain]
	if !ok {
		return nil, err
	}
	providers, err := GetProviders()
	if err != nil {
		return nil, err
	}
	providerData, ok := providers[domainData.Provider]
	if !ok {
		return nil, err
	}

	accessKey := providerData.AccessKey
	secret := providerData.Secret
	return &ProviderKeys{AccessKey: accessKey, Secret: secret}, nil
}
