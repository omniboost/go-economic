package economic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
	"text/template"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-economic/" + libraryVersion
	mediaType      = "application/json"
	charset        = "utf-8"

	demoGrantToken  = "demo"
	demoSecretToken = "demo"
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "restapi.e-conomic.com",
		Path:   "",
		// RawQuery: "demo=true",
	}
)

// NewClient returns a new InvoiceXpress Client client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{
		http: httpClient,
	}

	client.SetBaseURL(BaseURL)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)
	client.SetGrantToken(demoGrantToken)
	client.SetSecretToken(demoSecretToken)

	return client
}

// Client manages communication with InvoiceXpress Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug   bool
	baseURL url.URL

	// User agent for client
	userAgent string

	mediaType             string
	charset               string
	disallowUnknownFields bool

	grantToken  string
	secretToken string

	// Optional function called after every successful request made to the DO Clients
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Client) BaseURL() url.URL {
	return c.baseURL
}

func (c *Client) SetBaseURL(baseURL url.URL) {
	c.baseURL = baseURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c *Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c *Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Client) UserAgent() string {
	return userAgent
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.disallowUnknownFields = disallowUnknownFields
}

func (c *Client) SetGrantToken(token string) {
	c.grantToken = token
}

func (c Client) GrantToken() string {
	return c.grantToken
}

func (c *Client) SetSecretToken(token string) {
	c.secretToken = token
}

func (c Client) SecretToken() string {
	return c.secretToken
}

func (c *Client) GetEndpointURL(relative string, pathParams PathParams) (url.URL, error) {
	clientURL := c.BaseURL()
	relativeURL, err := url.Parse(relative)
	if err != nil {
		return clientURL, err
	}

	clientURL.Path = path.Join(clientURL.Path, relativeURL.Path)

	query := url.Values{}
	for k, v := range clientURL.Query() {
		query[k] = append(query[k], v...)
	}
	for k, v := range relativeURL.Query() {
		query[k] = append(query[k], v...)
	}
	clientURL.RawQuery = query.Encode()

	tmpl, err := template.New("endpoint_url").Parse(clientURL.Path)
	if err != nil {
		return clientURL, err
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()

	for k, v := range params {
		params[k] = QueryEscape(v)
	}

	err = tmpl.Execute(buf, params)
	if err != nil {
		return clientURL, err
	}
	clientURL.Path = buf.String()

	return clientURL, nil
}

func (c *Client) NewRequest(ctx context.Context, method string, URL url.URL, body interface{}) (*http.Request, error) {
	// convert body struct to json
	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// create new http request
	req, err := http.NewRequest(method, URL.String(), buf)
	if err != nil {
		return nil, err
	}

	// optionally pass along context
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	// set other headers
	req.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	req.Header.Add("Accept", c.MediaType())
	req.Header.Add("User-Agent", c.UserAgent())
	req.Header.Add("X-AgreementGrantToken", c.GrantToken())
	req.Header.Add("X-AppSecretToken", c.SecretToken())

	return req, nil
}

// Do sends an Client request and returns the Client response. The Client response is json decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, responseBody interface{}) (*http.Response, error) {
	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	// try to decode body into interface parameter
	if responseBody == nil {
		return httpResp, nil
	}

	errorResponse := &ErrorResponse{Response: httpResp}
	message := &Message{}
	message2 := &Message2{}
	message3 := &Message3{}
	err = c.Unmarshal(httpResp.Body, []interface{}{}, []interface{}{responseBody, errorResponse, message, message2, message3})
	if err != nil {
		return httpResp, err
	}

	if message3.Error() != "" {
		return httpResp, message3
	}

	if message2.Error() != "" {
		return httpResp, message2
	}

	if message.Error() != "" {
		return httpResp, message
	}

	if len(errorResponse.Message.Errors) > 0 {
		return httpResp, errorResponse
	}

	return httpResp, nil
}

func (c *Client) Unmarshal(r io.Reader, vv []interface{}, optionalVv []interface{}) error {
	if len(vv) == 0 && len(optionalVv) == 0 {
		return nil
	}

	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	for _, v := range vv {
		r := bytes.NewReader(b)
		dec := json.NewDecoder(r)

		err := dec.Decode(v)
		if err != nil && err != io.EOF {
			return err
		}
	}

	for _, v := range optionalVv {
		r := bytes.NewReader(b)
		dec := json.NewDecoder(r)

		_ = dec.Decode(v)
	}

	return nil
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a json response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; (c >= 200 && c <= 299) || c == 400 {
		return nil
	}

	err := checkContentType(r)
	if err != nil {
		return errors.New(r.Status)
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	if len(data) == 0 {
		return errorResponse
	}

	// convert json to struct
	err = json.Unmarshal(data, errorResponse)
	if err != nil {
		return err
	}

	if len(errorResponse.Message.Errors) > 0 {
		return errorResponse
	}

	return nil
}

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response `json:"-"`

	Message Message
}

func (r *ErrorResponse) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Message)
}

func (r *ErrorResponse) Error() string {
	return r.Message.Error()
}

// {
//   "message": "Validation failed. 2 errors found.",
//   "errorCode": "E04300",
//   "developerHint": "Inspect validation errors and correct your request.",
//   "logId": "5ab645d7a7708766-DUS",
//   "httpStatusCode": 400,
//   "errors": [
//     {
//       "arrayIndex": 0,
//       "account": {
//         "errors": [
//           {
//             "propertyName": "account",
//             "errorMessage": "Account(s) is not found or barred.",
//             "errorCode": "E04041",
//             "inputValue": "9999",
//             "developerHint": "You must provide an accessible account."
//           }
//         ]
//       },
//       "entries": {
//         "items": [
//           {
//             "arrayIndex": 10,
//             "Account": {
//               "errors": [
//                 {
//                   "propertyName": "Account",
//                   "errorMessage": "Account '9999' not found.",
//                   "errorCode": "E07150",
//                   "inputValue": 9999,
//                   "developerHint": "Find a list of accounts at https://restapi.e-conomic.com/accounts ."
//                 }
//               ]
//             }
//           }
//         ]
//       }
//     }
//   ],
//   "logTime": "2020-06-30T09:46:19",
//   "errorCount": 2
// }

type Message struct {
	Message        string          `json:"message"`
	ErrorCode      string          `json:"errorCode"`
	DeveloperHint  string          `json:"developerHint"`
	LogID          string          `json:"logId"`
	HTTPStatusCode int             `json:"httpStatusCode"`
	Errors         ErrorCollection `json:"errors"`
	LogTime        LogTime         `json:"logTime"`
	SchemaPath     URL             `json:"schemaPath"`
}

//	{
//	    "developerHint": "Inspect validation errors and correct your request.",
//	    "errorCode": "E04300",
//	    "errorCount": 1,
//	    "errors": {
//	        "customerGroup": {
//	            "errors": [
//	                {
//	                    "developerHint": "Find a list of customer groups at https://restapi.e-conomic.com/customer-groups .",
//	                    "errorCode": "E07140",
//	                    "errorMessage": "CustomerGroup '1' not found.",
//	                    "inputValue": 1,
//	                    "propertyName": "customerGroup"
//	                }
//	            ]
//	        }
//	    },
//	    "httpStatusCode": 400,
//	    "logId": "7dacf1e1192c3689-FRA",
//	    "logTime": "2023-06-21T16:29:38",
//	    "message": "Validation failed. 1 error found."
//	}
type Message2 struct {
	Message        string `json:"message"`
	ErrorCode      string `json:"errorCode"`
	DeveloperHint  string `json:"developerHint"`
	LogID          string `json:"logId"`
	HTTPStatusCode int    `json:"httpStatusCode"`
	Errors         struct {
		CustomerGroup struct {
			Errors []Error `json:"errors"`
		} `json:"customerGroup"`
	} `json:"errors"`
	LogTime    LogTime `json:"logTime"`
	SchemaPath URL     `json:"schemaPath"`
}

type Message3 struct {
	Message        string   `json:"message"`
	ErrorCode      string   `json:"errorCode"`
	DeveloperHint  string   `json:"developerHint"`
	LogID          string   `json:"logId"`
	HTTPStatusCode int      `json:"httpStatusCode"`
	Errors         []string `json:"errors"`
	LogTime        LogTime  `json:"logTime"`
	SchemaPath     URL      `json:"schemaPath"`
}

type ErrorCollection []struct {
	ArrayIndex int `json:"arrayIndex"`
	Account    struct {
		Errors []Error `json:"errors"`
	} `json:"account"`
	Entries struct {
		Items ErrorCollection `json:"items"`
	} `json:"entries"`
	Type struct {
		Errors []Error `json:"errors"`
	} `json:"type"`
}

func (cc ErrorCollection) Error() string {
	err := []string{}

	for _, c := range cc {
		for _, e := range c.Account.Errors {
			err = append(err, e.Error())
		}
		if e := c.Entries.Items.Error(); e != "" {
			err = append(err, c.Entries.Items.Error())
		}
	}

	for _, c := range cc {
		for _, e := range c.Type.Errors {
			err = append(err, e.Error())
		}
	}

	return strings.Join(err, ", ")
}

func (m Message) Error() string {
	err := []string{}
	if m.Message != "" {
		err = append(err, m.Message)
	}

	if m.DeveloperHint != "" {
		err = append(err, m.DeveloperHint)
	}

	if m.Errors.Error() != "" {
		err = append(err, m.Errors.Error())
	}

	return strings.Join(err, ", ")
}

func (m Message2) Error() string {
	err := []string{}
	if m.Message != "" {
		err = append(err, m.Message)
	}

	if m.DeveloperHint != "" {
		err = append(err, m.DeveloperHint)
	}

	for _, e := range m.Errors.CustomerGroup.Errors {
		if e.Error() != "" {
			err = append(err, e.Error())
		}
	}

	return strings.Join(err, ", ")
}

func (m Message3) Error() string {
	err := []string{}
	if m.Message != "" {
		err = append(err, m.Message)
	}

	if m.DeveloperHint != "" {
		err = append(err, m.DeveloperHint)
	}

	for _, e := range m.Errors {
		if e != "" {
			err = append(err, e)
		}
	}

	return strings.Join(err, ", ")
}

type Error struct {
	PropertyName  string      `json:"propertyName"`
	ErrorMessage  string      `json:"errorMessage"`
	ErrorCode     string      `json:"errorCode"`
	InputValue    interface{} `json:"inputValue"`
	DeveloperHint string      `json:"developerHint"`
}

func (e *Error) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err == nil {
		e.ErrorMessage = str
		log.Println("1")
		return nil
	}

	type alias Error
	a := alias(*e)
	err = json.Unmarshal(data, &a)
	if err != nil {
		log.Println("2")
		return err
	}

	*e = Error(a)
	log.Println("3")
	return nil
}

func (r Error) Error() string {
	if r.ErrorCode == "" && r.ErrorMessage != "" {
		return r.ErrorMessage
	}

	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}

type PathParams interface {
	Params() map[string]string
}
