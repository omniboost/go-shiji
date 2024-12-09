package shiji

import (
	"net/http"
	"net/url"
)

func (c *Client) NewGetPropertiesRequest() GetPropertiesRequest {
	return GetPropertiesRequest{
		client:      c,
		queryParams: c.NewGetPropertiesQueryParams(),
		pathParams:  c.NewGetPropertiesPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetPropertiesRequestBody(),
	}
}

type GetPropertiesRequest struct {
	client      *Client
	queryParams *GetPropertiesQueryParams
	pathParams  *GetPropertiesPathParams
	method      string
	headers     http.Header
	requestBody GetPropertiesRequestBody
}

func (c *Client) NewGetPropertiesQueryParams() *GetPropertiesQueryParams {
	return &GetPropertiesQueryParams{}
}

type GetPropertiesQueryParams struct {
	StatusCode   string `schema:"statusCode,omitempty"`
	SubsidiaryID string `schema:"subsidiaryId,omitempty"`
	Recursive    bool   `schema:"recursive,omitempty"`
	Sort         string `schema:"sort,omitempty"`
	PageNumber   int32  `schema:"pageNumber,omitempty"`
	PageSize     int32  `schema:"pageSize,omitempty"`
	Filter       string `schema:"filter,omitempty"`
}

func (p GetPropertiesQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetPropertiesRequest) QueryParams() *GetPropertiesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetPropertiesPathParams() *GetPropertiesPathParams {
	return &GetPropertiesPathParams{}
}

type GetPropertiesPathParams struct {
}

func (p *GetPropertiesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetPropertiesRequest) PathParams() *GetPropertiesPathParams {
	return r.pathParams
}

func (r *GetPropertiesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetPropertiesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetPropertiesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetPropertiesRequestBody() GetPropertiesRequestBody {
	return GetPropertiesRequestBody{}
}

type GetPropertiesRequestBody struct {
}

func (r *GetPropertiesRequest) RequestBody() *GetPropertiesRequestBody {
	return nil
}

func (r *GetPropertiesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetPropertiesRequest) SetRequestBody(body GetPropertiesRequestBody) {
	r.requestBody = body
}

func (r *GetPropertiesRequest) NewResponseBody() *GetPropertiesResponseBody {
	return &GetPropertiesResponseBody{}
}

type GetPropertiesResponseBody PaginatedResponse[PropertyListItem]

func (r *GetPropertiesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("api-gateway/configuration/v1/properties", r.PathParams())
	return &u
}

func (r *GetPropertiesRequest) Do() (GetPropertiesResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
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

func (r *GetPropertiesRequest) All() ([]PropertyListItem, error) {
	properties := []PropertyListItem{}
	for {
		resp, err := r.Do()
		if err != nil {
			return properties, err
		}

		// Break out of loop when no properties are found
		if len(resp.Results) == 0 {
			break
		}

		// Add properties to list
		properties = append(properties, resp.Results...)

		if len(properties) == resp.Paging.TotalCount {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return properties, nil
}
