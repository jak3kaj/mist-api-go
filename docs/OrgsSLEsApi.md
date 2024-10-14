# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgSitesSle**](OrgsSLEsApi.md#GetOrgSitesSle) | **Get** /api/v1/orgs/{org_id}/insights/sites-sle | getOrgSitesSle
[**GetOrgSle**](OrgsSLEsApi.md#GetOrgSle) | **Get** /api/v1/orgs/{org_id}/insights/{metric} | getOrgSle

# **GetOrgSitesSle**
> ResponseOrgSiteSle GetOrgSitesSle(ctx, orgId, optional)
getOrgSitesSle

Get Org Sites SLE

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSLEsApiGetOrgSitesSleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSLEsApiGetOrgSitesSleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sle** | [**optional.Interface of OrgSiteSleType**](.md)|  | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **interval** | **optional.String**| Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. | 
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**ResponseOrgSiteSle**](response_org_site_sle.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgSle**
> InlineResponse200101 GetOrgSle(ctx, orgId, metric, optional)
getOrgSle

Get Org SLEs (all/worst sites, Mx Edges, ...)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **metric** | **string**| see /api/v1/const/insight_metrics for available metrics | 
 **optional** | ***OrgsSLEsApiGetOrgSleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSLEsApiGetOrgSleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sle** | **optional.String**| see [listInsightMetrics]($e/Constants%20Definitions/listInsightMetrics) for more details | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **interval** | **optional.String**| Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 

### Return type

[**InlineResponse200101**](inline_response_200_101.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

