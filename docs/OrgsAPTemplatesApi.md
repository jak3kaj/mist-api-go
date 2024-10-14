# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgAptemplate**](OrgsAPTemplatesApi.md#CreateOrgAptemplate) | **Post** /api/v1/orgs/{org_id}/aptemplates | createOrgAptemplate
[**DeleteOrgAptemplate**](OrgsAPTemplatesApi.md#DeleteOrgAptemplate) | **Delete** /api/v1/orgs/{org_id}/aptemplates/{aptemplate_id} | deleteOrgAptemplate
[**GetOrgAptemplate**](OrgsAPTemplatesApi.md#GetOrgAptemplate) | **Get** /api/v1/orgs/{org_id}/aptemplates/{aptemplate_id} | getOrgAptemplate
[**ListOrgAptemplates**](OrgsAPTemplatesApi.md#ListOrgAptemplates) | **Get** /api/v1/orgs/{org_id}/aptemplates | listOrgAptemplates
[**UpdateOrgAptemplate**](OrgsAPTemplatesApi.md#UpdateOrgAptemplate) | **Put** /api/v1/orgs/{org_id}/aptemplates/{aptemplate_id} | updateOrgAptemplate

# **CreateOrgAptemplate**
> InlineResponse20030 CreateOrgAptemplate(ctx, orgId, optional)
createOrgAptemplate

Create Org AP Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAPTemplatesApiCreateOrgAptemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAPTemplatesApiCreateOrgAptemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdAptemplatesBody**](OrgIdAptemplatesBody.md)|  | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgAptemplate**
> DeleteOrgAptemplate(ctx, orgId, aptemplateId)
deleteOrgAptemplate

Delete existing AP Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **aptemplateId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgAptemplate**
> InlineResponse20030 GetOrgAptemplate(ctx, orgId, aptemplateId)
getOrgAptemplate

Get AP Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **aptemplateId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgAptemplates**
> []ApTemplate ListOrgAptemplates(ctx, orgId, optional)
listOrgAptemplates

Get List of Org AP Templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAPTemplatesApiListOrgAptemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAPTemplatesApiListOrgAptemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]ApTemplate**](ap_template.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgAptemplate**
> InlineResponse20030 UpdateOrgAptemplate(ctx, orgId, aptemplateId, optional)
updateOrgAptemplate

Update AP Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **aptemplateId** | [**string**](.md)|  | 
 **optional** | ***OrgsAPTemplatesApiUpdateOrgAptemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAPTemplatesApiUpdateOrgAptemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AptemplatesAptemplateIdBody**](AptemplatesAptemplateIdBody.md)|  | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

