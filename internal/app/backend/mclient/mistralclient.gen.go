// Package mclient provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package mclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// AddQuestionsRequest defines model for AddQuestionsRequest.
type AddQuestionsRequest struct {
	Filename     *string          `json:"filename,omitempty"`
	NewQuestions []QuestionAnswer `json:"new_questions"`
}

// GetChatResponseRequest defines model for GetChatResponseRequest.
type GetChatResponseRequest struct {
	Context *string `json:"context,omitempty"`
	Model   *string `json:"model,omitempty"`
	Prompt  string  `json:"prompt"`
}

// HTTPValidationError defines model for HTTPValidationError.
type HTTPValidationError struct {
	Detail *[]ValidationError `json:"detail,omitempty"`
}

// ProcessQuestionsRequest defines model for ProcessQuestionsRequest.
type ProcessQuestionsRequest struct {
	Filename      *string  `json:"filename,omitempty"`
	QuestionsList []string `json:"questions_list"`
	UseSaved      *bool    `json:"use_saved,omitempty"`
}

// QuestionAnswer defines model for QuestionAnswer.
type QuestionAnswer struct {
	Answer   string `json:"answer"`
	Question string `json:"question"`
}

// SaveVectorizedDataRequest defines model for SaveVectorizedDataRequest.
type SaveVectorizedDataRequest struct {
	Data       []QuestionAnswer `json:"data"`
	Embeddings [][]float32      `json:"embeddings"`
	Filename   *string          `json:"filename,omitempty"`
}

// ValidationError defines model for ValidationError.
type ValidationError struct {
	Loc  []ValidationError_Loc_Item `json:"loc"`
	Msg  string                     `json:"msg"`
	Type string                     `json:"type"`
}

// ValidationErrorLoc0 defines model for .
type ValidationErrorLoc0 = string

// ValidationErrorLoc1 defines model for .
type ValidationErrorLoc1 = int

// ValidationError_Loc_Item defines model for ValidationError.loc.Item.
type ValidationError_Loc_Item struct {
	union json.RawMessage
}

// ApiAddQuestionsAddQuestionsPostJSONRequestBody defines body for ApiAddQuestionsAddQuestionsPost for application/json ContentType.
type ApiAddQuestionsAddQuestionsPostJSONRequestBody = AddQuestionsRequest

// ApiGetChatResponseGetChatResponsePostJSONRequestBody defines body for ApiGetChatResponseGetChatResponsePost for application/json ContentType.
type ApiGetChatResponseGetChatResponsePostJSONRequestBody = GetChatResponseRequest

// ApiProcessQuestionsProcessQuestionsPostJSONRequestBody defines body for ApiProcessQuestionsProcessQuestionsPost for application/json ContentType.
type ApiProcessQuestionsProcessQuestionsPostJSONRequestBody = ProcessQuestionsRequest

// ApiSaveVectorizedDataSaveVectorizedDataPostJSONRequestBody defines body for ApiSaveVectorizedDataSaveVectorizedDataPost for application/json ContentType.
type ApiSaveVectorizedDataSaveVectorizedDataPostJSONRequestBody = SaveVectorizedDataRequest

// AsValidationErrorLoc0 returns the union data inside the ValidationError_Loc_Item as a ValidationErrorLoc0
func (t ValidationError_Loc_Item) AsValidationErrorLoc0() (ValidationErrorLoc0, error) {
	var body ValidationErrorLoc0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromValidationErrorLoc0 overwrites any union data inside the ValidationError_Loc_Item as the provided ValidationErrorLoc0
func (t *ValidationError_Loc_Item) FromValidationErrorLoc0(v ValidationErrorLoc0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeValidationErrorLoc0 performs a merge with any union data inside the ValidationError_Loc_Item, using the provided ValidationErrorLoc0
func (t *ValidationError_Loc_Item) MergeValidationErrorLoc0(v ValidationErrorLoc0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsValidationErrorLoc1 returns the union data inside the ValidationError_Loc_Item as a ValidationErrorLoc1
func (t ValidationError_Loc_Item) AsValidationErrorLoc1() (ValidationErrorLoc1, error) {
	var body ValidationErrorLoc1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromValidationErrorLoc1 overwrites any union data inside the ValidationError_Loc_Item as the provided ValidationErrorLoc1
func (t *ValidationError_Loc_Item) FromValidationErrorLoc1(v ValidationErrorLoc1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeValidationErrorLoc1 performs a merge with any union data inside the ValidationError_Loc_Item, using the provided ValidationErrorLoc1
func (t *ValidationError_Loc_Item) MergeValidationErrorLoc1(v ValidationErrorLoc1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t ValidationError_Loc_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ValidationError_Loc_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
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
	// ApiAddQuestionsAddQuestionsPostWithBody request with any body
	ApiAddQuestionsAddQuestionsPostWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	ApiAddQuestionsAddQuestionsPost(ctx context.Context, body ApiAddQuestionsAddQuestionsPostJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiGetChatResponseGetChatResponsePostWithBody request with any body
	ApiGetChatResponseGetChatResponsePostWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	ApiGetChatResponseGetChatResponsePost(ctx context.Context, body ApiGetChatResponseGetChatResponsePostJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiProcessQuestionsProcessQuestionsPostWithBody request with any body
	ApiProcessQuestionsProcessQuestionsPostWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	ApiProcessQuestionsProcessQuestionsPost(ctx context.Context, body ApiProcessQuestionsProcessQuestionsPostJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ApiSaveVectorizedDataSaveVectorizedDataPostWithBody request with any body
	ApiSaveVectorizedDataSaveVectorizedDataPostWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	ApiSaveVectorizedDataSaveVectorizedDataPost(ctx context.Context, body ApiSaveVectorizedDataSaveVectorizedDataPostJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) ApiAddQuestionsAddQuestionsPostWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiAddQuestionsAddQuestionsPostRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ApiAddQuestionsAddQuestionsPost(ctx context.Context, body ApiAddQuestionsAddQuestionsPostJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiAddQuestionsAddQuestionsPostRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ApiGetChatResponseGetChatResponsePostWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiGetChatResponseGetChatResponsePostRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ApiGetChatResponseGetChatResponsePost(ctx context.Context, body ApiGetChatResponseGetChatResponsePostJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiGetChatResponseGetChatResponsePostRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ApiProcessQuestionsProcessQuestionsPostWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiProcessQuestionsProcessQuestionsPostRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ApiProcessQuestionsProcessQuestionsPost(ctx context.Context, body ApiProcessQuestionsProcessQuestionsPostJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiProcessQuestionsProcessQuestionsPostRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ApiSaveVectorizedDataSaveVectorizedDataPostWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiSaveVectorizedDataSaveVectorizedDataPostRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ApiSaveVectorizedDataSaveVectorizedDataPost(ctx context.Context, body ApiSaveVectorizedDataSaveVectorizedDataPostJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewApiSaveVectorizedDataSaveVectorizedDataPostRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewApiAddQuestionsAddQuestionsPostRequest calls the generic ApiAddQuestionsAddQuestionsPost builder with application/json body
func NewApiAddQuestionsAddQuestionsPostRequest(server string, body ApiAddQuestionsAddQuestionsPostJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewApiAddQuestionsAddQuestionsPostRequestWithBody(server, "application/json", bodyReader)
}

// NewApiAddQuestionsAddQuestionsPostRequestWithBody generates requests for ApiAddQuestionsAddQuestionsPost with any type of body
func NewApiAddQuestionsAddQuestionsPostRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/add_questions")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewApiGetChatResponseGetChatResponsePostRequest calls the generic ApiGetChatResponseGetChatResponsePost builder with application/json body
func NewApiGetChatResponseGetChatResponsePostRequest(server string, body ApiGetChatResponseGetChatResponsePostJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewApiGetChatResponseGetChatResponsePostRequestWithBody(server, "application/json", bodyReader)
}

// NewApiGetChatResponseGetChatResponsePostRequestWithBody generates requests for ApiGetChatResponseGetChatResponsePost with any type of body
func NewApiGetChatResponseGetChatResponsePostRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/get_chat_response")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewApiProcessQuestionsProcessQuestionsPostRequest calls the generic ApiProcessQuestionsProcessQuestionsPost builder with application/json body
func NewApiProcessQuestionsProcessQuestionsPostRequest(server string, body ApiProcessQuestionsProcessQuestionsPostJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewApiProcessQuestionsProcessQuestionsPostRequestWithBody(server, "application/json", bodyReader)
}

// NewApiProcessQuestionsProcessQuestionsPostRequestWithBody generates requests for ApiProcessQuestionsProcessQuestionsPost with any type of body
func NewApiProcessQuestionsProcessQuestionsPostRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/process_questions")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewApiSaveVectorizedDataSaveVectorizedDataPostRequest calls the generic ApiSaveVectorizedDataSaveVectorizedDataPost builder with application/json body
func NewApiSaveVectorizedDataSaveVectorizedDataPostRequest(server string, body ApiSaveVectorizedDataSaveVectorizedDataPostJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewApiSaveVectorizedDataSaveVectorizedDataPostRequestWithBody(server, "application/json", bodyReader)
}

// NewApiSaveVectorizedDataSaveVectorizedDataPostRequestWithBody generates requests for ApiSaveVectorizedDataSaveVectorizedDataPost with any type of body
func NewApiSaveVectorizedDataSaveVectorizedDataPostRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/save_vectorized_data")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

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
	// ApiAddQuestionsAddQuestionsPostWithBodyWithResponse request with any body
	ApiAddQuestionsAddQuestionsPostWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiAddQuestionsAddQuestionsPostResponse, error)

	ApiAddQuestionsAddQuestionsPostWithResponse(ctx context.Context, body ApiAddQuestionsAddQuestionsPostJSONRequestBody, reqEditors ...RequestEditorFn) (*ApiAddQuestionsAddQuestionsPostResponse, error)

	// ApiGetChatResponseGetChatResponsePostWithBodyWithResponse request with any body
	ApiGetChatResponseGetChatResponsePostWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiGetChatResponseGetChatResponsePostResponse, error)

	ApiGetChatResponseGetChatResponsePostWithResponse(ctx context.Context, body ApiGetChatResponseGetChatResponsePostJSONRequestBody, reqEditors ...RequestEditorFn) (*ApiGetChatResponseGetChatResponsePostResponse, error)

	// ApiProcessQuestionsProcessQuestionsPostWithBodyWithResponse request with any body
	ApiProcessQuestionsProcessQuestionsPostWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiProcessQuestionsProcessQuestionsPostResponse, error)

	ApiProcessQuestionsProcessQuestionsPostWithResponse(ctx context.Context, body ApiProcessQuestionsProcessQuestionsPostJSONRequestBody, reqEditors ...RequestEditorFn) (*ApiProcessQuestionsProcessQuestionsPostResponse, error)

	// ApiSaveVectorizedDataSaveVectorizedDataPostWithBodyWithResponse request with any body
	ApiSaveVectorizedDataSaveVectorizedDataPostWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiSaveVectorizedDataSaveVectorizedDataPostResponse, error)

	ApiSaveVectorizedDataSaveVectorizedDataPostWithResponse(ctx context.Context, body ApiSaveVectorizedDataSaveVectorizedDataPostJSONRequestBody, reqEditors ...RequestEditorFn) (*ApiSaveVectorizedDataSaveVectorizedDataPostResponse, error)
}

type ApiAddQuestionsAddQuestionsPostResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *interface{}
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r ApiAddQuestionsAddQuestionsPostResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ApiAddQuestionsAddQuestionsPostResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ApiGetChatResponseGetChatResponsePostResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *interface{}
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r ApiGetChatResponseGetChatResponsePostResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ApiGetChatResponseGetChatResponsePostResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ApiProcessQuestionsProcessQuestionsPostResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *interface{}
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r ApiProcessQuestionsProcessQuestionsPostResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ApiProcessQuestionsProcessQuestionsPostResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ApiSaveVectorizedDataSaveVectorizedDataPostResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *interface{}
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r ApiSaveVectorizedDataSaveVectorizedDataPostResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ApiSaveVectorizedDataSaveVectorizedDataPostResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ApiAddQuestionsAddQuestionsPostWithBodyWithResponse request with arbitrary body returning *ApiAddQuestionsAddQuestionsPostResponse
func (c *ClientWithResponses) ApiAddQuestionsAddQuestionsPostWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiAddQuestionsAddQuestionsPostResponse, error) {
	rsp, err := c.ApiAddQuestionsAddQuestionsPostWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseApiAddQuestionsAddQuestionsPostResponse(rsp)
}

func (c *ClientWithResponses) ApiAddQuestionsAddQuestionsPostWithResponse(ctx context.Context, body ApiAddQuestionsAddQuestionsPostJSONRequestBody, reqEditors ...RequestEditorFn) (*ApiAddQuestionsAddQuestionsPostResponse, error) {
	rsp, err := c.ApiAddQuestionsAddQuestionsPost(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseApiAddQuestionsAddQuestionsPostResponse(rsp)
}

// ApiGetChatResponseGetChatResponsePostWithBodyWithResponse request with arbitrary body returning *ApiGetChatResponseGetChatResponsePostResponse
func (c *ClientWithResponses) ApiGetChatResponseGetChatResponsePostWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiGetChatResponseGetChatResponsePostResponse, error) {
	rsp, err := c.ApiGetChatResponseGetChatResponsePostWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseApiGetChatResponseGetChatResponsePostResponse(rsp)
}

func (c *ClientWithResponses) ApiGetChatResponseGetChatResponsePostWithResponse(ctx context.Context, body ApiGetChatResponseGetChatResponsePostJSONRequestBody, reqEditors ...RequestEditorFn) (*ApiGetChatResponseGetChatResponsePostResponse, error) {
	rsp, err := c.ApiGetChatResponseGetChatResponsePost(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseApiGetChatResponseGetChatResponsePostResponse(rsp)
}

// ApiProcessQuestionsProcessQuestionsPostWithBodyWithResponse request with arbitrary body returning *ApiProcessQuestionsProcessQuestionsPostResponse
func (c *ClientWithResponses) ApiProcessQuestionsProcessQuestionsPostWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiProcessQuestionsProcessQuestionsPostResponse, error) {
	rsp, err := c.ApiProcessQuestionsProcessQuestionsPostWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseApiProcessQuestionsProcessQuestionsPostResponse(rsp)
}

func (c *ClientWithResponses) ApiProcessQuestionsProcessQuestionsPostWithResponse(ctx context.Context, body ApiProcessQuestionsProcessQuestionsPostJSONRequestBody, reqEditors ...RequestEditorFn) (*ApiProcessQuestionsProcessQuestionsPostResponse, error) {
	rsp, err := c.ApiProcessQuestionsProcessQuestionsPost(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseApiProcessQuestionsProcessQuestionsPostResponse(rsp)
}

// ApiSaveVectorizedDataSaveVectorizedDataPostWithBodyWithResponse request with arbitrary body returning *ApiSaveVectorizedDataSaveVectorizedDataPostResponse
func (c *ClientWithResponses) ApiSaveVectorizedDataSaveVectorizedDataPostWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ApiSaveVectorizedDataSaveVectorizedDataPostResponse, error) {
	rsp, err := c.ApiSaveVectorizedDataSaveVectorizedDataPostWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseApiSaveVectorizedDataSaveVectorizedDataPostResponse(rsp)
}

func (c *ClientWithResponses) ApiSaveVectorizedDataSaveVectorizedDataPostWithResponse(ctx context.Context, body ApiSaveVectorizedDataSaveVectorizedDataPostJSONRequestBody, reqEditors ...RequestEditorFn) (*ApiSaveVectorizedDataSaveVectorizedDataPostResponse, error) {
	rsp, err := c.ApiSaveVectorizedDataSaveVectorizedDataPost(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseApiSaveVectorizedDataSaveVectorizedDataPostResponse(rsp)
}

// ParseApiAddQuestionsAddQuestionsPostResponse parses an HTTP response from a ApiAddQuestionsAddQuestionsPostWithResponse call
func ParseApiAddQuestionsAddQuestionsPostResponse(rsp *http.Response) (*ApiAddQuestionsAddQuestionsPostResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ApiAddQuestionsAddQuestionsPostResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseApiGetChatResponseGetChatResponsePostResponse parses an HTTP response from a ApiGetChatResponseGetChatResponsePostWithResponse call
func ParseApiGetChatResponseGetChatResponsePostResponse(rsp *http.Response) (*ApiGetChatResponseGetChatResponsePostResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ApiGetChatResponseGetChatResponsePostResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseApiProcessQuestionsProcessQuestionsPostResponse parses an HTTP response from a ApiProcessQuestionsProcessQuestionsPostWithResponse call
func ParseApiProcessQuestionsProcessQuestionsPostResponse(rsp *http.Response) (*ApiProcessQuestionsProcessQuestionsPostResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ApiProcessQuestionsProcessQuestionsPostResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseApiSaveVectorizedDataSaveVectorizedDataPostResponse parses an HTTP response from a ApiSaveVectorizedDataSaveVectorizedDataPostWithResponse call
func ParseApiSaveVectorizedDataSaveVectorizedDataPostResponse(rsp *http.Response) (*ApiSaveVectorizedDataSaveVectorizedDataPostResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ApiSaveVectorizedDataSaveVectorizedDataPostResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}
