# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Lookup**](AdminsLookupApi.md#Lookup) | **Post** /api/v1/login/lookup | lookup

# **Lookup**
> InlineResponse2004 Lookup(ctx, optional)
lookup

Login Lookup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AdminsLookupApiLookupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdminsLookupApiLookupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of LoginLookupBody**](LoginLookupBody.md)| Request Body | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

