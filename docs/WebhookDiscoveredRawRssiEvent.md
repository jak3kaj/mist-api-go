# WebhookDiscoveredRawRssiEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApLoc** | **[]float64** | coordinates (if any) of reporting AP (updated once in 60s per client) | [optional] [default to null]
**Beam** | **int32** | antenna index, from 1-8, clock-wise starting from the LED | [default to null]
**DeviceId** | **string** | device id of the reporting AP | [default to null]
**IbeaconMajor** | **int32** |  | [optional] [default to null]
**IbeaconMinor** | **int32** |  | [optional] [default to null]
**IbeaconUuid** | **string** |  | [optional] [default to null]
**IsAsset** | **bool** |  | [optional] [default to null]
**Mac** | **string** | MAC of the asset/ beacon | [default to null]
**MapId** | **string** |  | [default to null]
**MfgCompanyId** | **string** | BLE manufacturing company ID | [optional] [default to null]
**MfgData** | **string** | BLE manufacturing data in hex byte-string format (ie: “112233AABBCC”) | [optional] [default to null]
**OrgId** | **string** |  | [default to null]
**Rssi** | **float64** | signal strength | [default to null]
**ServicePackets** | [**[]ServicePacket**](service_packet.md) | list of service data packets heard from the asset/ beacon | [optional] [default to null]
**SiteId** | **string** |  | [default to null]
**Timestamp** | **int32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

