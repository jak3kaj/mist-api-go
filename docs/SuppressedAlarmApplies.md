# SuppressedAlarmApplies

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** | UUID of the current org (if provided, the alarms will be suppressed at org level) | [optional] [default to null]
**SiteIds** | **[]string** | List of UUID of the sites within the org (if provided, the alarms will be suppressed for all the mentioned sites under the org) | [optional] [default to null]
**SitegroupIds** | **[]string** | List of UUID of the site groups within the org (if provided, the alarms will be suppressed for all the sites under those site groups) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

