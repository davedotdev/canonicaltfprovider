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
			"string1": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"int1": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"bool1": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"float1": &schema.Schema{
				Type:     schema.TypeFloat,
				Required: true,
			},
			"map1": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"set2": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"int2": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"string2": {
							Type:     schema.TypeString,
							Required: true,
						},
						"bool2": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"float2": {
							Type:     schema.TypeFloat,
							Required: true,
						},
					},
				},
				MinItems: 1,
				MaxItems: 1,
			},
			"listset3": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"int3": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"string3": {
							Type:     schema.TypeString,
							Required: true,
						},
						"bool3": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"float3": {
							Type:     schema.TypeFloat,
							Required: true,
						},
					},
				},
				MinItems: 1,
				MaxItems: 1,
			},
			"setnested4": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"int4": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"string4": {
							Type:     schema.TypeString,
							Required: true,
						},
						"bool4": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"float4": {
							Type:     schema.TypeFloat,
							Required: true,
						},
						"set5": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"int5": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"string5": {
										Type:     schema.TypeString,
										Required: true,
									},
									"bool5": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"float5": {
										Type:     schema.TypeFloat,
										Required: true,
									},
									"set6": {
										Type:     schema.TypeSet,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"int6": {
													Type:     schema.TypeInt,
													Required: true,
												},
												"string6": {
													Type:     schema.TypeString,
													Required: true,
												},
												"bool6": {
													Type:     schema.TypeBool,
													Required: true,
												},
												"float6": {
													Type:     schema.TypeFloat,
													Required: true,
												},
											},
										},
									},
									"listset7": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"int7": {
													Type:     schema.TypeInt,
													Required: true,
												},
												"string7": {
													Type:     schema.TypeString,
													Required: true,
												},
												"bool7": {
													Type:     schema.TypeBool,
													Required: true,
												},
												"float7": {
													Type:     schema.TypeFloat,
													Required: true,
												},
											},
										},
									},
								},
							},
						},
						"listset8": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"int8": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"string8": {
										Type:     schema.TypeString,
										Required: true,
									},
									"bool8": {
										Type:     schema.TypeBool,
										Required: true,
									},
									"float8": {
										Type:     schema.TypeFloat,
										Required: true,
									},
								},
							},
						},
					},
				},
				MinItems: 1,
				MaxItems: 1,
			},
			"liststring9": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}
}
