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

// Initiate a Radiotap Packet Capture and Wired Packet Capture
type CaptureRadiotapwired struct {
	ApMac string `json:"ap_mac,omitempty"`
	Band *AllOfcaptureRadiotapwiredBand `json:"band,omitempty"`
	ClientMac string `json:"client_mac,omitempty"`
	// duration of the capture, in seconds
	Duration int32 `json:"duration,omitempty"`
	Format *AllOfcaptureRadiotapwiredFormat `json:"format,omitempty"`
	// max_len of each packet to capture
	MaxPktLen int32 `json:"max_pkt_len,omitempty"`
	// number of packets to capture, 0 for unlimited
	NumPackets int32 `json:"num_packets,omitempty"`
	// tcpdump expression for radiotap interface (802.11 + radio headers)
	RadiotapTcpdumpExpression string `json:"radiotap_tcpdump_expression,omitempty"`
	Ssid string `json:"ssid,omitempty"`
	// tcpdump expression common for wired,radiotap
	TcpdumpExpression string `json:"tcpdump_expression,omitempty"`
	// enum: `radiotap,wired`
	Type_ string `json:"type"`
	// tcpdump expression for wired
	WiredTcpdumpExpression string `json:"wired_tcpdump_expression,omitempty"`
	// tcpdump expression for radiotap interface (802.11)
	WirelessTcpdumpExpression string `json:"wireless_tcpdump_expression,omitempty"`
	// wlan id associated with the respective ssid.
	WlanId string `json:"wlan_id,omitempty"`
}
