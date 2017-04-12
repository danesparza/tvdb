package tvdb

type UpdatedRequest struct {
	FromTime int64 `json:"fromTime"`
	ToTime   int64 `json:"toTime"`
}

type UpdatedResponse struct {
	Id          int `json:"id"`
	LastUpdated int `json:"lastUpdated"`
}

type UpdatedResponses struct {
	Data []UpdatedResponse `json:"data"`
}
