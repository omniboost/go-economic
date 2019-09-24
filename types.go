package economic

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
