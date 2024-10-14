# InventoryUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableAutoConfig** | **bool** | if &#x60;op&#x60;&#x3D;&#x3D;&#x60;assign&#x60;, a **cloud-ready** switch/gateway will be managed/configured by Mist by default, this disabled the behavior | [optional] [default to false]
**Macs** | **[]string** | if &#x60;op&#x60;&#x3D;&#x3D;&#x60;assign&#x60;, &#x60;op&#x60;&#x3D;&#x3D;&#x60;unassign&#x60;, &#x60;op&#x60;&#x3D;&#x3D;&#x60;upgrade_to_mist&#x60;or &#x60;op&#x60;&#x3D;&#x3D;&#x60;downgrade_to_jsi&#x60; , list of MAC, e.g. [\&quot;5c5b350e0001\&quot;] | [optional] [default to null]
**Managed** | **bool** | if &#x60;op&#x60;&#x3D;&#x3D;&#x60;assign&#x60;, an **adopted** switch/gateway will not be managed/configured by Mist by default, this enables the behavior | [optional] [default to false]
**NoReassign** | **bool** | if &#x60;op&#x60;&#x3D;&#x3D;&#x60;assign&#x60;, if true, treat site assignment against an already assigned AP as error | [optional] [default to null]
**Op** | [***AllOfinventoryUpdateOp**](AllOfinventoryUpdateOp.md) |  | [default to null]
**Serials** | **[]string** | if &#x60;op&#x60;&#x3D;&#x3D;&#x60;delete&#x60;, list of serial numbers, e.g. [\&quot;FXLH2015150025\&quot;] | [optional] [default to null]
**SiteId** | **string** | if &#x60;op&#x60;&#x3D;&#x3D;&#x60;assign&#x60;, target site id | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

