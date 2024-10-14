# ResponseSiteDeviceUpgradeCounts

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadRequested** | **int32** | count of devices which cloud has requested to download firmware | [optional] [default to null]
**Downloaded** | **int32** | count of ap’s which have the firmware downloaded | [optional] [default to null]
**Failed** | **int32** | count of devices which have failed to upgrade | [optional] [default to null]
**RebootInProgress** | **int32** | count of devices which are rebooting | [optional] [default to null]
**Rebooted** | **int32** | count of devices which have rebooted successfully | [optional] [default to null]
**Scheduled** | **int32** | count of devices which cloud has scheduled an upgrade for | [optional] [default to null]
**Skipped** | **int32** | count of devices which skipped upgrade since requested version was same as running version. Use force to always upgrade | [optional] [default to null]
**Total** | **int32** | count of devices part of this upgrade | [optional] [default to null]
**Upgraded** | **int32** | count of devices which have upgraded successfully | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

