package models

// VMConfig model Api (http://developer.nutanix.com/reference/v2/#definitions-create.dto.uhura.VmConfigDTO)
type VMConfig struct {
	Name                  string                `json:"name" binding:"required"`
	MemomryMB             int                   `json:"memory_mb" binding:"required"`
	NumVCPUS              int                   `json:"num_vcpus" binding:"required"`
	VMFeatures            bool                  `json:"vm_features"`
	VMCustomizationConfig VMCustomizationConfig `json:"vm_customization_config"`
	UUID                  string                `json:"uuid"`
	Description           string                `json:"description"`
	NumCorePerVCPU        int                   `json:"num_cores_per_vcpu"`
	MemoryReservationMB   int                   `json:"memory_reservation_mb"`
	VMGPUS                []*VMGPU              `json:"vm_gpus"`
	Boot                  BootConfig            `json:"boot"`
	GuestOS               string                `json:"guest_os"`
	Timezone              string                `json:"timezone"`
	VCPUReservationHZ     int                   `json:"vcpu_reservation_hz"`
	VMNics                []*VMNic              `json:"vm_nics"`
	HAPriority            int                   `json:"ha_priority"`
	VMDisks               []*VMDisk             `json:"vm_disks"`
	StorageContainerUUID  string                `json:"storage_container_uuid"`
}

// VMCustomizationConfig model API (http://developer.nutanix.com/reference/v2/#definitions-create.dto.acropolis.VMCustomizationConfigDTO)
type VMCustomizationConfig struct {
	UserdataPath      string          `json:"userdata_path"`
	FreshInstall      bool            `json:"fresh_install"`
	FilesToInjectList []*FileToInject `json:"files_to_inject_list"`
	DataSourceType    string          `json:"datasource_type"`
	UserData          string          `json:"userdata"`
}

// FileToInject model API (http://developer.nutanix.com/reference/v2/#definitions-create.dto.acropolis.FileToInjectDTO)
type FileToInject struct {
	DestinationPath string `json:"destination_path"`
	SourcePath      string `json:"source_path"`
}

// VMGPU model API (http://developer.nutanix.com/reference/v2/#definitions-create.dto.uhura.GPUConfigDTO)
type VMGPU struct {
	GPUType   string `json:"gpu_type" binding:"required"`
	DeviceID  int    `json:"device_id" binding:"required"`
	GPUVendor string `json:"gpu_vendor"`
}

// BootConfig model API (http://developer.nutanix.com/reference/v2/#definitions-create.dto.acropolis.BootConfigDTO)
type BootConfig struct {
	BootDeviceType string        `json:"boot_device_type" binding:"required"`
	MACAddress     string        `json:"mac_address"`
	DiskAddress    VMDiskAddress `json:"disk_address"`
}

// VMNic model API (http://developer.nutanix.com/reference/v2/#definitions-create.dto.acropolis.VMNicSpecDTO)
type VMNic struct {
	NetworkUUID        string `json:"network_uuid" binding:"required"`
	RequestIP          bool   `json:"request_ip"`
	AdapterType        string `json:"adapter_type"`
	MACAddress         string `json:"mac_address"`
	RequestedIPAddress string `json:"requested_ip_address"`
	Model              string `json:"model"`
}

// VMDiskCreate model API (http://developer.nutanix.com/reference/v2/#definitions-create.dto.uhura.VmDiskCreateDTO)
type VMDiskCreate struct {
	StorageContainerUUID string `json:"storage_container_uuid"`
	Size                 int    `json:"storage"`
}

// VMDiskClone model API (http://developer.nutanix.com/reference/v2/#definitions-create.dto.uhura.VmDiskCloneDTO)
type VMDiskClone struct {
	StorageContainerUUID string        `json:"storage_container_uuid"`
	MinimumSize          int           `json:"minimum_size"`
	SnapshotGroupUUID    string        `json:"snapshot_group_uuid"`
	DiskAddress          VMDiskAddress `json:"disk_address"`
}

// VMDisk model API (http://developer.nutanix.com/reference/v2/#definitions-create.dto.uhura.VMDiskDTO)
type VMDisk struct {
	DiskAddress       VMDiskAddress `json:"disk_address"`
	IsCDROM           bool          `json:"is_cdrom"`
	VMDiskCreate      VMDiskCreate  `json:"vm_disk_create"`
	IsScsiPassThrough bool          `json:"is_scsi_pass_through"`
	IsEmpty           bool          `json:"is_empty"`
	IsThinProvisioned bool          `json:"is_thin_provisioned"`
	VMDiskClone       VMDiskClone   `json:"vm_disk_clone"`
}

// VMDiskAddress model API (http://developer.nutanix.com/reference/v2/#definitions-create.dto.acropolis.VMDiskAddressDTO)
type VMDiskAddress struct {
	NDFSFilePath    string `json:"ndfs_filepath"`
	VMDiskUUID      string `json:"vmdisk_uuid"`
	VolumeGroupUUID string `json:"volume_group_uuid"`
	DeviceIndex     int    `json:"device_index"`
	DeviceBus       string `json:"device_bus"`
}
