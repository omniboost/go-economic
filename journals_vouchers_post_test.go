package economic_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestJournalsVouchersPost(t *testing.T) {
	client := client()
	req := client.NewJournalsVouchersPostRequest()
	req.PathParams().JournalNumber = 1

	body := req.RequestBody()
	body.AccountingYear.Year = "2019"
	body.Journal.JournalNumber = 1

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
