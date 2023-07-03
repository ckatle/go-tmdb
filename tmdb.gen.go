// Package tmdb provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package tmdb

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Country defines model for Country.
type Country struct {
	EnglishName string `json:"english_name"`

	// Iso31661 ISO 3166-1 tag
	Iso31661   string  `json:"iso_3166_1"`
	NativeName *string `json:"native_name,omitempty"`
}

// Department defines model for Department.
type Department struct {
	// Department The name of the department
	Department string   `json:"department"`
	Jobs       []string `json:"jobs"`
}

// Error defines model for Error.
type Error struct {
	StatusCode    int32  `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
}

// Language defines model for Language.
type Language = string

// Unauthorized defines model for Unauthorized.
type Unauthorized = Error

// ConfigurationCountriesParams defines parameters for ConfigurationCountries.
type ConfigurationCountriesParams struct {
	// Language Pass a ISO 639-1 value to display translated data for the fields that support it.
	Language *Language `form:"language,omitempty" json:"language,omitempty"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// ConfigurationDetails request
	ConfigurationDetails(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ConfigurationCountries request
	ConfigurationCountries(ctx context.Context, params *ConfigurationCountriesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ConfigurationJobs request
	ConfigurationJobs(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) ConfigurationDetails(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewConfigurationDetailsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ConfigurationCountries(ctx context.Context, params *ConfigurationCountriesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewConfigurationCountriesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ConfigurationJobs(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewConfigurationJobsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewConfigurationDetailsRequest generates requests for ConfigurationDetails
func NewConfigurationDetailsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/configuration")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewConfigurationCountriesRequest generates requests for ConfigurationCountries
func NewConfigurationCountriesRequest(server string, params *ConfigurationCountriesParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/configuration/countries")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Language != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "language", runtime.ParamLocationQuery, *params.Language); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewConfigurationJobsRequest generates requests for ConfigurationJobs
func NewConfigurationJobsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/configuration/jobs")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// ConfigurationDetails request
	ConfigurationDetailsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ConfigurationDetailsResponse, error)

	// ConfigurationCountries request
	ConfigurationCountriesWithResponse(ctx context.Context, params *ConfigurationCountriesParams, reqEditors ...RequestEditorFn) (*ConfigurationCountriesResponse, error)

	// ConfigurationJobs request
	ConfigurationJobsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ConfigurationJobsResponse, error)
}

type ConfigurationDetailsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		ChangeKeys *[]string `json:"change_keys,omitempty"`
		Images     *struct {
			BackdropSizes *[]string `json:"backdrop_sizes,omitempty"`
			BaseUrl       *string   `json:"base_url,omitempty"`
			LogoSizes     *[]string `json:"logo_sizes,omitempty"`
			PosterSizes   *[]string `json:"poster_sizes,omitempty"`
			ProfileSizes  *[]string `json:"profile_sizes,omitempty"`
			SecureBaseUrl *string   `json:"secure_base_url,omitempty"`
			StillSizes    *[]string `json:"still_sizes,omitempty"`
		} `json:"images,omitempty"`
	}
	JSON401 *Error
}

// Status returns HTTPResponse.Status
func (r ConfigurationDetailsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ConfigurationDetailsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ConfigurationCountriesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Country
	JSON401      *Error
}

// Status returns HTTPResponse.Status
func (r ConfigurationCountriesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ConfigurationCountriesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ConfigurationJobsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Department
	JSON401      *Error
}

// Status returns HTTPResponse.Status
func (r ConfigurationJobsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ConfigurationJobsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ConfigurationDetailsWithResponse request returning *ConfigurationDetailsResponse
func (c *ClientWithResponses) ConfigurationDetailsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ConfigurationDetailsResponse, error) {
	rsp, err := c.ConfigurationDetails(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseConfigurationDetailsResponse(rsp)
}

// ConfigurationCountriesWithResponse request returning *ConfigurationCountriesResponse
func (c *ClientWithResponses) ConfigurationCountriesWithResponse(ctx context.Context, params *ConfigurationCountriesParams, reqEditors ...RequestEditorFn) (*ConfigurationCountriesResponse, error) {
	rsp, err := c.ConfigurationCountries(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseConfigurationCountriesResponse(rsp)
}

// ConfigurationJobsWithResponse request returning *ConfigurationJobsResponse
func (c *ClientWithResponses) ConfigurationJobsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ConfigurationJobsResponse, error) {
	rsp, err := c.ConfigurationJobs(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseConfigurationJobsResponse(rsp)
}

// ParseConfigurationDetailsResponse parses an HTTP response from a ConfigurationDetailsWithResponse call
func ParseConfigurationDetailsResponse(rsp *http.Response) (*ConfigurationDetailsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ConfigurationDetailsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			ChangeKeys *[]string `json:"change_keys,omitempty"`
			Images     *struct {
				BackdropSizes *[]string `json:"backdrop_sizes,omitempty"`
				BaseUrl       *string   `json:"base_url,omitempty"`
				LogoSizes     *[]string `json:"logo_sizes,omitempty"`
				PosterSizes   *[]string `json:"poster_sizes,omitempty"`
				ProfileSizes  *[]string `json:"profile_sizes,omitempty"`
				SecureBaseUrl *string   `json:"secure_base_url,omitempty"`
				StillSizes    *[]string `json:"still_sizes,omitempty"`
			} `json:"images,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseConfigurationCountriesResponse parses an HTTP response from a ConfigurationCountriesWithResponse call
func ParseConfigurationCountriesResponse(rsp *http.Response) (*ConfigurationCountriesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ConfigurationCountriesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Country
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseConfigurationJobsResponse parses an HTTP response from a ConfigurationJobsWithResponse call
func ParseConfigurationJobsResponse(rsp *http.Response) (*ConfigurationJobsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ConfigurationJobsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Department
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Details
	// (GET /configuration)
	ConfigurationDetails(ctx echo.Context) error
	// Countries
	// (GET /configuration/countries)
	ConfigurationCountries(ctx echo.Context, params ConfigurationCountriesParams) error
	// Jobs
	// (GET /configuration/jobs)
	ConfigurationJobs(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ConfigurationDetails converts echo context to params.
func (w *ServerInterfaceWrapper) ConfigurationDetails(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ConfigurationDetails(ctx)
	return err
}

// ConfigurationCountries converts echo context to params.
func (w *ServerInterfaceWrapper) ConfigurationCountries(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params ConfigurationCountriesParams
	// ------------- Optional query parameter "language" -------------

	err = runtime.BindQueryParameter("form", true, false, "language", ctx.QueryParams(), &params.Language)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter language: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ConfigurationCountries(ctx, params)
	return err
}

// ConfigurationJobs converts echo context to params.
func (w *ServerInterfaceWrapper) ConfigurationJobs(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ConfigurationJobs(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/configuration", wrapper.ConfigurationDetails)
	router.GET(baseURL+"/configuration/countries", wrapper.ConfigurationCountries)
	router.GET(baseURL+"/configuration/jobs", wrapper.ConfigurationJobs)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7xXW2/bRhP9K4v58uAAtGhZuSB882cXgdsESeGkQCuowogckRuTu8zu0Iki6L8Xs9SF",
	"tGRHRdq+LVe7czlzzuxoCamtamvIsIdkCTU6rIjJha83aPIGc5J1Rj51umZtDSTwHr1XqK5v3qkXo1en",
	"Q3WHZUOKrcq0r0tcKHZofIlMmcqQUc2tU1yQmmsqM6+4QFa+qWvrWGkeQARaDH9uyC0gAoMVQQLlJoII",
	"fFpQhW0oc2xKhgTInH68gQgq/PqGTM4FJM8jqLTZfJ1HUCMzObH958kYT79Nluerp6cn44vTP8LyCUTA",
	"i1q8eXba5LBarSJw5GtrPAUgPhpsuLBOf6NMvlNrmAzLEuu61CkKLvEnL+Asgb5iVZcEyXgJnpEbP01t",
	"RpC8jDbfFXkfkIVrc4elztTF+2t1S4tE/W4bVTWe1YxU7tAIhKjaQ7e0EKh8k6bkPSRzLD2tJl10njia",
	"QwL/i3eVjdtfffyTc9a1+fXrebFOL+ShtJlbV63XXlXae21yZZ3SbbADEBNrq+L00jaG3SIAkmVabmL5",
	"3tmaHGvaBBpB3dlaApm81L6YtsXu4AYXJrPO4X5pItDeTkfDFy+mw31aCh/lt9OhYswh6lq86vPk/BGe",
	"bLjx5JB/g6zv6O/EHNj0udFOyDPuJhD1IZhsr9rZJ0pZ3F1RjY6rDduOBzfrXezj9KEgJR6VnQdRds4e",
	"yPiTnQWLmqkKi70T6w10Dhd7+faMB1uH8mypKS2ol0VPPktoeQkJaMOj87aiumoqSJ69DAVtP4ZbB9ow",
	"5eTEw33lHchjq6vtbzNrS0Kzl9TmZNSLcM/JfqbihdLGaV7ciH7aNGeEjpyoMCQdfhDvYXtXk4K5buUr",
	"Ej1c17f2TpO6QsYZelInH95e/f+pdBcxozlQVfbWW3fkfHv5bDAaDAUFW5PBWkMCo8HZYARBG0WIM06t",
	"meu8cdh6XEJOB/j1q3TxwCzpar07KiNGXXppY1LosHmdQQKX3WNX7Sm414jPz86O6L+7ZtgnU1qgyWl6",
	"S4s+n3caxkweluh7DI9AV5i3RvsuZpjeZs7WU6+/0UNevozOzo5xIgWcNq7s3xYSJHEcIhhwlc0G1uUx",
	"x3V8yGZpc/t4MM+eHxNLbT2Te9zSq/OjLDk71yX9E0EFJdH0YZz88UB51mX54/mtDin+/s7e+/vuF7H0",
	"7Gz40Au+VUHcm0VCO2mqCuXthZ1oGHMvbaqnKZjI8b6E4zS83Gv2HhTza+Ig5VJ7lgdje0Od9N9b/1Q1",
	"njLFhbNNXtiGlTSa7yj9chtA1Bs/x4eB2B2Jt+OpDEA/1Ca2xX5seNrMOPsl/5fK2UXm+IJuXuujailr",
	"uaDQZJ0pwKsvJLVU1hxTwp/F5Q+WoDszd0cXuEjZOr8ZHZJxuwERXGJFFiK4qSnVWKrXDXnp3b9ZnRJM",
	"VtE9Q1faUcoi2p2tdi+Ye8cFOZgIl46iQ2cw+88YsQb6QTJ0hosAZHesGIfUPLm7jbxCx9y2Saz1gAuq",
	"ZHxYt8pRUNba23Lzn6zvdTVZ/RUAAP//Hk+J80UOAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
