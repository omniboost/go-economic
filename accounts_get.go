package economic

import (
	"net/http"
	"net/url"
)

func (c *Client) NewAccountsGetRequest() AccountsGetRequest {
	return AccountsGetRequest{
		client:      c,
		queryParams: c.NewAccountsGetQueryParams(),
		pathParams:  c.NewAccountsGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewAccountsGetRequestBody(),
	}
}

type AccountsGetRequest struct {
	client      *Client
	queryParams *AccountsGetQueryParams
	pathParams  *AccountsGetPathParams
	method      string
	headers     http.Header
	requestBody AccountsGetRequestBody
}

func (c *Client) NewAccountsGetQueryParams() *AccountsGetQueryParams {
	return &AccountsGetQueryParams{}
}

type AccountsGetQueryParams struct {
	// https://restdocs.e-conomic.com/#pagination
	SkipPages int `schema:"skippages,omitempty"`
	PageSize  int `schema:"pagesize,omitempty"`
	// https://restdocs.e-conomic.com/#filtering
	// Filterable properties: accountNumber, accountType, balance, barred, blockDirectEntries, debitCredit, name
	Filter string `schema:"filter,omitempty"`
	// https://restdocs.e-conomic.com/#sorting
	Sort string `schema:"sort,omitempty"`
}

func (r AccountsGetRequest) RequiredProperties() []string {
	return []string{
		"attention.self",
		"currency",
		"customerContact.self",
		"customerGroup",
		"customerGroup.self",
		"defaultDeliveryLocation.self",
		"invoices.self",
		"layout.self",
		"name",
		"paymentTerms",
		"paymentTerms.self",
		"salesPerson.self",
		"self",
		"templates.self",
		"totals.self",
		"vatZone",
		"vatZone.self",
	}
}

func (r AccountsGetRequest) FilterableProperties() []string {
	return []string{
		"address",
		"balance",
		"barred",
		"city",
		"corporateIdentificationNumber",
		"country",
		"creditLimit",
		"currency",
		"customerNumber",
		"ean",
		"email",
		"lastUpdated",
		"mobilePhone",
		"name",
		"publicEntryNumber",
		"telephoneAndFaxNumber",
		"vatNumber",
		"website",
		"zip",
	}
}

func (r AccountsGetRequest) SortableProperties() []string {
	return []string{
		"address",
		"balance",
		"city",
		"corporateIdentificationNumber",
		"country",
		"creditLimit",
		"currency",
		"customerNumber",
		"ean",
		"email",
		"lastUpdated",
		"mobilePhone",
		"name",
		"publicEntryNumber",
		"telephoneAndFaxNumber",
		"vatNumber",
		"website",
		"zip",
	}
}

func (p AccountsGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AccountsGetRequest) QueryParams() *AccountsGetQueryParams {
	return r.queryParams
}

func (c *Client) NewAccountsGetPathParams() *AccountsGetPathParams {
	return &AccountsGetPathParams{}
}

type AccountsGetPathParams struct {
}

func (p *AccountsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AccountsGetRequest) PathParams() *AccountsGetPathParams {
	return r.pathParams
}

func (r *AccountsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountsGetRequest) Method() string {
	return r.method
}

func (s *Client) NewAccountsGetRequestBody() AccountsGetRequestBody {
	return AccountsGetRequestBody{}
}

type AccountsGetRequestBody struct {
}

func (r *AccountsGetRequest) RequestBody() *AccountsGetRequestBody {
	return &r.requestBody
}

func (r *AccountsGetRequest) SetRequestBody(body AccountsGetRequestBody) {
	r.requestBody = body
}

func (r *AccountsGetRequest) NewResponseBody() *AccountsGetResponseBody {
	return &AccountsGetResponseBody{}
}

type AccountsGetResponseBody struct {
	Collection []struct {
		AccountNumber      int     `json:"accountNumber"`
		AccountType        string  `json:"accountType"`
		Balance            float64 `json:"balance"`
		BlockDirectEntries bool    `json:"blockDirectEntries"`
		DebitCredit        string  `json:"debitCredit"`
		Name               string  `json:"name"`
		AccountingYears    string  `json:"accountingYears"`
		Self               string  `json:"self"`
		VatAccount         struct {
			VatCode string `json:"vatCode"`
			Self    string `json:"self"`
		} `json:"vatAccount,omitempty"`
		TotalFromAccount struct {
			AccountNumber int    `json:"accountNumber"`
			Self          string `json:"self"`
		} `json:"totalFromAccount,omitempty"`
	} `json:"collection"`
	Pagination Pagination `json:"pagination"`
	Self       string     `json:"self"`
}

func (r *AccountsGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("accounts", r.PathParams())
}

func (r *AccountsGetRequest) Do() (AccountsGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), nil)
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
