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

// Unconnected clients statistics
type StatsUnconnectedClient struct {
	// mac address of the AP that heard the client
	ApMac string `json:"ap_mac"`
	// last seen timestamp
	LastSeen float64 `json:"last_seen"`
	// mac address of the (unconnected) client
	Mac string `json:"mac"`
	// device manufacture, through fingerprinting or OUI
	Manufacture string `json:"manufacture"`
	// map_id of the client (if known), or null
	MapId string `json:"map_id,omitempty"`
	// client RSSI observered by the AP that heard the client (in dBm)
	Rssi int32 `json:"rssi"`
	// x (in pixels) of user location on the map (if known)
	X float64 `json:"x,omitempty"`
	// y (in pixels) of user location on the map (if known)
	Y float64 `json:"y"`
}
