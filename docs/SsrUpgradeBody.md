# SsrUpgradeBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | [***Object**](.md) |  | [optional] [default to null]
**DeviceIds** | **[]string** | list of 128T device IDs to upgrade | [default to null]
**RebootAt** | **int32** | reboot start time in epoch seconds, default is start_time, -1 disables reboot | [optional] [default to null]
**StartTime** | **int32** | 128T firmware download start time in epoch seconds, default is now, -1 disables download | [optional] [default to null]
**Strategy** | [***Object**](.md) |  | [optional] [default to null]
**Version** | **string** | 128T firmware version to upgrade (e.g. 5.3.0-93) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

