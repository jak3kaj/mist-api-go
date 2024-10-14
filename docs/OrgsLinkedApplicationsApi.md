# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrgOauthAppAccounts**](OrgsLinkedApplicationsApi.md#AddOrgOauthAppAccounts) | **Post** /api/v1/orgs/{org_id}/setting/{app_name}/link_accounts | addOrgOauthAppAccounts
[**DeleteOrgOauthAppAuthorization**](OrgsLinkedApplicationsApi.md#DeleteOrgOauthAppAuthorization) | **Delete** /api/v1/orgs/{org_id}/setting/{app_name}/link_accounts/{account_id} | deleteOrgOauthAppAuthorization
[**GetOrgOauthAppLinkedStatus**](OrgsLinkedApplicationsApi.md#GetOrgOauthAppLinkedStatus) | **Get** /api/v1/orgs/{org_id}/setting/{app_name}/link_accounts | getOrgOauthAppLinkedStatus
[**LinkOrgToJuniperJuniperAccount**](OrgsLinkedApplicationsApi.md#LinkOrgToJuniperJuniperAccount) | **Post** /api/v1/orgs/{org_id}/setting/juniper/link_accounts | linkOrgToJuniperJuniperAccount
[**UnlinkOrgFromJuniperCustomerId**](OrgsLinkedApplicationsApi.md#UnlinkOrgFromJuniperCustomerId) | **Delete** /api/v1/orgs/{org_id}/setting/juniper/unlink_account | unlinkOrgFromJuniperCustomerId
[**UpdateOrgOauthAppAccounts**](OrgsLinkedApplicationsApi.md#UpdateOrgOauthAppAccounts) | **Put** /api/v1/orgs/{org_id}/setting/{app_name}/link_accounts | updateOrgOauthAppAccounts

# **AddOrgOauthAppAccounts**
> AddOrgOauthAppAccounts(ctx, orgId, appName, optional)
addOrgOauthAppAccounts

Add Jamf, VMware Authorization With Mist Portal

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **appName** | [**AppName2**](.md)| OAuth application name | 
 **optional** | ***OrgsLinkedApplicationsApiAddOrgOauthAppAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsLinkedApplicationsApiAddOrgOauthAppAccountsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AppNameLinkAccountsBody1**](AppNameLinkAccountsBody1.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgOauthAppAuthorization**
> DeleteOrgOauthAppAuthorization(ctx, orgId, appName, accountId)
deleteOrgOauthAppAuthorization

Delete Org Level OAuth Application Authorization With Mist Portal

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **appName** | [**AppName3**](.md)| OAuth application name | 
  **accountId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgOauthAppLinkedStatus**
> ResponseOauthAppLink GetOrgOauthAppLinkedStatus(ctx, orgId, appName, forward)
getOrgOauthAppLinkedStatus

Get Org Level OAuth Application Linked Status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **appName** | [**AppName**](.md)| OAuth application name | 
  **forward** | **string**| Mist portal url to which backend needs to redirect after succesful OAuth authorization. **Required** to get the &#x60;authorization_url&#x60; | 

### Return type

[**ResponseOauthAppLink**](response_oauth_app_link.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinkOrgToJuniperJuniperAccount**
> InlineResponse20064 LinkOrgToJuniperJuniperAccount(ctx, orgId, optional)
linkOrgToJuniperJuniperAccount

Link Juniper Accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsLinkedApplicationsApiLinkOrgToJuniperJuniperAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsLinkedApplicationsApiLinkOrgToJuniperJuniperAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of JuniperLinkAccountsBody**](JuniperLinkAccountsBody.md)|  | 

### Return type

[**InlineResponse20064**](inline_response_200_64.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnlinkOrgFromJuniperCustomerId**
> UnlinkOrgFromJuniperCustomerId(ctx, orgId, contentType)
unlinkOrgFromJuniperCustomerId

Unlink Juniper Customer ID `linked_by` field is only required if there are duplicate account_names.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **contentType** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgOauthAppAccounts**
> UpdateOrgOauthAppAccounts(ctx, orgId, appName, optional)
updateOrgOauthAppAccounts

Update Zoom, Teams, Intune Authorization.  Request Payload, These Field And Values Will Be Specific To Each Of The Third Party Apps Accounts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **appName** | [**AppName1**](.md)| OAuth application name | 
 **optional** | ***OrgsLinkedApplicationsApiUpdateOrgOauthAppAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsLinkedApplicationsApiUpdateOrgOauthAppAccountsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of AppNameLinkAccountsBody**](AppNameLinkAccountsBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

