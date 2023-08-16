package model

type SearchResult struct {
	ResultCount int `json:"resultCount"`
	Results []*Podcast `json:"results"`
}