# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOauth2UrlForLinking**](SelfOAuth2Api.md#GetOauth2UrlForLinking) | **Get** /api/v1/self/oauth/{provider} | getOauth2UrlForLinking
[**LinkOauth2MistAccount**](SelfOAuth2Api.md#LinkOauth2MistAccount) | **Post** /api/v1/self/oauth/{provider} | linkOauth2MistAccount

# **GetOauth2UrlForLinking**
> InlineResponse200182 GetOauth2UrlForLinking(ctx, provider, optional)
getOauth2UrlForLinking

Obtain Authorization URL for Linking

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **provider** | **string**|  | 
 **optional** | ***SelfOAuth2ApiGetOauth2UrlForLinkingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SelfOAuth2ApiGetOauth2UrlForLinkingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forward** | **optional.String**|  | 

### Return type

[**InlineResponse200182**](inline_response_200_182.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinkOauth2MistAccount**
> InlineResponse200183 LinkOauth2MistAccount(ctx, provider, optional)
linkOauth2MistAccount

Link Mist account with an OAuth2 Provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **provider** | **string**|  | 
 **optional** | ***SelfOAuth2ApiLinkOauth2MistAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SelfOAuth2ApiLinkOauth2MistAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OauthProviderBody1**](OauthProviderBody1.md)| Request Body | 

### Return type

[**InlineResponse200183**](inline_response_200_183.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

