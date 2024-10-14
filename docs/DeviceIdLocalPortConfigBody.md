# DeviceIdLocalPortConfigBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AeDisableLacp** | **bool** | To disable LACP support for the AE interface | [optional] [default to null]
**AeIdx** | **int32** | Users could force to use the designated AE name | [optional] [default to null]
**AeLacpSlow** | **bool** | to use fast timeout | [optional] [default to true]
**Aggregated** | **bool** |  | [optional] [default to false]
**Critical** | **bool** | if want to generate port up/down alarm | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**DisableAutoneg** | **bool** | if &#x60;speed&#x60; and &#x60;duplex&#x60; are specified, whether to disable autonegotiation | [optional] [default to false]
**Duplex** | [***Object**](.md) |  | [optional] [default to null]
**DynamicUsage** | **string** | Enable dynamic usage for this port. Set to &#x60;dynamic&#x60; to enable. | [optional] [default to null]
**Esilag** | **bool** |  | [optional] [default to null]
**Mtu** | **int32** | media maximum transmission unit (MTU) is the largest data unit that can be forwarded without fragmentation | [optional] [default to 1514]
**NoLocalOverwrite** | **bool** | prevent helpdesk to override the port config | [optional] [default to null]
**PoeDisabled** | **bool** |  | [optional] [default to false]
**Speed** | [***Object**](.md) |  | [optional] [default to null]
**Usage** | **string** | port usage name.   If EVPN is used, use &#x60;evpn_uplink&#x60;or &#x60;evpn_downlink&#x60; | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

