# SwitchNetwork

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Isolation** | **bool** | whether to stop clients to talk to each other, default is false (when enabled, a unique isolation_vlan_id is required) NOTE: this features requires uplink device to also a be Juniper device and &#x60;inter_switch_link&#x60; to be set | [optional] [default to false]
**IsolationVlanId** | **string** |  | [optional] [default to null]
**Subnet** | **string** | optional for pure switching, required when L3 / routing features are used | [optional] [default to null]
**VlanId** | [***AllOfswitchNetworkVlanId**](AllOfswitchNetworkVlanId.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

