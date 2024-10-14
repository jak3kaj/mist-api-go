# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteAsset**](SitesAssetsApi.md#CreateSiteAsset) | **Post** /api/v1/sites/{site_id}/assets | createSiteAsset
[**DeleteSiteAsset**](SitesAssetsApi.md#DeleteSiteAsset) | **Delete** /api/v1/sites/{site_id}/assets/{asset_id} | deleteSiteAsset
[**GetSiteAsset**](SitesAssetsApi.md#GetSiteAsset) | **Get** /api/v1/sites/{site_id}/assets/{asset_id} | getSiteAsset
[**ImportSiteAssets**](SitesAssetsApi.md#ImportSiteAssets) | **Post** /api/v1/sites/{site_id}/assets/import | importSiteAssets
[**ListSiteAssets**](SitesAssetsApi.md#ListSiteAssets) | **Get** /api/v1/sites/{site_id}/assets | listSiteAssets
[**UpdateSiteAsset**](SitesAssetsApi.md#UpdateSiteAsset) | **Put** /api/v1/sites/{site_id}/assets/{asset_id} | updateSiteAsset

# **CreateSiteAsset**
> InlineResponse20032 CreateSiteAsset(ctx, siteId, optional)
createSiteAsset

Create Site Asset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesAssetsApiCreateSiteAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesAssetsApiCreateSiteAssetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SiteIdAssetsBody**](SiteIdAssetsBody.md)| Request Body | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteAsset**
> DeleteSiteAsset(ctx, siteId, assetId)
deleteSiteAsset

Delete Site Asset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **assetId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteAsset**
> InlineResponse20032 GetSiteAsset(ctx, siteId, assetId)
getSiteAsset

Get Site Asset Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **assetId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportSiteAssets**
> ImportSiteAssets(ctx, siteId, optional)
importSiteAssets

Impert Site Assets.   It can be done via a CSV file or a JSON payload.  ## CSV File Format ```csv name,mac \"asset_name\",5c5b53010101 ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesAssetsApiImportSiteAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesAssetsApiImportSiteAssetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **optional.Interface of *os.File****optional.**|  | 
 **upsert** | [**optional.Interface of Upsert**](.md)| API will replace the assets with same mac if provided &#x60;upsert&#x60;&#x3D;&#x3D;&#x60;True&#x60;, otherwise will report in errors in response. | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteAssets**
> []Asset ListSiteAssets(ctx, siteId, optional)
listSiteAssets

Get List of Site Assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesAssetsApiListSiteAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesAssetsApiListSiteAssetsOpts struct
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

# **UpdateSiteAsset**
> InlineResponse20032 UpdateSiteAsset(ctx, siteId, assetId, optional)
updateSiteAsset

Update Site Asset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **assetId** | [**string**](.md)|  | 
 **optional** | ***SitesAssetsApiUpdateSiteAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesAssetsApiUpdateSiteAssetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AssetsAssetIdBody1**](AssetsAssetIdBody1.md)| Request Body | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

