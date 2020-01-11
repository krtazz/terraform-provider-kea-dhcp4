package main

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func hostLease() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "lease_name in lowercase",
				ForceNew:     true,
				ValidateFunc: validateName,
			},
			"mac_address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "mac address in lower-case",
			},
			"ip_address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ip address to assign for mac",
			},
		},
		Create: resourceCreateLease,
		Read:   resourceReadLease,
		//Update: resourceUpdateLease,
		Delete: resourceDeleteLease,
		//Exists: resourceExistsLease,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func validateName(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}
	whiteSpace := regexp.MustCompile(`\s+`)
	if whiteSpace.Match([]byte(value)) {
		errs = append(errs, fmt.Errorf("name cannot contain whitespace. Got %s", value))
		return warns, errs
	}
	return warns, errs
}

func resourceCreateLease(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*Client)

	lease := Reservations{
		Hostname:  d.Get("name").(string),
		Hwaddress: d.Get("mac_address").(string),
		Ipaddress: d.Get("ip_address").(string),
	}

	err := apiClient.NewLease(lease)

	if err != nil {
		return err
	}
	d.SetId(lease.Hostname)
	return nil
}

/*func resourceUpdateLease(d *schema.ResourceData, m interface{}) error {
	return nil
}*/
func resourceReadLease(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceDeleteLease(d *schema.ResourceData, m interface{}) error {
	return nil
}

/*func resourceExistsLease(d *schema.ResourceData, m interface{}) error {
	return nil
}*/
