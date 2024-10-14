# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteAnomalyEvents**](SitesAnomalyApi.md#GetSiteAnomalyEvents) | **Get** /api/v1/sites/{site_id}/anomaly/{metric} | getSiteAnomalyEvents
[**GetSiteAnomalyEventsForClient**](SitesAnomalyApi.md#GetSiteAnomalyEventsForClient) | **Get** /api/v1/sites/{site_id}/anomaly/client/{client_mac}/{metric} | getSiteAnomalyEventsForClient
[**GetSiteAnomalyEventsforDevice**](SitesAnomalyApi.md#GetSiteAnomalyEventsforDevice) | **Get** /api/v1/sites/{site_id}/anomaly/device/{device_mac}/{metric} | getSiteAnomalyEventsforDevice

# **GetSiteAnomalyEvents**
> ResponseAnomalySearch GetSiteAnomalyEvents(ctx, siteId, metric)
getSiteAnomalyEvents

Get Site Anomaly Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **metric** | **string**| see /api/v1/const/insight_metrics for available metrics | 

### Return type

[**ResponseAnomalySearch**](response_anomaly_search.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteAnomalyEventsForClient**
> ResponseAnomalySearch GetSiteAnomalyEventsForClient(ctx, siteId, clientMac, metric)
getSiteAnomalyEventsForClient

Get Client Anomaly Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **clientMac** | **string**|  | 
  **metric** | **string**| see /api/v1/const/insight_metrics for available metrics | 

### Return type

[**ResponseAnomalySearch**](response_anomaly_search.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteAnomalyEventsforDevice**
> ResponseAnomalySearch GetSiteAnomalyEventsforDevice(ctx, siteId, metric, deviceMac)
getSiteAnomalyEventsforDevice

Get Device Anomaly Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **metric** | **string**| see /api/v1/const/insight_metrics for available metrics | 
  **deviceMac** | **string**|  | 

### Return type

[**ResponseAnomalySearch**](response_anomaly_search.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

