package shiji

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-shiji/utils"
)

func (c *Client) NewPostIndividualProfileRequest() PostIndividualProfileRequest {
	return PostIndividualProfileRequest{
		client:      c,
		queryParams: c.NewPostIndividualProfileQueryParams(),
		pathParams:  c.NewPostIndividualProfilePathParams(),
		method:      http.MethodPost,
		headers:     c.NewPostIndividualProfileHeaders(),
		requestBody: c.NewPostIndividualProfileRequestBody(),
	}
}

type PostIndividualProfileRequest struct {
	client      *Client
	queryParams *PostIndividualProfileQueryParams
	pathParams  *PostIndividualProfilePathParams
	method      string
	headers     *PostIndividualProfileHeaders
	requestBody PostIndividualProfileRequestBody
}

func (c *Client) NewPostIndividualProfileQueryParams() *PostIndividualProfileQueryParams {
	return &PostIndividualProfileQueryParams{}
}

type PostIndividualProfileQueryParams struct{}

func (p PostIndividualProfileQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(CommaSeparatedQueryParam{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PostIndividualProfileRequest) QueryParams() *PostIndividualProfileQueryParams {
	return r.queryParams
}

func (c *Client) NewPostIndividualProfileHeaders() *PostIndividualProfileHeaders {
	return &PostIndividualProfileHeaders{}
}

type PostIndividualProfileHeaders struct {
	PropertyID string `schema:"AC-Property-ID,omitempty"`
}

func (r *PostIndividualProfileRequest) Headers() *PostIndividualProfileHeaders {
	return r.headers
}

func (c *Client) NewPostIndividualProfilePathParams() *PostIndividualProfilePathParams {
	return &PostIndividualProfilePathParams{}
}

type PostIndividualProfilePathParams struct {
}

func (p *PostIndividualProfilePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostIndividualProfileRequest) PathParams() *PostIndividualProfilePathParams {
	return r.pathParams
}

func (r *PostIndividualProfileRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostIndividualProfileRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostIndividualProfileRequest) Method() string {
	return r.method
}

func (s *Client) NewPostIndividualProfileRequestBody() PostIndividualProfileRequestBody {
	return PostIndividualProfileRequestBody{}
}

type PostIndividualProfileRequestBody struct {
}

func (r *PostIndividualProfileRequest) RequestBody() *PostIndividualProfileRequestBody {
	return nil
}

func (r *PostIndividualProfileRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *PostIndividualProfileRequest) SetRequestBody(body PostIndividualProfileRequestBody) {
	r.requestBody = body
}

func (r *PostIndividualProfileRequest) NewResponseBody() *PostIndividualProfileResponseBody {
	return &PostIndividualProfileResponseBody{}
}

type PostIndividualProfileResponseBody CreateIndividualV1

func (r *PostIndividualProfileRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/profiles/v1/individual", r.PathParams())
	return &u
}

func (r *PostIndividualProfileRequest) Do() (PostIndividualProfileResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Add headers
	req.Header.Add("AC-Property-ID", r.Headers().PropertyID)

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
