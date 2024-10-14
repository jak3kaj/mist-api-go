# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneOrgTemplate**](OrgsWLANTemplatesApi.md#CloneOrgTemplate) | **Post** /api/v1/orgs/{org_id}/templates/{template_id}/clone | cloneOrgTemplate
[**CreateOrgTemplate**](OrgsWLANTemplatesApi.md#CreateOrgTemplate) | **Post** /api/v1/orgs/{org_id}/templates | createOrgTemplate
[**DeleteOrgTemplate**](OrgsWLANTemplatesApi.md#DeleteOrgTemplate) | **Delete** /api/v1/orgs/{org_id}/templates/{template_id} | deleteOrgTemplate
[**GetOrgTemplate**](OrgsWLANTemplatesApi.md#GetOrgTemplate) | **Get** /api/v1/orgs/{org_id}/templates/{template_id} | getOrgTemplate
[**ListOrgTemplates**](OrgsWLANTemplatesApi.md#ListOrgTemplates) | **Get** /api/v1/orgs/{org_id}/templates | listOrgTemplates
[**UpdateOrgTemplate**](OrgsWLANTemplatesApi.md#UpdateOrgTemplate) | **Put** /api/v1/orgs/{org_id}/templates/{template_id} | updateOrgTemplate

# **CloneOrgTemplate**
> InlineResponse200103 CloneOrgTemplate(ctx, orgId, templateId, optional)
cloneOrgTemplate

Clone Org Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **templateId** | [**string**](.md)|  | 
 **optional** | ***OrgsWLANTemplatesApiCloneOrgTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWLANTemplatesApiCloneOrgTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of TemplateIdCloneBody**](TemplateIdCloneBody.md)| Request Body | 

### Return type

[**InlineResponse200103**](inline_response_200_103.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrgTemplate**
> InlineResponse200103 CreateOrgTemplate(ctx, orgId, optional)
createOrgTemplate

Create Org Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsWLANTemplatesApiCreateOrgTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWLANTemplatesApiCreateOrgTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdTemplatesBody**](OrgIdTemplatesBody.md)| Request Body | 

### Return type

[**InlineResponse200103**](inline_response_200_103.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgTemplate**
> DeleteOrgTemplate(ctx, orgId, templateId)
deleteOrgTemplate

Delete Org Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **templateId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgTemplate**
> InlineResponse200103 GetOrgTemplate(ctx, orgId, templateId)
getOrgTemplate

Get Org Template Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **templateId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200103**](inline_response_200_103.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgTemplates**
> []Template ListOrgTemplates(ctx, orgId, optional)
listOrgTemplates

Get List of Org WLAN Templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsWLANTemplatesApiListOrgTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWLANTemplatesApiListOrgTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Template**](template.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgTemplate**
> InlineResponse200103 UpdateOrgTemplate(ctx, orgId, templateId, optional)
updateOrgTemplate

Update Org Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **templateId** | [**string**](.md)|  | 
 **optional** | ***OrgsWLANTemplatesApiUpdateOrgTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWLANTemplatesApiUpdateOrgTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of TemplatesTemplateIdBody**](TemplatesTemplateIdBody.md)| Request Body | 

### Return type

[**InlineResponse200103**](inline_response_200_103.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

