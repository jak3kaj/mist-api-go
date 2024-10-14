/*
 * Mist API
 *
 * > Version: **2409.1.9** > > Date: **September 27, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates] (https://www.juniper.net/documentation/us/en/software/mist/api/http/getting-started/how-to-get-started)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 
 *
 * API version: 2409.1.9
 * Contact: tmunzer@juniper.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AllOfsiteSettingSwitch_ struct {
	AclPolicies []AclPolicy `json:"acl_policies,omitempty"`
	// ACL Tags to identify traffic source or destination. Key name is the tag name
	AclTags map[string]AclTag `json:"acl_tags,omitempty"`
	// additional CLI commands to append to the generated Junos config  **Note**: no check is done
	AdditionalConfigCmds []string `json:"additional_config_cmds,omitempty"`
	CreatedTime float64 `json:"created_time,omitempty"`
	DhcpSnooping *DhcpSnooping `json:"dhcp_snooping,omitempty"`
	// Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
	DnsServers []string `json:"dns_servers,omitempty"`
	// Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
	DnsSuffix []string `json:"dns_suffix,omitempty"`
	ExtraRoutes map[string]ExtraRoute `json:"extra_routes,omitempty"`
	// Property key is the destination CIDR (e.g. \"2a02:1234:420a:10c9::/64\")
	ExtraRoutes6 map[string]ExtraRoute6 `json:"extra_routes6,omitempty"`
	Id string `json:"id,omitempty"`
	// Org Networks that we'd like to import
	ImportOrgNetworks []string `json:"import_org_networks,omitempty"`
	MistNac *Object `json:"mist_nac,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	Name string `json:"name,omitempty"`
	// Property key is network name
	Networks map[string]SwitchNetwork `json:"networks,omitempty"`
	// list of NTP servers specific to this device. By default, those in Site Settings will be used
	NtpServers []string `json:"ntp_servers,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	// Junos OSPF areas
	OspfAreas map[string]OspfArea `json:"ospf_areas,omitempty"`
	// Property key is the port mirroring instance name port_mirroring can be added under device/site settings. It takes interface and ports as input for ingress, interface as input for egress and can take interface and port as output. A maximum 4 port mirrorings is allowed
	PortMirroring map[string]SwitchPortMirroringProperty `json:"port_mirroring,omitempty"`
	PortUsages map[string]SwitchPortUsage `json:"port_usages,omitempty"`
	RadiusConfig *Object `json:"radius_config,omitempty"`
	RemoteSyslog *RemoteSyslog `json:"remote_syslog,omitempty"`
	// by default, when we configure a device, we only clean up config we generates. Remove existing configs if enabled
	RemoveExistingConfigs bool `json:"remove_existing_configs,omitempty"`
	SnmpConfig *SnmpConfig `json:"snmp_config,omitempty"`
	SwitchMatching *Object `json:"switch_matching,omitempty"`
	SwitchMgmt *Object `json:"switch_mgmt,omitempty"`
	VrfConfig *VrfConfig `json:"vrf_config,omitempty"`
	// Property key is the network name
	VrfInstances map[string]SwitchVrfInstance `json:"vrf_instances,omitempty"`
}
