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

// bonjour gateway wlan settings
type WlanBonjour struct {
	// comma sperated list of additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses
	AdditionalVlanIds string `json:"additional_vlan_ids"`
	// whether to enable bonjour for this WLAN. Once enabled, limit_bcast is assumed true, allow_mdns is assumed false
	Enabled bool `json:"enabled,omitempty"`
	// what services are allowed.  Property key is the service name
	Services map[string]WlanBonjourServiceProperties `json:"services"`
}
