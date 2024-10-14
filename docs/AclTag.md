# AclTag

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GbpTag** | **int32** | required if - &#x60;type&#x60;&#x3D;&#x3D;&#x60;dynamic_gbp&#x60; (gbp_tag received from RADIUS) - &#x60;type&#x60;&#x3D;&#x3D;&#x60;static_gbp&#x60; (applying gbp tag against matching conditions) | [optional] [default to null]
**Macs** | **[]string** | required if  - &#x60;type&#x60;&#x3D;&#x3D;&#x60;mac&#x60; - &#x60;type&#x60;&#x3D;&#x3D;&#x60;static_gbp&#x60; if from matching mac | [optional] [default to null]
**Network** | **string** | if:   * &#x60;type&#x60;&#x3D;&#x3D;&#x60;mac&#x60; (optional. default is &#x60;any&#x60;)   * &#x60;type&#x60;&#x3D;&#x3D;&#x60;subnet&#x60; (optional. default is &#x60;any&#x60;)   * &#x60;type&#x60;&#x3D;&#x3D;&#x60;network&#x60;   * &#x60;type&#x60;&#x3D;&#x3D;&#x60;resource&#x60; (optional. default is &#x60;any&#x60;)   * &#x60;type&#x60;&#x3D;&#x3D;&#x60;static_gbp&#x60; if from matching network (vlan)&#x27; | [optional] [default to null]
**RadiusGroup** | **string** | required if:   * &#x60;type&#x60;&#x3D;&#x3D;&#x60;radius_group&#x60;   * &#x60;type&#x60;&#x3D;&#x3D;&#x60;static_gbp&#x60; if from matching radius_group | [optional] [default to null]
**Specs** | [**[]AclTagSpec**](acl_tag_spec.md) | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;resource&#x60; empty means unrestricted, i.e. any | [optional] [default to null]
**Subnets** | **[]string** | if  - &#x60;type&#x60;&#x3D;&#x3D;&#x60;subnet&#x60;  - &#x60;type&#x60;&#x3D;&#x3D;&#x60;resource&#x60; (optional. default is &#x60;any&#x60;) - &#x60;type&#x60;&#x3D;&#x3D;&#x60;static_gbp&#x60; if from matching subnet | [optional] [default to null]
**Type_** | [***AllOfaclTagType_**](AllOfaclTagType_.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

