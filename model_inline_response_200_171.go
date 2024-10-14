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

type InlineResponse200171 struct {
	Id string `json:"id"`
	// last seen timestamp
	LastSeen float64 `json:"last_seen"`
	// map_id of the sdk client (if known), or null
	MapId string `json:"map_id,omitempty"`
	// name of the sdk client (if provided)
	Name string `json:"name,omitempty"`
	NetworkConnection *any `json:"network_connection,omitempty"`
	// uuid of the sdk client
	Uuid string `json:"uuid"`
	// list of beacon_id’s of the sdk client is in and since when (if known)
	Vbeacons []SdkstatsWirelessClientVbeacon `json:"vbeacons,omitempty"`
	// x (in pixels) of user location on the map (if known)
	X float64 `json:"x,omitempty"`
	// y (in pixels) of user location on the map (if known)
	Y float64 `json:"y,omitempty"`
	// list of zone_id’s of the sdk client is in and since when (if known)
	Zones []SdkstatsWirelessClientZone `json:"zones,omitempty"`
}
