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
