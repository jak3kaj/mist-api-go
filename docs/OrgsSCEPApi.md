# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableOrgMistScep**](OrgsSCEPApi.md#DisableOrgMistScep) | **Delete** /api/v1/orgs/{org_id}/setting/mist_scep | disableOrgMistScep
[**GetOrgIssuedClientCertificates**](OrgsSCEPApi.md#GetOrgIssuedClientCertificates) | **Get** /api/v1/orgs/{org_id}/setting/mist_scep/client_certs | getOrgIssuedClientCertificates
[**GetOrgMistScep**](OrgsSCEPApi.md#GetOrgMistScep) | **Get** /api/v1/orgs/{org_id}/setting/mist_scep | getOrgMistScep
[**RevokeOrgIssuedClientCertificates**](OrgsSCEPApi.md#RevokeOrgIssuedClientCertificates) | **Post** /api/v1/orgs/{org_id}/setting/mist_scep/client_certs/revoke | revokeOrgIssuedClientCertificates
[**UpdateOrgMistScep**](OrgsSCEPApi.md#UpdateOrgMistScep) | **Put** /api/v1/orgs/{org_id}/setting/mist_scep | updateOrgMistScep

# **DisableOrgMistScep**
> SettingMistScepBody DisableOrgMistScep(ctx, orgId)
disableOrgMistScep

Disable Mist SCEP Org setting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**SettingMistScepBody**](setting_mist_scep_body.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgIssuedClientCertificates**
> InlineResponse20057 GetOrgIssuedClientCertificates(ctx, orgId, optional)
getOrgIssuedClientCertificates

Get Issued Client Certificates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSCEPApiGetOrgIssuedClientCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSCEPApiGetOrgIssuedClientCertificatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ssoNameId** | **optional.String**| sso_name_id obtained from NAC Portal | 
 **serialNumber** | **optional.String**| serial number of the certificate | 
 **deviceId** | **optional.String**| Device ID | 

### Return type

[**InlineResponse20057**](inline_response_200_57.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgMistScep**
> InlineResponse20056 GetOrgMistScep(ctx, orgId)
getOrgMistScep

Get Mist SCEP Org setting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20056**](inline_response_200_56.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeOrgIssuedClientCertificates**
> RevokeOrgIssuedClientCertificates(ctx, orgId, optional)
revokeOrgIssuedClientCertificates

Revoke Issued Client Certificates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSCEPApiRevokeOrgIssuedClientCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSCEPApiRevokeOrgIssuedClientCertificatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ClientCertsRevokeBody**](ClientCertsRevokeBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgMistScep**
> SettingMistScepBody UpdateOrgMistScep(ctx, orgId, optional)
updateOrgMistScep

Update Mist SCEP Org setting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSCEPApiUpdateOrgMistScepOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSCEPApiUpdateOrgMistScepOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SettingMistScepBody**](SettingMistScepBody.md)|  | 

### Return type

[**SettingMistScepBody**](setting_mist_scep_body.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

