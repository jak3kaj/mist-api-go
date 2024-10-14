# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListOrgPmaDashboards**](OrgsPremiumAnalyticsApi.md#ListOrgPmaDashboards) | **Get** /api/v1/orgs/{org_id}/pma/dashboards | listOrgPmaDashboards

# **ListOrgPmaDashboards**
> []PmaDashboard ListOrgPmaDashboards(ctx, orgId, optional)
listOrgPmaDashboards

Get List of premium analytics dashboards for this Org

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsPremiumAnalyticsApiListOrgPmaDashboardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsPremiumAnalyticsApiListOrgPmaDashboardsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]PmaDashboard**](pma_dashboard.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

