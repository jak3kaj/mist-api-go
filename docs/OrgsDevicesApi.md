# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgDeviceEvents**](OrgsDevicesApi.md#CountOrgDeviceEvents) | **Get** /api/v1/orgs/{org_id}/devices/events/count | countOrgDeviceEvents
[**CountOrgDeviceLastConfigs**](OrgsDevicesApi.md#CountOrgDeviceLastConfigs) | **Get** /api/v1/orgs/{org_id}/devices/last_config/count | countOrgDeviceLastConfigs
[**CountOrgDevices**](OrgsDevicesApi.md#CountOrgDevices) | **Get** /api/v1/orgs/{org_id}/devices/count | countOrgDevices
[**GetOrgJuniperDevicesCommand**](OrgsDevicesApi.md#GetOrgJuniperDevicesCommand) | **Get** /api/v1/orgs/{org_id}/ocdevices/outbound_ssh_cmd | getOrgJuniperDevicesCommand
[**ListOrgApsMacs**](OrgsDevicesApi.md#ListOrgApsMacs) | **Get** /api/v1/orgs/{org_id}/devices/radio_macs | listOrgApsMacs
[**ListOrgDevices**](OrgsDevicesApi.md#ListOrgDevices) | **Get** /api/v1/orgs/{org_id}/devices | listOrgDevices
[**SearchOrgDeviceEvents**](OrgsDevicesApi.md#SearchOrgDeviceEvents) | **Get** /api/v1/orgs/{org_id}/devices/events/search | searchOrgDeviceEvents
[**SearchOrgDeviceLastConfigs**](OrgsDevicesApi.md#SearchOrgDeviceLastConfigs) | **Get** /api/v1/orgs/{org_id}/devices/last_config/search | searchOrgDeviceLastConfigs
[**SearchOrgDevices**](OrgsDevicesApi.md#SearchOrgDevices) | **Get** /api/v1/orgs/{org_id}/devices/search | searchOrgDevices

# **CountOrgDeviceEvents**
> InlineResponse20016 CountOrgDeviceEvents(ctx, orgId, optional)
countOrgDeviceEvents

Count Org Devices Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsDevicesApiCountOrgDeviceEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsDevicesApiCountOrgDeviceEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of OrgDevicesEventsCountDistinct**](.md)|  | 
 **siteId** | **optional.String**| site id | 
 **ap** | **optional.String**| AP mac | 
 **apfw** | **optional.String**| AP Firmware | 
 **model** | **optional.String**| device model | 
 **text** | **optional.String**| event message | 
 **timestamp** | **optional.String**| event time | 
 **type_** | **optional.String**| see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
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

# **CountOrgDeviceLastConfigs**
> InlineResponse20016 CountOrgDeviceLastConfigs(ctx, orgId, optional)
countOrgDeviceLastConfigs

Counts the number of entries in device config history for distinct field with given filters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsDevicesApiCountOrgDeviceLastConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsDevicesApiCountOrgDeviceLastConfigsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of DeviceType**](.md)|  | 
 **distinct** | [**optional.Interface of OrgDevicesLastConfigsCountDistinct**](.md)|  | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **limit** | **optional.Int32**|  | [default to 100]

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountOrgDevices**
> InlineResponse20016 CountOrgDevices(ctx, orgId, optional)
countOrgDevices

Count Org Devices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsDevicesApiCountOrgDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsDevicesApiCountOrgDevicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of OrgDevicesCountDistinct**](.md)|  | 
 **hostname** | **optional.String**| partial / full hostname | 
 **siteId** | **optional.String**| site id | 
 **model** | **optional.String**| device model | 
 **managed** | **optional.String**| for switches and gateways, to filter on managed/unmanaged devices. enum: &#x60;true&#x60;, &#x60;false&#x60; | 
 **mac** | **optional.String**| AP mac | 
 **version** | **optional.String**| version | 
 **ipAddress** | **optional.String**|  | 
 **mxtunnelStatus** | [**optional.Interface of MxtunnelStatus**](.md)| MxTunnel status, enum: &#x60;up&#x60;, &#x60;down&#x60; | 
 **mxedgeId** | **optional.String**| Mist Edge id, if AP is connecting to a Mist Edge | 
 **lldpSystemName** | **optional.String**| LLDP system name | 
 **lldpSystemDesc** | **optional.String**| LLDP system description | 
 **lldpPortId** | **optional.String**| LLDP port id | 
 **lldpMgmtAddr** | **optional.String**| LLDP management ip address | 
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

# **GetOrgJuniperDevicesCommand**
> InlineResponse20050 GetOrgJuniperDevicesCommand(ctx, orgId, optional)
getOrgJuniperDevicesCommand

Get Org Juniper Devices command  Juniper devices can be managed/adopted by Mist. Currently outbound-ssh + netconf is used. A few lines of CLI commands are generated per-Org, allowing the Juniper devices to phone home to Mist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsDevicesApiGetOrgJuniperDevicesCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsDevicesApiGetOrgJuniperDevicesCommandOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | **optional.String**| site_id would be used for proxy config check of the site and automatic site assignment | 

### Return type

[**InlineResponse20050**](inline_response_200_50.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgApsMacs**
> []ApRadioMac ListOrgApsMacs(ctx, orgId, optional)
listOrgApsMacs

For some scenarios like E911 or security systems, the BSSIDs are required to identify which AP the client is connecting to. Then the location of the AP can be used as the approximate location of the client.  Each radio MAC can have 16 BSSIDs (enumerate the last octet from 0-F)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsDevicesApiListOrgApsMacsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsDevicesApiListOrgApsMacsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]ApRadioMac**](ap_radio_mac.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgDevices**
> InlineResponse20047 ListOrgDevices(ctx, orgId)
listOrgDevices

Get List of Org Devices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20047**](inline_response_200_47.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrgDeviceEvents**
> InlineResponse20048 SearchOrgDeviceEvents(ctx, orgId, optional)
searchOrgDeviceEvents

Search Org Devices Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsDevicesApiSearchOrgDeviceEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsDevicesApiSearchOrgDeviceEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **optional.String**| device mac | 
 **model** | **optional.String**| device model | 
 **deviceType** | [**optional.Interface of DeviceType**](.md)|  | 
 **text** | **optional.String**| event message | 
 **timestamp** | **optional.String**| event time | 
 **type_** | **optional.String**| see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **lastBy** | **optional.String**| Return last/recent event for passed in field | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse20048**](inline_response_200_48.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrgDeviceLastConfigs**
> InlineResponse20049 SearchOrgDeviceLastConfigs(ctx, orgId, optional)
searchOrgDeviceLastConfigs

Search Device Last Configs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsDevicesApiSearchOrgDeviceLastConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsDevicesApiSearchOrgDeviceLastConfigsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of DeviceType**](.md)|  | 
 **mac** | **optional.String**| Device MAC address | 
 **name** | **optional.String**| Devices Name | 
 **version** | **optional.String**| Device Version | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **limit** | **optional.Int32**|  | [default to 100]
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse20049**](inline_response_200_49.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrgDevices**
> ResponseDeviceSearch SearchOrgDevices(ctx, orgId, optional)
searchOrgDevices

Search Org Devices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsDevicesApiSearchOrgDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsDevicesApiSearchOrgDevicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostname** | **optional.String**| partial / full hostname | 
 **siteId** | **optional.String**| site id | 
 **model** | **optional.String**| device model | 
 **mac** | **optional.String**| AP mac | 
 **version** | **optional.String**| version | 
 **powerConstrained** | **optional.Bool**| power_constrained | 
 **ipAddress** | **optional.String**|  | 
 **mxtunnelStatus** | [**optional.Interface of MxtunnelStatus1**](.md)| MxTunnel status, up / down | 
 **mxedgeId** | **optional.String**| Mist Edge id, if AP is connecting to a Mist Edge | 
 **lldpSystemName** | **optional.String**| LLDP system name | 
 **lldpSystemDesc** | **optional.String**| LLDP system description | 
 **lldpPortId** | **optional.String**| LLDP port id | 
 **lldpMgmtAddr** | **optional.String**| LLDP management ip address | 
 **band24Bandwith** | **optional.Int32**| Bandwith of band_24 | 
 **band5Bandwith** | **optional.Int32**| Bandwith of band_5 | 
 **band6Bandwith** | **optional.Int32**| Bandwith of band_6 | 
 **band24Channel** | **optional.Int32**| Channel of band_24 | 
 **band5Channel** | **optional.Int32**| Channel of band_5 | 
 **band6Channel** | **optional.Int32**| Channel of band_6 | 
 **eth0PortSpeed** | **optional.Int32**| Port speed of eth0 | 
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

