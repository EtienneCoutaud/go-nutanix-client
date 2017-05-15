package nutanix

import (
	"bytes"
	"encoding/json"

	"github.com/etiennecoutaud/go-nutanix-client/endpoint"
	"github.com/etiennecoutaud/go-nutanix-client/models"
)

// GetVM GET, Use to get specific VM from uuid
func (c *Client) GetVM(uuid string) (*models.VMConfig, error) {
	// Build uri "/vms/{uuid}?include_vm_disk_config=true&include_vm_nic_config=true"
	var buffer bytes.Buffer
	buffer.WriteString(endpoint.VMS)
	buffer.WriteString(uuid)
	buffer.WriteString("?include_vm_nic_config=true")

	resdata, err := c.GetRequest(buffer.String())
	if err != nil {
		return nil, err
	}
	var result models.VMConfig
	err = json.Unmarshal(resdata, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetVMS GET, Use to get all VM
func (c *Client) GetVMS() (*models.VMConfigGet, error) {
	resdata, err := c.GetRequest(endpoint.VMS)
	if err != nil {
		return nil, err
	}
	var result models.VMConfigGet
	err = json.Unmarshal(resdata, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateVM  POST, Use to create a new VM (http://developer.nutanix.com/reference/v2/?python#vms-post)
func (c *Client) CreateVM(vm *models.VMConfigCreate) (*models.VMCreateResponse, error) {
	resdata, err := c.PostRequest(endpoint.VMS, vm)
	var result models.VMCreateResponse
	err = json.Unmarshal(resdata, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
