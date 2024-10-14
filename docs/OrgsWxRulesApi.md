# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgWxRule**](OrgsWxRulesApi.md#CreateOrgWxRule) | **Post** /api/v1/orgs/{org_id}/wxrules | createOrgWxRule
[**DeleteOrgWxRule**](OrgsWxRulesApi.md#DeleteOrgWxRule) | **Delete** /api/v1/orgs/{org_id}/wxrules/{wxrule_id} | deleteOrgWxRule
[**GetOrgWxRule**](OrgsWxRulesApi.md#GetOrgWxRule) | **Get** /api/v1/orgs/{org_id}/wxrules/{wxrule_id} | getOrgWxRule
[**ListOrgWxRules**](OrgsWxRulesApi.md#ListOrgWxRules) | **Get** /api/v1/orgs/{org_id}/wxrules | listOrgWxRules
[**ListOrgWxRulesDerived**](OrgsWxRulesApi.md#ListOrgWxRulesDerived) | **Get** /api/v1/orgs/{org_id}/wxrules/derived | listOrgWxRulesDerived
[**UpdateOrgWxRule**](OrgsWxRulesApi.md#UpdateOrgWxRule) | **Put** /api/v1/orgs/{org_id}/wxrules/{wxrule_id} | updateOrgWxRule

# **CreateOrgWxRule**
> InlineResponse200115 CreateOrgWxRule(ctx, orgId, optional)
createOrgWxRule

Create Org WxRule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsWxRulesApiCreateOrgWxRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWxRulesApiCreateOrgWxRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdWxrulesBody**](OrgIdWxrulesBody.md)| Request Body | 

### Return type

[**InlineResponse200115**](inline_response_200_115.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgWxRule**
> DeleteOrgWxRule(ctx, orgId, wxruleId)
deleteOrgWxRule

Delete Org WxRule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **wxruleId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgWxRule**
> InlineResponse200115 GetOrgWxRule(ctx, orgId, wxruleId)
getOrgWxRule

Get Org WxRule Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **wxruleId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200115**](inline_response_200_115.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgWxRules**
> []WxlanRule ListOrgWxRules(ctx, orgId, optional)
listOrgWxRules

Get List of Org WxRules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsWxRulesApiListOrgWxRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWxRulesApiListOrgWxRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]WxlanRule**](wxlan_rule.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgWxRulesDerived**
> []WxlanRule ListOrgWxRulesDerived(ctx, orgId)
listOrgWxRulesDerived

Get Derived Org WxRule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**[]WxlanRule**](wxlan_rule.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgWxRule**
> InlineResponse200115 UpdateOrgWxRule(ctx, orgId, wxruleId, optional)
updateOrgWxRule

Update Org WxRule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **wxruleId** | [**string**](.md)|  | 
 **optional** | ***OrgsWxRulesApiUpdateOrgWxRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWxRulesApiUpdateOrgWxRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of WxrulesWxruleIdBody**](WxrulesWxruleIdBody.md)| Request Body | 

### Return type

[**InlineResponse200115**](inline_response_200_115.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

