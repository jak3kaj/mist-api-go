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

type ApStatMeshDownlink struct {
	Band string `json:"band,omitempty"`
	Channel int32 `json:"channel,omitempty"`
	IdleTime int32 `json:"idle_time,omitempty"`
	LastSeen float64 `json:"last_seen,omitempty"`
	Proto string `json:"proto,omitempty"`
	Rssi int32 `json:"rssi,omitempty"`
	RxBps int32 `json:"rx_bps,omitempty"`
	RxBytes int32 `json:"rx_bytes,omitempty"`
	RxPackets int32 `json:"rx_packets,omitempty"`
	RxRate int32 `json:"rx_rate,omitempty"`
	RxRetries int32 `json:"rx_retries,omitempty"`
	SiteId string `json:"site_id,omitempty"`
	Snr int32 `json:"snr,omitempty"`
	TxBps int32 `json:"tx_bps,omitempty"`
	TxBytes int32 `json:"tx_bytes,omitempty"`
	TxPackets int32 `json:"tx_packets,omitempty"`
	TxRate int32 `json:"tx_rate,omitempty"`
	TxRetries int32 `json:"tx_retries,omitempty"`
}
