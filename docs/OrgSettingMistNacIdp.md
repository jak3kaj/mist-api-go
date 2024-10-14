# OrgSettingMistNacIdp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeRealms** | **[]string** | when the IDP of mxedge_proxy type, exclude the following realms from proxying in addition to other valid home realms in this org | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**UserRealms** | **[]string** | which realm should trigger this IDP. User Realm is extracted from:   * Username-AVP (&#x60;mist.com&#x60; from john@mist.com)   * Cert CN | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

