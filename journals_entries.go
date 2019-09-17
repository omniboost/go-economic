package economic

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) NewJournalsEntriesGetRequest() JournalsEntriesGetRequest {
	return JournalsEntriesGetRequest{
		client:      c,
		queryParams: c.NewJournalsEntriesGetQueryParams(),
		pathParams:  c.NewJournalsEntriesGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewJournalsEntriesGetRequestBody(),
	}
}

type JournalsEntriesGetRequest struct {
	client      *Client
	queryParams *JournalsEntriesGetQueryParams
	pathParams  *JournalsEntriesGetPathParams
	method      string
	headers     http.Header
	requestBody JournalsEntriesGetRequestBody
}

func (c *Client) NewJournalsEntriesGetQueryParams() *JournalsEntriesGetQueryParams {
	return &JournalsEntriesGetQueryParams{}
}

type JournalsEntriesGetQueryParams struct {
}

func (r JournalsEntriesGetRequest) RequiredProperties() []string {
	return []string{}
}

func (r JournalsEntriesGetRequest) FilterableProperties() []string {
	return []string{}
}

func (r JournalsEntriesGetRequest) SortableProperties() []string {
	return []string{}
}

func (p JournalsEntriesGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *JournalsEntriesGetRequest) QueryParams() *JournalsEntriesGetQueryParams {
	return r.queryParams
}

func (c *Client) NewJournalsEntriesGetPathParams() *JournalsEntriesGetPathParams {
	return &JournalsEntriesGetPathParams{}
}

type JournalsEntriesGetPathParams struct {
	JournalNumber int
}

func (p *JournalsEntriesGetPathParams) Params() map[string]string {
	return map[string]string{
		"journal_number": fmt.Sprint(p.JournalNumber),
	}
}

func (r *JournalsEntriesGetRequest) PathParams() *JournalsEntriesGetPathParams {
	return r.pathParams
}

func (r *JournalsEntriesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalsEntriesGetRequest) Method() string {
	return r.method
}

func (s *Client) NewJournalsEntriesGetRequestBody() JournalsEntriesGetRequestBody {
	return JournalsEntriesGetRequestBody{}
}

type JournalsEntriesGetRequestBody struct {
}

func (r *JournalsEntriesGetRequest) RequestBody() *JournalsEntriesGetRequestBody {
	return &r.requestBody
}

func (r *JournalsEntriesGetRequest) SetRequestBody(body JournalsEntriesGetRequestBody) {
	r.requestBody = body
}

func (r *JournalsEntriesGetRequest) NewResponseBody() *JournalsEntriesGetResponseBody {
	return &JournalsEntriesGetResponseBody{}
}

type JournalsEntriesGetResponseBody struct {
}

func (r *JournalsEntriesGetRequest) URL() (url.URL, error) {
	return r.client.GetEndpointURL("journals-experimental/{{.journal_number}}/entries", r.PathParams())
}

func (r *JournalsEntriesGetRequest) Do() (JournalsEntriesGetResponseBody, error) {
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
