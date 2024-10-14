# WebhookLocationCentrakEvent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MapId** | **string** | map id | [optional] [default to null]
**MfgCompanyId** | **int32** | optional, BLE manufacturing company ID | [optional] [default to null]
**MfgData** | **string** | optional, BLE manufacturing data in hex byte-string format (i.e. “112233AABBCC”) | [optional] [default to null]
**Timestamp** | **int32** | timestamp of the event, epoch | [optional] [default to null]
**WifiBeaconExtendedInfo** | [**[]WifiBeaconExtendedInfoItems**](wifi_beacon_extended_info_items.md) | optional, list of extended beacon info packets heard from the client, frame and sequence control included with the payload | [optional] [default to null]
**X** | **float64** | x, in meter | [optional] [default to null]
**Y** | **float64** | y, in meter | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

