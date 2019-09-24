package economic_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestCustomersGet(t *testing.T) {
	client := client()
	req := client.NewCustomersGetRequest()
	filter := fmt.Sprintf("name$eq:%s$or:name$eq:%s", "Bravo Tours", "leon@omniboost.io")
	req.QueryParams().Filter.Set(filter)

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
