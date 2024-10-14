# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgMxEdgeStats**](OrgsStatsMxEdgesApi.md#GetOrgMxEdgeStats) | **Get** /api/v1/orgs/{org_id}/stats/mxedges/{mxedge_id} | getOrgMxEdgeStats
[**ListOrgMxEdgesStats**](OrgsStatsMxEdgesApi.md#ListOrgMxEdgesStats) | **Get** /api/v1/orgs/{org_id}/stats/mxedges | listOrgMxEdgesStats

# **GetOrgMxEdgeStats**
> InlineResponse20077 GetOrgMxEdgeStats(ctx, orgId, mxedgeId, optional)
getOrgMxEdgeStats

Get Org MxEdge Details Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **mxedgeId** | [**string**](.md)|  | 
 **optional** | ***OrgsStatsMxEdgesApiGetOrgMxEdgeStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsStatsMxEdgesApiGetOrgMxEdgeStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forSite** | **optional.Bool**|  | [default to false]

### Return type

[**InlineResponse20077**](inline_response_200_77.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgMxEdgesStats**
> []StatsMxedge ListOrgMxEdgesStats(ctx, orgId, optional)
listOrgMxEdgesStats

Get List of Org MxEdge Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsStatsMxEdgesApiListOrgMxEdgesStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsStatsMxEdgesApiListOrgMxEdgesStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forSite** | **optional.Bool**| filter for site level mist edges | [default to false]
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

