# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteWanUsage**](SitesWANUsagesApi.md#CountSiteWanUsage) | **Get** /api/v1/sites/{site_id}/wan_usages/count | countSiteWanUsage
[**SearchSiteWanUsage**](SitesWANUsagesApi.md#SearchSiteWanUsage) | **Get** /api/v1/sites/{site_id}/wan_usages/search | searchSiteWanUsage

# **CountSiteWanUsage**
> InlineResponse20016 CountSiteWanUsage(ctx, siteId, optional)
countSiteWanUsage

Count Site WAN Uages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesWANUsagesApiCountSiteWanUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWANUsagesApiCountSiteWanUsageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **optional.String**| MAC address | 
 **peerMac** | **optional.String**| Peer MAC address | 
 **portId** | **optional.String**| Port ID for the device | 
 **peerPortId** | **optional.String**| Peer Port ID for the device | 
 **policy** | **optional.String**| policy for the wan path | 
 **tenant** | **optional.String**| tenant network in which the packet is sent | 
 **pathType** | **optional.String**| path_type of the port | 
 **distinct** | [**optional.Interface of WanUsagesCountDisctinct**](.md)|  | 
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

# **SearchSiteWanUsage**
> SearchWanUsage SearchSiteWanUsage(ctx, siteId, optional)
searchSiteWanUsage

Search Site WAN Uages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesWANUsagesApiSearchSiteWanUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWANUsagesApiSearchSiteWanUsageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **optional.String**| MAC address | 
 **peerMac** | **optional.String**| Peer MAC address | 
 **portId** | **optional.String**| Port ID for the device | 
 **peerPortId** | **optional.String**| Peer Port ID for the device | 
 **policy** | **optional.String**| policy for the wan path | 
 **tenant** | **optional.String**| tenant network in which the packet is sent | 
 **pathType** | **optional.String**| path_type of the port | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**SearchWanUsage**](search_wan_usage.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

