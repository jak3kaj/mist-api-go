# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteMxEdgeStats**](SitesStatsMxEdgesApi.md#GetSiteMxEdgeStats) | **Get** /api/v1/sites/{site_id}/stats/mxedges/{mxedge_id} | getSiteMxEdgeStats
[**ListSiteMxEdgesStats**](SitesStatsMxEdgesApi.md#ListSiteMxEdgesStats) | **Get** /api/v1/sites/{site_id}/stats/mxedges | listSiteMxEdgesStats

# **GetSiteMxEdgeStats**
> InlineResponse20077 GetSiteMxEdgeStats(ctx, siteId, mxedgeId, optional)
getSiteMxEdgeStats

Get One Site MxEdge Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mxedgeId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsMxEdgesApiGetSiteMxEdgeStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsMxEdgesApiGetSiteMxEdgeStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse20077**](inline_response_200_77.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteMxEdgesStats**
> []StatsMxedge ListSiteMxEdgesStats(ctx, siteId, optional)
listSiteMxEdgesStats

Get List of Site MxEdges Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsMxEdgesApiListSiteMxEdgesStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsMxEdgesApiListSiteMxEdgesStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]StatsMxedge**](stats_mxedge.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

