package models

// VirtualDisk model API
type VirtualDisk struct {
	AlertSummary          AlertSummary      `json:"alert_summary"`
	Stats                 map[string]string `json:"stats"`
	UUID                  string            `json:"uuid"`
	DiskAddress           string            `json:"disk_address"`
	AttachedVMName        string            `json:"attached_v_m_name"`
	VirtualDiskID         string            `json:"virtual_disk_id"`
	DiskCapacityInBytes   int               `json:"disk_capacity_in_bytes"`
	FlashModeEnabled      bool              `json:"flash_mode_enabled"`
	UsageStats            map[string]string `json:"usage_stats"`
	AttachedVMID          string            `json:"attached_vm_id"`
	ClusterUUID           string            `json:"cluster_uuid"`
	HealthSummary         HealthSummary     `json:"health_summary"`
	NutanixNFSFilePath    string            `json:"nutanix_n_f_s_file_path"`
	AttachedVMUUID        string            `json:"attached_vm_uuid"`
	AttachedVolumeGroupID string            `json:"attached_volume_group_id"`
}
