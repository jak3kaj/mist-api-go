# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteRogueEvents**](SitesRoguesApi.md#CountSiteRogueEvents) | **Get** /api/v1/sites/{site_id}/rogues/events/count | countSiteRogueEvents
[**GetSiteRogueAP**](SitesRoguesApi.md#GetSiteRogueAP) | **Get** /api/v1/sites/{site_id}/rogues/{rogue_bssid} | getSiteRogueAP
[**ListSiteRogueAPs**](SitesRoguesApi.md#ListSiteRogueAPs) | **Get** /api/v1/sites/{site_id}/insights/rogues | listSiteRogueAPs
[**ListSiteRogueClients**](SitesRoguesApi.md#ListSiteRogueClients) | **Get** /api/v1/sites/{site_id}/insights/rogues/clients | listSiteRogueClients
[**SearchSiteRogueEvents**](SitesRoguesApi.md#SearchSiteRogueEvents) | **Get** /api/v1/sites/{site_id}/rogues/events/search | searchSiteRogueEvents

# **CountSiteRogueEvents**
> InlineResponse20016 CountSiteRogueEvents(ctx, siteId, optional)
countSiteRogueEvents

Count Rogue Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesRoguesApiCountSiteRogueEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesRoguesApiCountSiteRogueEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of SiteRogueEventsCountDistinct**](.md)|  | 
 **type_** | [**optional.Interface of RogueType**](.md)|  | 
 **ssid** | **optional.String**| ssid of the network detected as threat | 
 **bssid** | **optional.String**| bssid of the network detected as threat | 
 **apMac** | **optional.String**| mac of the device that had strongest signal strength for ssid/bssid pair | 
 **channel** | **optional.String**| channel over which ap_mac heard ssid/bssid pair | 
 **seenOnLan** | **optional.Bool**| whether the reporting AP see a wireless client (on LAN) connecting to it | 
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

# **GetSiteRogueAP**
> InlineResponse200136 GetSiteRogueAP(ctx, siteId, rogueBssid)
getSiteRogueAP

Get Rogue AP Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **rogueBssid** | **string**|  | 

### Return type

[**InlineResponse200136**](inline_response_200_136.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteRogueAPs**
> InlineResponse200133 ListSiteRogueAPs(ctx, siteId, optional)
listSiteRogueAPs

Get List of Site Rogue/Neighbor APs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesRoguesApiListSiteRogueAPsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesRoguesApiListSiteRogueAPsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of RogueType**](.md)|  | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **interval** | **optional.String**| Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. | 

### Return type

[**InlineResponse200133**](inline_response_200_133.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteRogueClients**
> InlineResponse200134 ListSiteRogueClients(ctx, siteId, optional)
listSiteRogueClients

Get List of Site Rogue Clients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesRoguesApiListSiteRogueClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesRoguesApiListSiteRogueClientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **interval** | **optional.String**| Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. | 

### Return type

[**InlineResponse200134**](inline_response_200_134.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteRogueEvents**
> InlineResponse200135 SearchSiteRogueEvents(ctx, siteId, optional)
searchSiteRogueEvents

Search Rogue Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesRoguesApiSearchSiteRogueEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesRoguesApiSearchSiteRogueEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of RogueType**](.md)|  | 
 **ssid** | **optional.String**| ssid of the network detected as threat | 
 **bssid** | **optional.String**| bssid of the network detected as threat | 
 **apMac** | **optional.String**| mac of the device that had strongest signal strength for ssid/bssid pair | 
 **channel** | **optional.Int32**| channel over which ap_mac heard ssid/bssid pair | 
 **seenOnLan** | **optional.Bool**| whether the reporting AP see a wireless client (on LAN) connecting to it | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse200135**](inline_response_200_135.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

