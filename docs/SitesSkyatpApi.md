# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteSkyatpEvents**](SitesSkyatpApi.md#CountSiteSkyatpEvents) | **Get** /api/v1/sites/{site_id}/skyatp/events/count | countSiteSkyatpEvents
[**SearchSiteSkyatpEvents**](SitesSkyatpApi.md#SearchSiteSkyatpEvents) | **Get** /api/v1/sites/{site_id}/skyatp/events/search | searchSiteSkyatpEvents

# **CountSiteSkyatpEvents**
> InlineResponse20016 CountSiteSkyatpEvents(ctx, siteId, optional)
countSiteSkyatpEvents

Count by Distinct Attributes of Skyatp Events (WIP)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesSkyatpApiCountSiteSkyatpEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSkyatpApiCountSiteSkyatpEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of SiteSkyAtpEventsCountDistinct**](.md)|  | 
 **type_** | **optional.String**| event type, e.g. cc, fs, mw | 
 **mac** | **optional.String**| client MAC | 
 **deviceMac** | **optional.String**| device MAC | 
 **threatLevel** | **optional.Int32**| threat level | 
 **ipAddress** | **optional.String**|  | 
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

# **SearchSiteSkyatpEvents**
> InlineResponse200145 SearchSiteSkyatpEvents(ctx, siteId, optional)
searchSiteSkyatpEvents

Search Skyatp Events (WIP)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesSkyatpApiSearchSiteSkyatpEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSkyatpApiSearchSiteSkyatpEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| event type, e.g. cc, fs, mw | 
 **mac** | **optional.String**| client MAC | 
 **deviceMac** | **optional.String**| device MAC | 
 **threatLevel** | **optional.Int32**| threat level | 
 **ipAddress** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse200145**](inline_response_200_145.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

