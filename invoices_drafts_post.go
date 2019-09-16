package economic

import (
	"net/http"
	"net/url"

	"github.com/cydev/zero"
	"github.com/omniboost/go-economic/omitempty"
)

func (c *Client) NewInvoicesDraftsPostRequest() InvoicesDraftsPostRequest {
	return InvoicesDraftsPostRequest{
		client:      c,
		queryParams: c.NewInvoicesDraftsPostQueryParams(),
		pathParams:  c.NewInvoicesDraftsPostPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewInvoicesDraftsPostRequestBody(),
	}
}

type InvoicesDraftsPostRequest struct {
	client      *Client
	queryParams *InvoicesDraftsPostQueryParams
	pathParams  *InvoicesDraftsPostPathParams
	method      string
	headers     http.Header
	requestBody InvoicesDraftsPostRequestBody
}

func (c *Client) NewInvoicesDraftsPostQueryParams() *InvoicesDraftsPostQueryParams {
	return &InvoicesDraftsPostQueryParams{}
}

type InvoicesDraftsPostQueryParams struct {
}

func (r InvoicesDraftsPostRequest) RequiredProperties() []string {
	return []string{
		"currency",
		"customer",
		"date",
		"layout",
		"paymentTerms",
		"recipient",
		"recipient.name",
		"recipient.vatZone",
	}
}

func (r InvoicesDraftsPostRequest) FilterableProperties() []string {
	return []string{}
}

func (r InvoicesDraftsPostRequest) SortableProperties() []string {
	return []string{}
}

func (p InvoicesDraftsPostQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *InvoicesDraftsPostRequest) QueryParams() *InvoicesDraftsPostQueryParams {
	return r.queryParams
}

func (c *Client) NewInvoicesDraftsPostPathParams() *InvoicesDraftsPostPathParams {
	return &InvoicesDraftsPostPathParams{}
}

type InvoicesDraftsPostPathParams struct {
}

func (p *InvoicesDraftsPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *InvoicesDraftsPostRequest) PathParams() *InvoicesDraftsPostPathParams {
	return r.pathParams
}

func (r *InvoicesDraftsPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *InvoicesDraftsPostRequest) Method() string {
	return r.method
}

func (s *Client) NewInvoicesDraftsPostRequestBody() InvoicesDraftsPostRequestBody {
	return InvoicesDraftsPostRequestBody{}
}

type InvoicesDraftsPostRequestBody struct {
	Date                    Date                     `json:"date"`
	Currency                string                   `json:"currency"`
	ExchangeRate            float64                  `json:"exchangeRate,omitempty"`
	NetAmount               float64                  `json:"netAmount"`
	NetAmountInBaseCurrency float64                  `json:"netAmountInBaseCurrency"`
	GrossAmount             float64                  `json:"grossAmount"`
	MarginInBaseCurrency    float64                  `json:"marginInBaseCurrency"`
	MarginPercentage        float64                  `json:"marginPercentage"`
	VatAmount               float64                  `json:"vatAmount"`
	RoundingAmount          float64                  `json:"roundingAmount"`
	CostPriceInBaseCurrency float64                  `json:"costPriceInBaseCurrency"`
	PaymentTerms            InvoiceDraftPaymentTerms `json:"paymentTerms,omitempty"`
	Customer                struct {
		CustomerNumber int `json:"customerNumber"`
	} `json:"customer"`
	Recipient  InvoiceDraftRecipient `json:"recipient,omitempty"`
	Delivery   InvoiceDraftDelivery  `json:"delivery,omitempty"`
	References struct {
		Other string `json:"other"`
	} `json:"references"`
	Layout InvoiceDraftLayout `json:"layout,omitempty"`
	Lines  []InvoiceDraftLine `json:"lines,omitempty"`
}

func (r *InvoicesDraftsPostRequest) RequestBody() *InvoicesDraftsPostRequestBody {
	return &r.requestBody
}

func (r *InvoicesDraftsPostRequest) SetRequestBody(body InvoicesDraftsPostRequestBody) {
	r.requestBody = body
}

func (r *InvoicesDraftsPostRequest) NewResponseBody() *InvoicesDraftsPostResponseBody {
	return &InvoicesDraftsPostResponseBody{}
}

type InvoicesDraftsPostResponseBody struct {
	DraftInvoiceNumber int `json:"draftInvoiceNumber"`
	Soap               struct {
		CurrentInvoiceHandle struct {
			ID int `json:"id"`
		} `json:"currentInvoiceHandle"`
	} `json:"soap"`
	Templates struct {
		BookingInstructions string `json:"bookingInstructions"`
		Self                string `json:"self"`
	} `json:"templates"`
	Date                      string  `json:"date"`
	Currency                  string  `json:"currency"`
	ExchangeRate              float64 `json:"exchangeRate"`
	NetAmount                 float64 `json:"netAmount"`
	NetAmountInBaseCurrency   float64 `json:"netAmountInBaseCurrency"`
	GrossAmount               float64 `json:"grossAmount"`
	GrossAmountInBaseCurrency float64 `json:"grossAmountInBaseCurrency"`
	MarginInBaseCurrency      float64 `json:"marginInBaseCurrency"`
	MarginPercentage          float64 `json:"marginPercentage"`
	VatAmount                 float64 `json:"vatAmount"`
	RoundingAmount            float64 `json:"roundingAmount"`
	CostPriceInBaseCurrency   float64 `json:"costPriceInBaseCurrency"`
	DueDate                   string  `json:"dueDate"`
	PaymentTerms              struct {
		PaymentTermsNumber int    `json:"paymentTermsNumber"`
		Name               string `json:"name"`
		PaymentTermsType   string `json:"paymentTermsType"`
		Self               string `json:"self"`
	} `json:"paymentTerms"`
	Customer struct {
		CustomerNumber int    `json:"customerNumber"`
		Self           string `json:"self"`
	} `json:"customer"`
	Recipient struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		VatZone struct {
			Name               string `json:"name"`
			VatZoneNumber      int    `json:"vatZoneNumber"`
			EnabledForCustomer bool   `json:"enabledForCustomer"`
			EnabledForSupplier bool   `json:"enabledForSupplier"`
			Self               string `json:"self"`
		} `json:"vatZone"`
	} `json:"recipient"`
	Layout struct {
		LayoutNumber int    `json:"layoutNumber"`
		Self         string `json:"self"`
	} `json:"layout"`
	Pdf struct {
		Download string `json:"download"`
	} `json:"pdf"`
	Self string `json:"self"`
}

func (r *InvoicesDraftsPostRequest) URL() (url.URL, error) {
	return r.client.GetEndpointURL("invoices/drafts", r.PathParams())
}

func (r *InvoicesDraftsPostRequest) Do() (InvoicesDraftsPostResponseBody, error) {
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

func (i InvoicesDraftsPostRequestBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(i)
}

type InvoiceDraftLine struct {
	LineNumber int `json:"lineNumber"`
	SortKey    int `json:"sortKey"`
	Unit       struct {
		UnitNumber int    `json:"unitNumber"`
		Name       string `json:"name"`
	} `json:"unit,omitempty"`
	Product struct {
		ProductNumber string `json:"productNumber"`
	} `json:"product"`
	Quantity             float64 `json:"quantity"`
	UnitNetPrice         float64 `json:"unitNetPrice"`
	DiscountPercentage   float64 `json:"discountPercentage"`
	UnitCostPrice        float64 `json:"unitCostPrice"`
	TotalNetAmount       float64 `json:"totalNetAmount"`
	MarginInBaseCurrency float64 `json:"marginInBaseCurrency"`
	MarginPercentage     float64 `json:"marginPercentage"`
}

type InvoiceDraftDelivery struct {
	Address      string `json:"address"`
	Zip          string `json:"zip"`
	City         string `json:"city"`
	Country      string `json:"country"`
	DeliveryDate string `json:"deliveryDate"`
}

func (d InvoiceDraftDelivery) IsEmpty() bool {
	return zero.IsZero(d)
}

type InvoiceDraftRecipient struct {
	Name    string              `json:"name"`
	Address string              `json:"address"`
	Zip     string              `json:"zip"`
	City    string              `json:"city"`
	VatZone InvoiceDraftVatZone `json:"vatZone,omitempty"`
}

func (r InvoiceDraftRecipient) IsEmpty() bool {
	return zero.IsZero(r)
}

func (r InvoiceDraftRecipient) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

type InvoiceDraftPaymentTerms struct {
	PaymentTermsNumber int    `json:"paymentTermsNumber"`
	DaysOfCredit       int    `json:"daysOfCredit,omitempty"`
	Name               string `json:"name,omitempty"`
	PaymentTermsType   string `json:"paymentTermsType,omitempty"`
}

func (pt InvoiceDraftPaymentTerms) IsEmpty() bool {
	return zero.IsZero(pt)
}

type InvoiceDraftVatZone struct {
	Name               string `json:"name"`
	VatZoneNumber      int    `json:"vatZoneNumber,omitempty"`
	EnabledForCustomer bool   `json:"enabledForCustomer"`
	EnabledForSupplier bool   `json:"enabledForSupplier"`
}

func (vz InvoiceDraftVatZone) IsEmpty() bool {
	return zero.IsZero(vz)
}

type InvoiceDraftLayout struct {
	LayoutNumber int `json:"layoutNumber,omitempty"`
}

func (l InvoiceDraftLayout) IsEmpty() bool {
	return zero.IsZero(l)
}
