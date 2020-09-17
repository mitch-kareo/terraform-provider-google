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
)

func resourceAppEngineApplicationUrlDispatchRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppEngineApplicationUrlDispatchRulesCreate,
		Read:   resourceAppEngineApplicationUrlDispatchRulesRead,
		Update: resourceAppEngineApplicationUrlDispatchRulesUpdate,
		Delete: resourceAppEngineApplicationUrlDispatchRulesDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAppEngineApplicationUrlDispatchRulesImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"dispatch_rules": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `Rules to match an HTTP request and dispatch that request to a service.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"path": {
							Type:     schema.TypeString,
							Required: true,
							Description: `Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
The sum of the lengths of the domain and path may not exceed 100 characters.`,
						},
						"service": {
							Type:     schema.TypeString,
							Required: true,
							Description: `Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
The sum of the lengths of the domain and path may not exceed 100 characters.`,
						},
						"domain": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `Domain name to match against. The wildcard "*" is supported if specified before a period: "*.".
Defaults to matching all domains: "*".`,
							Default: "*",
						},
					},
				},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAppEngineApplicationUrlDispatchRulesCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := make(map[string]interface{})
	dispatchRulesProp, err := expandAppEngineApplicationUrlDispatchRulesDispatchRules(d.Get("dispatch_rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dispatch_rules"); !isEmptyValue(reflect.ValueOf(dispatchRulesProp)) && (ok || !reflect.DeepEqual(v, dispatchRulesProp)) {
		obj["dispatchRules"] = dispatchRulesProp
	}

	lockName, err := replaceVars(d, config, "apps/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}?updateMask=dispatch_rules")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ApplicationUrlDispatchRules: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, obj, d.Timeout(schema.TimeoutCreate), isAppEngineRetryableError)
	if err != nil {
		return fmt.Errorf("Error creating ApplicationUrlDispatchRules: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{project}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = appEngineOperationWaitTime(
		config, res, project, "Creating ApplicationUrlDispatchRules",
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create ApplicationUrlDispatchRules: %s", err)
	}

	log.Printf("[DEBUG] Finished creating ApplicationUrlDispatchRules %q: %#v", d.Id(), res)

	return resourceAppEngineApplicationUrlDispatchRulesRead(d, meta)
}

func resourceAppEngineApplicationUrlDispatchRulesRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	url, err := replaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, nil, isAppEngineRetryableError)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("AppEngineApplicationUrlDispatchRules %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ApplicationUrlDispatchRules: %s", err)
	}

	if err := d.Set("dispatch_rules", flattenAppEngineApplicationUrlDispatchRulesDispatchRules(res["dispatchRules"], d, config)); err != nil {
		return fmt.Errorf("Error reading ApplicationUrlDispatchRules: %s", err)
	}

	return nil
}

func resourceAppEngineApplicationUrlDispatchRulesUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	obj := make(map[string]interface{})
	dispatchRulesProp, err := expandAppEngineApplicationUrlDispatchRulesDispatchRules(d.Get("dispatch_rules"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dispatch_rules"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, dispatchRulesProp)) {
		obj["dispatchRules"] = dispatchRulesProp
	}

	lockName, err := replaceVars(d, config, "apps/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}?updateMask=dispatch_rules")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ApplicationUrlDispatchRules %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, obj, d.Timeout(schema.TimeoutUpdate), isAppEngineRetryableError)

	if err != nil {
		return fmt.Errorf("Error updating ApplicationUrlDispatchRules %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ApplicationUrlDispatchRules %q: %#v", d.Id(), res)
	}

	err = appEngineOperationWaitTime(
		config, res, project, "Updating ApplicationUrlDispatchRules",
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceAppEngineApplicationUrlDispatchRulesRead(d, meta)
}

func resourceAppEngineApplicationUrlDispatchRulesDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	lockName, err := replaceVars(d, config, "apps/{{project}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{AppEngineBasePath}}apps/{{project}}?updateMask=dispatch_rules")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting ApplicationUrlDispatchRules %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, obj, d.Timeout(schema.TimeoutDelete), isAppEngineRetryableError)
	if err != nil {
		return handleNotFoundError(err, d, "ApplicationUrlDispatchRules")
	}

	err = appEngineOperationWaitTime(
		config, res, project, "Deleting ApplicationUrlDispatchRules",
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting ApplicationUrlDispatchRules %q: %#v", d.Id(), res)
	return nil
}

func resourceAppEngineApplicationUrlDispatchRulesImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"(?P<project>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{project}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAppEngineApplicationUrlDispatchRulesDispatchRules(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"domain":  flattenAppEngineApplicationUrlDispatchRulesDispatchRulesDomain(original["domain"], d, config),
			"path":    flattenAppEngineApplicationUrlDispatchRulesDispatchRulesPath(original["path"], d, config),
			"service": flattenAppEngineApplicationUrlDispatchRulesDispatchRulesService(original["service"], d, config),
		})
	}
	return transformed
}
func flattenAppEngineApplicationUrlDispatchRulesDispatchRulesDomain(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAppEngineApplicationUrlDispatchRulesDispatchRulesPath(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenAppEngineApplicationUrlDispatchRulesDispatchRulesService(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandAppEngineApplicationUrlDispatchRulesDispatchRules(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDomain, err := expandAppEngineApplicationUrlDispatchRulesDispatchRulesDomain(original["domain"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDomain); val.IsValid() && !isEmptyValue(val) {
			transformed["domain"] = transformedDomain
		}

		transformedPath, err := expandAppEngineApplicationUrlDispatchRulesDispatchRulesPath(original["path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !isEmptyValue(val) {
			transformed["path"] = transformedPath
		}

		transformedService, err := expandAppEngineApplicationUrlDispatchRulesDispatchRulesService(original["service"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedService); val.IsValid() && !isEmptyValue(val) {
			transformed["service"] = transformedService
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAppEngineApplicationUrlDispatchRulesDispatchRulesDomain(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineApplicationUrlDispatchRulesDispatchRulesPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineApplicationUrlDispatchRulesDispatchRulesService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
