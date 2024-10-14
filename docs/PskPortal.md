# PskPortal

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | [***AllOfpskPortalAuth**](AllOfpskPortalAuth.md) |  | [optional] [default to null]
**BgImageUrl** | **string** |  | [optional] [default to null]
**CleanupPsk** | **bool** | used to cleanup exited psk when portal delete or ssid changed | [optional] [default to false]
**CreatedTime** | **float64** |  | [optional] [default to null]
**ExpireTime** | **int32** | unit min | [optional] [default to null]
**ExpiryNotificationTime** | **int32** | Number of days before psk is expired. Used as to when to start sending reminder notification when the psk is about to expire | [optional] [default to null]
**HidePsksCreatedByOtherAdmins** | **bool** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;admin&#x60; | [optional] [default to false]
**Id** | **string** |  | [optional] [default to null]
**MaxUsage** | **int32** | &#x60;max_usage&#x60;&#x3D;&#x3D;&#x60;0&#x60; means unlimited | [optional] [default to 0]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**NotificationRenewUrl** | **string** | optional, will include the link in the notification email the customer can either provide their own url or use the one generate from mist, or do a url shorterner against either | [optional] [default to null]
**NotifyExpiry** | **bool** | If set to true, reminder notification will be sent when psk is about to expire | [optional] [default to null]
**NotifyOnCreateOrEdit** | **bool** |  | [optional] [default to false]
**OrgId** | **string** |  | [optional] [default to null]
**PassphraseRules** | [***PskPortalPassphraseRules**](psk_portal_passphrase_rules.md) |  | [optional] [default to null]
**RequiredFields** | **[]string** | what information to ask for (email is required by default) | [optional] [default to null]
**Role** | **string** |  | [optional] [default to null]
**Ssid** | **string** | intended SSID | [default to null]
**Sso** | [***AllOfpskPortalSso**](AllOfpskPortalSso.md) |  | [optional] [default to null]
**TemplateUrl** | **string** | UI customization | [optional] [default to null]
**ThumbnailUrl** | **string** |  | [optional] [default to null]
**Type_** | [***AllOfpskPortalType_**](AllOfpskPortalType_.md) |  | [optional] [default to null]
**VlanId** | [***AllOfpskPortalVlanId**](AllOfpskPortalVlanId.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

