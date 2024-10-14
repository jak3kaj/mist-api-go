# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AckSiteAlarm**](SitesAlarmsApi.md#AckSiteAlarm) | **Post** /api/v1/sites/{site_id}/alarms/{alarm_id}/ack | ackSiteAlarm
[**AckSiteAllAlarms**](SitesAlarmsApi.md#AckSiteAllAlarms) | **Post** /api/v1/sites/{site_id}/alarms/ack_all | ackSiteAllAlarms
[**AckSiteMultipleAlarms**](SitesAlarmsApi.md#AckSiteMultipleAlarms) | **Post** /api/v1/sites/{site_id}/alarms/ack | AckSiteMultipleAlarms
[**CountSiteAlarms**](SitesAlarmsApi.md#CountSiteAlarms) | **Get** /api/v1/sites/{site_id}/alarms/count | countSiteAlarms
[**SearchSiteAlarms**](SitesAlarmsApi.md#SearchSiteAlarms) | **Get** /api/v1/sites/{site_id}/alarms/search | searchSiteAlarms
[**SubscribeSiteAlarms**](SitesAlarmsApi.md#SubscribeSiteAlarms) | **Post** /api/v1/sites/{site_id}/subscriptions | SubscribeSiteAlarms
[**UnackSiteAlarm**](SitesAlarmsApi.md#UnackSiteAlarm) | **Post** /api/v1/sites/{site_id}/alarms/{alarm_id}/unack | unackSiteAlarm
[**UnackSiteAllArlarms**](SitesAlarmsApi.md#UnackSiteAllArlarms) | **Post** /api/v1/sites/{site_id}/alarms/unack_all | unackSiteAllArlarms
[**UnackSiteMultipleAlarms**](SitesAlarmsApi.md#UnackSiteMultipleAlarms) | **Post** /api/v1/sites/{site_id}/alarms/unack | unackSiteMultipleAlarms
[**UnsubscribeSiteAlarms**](SitesAlarmsApi.md#UnsubscribeSiteAlarms) | **Delete** /api/v1/sites/{site_id}/subscriptions | UnsubscribeSiteAlarms

# **AckSiteAlarm**
> AckSiteAlarm(ctx, siteId, alarmId, optional)
ackSiteAlarm

Ack Site Alarm

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **alarmId** | [**string**](.md)|  | 
 **optional** | ***SitesAlarmsApiAckSiteAlarmOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesAlarmsApiAckSiteAlarmOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AlarmIdAckBody1**](AlarmIdAckBody1.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AckSiteAllAlarms**
> AckSiteAllAlarms(ctx, siteId, optional)
ackSiteAllAlarms

Ack all Site Alarms  **N.B.**: Batch size for multiple alarm ack and unack has to be less or or equal to 1000.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesAlarmsApiAckSiteAllAlarmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesAlarmsApiAckSiteAllAlarmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NoteString**](NoteString.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AckSiteMultipleAlarms**
> AckSiteMultipleAlarms(ctx, siteId, optional)
AckSiteMultipleAlarms

Ack multiple Site Alarms

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesAlarmsApiAckSiteMultipleAlarmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesAlarmsApiAckSiteMultipleAlarmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AlarmsAckBody1**](AlarmsAckBody1.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountSiteAlarms**
> InlineResponse20016 CountSiteAlarms(ctx, siteId, optional)
countSiteAlarms

Count Site Alarms

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesAlarmsApiCountSiteAlarmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesAlarmsApiCountSiteAlarmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of Distinct3**](.md)| Group by and count the alarms by some distinct field | 
 **ackAdminName** | **optional.String**| Name of the admins who have acked the alarms; accepts multiple values separated by comma | 
 **acked** | **optional.Bool**|  | 
 **type_** | **optional.String**| Key-name of the alarms; accepts multiple values separated by comma | 
 **severity** | **optional.String**| Alarm severity; accepts multiple values separated by comma | 
 **group** | **optional.String**| Alarm group name; accepts multiple values separated by comma | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteAlarms**
> AlarmSearchResult SearchSiteAlarms(ctx, siteId, optional)
searchSiteAlarms

Search Site Alarms

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesAlarmsApiSearchSiteAlarmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesAlarmsApiSearchSiteAlarmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| Key-name of the alarms; accepts multiple values separated by comma | 
 **ackAdminName** | **optional.String**| Name of the admins who have acked the alarms; accepts multiple values separated by comma | 
 **acked** | **optional.Bool**|  | 
 **severity** | **optional.String**| Alarm severity; accepts multiple values separated by comma | 
 **group** | **optional.String**| Alarm group name; accepts multiple values separated by comma | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**AlarmSearchResult**](alarm_search_result.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubscribeSiteAlarms**
> SubscribeSiteAlarms(ctx, siteId)
SubscribeSiteAlarms

Subscribe to Site Alarms

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnackSiteAlarm**
> UnackSiteAlarm(ctx, siteId, alarmId, optional)
unackSiteAlarm

Unack Site Alarm

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **alarmId** | [**string**](.md)|  | 
 **optional** | ***SitesAlarmsApiUnackSiteAlarmOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesAlarmsApiUnackSiteAlarmOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AlarmIdUnackBody**](AlarmIdUnackBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnackSiteAllArlarms**
> UnackSiteAllArlarms(ctx, siteId, optional)
unackSiteAllArlarms

Unack all Site Alarms  **N.B.**: Batch size for multiple alarm ack and unack has to be less or or equal to 1000.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesAlarmsApiUnackSiteAllArlarmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesAlarmsApiUnackSiteAllArlarmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AlarmsUnackAllBody1**](AlarmsUnackAllBody1.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnackSiteMultipleAlarms**
> UnackSiteMultipleAlarms(ctx, siteId, optional)
unackSiteMultipleAlarms

Unack multiple Site Alarms

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesAlarmsApiUnackSiteMultipleAlarmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesAlarmsApiUnackSiteMultipleAlarmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AlarmsUnackBody1**](AlarmsUnackBody1.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsubscribeSiteAlarms**
> UnsubscribeSiteAlarms(ctx, siteId)
UnsubscribeSiteAlarms

Unsubscribe to Site Alarms

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

