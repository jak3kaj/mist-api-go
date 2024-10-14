# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgAssetsByDistanceField**](OrgsStatsAssetsApi.md#CountOrgAssetsByDistanceField) | **Get** /api/v1/orgs/{org_id}/stats/assets/count | countOrgAssetsByDistanceField
[**ListOrgAssetsStats**](OrgsStatsAssetsApi.md#ListOrgAssetsStats) | **Get** /api/v1/orgs/{org_id}/stats/assets | listOrgAssetsStats
[**SearchOrgAssets**](OrgsStatsAssetsApi.md#SearchOrgAssets) | **Get** /api/v1/orgs/{org_id}/stats/assets/search | searchOrgAssets

# **CountOrgAssetsByDistanceField**
> InlineResponse20016 CountOrgAssetsByDistanceField(ctx, orgId, optional)
countOrgAssetsByDistanceField

Count Org Assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsStatsAssetsApiCountOrgAssetsByDistanceFieldOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsStatsAssetsApiCountOrgAssetsByDistanceFieldOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of OrgAssetCountDistinct**](.md)|  | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgAssetsStats**
> []StatsAsset ListOrgAssetsStats(ctx, orgId, optional)
listOrgAssetsStats

Get List of Org Assets Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsStatsAssetsApiListOrgAssetsStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsStatsAssetsApiListOrgAssetsStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]StatsAsset**](stats_asset.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrgAssets**
> InlineResponse20075 SearchOrgAssets(ctx, orgId, optional)
searchOrgAssets

Search for Org Assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsStatsAssetsApiSearchOrgAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsStatsAssetsApiSearchOrgAssetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | **optional.String**|  | 
 **mac** | **optional.String**|  | 
 **deviceName** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **mapId** | **optional.String**|  | 
 **ibeaconUuid** | **optional.String**|  | 
 **ibeaconMajor** | **optional.String**|  | 
 **ibeaconMinor** | **optional.String**|  | 
 **eddystoneUidNamespace** | **optional.String**|  | 
 **eddystoneUidInstance** | **optional.String**|  | 
 **eddystoneUrl** | **optional.String**|  | 
 **apMac** | **optional.String**|  | 
 **beam** | **optional.Int32**|  | 
 **rssi** | **optional.Int32**|  | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse20075**](inline_response_200_75.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

