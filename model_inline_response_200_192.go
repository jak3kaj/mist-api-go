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

type InlineResponse200192 struct {
	ApMac string `json:"ap_mac,omitempty"`
	// List of target APs to capture packets
	Aps []string `json:"aps,omitempty"`
	ClientMac string `json:"client_mac,omitempty"`
	Duration int32 `json:"duration,omitempty"`
	// List of APs where configuration attempt failed
	Failed []string `json:"failed,omitempty"`
	Format *any `json:"format,omitempty"`
	// Information on gateways to capture packets on if a gateway capture type is specified
	Gateways []string `json:"gateways,omitempty"`
	// unique id for the capture
	Id string `json:"id"`
	IncludesMcast bool `json:"includes_mcast,omitempty"`
	// max number of packets configured by user
	MaxNumPackets int32 `json:"max_num_packets,omitempty"`
	MaxPktLen int32 `json:"max_pkt_len,omitempty"`
	// total number of packets captured by all AP, not applicable for type [client, new_assoc]
	NumPackets int32 `json:"num_packets,omitempty"`
	// List of target APs successfully configured to capture packets
	Ok []string `json:"ok,omitempty"`
	PcapAps map[string]ResponsePcapAp `json:"pcap_aps,omitempty"`
	// when `type`==`radiotap`, radiotap_tcpdump_expression expression provided by the user
	RadiotapTcpdumpExpression string `json:"radiotap_tcpdump_expression,omitempty"`
	// when `type`==`scan`, scan_tcpdump_expression provided by the user
	ScanTcpdumpExpression string `json:"scan_tcpdump_expression,omitempty"`
	Ssid string `json:"ssid,omitempty"`
	StartedTime int32 `json:"started_time,omitempty"`
	// Information on switches to capture packets on if a switch capture type is specified. irb port interface is automatically added to capture as needed to ensure all desired packets are captured.
	Switches []string `json:"switches,omitempty"`
	// tcpdump expression provided by the user (common)
	TcpdumpExpression string `json:"tcpdump_expression,omitempty"`
	Type_ *any `json:"type"`
	// Required if `format`==`tzsp`. Remote host accessible to mxedges over the network for receiving the captured packets.
	TzspHost string `json:"tzsp_host,omitempty"`
	// if `format`==`tzsp`. Port on remote host for receiving the captured packets
	TzspPort int32 `json:"tzsp_port,omitempty"`
	// when `type`==`wired`, wired_tcpdump_expression provided by the user
	WiredTcpdumpExpression string `json:"wired_tcpdump_expression,omitempty"`
	// when `type`==`‘wireless’`, wireless_tcpdump_expression provided by the user
	WirelessTcpdumpExpression string `json:"wireless_tcpdump_expression,omitempty"`
}
