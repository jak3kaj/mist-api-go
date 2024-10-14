# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteRssiZone**](SitesRSSIZonesApi.md#CreateSiteRssiZone) | **Post** /api/v1/sites/{site_id}/rssizones | createSiteRssiZone
[**DeleteSiteRssiZone**](SitesRSSIZonesApi.md#DeleteSiteRssiZone) | **Delete** /api/v1/sites/{site_id}/rssizones/{rssizone_id} | deleteSiteRssiZone
[**GetSiteRssiZone**](SitesRSSIZonesApi.md#GetSiteRssiZone) | **Get** /api/v1/sites/{site_id}/rssizones/{rssizone_id} | getSiteRssiZone
[**ListSiteRssiZones**](SitesRSSIZonesApi.md#ListSiteRssiZones) | **Get** /api/v1/sites/{site_id}/rssizones | listSiteRssiZones
[**UpdateSiteRssiZone**](SitesRSSIZonesApi.md#UpdateSiteRssiZone) | **Put** /api/v1/sites/{site_id}/rssizones/{rssizone_id} | updateSiteRssiZone

# **CreateSiteRssiZone**
> InlineResponse200141 CreateSiteRssiZone(ctx, siteId, optional)
createSiteRssiZone

Create RSSI Zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesRSSIZonesApiCreateSiteRssiZoneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesRSSIZonesApiCreateSiteRssiZoneOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SiteIdRssizonesBody**](SiteIdRssizonesBody.md)| Request Body | 

### Return type

[**InlineResponse200141**](inline_response_200_141.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteRssiZone**
> DeleteSiteRssiZone(ctx, siteId, rssizoneId)
deleteSiteRssiZone

Delete Site RSSI Zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **rssizoneId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteRssiZone**
> []RssiZone GetSiteRssiZone(ctx, siteId, rssizoneId)
getSiteRssiZone

Get Site RSSI Zone details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **rssizoneId** | [**string**](.md)|  | 

### Return type

[**[]RssiZone**](rssi_zone.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteRssiZones**
> []RssiZone ListSiteRssiZones(ctx, siteId, optional)
listSiteRssiZones

Get List of Site RSSI Zone (RSSI-based)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesRSSIZonesApiListSiteRssiZonesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesRSSIZonesApiListSiteRssiZonesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]RssiZone**](rssi_zone.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSiteRssiZone**
> InlineResponse200142 UpdateSiteRssiZone(ctx, siteId, rssizoneId, optional)
updateSiteRssiZone

Update Site RSSI Zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **rssizoneId** | [**string**](.md)|  | 
 **optional** | ***SitesRSSIZonesApiUpdateSiteRssiZoneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesRSSIZonesApiUpdateSiteRssiZoneOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of RssizonesRssizoneIdBody**](RssizonesRssizoneIdBody.md)| Request Body | 

### Return type

[**InlineResponse200142**](inline_response_200_142.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

