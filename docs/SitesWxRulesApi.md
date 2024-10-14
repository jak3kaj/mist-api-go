# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteWxRule**](SitesWxRulesApi.md#CreateSiteWxRule) | **Post** /api/v1/sites/{site_id}/wxrules | createSiteWxRule
[**DeleteSiteWxRule**](SitesWxRulesApi.md#DeleteSiteWxRule) | **Delete** /api/v1/sites/{site_id}/wxrules/{wxrule_id} | deleteSiteWxRule
[**GetSiteWxRule**](SitesWxRulesApi.md#GetSiteWxRule) | **Get** /api/v1/sites/{site_id}/wxrules/{wxrule_id} | getSiteWxRule
[**GetSiteWxRulesDerived**](SitesWxRulesApi.md#GetSiteWxRulesDerived) | **Get** /api/v1/sites/{site_id}/wxrules/derived | getSiteWxRulesDerived
[**ListSiteWxRules**](SitesWxRulesApi.md#ListSiteWxRules) | **Get** /api/v1/sites/{site_id}/wxrules | listSiteWxRules
[**UpdateSiteWxRule**](SitesWxRulesApi.md#UpdateSiteWxRule) | **Put** /api/v1/sites/{site_id}/wxrules/{wxrule_id} | updateSiteWxRule

# **CreateSiteWxRule**
> InlineResponse200115 CreateSiteWxRule(ctx, siteId, optional)
createSiteWxRule

Create Site WxLan Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesWxRulesApiCreateSiteWxRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWxRulesApiCreateSiteWxRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SiteIdWxrulesBody**](SiteIdWxrulesBody.md)| Request Body | 

### Return type

[**InlineResponse200115**](inline_response_200_115.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteWxRule**
> DeleteSiteWxRule(ctx, siteId, wxruleId)
deleteSiteWxRule

Delete Site WxLan Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **wxruleId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteWxRule**
> InlineResponse200115 GetSiteWxRule(ctx, siteId, wxruleId)
getSiteWxRule

Get Site WxLan Rule Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **wxruleId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200115**](inline_response_200_115.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteWxRulesDerived**
> []WxlanRule GetSiteWxRulesDerived(ctx, siteId)
getSiteWxRulesDerived

Get Site WxLan Rule Derived

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

[**[]WxlanRule**](wxlan_rule.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteWxRules**
> []WxlanRule ListSiteWxRules(ctx, siteId, optional)
listSiteWxRules

Get List of Site WxLan Rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesWxRulesApiListSiteWxRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWxRulesApiListSiteWxRulesOpts struct
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

# **UpdateSiteWxRule**
> InlineResponse200115 UpdateSiteWxRule(ctx, siteId, wxruleId, optional)
updateSiteWxRule

Update Site WxLan Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **wxruleId** | [**string**](.md)|  | 
 **optional** | ***SitesWxRulesApiUpdateSiteWxRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWxRulesApiUpdateSiteWxRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of WxrulesWxruleIdBody1**](WxrulesWxruleIdBody1.md)| Request Body | 

### Return type

[**InlineResponse200115**](inline_response_200_115.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

