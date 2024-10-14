# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignOrgDeviceProfile**](OrgsDeviceProfilesApi.md#AssignOrgDeviceProfile) | **Post** /api/v1/orgs/{org_id}/deviceprofiles/{deviceprofile_id}/assign | assignOrgDeviceProfile
[**CreateOrgDeviceProfiles**](OrgsDeviceProfilesApi.md#CreateOrgDeviceProfiles) | **Post** /api/v1/orgs/{org_id}/deviceprofiles | createOrgDeviceProfiles
[**DeleteOrgDeviceProfile**](OrgsDeviceProfilesApi.md#DeleteOrgDeviceProfile) | **Delete** /api/v1/orgs/{org_id}/deviceprofiles/{deviceprofile_id} | deleteOrgDeviceProfile
[**GetOrgDeviceProfile**](OrgsDeviceProfilesApi.md#GetOrgDeviceProfile) | **Get** /api/v1/orgs/{org_id}/deviceprofiles/{deviceprofile_id} | getOrgDeviceProfile
[**ListOrgDeviceProfiles**](OrgsDeviceProfilesApi.md#ListOrgDeviceProfiles) | **Get** /api/v1/orgs/{org_id}/deviceprofiles | listOrgDeviceProfiles
[**UnassignOrgDeviceProfile**](OrgsDeviceProfilesApi.md#UnassignOrgDeviceProfile) | **Post** /api/v1/orgs/{org_id}/deviceprofiles/{deviceprofile_id}/unassign | unassignOrgDeviceProfile
[**UpdateOrgDeviceProfile**](OrgsDeviceProfilesApi.md#UpdateOrgDeviceProfile) | **Put** /api/v1/orgs/{org_id}/deviceprofiles/{deviceprofile_id} | updateOrgDeviceProfile

# **AssignOrgDeviceProfile**
> InlineResponse20046 AssignOrgDeviceProfile(ctx, orgId, deviceprofileId, optional)
assignOrgDeviceProfile

Assign Org Device Profile to Devices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **deviceprofileId** | [**string**](.md)|  | 
 **optional** | ***OrgsDeviceProfilesApiAssignOrgDeviceProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsDeviceProfilesApiAssignOrgDeviceProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceprofileIdAssignBody**](DeviceprofileIdAssignBody.md)| Request Body | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrgDeviceProfiles**
> InlineResponse20045 CreateOrgDeviceProfiles(ctx, orgId, optional)
createOrgDeviceProfiles

Create Device Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsDeviceProfilesApiCreateOrgDeviceProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsDeviceProfilesApiCreateOrgDeviceProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdDeviceprofilesBody**](OrgIdDeviceprofilesBody.md)| Request Body | 

### Return type

[**InlineResponse20045**](inline_response_200_45.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgDeviceProfile**
> DeleteOrgDeviceProfile(ctx, orgId, deviceprofileId)
deleteOrgDeviceProfile

Delete Org Device Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **deviceprofileId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgDeviceProfile**
> InlineResponse20045 GetOrgDeviceProfile(ctx, orgId, deviceprofileId)
getOrgDeviceProfile

Get Org device Profile Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **deviceprofileId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20045**](inline_response_200_45.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgDeviceProfiles**
> []Deviceprofile ListOrgDeviceProfiles(ctx, orgId, optional)
listOrgDeviceProfiles

Get List of Org Device Profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsDeviceProfilesApiListOrgDeviceProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsDeviceProfilesApiListOrgDeviceProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**optional.Interface of DeviceType**](.md)|  | 
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Deviceprofile**](deviceprofile.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnassignOrgDeviceProfile**
> InlineResponse20046 UnassignOrgDeviceProfile(ctx, orgId, deviceprofileId, optional)
unassignOrgDeviceProfile

Unassign Org Device Profile from Devices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **deviceprofileId** | [**string**](.md)|  | 
 **optional** | ***OrgsDeviceProfilesApiUnassignOrgDeviceProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsDeviceProfilesApiUnassignOrgDeviceProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceprofileIdUnassignBody**](DeviceprofileIdUnassignBody.md)| Request Body | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgDeviceProfile**
> InlineResponse20045 UpdateOrgDeviceProfile(ctx, orgId, deviceprofileId, optional)
updateOrgDeviceProfile

Update Org Device Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **deviceprofileId** | [**string**](.md)|  | 
 **optional** | ***OrgsDeviceProfilesApiUpdateOrgDeviceProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsDeviceProfilesApiUpdateOrgDeviceProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceprofilesDeviceprofileIdBody**](DeviceprofilesDeviceprofileIdBody.md)| Request Body | 

### Return type

[**InlineResponse20045**](inline_response_200_45.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

