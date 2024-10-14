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

type IfStatProperty struct {
	AddressMode string `json:"address_mode,omitempty"`
	Ips []string `json:"ips,omitempty"`
	NatAddresses []string `json:"nat_addresses,omitempty"`
	NetworkName string `json:"network_name,omitempty"`
	PortId string `json:"port_id,omitempty"`
	PortUsage string `json:"port_usage,omitempty"`
	RedundancyState string `json:"redundancy_state,omitempty"`
	RxBytes int32 `json:"rx_bytes,omitempty"`
	RxPkts int32 `json:"rx_pkts,omitempty"`
	ServpInfo *IfStatPropertyServpInfo `json:"servp_info,omitempty"`
	TxBytes int32 `json:"tx_bytes,omitempty"`
	TxPkts int32 `json:"tx_pkts,omitempty"`
	Up bool `json:"up,omitempty"`
	Vlan int32 `json:"vlan,omitempty"`
	WanName string `json:"wan_name,omitempty"`
	WanType string `json:"wan_type,omitempty"`
}
