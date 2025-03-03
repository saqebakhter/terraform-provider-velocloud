package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	velo "github.com/adeleporte/terraform-provider-velocloud/velocloud"
)

func dataSourceGateway() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGatewayRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enterpriseid": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceGatewayRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(*velo.Client)

	if !client.Operator {
		return diag.Errorf("Not logged as an operator")
	}

	EnterpriseName := d.Get("name").(string)
	id, err := velo.GetGateway(client, EnterpriseName)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%d", id))

	return diags
}
