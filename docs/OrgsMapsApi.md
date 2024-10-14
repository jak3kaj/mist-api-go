# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportOrgMapToSite**](OrgsMapsApi.md#ImportOrgMapToSite) | **Post** /api/v1/orgs/{org_id}/sites/{site_name}/maps/import | importOrgMapToSite
[**ImportOrgMaps**](OrgsMapsApi.md#ImportOrgMaps) | **Post** /api/v1/orgs/{org_id}/maps/import | importOrgMaps

# **ImportOrgMapToSite**
> InlineResponse2008 ImportOrgMapToSite(ctx, orgId, siteName, optional)
importOrgMapToSite

Import data from files is a multipart POST which has a file, an optional json, and an optional csv, to create floorplan, assign matching inventory to specific site, place ap if name or mac matches  #### Request  ``` \"json\": a JSON string describing your upload \"file\": a binary file ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **siteName** | **string**|  | 
 **optional** | ***OrgsMapsApiImportOrgMapToSiteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMapsApiImportOrgMapToSiteOpts struct
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

# **ImportOrgMaps**
> InlineResponse2008 ImportOrgMaps(ctx, orgId, optional)
importOrgMaps

Import data from files is a multipart POST which has a file, an optional json, and an optional csv, to create floorplan, assign matching inventory to specific site, place ap if name or mac matches  ### CSV File Format ```csv Vendor AP name,Mist AP Mac US Office AP-2 - 5c:5b:35:00:00:02,5c5b35000002 ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsMapsApiImportOrgMapsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMapsApiImportOrgMapsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoDeviceprofileAssignment** | **optional.**|  | 
 **csv** | **optional.Interface of *os.File****optional.**|  | 
 **file** | **optional.Interface of *os.File****optional.**|  | 
 **json** | [**optional.Interface of MapOrgImportFileJson**](.md)|  | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

