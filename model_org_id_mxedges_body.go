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

type OrgIdMxedgesBody struct {
	CreatedTime float64 `json:"created_time,omitempty"`
	ForSite bool `json:"for_site,omitempty"`
	Id string `json:"id,omitempty"`
	Magic string `json:"magic,omitempty"`
	Model string `json:"model"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	MxagentRegistered bool `json:"mxagent_registered,omitempty"`
	// MxCluster this MxEdge belongs to
	MxclusterId string `json:"mxcluster_id,omitempty"`
	MxedgeMgmt *MxedgeMgmt `json:"mxedge_mgmt,omitempty"`
	Name string `json:"name"`
	Note string `json:"note,omitempty"`
	NtpServers []string `json:"ntp_servers,omitempty"`
	OobIpConfig *any `json:"oob_ip_config,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	Proxy *any `json:"proxy,omitempty"`
	// list of services to run, tunterm only for now
	Services []string `json:"services,omitempty"`
	SiteId string `json:"site_id,omitempty"`
	TuntermDhcpdConfig *any `json:"tunterm_dhcpd_config,omitempty"`
	// Property key is a CIDR
	TuntermExtraRoutes map[string]MxedgeTuntermExtraRoute `json:"tunterm_extra_routes,omitempty"`
	TuntermIgmpSnoopingConfig *MxedgeTuntermIgmpSnoopingConfig `json:"tunterm_igmp_snooping_config,omitempty"`
	TuntermIpConfig *any `json:"tunterm_ip_config,omitempty"`
	TuntermMonitoring [][]TuntermMonitoringItem `json:"tunterm_monitoring,omitempty"`
	TuntermMulticastConfig *MxedgeTuntermMulticastConfig `json:"tunterm_multicast_config,omitempty"`
	// ip configs by VLAN ID. Property key is the VLAN ID
	TuntermOtherIpConfigs map[string]MxedgeTuntermOtherIpConfig `json:"tunterm_other_ip_configs,omitempty"`
	TuntermPortConfig *any `json:"tunterm_port_config,omitempty"`
	TuntermRegistered bool `json:"tunterm_registered,omitempty"`
	TuntermSwitchConfig *any `json:"tunterm_switch_config,omitempty"`
	Versions *any `json:"versions,omitempty"`
}
