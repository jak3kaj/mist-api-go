# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteBgpStats**](SitesStatsBGPPeersApi.md#CountSiteBgpStats) | **Get** /api/v1/sites/{site_id}/stats/bgp_peers/count | countSiteBgpStats
[**SearchSiteBgpStats**](SitesStatsBGPPeersApi.md#SearchSiteBgpStats) | **Get** /api/v1/sites/{site_id}/stats/bgp_peers/search | searchSiteBgpStats

# **CountSiteBgpStats**
> InlineResponse20016 CountSiteBgpStats(ctx, siteId, optional)
countSiteBgpStats

Count BGP Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsBGPPeersApiCountSiteBgpStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsBGPPeersApiCountSiteBgpStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **optional.String**|  | 
 **distinct** | **optional.String**|  | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteBgpStats**
> InlineResponse20076 SearchSiteBgpStats(ctx, siteId)
searchSiteBgpStats

Search BGP Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20076**](inline_response_200_76.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

