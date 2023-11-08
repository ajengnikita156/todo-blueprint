package model


type BulkDeleteRequest struct {
	ID []int `json:"id"`
}

type Count struct {
	Pending  int `json:"pending"`
	Progress int `json:"progress"`
	Done     int `json:"done"`
}