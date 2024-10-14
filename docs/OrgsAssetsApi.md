# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgAsset**](OrgsAssetsApi.md#CreateOrgAsset) | **Post** /api/v1/orgs/{org_id}/assets | createOrgAsset
[**DeleteOrgAsset**](OrgsAssetsApi.md#DeleteOrgAsset) | **Delete** /api/v1/orgs/{org_id}/assets/{asset_id} | deleteOrgAsset
[**GetOrgAsset**](OrgsAssetsApi.md#GetOrgAsset) | **Get** /api/v1/orgs/{org_id}/assets/{asset_id} | getOrgAsset
[**ImportOrgAssets**](OrgsAssetsApi.md#ImportOrgAssets) | **Post** /api/v1/orgs/{org_id}/assets/import | importOrgAssets
[**ListOrgAssets**](OrgsAssetsApi.md#ListOrgAssets) | **Get** /api/v1/orgs/{org_id}/assets | listOrgAssets
[**UpdateOrgAsset**](OrgsAssetsApi.md#UpdateOrgAsset) | **Put** /api/v1/orgs/{org_id}/assets/{asset_id} | updateOrgAsset

# **CreateOrgAsset**
> InlineResponse20032 CreateOrgAsset(ctx, orgId, optional)
createOrgAsset

Create Org Asset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAssetsApiCreateOrgAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAssetsApiCreateOrgAssetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdAssetsBody**](OrgIdAssetsBody.md)| Request Body | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgAsset**
> DeleteOrgAsset(ctx, orgId, assetId)
deleteOrgAsset

Delete Org Asset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **assetId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgAsset**
> InlineResponse20032 GetOrgAsset(ctx, orgId, assetId)
getOrgAsset

Get Org Asset Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **assetId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportOrgAssets**
> ImportOrgAssets(ctx, orgId, optional)
importOrgAssets

Impert Org Assets.   It can be done via a CSV file or a JSON payload.  #### CSV File Format ```csv name,mac \"asset_name\",5c5b53010101 ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAssetsApiImportOrgAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAssetsApiImportOrgAssetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **optional.Interface of *os.File****optional.**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgAssets**
> []Asset ListOrgAssets(ctx, orgId, optional)
listOrgAssets

Get List of Org Assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAssetsApiListOrgAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAssetsApiListOrgAssetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Asset**](asset.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgAsset**
> InlineResponse20032 UpdateOrgAsset(ctx, orgId, assetId, optional)
updateOrgAsset

Update Org Asset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **assetId** | [**string**](.md)|  | 
 **optional** | ***OrgsAssetsApiUpdateOrgAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAssetsApiUpdateOrgAssetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AssetsAssetIdBody**](AssetsAssetIdBody.md)| Request Body | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

