# AllOfsiteSettingWifi

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CiscoEnabled** | **bool** |  | [optional] [default to true]
**Disable11k** | **bool** | whether to disable 11k | [optional] [default to false]
**DisableRadiosWhenPowerConstrained** | **bool** |  | [optional] [default to false]
**EnableArpSpoofCheck** | **bool** | when proxy_arp is enabled, check for arp spoofing. | [optional] [default to false]
**EnableSharedRadioScanning** | **bool** |  | [optional] [default to true]
**Enabled** | **bool** | enable WIFI feature (using SUB-MAN license) | [optional] [default to true]
**LocateConnected** | **bool** | whether to locate connected clients | [optional] [default to true]
**LocateUnconnected** | **bool** | whether to locate unconnected clients | [optional] [default to false]
**MeshAllowDfs** | **bool** | whether to allow Mesh to use DFS channels. For DFS channels, Remote Mesh AP would have to do CAC when scanning for new Base AP, which is slow and will distrupt the connection. If roaming is desired, keep it disabled. | [optional] [default to false]
**MeshEnableCrm** | **bool** | used to enable/disable CRM | [optional] [default to false]
**MeshEnabled** | **bool** | whether to enable Mesh feature for the site | [optional] [default to false]
**MeshPsk** | **string** | optional passphrase of mesh networking, default is generated randomly | [optional] [default to null]
**MeshSsid** | **string** | optional ssid of mesh networking, default is based on site_id | [optional] [default to null]
**ProxyArp** | [***Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

