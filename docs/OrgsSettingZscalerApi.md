# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgZscalerCredential**](OrgsSettingZscalerApi.md#DeleteOrgZscalerCredential) | **Delete** /api/v1/orgs/{org_id}/setting/zscaler/setup | deleteOrgZscalerCredential
[**GetOrgZscalerCredential**](OrgsSettingZscalerApi.md#GetOrgZscalerCredential) | **Get** /api/v1/orgs/{org_id}/setting/zscaler/setup | getOrgZscalerCredential
[**SetupOrgZscalerCredential**](OrgsSettingZscalerApi.md#SetupOrgZscalerCredential) | **Post** /api/v1/orgs/{org_id}/setting/zscaler/setup | setupOrgZscalerCredential

# **DeleteOrgZscalerCredential**
> DeleteOrgZscalerCredential(ctx, orgId)
deleteOrgZscalerCredential

To delete Zscaler credential

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgZscalerCredential**
> InlineResponse20055 GetOrgZscalerCredential(ctx, orgId)
getOrgZscalerCredential

To get Zscaler credential

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20055**](inline_response_200_55.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetupOrgZscalerCredential**
> SetupOrgZscalerCredential(ctx, orgId, optional)
setupOrgZscalerCredential

To setup Zscaler credential

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSettingZscalerApiSetupOrgZscalerCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSettingZscalerApiSetupOrgZscalerCredentialOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ZscalerSetupBody**](ZscalerSetupBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

