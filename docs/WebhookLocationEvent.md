# WebhookLocationEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatteryVoltage** | **int32** |  | [optional] [default to null]
**EddystoneUidInstance** | **string** |  | [optional] [default to null]
**EddystoneUidNamespace** | **string** |  | [optional] [default to null]
**EddystoneUrlUrl** | **string** |  | [optional] [default to null]
**IbeaconMajor** | **int32** |  | [optional] [default to null]
**IbeaconMinor** | **int32** |  | [optional] [default to null]
**IbeaconUuid** | **string** |  | [optional] [default to null]
**Id** | **string** | unique id of the client (a client would have different id for different org) | [default to null]
**Mac** | **string** |  | [optional] [default to null]
**MapId** | **string** | map id | [default to null]
**MfgCompanyId** | **int32** | optional, BLE manufacturing company ID | [optional] [default to null]
**MfgData** | **string** | optional, BLE manufacturing data in hex byte-string format (ie \&quot;112233AABBCC\&quot;) | [optional] [default to null]
**Name** | **string** | name of the client, may be empty | [optional] [default to null]
**SiteId** | **string** |  | [default to null]
**Timestamp** | **int32** | timestamp of the event, epoch | [default to null]
**Type_** | **string** |  | [default to null]
**WifiBeaconExtendedInfo** | [**[]WifiBeaconExtendedInfoItems**](wifi_beacon_extended_info_items.md) | optional, list of extended beacon info packets heard from the client, frame and sequence control included with the payload | [optional] [default to null]
**X** | **float64** | x, in meter | [default to null]
**Y** | **float64** | y, in meter | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

