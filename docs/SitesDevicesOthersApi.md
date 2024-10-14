# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteOtherDeviceEvents**](SitesDevicesOthersApi.md#CountSiteOtherDeviceEvents) | **Get** /api/v1/sites/{site_id}/otherdevices/events/count | countSiteOtherDeviceEvents
[**ListSiteOtherDevices**](SitesDevicesOthersApi.md#ListSiteOtherDevices) | **Get** /api/v1/sites/{site_id}/otherdevices | listSiteOtherDevices
[**SearchSiteOtherDeviceEvents**](SitesDevicesOthersApi.md#SearchSiteOtherDeviceEvents) | **Get** /api/v1/sites/{site_id}/otherdevices/events/search | searchSiteOtherDeviceEvents

# **CountSiteOtherDeviceEvents**
> InlineResponse20016 CountSiteOtherDeviceEvents(ctx, siteId, optional)
countSiteOtherDeviceEvents

Count Site OtherDevices Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesOthersApiCountSiteOtherDeviceEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesOthersApiCountSiteOtherDeviceEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of SiteOtherDeviceEventsCountDistinct**](.md)|  | 
 **type_** | **optional.String**| see [listDeviceEventsDefinitions]($e/Constants%20Events/listOtherDeviceEventsDefinitions) | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteOtherDevices**
> []DeviceOther ListSiteOtherDevices(ctx, siteId, optional)
listSiteOtherDevices

Get List of Site other devices (3rd party devices)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesOthersApiListSiteOtherDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesOthersApiListSiteOtherDevicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vendor** | **optional.String**|  | 
 **mac** | **optional.String**|  | 
 **serial** | **optional.String**|  | 
 **model** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]DeviceOther**](device_other.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteOtherDeviceEvents**
> InlineResponse20053 SearchSiteOtherDeviceEvents(ctx, siteId, optional)
searchSiteOtherDeviceEvents

Search Site OtherDevices Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesOthersApiSearchSiteOtherDeviceEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesOthersApiSearchSiteOtherDeviceEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **optional.String**| mac | 
 **deviceMac** | **optional.String**| mac of attached device | 
 **vendor** | **optional.String**| vendor name | 
 **type_** | **optional.String**| see [listDeviceEventsDefinitions]($e/Constants%20Events/listOtherDeviceEventsDefinitions) | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]

### Return type

[**InlineResponse20053**](inline_response_200_53.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

