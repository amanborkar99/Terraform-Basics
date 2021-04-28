package zoom_corp

import (
	"fmt"
	"testing"
	"regexp"

	//"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	//"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	xl "terraform-provider-zoom/Client"
)

func TestAccUser_Basic(t *testing.T){
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: testAccCheckItemDestroy,
		Steps: []resource.TestStep{
			{
              Config: testAccCheckUserBasic(),
			  Check: resource.ComposeTestCheckFunc(
				  testAccCheckZoomUserExists("zoom_user.user"),
				  resource.TestCheckResourceAttr(
					"zoom_resource.user", "email", "amanborkar17@cse.iiitp.ac.in"),

				  resource.TestCheckResourceAttr(
					  "zoom_resource.user", "first_name", "Aman" ),

				  resource.TestCheckResourceAttr( 
					  "zoom_user.user", "last_name", "Borkar"),

				  resource.TestCheckResourceAttr( 
						"zoom_user.user", "type", "1"),
			  ),
			},
		},
	})
}

func testAccCheckUserBasic() string {
	return fmt.Sprintf(`
resource "zoom_resource" "user" {
	email = "amanborkar17@cse.iiitp.ac.in"
	first_name = "Aman"
	last_name = "Borkar"
	type = 1
}  
`)
}

func testAccCheckItemDestroy(s *terraform.State) error {
	apiClient := testAccProvider.Meta().(*xl.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "zoom_resource" {
			continue
		}

		_, err := apiClient.GetUser(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Alert still exists")
		}
		notFoundErr := "not found"
		expectedErr := regexp.MustCompile(notFoundErr)
		if !expectedErr.Match([]byte(err.Error())) {
			return fmt.Errorf("expected %s, got %s", notFoundErr, err)
		}
	}

	return nil
}
func testAccCheckZoomUserExists(resource string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("Not found: %s", resource)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}
		name := rs.Primary.ID
		apiClient := testAccProvider.Meta().(*xl.Client)
		_, err := apiClient.GetUser(name)
		if err != nil {
			return fmt.Errorf("error fetching item with resource %s. %s", resource, err)
		}
		return nil
	}
}

func TestAccUser_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckItemDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckUserUpdatePre(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoomUserExists("zoom_user.user"),
				  resource.TestCheckResourceAttr(
					"zoom_resource.user", "email", "amanborkar17@cse.iiitp.ac.in"),

				  resource.TestCheckResourceAttr(
					  "zoom_resource.user", "first_name", "Aman" ),

				  resource.TestCheckResourceAttr( 
					  "zoom_user.user", "last_name", "Borkar"),

				  resource.TestCheckResourceAttr( 
						"zoom_user.user", "type", "1"),
				),
			},
			{
				Config: testAccCheckUserUpdatePost(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoomUserExists("zoom_user.user"),
				  resource.TestCheckResourceAttr(
					"zoom_resource.user", "email", "amanborkar17@cse.iiitp.ac.in"),

				  resource.TestCheckResourceAttr(
					  "zoom_resource.user", "first_name", "Aman" ),

				  resource.TestCheckResourceAttr( 
					  "zoom_user.user", "last_name", "VIPBoss"),

				  resource.TestCheckResourceAttr( 
						"zoom_user.user", "type", "1"),
				),
			},
		},
	})
}

func testAccCheckUserUpdatePre() string {
	return fmt.Sprintf(`
	resource "zoom_resource" "user" {
		email = "amanborkar17@cse.iiitp.ac.in"
		first_name = "Aman"
		last_name = "Borkar"
		type = 1
	}  
`)
}

func testAccCheckUserUpdatePost() string {
	return fmt.Sprintf(`
	resource "zoom_resource" "user" {
		email = "amanborkar17@cse.iiitp.ac.in"
		first_name = "Aman"
		last_name = "VIPBoss"
		type = 1
	}  
`)
}
