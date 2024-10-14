# GatewayPortVpnPath

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BfdProfile** | [***AllOfgatewayPortVpnPathBfdProfile**](AllOfgatewayPortVpnPathBfdProfile.md) |  | [optional] [default to null]
**BfdUseTunnelMode** | **bool** | whether to use tunnel mode. SSR only | [optional] [default to false]
**Preference** | **int32** | for a given VPN, when &#x60;path_selection.strategy&#x60;&#x3D;&#x3D;&#x60;simple&#x60;, the preference for a path (lower is preferred) | [optional] [default to null]
**Role** | [***AllOfgatewayPortVpnPathRole**](AllOfgatewayPortVpnPathRole.md) |  | [optional] [default to null]
**TrafficShaping** | [***GatewayTrafficShaping**](gateway_traffic_shaping.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

