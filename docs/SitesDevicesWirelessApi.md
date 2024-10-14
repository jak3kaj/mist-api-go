# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteDeviceIotPort**](SitesDevicesWirelessApi.md#GetSiteDeviceIotPort) | **Get** /api/v1/sites/{site_id}/devices/{device_id}/iot | getSiteDeviceIotPort
[**GetSiteDeviceRadioChannels**](SitesDevicesWirelessApi.md#GetSiteDeviceRadioChannels) | **Get** /api/v1/sites/{site_id}/devices/ap_channels | getSiteDeviceRadioChannels
[**SetSiteDeviceIotPort**](SitesDevicesWirelessApi.md#SetSiteDeviceIotPort) | **Put** /api/v1/sites/{site_id}/devices/{device_id}/iot | setSiteDeviceIotPort

# **GetSiteDeviceIotPort**
> map[string]int32 GetSiteDeviceIotPort(ctx, siteId, deviceId)
getSiteDeviceIotPort

Returns the current state of each enabled IoT pin configured as an output.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 

### Return type

**map[string]int32**

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteDeviceRadioChannels**
> InlineResponse200123 GetSiteDeviceRadioChannels(ctx, siteId, optional)
getSiteDeviceRadioChannels

Get a list of allowed channels (per channel width)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesWirelessApiGetSiteDeviceRadioChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesWirelessApiGetSiteDeviceRadioChannelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **countryCode** | **optional.String**| country code for the site (for AP config generation), in [two-character](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) | 

### Return type

[**InlineResponse200123**](inline_response_200_123.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSiteDeviceIotPort**
> map[string]int32 SetSiteDeviceIotPort(ctx, siteId, deviceId, optional)
setSiteDeviceIotPort

**Note**: For each IoT pin referenced:  * The pin must be enabled using the Device `iot_config` API  * The pin must support the output direction

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesWirelessApiSetSiteDeviceIotPortOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesWirelessApiSetSiteDeviceIotPortOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of map[string]int32**](map.md)| Request Body | 

### Return type

**map[string]int32**

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

