package models

type EsResult struct {
	Total     int64
	PageIndex int
	PageSize  int
	Result    interface{}
}
