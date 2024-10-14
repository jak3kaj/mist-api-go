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

// IP AP settings
type ApIpConfig struct {
	// if `type`==`static`
	Dns []string `json:"dns,omitempty"`
	// required if `type`==`static`
	DnsSuffix []string `json:"dns_suffix,omitempty"`
	// required if `type`==`static`
	Gateway string `json:"gateway,omitempty"`
	Gateway6 string `json:"gateway6,omitempty"`
	// required if `type`==`static`
	Ip string `json:"ip,omitempty"`
	Ip6 string `json:"ip6,omitempty"`
	Mtu int32 `json:"mtu,omitempty"`
	// required if `type`==`static`
	Netmask string `json:"netmask,omitempty"`
	Netmask6 string `json:"netmask6,omitempty"`
	Type_ *AllOfapIpConfigType_ `json:"type,omitempty"`
	Type6 *AllOfapIpConfigType6 `json:"type6,omitempty"`
	// management vlan id, default is 1 (untagged)
	VlanId int32 `json:"vlan_id,omitempty"`
}
