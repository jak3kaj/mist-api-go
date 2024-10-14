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

// Switch storm control Only if `mode`!=`dynamic`
type SwitchPortUsageStormControl struct {
	// whether to disable storm control on broadcast traffic
	NoBroadcast bool `json:"no_broadcast,omitempty"`
	// whether to disable storm control on multicast traffic
	NoMulticast bool `json:"no_multicast,omitempty"`
	// whether to disable storm control on registered multicast traffic
	NoRegisteredMulticast bool `json:"no_registered_multicast,omitempty"`
	// whether to disable storm control on unknown unicast traffic
	NoUnknownUnicast bool `json:"no_unknown_unicast,omitempty"`
	// bandwidth-percentage, configures the storm control level as a percentage of the available bandwidth
	Percentage int32 `json:"percentage,omitempty"`
}
