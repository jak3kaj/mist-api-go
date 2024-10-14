# SwitchOobIpConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | **string** |  | [optional] [default to null]
**Ip** | **string** |  | [optional] [default to null]
**Netmask** | **string** | used only if &#x60;subnet&#x60; is not specified in &#x60;networks&#x60; | [optional] [default to null]
**Network** | **string** | optional, the network to be used for mgmt | [optional] [default to null]
**Type_** | [***AllOfswitchOobIpConfigType_**](AllOfswitchOobIpConfigType_.md) |  | [optional] [default to null]
**UseMgmtVrf** | **bool** | f supported on the platform. If enabled, DNS will be using this routing-instance, too | [optional] [default to false]
**UseMgmtVrfForHostOut** | **bool** | for host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

