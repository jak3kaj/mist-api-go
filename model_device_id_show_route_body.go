/*
 * Mist API
 *
 * > Version: **2409.1.9** > > Date: **September 27, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates] (https://www.juniper.net/documentation/us/en/software/mist/api/http/getting-started/how-to-get-started)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 
 *
 * API version: 2409.1.9
 * Contact: tmunzer@juniper.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DeviceIdShowRouteBody struct {
	// duration in sec for which refresh is enabled. Should be set only if interval is configured to non-zero value.
	Duration int32 `json:"duration,omitempty"`
	// rate at which output will refresh
	Interval int32 `json:"interval,omitempty"`
	// IP of the neighbor
	Neighbor string `json:"neighbor,omitempty"`
	Node *HaClusterNode `json:"node,omitempty"`
	// route prefix
	Prefix string `json:"prefix,omitempty"`
	Protocol *any `json:"protocol,omitempty"`
	// if specified, dump bot received and advertised, if not specified, both will be shown   * for SSR, show bgp neighbors 10.250.18.202 received-routes/advertised-routes   * for SRX and Switches, show route receive_protocol/advertise_protocol bgp 192.168.255.12'
	Route string `json:"route,omitempty"`
	// VRF name
	Vrf string `json:"vrf,omitempty"`
}
