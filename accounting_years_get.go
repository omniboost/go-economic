package economic

import (
	"net/http"
	"net/url"
)

func (c *Client) NewAccountingYearsGetRequest() AccountingYearsGetRequest {
	return AccountingYearsGetRequest{
		client:      c,
		queryParams: c.NewAccountingYearsGetQueryParams(),
		pathParams:  c.NewAccountingYearsGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewAccountingYearsGetRequestBody(),
	}
}

type AccountingYearsGetRequest struct {
	client      *Client
	queryParams *AccountingYearsGetQueryParams
	pathParams  *AccountingYearsGetPathParams
	method      string
	headers     http.Header
	requestBody AccountingYearsGetRequestBody
}

func (c *Client) NewAccountingYearsGetQueryParams() *AccountingYearsGetQueryParams {
	return &AccountingYearsGetQueryParams{}
}

type AccountingYearsGetQueryParams struct {
}

func (r AccountingYearsGetRequest) RequiredProperties() []string {
	return []string{}
}

func (r AccountingYearsGetRequest) FilterableProperties() []string {
	return []string{}
}

func (r AccountingYearsGetRequest) SortableProperties() []string {
	return []string{}
}

func (p AccountingYearsGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AccountingYearsGetRequest) QueryParams() *AccountingYearsGetQueryParams {
	return r.queryParams
}

func (c *Client) NewAccountingYearsGetPathParams() *AccountingYearsGetPathParams {
	return &AccountingYearsGetPathParams{}
}

type AccountingYearsGetPathParams struct {
}

func (p *AccountingYearsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AccountingYearsGetRequest) PathParams() *AccountingYearsGetPathParams {
	return r.pathParams
}

func (r *AccountingYearsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountingYearsGetRequest) Method() string {
	return r.method
}

func (s *Client) NewAccountingYearsGetRequestBody() AccountingYearsGetRequestBody {
	return AccountingYearsGetRequestBody{}
}

type AccountingYearsGetRequestBody struct {
}

func (r *AccountingYearsGetRequest) RequestBody() *AccountingYearsGetRequestBody {
	return &r.requestBody
}

func (r *AccountingYearsGetRequest) SetRequestBody(body AccountingYearsGetRequestBody) {
	r.requestBody = body
}

func (r *AccountingYearsGetRequest) NewResponseBody() *AccountingYearsGetResponseBody {
	return &AccountingYearsGetResponseBody{}
}

type AccountingYearsGetResponseBody struct {
	Collection []AccountingYear `json:"collection"`
	Pagination struct {
		MaxPageSizeAllowed   int    `json:"maxPageSizeAllowed"`
		SkipPages            int    `json:"skipPages"`
		PageSize             int    `json:"pageSize"`
		Results              int    `json:"results"`
		ResultsWithoutFilter int    `json:"resultsWithoutFilter"`
		FirstPage            string `json:"firstPage"`
		LastPage             string `json:"lastPage"`
	} `json:"pagination"`
	Self string `json:"self"`
}

func (r *AccountingYearsGetRequest) URL() (url.URL, error) {
	return r.client.GetEndpointURL("accounting-years", r.PathParams())
}

func (r *AccountingYearsGetRequest) Do() (AccountingYearsGetResponseBody, error) {
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

type AccountingYear struct {
	FromDate Date   `json:"fromDate"`
	ToDate   Date   `json:"toDate"`
	Closed   bool   `json:"closed,omitempty"`
	Year     string `json:"year"`
	Periods  string `json:"periods"`
	Entries  string `json:"entries"`
	Totals   string `json:"totals"`
	Vouchers string `json:"vouchers"`
	Self     string `json:"self"`
}
