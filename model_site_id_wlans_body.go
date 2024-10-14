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

type SiteIdWlansBody struct {
	// enable coa-immediate-update and address-change-immediate-update on the access profile.
	AcctImmediateUpdate bool `json:"acct_immediate_update,omitempty"`
	// how frequently should interim accounting be reported, 60-65535. default is 0 (use one specified in Access-Accept request from RADIUS Server). Very frequent messages can affect the performance of the radius server, 600 and up is recommended when enabled
	AcctInterimInterval int32 `json:"acct_interim_interval,omitempty"`
	// list of RADIUS accounting servers, optional, order matters where the first one is treated as primary
	AcctServers []RadiusAcctServer `json:"acct_servers,omitempty"`
	Airwatch *Object `json:"airwatch,omitempty"`
	// only applicable when limit_bcast==true, which allows or disallows ipv6 Neighbor Discovery packets to go through
	AllowIpv6Ndp bool `json:"allow_ipv6_ndp,omitempty"`
	// only applicable when limit_bcast==true, which allows mDNS / Bonjour packets to go through
	AllowMdns bool `json:"allow_mdns,omitempty"`
	// only applicable when `limit_bcast`==`true`, which allows SSDP
	AllowSsdp bool `json:"allow_ssdp,omitempty"`
	// list of device ids
	ApIds []string `json:"ap_ids,omitempty"`
	AppLimit *Object `json:"app_limit,omitempty"`
	AppQos *Object `json:"app_qos,omitempty"`
	ApplyTo *Object `json:"apply_to,omitempty"`
	// whether to enable smart arp filter
	ArpFilter bool `json:"arp_filter,omitempty"`
	Auth *Object `json:"auth,omitempty"`
	AuthServerSelection *Object `json:"auth_server_selection,omitempty"`
	// list of RADIUS authentication servers, at least one is needed if `auth type`==`eap`, order matters where the first one is treated as primary
	AuthServers []RadiusAuthServer `json:"auth_servers,omitempty"`
	// optional, up to 48 bytes, will be dynamically generated if not provided. used only for authentication servers
	AuthServersNasId string `json:"auth_servers_nas_id,omitempty"`
	// optional, NAS-IP-ADDRESS to use
	AuthServersNasIp string `json:"auth_servers_nas_ip,omitempty"`
	// radius auth session retries. Following fast timers are set if “fast_dot1x_timers” knob is enabled. ‘retries’  are set to value of auth_servers_retries. ‘max-requests’ is also set when setting auth_servers_retries and is set to default value to 3.
	AuthServersRetries int32 `json:"auth_servers_retries,omitempty"`
	// radius auth session timeout. Following fast timers are set if “fast_dot1x_timers” knob is enabled. ‘quite-period’  and ‘transmit-period’ are set to half the value of auth_servers_timeout. ‘supplicant-timeout’ is also set when setting auth_servers_timeout and is set to default value of 10.
	AuthServersTimeout int32 `json:"auth_servers_timeout,omitempty"`
	// `band` is deprecated and kept for backward compability. Use bands instead
	Band string `json:"band,omitempty"`
	// whether to enable band_steering, this works only when band==both
	BandSteer bool `json:"band_steer,omitempty"`
	// force dual_band capable client to connect to 5G
	BandSteerForceBand5 bool `json:"band_steer_force_band5,omitempty"`
	// list of radios that the wlan should apply to.
	Bands []Dot11Band `json:"bands,omitempty"`
	// whether to block the clients in the blacklist (up to first 256 macs)
	BlockBlacklistClients bool `json:"block_blacklist_clients,omitempty"`
	Bonjour *Object `json:"bonjour,omitempty"`
	CiscoCwa *Object `json:"cisco_cwa,omitempty"`
	// kbps
	ClientLimitDown int32 `json:"client_limit_down,omitempty"`
	// if downlink limiting per-client is enabled
	ClientLimitDownEnabled bool `json:"client_limit_down_enabled,omitempty"`
	// kbps
	ClientLimitUp int32 `json:"client_limit_up,omitempty"`
	// if uplink limiting per-client is enabled
	ClientLimitUpEnabled bool `json:"client_limit_up_enabled,omitempty"`
	// list of COA (change of authorization) servers, optional
	CoaServers []CoaServer `json:"coa_servers,omitempty"`
	CreatedTime float64 `json:"created_time,omitempty"`
	// some old WLAN drivers may not be compatible
	Disable11ax bool `json:"disable_11ax,omitempty"`
	// to disable ht or vht rates
	DisableHtVhtRates bool `json:"disable_ht_vht_rates,omitempty"`
	// whether to disable U-APSD
	DisableUapsd bool `json:"disable_uapsd,omitempty"`
	// disable sending v2 roam notification messages
	DisableV1RoamNotify bool `json:"disable_v1_roam_notify,omitempty"`
	// disable sending v2 roam notification messages
	DisableV2RoamNotify bool `json:"disable_v2_roam_notify,omitempty"`
	// whether to disable WMM
	DisableWmm bool `json:"disable_wmm,omitempty"`
	DnsServerRewrite *Object `json:"dns_server_rewrite,omitempty"`
	Dtim int32 `json:"dtim,omitempty"`
	DynamicPsk *Object `json:"dynamic_psk,omitempty"`
	DynamicVlan *Object `json:"dynamic_vlan,omitempty"`
	// enable AP-AP keycaching via multicast
	EnableLocalKeycaching bool `json:"enable_local_keycaching,omitempty"`
	// by default, we'd inspect all DHCP packets and drop those unrelated to the wireless client itself in the case where client is a wireless bridge (DHCP packets for other MACs will need to be orwarded), wireless_bridging can be enabled
	EnableWirelessBridging bool `json:"enable_wireless_bridging,omitempty"`
	// if the client bridge is doing DHCP on behalf of other devices (L2-NAT), enable dhcp_tracking will cut down DHCP response packets to be forwarded to wireless
	EnableWirelessBridgingDhcpTracking bool `json:"enable_wireless_bridging_dhcp_tracking,omitempty"`
	// if this wlan is enabled
	Enabled bool `json:"enabled,omitempty"`
	// if set to true, sets default fast-timers with values calculated from ‘auth_servers_timeout’ and ‘auth_server_retries’ .
	FastDot1xTimers bool `json:"fast_dot1x_timers,omitempty"`
	ForSite bool `json:"for_site,omitempty"`
	// whether to hide SSID in beacon
	HideSsid bool `json:"hide_ssid,omitempty"`
	// include hostname inside IE in AP beacons / probe responses
	HostnameIe bool `json:"hostname_ie,omitempty"`
	Hotspot20 *Object `json:"hotspot20,omitempty"`
	Id string `json:"id,omitempty"`
	InjectDhcpOption82 *WlanInjectDhcpOption82 `json:"inject_dhcp_option_82,omitempty"`
	Interface_ *Object `json:"interface,omitempty"`
	// whether to stop clients to talk to each other
	Isolation bool `json:"isolation,omitempty"`
	// if isolation is enabled, whether to deny clients to talk to L2 on the LAN
	L2Isolation bool `json:"l2_isolation,omitempty"`
	// legacy devices requires the Over-DS (for Fast BSS Transition) bit set (while our chip doesn’t support it). Warning! Enabling this will cause problem for iOS devices.
	LegacyOverds bool `json:"legacy_overds,omitempty"`
	// whether to limit broadcast packets going to wireless (i.e. only allow certain bcast packets to go through)
	LimitBcast bool `json:"limit_bcast,omitempty"`
	// limit probe response base on some heuristic rules
	LimitProbeResponse bool `json:"limit_probe_response,omitempty"`
	// max idle time in seconds
	MaxIdletime int32 `json:"max_idletime,omitempty"`
	// maximum number of client connected to the SSID. `0` means unlimited
	MaxNumClients int32 `json:"max_num_clients,omitempty"`
	MistNac *WlanMistNac `json:"mist_nac,omitempty"`
	ModifiedTime float64 `json:"modified_time,omitempty"`
	MspId string `json:"msp_id,omitempty"`
	// (deprecated, use mxtunnel_ids instead) when `interface`==`mxtunnel`, id of the Mist Tunnel
	MxtunnelId string `json:"mxtunnel_id,omitempty"`
	// when `interface`=`mxtunnel`, id of the Mist Tunnel
	MxtunnelIds []string `json:"mxtunnel_ids,omitempty"`
	// when `interface`=`site_medge`, name of the mxtunnel that in mxtunnels under Site Setting
	MxtunnelName []string `json:"mxtunnel_name,omitempty"`
	// whether to only allow client to use DNS that we’ve learned from DHCP response
	NoStaticDns bool `json:"no_static_dns,omitempty"`
	// whether to only allow client that we’ve learned from DHCP exchange to talk
	NoStaticIp bool `json:"no_static_ip,omitempty"`
	OrgId string `json:"org_id,omitempty"`
	Portal *Object `json:"portal,omitempty"`
	// list of hostnames without http(s):// (matched by substring)
	PortalAllowedHostnames []string `json:"portal_allowed_hostnames,omitempty"`
	// list of CIDRs
	PortalAllowedSubnets []string `json:"portal_allowed_subnets,omitempty"`
	// api secret (auto-generated) that can be used to sign guest authorization requests
	PortalApiSecret string `json:"portal_api_secret,omitempty"`
	// list of hostnames without http(s):// (matched by substring), this takes precedence over portal_allowed_hostnames
	PortalDeniedHostnames []string `json:"portal_denied_hostnames,omitempty"`
	// Url of portal background image
	PortalImage string `json:"portal_image,omitempty"`
	PortalSsoUrl string `json:"portal_sso_url,omitempty"`
	// N.B portal_template will be forked out of wlan objects soon. To fetch portal_template, please query portal_template_url. To update portal_template, use Wlan Portal Template.
	PortalTemplateUrl string `json:"portal_template_url,omitempty"`
	Qos *WlanQos `json:"qos,omitempty"`
	Radsec *Object `json:"radsec,omitempty"`
	Rateset map[string]WlanDatarates `json:"rateset,omitempty"`
	RoamMode *Object `json:"roam_mode,omitempty"`
	Schedule *Object `json:"schedule,omitempty"`
	SiteId string `json:"site_id,omitempty"`
	// whether to exclude this WLAN from SLE metrics
	SleExcluded bool `json:"sle_excluded,omitempty"`
	// the name of the SSID
	Ssid string `json:"ssid"`
	TemplateId string `json:"template_id,omitempty"`
	// Url of portal background image thumbnail
	Thumbnail string `json:"thumbnail,omitempty"`
	// if `auth.type`==’eap’ or ‘psk’, should only be set for legacy client, such as pre-2004, 802.11b devices
	UseEapolV1 bool `json:"use_eapol_v1,omitempty"`
	// if vlan tagging is enabled
	VlanEnabled bool `json:"vlan_enabled,omitempty"`
	VlanId *Object `json:"vlan_id,omitempty"`
	// vlan_ids to use when there’s no match from RA
	VlanIds []VlanIdWithVariable `json:"vlan_ids,omitempty"`
	// vlan pooling allows AP to place client on different VLAN using a deterministic algorithm
	VlanPooling bool `json:"vlan_pooling,omitempty"`
	// kbps
	WlanLimitDown int32 `json:"wlan_limit_down,omitempty"`
	// if downlink limiting for whole wlan is enabled
	WlanLimitDownEnabled bool `json:"wlan_limit_down_enabled,omitempty"`
	// kbps
	WlanLimitUp int32 `json:"wlan_limit_up,omitempty"`
	// if uplink limiting for whole wlan is enabled
	WlanLimitUpEnabled bool `json:"wlan_limit_up_enabled,omitempty"`
	// list of wxtag_ids
	WxtagIds []string `json:"wxtag_ids,omitempty"`
	// when `interface`=`wxtunnel`, id of the WXLAN Tunnel
	WxtunnelId string `json:"wxtunnel_id,omitempty"`
	// when `interface`=`wxtunnel`, remote tunnel identifier
	WxtunnelRemoteId string `json:"wxtunnel_remote_id,omitempty"`
}
