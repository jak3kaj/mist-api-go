# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOauth2AuthorizationUrlForLogin**](AdminsLoginOAuth2Api.md#GetOauth2AuthorizationUrlForLogin) | **Get** /api/v1/login/oauth/{provider} | getOauth2AuthorizationUrlForLogin
[**LoginOauth2**](AdminsLoginOAuth2Api.md#LoginOauth2) | **Post** /api/v1/login/oauth/{provider} | loginOauth2
[**UnlinkOauth2Provider**](AdminsLoginOAuth2Api.md#UnlinkOauth2Provider) | **Delete** /api/v1/login/oauth/{provider} | unlinkOauth2Provider

# **GetOauth2AuthorizationUrlForLogin**
> InlineResponse2005 GetOauth2AuthorizationUrlForLogin(ctx, provider, optional)
getOauth2AuthorizationUrlForLogin

Obtain Authorization URL for Login

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **provider** | **string**|  | 
 **optional** | ***AdminsLoginOAuth2ApiGetOauth2AuthorizationUrlForLoginOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdminsLoginOAuth2ApiGetOauth2AuthorizationUrlForLoginOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forward** | **optional.String**| callback URL | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoginOauth2**
> LoginOauth2(ctx, provider, optional)
loginOauth2

Login via OAuth2

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **provider** | **string**|  | 
 **optional** | ***AdminsLoginOAuth2ApiLoginOauth2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdminsLoginOAuth2ApiLoginOauth2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OauthProviderBody**](OauthProviderBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnlinkOauth2Provider**
> UnlinkOauth2Provider(ctx, provider)
unlinkOauth2Provider

Unlink OAuth2 Provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **provider** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

