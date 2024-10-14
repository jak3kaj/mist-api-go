# UpgradeOrgDevices

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanaryPhases** | **[]int32** | phases for canary deployment. Each phase represents percentage of devices that need to be upgraded in that phase. | [optional] [default to [1,10,50,100]]
**EnableP2p** | **bool** | whether to allow local AP-to-AP FW upgrade | [optional] [default to null]
**Force** | **bool** | true will force upgrade when requested version is same as running version | [optional] [default to false]
**MaxFailurePercentage** | **float64** | percentage of failures allowed across the entire upgrade(not applicable for &#x60;big_bang&#x60;) | [optional] [default to 5]
**MaxFailures** | **[]int32** | number of failures allowed within each phase. Only applicable for &#x60;canary&#x60;. Array length should be same as &#x60;canary_phases&#x60;. Will be used if provided, else &#x60;max_failure_percentage&#x60; will be used | [optional] [default to null]
**Models** | **[]string** |  | [optional] [default to null]
**P2pClusterSize** | **int32** |  | [optional] [default to 10]
**P2pParallelism** | **int32** | number of parallel p2p download batches to creat | [optional] [default to null]
**Reboot** | **bool** | Reboot device immediately after upgrade is completed (Available on Junos OS devices) | [optional] [default to false]
**RebootAt** | **float64** | reboot start time in epoch seconds, default is &#x60;start_time&#x60; | [optional] [default to null]
**RrmFirstBatchPercentage** | **int32** | percentage of AP’s that need to be present in the first rrm batch | [optional] [default to null]
**RrmMaxBatchPercentage** | **int32** | max percentage of AP’s that need to be present in each rrm batch | [optional] [default to null]
**RrmMeshUpgrade** | [***AllOfupgradeOrgDevicesRrmMeshUpgrade**](AllOfupgradeOrgDevicesRrmMeshUpgrade.md) |  | [optional] [default to null]
**RrmNodeOrder** | [***AllOfupgradeOrgDevicesRrmNodeOrder**](AllOfupgradeOrgDevicesRrmNodeOrder.md) |  | [optional] [default to null]
**RrmSlowRamp** | **bool** | true will make rrm batch sizes slowly ramp up | [optional] [default to null]
**SiteIds** | **[]string** |  | [optional] [default to null]
**Snapshot** | **bool** | Perform recovery snapshot after device is rebooted (Available on Junos OS devices) | [optional] [default to false]
**StartTime** | **float64** | upgrade start time in epoch seconds, default is now | [optional] [default to null]
**Strategy** | [***AllOfupgradeOrgDevicesStrategy**](AllOfupgradeOrgDevicesStrategy.md) |  | [optional] [default to null]
**Version** | **string** | specific version / stable | [optional] [default to latest]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

