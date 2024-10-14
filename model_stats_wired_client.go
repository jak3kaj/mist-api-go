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

type StatsWiredClient struct {
	// client authorization status
	AuthState string `json:"auth_state,omitempty"`
	// Device ID the client is connected to
	DeviceId string `json:"device_id,omitempty"`
	// port on AP where the wired client is connected
	EthPort string `json:"eth_port,omitempty"`
	// time when last Tx/Rx observed
	LastSeen float64 `json:"last_seen,omitempty"`
	// client mac
	Mac string `json:"mac"`
	// amount of traffic sent to client since client connects
	RxBytes float64 `json:"rx_bytes,omitempty"`
	// amount of traffic sent to client since client connects
	RxPkts float64 `json:"rx_pkts,omitempty"`
	SiteId string `json:"site_id,omitempty"`
	// amount of traffic received from client since client connects
	TxBytes float64 `json:"tx_bytes,omitempty"`
	// amount of traffic received from client since client connects
	TxPkts float64 `json:"tx_pkts,omitempty"`
	// how long, in seconds, has the client been connected
	Uptime float64 `json:"uptime,omitempty"`
	// vlan id, could be empty
	VlanId float64 `json:"vlan_id,omitempty"`
}
