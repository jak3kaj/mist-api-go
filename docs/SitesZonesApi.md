# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteZoneSessions**](SitesZonesApi.md#CountSiteZoneSessions) | **Get** /api/v1/sites/{site_id}/{zone_type}/count | countSiteZoneSessions
[**CreateSiteZone**](SitesZonesApi.md#CreateSiteZone) | **Post** /api/v1/sites/{site_id}/zones | createSiteZone
[**DeleteSiteZone**](SitesZonesApi.md#DeleteSiteZone) | **Delete** /api/v1/sites/{site_id}/zones/{zone_id} | deleteSiteZone
[**GetSiteZone**](SitesZonesApi.md#GetSiteZone) | **Get** /api/v1/sites/{site_id}/zones/{zone_id} | getSiteZone
[**ListSiteZones**](SitesZonesApi.md#ListSiteZones) | **Get** /api/v1/sites/{site_id}/zones | listSiteZones
[**SearchSiteZoneSessions**](SitesZonesApi.md#SearchSiteZoneSessions) | **Get** /api/v1/sites/{site_id}/{zone_type}/visits/search | searchSiteZoneSessions
[**UpdateSiteZone**](SitesZonesApi.md#UpdateSiteZone) | **Put** /api/v1/sites/{site_id}/zones/{zone_id} | updateSiteZone

# **CountSiteZoneSessions**
> InlineResponse20016 CountSiteZoneSessions(ctx, siteId, zoneType, optional)
countSiteZoneSessions

Count Site Zone Sessions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **zoneType** | [**ZoneType**](.md)|  | 
 **optional** | ***SitesZonesApiCountSiteZoneSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesZonesApiCountSiteZoneSessionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **distinct** | [**optional.Interface of Distinct9**](.md)|  | 
 **userType** | [**optional.Interface of UserType**](.md)| user type | 
 **user** | **optional.String**| client MAC / Asset MAC / SDK UUID | 
 **scopeId** | **optional.String**| if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;map&#x60;/&#x60;zone&#x60;/&#x60;rssizone&#x60;, the scope id | 
 **scope** | [**optional.Interface of Scope1**](.md)| scope | 
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

# **CreateSiteZone**
> InlineResponse200165 CreateSiteZone(ctx, siteId, optional)
createSiteZone

Create Site Zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesZonesApiCreateSiteZoneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesZonesApiCreateSiteZoneOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SiteIdZonesBody**](SiteIdZonesBody.md)| Request Body | 

### Return type

[**InlineResponse200165**](inline_response_200_165.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteZone**
> DeleteSiteZone(ctx, siteId, zoneId)
deleteSiteZone

Delete Site Zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **zoneId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteZone**
> InlineResponse200165 GetSiteZone(ctx, siteId, zoneId)
getSiteZone

Get Site Zone Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **zoneId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200165**](inline_response_200_165.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteZones**
> []Zone ListSiteZones(ctx, siteId, optional)
listSiteZones

Get List of Site Zones

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesZonesApiListSiteZonesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesZonesApiListSiteZonesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Zone**](zone.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteZoneSessions**
> InlineResponse200166 SearchSiteZoneSessions(ctx, siteId, zoneType, optional)
searchSiteZoneSessions

Search Zone Sessions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **zoneType** | [**ZoneType**](.md)|  | 
 **optional** | ***SitesZonesApiSearchSiteZoneSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesZonesApiSearchSiteZoneSessionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userType** | [**optional.Interface of UserType1**](.md)| user type, client (default) / sdkclient / asset | 
 **user** | **optional.String**| client MAC / Asset MAC / SDK UUID | 
 **scopeId** | **optional.String**| if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;map&#x60;/&#x60;zone&#x60;/&#x60;rssizone&#x60;, the scope id | 
 **scope** | [**optional.Interface of Scope2**](.md)| scope | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse200166**](inline_response_200_166.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSiteZone**
> InlineResponse200165 UpdateSiteZone(ctx, siteId, zoneId, optional)
updateSiteZone

Update Site Zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **zoneId** | [**string**](.md)|  | 
 **optional** | ***SitesZonesApiUpdateSiteZoneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesZonesApiUpdateSiteZoneOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ZonesZoneIdBody**](ZonesZoneIdBody.md)| Request Body | 

### Return type

[**InlineResponse200165**](inline_response_200_165.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

