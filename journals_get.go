package economic

import (
	"net/http"
	"net/url"
)

func (c *Client) NewJournalsGetRequest() JournalsGetRequest {
	return JournalsGetRequest{
		client:      c,
		queryParams: c.NewJournalsGetQueryParams(),
		pathParams:  c.NewJournalsGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewJournalsGetRequestBody(),
	}
}

type JournalsGetRequest struct {
	client      *Client
	queryParams *JournalsGetQueryParams
	pathParams  *JournalsGetPathParams
	method      string
	headers     http.Header
	requestBody JournalsGetRequestBody
}

func (c *Client) NewJournalsGetQueryParams() *JournalsGetQueryParams {
	return &JournalsGetQueryParams{}
}

type JournalsGetQueryParams struct {
	// https://restdocs.e-conomic.com/#pagination
	SkipPages int `schema:"skippages,omitempty"`
	PageSize  int `schema:"pagesize,omitempty"`
	// https://restdocs.e-conomic.com/#filtering
	// Filterable properties: accountNumber, accountType, balance, barred, blockDirectEntries, debitCredit, name
	Filter string `schema:"filter,omitempty"`
	// https://restdocs.e-conomic.com/#sorting
	Sort string `schema:"sort,omitempty"`
}

func (r JournalsGetRequest) RequiredProperties() []string {
	return []string{
		"self",
		"settings.contraAccounts.customerPayments.self",
		"settings.contraAccounts.financeVouchers.self",
		"settings.contraAccounts.supplierPayments.self",
		"templates.self",
	}
}

func (r JournalsGetRequest) FilterableProperties() []string {
	return []string{
		"journalNumber",
		"name",
	}
}

func (r JournalsGetRequest) SortableProperties() []string {
	return []string{
		"journalNumber",
		"name",
	}
}

func (p JournalsGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *JournalsGetRequest) QueryParams() *JournalsGetQueryParams {
	return r.queryParams
}

func (c *Client) NewJournalsGetPathParams() *JournalsGetPathParams {
	return &JournalsGetPathParams{}
}

type JournalsGetPathParams struct {
}

func (p *JournalsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *JournalsGetRequest) PathParams() *JournalsGetPathParams {
	return r.pathParams
}

func (r *JournalsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalsGetRequest) Method() string {
	return r.method
}

func (s *Client) NewJournalsGetRequestBody() JournalsGetRequestBody {
	return JournalsGetRequestBody{}
}

type JournalsGetRequestBody struct {
}

func (r *JournalsGetRequest) RequestBody() *JournalsGetRequestBody {
	return &r.requestBody
}

func (r *JournalsGetRequest) SetRequestBody(body JournalsGetRequestBody) {
	r.requestBody = body
}

func (r *JournalsGetRequest) NewResponseBody() *JournalsGetResponseBody {
	return &JournalsGetResponseBody{}
}

type JournalsGetResponseBody struct {
	Collection []Journal  `json:"collection"`
	Pagination Pagination `json:"pagination"`
	Self       string     `json:"self"`
}

func (r *JournalsGetRequest) URL() (url.URL, error) {
	return r.client.GetEndpointURL("journals-experimental", r.PathParams())
}

func (r *JournalsGetRequest) Do() (JournalsGetResponseBody, error) {
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
