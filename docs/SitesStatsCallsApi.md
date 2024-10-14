# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteCalls**](SitesStatsCallsApi.md#CountSiteCalls) | **Get** /api/v1/sites/{site_id}/stats/calls/count | countSiteCalls
[**ListSiteTroubleshootCalls**](SitesStatsCallsApi.md#ListSiteTroubleshootCalls) | **Get** /api/v1/sites/{site_id}/stats/calls/troubleshoot | listSiteTroubleshootCalls
[**SearchSiteCalls**](SitesStatsCallsApi.md#SearchSiteCalls) | **Get** /api/v1/sites/{site_id}/stats/calls/search | searchSiteCalls
[**TroubleshootSiteCall**](SitesStatsCallsApi.md#TroubleshootSiteCall) | **Get** /api/v1/sites/{site_id}/stats/calls/client/{client_mac}/troubleshoot | troubleshootSiteCall

# **CountSiteCalls**
> InlineResponse20016 CountSiteCalls(ctx, siteId, optional)
countSiteCalls

Count by Distinct Attributes of Calls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsCallsApiCountSiteCallsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsCallsApiCountSiteCallsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distrinct** | [**optional.Interface of CountSiteCallsDistrinct**](.md)|  | 
 **rating** | **optional.Int32**| feedback rating (e.g. \&quot;rating&#x3D;1\&quot; or \&quot;rating&#x3D;1,2\&quot;) | 
 **app** | **optional.String**|  | 
 **start** | **optional.String**|  | 
 **end** | **optional.String**|  | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteTroubleshootCalls**
> InlineResponse200170 ListSiteTroubleshootCalls(ctx, siteId, optional)
listSiteTroubleshootCalls

Summary of calls troubleshoot by site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsCallsApiListSiteTroubleshootCallsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsCallsApiListSiteTroubleshootCallsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ap** | **optional.String**| AP MAC | 
 **meetingId** | **optional.String**| meeting_id | 
 **mac** | **optional.String**| device identifier | 
 **app** | **optional.String**| Third party app name | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse200170**](inline_response_200_170.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteCalls**
> ResponseStatsCalls SearchSiteCalls(ctx, siteId, optional)
searchSiteCalls

Search Calls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsCallsApiSearchSiteCallsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsCallsApiSearchSiteCallsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **optional.String**| device identifier | 
 **app** | **optional.String**| Third party app name | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**ResponseStatsCalls**](response_stats_calls.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TroubleshootSiteCall**
> InlineResponse200169 TroubleshootSiteCall(ctx, siteId, clientMac, meetingId, optional)
troubleshootSiteCall

Troubleshoot a call

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **clientMac** | **string**|  | 
  **meetingId** | **string**| meeting_id | 
 **optional** | ***SitesStatsCallsApiTroubleshootSiteCallOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsCallsApiTroubleshootSiteCallOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **mac** | **optional.String**| device identifier | 
 **app** | **optional.String**| Third party app name | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse200169**](inline_response_200_169.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

