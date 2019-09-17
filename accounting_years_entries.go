package economic

import (
	"net/http"
	"net/url"
)

func (c *Client) NewAccountingYearsEntriesGetRequest() AccountingYearsEntriesGetRequest {
	return AccountingYearsEntriesGetRequest{
		client:      c,
		queryParams: c.NewAccountingYearsEntriesGetQueryParams(),
		pathParams:  c.NewAccountingYearsEntriesGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewAccountingYearsEntriesGetRequestBody(),
	}
}

type AccountingYearsEntriesGetRequest struct {
	client      *Client
	queryParams *AccountingYearsEntriesGetQueryParams
	pathParams  *AccountingYearsEntriesGetPathParams
	method      string
	headers     http.Header
	requestBody AccountingYearsEntriesGetRequestBody
}

func (c *Client) NewAccountingYearsEntriesGetQueryParams() *AccountingYearsEntriesGetQueryParams {
	return &AccountingYearsEntriesGetQueryParams{}
}

type AccountingYearsEntriesGetQueryParams struct {
}

func (r AccountingYearsEntriesGetRequest) RequiredProperties() []string {
	return []string{}
}

func (r AccountingYearsEntriesGetRequest) FilterableProperties() []string {
	return []string{}
}

func (r AccountingYearsEntriesGetRequest) SortableProperties() []string {
	return []string{}
}

func (p AccountingYearsEntriesGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AccountingYearsEntriesGetRequest) QueryParams() *AccountingYearsEntriesGetQueryParams {
	return r.queryParams
}

func (c *Client) NewAccountingYearsEntriesGetPathParams() *AccountingYearsEntriesGetPathParams {
	return &AccountingYearsEntriesGetPathParams{}
}

type AccountingYearsEntriesGetPathParams struct {
	AccountingYear string
}

func (p *AccountingYearsEntriesGetPathParams) Params() map[string]string {
	return map[string]string{
		"accounting_year": p.AccountingYear,
	}
}

func (r *AccountingYearsEntriesGetRequest) PathParams() *AccountingYearsEntriesGetPathParams {
	return r.pathParams
}

func (r *AccountingYearsEntriesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountingYearsEntriesGetRequest) Method() string {
	return r.method
}

func (s *Client) NewAccountingYearsEntriesGetRequestBody() AccountingYearsEntriesGetRequestBody {
	return AccountingYearsEntriesGetRequestBody{}
}

type AccountingYearsEntriesGetRequestBody struct {
}

func (r *AccountingYearsEntriesGetRequest) RequestBody() *AccountingYearsEntriesGetRequestBody {
	return &r.requestBody
}

func (r *AccountingYearsEntriesGetRequest) SetRequestBody(body AccountingYearsEntriesGetRequestBody) {
	r.requestBody = body
}

func (r *AccountingYearsEntriesGetRequest) NewResponseBody() *AccountingYearsEntriesGetResponseBody {
	return &AccountingYearsEntriesGetResponseBody{}
}

type AccountingYearsEntriesGetResponseBody struct {
}

func (r *AccountingYearsEntriesGetRequest) URL() (url.URL, error) {
	return r.client.GetEndpointURL("accounting-years/{{.accounting_year}}/entries", r.PathParams())
}

func (r *AccountingYearsEntriesGetRequest) Do() (AccountingYearsEntriesGetResponseBody, error) {
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
