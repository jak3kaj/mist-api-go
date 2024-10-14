# WlanAuth

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnticlogThreshold** | **int32** | SAE anti-clogging token threshold | [optional] [default to 16]
**EapReauth** | **bool** | whether to trigger EAP reauth when the session ends | [optional] [default to false]
**EnableMacAuth** | **bool** | whether to enable MAC Auth, uses the same auth_servers | [optional] [default to false]
**KeyIdx** | **int32** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;wep&#x60; | [optional] [default to 1]
**Keys** | **[]string** | when type&#x3D;wep, four 10-character or 26-character hex string, null can be used. All keys, if provided, have to be in the same length | [optional] [default to []]
**MultiPskOnly** | **bool** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;psk&#x60;, whether to only use multi_psk | [optional] [default to false]
**Owe** | [***AllOfwlanAuthOwe**](AllOfwlanAuthOwe.md) |  | [optional] [default to null]
**Pairwise** | [**[]OneOfwlanAuthPairwiseItems**](.md) | when &#x60;type&#x60;&#x3D;&#x60;psk&#x60; or &#x60;type&#x60;&#x3D;&#x60;eap&#x60;, one or more of &#x60;wpa1-ccmp&#x60;, &#x60;wpa1-tkip&#x60;, &#x60;wpa2-ccmp&#x60;, &#x60;wpa2-tkip&#x60;, &#x60;wpa3&#x60; | [optional] [default to null]
**PrivateWlan** | **bool** | when &#x60;multi_psk_only&#x60;&#x3D;&#x3D;&#x60;true&#x60;, whether private wlan is enabled | [optional] [default to false]
**Psk** | **string** | when &#x60;type&#x60;&#x3D;&#x3D;&#x60;psk&#x60;, 8-64 characters, or 64 hex characters | [optional] [default to null]
**Type_** | [***AllOfwlanAuthType_**](AllOfwlanAuthType_.md) |  | [default to null]
**WepAsSecondaryAuth** | **bool** | enable WEP as secondary auth | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

