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

type StatsMxedgeServiceStat struct {
	// external IP from ep-terminator’s point of view. valid only for service having its own cloud connection
	ExtIp string `json:"ext_ip,omitempty"`
	// timestamp when the last stats is seen (cloud unix time, in second). valid only for service having its own stats or whole system (last among last_seen of all services)
	LastSeen float64 `json:"last_seen,omitempty"`
	// package/service installation state.
	PackageState string `json:"package_state,omitempty"`
	// package/service installation state.
	PackageVersion string `json:"package_version,omitempty"`
	// service running state.
	RunningState string `json:"running_state,omitempty"`
	// service uptime.
	Uptime int32 `json:"uptime,omitempty"`
}
