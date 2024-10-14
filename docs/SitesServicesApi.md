# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteServicePathEvents**](SitesServicesApi.md#CountSiteServicePathEvents) | **Get** /api/v1/sites/{site_id}/services/events/count | countSiteServicePathEvents
[**ListSiteServicesDerived**](SitesServicesApi.md#ListSiteServicesDerived) | **Get** /api/v1/sites/{site_id}/services/derived | listSiteServicesDerived
[**SearchSiteServicePathEvents**](SitesServicesApi.md#SearchSiteServicePathEvents) | **Get** /api/v1/sites/{site_id}/services/events/search | searchSiteServicePathEvents

# **CountSiteServicePathEvents**
> InlineResponse20016 CountSiteServicePathEvents(ctx, siteId, optional)
countSiteServicePathEvents

Count Service Path Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesServicesApiCountSiteServicePathEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesServicesApiCountSiteServicePathEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of SiteServiceEventsCountDistinct**](.md)|  | 
 **type_** | **optional.String**| Event type, e.g. GW_SERVICE_PATH_DOWN | 
 **text** | **optional.String**| Description of the event including the reason it is triggered | 
 **vpnName** | **optional.String**| Peer name | 
 **vpnPath** | **optional.String**| Peer path name | 
 **policy** | **optional.String**| Service policy associated with that specific path | 
 **portId** | **optional.String**| Network interface | 
 **model** | **optional.String**| Device model | 
 **version** | **optional.String**| Device firmware version | 
 **timestamp** | **optional.Float64**| Start time, in epoch | 
 **mac** | **optional.String**| MAC address | 
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

# **ListSiteServicesDerived**
> []Service ListSiteServicesDerived(ctx, siteId, optional)
listSiteServicesDerived

Retrieves the list of Services available for the Site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesServicesApiListSiteServicesDerivedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesServicesApiListSiteServicesDerivedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resolve** | **optional.Bool**| whether resolve the site variables | [default to false]

### Return type

[**[]Service**](service.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteServicePathEvents**
> InlineResponse200143 SearchSiteServicePathEvents(ctx, siteId, optional)
searchSiteServicePathEvents

Search Service Path Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesServicesApiSearchSiteServicePathEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesServicesApiSearchSiteServicePathEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| Event type, e.g. GW_SERVICE_PATH_DOWN | 
 **text** | **optional.String**| Description of the event including the reason it is triggered | 
 **peerPortId** | **optional.String**| Port ID of the peer gateway | 
 **peerMac** | **optional.String**| MAC address of the peer gateway | 
 **vpnName** | **optional.String**| Peer name | 
 **vpnPath** | **optional.String**| Peer path name | 
 **policy** | **optional.String**| Service policy associated with that specific path | 
 **portId** | **optional.String**| Network interface | 
 **model** | **optional.String**| Device model | 
 **version** | **optional.String**| Device firmware version | 
 **timestamp** | **optional.Float64**| Start time, in epoch | 
 **mac** | **optional.String**| MAC address | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]

### Return type

[**InlineResponse200143**](inline_response_200_143.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

