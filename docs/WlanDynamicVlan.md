# WlanDynamicVlan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultVlanId** | [***AllOfwlanDynamicVlanDefaultVlanId**](AllOfwlanDynamicVlanDefaultVlanId.md) |  | [optional] [default to null]
**DefaultVlanIds** | [**[]WlanDynamicVlanDefaultVlanId**](wlan_dynamic_vlan_default_vlan_id.md) | Default VLAN ID(s) can be a number, a range of VLAN IDs, a variable or multiple numbers, ranges or variables as a VLAN pool. Default VLAN as a pool of VLANS requires 0.14.x or newer firmware | [optional] [default to null]
**Enabled** | **bool** | whether to enable dynamic vlan | [optional] [default to false]
**LocalVlanIds** | [**[]VlanIdWithVariable**](vlan_id_with_variable.md) | vlan_ids to be locally bridged | [optional] [default to null]
**Type_** | [***AllOfwlanDynamicVlanType_**](AllOfwlanDynamicVlanType_.md) |  | [optional] [default to null]
**Vlans** | **map[string]string** | map between vlan_id (as string) to airespace interface names (comma-separated) or null for stndard mapping   * if &#x60;dynamic_vlan.type&#x60;&#x3D;&#x3D;&#x60;standard&#x60;, property key is the Vlan ID and property value is \\\&quot;\\\&quot;   * if &#x60;dynamic_vlan.type&#x60;&#x3D;&#x3D;&#x60;airespace-interface-name&#x60;, property key is the Vlan ID and property value is the Airespace Interface Name | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

