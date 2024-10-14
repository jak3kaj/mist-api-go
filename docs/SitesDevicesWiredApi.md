# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSiteLocalSwitchPortConfig**](SitesDevicesWiredApi.md#DeleteSiteLocalSwitchPortConfig) | **Delete** /api/v1/sites/{site_id}/devices/{device_id}/local_port_config | deleteSiteLocalSwitchPortConfig
[**UpdateSiteLocalSwitchPortConfig**](SitesDevicesWiredApi.md#UpdateSiteLocalSwitchPortConfig) | **Put** /api/v1/sites/{site_id}/devices/{device_id}/local_port_config | updateSiteLocalSwitchPortConfig

# **DeleteSiteLocalSwitchPortConfig**
> DeleteSiteLocalSwitchPortConfig(ctx, siteId, deviceId)
deleteSiteLocalSwitchPortConfig

Sometimes HelpDesk Admin needs to change port configs

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

# **UpdateSiteLocalSwitchPortConfig**
> UpdateSiteLocalSwitchPortConfig(ctx, siteId, deviceId, optional)
updateSiteLocalSwitchPortConfig

Sometimes HelpDesk Admin needs to change port configs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesWiredApiUpdateSiteLocalSwitchPortConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesWiredApiUpdateSiteLocalSwitchPortConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdLocalPortConfigBody**](DeviceIdLocalPortConfigBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

