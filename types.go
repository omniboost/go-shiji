package shiji

type PaginatedResponse[T any] struct {
	Results []T `json:"results"`
	Paging  struct {
		PageNumber int `json:"pageNumber"`
		PageSize   int `json:"pageSize"`
		TotalCount int `json:"totalCount"`
	} `json:"paging"`
}

type ConfigurationSupportedType string
