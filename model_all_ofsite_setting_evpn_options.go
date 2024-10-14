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

type AllOfsiteSettingEvpnOptions struct {
	// optional, for dhcp_relay, unique loopback IPs are required for ERB or IPClos where we can set option-82 server_id-overrides
	AutoLoopbackSubnet string `json:"auto_loopback_subnet,omitempty"`
	// optional, for dhcp_relay, unique loopback IPs are required for ERB or IPClos where we can set option-82 server_id-overrides
	AutoLoopbackSubnet6 string `json:"auto_loopback_subnet6,omitempty"`
	// optional, this generates router_id automatically, if specified, `router_id_prefix` is ignored
	AutoRouterIdSubnet string `json:"auto_router_id_subnet,omitempty"`
	// optional, this generates router_id automatically, if specified, `router_id_prefix` is ignored
	AutoRouterIdSubnet6 string `json:"auto_router_id_subnet6,omitempty"`
	// optional, for ERB or CLOS, you can either use esilag to upstream routers or to also be the virtual-gateway when `routed_at` != `core`, whether to do virtual-gateway at core as well
	CoreAsBorder bool `json:"core_as_border,omitempty"`
	Overlay *EvpnOptionsOverlay `json:"overlay,omitempty"`
	// by default, JUNOS uses 00-00-5e-00-01-01 as the virtual-gateway-address's v4_mac if enabled, 00-00-5e-00-XX-YY will be used (where XX=vlan_id/256, YY=vlan_id%256)
	PerVlanVgaV4Mac bool `json:"per_vlan_vga_v4_mac,omitempty"`
	RoutedAt *any `json:"routed_at,omitempty"`
	Underlay *EvpnOptionsUnderlay `json:"underlay,omitempty"`
	// optional, for EX9200 only to seggregate virtual-switches
	VsInstances map[string]EvpnOptionsVsInstance `json:"vs_instances,omitempty"`
}
