# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgSecIntelProfile**](OrgsSecIntelProfilesApi.md#CreateOrgSecIntelProfile) | **Post** /api/v1/orgs/{org_id}/secintelprofiles | createOrgSecIntelProfile
[**DeleteOrgSecIntelProfile**](OrgsSecIntelProfilesApi.md#DeleteOrgSecIntelProfile) | **Delete** /api/v1/orgs/{org_id}/secintelprofiles/{secintelprofile_id} | deleteOrgSecIntelProfile
[**GetOrgSecIntelProfile**](OrgsSecIntelProfilesApi.md#GetOrgSecIntelProfile) | **Get** /api/v1/orgs/{org_id}/secintelprofiles/{secintelprofile_id} | getOrgSecIntelProfile
[**ListOrgSecIntelProfiles**](OrgsSecIntelProfilesApi.md#ListOrgSecIntelProfiles) | **Get** /api/v1/orgs/{org_id}/secintelprofiles | listOrgSecIntelProfiles
[**UpdateOrgSecIntelProfile**](OrgsSecIntelProfilesApi.md#UpdateOrgSecIntelProfile) | **Put** /api/v1/orgs/{org_id}/secintelprofiles/{secintelprofile_id} | updateOrgSecIntelProfile

# **CreateOrgSecIntelProfile**
> InlineResponse20013 CreateOrgSecIntelProfile(ctx, orgId, optional)
createOrgSecIntelProfile

Create Sec Intel Profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSecIntelProfilesApiCreateOrgSecIntelProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSecIntelProfilesApiCreateOrgSecIntelProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdSecintelprofilesBody**](OrgIdSecintelprofilesBody.md)| Request Body | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgSecIntelProfile**
> DeleteOrgSecIntelProfile(ctx, orgId, secintelprofileId)
deleteOrgSecIntelProfile

Delete Sec Intel Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **secintelprofileId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgSecIntelProfile**
> InlineResponse20013 GetOrgSecIntelProfile(ctx, orgId, secintelprofileId)
getOrgSecIntelProfile

Get Sec Intel Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **secintelprofileId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgSecIntelProfiles**
> []SecintelProfile ListOrgSecIntelProfiles(ctx, orgId)
listOrgSecIntelProfiles

Get List of Sec Intel Profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**[]SecintelProfile**](secintel_profile.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgSecIntelProfile**
> InlineResponse20013 UpdateOrgSecIntelProfile(ctx, orgId, secintelprofileId, optional)
updateOrgSecIntelProfile

Update Sec Intel Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **secintelprofileId** | [**string**](.md)|  | 
 **optional** | ***OrgsSecIntelProfilesApiUpdateOrgSecIntelProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSecIntelProfilesApiUpdateOrgSecIntelProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SecintelprofilesSecintelprofileIdBody**](SecintelprofilesSecintelprofileIdBody.md)| Request Body | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

