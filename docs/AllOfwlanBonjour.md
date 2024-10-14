# AllOfwlanBonjour

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalVlanIds** | **string** | comma sperated list of additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses | [default to null]
**Enabled** | **bool** | whether to enable bonjour for this WLAN. Once enabled, limit_bcast is assumed true, allow_mdns is assumed false | [optional] [default to false]
**Services** | [**map[string]WlanBonjourServiceProperties**](wlan_bonjour_service_properties.md) | what services are allowed.  Property key is the service name | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

