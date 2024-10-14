# OspfAreasNetwork

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthKeys** | **map[string]string** | Required if &#x60;auth_type&#x60;&#x3D;&#x3D;&#x60;md5&#x60;. Property key is the key number | [optional] [default to null]
**AuthPassword** | **string** | Required if &#x60;auth_type&#x60;&#x3D;&#x3D;&#x60;password&#x60;, the password, max length is 8 | [optional] [default to null]
**AuthType** | [***AllOfospfAreasNetworkAuthType**](AllOfospfAreasNetworkAuthType.md) |  | [optional] [default to null]
**BfdMinimumInterval** | **int32** |  | [optional] [default to null]
**DeadInterval** | **int32** |  | [optional] [default to null]
**ExportPolicy** | **string** |  | [optional] [default to null]
**HelloInterval** | **int32** |  | [optional] [default to null]
**ImportPolicy** | **string** |  | [optional] [default to null]
**InterfaceType** | [***AllOfospfAreasNetworkInterfaceType**](AllOfospfAreasNetworkInterfaceType.md) |  | [optional] [default to null]
**Metric** | **int32** |  | [optional] [default to null]
**NoReadvertiseToOverlay** | **bool** | by default, we&#x27;ll re-advertise all learned OSPF routes toward overlay | [optional] [default to false]
**Passive** | **bool** | whether to send OSPF-Hello | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

