package models

// Metadata model API (http://developer.nutanix.com/reference/v2/?go#definitions-get.base.Metadata)
type Metadata struct {
	Count              int    `json:"count"`
	FilterCriteria     string `json:"filter_criteria"`
	GrandTotalEntities string `json:"grand_total_entities"`
	PreviousCursor     string `json:"previous_cursor"`
	SortCriteria       string `json:"sort_criteria"`
	EndIndex           string `json:"end_index"`
	SearchString       string `json:"search_string"`
	NextCursor         string `json:"next_cursor"`
	TotalEntities      string `json:"total_entities"`
	Page               string `json:"page"`
	StartIndex         string `json:"start_index"`
}
