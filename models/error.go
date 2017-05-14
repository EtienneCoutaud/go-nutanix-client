package models

// ErrorInfo model API (http://developer.nutanix.com/reference/v2/?python#definitions-get.base.ErrorInfo)
type ErrorInfo struct {
	Count      int      `json:"count"`
	Message    string   `json:"message"`
	EntityList []string `json:"entity_list"`
}
