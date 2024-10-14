# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSiteMapImage**](SitesMapsApi.md#AddSiteMapImage) | **Post** /api/v1/sites/{site_id}/maps/{map_id}/image | addSiteMapImage
[**BulkAssignSiteApsToMap**](SitesMapsApi.md#BulkAssignSiteApsToMap) | **Post** /api/v1/sites/{site_id}/maps/{map_id}/set_map | bulkAssignSiteApsToMap
[**CreateSiteMap**](SitesMapsApi.md#CreateSiteMap) | **Post** /api/v1/sites/{site_id}/maps | createSiteMap
[**DeleteSiteMap**](SitesMapsApi.md#DeleteSiteMap) | **Delete** /api/v1/sites/{site_id}/maps/{map_id} | deleteSiteMap
[**DeleteSiteMapImage**](SitesMapsApi.md#DeleteSiteMapImage) | **Delete** /api/v1/sites/{site_id}/maps/{map_id}/image | deleteSiteMapImage
[**GetSiteMap**](SitesMapsApi.md#GetSiteMap) | **Get** /api/v1/sites/{site_id}/maps/{map_id} | getSiteMap
[**ImportSiteMaps**](SitesMapsApi.md#ImportSiteMaps) | **Post** /api/v1/sites/{site_id}/maps/import | importSiteMaps
[**ImportSiteWayfindings**](SitesMapsApi.md#ImportSiteWayfindings) | **Post** /api/v1/sites/{site_id}/maps/{map_id}/wayfinding/import | importSiteWayfindings
[**ListSiteMaps**](SitesMapsApi.md#ListSiteMaps) | **Get** /api/v1/sites/{site_id}/maps | listSiteMaps
[**ReplaceSiteMapImage**](SitesMapsApi.md#ReplaceSiteMapImage) | **Post** /api/v1/sites/{site_id}/maps/{map_id}/replace | replaceSiteMapImage
[**UpdateSiteMap**](SitesMapsApi.md#UpdateSiteMap) | **Put** /api/v1/sites/{site_id}/maps/{map_id} | updateSiteMap

# **AddSiteMapImage**
> AddSiteMapImage(ctx, siteId, mapId, optional)
addSiteMapImage

Add image map is a multipart POST which has an file (Image) and an optional json parameter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mapId** | [**string**](.md)|  | 
 **optional** | ***SitesMapsApiAddSiteMapImageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesMapsApiAddSiteMapImageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | **optional.Interface of *os.File****optional.**|  | 
 **json** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkAssignSiteApsToMap**
> InlineResponse200130 BulkAssignSiteApsToMap(ctx, siteId, mapId, optional)
bulkAssignSiteApsToMap

This API can be used to assign a list of AP Macs associated with site_id to the specified map_id. Note that map_id must be associated with corresponding site_id. This API obeys the following rules  1. if AP is unassigned to any Map, it gets associated with map_id  2. Any moved APs are returned in the response  3. If the AP is considered a locked AP, no action will be taken

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mapId** | [**string**](.md)|  | 
 **optional** | ***SitesMapsApiBulkAssignSiteApsToMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesMapsApiBulkAssignSiteApsToMapOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of MapIdSetMapBody**](MapIdSetMapBody.md)|  | 

### Return type

[**InlineResponse200130**](inline_response_200_130.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSiteMap**
> InlineResponse2009 CreateSiteMap(ctx, siteId, optional)
createSiteMap

Create Site Map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesMapsApiCreateSiteMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesMapsApiCreateSiteMapOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SiteIdMapsBody**](SiteIdMapsBody.md)| Request Body | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteMap**
> DeleteSiteMap(ctx, siteId, mapId)
deleteSiteMap

Delete Site Map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mapId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteMapImage**
> DeleteSiteMapImage(ctx, siteId, mapId)
deleteSiteMapImage

Delete Site Map Image

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mapId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteMap**
> InlineResponse2009 GetSiteMap(ctx, siteId, mapId)
getSiteMap

Get Site Map Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mapId** | [**string**](.md)|  | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportSiteMaps**
> InlineResponse2008 ImportSiteMaps(ctx, siteId, optional)
importSiteMaps

Import data from files is a multipart POST which has an file, an optional json, and an optional csv, to create floorplan, assign matching inventory to specific site, place ap if name or mac matches.  # Note This endpoint (at the site level), the AP must be already assigned to the site to be placed on the floorplan. If you want to place APs from the Org inventory, it is required to use the endpoint at the Org level [importOrgMaps](#operation/importOrgMaps)  # CSV File Format ```csv Vendor AP name,Mist AP Mac US Office AP-2,5c:5b:35:00:00:02 US Office AP-3,5c5b35000002 ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesMapsApiImportSiteMapsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesMapsApiImportSiteMapsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoDeviceprofileAssignment** | **optional.**|  | 
 **csv** | **optional.Interface of *os.File****optional.**|  | 
 **file** | **optional.Interface of *os.File****optional.**|  | 
 **json** | [**optional.Interface of MapImportJson**](.md)|  | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportSiteWayfindings**
> ImportSiteWayfindings(ctx, siteId, mapId, optional)
importSiteWayfindings

This imports the vendor map meta data into the Map JSON. This is required by the SDK and App in order to access/render the vendor Map properly.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mapId** | [**string**](.md)|  | 
 **optional** | ***SitesMapsApiImportSiteWayfindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesMapsApiImportSiteWayfindingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of WayfindingImportBody**](WayfindingImportBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteMaps**
> []ModelMap ListSiteMaps(ctx, siteId, optional)
listSiteMaps

Get List of Site Maps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesMapsApiListSiteMapsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesMapsApiListSiteMapsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]ModelMap**](map.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceSiteMapImage**
> ReplaceSiteMapImage(ctx, siteId, mapId, optional)
replaceSiteMapImage

Replace Map Image   This works like an PUT where the image will be replaced. If transform is provided, all the locations of the objects on the map (AP, Zone, Vbeacon, Beacon) will be transformed as well (relative to the new Map)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mapId** | [**string**](.md)|  | 
 **optional** | ***SitesMapsApiReplaceSiteMapImageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesMapsApiReplaceSiteMapImageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | **optional.Interface of *os.File****optional.**|  | 
 **json** | [**optional.Interface of MapSiteReplaceFileJson**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSiteMap**
> InlineResponse2009 UpdateSiteMap(ctx, siteId, mapId, optional)
updateSiteMap

Update Site Map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mapId** | [**string**](.md)|  | 
 **optional** | ***SitesMapsApiUpdateSiteMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesMapsApiUpdateSiteMapOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of MapsMapIdBody2**](MapsMapIdBody2.md)| Request Body | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

