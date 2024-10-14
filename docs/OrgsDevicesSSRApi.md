# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrg128TRegistrationCommands**](OrgsDevicesSSRApi.md#GetOrg128TRegistrationCommands) | **Get** /api/v1/orgs/{org_id}/128routers/register_cmd | getOrg128TRegistrationCommands

# **GetOrg128TRegistrationCommands**
> InlineResponse20051 GetOrg128TRegistrationCommands(ctx, orgId)
getOrg128TRegistrationCommands

128T devices can be managed/adopted by Mist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20051**](inline_response_200_51.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

