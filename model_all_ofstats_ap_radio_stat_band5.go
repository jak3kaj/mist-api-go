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

type AllOfstatsApRadioStatBand5 struct {
	Bandwidth *any `json:"bandwidth,omitempty"`
	// current channel the radio is running on
	Channel int32 `json:"channel,omitempty"`
	// Use dynamic chaining for downlink
	DynamicChainingEnalbed bool `json:"dynamic_chaining_enalbed,omitempty"`
	// radio (base) mac, it can have 16 bssids (e.g. 5c5b350001a0-5c5b350001af)
	Mac string `json:"mac,omitempty"`
	NoiseFloor int32 `json:"noise_floor,omitempty"`
	NumClients int32 `json:"num_clients,omitempty"`
	// how many WLANs are applied to the radio
	NumWlans int32 `json:"num_wlans,omitempty"`
	// transmit power (in dBm)
	Power int32 `json:"power,omitempty"`
	RxBytes int32 `json:"rx_bytes,omitempty"`
	RxPkts int32 `json:"rx_pkts,omitempty"`
	TxBytes int32 `json:"tx_bytes,omitempty"`
	TxPkts int32 `json:"tx_pkts,omitempty"`
	Usage string `json:"usage,omitempty"`
	// all utilization in percentage
	UtilAll int32 `json:"util_all,omitempty"`
	// reception of “No Packets” utilization in percentage, received frames with invalid PLCPs and CRS glitches as noise
	UtilNonWifi int32 `json:"util_non_wifi,omitempty"`
	// reception of “In BSS” utilization in percentage, only frames that are received from AP/STAs within the BSS
	UtilRxInBss int32 `json:"util_rx_in_bss,omitempty"`
	// reception of “Other BSS” utilization in percentage, all frames received from AP/STAs that are outside the BSS
	UtilRxOtherBss int32 `json:"util_rx_other_bss,omitempty"`
	// transmission utilization in percentage
	UtilTx int32 `json:"util_tx,omitempty"`
	// reception of “UnDecodable Wifi“ utilization in percentage, only Preamble, PLCP header is decoded, Rest is undecodable in this radio
	UtilUndecodableWifi int32 `json:"util_undecodable_wifi,omitempty"`
	// reception of “No Category” utilization in percentage, all 802.11 frames that are corrupted at the receiver
	UtilUnknownWifi int32 `json:"util_unknown_wifi,omitempty"`
}
