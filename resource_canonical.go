package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Set2 type
type Set2 struct {
	Int2    int
	String2 string
	Bool2   bool
	Float2  float64
}

// Listset3 type
type Listset3 struct {
	Int3    int
	String3 string
	Bool3   bool
	Float3  float64
}

// Setnested5 type
type Setnested5 struct {
	Int5       int
	String5    string
	Bool5      bool
	Float5     float64
	Setnested6 Setnested6
	Listset7   []Listset7
}

// Setnested4 type
type Setnested4 struct {
	Int4       int
	String4    string
	Bool4      bool
	Float4     float64
	Setnested5 Setnested5
	Listset8   []Listset8
}

// Setnested6 type
type Setnested6 struct {
	Int6    int
	String6 string
	Bool6   bool
	Float6  float64
}

// Listset7 type
type Listset7 struct {
	Int7    int
	String7 string
	Bool7   bool
	Float7  float64
}

// Listset8 type
type Listset8 struct {
	Int8    int
	String8 string
	Bool8   bool
	Float8  float64
}

// Contains all the data to do some pretend read/writes
type tCanonicalType struct {
	String1     string
	Int1        int
	Bool1       bool
	Float1      float64
	Map1        map[string]interface{}
	Set2        *schema.Set
	Listset3    []interface{}
	Setnested4  interface{}
	Liststring9 []interface{}
}

// Create func
func Create(d *schema.ResourceData, m interface{}) error {

	return nil

}

// Read func
func Read(d *schema.ResourceData, m interface{}) error {

	return nil
}

// Update func
func Update(d *schema.ResourceData, m interface{}) error {

	return nil

}

// Delete func
func Delete(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}

// Canonical is the schema for this canonical resource
func Canonical() *schema.Resource {
	return &schema.Resource{
		Create: Create,
		Read:   Read,
		Update: Update,
		Delete: Delete,

		Schema: map[string]*schema.Schema{
			"resource_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
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
				Type:     schema.TypeMap,
				Required: true,
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
						"setnested5": {
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
									"setnested6": {
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
					Type: schema.TypeString,
				},
			},
		},
	}
}
