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

type InlineResponse20066 struct {
	CreatedTime float64 `json:"created_time,omitempty"`
	ForSite bool `json:"for_site,omitempty"`
	Id string `json:"id,omitempty"`
	MistDas *any `json:"mist_das,omitempty"`
	MistNac *MxclusterNac `json:"mist_nac,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	MxedgeMgmt *MxedgeMgmt `json:"mxedge_mgmt,omitempty"`
	Name string `json:"name,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	Proxy *any `json:"proxy,omitempty"`
	Radsec *any `json:"radsec,omitempty"`
	RadsecTls *MxclusterRadsecTls `json:"radsec_tls,omitempty"`
	SiteId string `json:"site_id,omitempty"`
	// list of subnets where we allow AP to establish Mist Tunnels from
	TuntermApSubnets []string `json:"tunterm_ap_subnets,omitempty"`
	TuntermDhcpdConfig *any `json:"tunterm_dhcpd_config,omitempty"`
	// extra routes for Mist Tunneled VLANs. Property key is a CIDR
	TuntermExtraRoutes map[string]MxclusterTuntermExtraRoute `json:"tunterm_extra_routes,omitempty"`
	// hostnames or IPs where a Mist Tunnel will use as the Peer (i.e. they are reachable from AP)
	TuntermHosts []string `json:"tunterm_hosts,omitempty"`
	// list of index of tunterm_hosts
	TuntermHostsOrder []int32 `json:"tunterm_hosts_order,omitempty"`
	TuntermHostsSelection *any `json:"tunterm_hosts_selection,omitempty"`
	TuntermMonitoring [][]TuntermMonitoringItem `json:"tunterm_monitoring,omitempty"`
	TuntermMonitoringDisabled bool `json:"tunterm_monitoring_disabled,omitempty"`
}
