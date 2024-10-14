# MspLicenseAction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmendmentId** | **string** | required if &#x60;op&#x60;&#x3D;&#x3D;&#x60;unamend&#x60; | [optional] [default to null]
**DstOrgId** | **string** | required if &#x60;op&#x60;&#x3D;&#x3D;&#x60;amend&#x60;, destination org id | [optional] [default to null]
**Notes** | **string** | required if &#x60;op&#x60;&#x3D;&#x3D; &#x60;annotate&#x60; | [optional] [default to null]
**Op** | [***AllOfmspLicenseActionOp**](AllOfmspLicenseActionOp.md) |  | [default to null]
**Quantity** | **float64** | required if &#x60;op&#x60;&#x3D;&#x3D;&#x60;amend&#x60; | [optional] [default to null]
**SubscriptionId** | **string** | required if &#x60;op&#x60;&#x3D;&#x3D; &#x60;annotate&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

