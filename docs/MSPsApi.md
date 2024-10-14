# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMsp**](MSPsApi.md#CreateMsp) | **Post** /api/v1/msps | createMsp
[**DeleteMsp**](MSPsApi.md#DeleteMsp) | **Delete** /api/v1/msps/{msp_id} | deleteMsp
[**GetMspDetails**](MSPsApi.md#GetMspDetails) | **Get** /api/v1/msps/{msp_id} | getMspDetails
[**SearchMspOrgGroup**](MSPsApi.md#SearchMspOrgGroup) | **Get** /api/v1/msps/{msp_id}/search | searchMspOrgGroup
[**UpdateMsp**](MSPsApi.md#UpdateMsp) | **Put** /api/v1/msps/{msp_id} | updateMsp

# **CreateMsp**
> InlineResponse20010 CreateMsp(ctx, optional)
createMsp

Create MSP account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MSPsApiCreateMspOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsApiCreateMspOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of V1MspsBody**](V1MspsBody.md)| Request Body | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMsp**
> DeleteMsp(ctx, mspId)
deleteMsp

Deleting MSP removes the MSP and OrgGroup under the MSP as well as all privileges associated with them. It does not remove any Org or Admins

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMspDetails**
> InlineResponse20010 GetMspDetails(ctx, mspId)
getMspDetails

Get MSP Detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchMspOrgGroup**
> InlineResponse20011 SearchMspOrgGroup(ctx, mspId, type_, optional)
searchMspOrgGroup

Search in MSP Orgs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **type_** | [**ModelType**](.md)| orgs | 
 **optional** | ***MSPsApiSearchMspOrgGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsApiSearchMspOrgGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **q** | **optional.String**| search string | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsp**
> InlineResponse20010 UpdateMsp(ctx, mspId, optional)
updateMsp

Update MSP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
 **optional** | ***MSPsApiUpdateMspOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsApiUpdateMspOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MspsMspIdBody**](MspsMspIdBody.md)| Request Body | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

