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

type ResponseConfigHistorySearchItem struct {
	Channel24 int32 `json:"channel_24"`
	Channel5 int32 `json:"channel_5"`
	RadioMacs []string `json:"radio_macs,omitempty"`
	Radios []ResponseConfigHistorySearchItemRadio `json:"radios,omitempty"`
	SecpolicyViolated bool `json:"secpolicy_violated"`
	Ssids []string `json:"ssids,omitempty"`
	Ssids24 []string `json:"ssids_24,omitempty"`
	Ssids5 []string `json:"ssids_5,omitempty"`
	Timestamp float64 `json:"timestamp"`
	Version string `json:"version"`
	Wlans []ResponseConfigHistorySearchItemWlan `json:"wlans,omitempty"`
}
