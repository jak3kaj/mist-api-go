# VrrpGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthKey** | **string** | if &#x60;auth_type&#x60;&#x3D;&#x3D;&#x60;md5&#x60; | [optional] [default to null]
**AuthPassword** | **string** | if &#x60;auth_type&#x60;&#x3D;&#x3D;&#x60;simple&#x60; | [optional] [default to null]
**AuthType** | [***AllOfvrrpGroupAuthType**](AllOfvrrpGroupAuthType.md) |  | [optional] [default to null]
**Networks** | [**map[string]VrrpGroupNetwork**](vrrp_group_network.md) | Property key is the network name | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

