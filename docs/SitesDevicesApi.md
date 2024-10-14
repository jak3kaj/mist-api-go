# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSiteDeviceImage**](SitesDevicesApi.md#AddSiteDeviceImage) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/image{image_number} | addSiteDeviceImage
[**CountSiteDeviceConfigHistory**](SitesDevicesApi.md#CountSiteDeviceConfigHistory) | **Get** /api/v1/sites/{site_id}/devices/config_history/count | countSiteDeviceConfigHistory
[**CountSiteDeviceEvents**](SitesDevicesApi.md#CountSiteDeviceEvents) | **Get** /api/v1/sites/{site_id}/devices/events/count | countSiteDeviceEvents
[**CountSiteDeviceLastConfig**](SitesDevicesApi.md#CountSiteDeviceLastConfig) | **Get** /api/v1/sites/{site_id}/devices/last_config/count | countSiteDeviceLastConfig
[**CountSiteDevices**](SitesDevicesApi.md#CountSiteDevices) | **Get** /api/v1/sites/{site_id}/devices/count | countSiteDevices
[**DeleteSiteDeviceImage**](SitesDevicesApi.md#DeleteSiteDeviceImage) | **Delete** /api/v1/sites/{site_id}/devices/{device_id}/image{image_number} | deleteSiteDeviceImage
[**ExportSiteDevices**](SitesDevicesApi.md#ExportSiteDevices) | **Get** /api/v1/sites/{site_id}/devices/export | exportSiteDevices
[**GetSiteDevice**](SitesDevicesApi.md#GetSiteDevice) | **Get** /api/v1/sites/{site_id}/devices/{device_id} | getSiteDevice
[**ImportSiteDevices**](SitesDevicesApi.md#ImportSiteDevices) | **Post** /api/v1/sites/{site_id}/devices/import | importSiteDevices
[**ListSiteDevices**](SitesDevicesApi.md#ListSiteDevices) | **Get** /api/v1/sites/{site_id}/devices | listSiteDevices
[**RestartSiteMultipleDevices**](SitesDevicesApi.md#RestartSiteMultipleDevices) | **Post** /api/v1/sites/{site_id}/devices/restart | restartSiteMultipleDevices
[**SearchSiteDeviceConfigHistory**](SitesDevicesApi.md#SearchSiteDeviceConfigHistory) | **Get** /api/v1/sites/{site_id}/devices/config_history/search | searchSiteDeviceConfigHistory
[**SearchSiteDeviceEvents**](SitesDevicesApi.md#SearchSiteDeviceEvents) | **Get** /api/v1/sites/{site_id}/devices/events/search | searchSiteDeviceEvents
[**SearchSiteDeviceLastConfigs**](SitesDevicesApi.md#SearchSiteDeviceLastConfigs) | **Get** /api/v1/sites/{site_id}/devices/last_config/search | searchSiteDeviceLastConfigs
[**SearchSiteDevices**](SitesDevicesApi.md#SearchSiteDevices) | **Get** /api/v1/sites/{site_id}/devices/search | searchSiteDevices
[**UpdateSiteDevice**](SitesDevicesApi.md#UpdateSiteDevice) | **Put** /api/v1/sites/{site_id}/devices/{device_id} | updateSiteDevice

# **AddSiteDeviceImage**
> AddSiteDeviceImage(ctx, siteId, deviceId, imageNumber, optional)
addSiteDeviceImage

Attach up to 3 images to a device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
  **imageNumber** | **int32**|  | 
 **optional** | ***SitesDevicesApiAddSiteDeviceImageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesApiAddSiteDeviceImageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **file** | **optional.Interface of *os.File****optional.**|  | 
 **json** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountSiteDeviceConfigHistory**
> InlineResponse20016 CountSiteDeviceConfigHistory(ctx, siteId, optional)
countSiteDeviceConfigHistory

Counts the number of entries in device config history for distinct field with given filters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesApiCountSiteDeviceConfigHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesApiCountSiteDeviceConfigHistoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | **optional.String**|  | 
 **mac** | **optional.String**|  | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountSiteDeviceEvents**
> InlineResponse20016 CountSiteDeviceEvents(ctx, siteId, optional)
countSiteDeviceEvents

Counts the number of entries in ap events history for distinct field with given filters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesApiCountSiteDeviceEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesApiCountSiteDeviceEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of SiteDeviceEventsCountDistinct**](.md)|  | 
 **model** | **optional.String**|  | 
 **type_** | **optional.String**| see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **typeCode** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountSiteDeviceLastConfig**
> InlineResponse20016 CountSiteDeviceLastConfig(ctx, siteId, optional)
countSiteDeviceLastConfig

Counts the number of entries in device config history for distinct field with given filters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesApiCountSiteDeviceLastConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesApiCountSiteDeviceLastConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of SiteDeviceLastConfigCountDistinct**](.md)|  | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountSiteDevices**
> InlineResponse20016 CountSiteDevices(ctx, siteId, optional)
countSiteDevices

Counts the number of entries in ap events history for distinct field with given filters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesApiCountSiteDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesApiCountSiteDevicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of SiteDevicesCountDistinct**](.md)|  | 
 **hostname** | **optional.String**|  | 
 **model** | **optional.String**|  | 
 **mac** | **optional.String**|  | 
 **version** | **optional.String**|  | 
 **mxtunnelStatus** | **optional.String**|  | 
 **mxedgeId** | **optional.String**|  | 
 **lldpSystemName** | **optional.String**|  | 
 **lldpSystemDesc** | **optional.String**|  | 
 **lldpPortId** | **optional.String**|  | 
 **lldpMgmtAddr** | **optional.String**|  | 
 **mapId** | **optional.String**|  | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteDeviceImage**
> DeleteSiteDeviceImage(ctx, siteId, deviceId, imageNumber)
deleteSiteDeviceImage

Delete image from a device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
  **imageNumber** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportSiteDevices**
> *os.File ExportSiteDevices(ctx, siteId)
exportSiteDevices

To download the exported device information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteDevice**
> InlineResponse200122 GetSiteDevice(ctx, siteId, deviceId)
getSiteDevice

Get Device Configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200122**](inline_response_200_122.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportSiteDevices**
> []ConfigDevice ImportSiteDevices(ctx, siteId, optional)
importSiteDevices

Import Information for Multiple Devices  CSV format: ```csv mac,name,map_id,x,y,height,orientation,labels,band_24.power,band_24.bandwidth,band_24.channel,band_24.disabled,band_5.power,band_5.bandwidth,band_5.channel,band_5.disabled,band_6.power,band_6.bandwidth,band_6.channel,band_6.disabled 5c5b53010101,\"AP 1\",845a23bf-bed9-e43c-4c86-6fa474be7ae5,30,10,2.3,45,\"guest, campus, vip\",1,20,0,false,0,40,0,false,17,80,0,false ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesApiImportSiteDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesApiImportSiteDevicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **optional.Interface of *os.File****optional.**|  | 

### Return type

[**[]ConfigDevice**](config_device.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteDevices**
> []ConfigDevice ListSiteDevices(ctx, siteId, optional)
listSiteDevices

Get list of devices on the site.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesApiListSiteDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesApiListSiteDevicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of DeviceTypeWithAll**](.md)|  | 
 **name** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]ConfigDevice**](config_device.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartSiteMultipleDevices**
> RestartSiteMultipleDevices(ctx, siteId, optional)
restartSiteMultipleDevices

Note that only the devices that are connected will be restarted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesApiRestartSiteMultipleDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesApiRestartSiteMultipleDevicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DevicesRestartBody**](DevicesRestartBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteDeviceConfigHistory**
> InlineResponse20049 SearchSiteDeviceConfigHistory(ctx, siteId, optional)
searchSiteDeviceConfigHistory

Search for entries in device config history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesApiSearchSiteDeviceConfigHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesApiSearchSiteDeviceConfigHistoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of DeviceType**](.md)|  | 
 **mac** | **optional.String**| Device MAC Address | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse20049**](inline_response_200_49.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteDeviceEvents**
> InlineResponse200121 SearchSiteDeviceEvents(ctx, siteId, optional)
searchSiteDeviceEvents

Search Devices Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesApiSearchSiteDeviceEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesApiSearchSiteDeviceEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **optional.String**| device mac | 
 **model** | **optional.String**| device model | 
 **text** | **optional.String**| event message | 
 **timestamp** | **optional.String**| event time | 
 **type_** | **optional.String**| see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **lastBy** | **optional.String**| Return last/recent event for passed in field | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse200121**](inline_response_200_121.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteDeviceLastConfigs**
> InlineResponse20049 SearchSiteDeviceLastConfigs(ctx, siteId, optional)
searchSiteDeviceLastConfigs

Search Device Last Configs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesApiSearchSiteDeviceLastConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesApiSearchSiteDeviceLastConfigsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of DeviceType**](.md)|  | 
 **mac** | **optional.String**|  | 
 **version** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse20049**](inline_response_200_49.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteDevices**
> ResponseDeviceSearch SearchSiteDevices(ctx, siteId, optional)
searchSiteDevices

Search Device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesApiSearchSiteDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesApiSearchSiteDevicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostname** | **optional.String**| partial / full hostname | 
 **type_** | [**optional.Interface of DeviceType**](.md)|  | 
 **model** | **optional.String**| device model | 
 **mac** | **optional.String**| device MAC | 
 **version** | **optional.String**| version | 
 **powerConstrained** | **optional.Bool**| power_constrained | 
 **ipAddress** | **optional.String**|  | 
 **mxtunnelStatus** | [**optional.Interface of MxtunnelStatus2**](.md)| MxTunnel status, up / down | 
 **mxedgeId** | [**optional.Interface of string**](.md)| Mist Edge id, if AP is connecting to a Mist Edge | 
 **lldpSystemName** | **optional.String**| LLDP system name | 
 **lldpSystemDesc** | **optional.String**| LLDP system description | 
 **lldpPortId** | **optional.String**| LLDP port id | 
 **lldpMgmtAddr** | **optional.String**| LLDP management ip address | 
 **band24Channel** | **optional.Int32**| Channel of band_24 | 
 **band5Channel** | **optional.Int32**| Channel of band_5 | 
 **band6Channel** | **optional.Int32**| Channel of band_6 | 
 **band24Bandwith** | **optional.Int32**| Bandwidth of band_24 | 
 **band5Bandwith** | **optional.Int32**| Bandwidth of band_5 | 
 **band6Bandwith** | **optional.Int32**| Bandwidth of band_6 | 
 **eth0PortSpeed** | **optional.Int32**| Port speed of eth0 | 
 **sort** | [**optional.Interface of Sort2**](.md)| sort options | 
 **descSort** | [**optional.Interface of DescSort**](.md)| sort options in reverse order | 
 **stats** | **optional.Bool**| whether to return device stats | [default to false]
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**ResponseDeviceSearch**](response_device_search.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSiteDevice**
> InlineResponse200122 UpdateSiteDevice(ctx, siteId, deviceId, optional)
updateSiteDevice

Update Device Configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***SitesDevicesApiUpdateSiteDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesDevicesApiUpdateSiteDeviceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DevicesDeviceIdBody**](DevicesDeviceIdBody.md)| Request Body | 

### Return type

[**InlineResponse200122**](inline_response_200_122.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

