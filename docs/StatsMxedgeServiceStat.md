# StatsMxedgeServiceStat

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtIp** | **string** | external IP from ep-terminator’s point of view. valid only for service having its own cloud connection | [optional] [default to null]
**LastSeen** | **float64** | timestamp when the last stats is seen (cloud unix time, in second). valid only for service having its own stats or whole system (last among last_seen of all services) | [optional] [default to null]
**PackageState** | **string** | package/service installation state. | [optional] [default to null]
**PackageVersion** | **string** | package/service installation state. | [optional] [default to null]
**RunningState** | **string** | service running state. | [optional] [default to null]
**Uptime** | **int32** | service uptime. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

