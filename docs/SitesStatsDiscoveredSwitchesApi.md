# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteDiscoveredSwitches**](SitesStatsDiscoveredSwitchesApi.md#CountSiteDiscoveredSwitches) | **Get** /api/v1/sites/{site_id}/stats/discovered_switches/count | countSiteDiscoveredSwitches
[**GetSiteDiscoveredSwitchesMetrics**](SitesStatsDiscoveredSwitchesApi.md#GetSiteDiscoveredSwitchesMetrics) | **Get** /api/v1/sites/{site_id}/stats/discovered_switches/metrics | getSiteDiscoveredSwitchesMetrics
[**SearchSiteDiscoveredSwitches**](SitesStatsDiscoveredSwitchesApi.md#SearchSiteDiscoveredSwitches) | **Get** /api/v1/sites/{site_id}/stats/discovered_switches/search | searchSiteDiscoveredSwitches
[**SearchSiteDiscoveredSwitchesMetrics**](SitesStatsDiscoveredSwitchesApi.md#SearchSiteDiscoveredSwitchesMetrics) | **Get** /api/v1/sites/{site_id}/stats/discovered_switch_metrics/search | searchSiteDiscoveredSwitchesMetrics

# **CountSiteDiscoveredSwitches**
> InlineResponse20016 CountSiteDiscoveredSwitches(ctx, siteId, optional)
countSiteDiscoveredSwitches

Count Discovered Switches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsDiscoveredSwitchesApiCountSiteDiscoveredSwitchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsDiscoveredSwitchesApiCountSiteDiscoveredSwitchesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of SiteDiscoveredSwitchesCountDistinct**](.md)|  | 
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

# **GetSiteDiscoveredSwitchesMetrics**
> InlineResponse200178 GetSiteDiscoveredSwitchesMetrics(ctx, siteId, optional)
getSiteDiscoveredSwitchesMetrics

Discovered switches related metrics, lists related switch system names & details if not compliant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsDiscoveredSwitchesApiGetSiteDiscoveredSwitchesMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsDiscoveredSwitchesApiGetSiteDiscoveredSwitchesMetricsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **threshold** | **optional.String**| configurable # ap per switch threshold, default 12 | 
 **systemName** | **optional.String**| system name for switch level metrics, optional | 

### Return type

[**InlineResponse200178**](inline_response_200_178.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteDiscoveredSwitches**
> InlineResponse200179 SearchSiteDiscoveredSwitches(ctx, siteId, optional)
searchSiteDiscoveredSwitches

Search Discovered Switches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsDiscoveredSwitchesApiSearchSiteDiscoveredSwitchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsDiscoveredSwitchesApiSearchSiteDiscoveredSwitchesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adopted** | **optional.Bool**|  | 
 **systemName** | **optional.String**|  | 
 **hostname** | **optional.String**|  | 
 **vendor** | **optional.String**|  | 
 **model** | **optional.String**|  | 
 **version** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse200179**](inline_response_200_179.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteDiscoveredSwitchesMetrics**
> InlineResponse200177 SearchSiteDiscoveredSwitchesMetrics(ctx, siteId, optional)
searchSiteDiscoveredSwitchesMetrics

Search Discovered Switch Metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsDiscoveredSwitchesApiSearchSiteDiscoveredSwitchesMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsDiscoveredSwitchesApiSearchSiteDiscoveredSwitchesMetricsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | [**optional.Interface of Scope3**](.md)| metric scope | 
 **type_** | [**optional.Interface of Type5**](.md)| metric type | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse200177**](inline_response_200_177.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

