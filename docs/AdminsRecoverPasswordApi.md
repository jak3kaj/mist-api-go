# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecoverPassword**](AdminsRecoverPasswordApi.md#RecoverPassword) | **Post** /api/v1/recover | recoverPassword
[**VerifyRecoverPasssword**](AdminsRecoverPasswordApi.md#VerifyRecoverPasssword) | **Post** /api/v1/recover/verify/{token} | verifyRecoverPasssword

# **RecoverPassword**
> RecoverPassword(ctx, optional)
recoverPassword

Recover Password An email will also be sent to the user with a link to https://manage.mist.com/verify/recover?token=:token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AdminsRecoverPasswordApiRecoverPasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdminsRecoverPasswordApiRecoverPasswordOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of V1RecoverBody**](V1RecoverBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyRecoverPasssword**
> VerifyRecoverPasssword(ctx, token)
verifyRecoverPasssword

Verify Recover Password With correct verification, the user will be authenticated. UI can then prompt for new password

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

