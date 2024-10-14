# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteVBeacon**](SitesVBeaconsApi.md#CreateSiteVBeacon) | **Post** /api/v1/sites/{site_id}/vbeacons | createSiteVBeacon
[**DeleteSiteVBeacon**](SitesVBeaconsApi.md#DeleteSiteVBeacon) | **Delete** /api/v1/sites/{site_id}/vbeacons/{vbeacon_id} | deleteSiteVBeacon
[**GetSiteVBeacon**](SitesVBeaconsApi.md#GetSiteVBeacon) | **Get** /api/v1/sites/{site_id}/vbeacons/{vbeacon_id} | getSiteVBeacon
[**ListSiteVBeacons**](SitesVBeaconsApi.md#ListSiteVBeacons) | **Get** /api/v1/sites/{site_id}/vbeacons | listSiteVBeacons
[**UpdateSiteVBeacon**](SitesVBeaconsApi.md#UpdateSiteVBeacon) | **Put** /api/v1/sites/{site_id}/vbeacons/{vbeacon_id} | updateSiteVBeacon

# **CreateSiteVBeacon**
> InlineResponse200164 CreateSiteVBeacon(ctx, siteId, optional)
createSiteVBeacon

Create Virtual Beacon

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesVBeaconsApiCreateSiteVBeaconOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesVBeaconsApiCreateSiteVBeaconOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SiteIdVbeaconsBody**](SiteIdVbeaconsBody.md)| Request Body | 

### Return type

[**InlineResponse200164**](inline_response_200_164.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteVBeacon**
> DeleteSiteVBeacon(ctx, siteId, vbeaconId)
deleteSiteVBeacon

Delete Site Virtual Beacon

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **vbeaconId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteVBeacon**
> InlineResponse200164 GetSiteVBeacon(ctx, siteId, vbeaconId)
getSiteVBeacon

Get Site Virtual Beacon Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **vbeaconId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200164**](inline_response_200_164.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteVBeacons**
> []Vbeacon ListSiteVBeacons(ctx, siteId, optional)
listSiteVBeacons

Get List of Site Virtual Beacons

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesVBeaconsApiListSiteVBeaconsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesVBeaconsApiListSiteVBeaconsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Vbeacon**](vbeacon.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSiteVBeacon**
> InlineResponse200164 UpdateSiteVBeacon(ctx, siteId, vbeaconId, optional)
updateSiteVBeacon

Update Site Virtual Beacon

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **vbeaconId** | [**string**](.md)|  | 
 **optional** | ***SitesVBeaconsApiUpdateSiteVBeaconOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesVBeaconsApiUpdateSiteVBeaconOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of VbeaconsVbeaconIdBody**](VbeaconsVbeaconIdBody.md)| Request Body | 

### Return type

[**InlineResponse200164**](inline_response_200_164.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

