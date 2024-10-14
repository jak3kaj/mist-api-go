# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdoptOrgJsiDevice**](OrgsJSIApi.md#AdoptOrgJsiDevice) | **Get** /api/v1/orgs/{org_id}/jsi/devices/outbound_ssh_cmd | adoptOrgJsiDevice
[**CreateOrgJsiDeviceShellSession**](OrgsJSIApi.md#CreateOrgJsiDeviceShellSession) | **Post** /api/v1/orgs/{org_id}/jsi/devices/{device_mac}/shell | createOrgJsiDeviceShellSession
[**ListOrgJsiDevices**](OrgsJSIApi.md#ListOrgJsiDevices) | **Get** /api/v1/orgs/{org_id}/jsi/devices | listOrgJsiDevices
[**ListOrgJsiPastPurchases**](OrgsJSIApi.md#ListOrgJsiPastPurchases) | **Get** /api/v1/orgs/{org_id}/jsi/inventory | listOrgJsiPastPurchases

# **AdoptOrgJsiDevice**
> InlineResponse20050 AdoptOrgJsiDevice(ctx, orgId)
adoptOrgJsiDevice

Adopt JSI devices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20050**](inline_response_200_50.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrgJsiDeviceShellSession**
> WebsocketSessionWithUrl CreateOrgJsiDeviceShellSession(ctx, orgId, deviceMac)
createOrgJsiDeviceShellSession

Create Shell Session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **deviceMac** | **string**|  | 

### Return type

[**WebsocketSessionWithUrl**](websocket_session_with_url.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgJsiDevices**
> []JseDevice ListOrgJsiDevices(ctx, orgId, optional)
listOrgJsiDevices

Get List of Org devices that connected to JSI

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsJSIApiListOrgJsiDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsJSIApiListOrgJsiDevicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]
 **model** | **optional.String**| Device model | 
 **serial** | **optional.String**| Device serial | 
 **mac** | **optional.String**| Device MAC Address | 

### Return type

[**[]JseDevice**](jse_device.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgJsiPastPurchases**
> []JseInventoryItem ListOrgJsiPastPurchases(ctx, orgId, optional)
listOrgJsiPastPurchases

Get List of all devices purchased from the accounts associated with the Org

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsJSIApiListOrgJsiPastPurchasesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsJSIApiListOrgJsiPastPurchasesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]
 **model** | **optional.String**|  | 
 **serial** | **optional.String**|  | 

### Return type

[**[]JseInventoryItem**](jse_inventory_item.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

