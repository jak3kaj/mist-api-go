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

type AllOfgatewayOobIpConfigNode1 struct {
	// if `type`==`static`
	Gateway string `json:"gateway,omitempty"`
	Ip string `json:"ip,omitempty"`
	// used only if `subnet` is not specified in `networks`
	Netmask string `json:"netmask,omitempty"`
	Type_ *any `json:"type,omitempty"`
	// if supported on the platform. If enabled, DNS will be using this routing-instance, too
	UseMgmtVrf bool `json:"use_mgmt_vrf,omitempty"`
	// whether to use `mgmt_junos` for host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired
	UseMgmtVrfForHostOut bool `json:"use_mgmt_vrf_for_host_out,omitempty"`
	VlanId string `json:"vlan_id,omitempty"`
}
