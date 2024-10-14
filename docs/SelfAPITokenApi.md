# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiToken**](SelfAPITokenApi.md#CreateApiToken) | **Post** /api/v1/self/apitokens | createApiToken
[**DeleteApiToken**](SelfAPITokenApi.md#DeleteApiToken) | **Delete** /api/v1/self/apitokens/{apitoken_id} | deleteApiToken
[**GetApiToken**](SelfAPITokenApi.md#GetApiToken) | **Get** /api/v1/self/apitokens/{apitoken_id} | getApiToken
[**ListApiTokens**](SelfAPITokenApi.md#ListApiTokens) | **Get** /api/v1/self/apitokens | listApiTokens
[**UpdateApiToken**](SelfAPITokenApi.md#UpdateApiToken) | **Put** /api/v1/self/apitokens/{apitoken_id} | updateApiToken

# **CreateApiToken**
> []UserApitoken CreateApiToken(ctx, optional)
createApiToken

Create API Token Note that the key is only available during creation time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SelfAPITokenApiCreateApiTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SelfAPITokenApiCreateApiTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SelfApitokensBody**](SelfApitokensBody.md)|  | 

### Return type

[**[]UserApitoken**](user_apitoken.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApiToken**
> DeleteApiToken(ctx, apitokenId)
deleteApiToken

Delete an API Token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apitokenId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiToken**
> InlineResponse200181 GetApiToken(ctx, apitokenId)
getApiToken

Get User API Token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apitokenId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200181**](inline_response_200_181.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApiTokens**
> []UserApitoken ListApiTokens(ctx, )
listApiTokens

Get List of Current User API Tokens

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]UserApitoken**](user_apitoken.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApiToken**
> InlineResponse200181 UpdateApiToken(ctx, apitokenId, optional)
updateApiToken

Update User API Token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apitokenId** | [**string**](.md)|  | 
 **optional** | ***SelfAPITokenApiUpdateApiTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SelfAPITokenApiUpdateApiTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApitokensApitokenIdBody1**](ApitokensApitokenIdBody1.md)|  | 

### Return type

[**InlineResponse200181**](inline_response_200_181.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

