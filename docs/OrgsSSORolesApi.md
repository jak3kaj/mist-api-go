# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgSsoRole**](OrgsSSORolesApi.md#CreateOrgSsoRole) | **Post** /api/v1/orgs/{org_id}/ssoroles | createOrgSsoRole
[**DeleteOrgSsoRole**](OrgsSSORolesApi.md#DeleteOrgSsoRole) | **Delete** /api/v1/orgs/{org_id}/ssoroles/{ssorole_id} | deleteOrgSsoRole
[**GetOrgSsoRole**](OrgsSSORolesApi.md#GetOrgSsoRole) | **Get** /api/v1/orgs/{org_id}/ssoroles/{ssorole_id} | getOrgSsoRole
[**ListOrgSsoRoles**](OrgsSSORolesApi.md#ListOrgSsoRoles) | **Get** /api/v1/orgs/{org_id}/ssoroles | listOrgSsoRoles
[**UpdateOrgSsoRole**](OrgsSSORolesApi.md#UpdateOrgSsoRole) | **Put** /api/v1/orgs/{org_id}/ssoroles/{ssorole_id} | updateOrgSsoRole

# **CreateOrgSsoRole**
> InlineResponse200102 CreateOrgSsoRole(ctx, orgId, optional)
createOrgSsoRole

Create Org SSO Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSSORolesApiCreateOrgSsoRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSSORolesApiCreateOrgSsoRoleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdSsorolesBody**](OrgIdSsorolesBody.md)| Request Body | 

### Return type

[**InlineResponse200102**](inline_response_200_102.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgSsoRole**
> DeleteOrgSsoRole(ctx, orgId, ssoroleId)
deleteOrgSsoRole

Delete Org SSO Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **ssoroleId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgSsoRole**
> InlineResponse200102 GetOrgSsoRole(ctx, orgId, ssoroleId)
getOrgSsoRole

Get Org SSO Role Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **ssoroleId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200102**](inline_response_200_102.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgSsoRoles**
> []SsoRoleMsp ListOrgSsoRoles(ctx, orgId, optional)
listOrgSsoRoles

Get List of Org SSO Roles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSSORolesApiListOrgSsoRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSSORolesApiListOrgSsoRolesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]SsoRoleMsp**](sso_role_msp.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgSsoRole**
> InlineResponse200102 UpdateOrgSsoRole(ctx, orgId, ssoroleId, optional)
updateOrgSsoRole

Update Org SSO Role

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **ssoroleId** | [**string**](.md)|  | 
 **optional** | ***OrgsSSORolesApiUpdateOrgSsoRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSSORolesApiUpdateOrgSsoRoleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SsorolesSsoroleIdBody1**](SsorolesSsoroleIdBody1.md)| Request Body | 

### Return type

[**InlineResponse200102**](inline_response_200_102.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

