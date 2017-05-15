package nutanix

import (
	"encoding/json"

	"github.com/etiennecoutaud/go-nutanix-client/endpoint"
	"github.com/etiennecoutaud/go-nutanix-client/models"
)

func (c *Client) CreateNetwork(network *models.NetworkConfigCreate) (*models.NetworkCreateResponse, error) {
	resdata, err := c.PostRequest(endpoint.NETWORKS, network)
	var result models.NetworkCreateResponse
	err = json.Unmarshal(resdata, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
