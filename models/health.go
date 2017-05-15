package models

// HealthSummary model API
type HealthSummary struct {
	EntityTypeSummaries  []*EntityTypeHealthCheckSummary `json:"entity_type_summaries"`
	HealthCheckSummaries map[string]HealthCheckSummary   `json:"health_check_summaries"`
	HealthStatus         string                          `json:"health_status"`
}

// HealthCheckSummary model API
type HealthCheckSummary struct{}

// EntityTypeHealthCheckSummary model API
type EntityTypeHealthCheckSummary struct {
	FilterCriteria       string                          `json:"filter_criteria"`
	HealthSummary        map[HealthStatus]int            `json:"health_summary"`
	EntityType           string                          `json:"entity_type"`
	DetailedCheckSummary map[string]map[HealthStatus]int `json:"detailed_check_summary"`
	ChecksInError        map[string]HealthCheckErrors    `json:"checks_in_error"`
}

// HealthCheckErrors TODO
type HealthCheckErrors struct{}

// HealthStatus TODO
type HealthStatus struct{}
