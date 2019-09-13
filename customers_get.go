package economic

import (
	"net/http"
	"net/url"
	"time"
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
	// Filterable properties: accountNumber, accountType, balance, barred, blockDirectEntries, debitCredit, name
	Filter string `schema:"filter"`
	// https://restdocs.e-conomic.com/#sorting
	// Sortable properties: accountNumber, accountType, balance, blockDirectEntries, debitCredit, name
	// Default sorting: accountNumber : ascending
	Sort string `schema:"sort"`
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
	Collection []struct {
		CustomerNumber int    `json:"customerNumber"`
		Currency       string `json:"currency"`
		PaymentTerms   struct {
			PaymentTermsNumber int    `json:"paymentTermsNumber"`
			Self               string `json:"self"`
		} `json:"paymentTerms"`
		CustomerGroup struct {
			CustomerGroupNumber int    `json:"customerGroupNumber"`
			Self                string `json:"self"`
		} `json:"customerGroup"`
		Address                       string  `json:"address,omitempty"`
		Balance                       float64 `json:"balance"`
		DueAmount                     float64 `json:"dueAmount"`
		CorporateIdentificationNumber string  `json:"corporateIdentificationNumber,omitempty"`
		City                          string  `json:"city,omitempty"`
		Country                       string  `json:"country,omitempty"`
		Email                         string  `json:"email,omitempty"`
		Name                          string  `json:"name"`
		Zip                           string  `json:"zip,omitempty"`
		VatZone                       struct {
			VatZoneNumber int    `json:"vatZoneNumber"`
			Self          string `json:"self"`
		} `json:"vatZone"`
		LastUpdated time.Time `json:"lastUpdated"`
		Contacts    string    `json:"contacts"`
		Templates   struct {
			Invoice     string `json:"invoice"`
			InvoiceLine string `json:"invoiceLine"`
			Self        string `json:"self"`
		} `json:"templates"`
		Totals struct {
			Drafts string `json:"drafts"`
			Booked string `json:"booked"`
			Self   string `json:"self"`
		} `json:"totals"`
		DeliveryLocations string `json:"deliveryLocations"`
		Invoices          struct {
			Drafts string `json:"drafts"`
			Booked string `json:"booked"`
			Self   string `json:"self"`
		} `json:"invoices"`
		Self                  string  `json:"self"`
		Barred                bool    `json:"barred,omitempty"`
		CreditLimit           float64 `json:"creditLimit,omitempty"`
		TelephoneAndFaxNumber string  `json:"telephoneAndFaxNumber,omitempty"`
		Website               string  `json:"website,omitempty"`
		Layout                struct {
			LayoutNumber int    `json:"layoutNumber"`
			Self         string `json:"self"`
		} `json:"layout,omitempty"`
	} `json:"collection"`
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

func (r *CustomersGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("customers", r.PathParams())
}

func (r *CustomersGetRequest) Do() (CustomersGetResponseBody, error) {
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
