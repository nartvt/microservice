package transport

const (
	ParamLimit     = "limit"
	ParamPage      = "param"
	ParamSectionId = "sectionId"

	ParamUserId = "userId"

	DefaultLimit = 8
	DefgaultPage = 1
)

type Response struct {
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Total   int    `json:"total,omitempty"`
	NextUrl string `json:"next_url,omitempty"`
}
