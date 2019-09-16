package economic_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestAccountsGet(t *testing.T) {
	client := client()
	req := client.NewAccountsGetRequest()

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
