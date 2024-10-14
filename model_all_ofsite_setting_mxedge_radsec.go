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

type AllOfsiteSettingMxedgeRadsec struct {
	// list of RADIUS accounting servers, optional, order matters where the first one is treated as primary
	AcctServers []MxclusterRadsecAcctServer `json:"acct_servers,omitempty"`
	// list of RADIUS authentication servers, order matters where the first one is treated as primary
	AuthServers []MxclusterRadsecAuthServer `json:"auth_servers,omitempty"`
	// whether to enable service on Mist Edge i.e. RADIUS proxy over TLS
	Enabled bool `json:"enabled,omitempty"`
	// whether to match ssid in request message to select from a subset of RADIUS servers
	MatchSsid bool `json:"match_ssid,omitempty"`
	NasIpSource *any `json:"nas_ip_source,omitempty"`
	// hostnames or IPs for Mist AP to use as the TLS Server (i.e. they are reachable from AP) in addition to `tunterm_hosts`
	ProxyHosts []string `json:"proxy_hosts,omitempty"`
	ServerSelection *any `json:"server_selection,omitempty"`
	SrcIpSource *any `json:"src_ip_source,omitempty"`
}
