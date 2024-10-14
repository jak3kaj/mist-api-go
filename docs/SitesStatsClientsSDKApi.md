# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteSdkStats**](SitesStatsClientsSDKApi.md#GetSiteSdkStats) | **Get** /api/v1/sites/{site_id}/stats/sdkclients/{sdkclient_id} | getSiteSdkStats
[**GetSiteSdkStatsByMap**](SitesStatsClientsSDKApi.md#GetSiteSdkStatsByMap) | **Get** /api/v1/sites/{site_id}/stats/maps/{map_id}/sdkclients | getSiteSdkStatsByMap

# **GetSiteSdkStats**
> InlineResponse200171 GetSiteSdkStats(ctx, siteId, sdkclientId)
getSiteSdkStats

Get Detail Stats of a SdkClient

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **sdkclientId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200171**](inline_response_200_171.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSdkStatsByMap**
> []StatsSdkclient GetSiteSdkStatsByMap(ctx, siteId, mapId)
getSiteSdkStatsByMap

Get SdkClient Stats By Map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mapId** | [**string**](.md)|  | 

### Return type

[**[]StatsSdkclient**](stats_sdkclient.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

