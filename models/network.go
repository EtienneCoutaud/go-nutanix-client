package models

// NetworkConfig model API
type NetworkConfig struct {
	UUID             string   `json:"uuid"`
	VSwitchName      string   `json:"vswitch_name"`
	Name             string   `json:"name"`
	IPConfig         IPConfig `json:"ip_config"`
	Annotation       string   `json:"annotation"`
	VLANID           int      `json:"vlan_id" binding:"required"`
	LogicalTimestamp int      `json:"logical_timestamp"`
	NetworkType      string   `json:"network_type"`
}

// NetworkConfigCollection model API
type NetworkConfigCollection struct {
	Entities  []*NetworkConfig `json:"entities"`
	Metadata  Metadata         `json:"metadata"`
	ErrorInfo ErrorInfo        `json:"error_info"`
}

// NetworkConfigCreate model API
type NetworkConfigCreate struct {
	UUID             string   `json:"uuid"`
	VSwitchName      string   `json:"vswitch_name"`
	Name             string   `json:"name"`
	IPConfig         IPConfig `json:"ip_config"`
	Annotation       string   `json:"annotation"`
	VLANID           int      `json:"vlan_id" binding:"required"`
	LogicalTimestamp int      `json:"logical_timestamp"`
}

// NetworkCreateResponse model API when created new network
type NetworkCreateResponse struct {
	NetworkUUID string `json:"network_uuid"`
}

// IPConfig model API
type IPConfig struct {
	DefaultGateway    string      `json:"default_gateway"`
	NetworkAddress    string      `json:"network_address" binding:"required"`
	DHCPServerAddress string      `json:"dhcp_server_address"`
	DHCPOptions       DHCPOptions `json:"dhcp_options"`
	PrefixLength      int         `json:"prefix_length"`
	Pool              []*IPPool   `json:"pool"`
}

// IPPool model API
type IPPool struct {
	Range string `json:"range"`
}

// DHCPOptions model API
type DHCPOptions struct {
	DomainSearch      string `json:"domain_search"`
	BootFileName      string `json:"boot_file_name"`
	DomainNameServers string `json:"domain_name_servers"`
	DomainName        string `json:"domain_name"`
	TFTPServerName    string `json:"tftp_server_name"`
}
