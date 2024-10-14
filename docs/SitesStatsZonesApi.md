# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteZoneStats**](SitesStatsZonesApi.md#GetSiteZoneStats) | **Get** /api/v1/sites/{site_id}/stats/{zone_type}/{zone_id} | getSiteZoneStats
[**ListSiteZonesStats**](SitesStatsZonesApi.md#ListSiteZonesStats) | **Get** /api/v1/sites/{site_id}/stats/zones | listSiteZonesStats

# **GetSiteZoneStats**
> InlineResponse200176 GetSiteZoneStats(ctx, siteId, zoneType, zoneId)
getSiteZoneStats

Get Detail Zone Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **zoneType** | [**ZoneType**](.md)|  | 
  **zoneId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200176**](inline_response_200_176.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteZonesStats**
> []StatsZone ListSiteZonesStats(ctx, siteId, optional)
listSiteZonesStats

Get List of Site Zones Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsZonesApiListSiteZonesStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsZonesApiListSiteZonesStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mapId** | **optional.String**|  | 

### Return type

[**[]StatsZone**](stats_zone.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

