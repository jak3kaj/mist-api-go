# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSdkClient**](OrgsClientsSDKApi.md#UpdateSdkClient) | **Put** /api/v1/orgs/{org_id}/sdkclients/{sdkclient_id} | updateSdkClient

# **UpdateSdkClient**
> UpdateSdkClient(ctx, orgId, sdkclientId, optional)
updateSdkClient

Update SDK Client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **sdkclientId** | [**string**](.md)|  | 
 **optional** | ***OrgsClientsSDKApiUpdateSdkClientOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsClientsSDKApiUpdateSdkClientOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SdkclientsSdkclientIdBody**](SdkclientsSdkclientIdBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

