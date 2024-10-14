# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgAntivirusProfile**](OrgsAntivirusProfilesApi.md#CreateOrgAntivirusProfile) | **Post** /api/v1/orgs/{org_id}/avprofiles | createOrgAntivirusProfile
[**DeleteOrgAntivirusProfile**](OrgsAntivirusProfilesApi.md#DeleteOrgAntivirusProfile) | **Delete** /api/v1/orgs/{org_id}/avprofiles/{avprofiles_id} | deleteOrgAntivirusProfile
[**GetOrgAntivirusProfile**](OrgsAntivirusProfilesApi.md#GetOrgAntivirusProfile) | **Get** /api/v1/orgs/{org_id}/avprofiles/{avprofiles_id} | getOrgAntivirusProfile
[**ListOrgAntivirusProfiles**](OrgsAntivirusProfilesApi.md#ListOrgAntivirusProfiles) | **Get** /api/v1/orgs/{org_id}/avprofiles | listOrgAntivirusProfiles
[**UpdateOrgAntivirusProfile**](OrgsAntivirusProfilesApi.md#UpdateOrgAntivirusProfile) | **Put** /api/v1/orgs/{org_id}/avprofiles/{avprofiles_id} | updateOrgAntivirusProfile

# **CreateOrgAntivirusProfile**
> OrgIdAvprofilesBody CreateOrgAntivirusProfile(ctx, orgId, optional)
createOrgAntivirusProfile

Create getOrgServices Antivirus Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAntivirusProfilesApiCreateOrgAntivirusProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAntivirusProfilesApiCreateOrgAntivirusProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdAvprofilesBody**](OrgIdAvprofilesBody.md)|  | 

### Return type

[**OrgIdAvprofilesBody**](org_id_avprofiles_body.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgAntivirusProfile**
> DeleteOrgAntivirusProfile(ctx, orgId, avprofilesId)
deleteOrgAntivirusProfile

deleteOrgAntivirusProfile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **avprofilesId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgAntivirusProfile**
> OrgIdAvprofilesBody GetOrgAntivirusProfile(ctx, orgId, avprofilesId)
getOrgAntivirusProfile

Get Org Antivirus Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **avprofilesId** | [**string**](.md)|  | 

### Return type

[**OrgIdAvprofilesBody**](org_id_avprofiles_body.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgAntivirusProfiles**
> []Avprofile ListOrgAntivirusProfiles(ctx, orgId, optional)
listOrgAntivirusProfiles

Get List of Antivirus Profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAntivirusProfilesApiListOrgAntivirusProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAntivirusProfilesApiListOrgAntivirusProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Avprofile**](avprofile.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgAntivirusProfile**
> AvprofilesAvprofilesIdBody UpdateOrgAntivirusProfile(ctx, orgId, avprofilesId, optional)
updateOrgAntivirusProfile

update Org Antivirus Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **avprofilesId** | [**string**](.md)|  | 
 **optional** | ***OrgsAntivirusProfilesApiUpdateOrgAntivirusProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAntivirusProfilesApiUpdateOrgAntivirusProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AvprofilesAvprofilesIdBody**](AvprofilesAvprofilesIdBody.md)|  | 

### Return type

[**AvprofilesAvprofilesIdBody**](avprofiles_avprofiles_id_body.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

