# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgNacRule**](OrgsNACRulesApi.md#CreateOrgNacRule) | **Post** /api/v1/orgs/{org_id}/nacrules | createOrgNacRule
[**DeleteOrgNacRule**](OrgsNACRulesApi.md#DeleteOrgNacRule) | **Delete** /api/v1/orgs/{org_id}/nacrules/{nacrule_id} | deleteOrgNacRule
[**GetOrgNacRule**](OrgsNACRulesApi.md#GetOrgNacRule) | **Get** /api/v1/orgs/{org_id}/nacrules/{nacrule_id} | getOrgNacRule
[**ListOrgNacRules**](OrgsNACRulesApi.md#ListOrgNacRules) | **Get** /api/v1/orgs/{org_id}/nacrules | listOrgNacRules
[**UpdateOrgNacRule**](OrgsNACRulesApi.md#UpdateOrgNacRule) | **Put** /api/v1/orgs/{org_id}/nacrules/{nacrule_id} | updateOrgNacRule

# **CreateOrgNacRule**
> NacRule CreateOrgNacRule(ctx, orgId, optional)
createOrgNacRule

create Org NAC Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsNACRulesApiCreateOrgNacRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNACRulesApiCreateOrgNacRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdNacrulesBody**](OrgIdNacrulesBody.md)|  | 

### Return type

[**NacRule**](nac_rule.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgNacRule**
> DeleteOrgNacRule(ctx, orgId, nacruleId)
deleteOrgNacRule

Delete Org NAC Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **nacruleId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgNacRule**
> NacRule GetOrgNacRule(ctx, orgId, nacruleId)
getOrgNacRule

Get Org NAC Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **nacruleId** | [**string**](.md)|  | 

### Return type

[**NacRule**](nac_rule.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgNacRules**
> []NacRule ListOrgNacRules(ctx, orgId, optional)
listOrgNacRules

Get List of Org NAC Rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsNACRulesApiListOrgNacRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNACRulesApiListOrgNacRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]NacRule**](nac_rule.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgNacRule**
> NacRule UpdateOrgNacRule(ctx, orgId, nacruleId, optional)
updateOrgNacRule

Update Org NAC Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **nacruleId** | [**string**](.md)|  | 
 **optional** | ***OrgsNACRulesApiUpdateOrgNacRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNACRulesApiUpdateOrgNacRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of NacRule**](NacRule.md)|  | 

### Return type

[**NacRule**](nac_rule.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

