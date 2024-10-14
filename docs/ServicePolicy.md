# ServicePolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [***AllOfservicePolicyAction**](AllOfservicePolicyAction.md) |  | [optional] [default to null]
**Appqoe** | [***AllOfservicePolicyAppqoe**](AllOfservicePolicyAppqoe.md) |  | [optional] [default to null]
**Ewf** | [**[]ServicePolicyEwfRule**](service_policy_ewf_rule.md) |  | [optional] [default to null]
**Idp** | [***IdpConfig**](idp_config.md) |  | [optional] [default to null]
**LocalRouting** | **bool** | access within the same VRF | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**PathPreference** | **string** | by default, we derive all paths available and use them optionally, you can customize by using &#x60;path_preference&#x60; | [optional] [default to null]
**ServicepolicyId** | **string** | used to link servicepolicy defined at org level and overwrite some attributes | [optional] [default to null]
**Services** | **[]string** |  | [optional] [default to null]
**Tenants** | **[]string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

