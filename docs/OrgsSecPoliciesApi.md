# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgSecPolicies**](OrgsSecPoliciesApi.md#CreateOrgSecPolicies) | **Post** /api/v1/orgs/{org_id}/secpolicies | createOrgSecPolicies
[**DeleteOrgSecPolicy**](OrgsSecPoliciesApi.md#DeleteOrgSecPolicy) | **Delete** /api/v1/orgs/{org_id}/secpolicies/{secpolicy_id} | deleteOrgSecPolicy
[**GetOrgSecPolicy**](OrgsSecPoliciesApi.md#GetOrgSecPolicy) | **Get** /api/v1/orgs/{org_id}/secpolicies/{secpolicy_id} | getOrgSecPolicy
[**ListOrgSecPolicies**](OrgsSecPoliciesApi.md#ListOrgSecPolicies) | **Get** /api/v1/orgs/{org_id}/secpolicies | listOrgSecPolicies
[**UpdateOrgSecPolicies**](OrgsSecPoliciesApi.md#UpdateOrgSecPolicies) | **Put** /api/v1/orgs/{org_id}/secpolicies/{secpolicy_id} | updateOrgSecPolicies

# **CreateOrgSecPolicies**
> InlineResponse20091 CreateOrgSecPolicies(ctx, orgId, optional)
createOrgSecPolicies

Create Org Security Policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSecPoliciesApiCreateOrgSecPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSecPoliciesApiCreateOrgSecPoliciesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Secpolicy**](Secpolicy.md)|  | 

### Return type

[**InlineResponse20091**](inline_response_200_91.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgSecPolicy**
> DeleteOrgSecPolicy(ctx, orgId, secpolicyId)
deleteOrgSecPolicy

Delete Org Security Policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **secpolicyId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgSecPolicy**
> InlineResponse20091 GetOrgSecPolicy(ctx, orgId, secpolicyId)
getOrgSecPolicy

Get Org Security Policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **secpolicyId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20091**](inline_response_200_91.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgSecPolicies**
> []Secpolicy ListOrgSecPolicies(ctx, orgId, optional)
listOrgSecPolicies

Get List of Org Security Policies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSecPoliciesApiListOrgSecPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSecPoliciesApiListOrgSecPoliciesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Secpolicy**](secpolicy.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgSecPolicies**
> InlineResponse20091 UpdateOrgSecPolicies(ctx, orgId, secpolicyId, optional)
updateOrgSecPolicies

Update Org Security Policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **secpolicyId** | [**string**](.md)|  | 
 **optional** | ***OrgsSecPoliciesApiUpdateOrgSecPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSecPoliciesApiUpdateOrgSecPoliciesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SecpoliciesSecpolicyIdBody**](SecpoliciesSecpolicyIdBody.md)| Request Body | 

### Return type

[**InlineResponse20091**](inline_response_200_91.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

