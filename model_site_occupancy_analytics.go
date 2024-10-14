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

// Occupancy Analytics settings
type SiteOccupancyAnalytics struct {
	// indicate whether named BLE assets should be included in the zone occupancy calculation
	AssetsEnabled bool `json:"assets_enabled,omitempty"`
	// indicate whether connected WiFi clients should be included in the zone occupancy calculation
	ClientsEnabled bool `json:"clients_enabled,omitempty"`
	// minimum duration
	MinDuration int32 `json:"min_duration,omitempty"`
	// indicate whether SDK clients should be included in the zone occupancy calculation
	SdkclientsEnabled bool `json:"sdkclients_enabled,omitempty"`
	// indicate whether unconnected WiFi clients should be included in the zone occupancy calculation
	UnconnectedClientsEnabled bool `json:"unconnected_clients_enabled,omitempty"`
}
