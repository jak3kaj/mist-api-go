# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMspAdmin**](MSPsAdminsApi.md#GetMspAdmin) | **Get** /api/v1/msps/{msp_id}/admins/{admin_id} | getMspAdmin
[**InviteMspAdmin**](MSPsAdminsApi.md#InviteMspAdmin) | **Post** /api/v1/msps/{msp_id}/invites | inviteMspAdmin
[**ListMspAdmins**](MSPsAdminsApi.md#ListMspAdmins) | **Get** /api/v1/msps/{msp_id}/admins | listMspAdmins
[**RevokeMspAdmin**](MSPsAdminsApi.md#RevokeMspAdmin) | **Delete** /api/v1/msps/{msp_id}/admins/{admin_id} | revokeMspAdmin
[**UninviteMspAdmin**](MSPsAdminsApi.md#UninviteMspAdmin) | **Delete** /api/v1/msps/{msp_id}/invites/{invite_id} | uninviteMspAdmin
[**UpdateMspAdmin**](MSPsAdminsApi.md#UpdateMspAdmin) | **Put** /api/v1/msps/{msp_id}/admins/{admin_id} | updateMspAdmin
[**UpdateMspAdminInvite**](MSPsAdminsApi.md#UpdateMspAdminInvite) | **Put** /api/v1/msps/{msp_id}/invites/{invite_id} | updateMspAdminInvite

# **GetMspAdmin**
> InlineResponse20012 GetMspAdmin(ctx, mspId, adminId)
getMspAdmin

Get MSP Admins

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **adminId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InviteMspAdmin**
> InlineResponse20012 InviteMspAdmin(ctx, mspId, optional)
inviteMspAdmin

Invite MSP Admin  **Note**: An email will also be sent to the user with a link to https://manage.mist.com/verify/invite?token=:token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
 **optional** | ***MSPsAdminsApiInviteMspAdminOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsAdminsApiInviteMspAdminOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MspIdInvitesBody**](MspIdInvitesBody.md)| Request Body | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMspAdmins**
> []Admin ListMspAdmins(ctx, mspId)
listMspAdmins

Get List of MSP Admins

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 

### Return type

[**[]Admin**](admin.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeMspAdmin**
> RevokeMspAdmin(ctx, mspId, adminId)
revokeMspAdmin

This removes all privileges this admin has against the MSP. This goes deep all the way to the sites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **adminId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UninviteMspAdmin**
> UninviteMspAdmin(ctx, mspId, inviteId)
uninviteMspAdmin

Delete admin invite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **inviteId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMspAdmin**
> InlineResponse20012 UpdateMspAdmin(ctx, mspId, adminId, optional)
updateMspAdmin

Update MSP Admin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **adminId** | [**string**](.md)|  | 
 **optional** | ***MSPsAdminsApiUpdateMspAdminOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsAdminsApiUpdateMspAdminOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AdminsAdminIdBody**](AdminsAdminIdBody.md)| Request Body | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMspAdminInvite**
> InlineResponse20012 UpdateMspAdminInvite(ctx, mspId, inviteId, optional)
updateMspAdminInvite

Update MSP admin invite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **inviteId** | [**string**](.md)|  | 
 **optional** | ***MSPsAdminsApiUpdateMspAdminInviteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsAdminsApiUpdateMspAdminInviteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of InvitesInviteIdBody**](InvitesInviteIdBody.md)| Request Body | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

