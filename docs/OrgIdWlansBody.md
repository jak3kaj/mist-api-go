# OrgIdWlansBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctImmediateUpdate** | **bool** | enable coa-immediate-update and address-change-immediate-update on the access profile. | [optional] [default to false]
**AcctInterimInterval** | **int32** | how frequently should interim accounting be reported, 60-65535. default is 0 (use one specified in Access-Accept request from RADIUS Server). Very frequent messages can affect the performance of the radius server, 600 and up is recommended when enabled | [optional] [default to 0]
**AcctServers** | [**[]RadiusAcctServer**](radius_acct_server.md) | list of RADIUS accounting servers, optional, order matters where the first one is treated as primary | [optional] [default to null]
**Airwatch** | [***Object**](.md) |  | [optional] [default to null]
**AllowIpv6Ndp** | **bool** | only applicable when limit_bcast&#x3D;&#x3D;true, which allows or disallows ipv6 Neighbor Discovery packets to go through | [optional] [default to true]
**AllowMdns** | **bool** | only applicable when limit_bcast&#x3D;&#x3D;true, which allows mDNS / Bonjour packets to go through | [optional] [default to false]
**AllowSsdp** | **bool** | only applicable when &#x60;limit_bcast&#x60;&#x3D;&#x3D;&#x60;true&#x60;, which allows SSDP | [optional] [default to false]
**ApIds** | **[]string** | list of device ids | [optional] [default to null]
**AppLimit** | [***Object**](.md) |  | [optional] [default to null]
**AppQos** | [***Object**](.md) |  | [optional] [default to null]
**ApplyTo** | [***Object**](.md) |  | [optional] [default to null]
**ArpFilter** | **bool** | whether to enable smart arp filter | [optional] [default to false]
**Auth** | [***Object**](.md) |  | [optional] [default to null]
**AuthServerSelection** | [***Object**](.md) |  | [optional] [default to null]
**AuthServers** | [**[]RadiusAuthServer**](radius_auth_server.md) | list of RADIUS authentication servers, at least one is needed if &#x60;auth type&#x60;&#x3D;&#x3D;&#x60;eap&#x60;, order matters where the first one is treated as primary | [optional] [default to null]
**AuthServersNasId** | **string** | optional, up to 48 bytes, will be dynamically generated if not provided. used only for authentication servers | [optional] [default to null]
**AuthServersNasIp** | **string** | optional, NAS-IP-ADDRESS to use | [optional] [default to null]
**AuthServersRetries** | **int32** | radius auth session retries. Following fast timers are set if “fast_dot1x_timers” knob is enabled. ‘retries’  are set to value of auth_servers_retries. ‘max-requests’ is also set when setting auth_servers_retries and is set to default value to 3. | [optional] [default to 2]
**AuthServersTimeout** | **int32** | radius auth session timeout. Following fast timers are set if “fast_dot1x_timers” knob is enabled. ‘quite-period’  and ‘transmit-period’ are set to half the value of auth_servers_timeout. ‘supplicant-timeout’ is also set when setting auth_servers_timeout and is set to default value of 10. | [optional] [default to 5]
**Band** | **string** | &#x60;band&#x60; is deprecated and kept for backward compability. Use bands instead | [optional] [default to null]
**BandSteer** | **bool** | whether to enable band_steering, this works only when band&#x3D;&#x3D;both | [optional] [default to false]
**BandSteerForceBand5** | **bool** | force dual_band capable client to connect to 5G | [optional] [default to false]
**Bands** | [**[]Dot11Band**](dot11_band.md) | list of radios that the wlan should apply to. | [optional] [default to null]
**BlockBlacklistClients** | **bool** | whether to block the clients in the blacklist (up to first 256 macs) | [optional] [default to false]
**Bonjour** | [***Object**](.md) |  | [optional] [default to null]
**CiscoCwa** | [***Object**](.md) |  | [optional] [default to null]
**ClientLimitDown** | **int32** | kbps | [optional] [default to null]
**ClientLimitDownEnabled** | **bool** | if downlink limiting per-client is enabled | [optional] [default to false]
**ClientLimitUp** | **int32** | kbps | [optional] [default to null]
**ClientLimitUpEnabled** | **bool** | if uplink limiting per-client is enabled | [optional] [default to false]
**CoaServers** | [**[]CoaServer**](coa_server.md) | list of COA (change of authorization) servers, optional | [optional] [default to null]
**CreatedTime** | **float64** |  | [optional] [default to null]
**Disable11ax** | **bool** | some old WLAN drivers may not be compatible | [optional] [default to false]
**DisableHtVhtRates** | **bool** | to disable ht or vht rates | [optional] [default to false]
**DisableUapsd** | **bool** | whether to disable U-APSD | [optional] [default to false]
**DisableV1RoamNotify** | **bool** | disable sending v2 roam notification messages | [optional] [default to false]
**DisableV2RoamNotify** | **bool** | disable sending v2 roam notification messages | [optional] [default to false]
**DisableWmm** | **bool** | whether to disable WMM | [optional] [default to false]
**DnsServerRewrite** | [***Object**](.md) |  | [optional] [default to null]
**Dtim** | **int32** |  | [optional] [default to 2]
**DynamicPsk** | [***Object**](.md) |  | [optional] [default to null]
**DynamicVlan** | [***Object**](.md) |  | [optional] [default to null]
**EnableLocalKeycaching** | **bool** | enable AP-AP keycaching via multicast | [optional] [default to false]
**EnableWirelessBridging** | **bool** | by default, we&#x27;d inspect all DHCP packets and drop those unrelated to the wireless client itself in the case where client is a wireless bridge (DHCP packets for other MACs will need to be orwarded), wireless_bridging can be enabled | [optional] [default to false]
**EnableWirelessBridgingDhcpTracking** | **bool** | if the client bridge is doing DHCP on behalf of other devices (L2-NAT), enable dhcp_tracking will cut down DHCP response packets to be forwarded to wireless | [optional] [default to false]
**Enabled** | **bool** | if this wlan is enabled | [optional] [default to true]
**FastDot1xTimers** | **bool** | if set to true, sets default fast-timers with values calculated from ‘auth_servers_timeout’ and ‘auth_server_retries’ . | [optional] [default to false]
**ForSite** | **bool** |  | [optional] [default to null]
**HideSsid** | **bool** | whether to hide SSID in beacon | [optional] [default to false]
**HostnameIe** | **bool** | include hostname inside IE in AP beacons / probe responses | [optional] [default to false]
**Hotspot20** | [***Object**](.md) |  | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**InjectDhcpOption82** | [***WlanInjectDhcpOption82**](wlan_inject_dhcp_option_82.md) |  | [optional] [default to null]
**Interface_** | [***Object**](.md) |  | [optional] [default to null]
**Isolation** | **bool** | whether to stop clients to talk to each other | [optional] [default to false]
**L2Isolation** | **bool** | if isolation is enabled, whether to deny clients to talk to L2 on the LAN | [optional] [default to false]
**LegacyOverds** | **bool** | legacy devices requires the Over-DS (for Fast BSS Transition) bit set (while our chip doesn’t support it). Warning! Enabling this will cause problem for iOS devices. | [optional] [default to false]
**LimitBcast** | **bool** | whether to limit broadcast packets going to wireless (i.e. only allow certain bcast packets to go through) | [optional] [default to false]
**LimitProbeResponse** | **bool** | limit probe response base on some heuristic rules | [optional] [default to false]
**MaxIdletime** | **int32** | max idle time in seconds | [optional] [default to 1800]
**MaxNumClients** | **int32** | maximum number of client connected to the SSID. &#x60;0&#x60; means unlimited | [optional] [default to 0]
**MistNac** | [***WlanMistNac**](wlan_mist_nac.md) |  | [optional] [default to null]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**MspId** | **string** |  | [optional] [default to null]
**MxtunnelId** | **string** | (deprecated, use mxtunnel_ids instead) when &#x60;interface&#x60;&#x3D;&#x3D;&#x60;mxtunnel&#x60;, id of the Mist Tunnel | [optional] [default to null]
**MxtunnelIds** | **[]string** | when &#x60;interface&#x60;&#x3D;&#x60;mxtunnel&#x60;, id of the Mist Tunnel | [optional] [default to null]
**MxtunnelName** | **[]string** | when &#x60;interface&#x60;&#x3D;&#x60;site_medge&#x60;, name of the mxtunnel that in mxtunnels under Site Setting | [optional] [default to null]
**NoStaticDns** | **bool** | whether to only allow client to use DNS that we’ve learned from DHCP response | [optional] [default to false]
**NoStaticIp** | **bool** | whether to only allow client that we’ve learned from DHCP exchange to talk | [optional] [default to false]
**OrgId** | **string** |  | [optional] [default to null]
**Portal** | [***Object**](.md) |  | [optional] [default to null]
**PortalAllowedHostnames** | **[]string** | list of hostnames without http(s):// (matched by substring) | [optional] [default to []]
**PortalAllowedSubnets** | **[]string** | list of CIDRs | [optional] [default to []]
**PortalApiSecret** | **string** | api secret (auto-generated) that can be used to sign guest authorization requests | [optional] [default to null]
**PortalDeniedHostnames** | **[]string** | list of hostnames without http(s):// (matched by substring), this takes precedence over portal_allowed_hostnames | [optional] [default to []]
**PortalImage** | **string** | Url of portal background image | [optional] [default to null]
**PortalSsoUrl** | **string** |  | [optional] [default to null]
**PortalTemplateUrl** | **string** | N.B portal_template will be forked out of wlan objects soon. To fetch portal_template, please query portal_template_url. To update portal_template, use Wlan Portal Template. | [optional] [default to null]
**Qos** | [***WlanQos**](wlan_qos.md) |  | [optional] [default to null]
**Radsec** | [***Object**](.md) |  | [optional] [default to null]
**Rateset** | [**map[string]WlanDatarates**](wlan_datarates.md) |  | [optional] [default to null]
**RoamMode** | [***Object**](.md) |  | [optional] [default to null]
**Schedule** | [***Object**](.md) |  | [optional] [default to null]
**SiteId** | **string** |  | [optional] [default to null]
**SleExcluded** | **bool** | whether to exclude this WLAN from SLE metrics | [optional] [default to false]
**Ssid** | **string** | the name of the SSID | [default to null]
**TemplateId** | **string** |  | [optional] [default to null]
**Thumbnail** | **string** | Url of portal background image thumbnail | [optional] [default to null]
**UseEapolV1** | **bool** | if &#x60;auth.type&#x60;&#x3D;&#x3D;’eap’ or ‘psk’, should only be set for legacy client, such as pre-2004, 802.11b devices | [optional] [default to false]
**VlanEnabled** | **bool** | if vlan tagging is enabled | [optional] [default to false]
**VlanId** | [***Object**](.md) |  | [optional] [default to null]
**VlanIds** | [**[]VlanIdWithVariable**](vlan_id_with_variable.md) | vlan_ids to use when there’s no match from RA | [optional] [default to null]
**VlanPooling** | **bool** | vlan pooling allows AP to place client on different VLAN using a deterministic algorithm | [optional] [default to false]
**WlanLimitDown** | **int32** | kbps | [optional] [default to null]
**WlanLimitDownEnabled** | **bool** | if downlink limiting for whole wlan is enabled | [optional] [default to false]
**WlanLimitUp** | **int32** | kbps | [optional] [default to null]
**WlanLimitUpEnabled** | **bool** | if uplink limiting for whole wlan is enabled | [optional] [default to false]
**WxtagIds** | **[]string** | list of wxtag_ids | [optional] [default to null]
**WxtunnelId** | **string** | when &#x60;interface&#x60;&#x3D;&#x60;wxtunnel&#x60;, id of the WXLAN Tunnel | [optional] [default to null]
**WxtunnelRemoteId** | **string** | when &#x60;interface&#x60;&#x3D;&#x60;wxtunnel&#x60;, remote tunnel identifier | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

