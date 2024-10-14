# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgIdpProfile**](OrgsIDPProfilesApi.md#CreateOrgIdpProfile) | **Post** /api/v1/orgs/{org_id}/idpprofiles | createOrgIdpProfile
[**DeleteOrgIdpProfile**](OrgsIDPProfilesApi.md#DeleteOrgIdpProfile) | **Delete** /api/v1/orgs/{org_id}/idpprofiles/{idpprofile_id} | deleteOrgIdpProfile
[**GetOrgIdpProfile**](OrgsIDPProfilesApi.md#GetOrgIdpProfile) | **Get** /api/v1/orgs/{org_id}/idpprofiles/{idpprofile_id} | getOrgIdpProfile
[**ListOrgIdpProfiles**](OrgsIDPProfilesApi.md#ListOrgIdpProfiles) | **Get** /api/v1/orgs/{org_id}/idpprofiles | listOrgIdpProfiles
[**UpdateOrgIdpProfile**](OrgsIDPProfilesApi.md#UpdateOrgIdpProfile) | **Put** /api/v1/orgs/{org_id}/idpprofiles/{idpprofile_id} | updateOrgIdpProfile

# **CreateOrgIdpProfile**
> InlineResponse20061 CreateOrgIdpProfile(ctx, orgId, optional)
createOrgIdpProfile

Create Org IDP Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsIDPProfilesApiCreateOrgIdpProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsIDPProfilesApiCreateOrgIdpProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdIdpprofilesBody**](OrgIdIdpprofilesBody.md)|  | 

### Return type

[**InlineResponse20061**](inline_response_200_61.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgIdpProfile**
> DeleteOrgIdpProfile(ctx, orgId, idpprofileId)
deleteOrgIdpProfile

Delete Org IDP Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **idpprofileId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgIdpProfile**
> InlineResponse20061 GetOrgIdpProfile(ctx, orgId, idpprofileId)
getOrgIdpProfile

Get Org IDP Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **idpprofileId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20061**](inline_response_200_61.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgIdpProfiles**
> []IdpProfile ListOrgIdpProfiles(ctx, orgId, optional)
listOrgIdpProfiles

get the list of Org IDP Profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsIDPProfilesApiListOrgIdpProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsIDPProfilesApiListOrgIdpProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]IdpProfile**](idp_profile.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgIdpProfile**
> InlineResponse20061 UpdateOrgIdpProfile(ctx, orgId, idpprofileId, optional)
updateOrgIdpProfile

Update Org IDP Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **idpprofileId** | [**string**](.md)|  | 
 **optional** | ***OrgsIDPProfilesApiUpdateOrgIdpProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsIDPProfilesApiUpdateOrgIdpProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of IdpprofilesIdpprofileIdBody**](IdpprofilesIdpprofileIdBody.md)|  | 

### Return type

[**InlineResponse20061**](inline_response_200_61.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

