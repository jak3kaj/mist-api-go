# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGatewayDefaultConfig**](ConstantsModelsApi.md#GetGatewayDefaultConfig) | **Get** /api/v1/const/default_gateway_config | getGatewayDefaultConfig
[**ListDeviceModels**](ConstantsModelsApi.md#ListDeviceModels) | **Get** /api/v1/const/device_models | listDeviceModels
[**ListMxEdgeModels**](ConstantsModelsApi.md#ListMxEdgeModels) | **Get** /api/v1/const/mxedge_models | listMxEdgeModels
[**ListSupportedOtherDeviceModels**](ConstantsModelsApi.md#ListSupportedOtherDeviceModels) | **Get** /api/v1/const/otherdevice_models | listSupportedOtherDeviceModels

# **GetGatewayDefaultConfig**
> interface{} GetGatewayDefaultConfig(ctx, model, optional)
getGatewayDefaultConfig

Generate Default Gateway Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **model** | **string**| model the default gateway config is intended (as the default LAN/WAN port can differ) | 
 **optional** | ***ConstantsModelsApiGetGatewayDefaultConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConstantsModelsApiGetGatewayDefaultConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ha** | **optional.String**| whether the config is intended for HA | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDeviceModels**
> []ConstDeviceModel ListDeviceModels(ctx, )
listDeviceModels

Get list of AP device models for the Mist Site

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ConstDeviceModel**](const_device_model.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMxEdgeModels**
> []ConstMxedgeModel ListMxEdgeModels(ctx, )
listMxEdgeModels

Get List of available Mx Edge models

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ConstMxedgeModel**](const_mxedge_model.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSupportedOtherDeviceModels**
> []ConstOtherDeviceModel ListSupportedOtherDeviceModels(ctx, )
listSupportedOtherDeviceModels

Supported OtherDevice Models

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ConstOtherDeviceModel**](const_other_device_model.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

