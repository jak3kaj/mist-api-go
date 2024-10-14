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

type StatsWanClient struct {
	DhcpExpireTime float64 `json:"dhcp_expire_time,omitempty"`
	DhcpStartTime float64 `json:"dhcp_start_time,omitempty"`
	Hostname []string `json:"hostname,omitempty"`
	Ip []string `json:"ip,omitempty"`
	IpSrc string `json:"ip_src,omitempty"`
	LastHostname string `json:"last_hostname,omitempty"`
	LastIp string `json:"last_ip,omitempty"`
	Mfg string `json:"mfg,omitempty"`
	Network string `json:"network,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	SiteId string `json:"site_id,omitempty"`
	Timestamp float64 `json:"timestamp,omitempty"`
	Wcid string `json:"wcid,omitempty"`
}
