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

type InlineResponse200117 struct {
	CreatedTime float64 `json:"created_time,omitempty"`
	Dmvpn *Object `json:"dmvpn,omitempty"`
	// determined during creation time and cannot be toggled. A management tunnel cannot be used by wxlan rule or by wlan
	ForMgmt bool `json:"for_mgmt,omitempty"`
	ForSite bool `json:"for_site,omitempty"`
	// in seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by hello_retries.
	HelloInterval int32 `json:"hello_interval,omitempty"`
	HelloRetries int32 `json:"hello_retries,omitempty"`
	// optional, overwrite the hostname in SCCRQ control message, default is \\u201C\\u201D or null, %H and %M can be used, which will be replace with corresponding values:   * %H: name of the ap if provided (and will be stripped so it can be used for hostname) and fallbacks to MAC   * %M: MAC (e.g. 5c5b350e0060)
	Hostname string `json:"hostname,omitempty"`
	Id string `json:"id,omitempty"`
	Ipsec *Object `json:"ipsec,omitempty"`
	// whether it’s static/unmanaged (i.e. no control session). As the session configurations are not compatible, cannot be toggled.
	IsStatic bool `json:"is_static,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	// 0 to enable PMTU, 552-1500 to start PMTU with a lower MTU
	Mtu int32 `json:"mtu,omitempty"`
	// The name of the tunnel
	Name string `json:"name"`
	OrgId string `json:"org_id,omitempty"`
	// list of remote peers’ IP or hostname
	Peers []string `json:"peers,omitempty"`
	// optional, overwrite the router-id in SCCRQ control message, default is “0” or null, can also be an IPv4 address
	RouterId string `json:"router_id,omitempty"`
	// secret, ‘’ if no auth is used
	Secret string `json:"secret,omitempty"`
	// sessions to be established with the tunnel. Has to be >= 1 in order for this tunnel to be useful. For management tunnel, it can only have 1
	Sessions []WxlanTunnelSession `json:"sessions,omitempty"`
	SiteId string `json:"site_id,omitempty"`
	// udp port if `use_udp`==`true`
	UdpPort int32 `json:"udp_port,omitempty"`
	// whether to use UDP instead of IP (proto=115, which is default of L2TPv3)
	UseUdp bool `json:"use_udp,omitempty"`
}
