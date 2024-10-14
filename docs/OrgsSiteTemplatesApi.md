# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgSiteTemplates**](OrgsSiteTemplatesApi.md#CreateOrgSiteTemplates) | **Post** /api/v1/orgs/{org_id}/sitetemplates | createOrgSiteTemplates
[**DeleteOrgSiteTemplate**](OrgsSiteTemplatesApi.md#DeleteOrgSiteTemplate) | **Delete** /api/v1/orgs/{org_id}/sitetemplates/{sitetemplate_id} | deleteOrgSiteTemplate
[**GetOrgSiteTemplate**](OrgsSiteTemplatesApi.md#GetOrgSiteTemplate) | **Get** /api/v1/orgs/{org_id}/sitetemplates/{sitetemplate_id} | getOrgSiteTemplate
[**ListOrgSiteTemplates**](OrgsSiteTemplatesApi.md#ListOrgSiteTemplates) | **Get** /api/v1/orgs/{org_id}/sitetemplates | listOrgSiteTemplates
[**UpdateOrgSiteTemplate**](OrgsSiteTemplatesApi.md#UpdateOrgSiteTemplate) | **Put** /api/v1/orgs/{org_id}/sitetemplates/{sitetemplate_id} | updateOrgSiteTemplate

# **CreateOrgSiteTemplates**
> InlineResponse200100 CreateOrgSiteTemplates(ctx, orgId, optional)
createOrgSiteTemplates

Create Org Site Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSiteTemplatesApiCreateOrgSiteTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSiteTemplatesApiCreateOrgSiteTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SiteTemplate**](SiteTemplate.md)|  | 

### Return type

[**InlineResponse200100**](inline_response_200_100.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgSiteTemplate**
> DeleteOrgSiteTemplate(ctx, orgId, sitetemplateId)
deleteOrgSiteTemplate

Delete Org Site Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **sitetemplateId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgSiteTemplate**
> InlineResponse200100 GetOrgSiteTemplate(ctx, orgId, sitetemplateId)
getOrgSiteTemplate

Get Org Site Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **sitetemplateId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200100**](inline_response_200_100.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgSiteTemplates**
> []SiteTemplate ListOrgSiteTemplates(ctx, orgId, optional)
listOrgSiteTemplates

Get List of Org Site Templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSiteTemplatesApiListOrgSiteTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSiteTemplatesApiListOrgSiteTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]SiteTemplate**](site_template.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgSiteTemplate**
> InlineResponse200100 UpdateOrgSiteTemplate(ctx, orgId, sitetemplateId, optional)
updateOrgSiteTemplate

Update Org Site Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **sitetemplateId** | [**string**](.md)|  | 
 **optional** | ***OrgsSiteTemplatesApiUpdateOrgSiteTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSiteTemplatesApiUpdateOrgSiteTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SiteTemplate**](SiteTemplate.md)|  | 

### Return type

[**InlineResponse200100**](inline_response_200_100.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

