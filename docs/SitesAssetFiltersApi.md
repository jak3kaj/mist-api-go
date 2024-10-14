# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteAssetFilters**](SitesAssetFiltersApi.md#CreateSiteAssetFilters) | **Post** /api/v1/sites/{site_id}/assetfilters | createSiteAssetFilters
[**DeleteSiteAssetFilter**](SitesAssetFiltersApi.md#DeleteSiteAssetFilter) | **Delete** /api/v1/sites/{site_id}/assetfilters/{assetfilter_id} | deleteSiteAssetFilter
[**GetSiteAssetFilter**](SitesAssetFiltersApi.md#GetSiteAssetFilter) | **Get** /api/v1/sites/{site_id}/assetfilters/{assetfilter_id} | getSiteAssetFilter
[**ListSiteAssetFilters**](SitesAssetFiltersApi.md#ListSiteAssetFilters) | **Get** /api/v1/sites/{site_id}/assetfilters | listSiteAssetFilters
[**UpdateSiteAssetFilter**](SitesAssetFiltersApi.md#UpdateSiteAssetFilter) | **Put** /api/v1/sites/{site_id}/assetfilters/{assetfilter_id} | updateSiteAssetFilter

# **CreateSiteAssetFilters**
> InlineResponse20033 CreateSiteAssetFilters(ctx, siteId, optional)
createSiteAssetFilters

Create Site Asset Filter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesAssetFiltersApiCreateSiteAssetFiltersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesAssetFiltersApiCreateSiteAssetFiltersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SiteIdAssetfiltersBody**](SiteIdAssetfiltersBody.md)| Request Body | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteAssetFilter**
> DeleteSiteAssetFilter(ctx, siteId, assetfilterId)
deleteSiteAssetFilter

Deletes an existing BLE asset filter for the given site.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **assetfilterId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteAssetFilter**
> InlineResponse20033 GetSiteAssetFilter(ctx, siteId, assetfilterId)
getSiteAssetFilter

Get Site Asset Filter Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **assetfilterId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteAssetFilters**
> []AssetFilter ListSiteAssetFilters(ctx, siteId, optional)
listSiteAssetFilters

Get List of Site Asset Filters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesAssetFiltersApiListSiteAssetFiltersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesAssetFiltersApiListSiteAssetFiltersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]AssetFilter**](asset_filter.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSiteAssetFilter**
> InlineResponse20033 UpdateSiteAssetFilter(ctx, siteId, assetfilterId, optional)
updateSiteAssetFilter

Updates an existing BLE asset filter for the given site.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **assetfilterId** | [**string**](.md)|  | 
 **optional** | ***SitesAssetFiltersApiUpdateSiteAssetFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesAssetFiltersApiUpdateSiteAssetFilterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AssetfiltersAssetfilterIdBody1**](AssetfiltersAssetfilterIdBody1.md)| Request Body | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

