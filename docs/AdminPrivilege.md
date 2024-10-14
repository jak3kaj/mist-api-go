# AdminPrivilege

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MspId** | **string** |  | [optional] [default to null]
**MspLogoUrl** | **string** | logo of the MSP (if the MSP belongs to an Advanced tier) | [optional] [default to null]
**MspName** | **string** | name of the MSP (if the org belongs to an MSP) | [optional] [default to null]
**MspUrl** | **string** | custom url of the MSP (if the MSP belongs to an Advanced tier) | [optional] [default to null]
**Name** | **string** | name of the org/site/MSP depending on object scope | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**OrgName** | **string** | name of the org (for a site belonging to org) | [optional] [default to null]
**OrggroupIds** | **[]string** | if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;orggroup&#x60; | [optional] [default to null]
**Role** | [***AllOfadminPrivilegeRole**](AllOfadminPrivilegeRole.md) |  | [default to null]
**Scope** | [***AllOfadminPrivilegeScope**](AllOfadminPrivilegeScope.md) |  | [default to null]
**SiteId** | **string** |  | [optional] [default to null]
**SitegroupIds** | **[]string** |  | [optional] [default to null]
**Views** | [***AllOfadminPrivilegeViews**](AllOfadminPrivilegeViews.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

