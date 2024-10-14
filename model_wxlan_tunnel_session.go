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

type WxlanTunnelSession struct {
	// if `use_ap_as_session_ids`==`true`, only apmac is supported right now. This is the name WLAN should use for wxtunnel_remote_id
	ApAsSessionId string `json:"ap_as_session_id,omitempty"`
	// optional, user-specified string for display purpose
	Comment string `json:"comment,omitempty"`
	EnableCookie bool `json:"enable_cookie,omitempty"`
	Ethertype *AllOfwxlanTunnelSessionEthertype `json:"ethertype,omitempty"`
	// 1-2147483647
	LocalSessionId int32 `json:"local_session_id,omitempty"`
	// optional. Enables the pseudo 802.1ad QinQ mode where the AP device drops the outer vlan tag (QinQ). This mode is useful when tunneling Mist AP’s to some aggregation routers.
	Pseudo8021adEnabled bool `json:"pseudo_802.1ad_enabled,omitempty"`
	// remote-id of the session, has to be unique in the same tunnel
	RemoteId string `json:"remote_id,omitempty"`
	// 1-2147483647
	RemoteSessionId int32 `json:"remote_session_id,omitempty"`
	// whether to use AP (last 4 bytes of MAC currently) as session ids
	UseApAsSessionIds bool `json:"use_ap_as_session_ids,omitempty"`
}
