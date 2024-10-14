# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteJseInfo**](SitesJSEApi.md#GetSiteJseInfo) | **Get** /api/v1/sites/{site_id}/setting/jse/info | getSiteJseInfo

# **GetSiteJseInfo**
> InlineResponse20063 GetSiteJseInfo(ctx, siteId)
getSiteJseInfo

Retrieves the list of JSE orgs associated with the account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20063**](inline_response_200_63.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

