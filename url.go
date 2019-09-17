package economic

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type URL url.URL

func NewURL(s string) (*URL, error) {
	uu, err := url.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("bad URL: %v", err)
	}
	nu := URL(*uu)
	return &nu, nil
}

func (u URL) String() string {
	uu := url.URL(u)
	return uu.String()
}

func (u *URL) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	nu, err := NewURL(s)
	if err != nil {
		return err
	}
	*u = *nu
	return nil
}

func (u URL) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

func QueryEscape(s string) string {
	mm := map[string]string{
		"<":  "0",
		">":  "1",
		"*":  "2",
		"%":  "3",
		":":  "4",
		"&":  "5",
		"/":  "6",
		"\\": "7",
		"_":  "8",
		" ":  "9",
		"?":  "10",
		".":  "11",
		"#":  "12",
		"+":  "13",
	}

	for k, v := range mm {
		r := "_" + v + "_"
		s = strings.Replace(s, k, r, -1)
	}

	return s
}
