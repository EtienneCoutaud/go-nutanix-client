package models

// ResultContentDescriptorEntity model API
type ResultContentDescriptorEntity struct {
	Entities  []*ResultContentDescriptor `json:"entities"`
	Metadata  Metadata                   `json:"metadata"`
	ErrorInfo ErrorInfo                  `json:"error_info"`
}

// ResultContentDescriptor model API
type ResultContentDescriptor struct {
	Perspectives []*Perspective `json:"perspectives"`
	Metadata     Metadata       `json:"metadata"`
	ErrorInfo    ErrorInfo      `json:"error_info"`
}

// Perspective model API
type Perspective struct {
	Metadata ResultMetadata `json:"metadata"`
	Type     string         `json:"string"`
	Format   string         `json:"format"`
}
