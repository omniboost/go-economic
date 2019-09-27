package economic_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCustomersPost(t *testing.T) {
	client := client()
	req := client.NewCustomersPostRequest()

	// body := req.RequestBody()

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
