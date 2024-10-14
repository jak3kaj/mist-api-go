# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AckOrgAlarm**](OrgsAlarmsApi.md#AckOrgAlarm) | **Post** /api/v1/orgs/{org_id}/alarms/{alarm_id}/ack | ackOrgAlarm
[**AckOrgAllAlarms**](OrgsAlarmsApi.md#AckOrgAllAlarms) | **Post** /api/v1/orgs/{org_id}/alarms/ack_all | ackOrgAllAlarms
[**AckOrgMultipleAlarms**](OrgsAlarmsApi.md#AckOrgMultipleAlarms) | **Post** /api/v1/orgs/{org_id}/alarms/ack | ackOrgMultipleAlarms
[**CountOrgAlarms**](OrgsAlarmsApi.md#CountOrgAlarms) | **Get** /api/v1/orgs/{org_id}/alarms/count | countOrgAlarms
[**SearchOrgAlarms**](OrgsAlarmsApi.md#SearchOrgAlarms) | **Get** /api/v1/orgs/{org_id}/alarms/search | searchOrgAlarms
[**UnackOrgAllArlarms**](OrgsAlarmsApi.md#UnackOrgAllArlarms) | **Post** /api/v1/orgs/{org_id}/alarms/unack_all | unackOrgAllArlarms
[**UnackOrgMultipleAlarms**](OrgsAlarmsApi.md#UnackOrgMultipleAlarms) | **Post** /api/v1/orgs/{org_id}/alarms/unack | unackOrgMultipleAlarms

# **AckOrgAlarm**
> AckOrgAlarm(ctx, orgId, alarmId, optional)
ackOrgAlarm

Ack Org Alarm

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **alarmId** | [**string**](.md)|  | 
 **optional** | ***OrgsAlarmsApiAckOrgAlarmOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAlarmsApiAckOrgAlarmOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AlarmIdAckBody**](AlarmIdAckBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AckOrgAllAlarms**
> AckOrgAllAlarms(ctx, orgId, optional)
ackOrgAllAlarms

Ack all Org Alarms  **N.B.**: Batch size for multiple alarm ack and unack has to be less or or equal to 1000.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAlarmsApiAckOrgAllAlarmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAlarmsApiAckOrgAllAlarmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AlarmsAckAllBody**](AlarmsAckAllBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AckOrgMultipleAlarms**
> AckOrgMultipleAlarms(ctx, orgId, optional)
ackOrgMultipleAlarms

Ack multiple Org Alarms

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAlarmsApiAckOrgMultipleAlarmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAlarmsApiAckOrgMultipleAlarmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AlarmsAckBody**](AlarmsAckBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountOrgAlarms**
> InlineResponse20016 CountOrgAlarms(ctx, orgId, optional)
countOrgAlarms

Count Org Alarms

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAlarmsApiCountOrgAlarmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAlarmsApiCountOrgAlarmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | **optional.String**|  | 
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

# **SearchOrgAlarms**
> AlarmSearchResult SearchOrgAlarms(ctx, orgId, optional)
searchOrgAlarms

Search Org Alarms

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAlarmsApiSearchOrgAlarmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAlarmsApiSearchOrgAlarmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | [**optional.Interface of string**](.md)| site ID | 
 **type_** | **optional.String**| alarm type | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]

### Return type

[**AlarmSearchResult**](alarm_search_result.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnackOrgAllArlarms**
> UnackOrgAllArlarms(ctx, orgId, optional)
unackOrgAllArlarms

Unack all Org Alarms  **N.B.**: Batch size for multiple alarm ack and unack has to be less or or equal to 1000.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAlarmsApiUnackOrgAllArlarmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAlarmsApiUnackOrgAllArlarmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AlarmsUnackAllBody**](AlarmsUnackAllBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnackOrgMultipleAlarms**
> UnackOrgMultipleAlarms(ctx, orgId, optional)
unackOrgMultipleAlarms

Unack multiple Org Alarms

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAlarmsApiUnackOrgMultipleAlarmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAlarmsApiUnackOrgMultipleAlarmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AlarmsUnackBody**](AlarmsUnackBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

