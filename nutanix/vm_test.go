package nutanix

import (
	"fmt"
	"testing"

	"github.com/etiennecoutaud/go-nutanix-sdk/models"
)

func GenerateVMConfigBasic() *models.VMCreateConfig {
	return &models.VMCreateConfig{
		Name:     "test",
		MemoryMB: 1,
	}
}

func TestCreateVM(t *testing.T) {
	var client Client
	// Nil test
	_, err := client.CreateVM(nil)
	fmt.Println(err)
}
