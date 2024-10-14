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

// Initiate a Switch (Junos) Packet Capture
type CaptureSwitch struct {
	// duration of the capture, in seconds
	Duration int32 `json:"duration,omitempty"`
	Format *AllOfcaptureSwitchFormat `json:"format,omitempty"`
	// max_len of each packet to capture
	MaxPktLen int32 `json:"max_pkt_len,omitempty"`
	// number of packets to capture, 0 for unlimited
	NumPackets int32 `json:"num_packets,omitempty"`
	// Property key is the port name. 6 ports max per switch supported, or 5 max with irb port auto-included into capture request
	Ports map[string]CaptureSwitchPortsTcpdumpExpression `json:"ports,omitempty"`
	// Property key is the switch mac
	Switches map[string]CaptureSwitchSwitches `json:"switches,omitempty"`
	// tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist.
	TcpdumpExpression string `json:"tcpdump_expression,omitempty"`
	// enum: `switch`
	Type_ string `json:"type"`
}
