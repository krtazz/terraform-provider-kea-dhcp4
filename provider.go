package main

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider definition for kea-dhcpd
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"kea_server_address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP or FQDN of host which serves Kea Control Agent API",
			},
			"kea_server_username": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "admin",
				Description: "HTTP basic auth username (if configured)",
			},
			"kea_server_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "password",
				Description: "HTTP basic auth password (if configured)",
			},
			"kea_server_configfile": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Path to Kea DHCP4 server config (on server, required for config-write command)",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"kea-dhcp4_host_lease": hostLease(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Server:     d.Get("kea_server_address").(string),
		Username:   d.Get("kea_server_username").(string),
		Password:   d.Get("kea_server_password").(string),
		Configfile: d.Get("kea_server_configfile").(string),
	}

	log.Printf("[DEBUG] Configuring Server for kea-dhcpd:  '%s': %v", config.Server, d)
	log.Printf("[DEBUG] Configuring Username for kea-dhcpd:  '%s': %v", config.Username, d)
	log.Printf("[DEBUG] Configuring config file for kea-dhcpd:  '%s': %v", config.Configfile, d)

	client, err := config.Client()
	if err != nil {
		return nil, err
	}
	return client, nil
}
