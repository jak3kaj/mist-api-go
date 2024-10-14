# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteInsightMetrics**](SitesInsightsApi.md#GetSiteInsightMetrics) | **Get** /api/v1/sites/{site_id}/insights/{metric} | getSiteInsightMetrics
[**GetSiteInsightMetricsForClient**](SitesInsightsApi.md#GetSiteInsightMetricsForClient) | **Get** /api/v1/sites/{site_id}/insights/client/{client_mac}/{metric} | getSiteInsightMetricsForClient
[**GetSiteInsightMetricsForDevice**](SitesInsightsApi.md#GetSiteInsightMetricsForDevice) | **Get** /api/v1/sites/{site_id}/insights/device/{device_mac}/{metric} | getSiteInsightMetricsForDevice

# **GetSiteInsightMetrics**
> InlineResponse200126 GetSiteInsightMetrics(ctx, siteId, metric, optional)
getSiteInsightMetrics

Get Site Insight Metrics See metrics possibilities at /api/v1/const/insight_metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **metric** | **string**| see /api/v1/const/insight_metrics for available metrics | 
 **optional** | ***SitesInsightsApiGetSiteInsightMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesInsightsApiGetSiteInsightMetricsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **interval** | **optional.String**| Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. | 
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse200126**](inline_response_200_126.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteInsightMetricsForClient**
> InlineResponse200126 GetSiteInsightMetricsForClient(ctx, siteId, clientMac, metric, optional)
getSiteInsightMetricsForClient

Get Client Insight Metrics See metrics possibilities at /api/v1/const/insight_metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **clientMac** | **string**|  | 
  **metric** | **string**| see /api/v1/const/insight_metrics for available metrics | 
 **optional** | ***SitesInsightsApiGetSiteInsightMetricsForClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesInsightsApiGetSiteInsightMetricsForClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **interval** | **optional.String**| Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. | 
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse200126**](inline_response_200_126.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteInsightMetricsForDevice**
> InlineResponse200127 GetSiteInsightMetricsForDevice(ctx, siteId, metric, deviceMac, optional)
getSiteInsightMetricsForDevice

Get AP Insight Metrics See metrics possibilities at /api/v1/const/insight_metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **metric** | **string**| see /api/v1/const/insight_metrics for available metrics | 
  **deviceMac** | **string**|  | 
 **optional** | ***SitesInsightsApiGetSiteInsightMetricsForDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesInsightsApiGetSiteInsightMetricsForDeviceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **interval** | **optional.String**| Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. | 
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse200127**](inline_response_200_127.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

