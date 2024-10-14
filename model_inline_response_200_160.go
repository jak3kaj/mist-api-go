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

type InlineResponse200160 struct {
	By string `json:"by,omitempty"`
	DeviceType *Object `json:"device_type,omitempty"`
	Failed bool `json:"failed,omitempty"`
	Latency int32 `json:"latency,omitempty"`
	Mac string `json:"mac,omitempty"`
	PortId string `json:"port_id,omitempty"`
	Reason string `json:"reason,omitempty"`
	RxMbps int32 `json:"rx_mbps,omitempty"`
	StartTime int32 `json:"start_time,omitempty"`
	Status string `json:"status,omitempty"`
	Timestamp float64 `json:"timestamp,omitempty"`
	TxMbps int32 `json:"tx_mbps,omitempty"`
	Type_ *Object `json:"type,omitempty"`
	VlanId int32 `json:"vlan_id,omitempty"`
}
