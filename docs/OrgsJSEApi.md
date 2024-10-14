# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgJsecCredential**](OrgsJSEApi.md#DeleteOrgJsecCredential) | **Delete** /api/v1/orgs/{org_id}/setting/jse/setup | deleteOrgJsecCredential
[**GetOrgJseInfo**](OrgsJSEApi.md#GetOrgJseInfo) | **Get** /api/v1/orgs/{org_id}/setting/jse/info | getOrgJseInfo
[**GetOrgJsecCredential**](OrgsJSEApi.md#GetOrgJsecCredential) | **Get** /api/v1/orgs/{org_id}/setting/jse/setup | getOrgJsecCredential
[**SetupOrgJsecCredential**](OrgsJSEApi.md#SetupOrgJsecCredential) | **Post** /api/v1/orgs/{org_id}/setting/jse/setup | setupOrgJsecCredential

# **DeleteOrgJsecCredential**
> DeleteOrgJsecCredential(ctx, orgId)
deleteOrgJsecCredential

Delete JSE credential

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

# **GetOrgJseInfo**
> InlineResponse20063 GetOrgJseInfo(ctx, orgId)
getOrgJseInfo

Retrieves the list of JSE orgs associated with the account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20063**](inline_response_200_63.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgJsecCredential**
> AccountJseInfo GetOrgJsecCredential(ctx, orgId)
getOrgJsecCredential

Get Org JSE Credential

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**AccountJseInfo**](account_jse_info.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetupOrgJsecCredential**
> AccountJseInfo SetupOrgJsecCredential(ctx, orgId, optional)
setupOrgJsecCredential

in JSE UI:  1. Create custom role with Read access to service_location and RW access to site and IPSec profile APIs.  2. Create a user with the above custom role. - email: john@abc.com  3. Activate the user in the JSE account.  4. Create the service locations on the JSE account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsJSEApiSetupOrgJsecCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsJSEApiSetupOrgJsecCredentialOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of JseSetupBody**](JseSetupBody.md)|  | 

### Return type

[**AccountJseInfo**](account_jse_info.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

