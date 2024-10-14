# ResponseSiteDeviceUpgrade

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counts** | [***AllOfresponseSiteDeviceUpgradeCounts**](AllOfresponseSiteDeviceUpgradeCounts.md) |  | [optional] [default to null]
**CurrentPhase** | **int32** | current canary or rrm phase in progress | [optional] [default to null]
**EnableP2p** | **bool** | whether to allow local AP-to-AP FW upgrade | [optional] [default to null]
**Force** | **bool** | whether to force upgrade when requested version is same as running version | [optional] [default to null]
**Id** | **string** | unique id for the upgrade | [default to null]
**MaxFailurePercentage** | **int32** | percentage of failures allowed | [optional] [default to null]
**MaxFailures** | **[]int32** | number of failures allowed within a canary phase or serial rollout | [optional] [default to null]
**RebootAt** | **int32** | reboot start time in epoch | [optional] [default to null]
**StartTime** | **float64** | firmware download start time in epoch | [optional] [default to null]
**Status** | [***AllOfresponseSiteDeviceUpgradeStatus**](AllOfresponseSiteDeviceUpgradeStatus.md) |  | [optional] [default to null]
**Strategy** | [***AllOfresponseSiteDeviceUpgradeStrategy**](AllOfresponseSiteDeviceUpgradeStrategy.md) |  | [optional] [default to null]
**TargetVersion** | **string** | version to upgrade to | [optional] [default to null]
**UpgradePlan** | [***interface{}**](interface{}.md) | a dictionary of rrm phase number to devices part of that phase | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

