package economic

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) NewJournalsVouchersByNumberGetRequest() JournalsVouchersByNumberGetRequest {
	return JournalsVouchersByNumberGetRequest{
		client:      c,
		queryParams: c.NewJournalsVouchersByNumberGetQueryParams(),
		pathParams:  c.NewJournalsVouchersByNumberGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewJournalsVouchersByNumberGetRequestBody(),
	}
}

type JournalsVouchersByNumberGetRequest struct {
	client      *Client
	queryParams *JournalsVouchersByNumberGetQueryParams
	pathParams  *JournalsVouchersByNumberGetPathParams
	method      string
	headers     http.Header
	requestBody JournalsVouchersByNumberGetRequestBody
}

func (c *Client) NewJournalsVouchersByNumberGetQueryParams() *JournalsVouchersByNumberGetQueryParams {
	return &JournalsVouchersByNumberGetQueryParams{}
}

type JournalsVouchersByNumberGetQueryParams struct {
}

func (r JournalsVouchersByNumberGetRequest) RequiredProperties() []string {
	return []string{}
}

func (r JournalsVouchersByNumberGetRequest) FilterableProperties() []string {
	return []string{}
}

func (r JournalsVouchersByNumberGetRequest) SortableProperties() []string {
	return []string{}
}

func (p JournalsVouchersByNumberGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *JournalsVouchersByNumberGetRequest) QueryParams() *JournalsVouchersByNumberGetQueryParams {
	return r.queryParams
}

func (c *Client) NewJournalsVouchersByNumberGetPathParams() *JournalsVouchersByNumberGetPathParams {
	return &JournalsVouchersByNumberGetPathParams{}
}

type JournalsVouchersByNumberGetPathParams struct {
	JournalNumber  int
	AccountingYear string
	VoucherNumber  int
}

func (p *JournalsVouchersByNumberGetPathParams) Params() map[string]string {
	return map[string]string{
		"journal_number":  fmt.Sprint(p.JournalNumber),
		"accounting_year": fmt.Sprint(p.AccountingYear),
		"voucher_number":  fmt.Sprint(p.VoucherNumber),
	}
}

func (r *JournalsVouchersByNumberGetRequest) PathParams() *JournalsVouchersByNumberGetPathParams {
	return r.pathParams
}

func (r *JournalsVouchersByNumberGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalsVouchersByNumberGetRequest) Method() string {
	return r.method
}

func (s *Client) NewJournalsVouchersByNumberGetRequestBody() JournalsVouchersByNumberGetRequestBody {
	return JournalsVouchersByNumberGetRequestBody{}
}

type JournalsVouchersByNumberGetRequestBody struct {
}

func (r *JournalsVouchersByNumberGetRequest) RequestBody() *JournalsVouchersByNumberGetRequestBody {
	return &r.requestBody
}

func (r *JournalsVouchersByNumberGetRequest) SetRequestBody(body JournalsVouchersByNumberGetRequestBody) {
	r.requestBody = body
}

func (r *JournalsVouchersByNumberGetRequest) NewResponseBody() *JournalsVouchersByNumberGetResponseBody {
	return &JournalsVouchersByNumberGetResponseBody{}
}

type JournalsVouchersByNumberGetResponseBody struct {
	AccountingYear struct {
		Year string `json:"year"`
		Self string `json:"self"`
	} `json:"accountingYear"`
	Journal struct {
		JournalNumber int    `json:"journalNumber"`
		Self          string `json:"self"`
	} `json:"journal"`
	Entries struct {
		FinanceVouchers []struct {
			Account struct {
				AccountNumber int    `json:"accountNumber"`
				Self          string `json:"self"`
			} `json:"account"`
			VatAccount struct {
				VatCode string `json:"vatCode"`
				Self    string `json:"self"`
			} `json:"vatAccount,omitempty"`
			VatAmount             float64 `json:"vatAmount"`
			VatAmountBaseCurrency float64 `json:"vatAmountBaseCurrency"`
			Text                  string  `json:"text"`
			Journal               struct {
				JournalNumber int    `json:"journalNumber"`
				Self          string `json:"self"`
			} `json:"journal"`
			Amount   float64 `json:"amount"`
			Currency struct {
				Code string `json:"code"`
				Self string `json:"self"`
			} `json:"currency"`
			Date         string  `json:"date"`
			ExchangeRate float64 `json:"exchangeRate"`
			EntryType    string  `json:"entryType"`
			Voucher      struct {
				AccountingYear struct {
					Year string `json:"year"`
					Self string `json:"self"`
				} `json:"accountingYear"`
				VoucherNumber int    `json:"voucherNumber"`
				Attachment    string `json:"attachment"`
				Self          string `json:"self"`
			} `json:"voucher"`
			AmountDefaultCurrency    float64 `json:"amountDefaultCurrency"`
			Remainder                float64 `json:"remainder"`
			RemainderDefaultCurrency float64 `json:"remainderDefaultCurrency"`
			JournalEntryNumber       int     `json:"journalEntryNumber"`
			MetaData                 struct {
				Delete struct {
					Description string `json:"description"`
					Href        string `json:"href"`
					HTTPMethod  string `json:"httpMethod"`
				} `json:"delete"`
			} `json:"metaData"`
			Self string `json:"self"`
		} `json:"financeVouchers"`
		ManualCustomerInvoices []struct {
			Customer struct {
				CustomerNumber int    `json:"customerNumber"`
				Self           string `json:"self"`
			} `json:"customer"`
			CustomerInvoice int    `json:"customerInvoice"`
			DueDate         string `json:"dueDate"`
			Templates       struct {
				CustomerPayment string `json:"customerPayment"`
				Self            string `json:"self"`
			} `json:"templates"`
			Text    string `json:"text"`
			Journal struct {
				JournalNumber int    `json:"journalNumber"`
				Self          string `json:"self"`
			} `json:"journal"`
			Amount   float64 `json:"amount"`
			Currency struct {
				Code string `json:"code"`
				Self string `json:"self"`
			} `json:"currency"`
			Date         string  `json:"date"`
			ExchangeRate float64 `json:"exchangeRate"`
			EntryType    string  `json:"entryType"`
			Voucher      struct {
				AccountingYear struct {
					Year string `json:"year"`
					Self string `json:"self"`
				} `json:"accountingYear"`
				VoucherNumber int    `json:"voucherNumber"`
				Attachment    string `json:"attachment"`
				Self          string `json:"self"`
			} `json:"voucher"`
			AmountDefaultCurrency    float64 `json:"amountDefaultCurrency"`
			Remainder                float64 `json:"remainder"`
			RemainderDefaultCurrency float64 `json:"remainderDefaultCurrency"`
			JournalEntryNumber       int     `json:"journalEntryNumber"`
			MetaData                 struct {
				Delete struct {
					Description string `json:"description"`
					Href        string `json:"href"`
					HTTPMethod  string `json:"httpMethod"`
				} `json:"delete"`
			} `json:"metaData"`
			Self string `json:"self"`
		} `json:"manualCustomerInvoices"`
	} `json:"entries"`
	VoucherNumber int    `json:"voucherNumber"`
	Attachment    string `json:"attachment"`
	Self          string `json:"self"`
}

func (r *JournalsVouchersByNumberGetRequest) URL() (url.URL, error) {
	return r.client.GetEndpointURL("journals-experimental/{{.journal_number}}/vouchers/{{.accounting_year}}-{{.voucher_number}}", r.PathParams())
}

func (r *JournalsVouchersByNumberGetRequest) Do() (JournalsVouchersByNumberGetResponseBody, error) {
	u, err := r.URL()
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), u, nil)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
