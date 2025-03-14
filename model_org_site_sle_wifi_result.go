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

type OrgSiteSleWifiResult struct {
	ApAvailability float64 `json:"ap-availability,omitempty"`
	ApHealth float64 `json:"ap-health,omitempty"`
	Capacity float64 `json:"capacity,omitempty"`
	Coverage float64 `json:"coverage,omitempty"`
	NumAps float64 `json:"num_aps,omitempty"`
	NumClients float64 `json:"num_clients,omitempty"`
	Roaming float64 `json:"roaming,omitempty"`
	SiteId string `json:"site_id"`
	SuccessfulConnect float64 `json:"successful-connect,omitempty"`
	Throughput float64 `json:"throughput,omitempty"`
	TimeToConnect float64 `json:"time-to-connect,omitempty"`
}
