package economic

import (
	"net/http"
	"net/url"
)

func (c *Client) NewCustomersGetRequest() CustomersGetRequest {
	return CustomersGetRequest{
		client:      c,
		queryParams: c.NewCustomersGetQueryParams(),
		pathParams:  c.NewCustomersGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewCustomersGetRequestBody(),
	}
}

type CustomersGetRequest struct {
	client      *Client
	queryParams *CustomersGetQueryParams
	pathParams  *CustomersGetPathParams
	method      string
	headers     http.Header
	requestBody CustomersGetRequestBody
}

func (c *Client) NewCustomersGetQueryParams() *CustomersGetQueryParams {
	return &CustomersGetQueryParams{}
}

type CustomersGetQueryParams struct {
	// https://restdocs.e-conomic.com/#pagination
	SkipPages int `schema:"skippages,omitempty"`
	PageSize  int `schema:"pagesize,omitempty"`
	// https://restdocs.e-conomic.com/#filtering
	Filter Filter `schema:"filter,omitempty"`
	// https://restdocs.e-conomic.com/#sorting
	Sort string `schema:"sort,omitempty"`
}

func (r CustomersGetRequest) RequiredProperties() []string {
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

func (r CustomersGetRequest) FilterableProperties() []string {
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

func (r CustomersGetRequest) SortableProperties() []string {
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

func (p CustomersGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomersGetRequest) QueryParams() *CustomersGetQueryParams {
	return r.queryParams
}

func (c *Client) NewCustomersGetPathParams() *CustomersGetPathParams {
	return &CustomersGetPathParams{}
}

type CustomersGetPathParams struct {
}

func (p *CustomersGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomersGetRequest) PathParams() *CustomersGetPathParams {
	return r.pathParams
}

func (r *CustomersGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomersGetRequest) Method() string {
	return r.method
}

func (s *Client) NewCustomersGetRequestBody() CustomersGetRequestBody {
	return CustomersGetRequestBody{}
}

type CustomersGetRequestBody struct {
}

func (r *CustomersGetRequest) RequestBody() *CustomersGetRequestBody {
	return &r.requestBody
}

func (r *CustomersGetRequest) SetRequestBody(body CustomersGetRequestBody) {
	r.requestBody = body
}

func (r *CustomersGetRequest) NewResponseBody() *CustomersGetResponseBody {
	return &CustomersGetResponseBody{}
}

type CustomersGetResponseBody struct {
	Collection []Customer `json:"collection"`
	Pagination Pagination `json:"pagination"`
	MetaData   struct {
		Create struct {
			Description string `json:"description"`
			Href        string `json:"href"`
			HTTPMethod  string `json:"httpMethod"`
		} `json:"create"`
	} `json:"metaData"`
	Self string `json:"self"`
}

func (r *CustomersGetRequest) URL() (url.URL, error) {
	return r.client.GetEndpointURL("customers", r.PathParams())
}

func (r *CustomersGetRequest) Do() (CustomersGetResponseBody, error) {
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
