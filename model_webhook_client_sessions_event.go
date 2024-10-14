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

type WebhookClientSessionsEvent struct {
	// mac address of the AP the client roamed or disconnected from
	Ap string `json:"ap"`
	// user-friendly name of the AP the client roamed or disconnected from.
	ApName string `json:"ap_name"`
	// 5GHz or 2.4GHz band
	Band string `json:"band"`
	Bssid string `json:"bssid"`
	// device family E.g. “Mac”, “iPhone”, “Apple watch”
	ClientFamily string `json:"client_family"`
	// device manufacturer E.g. “Apple”
	ClientManufacture string `json:"client_manufacture"`
	// device model E.g. “8+”, “XS”
	ClientModel string `json:"client_model"`
	// device operating system E.g. “Mojave”, “Windows 10”, “Linux”
	ClientOs string `json:"client_os"`
	// time when the user connects
	Connect int32 `json:"connect"`
	// floating point connect timestamp with millisecond precision
	ConnectFloat float64 `json:"connect_float"`
	// time when the user disconnects
	Disconnect int32 `json:"disconnect"`
	// floating point disconnect timestamp with millisecond precision
	DisconnectFloat float64 `json:"disconnect_float"`
	// the duration of the roamed or complete session indicated by termination_reason field.
	Duration int32 `json:"duration"`
	// the client’s mac
	Mac string `json:"mac"`
	// the AP the client has roamed to.
	NextAp string `json:"next_ap"`
	OrgId string `json:"org_id"`
	// latest average RSSI before the user disconnects
	Rssi float64 `json:"rssi"`
	SiteId string `json:"site_id"`
	SiteName string `json:"site_name"`
	Ssid string `json:"ssid"`
	// 1 disassociate - when the client disassociates. 2 inactive - when the client is timeout. 3 roamed - when the client is roamed between APs
	TerminationReason int32 `json:"termination_reason"`
	Timestamp float64 `json:"timestamp"`
	// schema version of this message
	Version float64 `json:"version"`
	WlanId string `json:"wlan_id"`
}
