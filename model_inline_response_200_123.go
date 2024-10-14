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

type InlineResponse200123 struct {
	Band2440mhzAllowed bool `json:"band24_40mhz_allowed"`
	// Property key is the channel width
	Band24Channels map[string]Object `json:"band24_channels"`
	Band24Enabled bool `json:"band24_enabled"`
	// Property key is the channel width
	Band5Channels map[string]Object `json:"band5_channels"`
	Band5Enabled bool `json:"band5_enabled"`
	// Property key is the channel width
	Band6Channels map[string]Object `json:"band6_channels,omitempty"`
	Band6Enabled bool `json:"band6_enabled,omitempty"`
	Certified bool `json:"certified"`
	Code int32 `json:"code"`
	DfsOk bool `json:"dfs_ok"`
	Key string `json:"key"`
	Name string `json:"name"`
	Uses string `json:"uses"`
}
