# EvpnTopologySwitch

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [***AllOfevpnTopologySwitchConfig**](AllOfevpnTopologySwitchConfig.md) |  | [optional] [default to null]
**DeviceprofileId** | **string** |  | [optional] [default to null]
**Downlinks** | **[]string** |  | [optional] [default to null]
**Esilaglinks** | **[]string** |  | [optional] [default to null]
**EvpnId** | **int32** |  | [optional] [default to null]
**Mac** | **string** |  | [optional] [default to null]
**Model** | **string** |  | [optional] [default to null]
**Pod** | **int32** | optionally, for distribution / access / esilag-access, they can be placed into different pods. e.g.    * for CLOS, to group dist / access switches into pods   * for ERB/CRB, to group dist / esilag-access into pods | [optional] [default to 1]
**Pods** | **[]int32** | by default, core switches are assumed to be connecting all pods.  if you want to limit the pods, you can specify pods. | [optional] [default to null]
**Role** | [***AllOfevpnTopologySwitchRole**](AllOfevpnTopologySwitchRole.md) |  | [optional] [default to null]
**SiteId** | **string** |  | [optional] [default to null]
**SuggestedDownlinks** | **[]string** |  | [optional] [default to null]
**SuggestedEsilaglinks** | **[]string** |  | [optional] [default to null]
**SuggestedUplinks** | **[]string** |  | [optional] [default to null]
**Uplinks** | **[]string** | if not specified in the request, suggested ones will be used | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

