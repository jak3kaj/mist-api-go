# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgNacCrl**](OrgsNACCRLApi.md#DeleteOrgNacCrl) | **Delete** /api/v1/orgs/{org_id}/setting/mist_nac_crls/{naccrl_id} | deleteOrgNacCrl
[**GetOrgNacCrl**](OrgsNACCRLApi.md#GetOrgNacCrl) | **Get** /api/v1/orgs/{org_id}/setting/mist_nac_crls | getOrgNacCrl
[**ImportOrgNacCrl**](OrgsNACCRLApi.md#ImportOrgNacCrl) | **Post** /api/v1/orgs/{org_id}/setting/mist_nac_crls | importOrgNacCrl

# **DeleteOrgNacCrl**
> DeleteOrgNacCrl(ctx, orgId, naccrlId)
deleteOrgNacCrl

Delete NAC Org CRL file is a DELETE request to delete CRL file identified by its ID (ID assigned on file upload/creation)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **naccrlId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgNacCrl**
> ResponseNacCrlFiles GetOrgNacCrl(ctx, orgId)
getOrgNacCrl

Returns all uploaded CRL file IDs with names for the orgI

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**ResponseNacCrlFiles**](response_nac_crl_files.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportOrgNacCrl**
> NacCrlFile ImportOrgNacCrl(ctx, orgId, optional)
importOrgNacCrl

Import NAC Org CRL file is a multipart POST which has a .crl file to allow user to manually upload a Certificate Revocation List file. We support one file upload per Issuer. Re-uploads for the same issuer will overwrite the existing file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsNACCRLApiImportOrgNacCrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNACCRLApiImportOrgNacCrlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **optional.Interface of *os.File****optional.**|  | 
 **json** | **optional.**|  | 

### Return type

[**NacCrlFile**](nac_crl_file.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

