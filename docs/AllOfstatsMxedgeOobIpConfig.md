# AllOfstatsMxedgeOobIpConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autoconf6** | **bool** |  | [optional] [default to true]
**Dhcp6** | **bool** |  | [optional] [default to true]
**Dns** | **[]string** | IPv4 ignored if &#x60;type&#x60;!&#x3D;&#x60;static&#x60; IPv6 ignored if &#x60;type6&#x60;!&#x3D;&#x60;static&#x60; | [optional] [default to ["8.8.8.8","8.8.4.4","2001:4860:4860::8888","2001:4860:4860::8844"]]
**Gateway** | **string** | if &#x60;type&#x60;&#x3D;&#x60;static&#x60; | [optional] [default to null]
**Gateway6** | **string** |  | [optional] [default to null]
**Ip** | **string** | if &#x60;type&#x60;&#x3D;&#x60;static&#x60; | [optional] [default to null]
**Ip6** | **string** |  | [optional] [default to null]
**Netmask** | **string** | if &#x60;type&#x60;&#x3D;&#x60;static&#x60; | [optional] [default to null]
**Netmask6** | **string** |  | [optional] [default to null]
**Type_** | [***Object**](.md) |  | [optional] [default to null]
**Type6** | [***Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

