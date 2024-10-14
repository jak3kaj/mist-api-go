# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteWirelessClientStats**](SitesStatsClientsWirelessApi.md#GetSiteWirelessClientStats) | **Get** /api/v1/sites/{site_id}/stats/clients/{client_mac} | getSiteWirelessClientStats
[**GetSiteWirelessClientsStatsByMap**](SitesStatsClientsWirelessApi.md#GetSiteWirelessClientsStatsByMap) | **Get** /api/v1/sites/{site_id}/stats/maps/{map_id}/clients | getSiteWirelessClientsStatsByMap
[**ListSiteUnconnectedClientStats**](SitesStatsClientsWirelessApi.md#ListSiteUnconnectedClientStats) | **Get** /api/v1/sites/{site_id}/stats/maps/{map_id}/unconnected_clients | listSiteUnconnectedClientStats
[**ListSiteWirelessClientsStats**](SitesStatsClientsWirelessApi.md#ListSiteWirelessClientsStats) | **Get** /api/v1/sites/{site_id}/stats/clients | listSiteWirelessClientsStats

# **GetSiteWirelessClientStats**
> []StatsClient GetSiteWirelessClientStats(ctx, siteId, clientMac, optional)
getSiteWirelessClientStats

Get Site Client Stats Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **clientMac** | **string**|  | 
 **optional** | ***SitesStatsClientsWirelessApiGetSiteWirelessClientStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsClientsWirelessApiGetSiteWirelessClientStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wired** | **optional.Bool**|  | [default to false]

### Return type

[**[]StatsClient**](stats_client.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteWirelessClientsStatsByMap**
> []StatsWirelessClient GetSiteWirelessClientsStatsByMap(ctx, siteId, mapId, optional)
getSiteWirelessClientsStatsByMap

Get Site Clients Stats By Map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mapId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsClientsWirelessApiGetSiteWirelessClientsStatsByMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsClientsWirelessApiGetSiteWirelessClientsStatsByMapOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]StatsWirelessClient**](stats_wireless_client.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteUnconnectedClientStats**
> []StatsUnconnectedClient ListSiteUnconnectedClientStats(ctx, siteId, mapId)
listSiteUnconnectedClientStats

Get List of Site Unconnected Client Location

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mapId** | [**string**](.md)|  | 

### Return type

[**[]StatsUnconnectedClient**](stats_unconnected_client.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteWirelessClientsStats**
> []StatsClient ListSiteWirelessClientsStats(ctx, siteId, optional)
listSiteWirelessClientsStats

Get List of Site All Clients Stats Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsClientsWirelessApiListSiteWirelessClientsStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsClientsWirelessApiListSiteWirelessClientsStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wired** | **optional.Bool**|  | [default to false]
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**[]StatsClient**](stats_client.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

