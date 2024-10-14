# WebhookSdkclientScanDataEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionAp** | **string** | mac address of the AP the client is connected to | [default to null]
**ConnectionBand** | **string** | 5GHz or 2.4GHz band, of the BSSID the client is connected to | [default to null]
**ConnectionBssid** | **string** | BSSID of the AP the client is connected to | [default to null]
**ConnectionChannel** | **int32** | channel of the band the client is connected to | [default to null]
**ConnectionRssi** | **float64** | RSSI of the client’s connection to the AP/BSSID | [default to null]
**LastSeen** | **float64** | time client last seen with scan data | [default to null]
**Mac** | **string** | the client’s mac | [default to null]
**ScanData** | [**[]WebhookSdkclientScanDataEventScanDataItem**](webhook_sdkclient_scan_data_event_scan_data_item.md) |  | [optional] [default to null]
**SiteId** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

