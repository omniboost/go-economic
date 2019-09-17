package economic_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestAccountingYearsEntriesGet(t *testing.T) {
	client := client()
	req := client.NewAccountingYearsEntriesGetRequest()
	req.PathParams().AccountingYear = "2018/2019"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
