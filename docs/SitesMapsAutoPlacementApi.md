# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearSiteApAutoOrient**](SitesMapsAutoPlacementApi.md#ClearSiteApAutoOrient) | **Post** /api/v1/sites/{site_id}/maps/{map_id}/clear_auto_orient | clearSiteApAutoOrient
[**ClearSiteApAutoplacement**](SitesMapsAutoPlacementApi.md#ClearSiteApAutoplacement) | **Post** /api/v1/sites/{site_id}/maps/{map_id}/clear_autoplacement | clearSiteApAutoplacement
[**ConfirmSiteApLocalizationData**](SitesMapsAutoPlacementApi.md#ConfirmSiteApLocalizationData) | **Post** /api/v1/sites/{site_id}/maps/{map_id}/use_auto_ap_values | confirmSiteApLocalizationData
[**DeleteSiteApAutoOrientation**](SitesMapsAutoPlacementApi.md#DeleteSiteApAutoOrientation) | **Delete** /api/v1/sites/{site_id}/maps/{map_id}/auto_orient | deleteSiteApAutoOrientation
[**DeleteSiteApAutoplacement**](SitesMapsAutoPlacementApi.md#DeleteSiteApAutoplacement) | **Delete** /api/v1/sites/{site_id}/maps/{map_id}/auto_placement | deleteSiteApAutoplacement
[**GetSiteApAutoPlacement**](SitesMapsAutoPlacementApi.md#GetSiteApAutoPlacement) | **Get** /api/v1/sites/{site_id}/maps/{map_id}/auto_placement | getSiteApAutoPlacement
[**RunSiteApAutoplacement**](SitesMapsAutoPlacementApi.md#RunSiteApAutoplacement) | **Post** /api/v1/sites/{site_id}/maps/{map_id}/auto_placement | runSiteApAutoplacement
[**StartSiteApAutoOrientation**](SitesMapsAutoPlacementApi.md#StartSiteApAutoOrientation) | **Post** /api/v1/sites/{site_id}/maps/{map_id}/auto_orient | startSiteApAutoOrientation

# **ClearSiteApAutoOrient**
> ClearSiteApAutoOrient(ctx, siteId, mapId, optional)
clearSiteApAutoOrient

This API is used to destroy the autoorientations of a map or subset of APs on a map.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mapId** | [**string**](.md)|  | 
 **optional** | ***SitesMapsAutoPlacementApiClearSiteApAutoOrientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesMapsAutoPlacementApiClearSiteApAutoOrientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of MacAddresses**](MacAddresses.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearSiteApAutoplacement**
> ClearSiteApAutoplacement(ctx, siteId, mapId, optional)
clearSiteApAutoplacement

This API is used to destroy the cached autoplacement locations of a map or subset of APs on a map.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mapId** | [**string**](.md)|  | 
 **optional** | ***SitesMapsAutoPlacementApiClearSiteApAutoplacementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesMapsAutoPlacementApiClearSiteApAutoplacementOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of MacAddresses**](MacAddresses.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfirmSiteApLocalizationData**
> ConfirmSiteApLocalizationData(ctx, siteId, mapId, optional)
confirmSiteApLocalizationData

This API is used to accept or reject the cached autoplacement and auto orientation values of a map or subset of APs on a map. A rejected AP will retain its current X,Y and orientation until accpeted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mapId** | [**string**](.md)|  | 
 **optional** | ***SitesMapsAutoPlacementApiConfirmSiteApLocalizationDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesMapsAutoPlacementApiConfirmSiteApLocalizationDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of MapIdUseAutoApValuesBody**](MapIdUseAutoApValuesBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteApAutoOrientation**
> DeleteSiteApAutoOrientation(ctx, mapId, siteId)
deleteSiteApAutoOrientation

This API is called to force stop auto placement for a given map

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mapId** | [**string**](.md)|  | 
  **siteId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteApAutoplacement**
> DeleteSiteApAutoplacement(ctx, siteId, mapId)
deleteSiteApAutoplacement

This API is called to force stop auto placement for a given map

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

# **GetSiteApAutoPlacement**
> InlineResponse200132 GetSiteApAutoPlacement(ctx, siteId, mapId)
getSiteApAutoPlacement

This API is called to view the current status of auto placement for a given map.   #### Status Descriptions  | Status | Description | | --- | --- | | `pending` | Autoplacement has not been requested for this map | | `inprogress` | Autoplacement is currently processing | | `done` | The autoplacement process has completed | | `data_needed` | Additional position data is required for autoplacement. Users should verify the requested anchor APs have a position on the map | | `invalid_model` | Autoplacement is not supported on the model of the APs on the map | | `invalid_version` | Autoplacement is not supported with the APs current firmware version | | `error` | There was an error in the autoplacement process |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mapId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200132**](inline_response_200_132.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RunSiteApAutoplacement**
> RunSiteApAutoplacement(ctx, siteId, mapId, optional)
runSiteApAutoplacement

This API is called to trigger a map for auto placement. For auto placement feature to work, RTT-FTM data need to be collected from the APs on the map. This scan is disruptive and therefore the user must be notified of service disrution during the functioning of auto placement Repeated POST to this endpoint while a map is still running will be rejected.  List of devices to provide suggestions for is an optional parameter that can be given to this API. This will provide autoplacement suggestions only for the devices specified. If no list of devices is provided, all APs asociated with that map are considered by default

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mapId** | [**string**](.md)|  | 
 **optional** | ***SitesMapsAutoPlacementApiRunSiteApAutoplacementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesMapsAutoPlacementApiRunSiteApAutoplacementOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AutoPlacement**](AutoPlacement.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartSiteApAutoOrientation**
> InlineResponse200131 StartSiteApAutoOrientation(ctx, mapId, siteId, optional)
startSiteApAutoOrientation

This API is called to trigger a map for auto orientation. For auto orient feature to work, BLE data needs to be collected from the APs on the map. This precess is not disruptive unlike FTM collection. Repeated POST to this endpoint while a map is still running will be rejected.  List of devices to provide suggestions for is an optional parameter that can be given to this API. This will provide auto orient suggestions only for the devices specified. If no list of devices is provided, all APs asociated with that map are considered by default

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mapId** | [**string**](.md)|  | 
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesMapsAutoPlacementApiStartSiteApAutoOrientationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesMapsAutoPlacementApiStartSiteApAutoOrientationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AutoOrient**](AutoOrient.md)|  | 

### Return type

[**InlineResponse200131**](inline_response_200_131.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

