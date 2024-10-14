# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMspSle**](MSPsSLEsApi.md#GetMspSle) | **Get** /api/v1/msps/{msp_id}/insights/{metric} | getMspSle

# **GetMspSle**
> InlineResponse20023 GetMspSle(ctx, mspId, metric, optional)
getMspSle

Get MSP SLEs (all/worst Orgs ...)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **metric** | **string**| see /api/v1/const/insight_metrics for available metrics | 
 **optional** | ***MSPsSLEsApiGetMspSleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsSLEsApiGetMspSleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sle** | **optional.String**| see /api/v1/const/insight_metrics for more details | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **interval** | **optional.String**| Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

