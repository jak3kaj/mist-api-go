# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InviteOrgAdmin**](OrgsAdminsApi.md#InviteOrgAdmin) | **Post** /api/v1/orgs/{org_id}/invites | inviteOrgAdmin
[**ListOrgAdmins**](OrgsAdminsApi.md#ListOrgAdmins) | **Get** /api/v1/orgs/{org_id}/admins | listOrgAdmins
[**RevokeOrgAdmin**](OrgsAdminsApi.md#RevokeOrgAdmin) | **Delete** /api/v1/orgs/{org_id}/admins/{admin_id} | revokeOrgAdmin
[**UninviteOrgAdmin**](OrgsAdminsApi.md#UninviteOrgAdmin) | **Delete** /api/v1/orgs/{org_id}/invites/{invite_id} | uninviteOrgAdmin
[**UpdateOrgAdmin**](OrgsAdminsApi.md#UpdateOrgAdmin) | **Put** /api/v1/orgs/{org_id}/admins/{admin_id} | updateOrgAdmin
[**UpdateOrgAdminInvite**](OrgsAdminsApi.md#UpdateOrgAdminInvite) | **Put** /api/v1/orgs/{org_id}/invites/{invite_id} | updateOrgAdminInvite

# **InviteOrgAdmin**
> InviteOrgAdmin(ctx, orgId, optional)
inviteOrgAdmin

If the request is successful, an email will also be sent to the user with a link to ```https://manage.mist.com/verify/invite?token=:token&expire=1459632743&org=OrgName```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsAdminsApiInviteOrgAdminOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAdminsApiInviteOrgAdminOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdInvitesBody**](OrgIdInvitesBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgAdmins**
> []Admin ListOrgAdmins(ctx, orgId)
listOrgAdmins

Get List of people who can manage the Site/Org under the Org

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**[]Admin**](admin.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeOrgAdmin**
> RevokeOrgAdmin(ctx, orgId, adminId)
revokeOrgAdmin

This removes all privileges this admin has against the org

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **adminId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UninviteOrgAdmin**
> UninviteOrgAdmin(ctx, orgId, inviteId)
uninviteOrgAdmin

Delete Admin Invite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **inviteId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgAdmin**
> InlineResponse20012 UpdateOrgAdmin(ctx, orgId, adminId, optional)
updateOrgAdmin

Invite Org Admin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **adminId** | [**string**](.md)|  | 
 **optional** | ***OrgsAdminsApiUpdateOrgAdminOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAdminsApiUpdateOrgAdminOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AdminsAdminIdBody1**](AdminsAdminIdBody1.md)| Request Body | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgAdminInvite**
> UpdateOrgAdminInvite(ctx, orgId, inviteId, optional)
updateOrgAdminInvite

Update Admin Invite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **inviteId** | [**string**](.md)|  | 
 **optional** | ***OrgsAdminsApiUpdateOrgAdminInviteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsAdminsApiUpdateOrgAdminInviteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of InvitesInviteIdBody1**](InvitesInviteIdBody1.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

