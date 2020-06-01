/*
***********************************************
Licensed under BSD-3-Clause - see license file.
***********************************************
*/

package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider returns a Terraform ResourceProvider.
func Provider() *schema.Provider {
	return &schema.Provider{

		ResourcesMap: map[string]*schema.Resource{
			"canonical": Canonical(),
		},
	}
}
