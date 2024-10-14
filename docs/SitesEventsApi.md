# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteSystemEvents**](SitesEventsApi.md#CountSiteSystemEvents) | **Get** /api/v1/sites/{site_id}/events/system/count | countSiteSystemEvents
[**GetSiteRoamingEvents**](SitesEventsApi.md#GetSiteRoamingEvents) | **Get** /api/v1/sites/{site_id}/events/fast_roam | getSiteRoamingEvents
[**SearchSiteSystemEvents**](SitesEventsApi.md#SearchSiteSystemEvents) | **Get** /api/v1/sites/{site_id}/events/system/search | searchSiteSystemEvents

# **CountSiteSystemEvents**
> InlineResponse20016 CountSiteSystemEvents(ctx, siteId, optional)
countSiteSystemEvents

Count System Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesEventsApiCountSiteSystemEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesEventsApiCountSiteSystemEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of SiteSystemEventsCountDistinct**](.md)|  | 
 **type_** | **optional.String**| see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteRoamingEvents**
> InlineResponse200125 GetSiteRoamingEvents(ctx, siteId, optional)
getSiteRoamingEvents

Get Roaming Events data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesEventsApiGetSiteRoamingEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesEventsApiGetSiteRoamingEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of Type2**](.md)| event type | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse200125**](inline_response_200_125.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteSystemEvents**
> InlineResponse20048 SearchSiteSystemEvents(ctx, siteId, optional)
searchSiteSystemEvents

Search System Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesEventsApiSearchSiteSystemEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesEventsApiSearchSiteSystemEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse20048**](inline_response_200_48.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

