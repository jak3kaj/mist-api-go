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

type MxedgeTuntermIgmpSnoopingQuerier struct {
	// querier’s query response interval, in tenths-of-seconds
	MaxResponseTime int32 `json:"max_response_time,omitempty"`
	// the MTU we use (needed when forming large IGMPv3 Reports)
	Mtu int32 `json:"mtu,omitempty"`
	// querier’s query interval, in seconds
	QueryInterval int32 `json:"query_interval,omitempty"`
	// querier’s robustness
	Robustness int32 `json:"robustness,omitempty"`
	// querier’s maximum protocol version
	Version int32 `json:"version,omitempty"`
}
