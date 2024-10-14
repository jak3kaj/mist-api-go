# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgRfTemplate**](OrgsRFTemplatesApi.md#CreateOrgRfTemplate) | **Post** /api/v1/orgs/{org_id}/rftemplates | createOrgRfTemplate
[**DeleteOrgRfTemplate**](OrgsRFTemplatesApi.md#DeleteOrgRfTemplate) | **Delete** /api/v1/orgs/{org_id}/rftemplates/{rftemplate_id} | deleteOrgRfTemplate
[**GetOrgRfTemplate**](OrgsRFTemplatesApi.md#GetOrgRfTemplate) | **Get** /api/v1/orgs/{org_id}/rftemplates/{rftemplate_id} | getOrgRfTemplate
[**ListOrgRfTemplates**](OrgsRFTemplatesApi.md#ListOrgRfTemplates) | **Get** /api/v1/orgs/{org_id}/rftemplates | listOrgRfTemplates
[**UpdateOrgRfTemplate**](OrgsRFTemplatesApi.md#UpdateOrgRfTemplate) | **Put** /api/v1/orgs/{org_id}/rftemplates/{rftemplate_id} | updateOrgRfTemplate

# **CreateOrgRfTemplate**
> InlineResponse20087 CreateOrgRfTemplate(ctx, orgId, optional)
createOrgRfTemplate

Create Org RF Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsRFTemplatesApiCreateOrgRfTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsRFTemplatesApiCreateOrgRfTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdRftemplatesBody**](OrgIdRftemplatesBody.md)| Request Body | 

### Return type

[**InlineResponse20087**](inline_response_200_87.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgRfTemplate**
> DeleteOrgRfTemplate(ctx, orgId, rftemplateId)
deleteOrgRfTemplate

Delete Org RF Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **rftemplateId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgRfTemplate**
> InlineResponse20087 GetOrgRfTemplate(ctx, orgId, rftemplateId)
getOrgRfTemplate

Get Org RF Template Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **rftemplateId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20087**](inline_response_200_87.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgRfTemplates**
> []RfTemplate ListOrgRfTemplates(ctx, orgId, optional)
listOrgRfTemplates

Get List of Org RF Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsRFTemplatesApiListOrgRfTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsRFTemplatesApiListOrgRfTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]RfTemplate**](rf_template.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgRfTemplate**
> InlineResponse20087 UpdateOrgRfTemplate(ctx, orgId, rftemplateId, optional)
updateOrgRfTemplate

Update Org RF Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **rftemplateId** | [**string**](.md)|  | 
 **optional** | ***OrgsRFTemplatesApiUpdateOrgRfTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsRFTemplatesApiUpdateOrgRfTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of RftemplatesRftemplateIdBody**](RftemplatesRftemplateIdBody.md)| Request Body | 

### Return type

[**InlineResponse20087**](inline_response_200_87.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

