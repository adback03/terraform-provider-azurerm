package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
)

func TestAccDataSourceAzureRMPPostgreSqlServer_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "data.azurerm_postgresql_server", "test")
	version := "9.5"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckAzureRMPostgreSQLServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAzureRMPostgreSqlServer_basic(data, version),
				Check: resource.ComposeTestCheckFunc(
					testCheckAzureRMPostgreSQLServerExists(data.ResourceName),
					resource.TestCheckResourceAttrSet(data.ResourceName, "location"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "fqdn"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "version"),
					resource.TestCheckResourceAttrSet(data.ResourceName, "administrator_login"),
					resource.TestCheckResourceAttr(data.ResourceName, "tags.%", "0"),
				),
			},
		},
	})
}

func testAccDataSourceAzureRMPostgreSqlServer_basic(data acceptance.TestData, version string) string {
	template := testAccAzureRMPostgreSQLServer_basic(data, version)
	return fmt.Sprintf(`
%s

data "azurerm_postgresql_server" "test" {
  name                = "${azurerm_postgresql_server.test.name}"
  resource_group_name = "${azurerm_resource_group.test.name}"
}
`, template)
}
