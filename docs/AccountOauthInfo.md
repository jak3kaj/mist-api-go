# AccountOauthInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Linked app(zoom/teams/intune) account id | [optional] [default to null]
**Company** | **string** | Name of the company whose account mist has subscribed to | [optional] [default to null]
**Error_** | **string** | This error is provided when the account fails to fetch token/data | [optional] [default to null]
**Errors** | **[]string** |  | [optional] [default to null]
**LastStatus** | **string** | Is the last data pull for account is successful or not | [optional] [default to null]
**LastSync** | **int64** | Last data pull timestamp, background jobs that pull account data | [optional] [default to null]
**LinkedBy** | **string** | First name of the user who linked the account | [optional] [default to null]
**MaxDailyApiRequests** | **int32** | Zoom daily api request quota, https://developers.zoom.us/docs/api/rest/rate-limits/ | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

