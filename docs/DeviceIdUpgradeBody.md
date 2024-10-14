# DeviceIdUpgradeBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reboot** | **bool** | Reboot device immediately after upgrade is completed (Available on Junos OS devices) | [optional] [default to false]
**RebootAt** | **int32** | reboot start time in epoch | [optional] [default to null]
**Snapshot** | **bool** | Perform recovery snapshot after device is rebooted (Available on Junos OS devices) | [optional] [default to false]
**StartTime** | **float64** | firmware download start time in epoch | [optional] [default to null]
**Version** | **string** | specific version / &#x60;stable&#x60;, default is to use the latest | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

