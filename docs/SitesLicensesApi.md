# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteLicenseUsage**](SitesLicensesApi.md#GetSiteLicenseUsage) | **Get** /api/v1/sites/{site_id}/licenses/usages | getSiteLicenseUsage

# **GetSiteLicenseUsage**
> InlineResponse200128 GetSiteLicenseUsage(ctx, siteId)
getSiteLicenseUsage

This shows license usage (i.e. needed) based on the features enabled for the site.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200128**](inline_response_200_128.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

