# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteAssets**](SitesStatsAssetsApi.md#CountSiteAssets) | **Get** /api/v1/sites/{site_id}/stats/assets/count | countSiteAssets
[**GetSiteAssetStats**](SitesStatsAssetsApi.md#GetSiteAssetStats) | **Get** /api/v1/sites/{site_id}/stats/assets/asset_id | getSiteAssetStats
[**GetSiteAssetsOfInterest**](SitesStatsAssetsApi.md#GetSiteAssetsOfInterest) | **Get** /api/v1/sites/{site_id}/stats/filtered_assets | getSiteAssetsOfInterest
[**GetSiteDiscoveredAssetByMap**](SitesStatsAssetsApi.md#GetSiteDiscoveredAssetByMap) | **Get** /api/v1/sites/{site_id}/stats/maps/{map_id}/discovered_assets | getSiteDiscoveredAssetByMap
[**ListSiteAssetsStats**](SitesStatsAssetsApi.md#ListSiteAssetsStats) | **Get** /api/v1/sites/{site_id}/stats/assets | listSiteAssetsStats
[**ListSiteDiscoveredAssets**](SitesStatsAssetsApi.md#ListSiteDiscoveredAssets) | **Get** /api/v1/sites/{site_id}/stats/discovered_assets | listSiteDiscoveredAssets
[**SearchSiteAssets**](SitesStatsAssetsApi.md#SearchSiteAssets) | **Get** /api/v1/sites/{site_id}/stats/assets/search | searchSiteAssets

# **CountSiteAssets**
> InlineResponse20016 CountSiteAssets(ctx, siteId, optional)
countSiteAssets

Count Asset by distinct field

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsAssetsApiCountSiteAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsAssetsApiCountSiteAssetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of SiteAssetsCountDistinct**](.md)|  | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteAssetStats**
> InlineResponse200168 GetSiteAssetStats(ctx, siteId, optional)
getSiteAssetStats

Get Site Asset Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsAssetsApiGetSiteAssetStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsAssetsApiGetSiteAssetStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse200168**](inline_response_200_168.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteAssetsOfInterest**
> []AssetOfInterest GetSiteAssetsOfInterest(ctx, siteId, optional)
getSiteAssetsOfInterest

Get a list of BLE beacons that matches Asset or AssetFilter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsAssetsApiGetSiteAssetsOfInterestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsAssetsApiGetSiteAssetsOfInterestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]AssetOfInterest**](asset_of_interest.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteDiscoveredAssetByMap**
> []StatsAsset GetSiteDiscoveredAssetByMap(ctx, siteId, mapId)
getSiteDiscoveredAssetByMap

Get a list of BLE beacons that we discovered (whether they’ re defined as assets or not)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mapId** | [**string**](.md)|  | 

### Return type

[**[]StatsAsset**](stats_asset.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteAssetsStats**
> []StatsAsset ListSiteAssetsStats(ctx, siteId, optional)
listSiteAssetsStats

Get List of Site Assets Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsAssetsApiListSiteAssetsStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsAssetsApiListSiteAssetsStatsOpts struct
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

# **ListSiteDiscoveredAssets**
> []Asset ListSiteDiscoveredAssets(ctx, siteId, optional)
listSiteDiscoveredAssets

Get List of Site Discovered BLE Assets that doesn’t match any of the Asset / Assetfilters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsAssetsApiListSiteDiscoveredAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsAssetsApiListSiteDiscoveredAssetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
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

# **SearchSiteAssets**
> InlineResponse20075 SearchSiteAssets(ctx, siteId, optional)
searchSiteAssets

Assets Search

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsAssetsApiSearchSiteAssetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsAssetsApiSearchSiteAssetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **optional.String**|  | 
 **mapId** | **optional.String**|  | 
 **ibeaconUuid** | **optional.String**|  | 
 **ibeaconMajor** | **optional.Int32**|  | 
 **ibeaconMinor** | **optional.Int32**|  | 
 **eddystoneUidNamespace** | **optional.String**|  | 
 **eddystoneUidInstance** | **optional.String**|  | 
 **eddystoneUrl** | **optional.String**|  | 
 **deviceName** | **optional.String**|  | 
 **by** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **apMac** | **optional.String**|  | 
 **beam** | **optional.String**|  | 
 **rssi** | **optional.String**|  | 
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

