package hashicups

import (
	"context"

	"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("h_name", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("h_pass", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"hashicups_coffees": dataSourceCoffees(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(cmt context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	u := d.Get("username").(string)
	p := d.Get("password").(string)

	var diags diag.Diagnostics

	if (u != "") && (p != "") {
		c, err := hashicups.NewClient(nil, &u, &p)
		if err != nil {
			return nil, diag.FromErr(err)
		}
		return c, diags
	}

	c, err := hashicups.NewClient(nil, nil, nil)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags

}
