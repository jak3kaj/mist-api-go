# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgApiToken**](OrgsAPITokensApi.md#CreateOrgApiToken) | **Post** /api/v1/orgs/{org_id}/apitokens | createOrgApiToken
[**DeleteOrgApiToken**](OrgsAPITokensApi.md#DeleteOrgApiToken) | **Delete** /api/v1/orgs/{org_id}/apitokens/{apitoken_id} | deleteOrgApiToken
[**GetOrgApiToken**](OrgsAPITokensApi.md#GetOrgApiToken) | **Get** /api/v1/orgs/{org_id}/apitokens/{apitoken_id} | getOrgApiToken
[**ListOrgApiTokens**](OrgsAPITokensApi.md#ListOrgApiTokens) | **Get** /api/v1/orgs/{org_id}/apitokens | listOrgApiTokens
[**UpdateOrgApiToken**](OrgsAPITokensApi.md#UpdateOrgApiToken) | **Put** /api/v1/orgs/{org_id}/apitokens/{apitoken_id} | updateOrgApiToken

# **CreateOrgApiToken**
> InlineResponse20031 CreateOrgApiToken(ctx, orgId, optional)
createOrgApiToken

Create Org API Token Note that the token key is only available during creation time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAPITokensApiCreateOrgApiTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAPITokensApiCreateOrgApiTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgApitoken**](OrgApitoken.md)|  | 

### Return type

[**InlineResponse20031**](inline_response_200_31.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgApiToken**
> DeleteOrgApiToken(ctx, orgId, apitokenId)
deleteOrgApiToken

Delete Org API Token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **apitokenId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgApiToken**
> InlineResponse20031 GetOrgApiToken(ctx, orgId, apitokenId)
getOrgApiToken

Get Org API Token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **apitokenId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20031**](inline_response_200_31.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgApiTokens**
> []OrgApitoken ListOrgApiTokens(ctx, orgId)
listOrgApiTokens

Get List of Org API Tokens

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**[]OrgApitoken**](org_apitoken.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgApiToken**
> UpdateOrgApiToken(ctx, orgId, apitokenId, optional)
updateOrgApiToken

Update Org API Token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **apitokenId** | [**string**](.md)|  | 
 **optional** | ***OrgsAPITokensApiUpdateOrgApiTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAPITokensApiUpdateOrgApiTokenOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApitokensApitokenIdBody**](ApitokensApitokenIdBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

