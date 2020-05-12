/*

***********************************************
Licensed under BSD-3-Clause - see license file.
***********************************************

Release Grade

__Helpful Info__

https://github.com/hashicorp/terraform/issues/17579
Talks about how resource names are external to resources and not accessible by code.
Separation of concerns etc.

*/

package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func tCreate(d *schema.ResourceData, m interface{}) error {
	// No purpose

	return tRead(d, m)
}

func tRead(d *schema.ResourceData, m interface{}) error {

	return nil
}

func tUpdate(d *schema.ResourceData, m interface{}) error {

	return tRead(d, m)
}

func tDelete(d *schema.ResourceData, m interface{}) error {

	return nil
}

// Canonical is the schema for this canonical resource
func Canonical() *schema.Resource {
	return &schema.Resource{
		Create: tCreate,
		Read:   tRead,
		Update: tUpdate,
		Delete: tDelete,

		Schema: map[string]*schema.Schema{
			"stringtest": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"inttest": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"booltest": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"floattest": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true,
			},
			"maptest": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"settest": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"setinttest": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"setstringtest": {
							Type:     schema.TypeString,
							Required: true,
						},
						"setbooltest": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"setfloattest": {
							Type:     schema.TypeFloat,
							Required: true,
						},
					},
				},
				MinItems: 1,
				MaxItems: 1,
			},
			"listsettest": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"setinttest": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"setstringtest": {
							Type:     schema.TypeString,
							Required: true,
						},
						"setbooltest": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"setfloattest": {
							Type:     schema.TypeFloat,
							Required: true,
						},
					},
				},
				MinItems: 1,
				MaxItems: 1,
			},
			"liststringtest": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				MinItems: 1,
				MaxItems: 1,
			},
		},
	}
}
