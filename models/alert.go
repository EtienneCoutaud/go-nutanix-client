package models

// Alert model API
type Alert struct {
	Resolved                      bool              `json:"resolved"`
	PossibleCauses                []*AlertCause     `json:"possible_causes"`
	ResolvedTimeStampInUsecs      int               `json:"resolved_time_stamp_in_usecs"`
	AlertDetails                  AlertDetails      `json:"alert_details"`
	ContextValues                 []string          `json:"context_values"`
	ClusterUUID                   string            `json:"cluster_uuid"`
	Message                       string            `json:"message"`
	LastOccurenceTimeStampInUsecs int               `json:"last_occurence_time_stamps_in_usecs"`
	ID                            string            `json:"id"`
	Severity                      string            `json:"severity"`
	CheckID                       string            `json:"check_id"`
	UserDefined                   bool              `json:"user_defined"`
	AlertTitle                    string            `json:"alert_title"`
	DetailedMessage               string            `json:"detailed_message"`
	ImpactTypes                   []string          `json:"impact_types"`
	CreatedTimeStampInUsecs       int               `json:"created_time_stamp_in_usecs"`
	AffectedEntities              []*AffectedEntity `json:"affected_entities"`
	AutoResolved                  bool              `json:"auto_resolved"`
	ContextTypes                  []string          `json:"context_types"`
	AcknowledgedByUsername        string            `json:"acknowledged_by_username"`
	ServiceVMID                   string            `json:"service_v_m_id"`
	NodeUUID                      string            `json:"node_uuid"`
	Classifications               []string          `json:"classifications"`
}

// AffectedEntity model API
type AffectedEntity struct {
	UUID                  string `json:"uuid"`
	EntityTypeDisplayName string `json:"entity_type_display_name"`
	ID                    string `json:"id"`
	EntityName            string `json:"entity_name"`
	EntityType            string `json:"entity_type"`
}

// AlertDetails model API
type AlertDetails struct {
	SeverityTrails []*SeverityTrail `json:"severity_trails"`
	MetricDetails  []*MetricDetail  `json:"metric_details"`
}

// SeverityTrail model API
type SeverityTrail struct {
	Severity                       string `json:"severity"`
	SeverityChangeTimestampInUsecs int    `json:"severity_change_timestamp_in_usecs"`
}

// AlertSummary model API
type AlertSummary struct {
	AlertSummaries map[string]AlertCollection `json:"alert_summaries"`
}

// AlertCollection model API
type AlertCollection struct {
	Entities  []*Alert  `json:"entities"`
	Metadata  Metadata  `json:"metadata"`
	ErrorInfo ErrorInfo `json:"error_info"`
}

// AlertCause model API
type AlertCause struct {
	HasGlobalDetails     string                        `json:"has_global_details"`
	AlertNamespace       string                        `json:"alert_namespace"`
	Title                string                        `json:"title"`
	ChartQueries         ResultContentDescriptorEntity `json:"chart_queries"`
	AdditionalLinks      string                        `json:"additional_links"`
	Actions              string                        `json:"actions"`
	ParentCauseID        string                        `json:"parent_cause_id"`
	Details              string                        `json:"details"`
	TroubleshootingSteps string                        `json:"troubleshooting_steps"`
	Cause                string                        `json:"cause"`
	ID                   string                        `json:"id"`
	NumResolvedCount     int                           `json:"num_resolved_count"`
}
