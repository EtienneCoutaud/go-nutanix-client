package models

// MetricDetail model API
type MetricDetail struct {
	MetricDisplayName  string               `json:"metric_display_name"`
	EntityID           string               `json:"entity_id"`
	MetricName         string               `json:"metric_name"`
	DataType           string               `json:"data_type"`
	EntityType         string               `json:"entity_type"`
	MetricValueDetails []*MetricValueDetail `json:"metric_value_details"`
	ComparisonOperator string               `json:"comparison_operator"`
	Unit               string               `json:"unit"`
}

// MetricValueDetail model API
type MetricValueDetail struct {
	StateChangeTimeStampInUsecs int    `json:"state_change_time_stamp_in_usecs"`
	MetricThresholdValue        string `json:"metric_threshold_value"`
	MetricValue                 string `json:"metric_value"`
}
