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

// client events
type EventsClient struct {
	Ap string `json:"ap,omitempty"`
	Band *AllOfeventsClientBand `json:"band"`
	Bssid string `json:"bssid,omitempty"`
	Channel int32 `json:"channel,omitempty"`
	Proto *AllOfeventsClientProto `json:"proto"`
	Ssid string `json:"ssid,omitempty"`
	Text string `json:"text,omitempty"`
	Timestamp float64 `json:"timestamp"`
	// event type, e.g. MARVIS_EVENT_CLIENT_FBT_FAILURE
	Type_ string `json:"type,omitempty"`
	// for assoc/disassoc events
	TypeCode int32 `json:"type_code,omitempty"`
	WlanId string `json:"wlan_id,omitempty"`
}
