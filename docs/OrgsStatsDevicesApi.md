# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListOrgDevicesStats**](OrgsStatsDevicesApi.md#ListOrgDevicesStats) | **Get** /api/v1/orgs/{org_id}/stats/devices | listOrgDevicesStats

# **ListOrgDevicesStats**
> []StatsDevice ListOrgDevicesStats(ctx, orgId, optional)
listOrgDevicesStats

Get List of Org Devices stats This API renders some high-level device stats, pagination is assumed and returned in response header (as the response is an array)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsStatsDevicesApiListOrgDevicesStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsStatsDevicesApiListOrgDevicesStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of DeviceTypeWithAll**](.md)|  | 
 **status** | [**optional.Interface of DeviceStatus**](.md)|  | 
 **siteId** | **optional.String**|  | 
 **mac** | **optional.String**|  | 
 **evpntopoId** | **optional.String**| EVPN Topology ID | 
 **evpnUnused** | **optional.String**| if &#x60;evpn_unused&#x60;&#x3D;&#x3D;&#x60;true&#x60;, find EVPN eligible switches which don’t belong to any EVPN Topology yet | 
 **fields** | **optional.String**| list of additional fields requests, comma separeted, or &#x60;fields&#x3D;*&#x60; for all of them | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
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

