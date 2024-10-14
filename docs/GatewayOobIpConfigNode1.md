# GatewayOobIpConfigNode1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;static&#x60; | [optional] [default to null]
**Ip** | **string** |  | [optional] [default to null]
**Netmask** | **string** | used only if &#x60;subnet&#x60; is not specified in &#x60;networks&#x60; | [optional] [default to null]
**Type_** | [***AllOfgatewayOobIpConfigNode1Type_**](AllOfgatewayOobIpConfigNode1Type_.md) |  | [optional] [default to null]
**UseMgmtVrf** | **bool** | if supported on the platform. If enabled, DNS will be using this routing-instance, too | [optional] [default to false]
**UseMgmtVrfForHostOut** | **bool** | whether to use &#x60;mgmt_junos&#x60; for host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired | [optional] [default to false]
**VlanId** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

