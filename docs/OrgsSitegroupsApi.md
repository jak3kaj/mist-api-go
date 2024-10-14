# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgSiteGroup**](OrgsSitegroupsApi.md#CreateOrgSiteGroup) | **Post** /api/v1/orgs/{org_id}/sitegroups | createOrgSiteGroup
[**DeleteOrgSiteGroup**](OrgsSitegroupsApi.md#DeleteOrgSiteGroup) | **Delete** /api/v1/orgs/{org_id}/sitegroups/{sitegroup_id} | deleteOrgSiteGroup
[**GetOrgSiteGroup**](OrgsSitegroupsApi.md#GetOrgSiteGroup) | **Get** /api/v1/orgs/{org_id}/sitegroups/{sitegroup_id} | getOrgSiteGroup
[**ListOrgSiteGroups**](OrgsSitegroupsApi.md#ListOrgSiteGroups) | **Get** /api/v1/orgs/{org_id}/sitegroups | listOrgSiteGroups
[**UpdateOrgSiteGroup**](OrgsSitegroupsApi.md#UpdateOrgSiteGroup) | **Put** /api/v1/orgs/{org_id}/sitegroups/{sitegroup_id} | updateOrgSiteGroup

# **CreateOrgSiteGroup**
> InlineResponse20097 CreateOrgSiteGroup(ctx, orgId, optional)
createOrgSiteGroup

Create Org Site Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSitegroupsApiCreateOrgSiteGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSitegroupsApiCreateOrgSiteGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdSitegroupsBody**](OrgIdSitegroupsBody.md)| Request Body | 

### Return type

[**InlineResponse20097**](inline_response_200_97.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgSiteGroup**
> DeleteOrgSiteGroup(ctx, orgId, sitegroupId)
deleteOrgSiteGroup

Delete Org Site Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **sitegroupId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgSiteGroup**
> InlineResponse20097 GetOrgSiteGroup(ctx, orgId, sitegroupId)
getOrgSiteGroup

Get Org Site Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **sitegroupId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20097**](inline_response_200_97.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgSiteGroups**
> []Sitegroup ListOrgSiteGroups(ctx, orgId, optional)
listOrgSiteGroups

Get List of Org Site Groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSitegroupsApiListOrgSiteGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSitegroupsApiListOrgSiteGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Sitegroup**](sitegroup.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgSiteGroup**
> InlineResponse20097 UpdateOrgSiteGroup(ctx, orgId, sitegroupId, optional)
updateOrgSiteGroup

Update Org Site Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **sitegroupId** | [**string**](.md)|  | 
 **optional** | ***OrgsSitegroupsApiUpdateOrgSiteGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSitegroupsApiUpdateOrgSiteGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SitegroupsSitegroupIdBody**](SitegroupsSitegroupIdBody.md)| Request Body | 

### Return type

[**InlineResponse20097**](inline_response_200_97.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

