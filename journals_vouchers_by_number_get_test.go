package economic_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestJournalsVouchersByNumberGet(t *testing.T) {
	client := client()
	req := client.NewJournalsVouchersByNumberGetRequest()
	req.PathParams().JournalNumber = 965
	req.PathParams().AccountingYear = 2019
	req.PathParams().VoucherNumber = 6731

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
