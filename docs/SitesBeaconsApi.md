# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteBeacon**](SitesBeaconsApi.md#CreateSiteBeacon) | **Post** /api/v1/sites/{site_id}/beacons | createSiteBeacon
[**DeleteSiteBeacons**](SitesBeaconsApi.md#DeleteSiteBeacons) | **Delete** /api/v1/sites/{site_id}/beacons/{beacon_id} | deleteSiteBeacons
[**GetSiteBeacon**](SitesBeaconsApi.md#GetSiteBeacon) | **Get** /api/v1/sites/{site_id}/beacons/{beacon_id} | getSiteBeacon
[**ListSiteBeacons**](SitesBeaconsApi.md#ListSiteBeacons) | **Get** /api/v1/sites/{site_id}/beacons | listSiteBeacons
[**UpdateSiteBeacons**](SitesBeaconsApi.md#UpdateSiteBeacons) | **Put** /api/v1/sites/{site_id}/beacons/{beacon_id} | updateSiteBeacons

# **CreateSiteBeacon**
> InlineResponse200118 CreateSiteBeacon(ctx, siteId, optional)
createSiteBeacon

Create Site Beacon

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesBeaconsApiCreateSiteBeaconOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesBeaconsApiCreateSiteBeaconOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SiteIdBeaconsBody**](SiteIdBeaconsBody.md)| Request Body | 

### Return type

[**InlineResponse200118**](inline_response_200_118.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteBeacons**
> DeleteSiteBeacons(ctx, siteId, beaconId)
deleteSiteBeacons

Delete Site Beacon

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **beaconId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteBeacon**
> InlineResponse200118 GetSiteBeacon(ctx, siteId, beaconId)
getSiteBeacon

Get Site Beacon Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **beaconId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200118**](inline_response_200_118.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteBeacons**
> []Beacon ListSiteBeacons(ctx, siteId, optional)
listSiteBeacons

Get List of Site Beacons

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesBeaconsApiListSiteBeaconsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesBeaconsApiListSiteBeaconsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Beacon**](beacon.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSiteBeacons**
> InlineResponse200118 UpdateSiteBeacons(ctx, siteId, beaconId, optional)
updateSiteBeacons

Update Site Beacon

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **beaconId** | [**string**](.md)|  | 
 **optional** | ***SitesBeaconsApiUpdateSiteBeaconsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesBeaconsApiUpdateSiteBeaconsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of BeaconsBeaconIdBody**](BeaconsBeaconIdBody.md)| Request Body | 

### Return type

[**InlineResponse200118**](inline_response_200_118.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

