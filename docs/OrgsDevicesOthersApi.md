# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgOtherDeviceEvents**](OrgsDevicesOthersApi.md#CountOrgOtherDeviceEvents) | **Get** /api/v1/orgs/{org_id}/otherdevices/events/count | countOrgOtherDeviceEvents
[**DeleteOrgOtherDevice**](OrgsDevicesOthersApi.md#DeleteOrgOtherDevice) | **Delete** /api/v1/orgs/{org_id}/otherdevices/{device_mac} | deleteOrgOtherDevice
[**GetOrgOtherDevice**](OrgsDevicesOthersApi.md#GetOrgOtherDevice) | **Get** /api/v1/orgs/{org_id}/otherdevices/{device_mac} | getOrgOtherDevice
[**ListOrgOtherDevices**](OrgsDevicesOthersApi.md#ListOrgOtherDevices) | **Get** /api/v1/orgs/{org_id}/otherdevices | listOrgOtherDevices
[**RebootOrgOtherDevice**](OrgsDevicesOthersApi.md#RebootOrgOtherDevice) | **Post** /api/v1/orgs/{org_id}/otherdevices/{device_mac}/reboot | rebootOrgOtherDevice
[**SearchOrgOtherDeviceEvents**](OrgsDevicesOthersApi.md#SearchOrgOtherDeviceEvents) | **Get** /api/v1/orgs/{org_id}/otherdevices/events/search | searchOrgOtherDeviceEvents
[**UpdateOrgOtherDevice**](OrgsDevicesOthersApi.md#UpdateOrgOtherDevice) | **Put** /api/v1/orgs/{org_id}/otherdevices/{device_mac} | updateOrgOtherDevice
[**UpdateOrgOtherDevices**](OrgsDevicesOthersApi.md#UpdateOrgOtherDevices) | **Put** /api/v1/orgs/{org_id}/otherdevices | updateOrgOtherDevices

# **CountOrgOtherDeviceEvents**
> InlineResponse20016 CountOrgOtherDeviceEvents(ctx, orgId, optional)
countOrgOtherDeviceEvents

Count Org OtherDevices Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsDevicesOthersApiCountOrgOtherDeviceEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsDevicesOthersApiCountOrgOtherDeviceEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of OrgOtherdevicesEventsCountDistinct**](.md)|  | 
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

# **DeleteOrgOtherDevice**
> DeleteOrgOtherDevice(ctx, orgId, deviceMac)
deleteOrgOtherDevice

Delete OtherDevice

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **deviceMac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgOtherDevice**
> InlineResponse20054 GetOrgOtherDevice(ctx, orgId, deviceMac)
getOrgOtherDevice

Get Org other device (3rd party device)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **deviceMac** | **string**|  | 

### Return type

[**InlineResponse20054**](inline_response_200_54.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgOtherDevices**
> []DeviceOther ListOrgOtherDevices(ctx, orgId, optional)
listOrgOtherDevices

Get List of Org other devices (3rd party devices)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsDevicesOthersApiListOrgOtherDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsDevicesOthersApiListOrgOtherDevicesOpts struct
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

# **RebootOrgOtherDevice**
> RebootOrgOtherDevice(ctx, orgId, deviceMac)
rebootOrgOtherDevice

Reboot OtherDevice

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **deviceMac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrgOtherDeviceEvents**
> InlineResponse20053 SearchOrgOtherDeviceEvents(ctx, orgId, optional)
searchOrgOtherDeviceEvents

Search Org OtherDevices Events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsDevicesOthersApiSearchOrgOtherDeviceEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsDevicesOthersApiSearchOrgOtherDeviceEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | **optional.String**| site id | 
 **mac** | **optional.String**| mac | 
 **deviceMac** | **optional.String**| mac of attached device | 
 **model** | **optional.String**| device model | 
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

# **UpdateOrgOtherDevice**
> UpdateOrgOtherDevice(ctx, orgId, deviceMac, optional)
updateOrgOtherDevice

If the Site / Device cannot be identified, a manual association can be made

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **deviceMac** | **string**|  | 
 **optional** | ***OrgsDevicesOthersApiUpdateOrgOtherDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsDevicesOthersApiUpdateOrgOtherDeviceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OtherdevicesDeviceMacBody**](OtherdevicesDeviceMacBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgOtherDevices**
> UpdateOrgOtherDevices(ctx, orgId, optional)
updateOrgOtherDevices

Assign or unassign OtherDevices to and from a site.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsDevicesOthersApiUpdateOrgOtherDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsDevicesOthersApiUpdateOrgOtherDevicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdOtherdevicesBody**](OrgIdOtherdevicesBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

