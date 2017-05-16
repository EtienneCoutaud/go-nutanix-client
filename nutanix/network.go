package nutanix

import (
	"bytes"
	"encoding/json"

	"github.com/etiennecoutaud/go-nutanix-client/endpoint"
	"github.com/etiennecoutaud/go-nutanix-client/models"
)

func (c *Client) CreateNetwork(network *models.NetworkConfigCreate) (*models.NetworkCreateResponse, error) {
	resdata, err := c.PostRequest(endpoint.NETWORKS, network)
	if err != nil {
		return nil, err
	}
	var result models.NetworkCreateResponse
	err = json.Unmarshal(resdata, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetNetworks() (*models.NetworkConfigCollection, error) {
	resdata, err := c.GetRequest(endpoint.NETWORKS)
	if err != nil {
		return nil, err
	}
	var result models.NetworkConfigCollection
	err = json.Unmarshal(resdata, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetNetwork(uuid string) (*models.NetworkConfig, error) {
	var buffer bytes.Buffer
	buffer.WriteString(endpoint.NETWORKS)
	buffer.WriteString(uuid)

	resdata, err := c.GetRequest(buffer.String())
	if err != nil {
		return nil, err
	}
	var result models.NetworkConfig
	err = json.Unmarshal(resdata, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
