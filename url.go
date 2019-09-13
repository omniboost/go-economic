package economic

import (
	"encoding/json"
	"fmt"
	"net/url"
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
