# AllOfsiteSettingZoneOccupancyAlert

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailNotifiers** | **[]string** | list of email addresses to send email notifications when the alert threshold is reached | [optional] [default to null]
**Enabled** | **bool** | indicate whether zone occupancy alert is enabled for the site | [optional] [default to false]
**Threshold** | **int32** | sending zone-occupancy-alert webhook message only if a zone stays non-compliant (i.e. actual occupancy &gt; occupancy_limit) for a minimum duration specified in the threshold, in minutes | [optional] [default to 5]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

