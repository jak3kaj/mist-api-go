# AllOfdeviceprofileApIpConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns** | **[]string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;static&#x60; | [optional] [default to null]
**DnsSuffix** | **[]string** | required if &#x60;type&#x60;&#x3D;&#x3D;&#x60;static&#x60; | [optional] [default to null]
**Gateway** | **string** | required if &#x60;type&#x60;&#x3D;&#x3D;&#x60;static&#x60; | [optional] [default to null]
**Gateway6** | **string** |  | [optional] [default to null]
**Ip** | **string** | required if &#x60;type&#x60;&#x3D;&#x3D;&#x60;static&#x60; | [optional] [default to null]
**Ip6** | **string** |  | [optional] [default to null]
**Mtu** | **int32** |  | [optional] [default to null]
**Netmask** | **string** | required if &#x60;type&#x60;&#x3D;&#x3D;&#x60;static&#x60; | [optional] [default to null]
**Netmask6** | **string** |  | [optional] [default to null]
**Type_** | [***Object**](.md) |  | [optional] [default to null]
**Type6** | [***Object**](.md) |  | [optional] [default to null]
**VlanId** | **int32** | management vlan id, default is 1 (untagged) | [optional] [default to 1]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

