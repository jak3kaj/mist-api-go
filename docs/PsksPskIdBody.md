# PsksPskIdBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminSsoId** | **string** | sso id for psk created from psk portal | [optional] [default to null]
**CreatedTime** | **float64** |  | [optional] [default to null]
**Email** | **string** | email to send psk expiring notifications to | [optional] [default to null]
**ExpireTime** | **int32** | Expire time for this PSK key (epoch time in seconds). Default &#x60;null&#x60; (as no expiration) | [optional] [default to null]
**ExpiryNotificationTime** | **int32** | Number of days before psk is expired. Used as to when to start sending reminder notification when the psk is about to expire | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**Mac** | **string** | if &#x60;usage&#x60;&#x3D;&#x3D;&#x60;single&#x60;, the mac that this PSK ties to, empty if &#x60;auto-binding&#x60; | [optional] [default to null]
**Macs** | **[]string** | if &#x60;usage&#x60;&#x3D;&#x3D;&#x60;macs&#x60;, this list contains N number of client mac addresses or mac patterns(11:22:*) or both. This list is capped at 5000 | [optional] [default to null]
**MaxUsage** | **int32** | For Org PSK Only. Max concurrent users for this PSK key. Default is 0 (unlimited) | [optional] [default to 0]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Note** | **string** |  | [optional] [default to null]
**NotifyExpiry** | **bool** | If set to true, reminder notification will be sent when psk is about to expire | [optional] [default to false]
**NotifyOnCreateOrEdit** | **bool** | If set to true, notification will be sent when psk is created or edited | [optional] [default to null]
**OldPassphrase** | **string** | previous passphrase of the PSK if it has been rotated | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**Passphrase** | **string** | passphrase of the PSK (8-63 character or 64 in hex) | [default to null]
**Role** | **string** |  | [optional] [default to null]
**SiteId** | **string** |  | [optional] [default to null]
**Ssid** | **string** | SSID this PSK should be applicable to | [default to null]
**Usage** | [***Object**](.md) |  | [optional] [default to null]
**VlanId** | [***Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

