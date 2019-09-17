package economic

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) NewJournalsVouchersGetRequest() JournalsVouchersGetRequest {
	return JournalsVouchersGetRequest{
		client:      c,
		queryParams: c.NewJournalsVouchersGetQueryParams(),
		pathParams:  c.NewJournalsVouchersGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewJournalsVouchersGetRequestBody(),
	}
}

type JournalsVouchersGetRequest struct {
	client      *Client
	queryParams *JournalsVouchersGetQueryParams
	pathParams  *JournalsVouchersGetPathParams
	method      string
	headers     http.Header
	requestBody JournalsVouchersGetRequestBody
}

func (c *Client) NewJournalsVouchersGetQueryParams() *JournalsVouchersGetQueryParams {
	return &JournalsVouchersGetQueryParams{}
}

type JournalsVouchersGetQueryParams struct {
}

func (r JournalsVouchersGetRequest) RequiredProperties() []string {
	return []string{}
}

func (r JournalsVouchersGetRequest) FilterableProperties() []string {
	return []string{}
}

func (r JournalsVouchersGetRequest) SortableProperties() []string {
	return []string{}
}

func (p JournalsVouchersGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *JournalsVouchersGetRequest) QueryParams() *JournalsVouchersGetQueryParams {
	return r.queryParams
}

func (c *Client) NewJournalsVouchersGetPathParams() *JournalsVouchersGetPathParams {
	return &JournalsVouchersGetPathParams{}
}

type JournalsVouchersGetPathParams struct {
	JournalNumber int
}

func (p *JournalsVouchersGetPathParams) Params() map[string]string {
	return map[string]string{
		"journal_number": fmt.Sprint(p.JournalNumber),
	}
}

func (r *JournalsVouchersGetRequest) PathParams() *JournalsVouchersGetPathParams {
	return r.pathParams
}

func (r *JournalsVouchersGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalsVouchersGetRequest) Method() string {
	return r.method
}

func (s *Client) NewJournalsVouchersGetRequestBody() JournalsVouchersGetRequestBody {
	return JournalsVouchersGetRequestBody{}
}

type JournalsVouchersGetRequestBody struct {
}

func (r *JournalsVouchersGetRequest) RequestBody() *JournalsVouchersGetRequestBody {
	return &r.requestBody
}

func (r *JournalsVouchersGetRequest) SetRequestBody(body JournalsVouchersGetRequestBody) {
	r.requestBody = body
}

func (r *JournalsVouchersGetRequest) NewResponseBody() *JournalsVouchersGetResponseBody {
	return &JournalsVouchersGetResponseBody{}
}

type JournalsVouchersGetResponseBody struct {
}

func (r *JournalsVouchersGetRequest) URL() (url.URL, error) {
	return r.client.GetEndpointURL("journals-experimental/{{.journal_number}}/vouchers", r.PathParams())
}

func (r *JournalsVouchersGetRequest) Do() (JournalsVouchersGetResponseBody, error) {
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
