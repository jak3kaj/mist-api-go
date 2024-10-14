# WebhookAssetRawEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** | asset id | [default to null]
**Beam** | **int32** | antenna index, from 1-8, clock-wise starting from the LED | [default to null]
**DeviceId** | **string** | device where the asset reading is from | [default to null]
**IbeaconMajor** | **int32** | iBeacon major | [optional] [default to null]
**IbeaconMinor** | **int32** | iBeacon minor | [optional] [default to null]
**IbeaconUuid** | **string** | iBeacon UUID | [optional] [default to null]
**Mac** | **string** | MAC of the beacon | [default to null]
**MapId** | **string** | map id | [default to null]
**MfgCompanyId** | **float64** | optional, BLE manufacturing company ID | [default to null]
**MfgData** | **string** | optional, BLE manufacturing data in hex byte-string format (ie: “112233AABBCC”) | [default to null]
**Rssi** | **float64** | signal strength | [default to null]
**ServiceDataData** | **string** | optional, data from service data | [optional] [default to null]
**ServiceDataLastRxTime** | **int32** | optional, last data transmit time from service data | [optional] [default to null]
**ServiceDataRxCnt** | **int32** | optional, data transmit count from service data | [optional] [default to null]
**ServiceDataUuid** | **string** | optional, UUID from service data | [optional] [default to null]
**ServicePackets** | [**[]ServicePacket**](service_packet.md) | list of service data packets heard from the asset/ beacon | [optional] [default to null]
**SiteId** | **string** |  | [default to null]
**Timestamp** | **float64** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

