// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceAccessContextManagerServicePerimeters() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessContextManagerServicePerimetersCreate,
		Read:   resourceAccessContextManagerServicePerimetersRead,
		Update: resourceAccessContextManagerServicePerimetersUpdate,
		Delete: resourceAccessContextManagerServicePerimetersDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAccessContextManagerServicePerimetersImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(6 * time.Minute),
			Update: schema.DefaultTimeout(6 * time.Minute),
			Delete: schema.DefaultTimeout(6 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"parent": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The AccessPolicy this ServicePerimeter lives in.
Format: accessPolicies/{policy_id}`,
			},
			"service_perimeters": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: `The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.`,
				Elem:        accesscontextmanagerServicePerimetersServicePerimetersSchema(),
				// Default schema.HashSchema is used.
			},
		},
	}
}

func accesscontextmanagerServicePerimetersServicePerimetersSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Resource name for the ServicePerimeter. The short_name component must
begin with a letter and only include alphanumeric and '_'.
Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}`,
			},
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Human readable title. Must be unique within the Policy.`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `Description of the ServicePerimeter and its use. Does not affect
behavior.`,
			},
			"perimeter_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"PERIMETER_TYPE_REGULAR", "PERIMETER_TYPE_BRIDGE", ""}, false),
				Description: `Specifies the type of the Perimeter. There are two types: regular and
bridge. Regular Service Perimeter contains resources, access levels,
and restricted services. Every resource can be in at most
ONE regular Service Perimeter.

In addition to being in a regular service perimeter, a resource can also
be in zero or more perimeter bridges. A perimeter bridge only contains
resources. Cross project operations are permitted if all effected
resources share some perimeter (whether bridge or regular). Perimeter
Bridge does not contain access levels or services: those are governed
entirely by the regular perimeter that resource is in.

Perimeter Bridges are typically useful when building more complex
topologies with many independent perimeters that need to share some data
with a common perimeter, but should not be able to share data among
themselves. Default value: "PERIMETER_TYPE_REGULAR" Possible values: ["PERIMETER_TYPE_REGULAR", "PERIMETER_TYPE_BRIDGE"]`,
				Default: "PERIMETER_TYPE_REGULAR",
			},
			"spec": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Proposed (or dry run) ServicePerimeter configuration.
This configuration allows to specify and test ServicePerimeter configuration
without enforcing actual access restrictions. Only allowed to be set when
the 'useExplicitDryRunSpec' flag is set.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_levels": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `A list of AccessLevel resource names that allow resources within
the ServicePerimeter to be accessed from the internet.
AccessLevels listed must be in the same policy as this
ServicePerimeter. Referencing a nonexistent AccessLevel is a
syntax error. If no AccessLevel names are listed, resources within
the perimeter can only be accessed via GCP calls with request
origins within the perimeter. For Service Perimeter Bridge, must
be empty.

Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"resources": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `A list of GCP resources that are inside of the service perimeter.
Currently only projects are allowed.
Format: projects/{project_number}`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"restricted_services": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `GCP services that are subject to the Service Perimeter
restrictions. Must contain a list of services. For example, if
'storage.googleapis.com' is specified, access to the storage
buckets inside the perimeter must meet the perimeter's access
restrictions.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"vpc_accessible_services": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `Specifies how APIs are allowed to communicate within the Service
Perimeter.`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowed_services": {
										Type:     schema.TypeList,
										Optional: true,
										Description: `The list of APIs usable within the Service Perimeter.
Must be empty unless 'enableRestriction' is True.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"enable_restriction": {
										Type:     schema.TypeBool,
										Optional: true,
										Description: `Whether to restrict API calls within the Service Perimeter to the
list of APIs specified in 'allowedServices'.`,
									},
								},
							},
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `ServicePerimeter configuration. Specifies sets of resources,
restricted services and access levels that determine
perimeter content and boundaries.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_levels": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `A list of AccessLevel resource names that allow resources within
the ServicePerimeter to be accessed from the internet.
AccessLevels listed must be in the same policy as this
ServicePerimeter. Referencing a nonexistent AccessLevel is a
syntax error. If no AccessLevel names are listed, resources within
the perimeter can only be accessed via GCP calls with request
origins within the perimeter. For Service Perimeter Bridge, must
be empty.

Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"resources": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `A list of GCP resources that are inside of the service perimeter.
Currently only projects are allowed.
Format: projects/{project_number}`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"restricted_services": {
							Type:     schema.TypeSet,
							Optional: true,
							Description: `GCP services that are subject to the Service Perimeter
restrictions. Must contain a list of services. For example, if
'storage.googleapis.com' is specified, access to the storage
buckets inside the perimeter must meet the perimeter's access
restrictions.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set: schema.HashString,
						},
						"vpc_accessible_services": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `Specifies how APIs are allowed to communicate within the Service
Perimeter.`,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowed_services": {
										Type:     schema.TypeSet,
										Optional: true,
										Description: `The list of APIs usable within the Service Perimeter.
Must be empty unless 'enableRestriction' is True.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Set: schema.HashString,
									},
									"enable_restriction": {
										Type:     schema.TypeBool,
										Optional: true,
										Description: `Whether to restrict API calls within the Service Perimeter to the
list of APIs specified in 'allowedServices'.`,
									},
								},
							},
						},
					},
				},
			},
			"use_explicit_dry_run_spec": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
for all Service Perimeters, and that spec is identical to the status for those
Service Perimeters. When this flag is set, it inhibits the generation of the
implicit spec, thereby allowing the user to explicitly provide a
configuration ("spec") to use in a dry-run version of the Service Perimeter.
This allows the user to test changes to the enforced config ("status") without
actually enforcing them. This testing is done through analyzing the differences
between currently enforced and suggested restrictions. useExplicitDryRunSpec must
bet set to True if any of the fields in the spec are set to non-default values.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the AccessPolicy was created in UTC.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Time the AccessPolicy was updated in UTC.`,
			},
		},
	}
}

func resourceAccessContextManagerServicePerimetersCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	servicePerimetersProp, err := expandAccessContextManagerServicePerimetersServicePerimeters(d.Get("service_perimeters"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service_perimeters"); !isEmptyValue(reflect.ValueOf(servicePerimetersProp)) && (ok || !reflect.DeepEqual(v, servicePerimetersProp)) {
		obj["servicePerimeters"] = servicePerimetersProp
	}
	parentProp, err := expandAccessContextManagerServicePerimetersParent(d.Get("parent"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent"); !isEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{parent}}/servicePerimeters:replaceAll")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ServicePerimeters: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating ServicePerimeters: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{parent}}/servicePerimeters")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = accessContextManagerOperationWaitTime(
		config, res, "Creating ServicePerimeters",
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ServicePerimeters: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ServicePerimeters %q: %#v", d.Id(), res)

	return resourceAccessContextManagerServicePerimetersRead(d, meta)
}

func resourceAccessContextManagerServicePerimetersRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{parent}}/servicePerimeters")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("AccessContextManagerServicePerimeters %q", d.Id()))
	}

	if err := d.Set("service_perimeters", flattenAccessContextManagerServicePerimetersServicePerimeters(res["servicePerimeters"], d, config)); err != nil {
		return fmt.Errorf("Error reading ServicePerimeters: %s", err)
	}

	return nil
}

func resourceAccessContextManagerServicePerimetersUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	billingProject := ""

	obj := make(map[string]interface{})
	servicePerimetersProp, err := expandAccessContextManagerServicePerimetersServicePerimeters(d.Get("service_perimeters"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service_perimeters"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, servicePerimetersProp)) {
		obj["servicePerimeters"] = servicePerimetersProp
	}
	parentProp, err := expandAccessContextManagerServicePerimetersParent(d.Get("parent"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{parent}}/servicePerimeters:replaceAll")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ServicePerimeters %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating ServicePerimeters %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ServicePerimeters %q: %#v", d.Id(), res)
	}

	err = accessContextManagerOperationWaitTime(
		config, res, "Updating ServicePerimeters",
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceAccessContextManagerServicePerimetersRead(d, meta)
}

func resourceAccessContextManagerServicePerimetersDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	obj["servicePerimeters"] = []string{}

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{parent}}/servicePerimeters:replaceAll")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting servicePerimeters %q: %#v", d.Id(), obj)
	res, err := sendRequestWithTimeout(config, "POST", "", url, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error deleting ServicePerimeters %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished deleting ServicePerimeters %q: %#v", d.Id(), res)
	}

	err = accessContextManagerOperationWaitTime(
		config, res, "Updating ServicePerimeters",
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return nil
}

func resourceAccessContextManagerServicePerimetersImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	parts, err := getImportIdQualifiers([]string{"accessPolicies/(?P<accessPolicy>[^/]+)/(.+)"}, d, config, d.Id())
	if err != nil {
		return nil, err
	}

	d.Set("parent", fmt.Sprintf("accessPolicies/%s", parts["accessPolicy"]))
	return []*schema.ResourceData{d}, nil
}

func flattenAccessContextManagerServicePerimetersServicePerimeters(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(accesscontextmanagerServicePerimetersServicePerimetersSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"name":                      flattenAccessContextManagerServicePerimetersServicePerimetersName(original["name"], d, config),
			"title":                     flattenAccessContextManagerServicePerimetersServicePerimetersTitle(original["title"], d, config),
			"description":               flattenAccessContextManagerServicePerimetersServicePerimetersDescription(original["description"], d, config),
			"create_time":               flattenAccessContextManagerServicePerimetersServicePerimetersCreateTime(original["createTime"], d, config),
			"update_time":               flattenAccessContextManagerServicePerimetersServicePerimetersUpdateTime(original["updateTime"], d, config),
			"perimeter_type":            flattenAccessContextManagerServicePerimetersServicePerimetersPerimeterType(original["perimeterType"], d, config),
			"status":                    flattenAccessContextManagerServicePerimetersServicePerimetersStatus(original["status"], d, config),
			"spec":                      flattenAccessContextManagerServicePerimetersServicePerimetersSpec(original["spec"], d, config),
			"use_explicit_dry_run_spec": flattenAccessContextManagerServicePerimetersServicePerimetersUseExplicitDryRunSpec(original["useExplicitDryRunSpec"], d, config),
		})
	}
	return transformed
}
func flattenAccessContextManagerServicePerimetersServicePerimetersName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimetersServicePerimetersTitle(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimetersServicePerimetersDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimetersServicePerimetersCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimetersServicePerimetersUpdateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimetersServicePerimetersPerimeterType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil || isEmptyValue(reflect.ValueOf(v)) {
		return "PERIMETER_TYPE_REGULAR"
	}

	return v
}

func flattenAccessContextManagerServicePerimetersServicePerimetersStatus(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["resources"] =
		flattenAccessContextManagerServicePerimetersServicePerimetersStatusResources(original["resources"], d, config)
	transformed["access_levels"] =
		flattenAccessContextManagerServicePerimetersServicePerimetersStatusAccessLevels(original["accessLevels"], d, config)
	transformed["restricted_services"] =
		flattenAccessContextManagerServicePerimetersServicePerimetersStatusRestrictedServices(original["restrictedServices"], d, config)
	transformed["vpc_accessible_services"] =
		flattenAccessContextManagerServicePerimetersServicePerimetersStatusVPCAccessibleServices(original["vpcAccessibleServices"], d, config)
	return []interface{}{transformed}
}
func flattenAccessContextManagerServicePerimetersServicePerimetersStatusResources(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimetersServicePerimetersStatusAccessLevels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimetersServicePerimetersStatusRestrictedServices(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenAccessContextManagerServicePerimetersServicePerimetersStatusVPCAccessibleServices(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enable_restriction"] =
		flattenAccessContextManagerServicePerimetersServicePerimetersStatusVPCAccessibleServicesEnableRestriction(original["enableRestriction"], d, config)
	transformed["allowed_services"] =
		flattenAccessContextManagerServicePerimetersServicePerimetersStatusVPCAccessibleServicesAllowedServices(original["allowedServices"], d, config)
	return []interface{}{transformed}
}
func flattenAccessContextManagerServicePerimetersServicePerimetersStatusVPCAccessibleServicesEnableRestriction(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimetersServicePerimetersStatusVPCAccessibleServicesAllowedServices(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenAccessContextManagerServicePerimetersServicePerimetersSpec(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["resources"] =
		flattenAccessContextManagerServicePerimetersServicePerimetersSpecResources(original["resources"], d, config)
	transformed["access_levels"] =
		flattenAccessContextManagerServicePerimetersServicePerimetersSpecAccessLevels(original["accessLevels"], d, config)
	transformed["restricted_services"] =
		flattenAccessContextManagerServicePerimetersServicePerimetersSpecRestrictedServices(original["restrictedServices"], d, config)
	transformed["vpc_accessible_services"] =
		flattenAccessContextManagerServicePerimetersServicePerimetersSpecVPCAccessibleServices(original["vpcAccessibleServices"], d, config)
	return []interface{}{transformed}
}
func flattenAccessContextManagerServicePerimetersServicePerimetersSpecResources(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimetersServicePerimetersSpecAccessLevels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimetersServicePerimetersSpecRestrictedServices(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimetersServicePerimetersSpecVPCAccessibleServices(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enable_restriction"] =
		flattenAccessContextManagerServicePerimetersServicePerimetersSpecVPCAccessibleServicesEnableRestriction(original["enableRestriction"], d, config)
	transformed["allowed_services"] =
		flattenAccessContextManagerServicePerimetersServicePerimetersSpecVPCAccessibleServicesAllowedServices(original["allowedServices"], d, config)
	return []interface{}{transformed}
}
func flattenAccessContextManagerServicePerimetersServicePerimetersSpecVPCAccessibleServicesEnableRestriction(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimetersServicePerimetersSpecVPCAccessibleServicesAllowedServices(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAccessContextManagerServicePerimetersServicePerimetersUseExplicitDryRunSpec(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandAccessContextManagerServicePerimetersServicePerimeters(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandAccessContextManagerServicePerimetersServicePerimetersName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedTitle, err := expandAccessContextManagerServicePerimetersServicePerimetersTitle(original["title"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !isEmptyValue(val) {
			transformed["title"] = transformedTitle
		}

		transformedDescription, err := expandAccessContextManagerServicePerimetersServicePerimetersDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !isEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedCreateTime, err := expandAccessContextManagerServicePerimetersServicePerimetersCreateTime(original["create_time"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCreateTime); val.IsValid() && !isEmptyValue(val) {
			transformed["createTime"] = transformedCreateTime
		}

		transformedUpdateTime, err := expandAccessContextManagerServicePerimetersServicePerimetersUpdateTime(original["update_time"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUpdateTime); val.IsValid() && !isEmptyValue(val) {
			transformed["updateTime"] = transformedUpdateTime
		}

		transformedPerimeterType, err := expandAccessContextManagerServicePerimetersServicePerimetersPerimeterType(original["perimeter_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPerimeterType); val.IsValid() && !isEmptyValue(val) {
			transformed["perimeterType"] = transformedPerimeterType
		}

		transformedStatus, err := expandAccessContextManagerServicePerimetersServicePerimetersStatus(original["status"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedStatus); val.IsValid() && !isEmptyValue(val) {
			transformed["status"] = transformedStatus
		}

		transformedSpec, err := expandAccessContextManagerServicePerimetersServicePerimetersSpec(original["spec"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSpec); val.IsValid() && !isEmptyValue(val) {
			transformed["spec"] = transformedSpec
		}

		transformedUseExplicitDryRunSpec, err := expandAccessContextManagerServicePerimetersServicePerimetersUseExplicitDryRunSpec(original["use_explicit_dry_run_spec"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUseExplicitDryRunSpec); val.IsValid() && !isEmptyValue(val) {
			transformed["useExplicitDryRunSpec"] = transformedUseExplicitDryRunSpec
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersTitle(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersCreateTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersUpdateTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersPerimeterType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersStatus(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResources, err := expandAccessContextManagerServicePerimetersServicePerimetersStatusResources(original["resources"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResources); val.IsValid() && !isEmptyValue(val) {
		transformed["resources"] = transformedResources
	}

	transformedAccessLevels, err := expandAccessContextManagerServicePerimetersServicePerimetersStatusAccessLevels(original["access_levels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAccessLevels); val.IsValid() && !isEmptyValue(val) {
		transformed["accessLevels"] = transformedAccessLevels
	}

	transformedRestrictedServices, err := expandAccessContextManagerServicePerimetersServicePerimetersStatusRestrictedServices(original["restricted_services"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRestrictedServices); val.IsValid() && !isEmptyValue(val) {
		transformed["restrictedServices"] = transformedRestrictedServices
	}

	transformedVPCAccessibleServices, err := expandAccessContextManagerServicePerimetersServicePerimetersStatusVPCAccessibleServices(original["vpc_accessible_services"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVPCAccessibleServices); val.IsValid() && !isEmptyValue(val) {
		transformed["vpcAccessibleServices"] = transformedVPCAccessibleServices
	}

	return transformed, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersStatusResources(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersStatusAccessLevels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersStatusRestrictedServices(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersStatusVPCAccessibleServices(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnableRestriction, err := expandAccessContextManagerServicePerimetersServicePerimetersStatusVPCAccessibleServicesEnableRestriction(original["enable_restriction"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableRestriction); val.IsValid() && !isEmptyValue(val) {
		transformed["enableRestriction"] = transformedEnableRestriction
	}

	transformedAllowedServices, err := expandAccessContextManagerServicePerimetersServicePerimetersStatusVPCAccessibleServicesAllowedServices(original["allowed_services"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedServices); val.IsValid() && !isEmptyValue(val) {
		transformed["allowedServices"] = transformedAllowedServices
	}

	return transformed, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersStatusVPCAccessibleServicesEnableRestriction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersStatusVPCAccessibleServicesAllowedServices(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersSpec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResources, err := expandAccessContextManagerServicePerimetersServicePerimetersSpecResources(original["resources"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResources); val.IsValid() && !isEmptyValue(val) {
		transformed["resources"] = transformedResources
	}

	transformedAccessLevels, err := expandAccessContextManagerServicePerimetersServicePerimetersSpecAccessLevels(original["access_levels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAccessLevels); val.IsValid() && !isEmptyValue(val) {
		transformed["accessLevels"] = transformedAccessLevels
	}

	transformedRestrictedServices, err := expandAccessContextManagerServicePerimetersServicePerimetersSpecRestrictedServices(original["restricted_services"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRestrictedServices); val.IsValid() && !isEmptyValue(val) {
		transformed["restrictedServices"] = transformedRestrictedServices
	}

	transformedVPCAccessibleServices, err := expandAccessContextManagerServicePerimetersServicePerimetersSpecVPCAccessibleServices(original["vpc_accessible_services"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVPCAccessibleServices); val.IsValid() && !isEmptyValue(val) {
		transformed["vpcAccessibleServices"] = transformedVPCAccessibleServices
	}

	return transformed, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersSpecResources(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersSpecAccessLevels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersSpecRestrictedServices(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersSpecVPCAccessibleServices(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnableRestriction, err := expandAccessContextManagerServicePerimetersServicePerimetersSpecVPCAccessibleServicesEnableRestriction(original["enable_restriction"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableRestriction); val.IsValid() && !isEmptyValue(val) {
		transformed["enableRestriction"] = transformedEnableRestriction
	}

	transformedAllowedServices, err := expandAccessContextManagerServicePerimetersServicePerimetersSpecVPCAccessibleServicesAllowedServices(original["allowed_services"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedServices); val.IsValid() && !isEmptyValue(val) {
		transformed["allowedServices"] = transformedAllowedServices
	}

	return transformed, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersSpecVPCAccessibleServicesEnableRestriction(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersSpecVPCAccessibleServicesAllowedServices(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimetersServicePerimetersUseExplicitDryRunSpec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerServicePerimetersParent(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
