package economic_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestJournalsEntriesGet(t *testing.T) {
	client := client()
	req := client.NewJournalsEntriesGetRequest()
	req.PathParams().JournalNumber = 8

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
