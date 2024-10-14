# VirtualChassisMemberUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | **string** | Required if &#x60;op&#x60;&#x3D;&#x3D;&#x60;add&#x60; or &#x60;op&#x60;&#x3D;&#x3D;&#x60;preprovision&#x60;. | [optional] [default to null]
**Member** | **int32** | Required if &#x60;op&#x60;&#x3D;&#x3D;&#x60;remove&#x60; or &#x60;op&#x60;&#x3D;&#x3D;&#x60;preprovision&#x60;. Optional if &#x60;op&#x60;&#x3D;&#x3D;&#x60;add&#x60; | [optional] [default to null]
**VcPorts** | **[]string** | Required if &#x60;op&#x60;&#x3D;&#x3D;&#x60;add&#x60; or &#x60;op&#x60;&#x3D;&#x3D;&#x60;preprovision&#x60; | [optional] [default to null]
**VcRole** | [***AllOfvirtualChassisMemberUpdateVcRole**](AllOfvirtualChassisMemberUpdateVcRole.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

