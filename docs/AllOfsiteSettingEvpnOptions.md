# AllOfsiteSettingEvpnOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoLoopbackSubnet** | **string** | optional, for dhcp_relay, unique loopback IPs are required for ERB or IPClos where we can set option-82 server_id-overrides | [optional] [default to null]
**AutoLoopbackSubnet6** | **string** | optional, for dhcp_relay, unique loopback IPs are required for ERB or IPClos where we can set option-82 server_id-overrides | [optional] [default to null]
**AutoRouterIdSubnet** | **string** | optional, this generates router_id automatically, if specified, &#x60;router_id_prefix&#x60; is ignored | [optional] [default to null]
**AutoRouterIdSubnet6** | **string** | optional, this generates router_id automatically, if specified, &#x60;router_id_prefix&#x60; is ignored | [optional] [default to fd31:5700:1::/64]
**CoreAsBorder** | **bool** | optional, for ERB or CLOS, you can either use esilag to upstream routers or to also be the virtual-gateway when &#x60;routed_at&#x60; !&#x3D; &#x60;core&#x60;, whether to do virtual-gateway at core as well | [optional] [default to false]
**Overlay** | [***EvpnOptionsOverlay**](evpn_options_overlay.md) |  | [optional] [default to null]
**PerVlanVgaV4Mac** | **bool** | by default, JUNOS uses 00-00-5e-00-01-01 as the virtual-gateway-address&#x27;s v4_mac if enabled, 00-00-5e-00-XX-YY will be used (where XX&#x3D;vlan_id/256, YY&#x3D;vlan_id%256) | [optional] [default to false]
**RoutedAt** | [***Object**](.md) |  | [optional] [default to null]
**Underlay** | [***EvpnOptionsUnderlay**](evpn_options_underlay.md) |  | [optional] [default to null]
**VsInstances** | [**map[string]EvpnOptionsVsInstance**](evpn_options_vs_instance.md) | optional, for EX9200 only to seggregate virtual-switches | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

