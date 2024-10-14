# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgServicePolicy**](OrgsServicePoliciesApi.md#CreateOrgServicePolicy) | **Post** /api/v1/orgs/{org_id}/servicepolicies | createOrgServicePolicy
[**DeleteOrgServicePolicy**](OrgsServicePoliciesApi.md#DeleteOrgServicePolicy) | **Delete** /api/v1/orgs/{org_id}/servicepolicies/{servicepolicy_id} | deleteOrgServicePolicy
[**GetOrgServicePolicy**](OrgsServicePoliciesApi.md#GetOrgServicePolicy) | **Get** /api/v1/orgs/{org_id}/servicepolicies/{servicepolicy_id} | getOrgServicePolicy
[**ListOrgServicePolicies**](OrgsServicePoliciesApi.md#ListOrgServicePolicies) | **Get** /api/v1/orgs/{org_id}/servicepolicies | listOrgServicePolicies
[**UpdateOrgServicePolicy**](OrgsServicePoliciesApi.md#UpdateOrgServicePolicy) | **Put** /api/v1/orgs/{org_id}/servicepolicies/{servicepolicy_id} | updateOrgServicePolicy

# **CreateOrgServicePolicy**
> OrgServicePolicy CreateOrgServicePolicy(ctx, orgId, optional)
createOrgServicePolicy

Create Org Serrvice Policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsServicePoliciesApiCreateOrgServicePolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsServicePoliciesApiCreateOrgServicePolicyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdServicepoliciesBody**](OrgIdServicepoliciesBody.md)|  | 

### Return type

[**OrgServicePolicy**](org_service_policy.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgServicePolicy**
> DeleteOrgServicePolicy(ctx, orgId, servicepolicyId)
deleteOrgServicePolicy

Delete Org Service Policuy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **servicepolicyId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgServicePolicy**
> OrgServicePolicy GetOrgServicePolicy(ctx, orgId, servicepolicyId)
getOrgServicePolicy

Get Org Service Policy Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **servicepolicyId** | [**string**](.md)|  | 

### Return type

[**OrgServicePolicy**](org_service_policy.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgServicePolicies**
> []OrgServicePolicy ListOrgServicePolicies(ctx, orgId, optional)
listOrgServicePolicies

Get List of Org Service Policies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsServicePoliciesApiListOrgServicePoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsServicePoliciesApiListOrgServicePoliciesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]OrgServicePolicy**](org_service_policy.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgServicePolicy**
> OrgServicePolicy UpdateOrgServicePolicy(ctx, orgId, servicepolicyId, optional)
updateOrgServicePolicy

Update Org Serrvice Policy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **servicepolicyId** | [**string**](.md)|  | 
 **optional** | ***OrgsServicePoliciesApiUpdateOrgServicePolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsServicePoliciesApiUpdateOrgServicePolicyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ServicepoliciesServicepolicyIdBody**](ServicepoliciesServicepolicyIdBody.md)|  | 

### Return type

[**OrgServicePolicy**](org_service_policy.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

