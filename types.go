package economic

import "github.com/cydev/zero"

type Pagination struct {
	MaxPageSizeAllowed   int    `json:"maxPageSizeAllowed"`
	SkipPages            int    `json:"skipPages"`
	PageSize             int    `json:"pageSize"`
	Results              int    `json:"results"`
	ResultsWithoutFilter int    `json:"resultsWithoutFilter"`
	FirstPage            string `json:"firstPage"`
	NextPage             string `json:"nextPage"`
	LastPage             string `json:"lastPage"`
}

// https://restdocs.e-conomic.com/#filtering
type Filter struct {
	raw string
}

func (f *Filter) Set(s string) {
	f.raw = s
}

func (f Filter) String() string {
	return f.raw
}

func (f Filter) MarshalSchema() string {
	return f.raw
}

func (f Filter) IsZero() bool {
	return f.raw == ""
}

type Account struct {
	AccountNumber int    `json:"accountNumber"`
	Self          string `json:"self,omitempty"`
}

func (a Account) IsEmpty() bool {
	return zero.IsZero(a)
}

type Voucher struct {
	// Journal voucher number must be between 1-999999999.
	VoucherNumber int `json:"voucherNumber"`
	//  A unique link reference to the voucher item.
	Self string `json:"self,omitempty"`
}

func (v Voucher) IsEmpty() bool {
	return zero.IsZero(v)
}
