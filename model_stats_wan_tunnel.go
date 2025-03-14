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

type StatsWanTunnel struct {
	// authentication algorithm
	AuthAlgo string `json:"auth_algo,omitempty"`
	// encryption algorithm
	EncryptAlgo string `json:"encrypt_algo,omitempty"`
	// ike version
	IkeVersion string `json:"ike_version,omitempty"`
	// ip address
	Ip string `json:"ip,omitempty"`
	// reason of why the tunnel is down
	LastEvent string `json:"last_event,omitempty"`
	// router mac address
	Mac string `json:"mac,omitempty"`
	// node0/node1
	Node string `json:"node,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	// peer host
	PeerHost string `json:"peer_host,omitempty"`
	// peer ip address
	PeerIp string `json:"peer_ip,omitempty"`
	Priority *AllOfstatsWanTunnelPriority `json:"priority,omitempty"`
	Protocol *AllOfstatsWanTunnelProtocol `json:"protocol,omitempty"`
	RxBytes int32 `json:"rx_bytes,omitempty"`
	RxPkts int32 `json:"rx_pkts,omitempty"`
	SiteId string `json:"site_id,omitempty"`
	// Mist Tunnel Name
	TunnelName string `json:"tunnel_name,omitempty"`
	TxBytes int32 `json:"tx_bytes,omitempty"`
	TxPkts int32 `json:"tx_pkts,omitempty"`
	Up bool `json:"up,omitempty"`
	// duration from first (or last) SA was established
	Uptime int32 `json:"uptime,omitempty"`
	// wan interface name
	WanName string `json:"wan_name,omitempty"`
}
