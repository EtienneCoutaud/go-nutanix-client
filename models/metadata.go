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

// ResultMetadata model API
type ResultMetadata struct {
	Instances  []*EntityMetadata `json:"instances"`
	Title      string            `json:"title"`
	Filters    []*Filter         `json:"filters"`
	EntityType string            `json:"entity_type"`
}

// EntityMetadata model API
type EntityMetadata struct {
	AdditionalFields map[string]string `json:"additional_fields"`
	EntityID         string            `json:"entity_id"`
	ContextHintField string            `json:"context_hint_field"`
	EntityName       string            `json:"entity_name"`
	EntityType       string            `json:"entity_type"`
	ClusterUUID      string            `json:"cluster_uuid"`
	ContextHintValue string            `json:"context_hint_value"`
}

// Filter model API
type Filter struct {
	Operator      string `json:"operator"`
	RHS           string `json:"rhs"`
	LHS           string `json:"lhs"`
	DisplayForRHS string `json:"display_for_rhs"`
}
