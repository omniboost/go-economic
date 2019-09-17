package economic_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/omniboost/go-economic"
)

func TestInvoicesDraftsPost(t *testing.T) {
	client := client()
	req := client.NewInvoicesDraftsPostRequest()

	body := req.RequestBody()
	body.Date = economic.Date{time.Now()}
	body.Currency = "EUR"
	body.Customer.CustomerNumber = 1
	body.PaymentTerms.PaymentTermsNumber = 14
	// body.PaymentTerms.PaymentTermsType = "dueDate"
	body.Recipient = economic.InvoiceDraftRecipient{
		Name: "Kees Zorge",
		VatZone: economic.InvoiceDraftVatZone{
			VatZoneNumber: 1,
		},
	}
	body.Layout = economic.InvoiceDraftLayout{
		LayoutNumber: 19,
	}
	line := economic.InvoiceDraftLine{}
	body.Lines = append(body.Lines, line)

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
