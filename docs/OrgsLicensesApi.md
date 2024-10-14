# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClaimOrgLicense**](OrgsLicensesApi.md#ClaimOrgLicense) | **Post** /api/v1/orgs/{org_id}/claim | claimOrgLicense
[**GetOrgLicencesBySite**](OrgsLicensesApi.md#GetOrgLicencesBySite) | **Get** /api/v1/orgs/{org_id}/licenses/usages | getOrgLicencesBySite
[**GetOrgLicencesSummary**](OrgsLicensesApi.md#GetOrgLicencesSummary) | **Get** /api/v1/orgs/{org_id}/licenses | getOrgLicencesSummary
[**MoveOrDeleteOrgLicenseToAnotherOrg**](OrgsLicensesApi.md#MoveOrDeleteOrgLicenseToAnotherOrg) | **Put** /api/v1/orgs/{org_id}/licenses | moveOrDeleteOrgLicenseToAnotherOrg

# **ClaimOrgLicense**
> InlineResponse20017 ClaimOrgLicense(ctx, orgId, optional)
claimOrgLicense

Claim Org licenses / activation codes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsLicensesApiClaimOrgLicenseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsLicensesApiClaimOrgLicenseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdClaimBody**](OrgIdClaimBody.md)| Request Body | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgLicencesBySite**
> []LicenseUsageOrg GetOrgLicencesBySite(ctx, orgId)
getOrgLicencesBySite

Get Licenses Usage by Sites This shows license usage (i.e. needed) based on the features enabled for the site.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**[]LicenseUsageOrg**](license_usage_org.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgLicencesSummary**
> InlineResponse20018 GetOrgLicencesSummary(ctx, orgId)
getOrgLicencesSummary

Get the list of licenses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveOrDeleteOrgLicenseToAnotherOrg**
> MoveOrDeleteOrgLicenseToAnotherOrg(ctx, orgId, optional)
moveOrDeleteOrgLicenseToAnotherOrg

Move, Undo Move or Delete Org License to Another Org If the admin has admin privilege against the `org_id` and `dst_org_id`, he can move some of the licenses to another Org. Given that:  1. the specified license is currently active  2. there’s enough licenses left in the specified license (by subscription_id)  3. there will still be enough entitled licenses for the type of license after the amendment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsLicensesApiMoveOrDeleteOrgLicenseToAnotherOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsLicensesApiMoveOrDeleteOrgLicenseToAnotherOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdLicensesBody**](OrgIdLicensesBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

