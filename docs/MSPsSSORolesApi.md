# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMspSsoRole**](MSPsSSORolesApi.md#CreateMspSsoRole) | **Post** /api/v1/msps/{msp_id}/ssoroles | createMspSsoRole
[**DeleteMspSsoRole**](MSPsSSORolesApi.md#DeleteMspSsoRole) | **Delete** /api/v1/msps/{msp_id}/ssoroles/{ssorole_id} | deleteMspSsoRole
[**ListMspSsoRoles**](MSPsSSORolesApi.md#ListMspSsoRoles) | **Get** /api/v1/msps/{msp_id}/ssoroles | listMspSsoRoles
[**UpdateMspSsoRole**](MSPsSSORolesApi.md#UpdateMspSsoRole) | **Put** /api/v1/msps/{msp_id}/ssoroles/{ssorole_id} | updateMspSsoRole

# **CreateMspSsoRole**
> InlineResponse20024 CreateMspSsoRole(ctx, mspId, optional)
createMspSsoRole

Create MSP Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
 **optional** | ***MSPsSSORolesApiCreateMspSsoRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsSSORolesApiCreateMspSsoRoleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MspIdSsorolesBody**](MspIdSsorolesBody.md)| Request Body | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMspSsoRole**
> DeleteMspSsoRole(ctx, mspId, ssoroleId)
deleteMspSsoRole

Delete MSP SSO Roles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **ssoroleId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMspSsoRoles**
> []SsoRoleOrg ListMspSsoRoles(ctx, mspId)
listMspSsoRoles

Get List of MSP SSO Roles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 

### Return type

[**[]SsoRoleOrg**](sso_role_org.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMspSsoRole**
> InlineResponse20024 UpdateMspSsoRole(ctx, mspId, ssoroleId, optional)
updateMspSsoRole

Update SSO Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **ssoroleId** | [**string**](.md)|  | 
 **optional** | ***MSPsSSORolesApiUpdateMspSsoRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsSSORolesApiUpdateMspSsoRoleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SsorolesSsoroleIdBody**](SsorolesSsoroleIdBody.md)| Request Body | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

