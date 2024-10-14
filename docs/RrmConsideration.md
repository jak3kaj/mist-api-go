# RrmConsideration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **int32** |  | [default to null]
**Noise** | **float64** |  | [default to null]
**OtherRssi** | **float64** | the avg RSSI heard from other APs (that does NOT belongs to the same site) | [optional] [default to null]
**OtherSsid** | **string** | SSID from other AP that we heard from with the max RSSI | [optional] [default to null]
**UtilScore** | **float64** | utilization score, 0-1, lower means less utilization (cleaner RF) | [default to null]
**UtilScoreNonWifi** | **float64** | non-wifi utilization score, 0-1, lower means less utilization (cleaner RF) | [default to null]
**UtilScoreOther** | **float64** | other utilization score, 0-1, lower means less utilization (cleaner RF) | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

