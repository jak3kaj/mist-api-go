# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgNetworkTemplate**](OrgsNetworkTemplatesApi.md#CreateOrgNetworkTemplate) | **Post** /api/v1/orgs/{org_id}/networktemplates | createOrgNetworkTemplate
[**DeleteOrgNetworkTemplate**](OrgsNetworkTemplatesApi.md#DeleteOrgNetworkTemplate) | **Delete** /api/v1/orgs/{org_id}/networktemplates/{networktemplate_id} | deleteOrgNetworkTemplate
[**GetOrgNetworkTemplate**](OrgsNetworkTemplatesApi.md#GetOrgNetworkTemplate) | **Get** /api/v1/orgs/{org_id}/networktemplates/{networktemplate_id} | getOrgNetworkTemplate
[**ListOrgNetworkTemplates**](OrgsNetworkTemplatesApi.md#ListOrgNetworkTemplates) | **Get** /api/v1/orgs/{org_id}/networktemplates | listOrgNetworkTemplates
[**UpdateOrgNetworkTemplates**](OrgsNetworkTemplatesApi.md#UpdateOrgNetworkTemplates) | **Put** /api/v1/orgs/{org_id}/networktemplates/{networktemplate_id} | updateOrgNetworkTemplates

# **CreateOrgNetworkTemplate**
> InlineResponse20081 CreateOrgNetworkTemplate(ctx, orgId, optional)
createOrgNetworkTemplate

Update Org Network Templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsNetworkTemplatesApiCreateOrgNetworkTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNetworkTemplatesApiCreateOrgNetworkTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdNetworktemplatesBody**](OrgIdNetworktemplatesBody.md)| Request Body | 

### Return type

[**InlineResponse20081**](inline_response_200_81.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgNetworkTemplate**
> DeleteOrgNetworkTemplate(ctx, orgId, networktemplateId)
deleteOrgNetworkTemplate

Delete Org Network Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **networktemplateId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgNetworkTemplate**
> InlineResponse20082 GetOrgNetworkTemplate(ctx, orgId, networktemplateId)
getOrgNetworkTemplate

Get Org Network Templates Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **networktemplateId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20082**](inline_response_200_82.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgNetworkTemplates**
> []NetworkTemplate ListOrgNetworkTemplates(ctx, orgId, optional)
listOrgNetworkTemplates

Get List of Org Network Templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsNetworkTemplatesApiListOrgNetworkTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNetworkTemplatesApiListOrgNetworkTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]NetworkTemplate**](network_template.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgNetworkTemplates**
> InlineResponse20083 UpdateOrgNetworkTemplates(ctx, orgId, networktemplateId, optional)
updateOrgNetworkTemplates

Update Org Network Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **networktemplateId** | [**string**](.md)|  | 
 **optional** | ***OrgsNetworkTemplatesApiUpdateOrgNetworkTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNetworkTemplatesApiUpdateOrgNetworkTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of NetworktemplatesNetworktemplateIdBody**](NetworktemplatesNetworktemplateIdBody.md)| Request Body | 

### Return type

[**InlineResponse20083**](inline_response_200_83.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

