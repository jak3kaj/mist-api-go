# ResponseDeviceUpgradeCounts

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadRequested** | **[]string** | list of devices MAC Addresses which cloud has requested to download firmware | [optional] [default to null]
**Downloaded** | **[]string** | list of devices MAC Addresses which have the firmware downloaded | [optional] [default to null]
**Failed** | **[]string** | list of devices MAC Addresses which have failed to upgrade | [optional] [default to null]
**RebootInProgress** | **[]string** | list of devices MAC Addresses which are rebooting | [optional] [default to null]
**Rebooted** | **[]string** | list of devices MAC Addresses which have rebooted successfully | [optional] [default to null]
**Scheduled** | **[]string** | list of devices MAC Addresses which cloud has scheduled an upgrade for | [optional] [default to null]
**Skipped** | **[]string** | list of devices MAC Addresses which skipped upgrade since requested version was same as running version. Use force to always upgrade | [optional] [default to null]
**Total** | **int32** | count of devices part of this upgrade | [optional] [default to null]
**Upgraded** | **[]string** | count of devices which have upgraded successfully | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

