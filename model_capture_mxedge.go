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

// Initiate a Wireless Packet Capture
type CaptureMxedge struct {
	// duration of the capture, in seconds
	Duration int32 `json:"duration,omitempty"`
	Format *AllOfcaptureMxedgeFormat `json:"format,omitempty"`
	// max_len of each packet to capture
	MaxPktLen int32 `json:"max_pkt_len,omitempty"`
	Mxedges map[string]CaptureMxedgeMxedges `json:"mxedges,omitempty"`
	// number of packets to capture, 0 for unlimited
	NumPackets int32 `json:"num_packets,omitempty"`
	// enum: `mxedge`
	Type_ string `json:"type"`
	// Required if `format`==`tzsp`. Remote host accessible to mxedges over the network for receiving the captured packets.
	TzspHost string `json:"tzsp_host,omitempty"`
	// if `format`==`tzsp`. Port on remote host for receiving the captured packets
	TzspPort int32 `json:"tzsp_port,omitempty"`
}
