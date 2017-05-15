package nutanix

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/etiennecoutaud/go-nutanix-sdk/client"
	"github.com/etiennecoutaud/go-nutanix-sdk/endpoint"
	"github.com/etiennecoutaud/go-nutanix-sdk/models"
)

// Client struct to request API
type Client struct {
	*client.Client
}

// GetVMS
func (c *Client) GetVMS() (*models.VMConfigGet, error) {
	req, err := c.NewRequest("GET", endpoint.VMS, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	resdata, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result models.VMConfigGet
	err = json.Unmarshal(resdata, &resdata)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Fatal("Error when created VM, code:", resp.StatusCode)
		return nil, errors.New("Fail to create VM")
	}
	return &result, nil
}

// CreateVM  POST, Use to create a new VM (http://developer.nutanix.com/reference/v2/?python#vms-post)
func (c *Client) CreateVM(vm *models.VMConfigCreate) (*models.VMCreateResponse, error) {
	if vm == nil {
		return nil, errors.New("VM Body nil")
	}
	data, err := json.Marshal(vm)
	if err != nil {
		return nil, err
	}
	req, err := c.NewRequest("POST", endpoint.VMS, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	resdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result models.VMCreateResponse
	err = json.Unmarshal(resdata, &result)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Fatal("Error when created VM, code:", resp.StatusCode)
		return nil, errors.New("Fail to create VM")
	}
	return &result, nil
}
