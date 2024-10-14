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

type WebhookNacAccountingEvent struct {
	// mac address of the AP the client roamed or disconnected from
	Ap string `json:"ap,omitempty"`
	// radius authentication type
	AuthType string `json:"auth_type,omitempty"`
	// it’s the MAC physical address of the access point
	Bssid string `json:"bssid,omitempty"`
	// IP Address of client
	ClientIp string `json:"client_ip,omitempty"`
	// client type E.g. “wired”, “wireless”, “vty”
	ClientType string `json:"client_type,omitempty"`
	// the client’s mac
	Mac string `json:"mac,omitempty"`
	// NAS Device vendor name E.g. “Juniper”, “Cisco”
	NasVendor string `json:"nas_vendor,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	// number of packets received
	RxPkts int32 `json:"rx_pkts,omitempty"`
	SiteId string `json:"site_id,omitempty"`
	// ESSID
	Ssid string `json:"ssid,omitempty"`
	// sampling time (in epoch)
	Timestamp float64 `json:"timestamp,omitempty"`
	// number of packets sent
	TxPkts int32 `json:"tx_pkts,omitempty"`
	// type of event. E.g. “ACCOUNTING_START”, “ACCOUNTING_UPDATE”, “ACCOUNTING_STOP”
	Type_ string `json:"type,omitempty"`
	// username authenticated with
	Username string `json:"username,omitempty"`
}
