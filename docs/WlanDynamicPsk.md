# WlanDynamicPsk

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultPsk** | **string** | default PSK to use if cloud WLC is not available, 8-63 characters | [optional] [default to null]
**DefaultVlanId** | [***AllOfwlanDynamicPskDefaultVlanId**](AllOfwlanDynamicPskDefaultVlanId.md) |  | [optional] [default to null]
**Enabled** | **bool** |  | [optional] [default to false]
**ForceLookup** | **bool** | when 11r is enabled, we&#x27;ll try to use the cached PMK, this can be disabled &#x60;false&#x60; means auto | [optional] [default to false]
**Source** | [***AllOfwlanDynamicPskSource**](AllOfwlanDynamicPskSource.md) |  | [optional] [default to null]
**VlanIds** | [**[]VlanIdWithVariable**](vlan_id_with_variable.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

