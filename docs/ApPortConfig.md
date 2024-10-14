# ApPortConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | **bool** |  | [optional] [default to false]
**DynamicVlan** | [***AllOfapPortConfigDynamicVlan**](AllOfapPortConfigDynamicVlan.md) |  | [optional] [default to null]
**EnableMacAuth** | **bool** |  | [optional] [default to false]
**Forwarding** | [***AllOfapPortConfigForwarding**](AllOfapPortConfigForwarding.md) |  | [optional] [default to null]
**MacAuthPreferred** | **bool** | when &#x60;true&#x60;, we&#x27;ll do dot1x then mac_auth. enable this to prefer mac_auth | [optional] [default to false]
**MacAuthProtocol** | [***AllOfapPortConfigMacAuthProtocol**](AllOfapPortConfigMacAuthProtocol.md) |  | [optional] [default to null]
**MistNac** | [***WlanMistNac**](wlan_mist_nac.md) |  | [optional] [default to null]
**MxTunnelId** | **string** | if &#x60;forwarding&#x60;&#x3D;&#x3D;&#x60;mxtunnel&#x60;, vlan_ids comes from mxtunnel | [optional] [default to null]
**MxtunnelName** | **string** | if &#x60;forwarding&#x60;&#x3D;&#x3D;&#x60;site_mxedge&#x60;, vlan_ids comes from site_mxedge (&#x60;mxtunnels&#x60; under site setting) | [optional] [default to null]
**PortAuth** | [***AllOfapPortConfigPortAuth**](AllOfapPortConfigPortAuth.md) |  | [optional] [default to null]
**PortVlanId** | **int32** | if &#x60;forwrding&#x60;&#x3D;&#x3D;&#x60;limited&#x60; | [optional] [default to null]
**RadiusConfig** | [***AllOfapPortConfigRadiusConfig**](AllOfapPortConfigRadiusConfig.md) |  | [optional] [default to null]
**Radsec** | [***AllOfapPortConfigRadsec**](AllOfapPortConfigRadsec.md) |  | [optional] [default to null]
**VlanId** | **int32** | optional to specify the vlan id for a tunnel if forwarding is for &#x60;wxtunnel&#x60;, &#x60;mxtunnel&#x60; or &#x60;site_mxedge&#x60;.   * if vlan_id is not specified then it will use first one in vlan_ids[] of the mxtunnel.   * if forwarding &#x3D;&#x3D; site_mxedge, vlan_ids comes from site_mxedge (&#x60;mxtunnels&#x60; under site setting) | [optional] [default to null]
**VlandIds** | **[]int32** | if &#x60;forwrding&#x60;&#x3D;&#x3D;&#x60;limited&#x60; | [optional] [default to null]
**WxtunnelId** | **string** | if &#x60;forwarding&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60;, the port is bridged to the vlan of the session | [optional] [default to null]
**WxtunnelRemoteId** | **string** | if &#x60;forwarding&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60;, the port is bridged to the vlan of the session | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

