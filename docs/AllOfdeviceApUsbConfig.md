# AllOfdeviceApUsbConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cacert** | **string** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;imagotag&#x60; | [optional] [default to null]
**Channel** | **int32** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;imagotag&#x60;, channel selection, not needed by default, required for manual channel override only | [optional] [default to null]
**Enabled** | **bool** | whether to enable any usb config | [optional] [default to null]
**Host** | **string** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;imagotag&#x60; | [optional] [default to null]
**Port** | **int32** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;imagotag&#x60; | [optional] [default to 0]
**Type_** | [***Object**](.md) |  | [optional] [default to null]
**VerifyCert** | **bool** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;imagotag&#x60;, whether to turn on SSL verification | [optional] [default to null]
**VlanId** | **int32** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;solum&#x60; or &#x60;type&#x60;&#x3D;&#x3D;&#x60;hanshow&#x60; | [optional] [default to 1]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

