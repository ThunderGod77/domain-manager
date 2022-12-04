package goDaddy

import (
	"fmt"

	"testing"
)

func TestAddRecords(t *testing.T) {

	accessKey, secret, domain := GetEnv()
	fmt.Println(domain)
	records, err := GetRecords(accessKey, secret, "", "", domain)
	if err != nil || len(records) == 0 {
		t.Fatal("function get records failed", err)
	}

}
