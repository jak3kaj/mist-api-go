# SiteIdWxrulesBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [***Object**](.md) |  | [optional] [default to null]
**ApplyTags** | **[]string** |  | [optional] [default to null]
**BlockedApps** | **[]string** | blocked apps (always blocking, ignoring action), the key of Get Application List | [optional] [default to null]
**CreatedTime** | **float64** |  | [optional] [default to null]
**DstAllowWxtags** | **[]string** | List of WxTag UUID to indicate these tags are allowed access | [default to null]
**DstDenyWxtags** | **[]string** | List of WxTag UUID to indicate these tags are blocked access | [default to null]
**DstWxtags** | **[]string** | List of WxTag UUID | [optional] [default to null]
**Enabled** | **bool** |  | [optional] [default to true]
**ForSite** | **bool** |  | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**Order** | **int32** | the order how rules would be looked up, &gt; 0 and bigger order got matched first, -1 means LAST, uniqueness not checked | [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**SiteId** | **string** |  | [optional] [default to null]
**SrcWxtags** | **[]string** | List of WxTag UUID to determine if this rule would match | [default to null]
**TemplateId** | **string** | Only for Org Level WxRule | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

