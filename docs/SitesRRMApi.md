# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteCurrentChannelPlanning**](SitesRRMApi.md#GetSiteCurrentChannelPlanning) | **Get** /api/v1/sites/{site_id}/rrm/current | getSiteCurrentChannelPlanning
[**GetSiteCurrentRrmConsiderations**](SitesRRMApi.md#GetSiteCurrentRrmConsiderations) | **Get** /api/v1/sites/{site_id}/rrm/current/devices/{device_id}/band/{band} | getSiteCurrentRrmConsiderations
[**GetSiteCurrentRrmNeighbors**](SitesRRMApi.md#GetSiteCurrentRrmNeighbors) | **Get** /api/v1/sites/{site_id}/rrm/neighbors/band/{band} | getSiteCurrentRrmNeighbors
[**GetSiteRrmEvents**](SitesRRMApi.md#GetSiteRrmEvents) | **Get** /api/v1/sites/{site_id}/rrm/events | getSiteRrmEvents

# **GetSiteCurrentChannelPlanning**
> InlineResponse200137 GetSiteCurrentChannelPlanning(ctx, siteId)
getSiteCurrentChannelPlanning

Get Current Channel Planning

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200137**](inline_response_200_137.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteCurrentRrmConsiderations**
> InlineResponse200138 GetSiteCurrentRrmConsiderations(ctx, siteId, deviceId, band)
getSiteCurrentRrmConsiderations

Get Current RRM Considerations for an AP on a Specific Band

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
  **band** | [**Band8**](.md)| 802.11 Band | 

### Return type

[**InlineResponse200138**](inline_response_200_138.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteCurrentRrmNeighbors**
> InlineResponse200140 GetSiteCurrentRrmNeighbors(ctx, siteId, band, optional)
getSiteCurrentRrmNeighbors

Get Current RRM observed neighbors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **band** | [**Band10**](.md)| 802.11 Band | 
 **optional** | ***SitesRRMApiGetSiteCurrentRrmNeighborsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesRRMApiGetSiteCurrentRrmNeighborsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse200140**](inline_response_200_140.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteRrmEvents**
> InlineResponse200139 GetSiteRrmEvents(ctx, siteId, optional)
getSiteRrmEvents

Get Site RRM Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesRRMApiGetSiteRrmEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesRRMApiGetSiteRrmEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **band** | [**optional.Interface of Band9**](.md)| 802.11 Band | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse200139**](inline_response_200_139.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

