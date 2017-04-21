package tvdb

//UpdatedRequest must be used when making the requst
type UpdatedRequest struct {
	FromTime int64 `json:"fromTime"`
	ToTime   int64 `json:"toTime"`
}

//UpdatedResponse a single UpdatedResponse
type UpdatedResponse struct {
	ID          int `json:"id"`
	LastUpdated int `json:"lastUpdated"`
}

//UpdatedResponses contains the response returned by the API Server
type UpdatedResponses struct {
	Data []UpdatedResponse `json:"data"`
}
