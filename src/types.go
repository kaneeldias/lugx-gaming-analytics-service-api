package main

type RecordPageViewRequest struct {
	Path string `json:"path"`
}

type RecordClickRequest struct {
	Path    string `json:"path"`
	Element string `json:"element"`
}
