# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgMarvisClientInvites**](OrgsClientsMarvisApi.md#CreateOrgMarvisClientInvites) | **Post** /api/v1/orgs/{org_id}/marvisinvites | createOrgMarvisClientInvites
[**DeleteOrgMarvisClientInvite**](OrgsClientsMarvisApi.md#DeleteOrgMarvisClientInvite) | **Delete** /api/v1/orgs/{org_id}/marvisinvites/{marvisinvite_id} | deleteOrgMarvisClientInvite
[**GetOrgMarvisClientInvites**](OrgsClientsMarvisApi.md#GetOrgMarvisClientInvites) | **Get** /api/v1/orgs/{org_id}/marvisinvites/{marvisinvite_id} | getOrgMarvisClientInvites
[**ListOrgMarvisClientInvites**](OrgsClientsMarvisApi.md#ListOrgMarvisClientInvites) | **Get** /api/v1/orgs/{org_id}/marvisinvites | listOrgMarvisClientInvites
[**UpdateOrgMarvisClientInvite**](OrgsClientsMarvisApi.md#UpdateOrgMarvisClientInvite) | **Put** /api/v1/orgs/{org_id}/marvisinvites/{marvisinvite_id} | updateOrgMarvisClientInvite

# **CreateOrgMarvisClientInvites**
> InlineResponse20036 CreateOrgMarvisClientInvites(ctx, orgId, optional)
createOrgMarvisClientInvites

Create Org Marvis Client Invites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsClientsMarvisApiCreateOrgMarvisClientInvitesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsClientsMarvisApiCreateOrgMarvisClientInvitesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdMarvisinvitesBody**](OrgIdMarvisinvitesBody.md)|  | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgMarvisClientInvite**
> DeleteOrgMarvisClientInvite(ctx, orgId, marvisinviteId)
deleteOrgMarvisClientInvite

Delete Org Marvis Client Invite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **marvisinviteId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgMarvisClientInvites**
> InlineResponse20036 GetOrgMarvisClientInvites(ctx, orgId, marvisinviteId)
getOrgMarvisClientInvites

Get Org Marvis Client Invite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **marvisinviteId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgMarvisClientInvites**
> []MarvisClient ListOrgMarvisClientInvites(ctx, orgId)
listOrgMarvisClientInvites

List Org Marvis Client Invites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**[]MarvisClient**](marvis_client.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgMarvisClientInvite**
> InlineResponse20036 UpdateOrgMarvisClientInvite(ctx, orgId, marvisinviteId, optional)
updateOrgMarvisClientInvite

Update Org Marvis Client Invite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **marvisinviteId** | [**string**](.md)|  | 
 **optional** | ***OrgsClientsMarvisApiUpdateOrgMarvisClientInviteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsClientsMarvisApiUpdateOrgMarvisClientInviteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of MarvisinvitesMarvisinviteIdBody**](MarvisinvitesMarvisinviteIdBody.md)|  | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

