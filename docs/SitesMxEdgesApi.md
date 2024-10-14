# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteMxEdgeEvents**](SitesMxEdgesApi.md#CountSiteMxEdgeEvents) | **Get** /api/v1/sites/{site_id}/mxedges/events/count | countSiteMxEdgeEvents
[**CreateSiteMxEdge**](SitesMxEdgesApi.md#CreateSiteMxEdge) | **Post** /api/v1/sites/{site_id}/mxedges | createSiteMxEdge
[**DeleteSiteMxEdge**](SitesMxEdgesApi.md#DeleteSiteMxEdge) | **Delete** /api/v1/sites/{site_id}/mxedges/{mxedge_id} | deleteSiteMxEdge
[**GetSiteMxEdge**](SitesMxEdgesApi.md#GetSiteMxEdge) | **Get** /api/v1/sites/{site_id}/mxedges/{mxedge_id} | getSiteMxEdge
[**ListSiteMxEdges**](SitesMxEdgesApi.md#ListSiteMxEdges) | **Get** /api/v1/sites/{site_id}/mxedges | listSiteMxEdges
[**SearchSiteMistEdgeEvents**](SitesMxEdgesApi.md#SearchSiteMistEdgeEvents) | **Get** /api/v1/sites/{site_id}/mxedges/events/search | searchSiteMistEdgeEvents
[**UpdateSiteMxEdge**](SitesMxEdgesApi.md#UpdateSiteMxEdge) | **Put** /api/v1/sites/{site_id}/mxedges/{mxedge_id} | updateSiteMxEdge
[**UploadSiteMxEdgeSupportFiles**](SitesMxEdgesApi.md#UploadSiteMxEdgeSupportFiles) | **Post** /api/v1/sites/{site_id}/mxedges/{mxedge_id}/support | uploadSiteMxEdgeSupportFiles

# **CountSiteMxEdgeEvents**
> InlineResponse20016 CountSiteMxEdgeEvents(ctx, siteId, optional)
countSiteMxEdgeEvents

Count Mist Edge Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesMxEdgesApiCountSiteMxEdgeEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesMxEdgesApiCountSiteMxEdgeEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of SiteMxedgeEventsCountDistinct**](.md)|  | 
 **mxedgeId** | **optional.String**| mist edge id | 
 **mxclusterId** | **optional.String**| mist edge cluster id | 
 **type_** | **optional.String**| see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **service** | **optional.String**| service running on mist edge(mxagent, tunterm etc) | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSiteMxEdge**
> InlineResponse20067 CreateSiteMxEdge(ctx, siteId, optional)
createSiteMxEdge

Create Site Mist Edge

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesMxEdgesApiCreateSiteMxEdgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesMxEdgesApiCreateSiteMxEdgeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Mxedge**](Mxedge.md)|  | 

### Return type

[**InlineResponse20067**](inline_response_200_67.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteMxEdge**
> DeleteSiteMxEdge(ctx, siteId, mxedgeId)
deleteSiteMxEdge

Delete Site Mist Edge

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mxedgeId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteMxEdge**
> GetSiteMxEdge(ctx, siteId, mxedgeId)
getSiteMxEdge

get Site Mist Edge

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mxedgeId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteMxEdges**
> []Mxedge ListSiteMxEdges(ctx, siteId, optional)
listSiteMxEdges

Get List of Site Mist Edges

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesMxEdgesApiListSiteMxEdgesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesMxEdgesApiListSiteMxEdgesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Mxedge**](mxedge.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteMistEdgeEvents**
> InlineResponse20069 SearchSiteMistEdgeEvents(ctx, siteId, optional)
searchSiteMistEdgeEvents

Search Site Mist Edge Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesMxEdgesApiSearchSiteMistEdgeEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesMxEdgesApiSearchSiteMistEdgeEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mxedgeId** | **optional.String**| mist edge id | 
 **mxclusterId** | **optional.String**| mist edge cluster id | 
 **type_** | **optional.String**| see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **service** | **optional.String**| service running on mist edge(mxagent, tunterm etc) | 
 **component** | **optional.String**| component like PS1, PS2 | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]

### Return type

[**InlineResponse20069**](inline_response_200_69.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSiteMxEdge**
> InlineResponse20067 UpdateSiteMxEdge(ctx, siteId, mxedgeId, optional)
updateSiteMxEdge

Update Site Mist Edge settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mxedgeId** | [**string**](.md)|  | 
 **optional** | ***SitesMxEdgesApiUpdateSiteMxEdgeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesMxEdgesApiUpdateSiteMxEdgeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of Mxedge**](Mxedge.md)|  | 

### Return type

[**InlineResponse20067**](inline_response_200_67.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadSiteMxEdgeSupportFiles**
> UploadSiteMxEdgeSupportFiles(ctx, siteId, mxedgeId)
uploadSiteMxEdgeSupportFiles

Support / Upload Mist Edge support files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mxedgeId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

