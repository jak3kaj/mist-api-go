# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateSdkInvite**](OrgsSDKInvitesApi.md#ActivateSdkInvite) | **Post** /api/v1/mobile/verify/{secret} | activateSdkInvite
[**CreateSdkInvite**](OrgsSDKInvitesApi.md#CreateSdkInvite) | **Post** /api/v1/orgs/{org_id}/sdkinvites | createSdkInvite
[**GetSdkInvite**](OrgsSDKInvitesApi.md#GetSdkInvite) | **Get** /api/v1/orgs/{org_id}/sdkinvites/{sdkinvite_id} | getSdkInvite
[**GetSdkInviteQrCode**](OrgsSDKInvitesApi.md#GetSdkInviteQrCode) | **Get** /api/v1/orgs/{org_id}/sdkinvites/{sdkinvite_id}/qrcode | getSdkInviteQrCode
[**ListSdkInvites**](OrgsSDKInvitesApi.md#ListSdkInvites) | **Get** /api/v1/orgs/{org_id}/sdkinvites | listSdkInvites
[**RevokeSdkInvite**](OrgsSDKInvitesApi.md#RevokeSdkInvite) | **Delete** /api/v1/orgs/{org_id}/sdkinvites/{sdkinvite_id} | revokeSdkInvite
[**SendSdkInviteEmail**](OrgsSDKInvitesApi.md#SendSdkInviteEmail) | **Post** /api/v1/orgs/{org_id}/sdkinvites/{sdkinvite_id}/email | sendSdkInviteEmail
[**SendSdkInviteSms**](OrgsSDKInvitesApi.md#SendSdkInviteSms) | **Post** /api/v1/orgs/{org_id}/sdkinvites/{sdkinvite_id}/sms | sendSdkInviteSms
[**UpdateSdkInvite**](OrgsSDKInvitesApi.md#UpdateSdkInvite) | **Put** /api/v1/orgs/{org_id}/sdkinvites/{sdkinvite_id} | updateSdkInvite

# **ActivateSdkInvite**
> InlineResponse20088 ActivateSdkInvite(ctx, secret, optional)
activateSdkInvite

Verify secret

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **secret** | **string**|  | 
 **optional** | ***OrgsSDKInvitesApiActivateSdkInviteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSDKInvitesApiActivateSdkInviteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of VerifySecretBody**](VerifySecretBody.md)|  | 

### Return type

[**InlineResponse20088**](inline_response_200_88.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSdkInvite**
> InlineResponse20089 CreateSdkInvite(ctx, orgId, optional)
createSdkInvite

Create SDK Invite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSDKInvitesApiCreateSdkInviteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSDKInvitesApiCreateSdkInviteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdSdkinvitesBody**](OrgIdSdkinvitesBody.md)| Request Body | 

### Return type

[**InlineResponse20089**](inline_response_200_89.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSdkInvite**
> InlineResponse20089 GetSdkInvite(ctx, orgId, sdkinviteId)
getSdkInvite

Get SDK Invite Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **sdkinviteId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20089**](inline_response_200_89.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSdkInviteQrCode**
> *os.File GetSdkInviteQrCode(ctx, orgId, sdkinviteId)
getSdkInviteQrCode

Revoke SDK Invite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **sdkinviteId** | [**string**](.md)|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSdkInvites**
> []Sdkinvite ListSdkInvites(ctx, orgId)
listSdkInvites

Get List of Org SDK Invites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**[]Sdkinvite**](sdkinvite.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeSdkInvite**
> RevokeSdkInvite(ctx, orgId, sdkinviteId)
revokeSdkInvite

Revoke SDK Invite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **sdkinviteId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendSdkInviteEmail**
> SendSdkInviteEmail(ctx, orgId, sdkinviteId, optional)
sendSdkInviteEmail

Send SDK Invite by Email

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **sdkinviteId** | [**string**](.md)|  | 
 **optional** | ***OrgsSDKInvitesApiSendSdkInviteEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSDKInvitesApiSendSdkInviteEmailOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SdkinviteIdEmailBody**](SdkinviteIdEmailBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendSdkInviteSms**
> SendSdkInviteSms(ctx, orgId, sdkinviteId, optional)
sendSdkInviteSms

Send SDK Invite by SMS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **sdkinviteId** | [**string**](.md)|  | 
 **optional** | ***OrgsSDKInvitesApiSendSdkInviteSmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSDKInvitesApiSendSdkInviteSmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SdkinviteIdSmsBody**](SdkinviteIdSmsBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSdkInvite**
> InlineResponse20089 UpdateSdkInvite(ctx, orgId, sdkinviteId, optional)
updateSdkInvite

Update SDK Invite

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **sdkinviteId** | [**string**](.md)|  | 
 **optional** | ***OrgsSDKInvitesApiUpdateSdkInviteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSDKInvitesApiUpdateSdkInviteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SdkinvitesSdkinviteIdBody**](SdkinvitesSdkinviteIdBody.md)| Request Body | 

### Return type

[**InlineResponse20089**](inline_response_200_89.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

