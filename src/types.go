package main

type RecordPageViewRequest struct {
	Path string `json:"path"`
}

type RecordClickRequest struct {
	Path    string `json:"path"`
	Element string `json:"element"`
}

type RecordPageTimeRequest struct {
	Path      string `json:"path"`
	TimeSpent int64  `json:"time_spent"`
}
