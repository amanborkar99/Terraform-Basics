package zoom_corp

import (
	"context"
	"os"
    //e "errors"

	//xl "terraform-provider-zoom/Client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)
type client struct{
	token string
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			// "Username": &schema.Schema{
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },
			// "Password": &schema.Schema{
			// 	Type:      schema.TypeString,
			// 	Sensitive: true,
			// 	Optional:  true,
			// },
			 "token": &schema.Schema{
			 	Type:      schema.TypeString,
			 	Optional:  true,
				//DefaultFunc: schema.EnvDefaultFunc("TOKEN_ZOOM", nil),

			 },
		},
		ResourcesMap: map[string]*schema.Resource{

		},
		DataSourcesMap: map[string]*schema.Resource{
           "zoom_corp_data": dataSourceUser(),
		},
		ConfigureContextFunc: providerConfig,
	}
}

func providerConfig(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics){
   
	//us_name := d.Get("Username").(string)
	//us_pass := d.Get("Password").(string)
	us_token := d.Get("token").(string)
    
	os.Setenv("ttk", us_token)
	var diags diag.Diagnostics
	 c:=client{}
	 c.token = us_token
	  return c, diags

   

	}
		
	
	
		

	//return c, diags

