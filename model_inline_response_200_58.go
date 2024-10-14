/*
 * Mist API
 *
 * > Version: **2409.1.9** > > Date: **September 27, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates] (https://www.juniper.net/documentation/us/en/software/mist/api/http/getting-started/how-to-get-started)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 
 *
 * API version: 2409.1.9
 * Contact: tmunzer@juniper.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mist

type InlineResponse20058 struct {
	// additional CLI commands to append to the generated Junos config  **Note**: no check is done
	AdditionalConfigCmds []string `json:"additional_config_cmds,omitempty"`
	BgpConfig map[string]BgpConfig `json:"bgp_config,omitempty"`
	CreatedTime float64 `json:"created_time,omitempty"`
	DhcpdConfig *map[string]DhcpdConfigProperty `json:"dhcpd_config,omitempty"`
	DnsOverride bool `json:"dnsOverride,omitempty"`
	// Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
	DnsServers []string `json:"dns_servers,omitempty"`
	// Global dns settings. To keep compatibility, dns settings in `ip_config` and `oob_ip_config` will overwrite this setting
	DnsSuffix []string `json:"dns_suffix,omitempty"`
	// Property key is the destination CIDR (e.g. \"10.0.0.0/8\")
	ExtraRoutes map[string]GatewayExtraRoute `json:"extra_routes,omitempty"`
	// Property key is the destination CIDR (e.g. \"2a02:1234:420a:10c9::/64\")
	ExtraRoutes6 map[string]GatewayExtraRoute `json:"extra_routes6,omitempty"`
	GatewayMatching *any `json:"gateway_matching,omitempty"`
	Id string `json:"id,omitempty"`
	// Property key is the profile name
	IdpProfiles map[string]IdpProfile `json:"idp_profiles,omitempty"`
	// Property key is the network name
	IpConfigs map[string]GatewayIpConfigProperty `json:"ip_configs,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	Name string `json:"name"`
	Networks []Network `json:"networks,omitempty"`
	NtpOverride bool `json:"ntpOverride,omitempty"`
	// list of NTP servers specific to this device. By default, those in Site Settings will be used
	NtpServers []string `json:"ntp_servers,omitempty"`
	OobIpConfig *any `json:"oob_ip_config,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	// Property key is the path name
	PathPreferences map[string]GatewayPathPreferences `json:"path_preferences,omitempty"`
	// Property key is the port(s) name or range (e.g. \"ge-0/0/0-10\")
	PortConfig map[string]GatewayPortConfig `json:"port_config,omitempty"`
	// auto assigned if not set
	RouterId string `json:"router_id,omitempty"`
	// Property key is the routing policy name
	RoutingPolicies map[string]RoutingPolicy `json:"routing_policies,omitempty"`
	ServicePolicies []ServicePolicy `json:"service_policies,omitempty"`
	// Property key is the tunnel name
	TunnelConfigs map[string]TunnelConfigs `json:"tunnel_configs,omitempty"`
	TunnelProviderOptions *TunnelProviderOptions `json:"tunnel_provider_options,omitempty"`
	Type_ *any `json:"type,omitempty"`
	VrfConfig *VrfConfig `json:"vrf_config,omitempty"`
	// Property key is the network name
	VrfInstances map[string]GatewayVrfInstance `json:"vrf_instances,omitempty"`
}
