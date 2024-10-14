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

type AllOfsiteSettingMxtunnels struct {
	AdditionalMxtunnels map[string]SiteMxtunnelAdditionalMxtunnel `json:"additional_mxtunnels,omitempty"`
	// list of subnets where we allow AP to establish Mist Tunnels from
	ApSubnets []string `json:"ap_subnets,omitempty"`
	AutoPreemption *any `json:"auto_preemption,omitempty"`
	// for AP, how to connect to tunterm or radsecproxy
	Clusters []SiteMxtunnelCluster `json:"clusters,omitempty"`
	CreatedTime float64 `json:"created_time,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	ForSite bool `json:"for_site,omitempty"`
	// in seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by hello_retries
	HelloInterval int32 `json:"hello_interval,omitempty"`
	HelloRetries int32 `json:"hello_retries,omitempty"`
	// hostnames or IPs where a Mist Tunnel will use as the Peer (i.e. they are reachable from AP)
	Hosts []string `json:"hosts,omitempty"`
	Id string `json:"id,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	// 0 to enable PMTU, 552-1500 to start PMTU with a lower MTU
	Mtu int32 `json:"mtu,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	Protocol *any `json:"protocol,omitempty"`
	Radsec *SiteMxtunnelRadsec `json:"radsec,omitempty"`
	SiteId string `json:"site_id,omitempty"`
	// list of vlan_ids that will be used
	VlanIds []int32 `json:"vlan_ids,omitempty"`
}
