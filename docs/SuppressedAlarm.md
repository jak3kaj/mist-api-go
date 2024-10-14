# SuppressedAlarm

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applies** | [***AllOfsuppressedAlarmApplies**](AllOfsuppressedAlarmApplies.md) |  | [optional] [default to null]
**Duration** | **float64** | duration, in seconds. Maximum duration is 86400 * 180 (180 days). 0 is to un-suppress alarms | [optional] [default to 3600]
**ScheduledTime** | **int32** | poch_time in seconds, Default as now, accepted time range is from now to now + 7 days | [optional] [default to null]
**Scope** | [***AllOfsuppressedAlarmScope**](AllOfsuppressedAlarmScope.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

