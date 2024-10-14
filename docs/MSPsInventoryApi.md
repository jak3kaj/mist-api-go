# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMspInventoryByMac**](MSPsInventoryApi.md#GetMspInventoryByMac) | **Get** /api/v1/msps/{msp_id}/inventory/{device_mac} | getMspInventoryByMac

# **GetMspInventoryByMac**
> InlineResponse20014 GetMspInventoryByMac(ctx, mspId, deviceMac)
getMspInventoryByMac

Get Inventoy By device MAC address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **deviceMac** | **string**|  | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

