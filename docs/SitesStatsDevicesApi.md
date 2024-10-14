# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteAllClientsStatsByDevice**](SitesStatsDevicesApi.md#GetSiteAllClientsStatsByDevice) | **Get** /api/v1/sites/{site_id}/stats/devices/{device_id}/clients | getSiteAllClientsStatsByDevice
[**GetSiteDeviceStats**](SitesStatsDevicesApi.md#GetSiteDeviceStats) | **Get** /api/v1/sites/{site_id}/stats/devices/{device_id} | getSiteDeviceStats
[**GetSiteGatewayMetrics**](SitesStatsDevicesApi.md#GetSiteGatewayMetrics) | **Get** /api/v1/sites/{site_id}/stats/gateways/metrics | getSiteGatewayMetrics
[**GetSiteSwitchesMetrics**](SitesStatsDevicesApi.md#GetSiteSwitchesMetrics) | **Get** /api/v1/sites/{site_id}/stats/switches/metrics | getSiteSwitchesMetrics
[**ListSiteDevicesStats**](SitesStatsDevicesApi.md#ListSiteDevicesStats) | **Get** /api/v1/sites/{site_id}/stats/devices | listSiteDevicesStats

# **GetSiteAllClientsStatsByDevice**
> []StatsWirelessClient GetSiteAllClientsStatsByDevice(ctx, siteId, deviceId)
getSiteAllClientsStatsByDevice

Get wireless client stat by Device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 

### Return type

[**[]StatsWirelessClient**](stats_wireless_client.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteDeviceStats**
> InlineResponse200172 GetSiteDeviceStats(ctx, siteId, deviceId, optional)
getSiteDeviceStats

Get Site Device Stats Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsDevicesApiGetSiteDeviceStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsDevicesApiGetSiteDeviceStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **optional.String**| list of additional fields requests, comma separeted, or &#x60;fields&#x3D;*&#x60; for all of them | 

### Return type

[**InlineResponse200172**](inline_response_200_172.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteGatewayMetrics**
> InlineResponse200173 GetSiteGatewayMetrics(ctx, siteId)
getSiteGatewayMetrics

Get Site Gateway Metrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200173**](inline_response_200_173.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSwitchesMetrics**
> InlineResponse200174 GetSiteSwitchesMetrics(ctx, siteId, optional)
getSiteSwitchesMetrics

Get version compliance metrics for managed or monitored switches

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsDevicesApiGetSiteSwitchesMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsDevicesApiGetSiteSwitchesMetricsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of SwitchMetricType**](.md)|  | 
 **scope** | [**optional.Interface of SwitchMetricScope**](.md)|  | 
 **switchMac** | **optional.String**| switch mac, used only with metric &#x60;type&#x60;&#x3D;&#x3D;&#x60;active_ports_summary&#x60; | 

### Return type

[**InlineResponse200174**](inline_response_200_174.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteDevicesStats**
> []StatsDevice ListSiteDevicesStats(ctx, siteId, optional)
listSiteDevicesStats

Get List of Site Devices Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsDevicesApiListSiteDevicesStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsDevicesApiListSiteDevicesStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of DeviceTypeWithAll**](.md)|  | 
 **status** | [**optional.Interface of StatDeviceStatusFilter**](.md)|  | 
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]StatsDevice**](stats_device.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

