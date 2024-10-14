# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgGatewayTemplate**](OrgsGatewayTemplatesApi.md#CreateOrgGatewayTemplate) | **Post** /api/v1/orgs/{org_id}/gatewaytemplates | createOrgGatewayTemplate
[**DeleteOrgGatewayTemplate**](OrgsGatewayTemplatesApi.md#DeleteOrgGatewayTemplate) | **Delete** /api/v1/orgs/{org_id}/gatewaytemplates/{gatewaytemplate_id} | deleteOrgGatewayTemplate
[**GetOrgGatewayTemplate**](OrgsGatewayTemplatesApi.md#GetOrgGatewayTemplate) | **Get** /api/v1/orgs/{org_id}/gatewaytemplates/{gatewaytemplate_id} | getOrgGatewayTemplate
[**ListOrgGatewayTemplates**](OrgsGatewayTemplatesApi.md#ListOrgGatewayTemplates) | **Get** /api/v1/orgs/{org_id}/gatewaytemplates | listOrgGatewayTemplates
[**UpdateOrgGatewayTemplate**](OrgsGatewayTemplatesApi.md#UpdateOrgGatewayTemplate) | **Put** /api/v1/orgs/{org_id}/gatewaytemplates/{gatewaytemplate_id} | updateOrgGatewayTemplate

# **CreateOrgGatewayTemplate**
> InlineResponse20058 CreateOrgGatewayTemplate(ctx, orgId, optional)
createOrgGatewayTemplate

Create Org Gateway Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsGatewayTemplatesApiCreateOrgGatewayTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsGatewayTemplatesApiCreateOrgGatewayTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdGatewaytemplatesBody**](OrgIdGatewaytemplatesBody.md)| Gateway Template | 

### Return type

[**InlineResponse20058**](inline_response_200_58.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgGatewayTemplate**
> DeleteOrgGatewayTemplate(ctx, orgId, gatewaytemplateId)
deleteOrgGatewayTemplate

Delete Organization Gateway Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **gatewaytemplateId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgGatewayTemplate**
> InlineResponse20058 GetOrgGatewayTemplate(ctx, orgId, gatewaytemplateId)
getOrgGatewayTemplate

Get Organization Gateway Template details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **gatewaytemplateId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20058**](inline_response_200_58.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgGatewayTemplates**
> []GatewayTemplate ListOrgGatewayTemplates(ctx, orgId, optional)
listOrgGatewayTemplates

Get List of Org Gateway Templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsGatewayTemplatesApiListOrgGatewayTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsGatewayTemplatesApiListOrgGatewayTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]GatewayTemplate**](gateway_template.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgGatewayTemplate**
> InlineResponse20058 UpdateOrgGatewayTemplate(ctx, orgId, gatewaytemplateId, optional)
updateOrgGatewayTemplate

Update Organization Gateway Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **gatewaytemplateId** | [**string**](.md)|  | 
 **optional** | ***OrgsGatewayTemplatesApiUpdateOrgGatewayTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsGatewayTemplatesApiUpdateOrgGatewayTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of GatewaytemplatesGatewaytemplateIdBody**](GatewaytemplatesGatewaytemplateIdBody.md)| Gateway Template | 

### Return type

[**InlineResponse20058**](inline_response_200_58.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

