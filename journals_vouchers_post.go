package economic

import (
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) NewJournalsVouchersPostRequest() JournalsVouchersPostRequest {
	return JournalsVouchersPostRequest{
		client:      c,
		queryParams: c.NewJournalsVouchersPostQueryParams(),
		pathParams:  c.NewJournalsVouchersPostPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewJournalsVouchersPostRequestBody(),
	}
}

type JournalsVouchersPostRequest struct {
	client      *Client
	queryParams *JournalsVouchersPostQueryParams
	pathParams  *JournalsVouchersPostPathParams
	method      string
	headers     http.Header
	requestBody JournalsVouchersPostRequestBody
}

func (c *Client) NewJournalsVouchersPostQueryParams() *JournalsVouchersPostQueryParams {
	return &JournalsVouchersPostQueryParams{}
}

type JournalsVouchersPostQueryParams struct {
}

func (r JournalsVouchersPostRequest) RequiredProperties() []string {
	return []string{
		"self",
		"settings.contraAccounts.customerPayments.self",
		"settings.contraAccounts.financeVouchers.self",
		"settings.contraAccounts.supplierPayments.self",
		"templates.self",
	}
}

func (r JournalsVouchersPostRequest) FilterableProperties() []string {
	return []string{
		"journalNumber",
		"name",
	}
}

func (r JournalsVouchersPostRequest) SortableProperties() []string {
	return []string{
		"journalNumber",
		"name",
	}
}

func (p JournalsVouchersPostQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *JournalsVouchersPostRequest) QueryParams() *JournalsVouchersPostQueryParams {
	return r.queryParams
}

func (c *Client) NewJournalsVouchersPostPathParams() *JournalsVouchersPostPathParams {
	return &JournalsVouchersPostPathParams{
		JournalNumber: 0,
	}
}

type JournalsVouchersPostPathParams struct {
	JournalNumber int
}

func (p *JournalsVouchersPostPathParams) Params() map[string]string {
	return map[string]string{
		"journal_number": strconv.Itoa(p.JournalNumber),
	}
}

func (r *JournalsVouchersPostRequest) PathParams() *JournalsVouchersPostPathParams {
	return r.pathParams
}

func (r *JournalsVouchersPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalsVouchersPostRequest) Method() string {
	return r.method
}

func (s *Client) NewJournalsVouchersPostRequestBody() JournalsVouchersPostRequestBody {
	return JournalsVouchersPostRequestBody{}
}

type JournalsVouchersPostRequestBody struct {
	AccountingYear struct {
		Year string `json:"year"`
	} `json:"accountingYear"`
	Journal struct {
		JournalNumber int    `json:"journalNumber"`
		Self          string `json:"self,omitempty"`
	} `json:"journal"`
	Entries struct {
		SupplierInvoices       []SupplierInvoice       `json:"supplierInvoices,omitempty"`
		SupplierPayments       []SupplierPayment       `json:"supplierPayments,omitempty"`
		CustomerPayments       []CustomerPayment       `json:"customerPayments,omitempty"`
		ManualCustomerInvoices []ManualCustomerInvoice `json:"manualCustomerInvoices,omitempty"`
	} `json:"entries"`
}

func (r *JournalsVouchersPostRequest) RequestBody() *JournalsVouchersPostRequestBody {
	return &r.requestBody
}

func (r *JournalsVouchersPostRequest) SetRequestBody(body JournalsVouchersPostRequestBody) {
	r.requestBody = body
}

func (r *JournalsVouchersPostRequest) NewResponseBody() *JournalsVouchersPostResponseBody {
	return &JournalsVouchersPostResponseBody{}
}

type JournalsVouchersPostResponseBody struct {
}

func (r *JournalsVouchersPostRequest) URL() (url.URL, error) {
	return r.client.GetEndpointURL("journals-experimental/{{.journal_number}}/vouchers", r.PathParams())
}

func (r *JournalsVouchersPostRequest) Do() (JournalsVouchersPostResponseBody, error) {
	u, err := r.URL()
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), u, r.RequestBody())
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

type SupplierInvoice struct {
	Supplier struct {
		SupplierNumber int    `json:"supplierNumber"`
		Self           string `json:"self,omitempty"`
	} `json:"supplier"`
	Amount  int `json:"amount"`
	Account struct {
		AccountNumber int    `json:"accountNumber"`
		Self          string `json:"self,omitempty"`
	} `json:"account"`
	ContraAccount struct {
		AccountNumber int    `json:"accountNumber"`
		Self          string `json:"self,omitempty"`
	} `json:"contraAccount"`
	Currency struct {
		Code string `json:"code"`
		Self string `json:"self,omitempty"`
	} `json:"currency"`
	Date string `json:"date"`
	Type string `json:"type"`
	Self string `json:"self,omitempty"`
}

type SupplierPayment struct {
	Supplier struct {
		SupplierNumber int    `json:"supplierNumber"`
		Self           string `json:"self,omitempty"`
	} `json:"supplier"`
	SupplierInvoiceNumber string `json:"supplierInvoiceNumber"`
	Text                  string `json:"text"`
	Amount                int    `json:"amount"`
	Account               struct {
		AccountNumber int    `json:"accountNumber"`
		Self          string `json:"self,omitempty"`
	} `json:"account"`
	ContraAccount struct {
		AccountNumber int    `json:"accountNumber"`
		Self          string `json:"sel,omitemptyf"`
	} `json:"contraAccount"`
	Currency struct {
		Code string `json:"code"`
		Self string `json:"self,omitempty"`
	} `json:"currency"`
	Date string `json:"date"`
	Type string `json:"type"`
	Self string `json:"self,omitempty"`
}

type CustomerPayment struct {
	Customer struct {
		CustomerNumber int    `json:"customerNumber"`
		Self           string `json:"self,omitempty"`
	} `json:"customer"`
	Text    string `json:"text"`
	Amount  int    `json:"amount"`
	Account struct {
		AccountNumber int    `json:"accountNumber"`
		Self          string `json:"self,omitempty"`
	} `json:"account"`
	ContraAccount struct {
		AccountNumber int    `json:"accountNumber"`
		Self          string `json:"self,omitempty"`
	} `json:"contraAccount"`
	Currency struct {
		Code string `json:"code"`
		Self string `json:"self,omitempty"`
	} `json:"currency"`
	Date string `json:"date"`
	Type string `json:"type"`
	Self string `json:"self,omitempty"`
}

type ManualCustomerInvoice struct {
	Customer struct {
		CustomerNumber int    `json:"customerNumber"`
		Self           string `json:"self,omitempty"`
	} `json:"customer"`
	CustomerInvoiceNumber string `json:"customerInvoiceNumber"`
	Text                  string `json:"text"`
	Amount                int    `json:"amount"`
	Account               struct {
		AccountNumber int    `json:"accountNumber"`
		Self          string `json:"self,omitempty"`
	} `json:"account"`
	ContraAccount struct {
		AccountNumber int    `json:"accountNumber"`
		Self          string `json:"self,omitempty"`
	} `json:"contraAccount"`
	Currency struct {
		Code string `json:"code"`
		Self string `json:"self,omitempty"`
	} `json:"currency"`
	Date string `json:"date"`
	Type string `json:"type"`
	Self string `json:"self,omitempty"`
}
