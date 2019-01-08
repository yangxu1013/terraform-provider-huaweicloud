// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file at
//     https://www.github.com/huaweicloud/magic-modules
//
// ----------------------------------------------------------------------------

package huaweicloud

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/huaweicloud/golangsdk"
)

func resourceDwsCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceDwsClusterCreate,
		Read:   resourceDwsClusterRead,
		Delete: resourceDwsClusterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"network_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"node_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"number_of_node": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"security_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"user_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"user_pwd": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"availability_zone": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},

			"port": {
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},

			"public_ip": {
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"eip_id": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
							ForceNew: true,
						},
						"public_bind_type": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},

			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"endpoints": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connect_info": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"jdbc_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"public_endpoints": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"jdbc_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_connect_info": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"recent_event": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"sub_status": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"task_status": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"updated": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDwsClusterCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sdkClient(GetRegion(d, config), "dws")
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	opts := make(map[string]interface{})
	userPwdProp := d.Get("user_pwd")
	e, err := isEmptyValue(reflect.ValueOf(userPwdProp))
	if err != nil {
		return err
	}
	if !e {
		opts["user_pwd"] = userPwdProp
	}

	availabilityZoneProp := d.Get("availability_zone")
	e, err = isEmptyValue(reflect.ValueOf(availabilityZoneProp))
	if err != nil {
		return err
	}
	if !e {
		opts["availability_zone"] = availabilityZoneProp
	}

	nameProp := d.Get("name")
	e, err = isEmptyValue(reflect.ValueOf(nameProp))
	if err != nil {
		return err
	}
	if !e {
		opts["name"] = nameProp
	}

	networkIDProp := d.Get("network_id")
	e, err = isEmptyValue(reflect.ValueOf(networkIDProp))
	if err != nil {
		return err
	}
	if !e {
		opts["subnet_id"] = networkIDProp
	}

	nodeTypeProp := d.Get("node_type")
	e, err = isEmptyValue(reflect.ValueOf(nodeTypeProp))
	if err != nil {
		return err
	}
	if !e {
		opts["node_type"] = nodeTypeProp
	}

	numberOFNodeProp := d.Get("number_of_node")
	e, err = isEmptyValue(reflect.ValueOf(numberOFNodeProp))
	if err != nil {
		return err
	}
	if !e {
		opts["number_of_node"] = numberOFNodeProp
	}

	portProp := d.Get("port")
	e, err = isEmptyValue(reflect.ValueOf(portProp))
	if err != nil {
		return err
	}
	if !e {
		opts["port"] = portProp
	}

	publicIPProp, err := expandDwsClusterPublicIP(d.Get("public_ip"))
	if err != nil {
		return err
	}
	e, err = isEmptyValue(reflect.ValueOf(publicIPProp))
	if err != nil {
		return err
	}
	if !e {
		opts["public_ip"] = publicIPProp
	}

	securityGroupIDProp := d.Get("security_group_id")
	e, err = isEmptyValue(reflect.ValueOf(securityGroupIDProp))
	if err != nil {
		return err
	}
	if !e {
		opts["security_group_id"] = securityGroupIDProp
	}

	userNameProp := d.Get("user_name")
	e, err = isEmptyValue(reflect.ValueOf(userNameProp))
	if err != nil {
		return err
	}
	if !e {
		opts["user_name"] = userNameProp
	}

	vpcIDProp := d.Get("vpc_id")
	e, err = isEmptyValue(reflect.ValueOf(vpcIDProp))
	if err != nil {
		return err
	}
	if !e {
		opts["vpc_id"] = vpcIDProp
	}

	url, err := replaceVars(d, "clusters", nil)
	if err != nil {
		return err
	}
	url = client.ServiceURL(url)

	log.Printf("[DEBUG] Creating new Cluster: %#v", opts)
	r := golangsdk.Result{}
	_, r.Err = client.Post(
		url,
		&map[string]interface{}{"cluster": opts},
		&r.Body,
		&golangsdk.RequestOpts{OkCodes: successHTTPCodes})
	if r.Err != nil {
		return fmt.Errorf("Error creating Cluster: %s", r.Err)
	}

	pathParameters := map[string][]string{
		"id": {"cluster", "id"},
	}
	var data = make(map[string]string)
	for key, path := range pathParameters {
		value, err := navigateMap(r.Body, path)
		if err != nil {
			return fmt.Errorf("Error retrieving async operation path parameter: %s", err)
		}
		data[key] = value.(string)
	}
	url, err = replaceVars(d, "clusters/{id}", data)
	if err != nil {
		return err
	}
	url = client.ServiceURL(url)

	obj, err := waitToFinish(
		[]string{"AVAILABLE"},
		[]string{"CREATING", "Pending"},
		d.Timeout(schema.TimeoutCreate),
		1*time.Second,
		func() (interface{}, string, error) {
			r := golangsdk.Result{}
			_, r.Err = client.Get(
				url, &r.Body,
				&golangsdk.RequestOpts{MoreHeaders: map[string]string{"Content-Type": "application/json"}})
			if r.Err != nil {
				return nil, "", nil
			}

			code, err := navigateMap(r.Body, []string{"failed_reasons", "error_code"})
			if err == nil {
				if v, err := convertToInt(code); err == nil {
					msg, err := navigateMap(r.Body, []string{"failed_reasons", "error_msg"})
					if err != nil {
						return nil, "", fmt.Errorf("async operation failed: %v", msg)
					}
					return nil, "", fmt.Errorf("async operation failed: error code = %v", v)
				}
			}

			status, err := navigateMap(r.Body, []string{"cluster", "status"})
			if err != nil {
				return nil, "", nil
			}
			return r.Body, status.(string), nil
		},
	)

	if err != nil {
		return err
	}
	id, err := navigateMap(obj, []string{"cluster", "id"})
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id.(string))

	return resourceDwsClusterRead(d, meta)
}

func resourceDwsClusterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sdkClient(GetRegion(d, config), "dws")
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	url, err := replaceVars(d, "clusters/{id}", nil)
	if err != nil {
		return err
	}
	url = client.ServiceURL(url)

	r := golangsdk.Result{}
	_, r.Err = client.Get(
		url, &r.Body,
		&golangsdk.RequestOpts{MoreHeaders: map[string]string{"Content-Type": "application/json"}})
	if r.Err != nil {
		return fmt.Errorf("Error reading %s: %s", fmt.Sprintf("DwsCluster %q", d.Id()), r.Err)
	}
	v, err := navigateMap(r.Body, []string{"cluster"})
	if err != nil {
		return fmt.Errorf("Error reading %s: the result does not contain cluster", fmt.Sprintf("DwsCluster %q", d.Id()))
	}
	res := v.(map[string]interface{})

	if v, ok := res["availability_zone"]; ok {
		if err := d.Set("availability_zone", v); err != nil {
			return fmt.Errorf("Error reading Cluster:availability_zone, err: %s", err)
		}
	}

	if v, ok := res["created"]; ok {
		if err := d.Set("created", v); err != nil {
			return fmt.Errorf("Error reading Cluster:created, err: %s", err)
		}
	}

	if v, ok := res["endpoints"]; ok {
		endpointsProp, err := flattenDwsClusterEndpoints(v)
		if err != nil {
			return fmt.Errorf("Error reading Cluster:endpoints, err: %s", err)
		}
		if err := d.Set("endpoints", endpointsProp); err != nil {
			return fmt.Errorf("Error reading Cluster:endpoints, err: %s", err)
		}
	}

	if v, ok := res["id"]; ok {
		if err := d.Set("id", v); err != nil {
			return fmt.Errorf("Error reading Cluster:id, err: %s", err)
		}
	}

	if v, ok := res["port"]; ok {
		if err := d.Set("port", v); err != nil {
			return fmt.Errorf("Error reading Cluster:port, err: %s", err)
		}
	}

	if v, ok := res["public_endpoints"]; ok {
		publicEndpointsProp, err := flattenDwsClusterPublicEndpoints(v)
		if err != nil {
			return fmt.Errorf("Error reading Cluster:public_endpoints, err: %s", err)
		}
		if err := d.Set("public_endpoints", publicEndpointsProp); err != nil {
			return fmt.Errorf("Error reading Cluster:public_endpoints, err: %s", err)
		}
	}

	if v, ok := res["public_ip"]; ok {
		publicIPProp, err := flattenDwsClusterPublicIP(v)
		if err != nil {
			return fmt.Errorf("Error reading Cluster:public_ip, err: %s", err)
		}
		if err := d.Set("public_ip", publicIPProp); err != nil {
			return fmt.Errorf("Error reading Cluster:public_ip, err: %s", err)
		}
	}

	if v, ok := res["recent_event"]; ok {
		if err := d.Set("recent_event", v); err != nil {
			return fmt.Errorf("Error reading Cluster:recent_event, err: %s", err)
		}
	}

	if v, ok := res["status"]; ok {
		if err := d.Set("status", v); err != nil {
			return fmt.Errorf("Error reading Cluster:status, err: %s", err)
		}
	}

	if v, ok := res["sub_status"]; ok {
		if err := d.Set("sub_status", v); err != nil {
			return fmt.Errorf("Error reading Cluster:sub_status, err: %s", err)
		}
	}

	if v, ok := res["task_status"]; ok {
		if err := d.Set("task_status", v); err != nil {
			return fmt.Errorf("Error reading Cluster:task_status, err: %s", err)
		}
	}

	if v, ok := res["updated"]; ok {
		if err := d.Set("updated", v); err != nil {
			return fmt.Errorf("Error reading Cluster:updated, err: %s", err)
		}
	}

	if v, ok := res["version"]; ok {
		if err := d.Set("version", v); err != nil {
			return fmt.Errorf("Error reading Cluster:version, err: %s", err)
		}
	}

	return nil
}

func resourceDwsClusterDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	client, err := config.sdkClient(GetRegion(d, config), "dws")
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	url, err := replaceVars(d, "clusters/{id}", nil)
	if err != nil {
		return err
	}
	url = client.ServiceURL(url)

	log.Printf("[DEBUG] Deleting Cluster %q", d.Id())
	r := golangsdk.Result{}
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		OkCodes:      successHTTPCodes,
		JSONResponse: &r.Body,
		MoreHeaders:  map[string]string{"Content-Type": "application/json"},
		JSONBody: map[string]interface{}{
			"keep_last_manual_snapshot": 0,
		},
	})
	if r.Err != nil {
		return fmt.Errorf("Error deleting Cluster %q: %s", d.Id(), r.Err)
	}

	_, err = waitToFinish(
		[]string{"Done"}, []string{"Pending"},
		d.Timeout(schema.TimeoutDelete),
		1*time.Second,
		func() (interface{}, string, error) {
			resp, err := client.Get(
				url, nil,
				&golangsdk.RequestOpts{MoreHeaders: map[string]string{"Content-Type": "application/json"}})
			if err != nil {
				if _, ok := err.(golangsdk.ErrDefault404); ok {
					return resp, "Done", nil
				}
				return nil, "", nil
			}
			return resp, "Pending", nil
		},
	)
	return err
}

func flattenDwsClusterEndpoints(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	l := v.([]interface{})
	result := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		if val, ok := original["connect_info"]; ok {
			transformed["connect_info"] = val
		}

		if val, ok := original["jdbc_url"]; ok {
			transformed["jdbc_url"] = val
		}

		result = append(result, transformed)
	}

	return result, nil
}

func flattenDwsClusterPublicEndpoints(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	l := v.([]interface{})
	result := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		if val, ok := original["jdbc_url"]; ok {
			transformed["jdbc_url"] = val
		}

		if val, ok := original["public_connect_info"]; ok {
			transformed["public_connect_info"] = val
		}

		result = append(result, transformed)
	}

	return result, nil
}

func flattenDwsClusterPublicIP(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})

	if val, ok := original["eip_id"]; ok {
		transformed["eip_id"] = val
	}

	if val, ok := original["public_bind_type"]; ok {
		transformed["public_bind_type"] = val
	}

	return []interface{}{transformed}, nil
}

func expandDwsClusterPublicIP(v interface{}) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	eipIDProp := original["eip_id"]
	e, err := isEmptyValue(reflect.ValueOf(eipIDProp))
	if err != nil {
		return nil, err
	}
	if !e {
		transformed["eip_id"] = eipIDProp
	}

	publicBindTypeProp := original["public_bind_type"]
	e, err = isEmptyValue(reflect.ValueOf(publicBindTypeProp))
	if err != nil {
		return nil, err
	}
	if !e {
		transformed["public_bind_type"] = publicBindTypeProp
	}

	return transformed, nil
}
