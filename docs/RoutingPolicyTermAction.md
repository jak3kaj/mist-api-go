# RoutingPolicyTermAction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accept** | **bool** |  | [optional] [default to null]
**AddCommunity** | **[]string** |  | [optional] [default to null]
**AddTargetVrfs** | **[]string** | for SSR, hub decides how VRF routes are leaked on spoke | [optional] [default to null]
**Community** | **[]string** | when used as export policy, optional | [optional] [default to null]
**ExcludeAsPath** | **[]string** | when used as export policy, optional. To exclude certain AS | [optional] [default to null]
**ExcludeCommunity** | **[]string** |  | [optional] [default to null]
**ExportCommunitites** | **[]string** | when used as export policy, optional | [optional] [default to null]
**LocalPreference** | **string** | optional, for an import policy, local_preference can be changed | [optional] [default to null]
**PrependAsPath** | **[]string** | when used as export policy, optional. By default, the local AS will be prepended, to change it | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

