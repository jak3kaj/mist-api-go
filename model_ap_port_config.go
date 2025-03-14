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

type ApPortConfig struct {
	Disabled bool `json:"disabled,omitempty"`
	DynamicVlan *AllOfapPortConfigDynamicVlan `json:"dynamic_vlan,omitempty"`
	EnableMacAuth bool `json:"enable_mac_auth,omitempty"`
	Forwarding *AllOfapPortConfigForwarding `json:"forwarding,omitempty"`
	// when `true`, we'll do dot1x then mac_auth. enable this to prefer mac_auth
	MacAuthPreferred bool `json:"mac_auth_preferred,omitempty"`
	MacAuthProtocol *AllOfapPortConfigMacAuthProtocol `json:"mac_auth_protocol,omitempty"`
	MistNac *WlanMistNac `json:"mist_nac,omitempty"`
	// if `forwarding`==`mxtunnel`, vlan_ids comes from mxtunnel
	MxTunnelId string `json:"mx_tunnel_id,omitempty"`
	// if `forwarding`==`site_mxedge`, vlan_ids comes from site_mxedge (`mxtunnels` under site setting)
	MxtunnelName string `json:"mxtunnel_name,omitempty"`
	PortAuth *AllOfapPortConfigPortAuth `json:"port_auth,omitempty"`
	// if `forwrding`==`limited`
	PortVlanId int32 `json:"port_vlan_id,omitempty"`
	RadiusConfig *AllOfapPortConfigRadiusConfig `json:"radius_config,omitempty"`
	Radsec *AllOfapPortConfigRadsec `json:"radsec,omitempty"`
	// optional to specify the vlan id for a tunnel if forwarding is for `wxtunnel`, `mxtunnel` or `site_mxedge`.   * if vlan_id is not specified then it will use first one in vlan_ids[] of the mxtunnel.   * if forwarding == site_mxedge, vlan_ids comes from site_mxedge (`mxtunnels` under site setting)
	VlanId int32 `json:"vlan_id,omitempty"`
	// if `forwrding`==`limited`
	VlandIds []int32 `json:"vland_ids,omitempty"`
	// if `forwarding`==`wxtunnel`, the port is bridged to the vlan of the session
	WxtunnelId string `json:"wxtunnel_id,omitempty"`
	// if `forwarding`==`wxtunnel`, the port is bridged to the vlan of the session
	WxtunnelRemoteId string `json:"wxtunnel_remote_id,omitempty"`
}
