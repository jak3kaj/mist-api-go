# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgAssetFilters**](OrgsAssetFiltersApi.md#CreateOrgAssetFilters) | **Post** /api/v1/orgs/{org_id}/assetfilters | createOrgAssetFilters
[**DeleteOrgAssetFilter**](OrgsAssetFiltersApi.md#DeleteOrgAssetFilter) | **Delete** /api/v1/orgs/{org_id}/assetfilters/{assetfilter_id} | deleteOrgAssetFilter
[**GetOrgAssetFilter**](OrgsAssetFiltersApi.md#GetOrgAssetFilter) | **Get** /api/v1/orgs/{org_id}/assetfilters/{assetfilter_id} | getOrgAssetFilter
[**ListOrgAssetFilters**](OrgsAssetFiltersApi.md#ListOrgAssetFilters) | **Get** /api/v1/orgs/{org_id}/assetfilters | listOrgAssetFilters
[**UpdateOrgAssetFilters**](OrgsAssetFiltersApi.md#UpdateOrgAssetFilters) | **Put** /api/v1/orgs/{org_id}/assetfilters/{assetfilter_id} | updateOrgAssetFilters

# **CreateOrgAssetFilters**
> InlineResponse20033 CreateOrgAssetFilters(ctx, orgId, optional)
createOrgAssetFilters

Create Asset Filter  Creates a single BLE asset filter for the given site. Any subset of filter properties can be included in the filter. A matching asset must meet the conditions of all given filter properties (logical ‘AND’).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAssetFiltersApiCreateOrgAssetFiltersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAssetFiltersApiCreateOrgAssetFiltersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdAssetfiltersBody**](OrgIdAssetfiltersBody.md)|  | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgAssetFilter**
> DeleteOrgAssetFilter(ctx, orgId, assetfilterId)
deleteOrgAssetFilter

Deletes an existing BLE asset filter for the given site.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **assetfilterId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgAssetFilter**
> InlineResponse20033 GetOrgAssetFilter(ctx, orgId, assetfilterId)
getOrgAssetFilter

Get Org Asset Filter Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **assetfilterId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgAssetFilters**
> []AssetFilter ListOrgAssetFilters(ctx, orgId, optional)
listOrgAssetFilters

Get List of Org BLE asset filters.  Each asset filter in the list operates independently. For a filter object to match an asset, all of the filter properties must match (logical ‘AND’ of each filter property). For example, the “Visitor Tags” filter below will match an asset when both the “ibeacon\\_uuid” and “ibeacon_major” properties match the asset. All non-matching assets are ignored.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAssetFiltersApiListOrgAssetFiltersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAssetFiltersApiListOrgAssetFiltersOpts struct
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

# **UpdateOrgAssetFilters**
> InlineResponse20033 UpdateOrgAssetFilters(ctx, orgId, assetfilterId, optional)
updateOrgAssetFilters

Updates an existing BLE asset filter for the given site.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **assetfilterId** | [**string**](.md)|  | 
 **optional** | ***OrgsAssetFiltersApiUpdateOrgAssetFiltersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAssetFiltersApiUpdateOrgAssetFiltersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AssetfiltersAssetfilterIdBody**](AssetfiltersAssetfilterIdBody.md)| Request Body | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

