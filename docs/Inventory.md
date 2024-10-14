# Inventory

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adopted** | **bool** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;type&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; whether the switch/gateway is adopted | [optional] [default to null]
**Connected** | **bool** | whether the device is connected | [optional] [default to null]
**CreatedTime** | **float64** | inventory created time, in epoch | [optional] [default to null]
**DeviceprofileId** | **string** | deviceprofile id if assigned, null if not assigned | [optional] [default to null]
**Hostname** | **string** | hostname reported by the device | [optional] [default to null]
**HwRev** | **string** | device hardware revision number | [optional] [default to null]
**Id** | **string** | device id | [optional] [default to null]
**Jsi** | **bool** |  | [optional] [default to null]
**Mac** | **string** | device MAC address | [optional] [default to null]
**Magic** | **string** | device claim code | [optional] [default to null]
**Model** | **string** | device model | [optional] [default to null]
**ModifiedTime** | **float64** | inventory last modified time, in epoch | [optional] [default to null]
**Name** | **string** | device name if configured | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**Serial** | **string** | device serial | [optional] [default to null]
**SiteId** | **string** |  | [optional] [default to null]
**Sku** | **string** | device stock keeping unit | [optional] [default to null]
**Type_** | [***AllOfinventoryType_**](AllOfinventoryType_.md) |  | [optional] [default to null]
**VcMac** | **string** | only if &#x60;type&#x60;&#x3D;&#x3D;&#x60;switch&#x60;, MAC Address of the Virtual Chassis | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

