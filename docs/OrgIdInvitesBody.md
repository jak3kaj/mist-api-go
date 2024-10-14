# OrgIdInvitesBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminId** | **string** |  | [optional] [default to null]
**ComplianceStatus** | [***Object**](.md) |  | [optional] [default to null]
**Email** | **string** | if admin account is not an Org API Token | [optional] [default to null]
**EnableTwoFactor** | **bool** | if admin account is not an Org API Token | [optional] [default to null]
**ExpireTime** | **int32** |  | [optional] [default to null]
**FirstName** | **string** | if admin account is not an Org API Token for an invite, this is the original first name used | [optional] [default to null]
**Hours** | **int32** | if admin account is not an Org API Token, how long the invite should be valid | [optional] [default to 24]
**LastName** | **string** | if admin account is not an Org API Token for an invite, this is the original last name used | [optional] [default to null]
**Name** | **string** | for Org API Token Only | [optional] [default to null]
**NoTracking** | **bool** | when it doesn’t exist, it’s assumed true on EU (i.e. no tracking, the user has to opt-in); otherwise, the user would have to opt-out | [optional] [default to null]
**OauthGoogle** | **bool** | if admin account is not an Org API Token | [optional] [default to null]
**PasswordModifiedTime** | **float64** | password last modified time, in epoch | [optional] [default to null]
**Phone** | **string** | if admin account is not an Org API Token phone number (numbers only, including country code) | [optional] [default to null]
**Phone2** | **string** | if admin account is not an Org API Token secondary phone number (numbers only, including country code) | [optional] [default to null]
**Privileges** | [**[]AdminPrivilege**](admin_privilege.md) | list of privileges the admin has | [optional] [default to null]
**SessionExpiry** | **int64** |  | [optional] [default to null]
**Tags** | **[]string** |  | [optional] [default to null]
**TwoFactorVerified** | **bool** | if admin account is not an Org API Token two factor status | [optional] [default to null]
**ViaSso** | **bool** | if admin account is not an Org API Token an admin login via_sso is more restircted. (password and email cannot be changed) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

