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

type AllOfreplaceDeviceTuntermPortConfig struct {
	// list of ports to be used for downstream (to AP) purpose
	DownstreamPorts []string `json:"downstream_ports,omitempty"`
	// weather to separate upstream / downstream ports. default is false where all ports will be used.
	SeparateUpstreamDownstream bool `json:"separate_upstream_downstream,omitempty"`
	// native VLAN id for upstream ports
	UpstreamPortVlanId int32 `json:"upstream_port_vlan_id,omitempty"`
	// list of ports to be used for upstrea purpose (to LAN)
	UpstreamPorts []string `json:"upstream_ports,omitempty"`
}
