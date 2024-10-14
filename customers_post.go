package economic

import (
	"net/http"
	"net/url"
	"strconv"
)

func (c *Client) NewCustomersPostRequest() CustomersPostRequest {
	return CustomersPostRequest{
		client:      c,
		queryParams: c.NewCustomersPostQueryParams(),
		pathParams:  c.NewCustomersPostPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewCustomersPostRequestBody(),
	}
}

type CustomersPostRequest struct {
	client      *Client
	queryParams *CustomersPostQueryParams
	pathParams  *CustomersPostPathParams
	method      string
	headers     http.Header
	requestBody CustomersPostRequestBody
}

func (c *Client) NewCustomersPostQueryParams() *CustomersPostQueryParams {
	return &CustomersPostQueryParams{}
}

type CustomersPostQueryParams struct {
}

func (r CustomersPostRequest) RequiredProperties() []string {
	return []string{
		"self",
		"settings.contraAccounts.customerPayments.self",
		"settings.contraAccounts.financeVouchers.self",
		"settings.contraAccounts.supplierPayments.self",
		"templates.self",
	}
}

func (r CustomersPostRequest) FilterableProperties() []string {
	return []string{
		"journalNumber",
		"name",
	}
}

func (r CustomersPostRequest) SortableProperties() []string {
	return []string{
		"journalNumber",
		"name",
	}
}

func (p CustomersPostQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomersPostRequest) QueryParams() *CustomersPostQueryParams {
	return r.queryParams
}

func (c *Client) NewCustomersPostPathParams() *CustomersPostPathParams {
	return &CustomersPostPathParams{
		JournalNumber: 0,
	}
}

type CustomersPostPathParams struct {
	JournalNumber int
}

func (p *CustomersPostPathParams) Params() map[string]string {
	return map[string]string{
		"journal_number": strconv.Itoa(p.JournalNumber),
	}
}

func (r *CustomersPostRequest) PathParams() *CustomersPostPathParams {
	return r.pathParams
}

func (r *CustomersPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomersPostRequest) Method() string {
	return r.method
}

func (s *Client) NewCustomersPostRequestBody() CustomersPostRequestBody {
	return CustomersPostRequestBody{}
}

type CustomersPostRequestBody CustomerPOSTSchema

// type CustomersPostRequestBody struct {
// 	CustomerPOSTSchema
// }

func (r *CustomersPostRequest) RequestBody() *CustomersPostRequestBody {
	return &r.requestBody
}

func (r *CustomersPostRequest) SetRequestBody(body CustomersPostRequestBody) {
	r.requestBody = body
}

func (r *CustomersPostRequest) NewResponseBody() *CustomersPostResponseBody {
	return &CustomersPostResponseBody{}
}

type CustomersPostResponseBody Customer

func (r *CustomersPostRequest) URL() (url.URL, error) {
	return r.client.GetEndpointURL("customers", r.PathParams())
}

func (r *CustomersPostRequest) Do() (CustomersPostResponseBody, error) {
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
