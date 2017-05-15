package nutanix

import (
	"fmt"
	"testing"

	"github.com/etiennecoutaud/go-nutanix-client/models"
)

func GenerateVMConfigBasic() *models.VMConfigCreate {
	return &models.VMConfigCreate{
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
