# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearSiteMlOverwriteForDevice**](SitesLocationApi.md#ClearSiteMlOverwriteForDevice) | **Delete** /api/v1/sites/{site_id}/location/ml/device/{device_id} | clearSiteMlOverwriteForDevice
[**ClearSiteMlOverwriteForMap**](SitesLocationApi.md#ClearSiteMlOverwriteForMap) | **Delete** /api/v1/sites/{site_id}/location/ml/map/{map_id} | clearSiteMlOverwriteForMap
[**GetSiteBeamCoverageOverview**](SitesLocationApi.md#GetSiteBeamCoverageOverview) | **Get** /api/v1/sites/{site_id}/location/coverage | getSiteBeamCoverageOverview
[**GetSiteDefaultPlfForModels**](SitesLocationApi.md#GetSiteDefaultPlfForModels) | **Get** /api/v1/sites/{site_id}/location/ml/defaults | getSiteDefaultPlfForModels
[**GetSiteMachineLearningCurrentStat**](SitesLocationApi.md#GetSiteMachineLearningCurrentStat) | **Get** /api/v1/sites/{site_id}/location/ml/current | getSiteMachineLearningCurrentStat
[**OverwriteSiteMlForDevice**](SitesLocationApi.md#OverwriteSiteMlForDevice) | **Put** /api/v1/sites/{site_id}/location/ml/device/{device_id} | overwriteSiteMlForDevice
[**OverwriteSiteMlForMap**](SitesLocationApi.md#OverwriteSiteMlForMap) | **Put** /api/v1/sites/{site_id}/location/ml/map/{map_id} | overwriteSiteMlForMap
[**ResetSiteMlStatsByMap**](SitesLocationApi.md#ResetSiteMlStatsByMap) | **Post** /api/v1/sites/{site_id}/location/ml/reset/map/{map_id} | resetSiteMlStatsByMap

# **ClearSiteMlOverwriteForDevice**
> ClearSiteMlOverwriteForDevice(ctx, siteId, deviceId)
clearSiteMlOverwriteForDevice

Clear ML Overwrite for Device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearSiteMlOverwriteForMap**
> ClearSiteMlOverwriteForMap(ctx, siteId, mapId)
clearSiteMlOverwriteForMap

Clear ML Overwrite for Map

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

# **GetSiteBeamCoverageOverview**
> InlineResponse200129 GetSiteBeamCoverageOverview(ctx, siteId, optional)
getSiteBeamCoverageOverview

Get Beam Coverage Overview

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesLocationApiGetSiteBeamCoverageOverviewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesLocationApiGetSiteBeamCoverageOverviewOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mapId** | **optional.String**| map_id (filter by map_id) | 
 **type_** | [**optional.Interface of RfClientType**](.md)|  | 
 **clientType** | **optional.String**| client_type (as filter. optional) | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **resolution** | [**optional.Interface of Resolution**](.md)|  | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 

### Return type

[**InlineResponse200129**](inline_response_200_129.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteDefaultPlfForModels**
> []interface{} GetSiteDefaultPlfForModels(ctx, siteId)
getSiteDefaultPlfForModels

Get Default PLF for Models

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteMachineLearningCurrentStat**
> []interface{} GetSiteMachineLearningCurrentStat(ctx, siteId, optional)
getSiteMachineLearningCurrentStat

Get Machine Learning Current Stat For each VBLE AP, it has ML model parameters (e.g. Path-loss-estimate, Intercept) as well as completion indicators (Level and PercentageComplete). For the completeness, ML takes N sample to finish its first level and use N*0.25 samples to complete each successive level. When a device is moved, the completeness will be reset as it has to re-learn.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesLocationApiGetSiteMachineLearningCurrentStatOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesLocationApiGetSiteMachineLearningCurrentStatOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mapId** | **optional.String**| map_id (as filter, optional) | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OverwriteSiteMlForDevice**
> []interface{} OverwriteSiteMlForDevice(ctx, siteId, deviceId, optional)
overwriteSiteMlForDevice

Overwrite ML For Device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***SitesLocationApiOverwriteSiteMlForDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesLocationApiOverwriteSiteMlForDeviceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of map[string]MlOverwriteAdditionalProperties**](map.md)| Request Body | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OverwriteSiteMlForMap**
> []interface{} OverwriteSiteMlForMap(ctx, siteId, mapId, optional)
overwriteSiteMlForMap

Overwrite ML For Map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mapId** | [**string**](.md)|  | 
 **optional** | ***SitesLocationApiOverwriteSiteMlForMapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesLocationApiOverwriteSiteMlForMapOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of map[string]MlOverwriteAdditionalProperties**](map.md)| Request Body | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetSiteMlStatsByMap**
> ResetSiteMlStatsByMap(ctx, siteId, mapId)
resetSiteMlStatsByMap

Reset ML Stats by Map

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

