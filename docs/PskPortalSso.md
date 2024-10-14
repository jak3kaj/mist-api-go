# PskPortalSso

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedRoles** | **[]string** | // allowed roles for accessing psk portal, if none, any role is permitted | [optional] [default to null]
**IdpCert** | **string** |  | [optional] [default to null]
**IdpSignAlgo** | [***AllOfpskPortalSsoIdpSignAlgo**](AllOfpskPortalSsoIdpSignAlgo.md) |  | [optional] [default to null]
**IdpSsoUrl** | **string** |  | [optional] [default to null]
**Issuer** | **string** |  | [optional] [default to null]
**NameidFormat** | **string** |  | [optional] [default to null]
**RoleMapping** | **map[string]string** | Property key is the role name, property value is the SSO Attribute | [optional] [default to null]
**UseSsoRoleForPskRole** | **bool** | if enabled, the &#x60;role&#x60; above will be ignored | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

