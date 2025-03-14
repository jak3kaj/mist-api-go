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

// Property key is the network name. Networks to participate in an OSPF area
type OspfAreasNetwork struct {
	// Required if `auth_type`==`md5`. Property key is the key number
	AuthKeys map[string]string `json:"auth_keys,omitempty"`
	// Required if `auth_type`==`password`, the password, max length is 8
	AuthPassword string `json:"auth_password,omitempty"`
	AuthType *AllOfospfAreasNetworkAuthType `json:"auth_type,omitempty"`
	BfdMinimumInterval int32 `json:"bfd_minimum_interval,omitempty"`
	DeadInterval int32 `json:"dead_interval,omitempty"`
	ExportPolicy string `json:"export_policy,omitempty"`
	HelloInterval int32 `json:"hello_interval,omitempty"`
	ImportPolicy string `json:"import_policy,omitempty"`
	InterfaceType *AllOfospfAreasNetworkInterfaceType `json:"interface_type,omitempty"`
	Metric int32 `json:"metric,omitempty"`
	// by default, we'll re-advertise all learned OSPF routes toward overlay
	NoReadvertiseToOverlay bool `json:"no_readvertise_to_overlay,omitempty"`
	// whether to send OSPF-Hello
	Passive bool `json:"passive,omitempty"`
}
