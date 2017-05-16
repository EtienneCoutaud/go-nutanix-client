package nutanix

import (
	"bytes"
	"encoding/json"

	"github.com/etiennecoutaud/go-nutanix-client/endpoint"
	"github.com/etiennecoutaud/go-nutanix-client/models"
)

// GetVM GET, Use to get specific VM from uuid
func (c *Client) GetVM(uuid string) (*models.VM, error) {
	// Build uri "/vms/{uuid}?include_vm_disk_config=true&include_vm_nic_config=true"
	var buffer bytes.Buffer
	buffer.WriteString(endpoint.VMS)
	buffer.WriteString(uuid)
	buffer.WriteString("?include_vm_nic_config=true")

	resdata, err := c.GetRequest(buffer.String())
	if err != nil {
		return nil, err
	}
	var result models.VM
	err = json.Unmarshal(resdata, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetVMS GET, Use to get all VM
func (c *Client) GetVMS() (*models.VMCollection, error) {
	resdata, err := c.GetRequest(endpoint.VMS)
	if err != nil {
		return nil, err
	}
	var result models.VMCollection
	err = json.Unmarshal(resdata, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateVM  POST, Use to create a new VM (http://developer.nutanix.com/reference/v2/?python#vms-post)
func (c *Client) CreateVM(vm *models.VMConfig) (*models.ResponseTasks, error) {
	resdata, err := c.PostRequest(endpoint.VMS, vm)
	if err != nil {
		return nil, err
	}
	var result models.ResponseTasks
	err = json.Unmarshal(resdata, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateVM PUT, Use to update VM
func (c *Client) UpdateVM(vm *models.VMConfig) (*models.ResponseTasks, error) {
	var buffer bytes.Buffer
	buffer.WriteString(endpoint.VMS)
	buffer.WriteString(vm.UUID)

	resdata, err := c.PutRequest(buffer.String(), vm)
	if err != nil {
		return nil, err
	}
	var result models.ResponseTasks
	err = json.Unmarshal(resdata, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteVM DELETE, Use to delete specific VM from uuid
func (c *Client) DeleteVM(uuid string) (*models.ResponseTasks, error) {
	var buffer bytes.Buffer
	buffer.WriteString(endpoint.VMS)
	buffer.WriteString(uuid)

	resdata, err := c.DeleteRequest(buffer.String())
	if err != nil {
		return nil, err
	}
	var result models.ResponseTasks
	err = json.Unmarshal(resdata, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
