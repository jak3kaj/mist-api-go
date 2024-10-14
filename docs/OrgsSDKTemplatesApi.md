# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSdkTemplate**](OrgsSDKTemplatesApi.md#CreateSdkTemplate) | **Post** /api/v1/orgs/{org_id}/sdktemplates | createSdkTemplate
[**DeleteSdkTemplate**](OrgsSDKTemplatesApi.md#DeleteSdkTemplate) | **Delete** /api/v1/orgs/{org_id}/sdktemplates/{sdktemplate_id} | deleteSdkTemplate
[**GetSdkTemplate**](OrgsSDKTemplatesApi.md#GetSdkTemplate) | **Get** /api/v1/orgs/{org_id}/sdktemplates/{sdktemplate_id} | getSdkTemplate
[**ListSdkTemplates**](OrgsSDKTemplatesApi.md#ListSdkTemplates) | **Get** /api/v1/orgs/{org_id}/sdktemplates | listSdkTemplates
[**UpdateSdkTemplate**](OrgsSDKTemplatesApi.md#UpdateSdkTemplate) | **Put** /api/v1/orgs/{org_id}/sdktemplates/{sdktemplate_id} | updateSdkTemplate

# **CreateSdkTemplate**
> InlineResponse20090 CreateSdkTemplate(ctx, orgId, optional)
createSdkTemplate

Create SDK Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSDKTemplatesApiCreateSdkTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSDKTemplatesApiCreateSdkTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdSdktemplatesBody**](OrgIdSdktemplatesBody.md)| Request Body | 

### Return type

[**InlineResponse20090**](inline_response_200_90.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSdkTemplate**
> DeleteSdkTemplate(ctx, orgId, sdktemplateId)
deleteSdkTemplate

Delete SDK Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **sdktemplateId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSdkTemplate**
> InlineResponse20090 GetSdkTemplate(ctx, orgId, sdktemplateId)
getSdkTemplate

Get SDK Template Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **sdktemplateId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20090**](inline_response_200_90.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSdkTemplates**
> []Sdktemplate ListSdkTemplates(ctx, orgId)
listSdkTemplates

Get List of Org SDK Templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**[]Sdktemplate**](sdktemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSdkTemplate**
> InlineResponse20090 UpdateSdkTemplate(ctx, orgId, sdktemplateId, optional)
updateSdkTemplate

Update SDK Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **sdktemplateId** | [**string**](.md)|  | 
 **optional** | ***OrgsSDKTemplatesApiUpdateSdkTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSDKTemplatesApiUpdateSdkTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SdktemplatesSdktemplateIdBody**](SdktemplatesSdktemplateIdBody.md)| Request Body | 

### Return type

[**InlineResponse20090**](inline_response_200_90.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

