# WebhookSdkclientScanDataEventScanDataItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | **string** | mac address of the AP associated with the BSSID scanned | [default to null]
**Band** | [***AllOfwebhookSdkclientScanDataEventScanDataItemBand**](AllOfwebhookSdkclientScanDataEventScanDataItemBand.md) |  | [default to null]
**Bssid** | **string** | BSSID found during client’s background scan for wifi | [default to null]
**Channel** | **int32** | channel of the band found in the scan | [default to null]
**Rssi** | **float64** | client’s RSSI relative to the BSSID scanned | [default to null]
**Ssid** | **string** | ESSID containing the BSSID scanned | [default to null]
**Timestamp** | **float64** | time the scan of the particular BSSID occurred | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

