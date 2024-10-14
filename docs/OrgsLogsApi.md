# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgAuditLogs**](OrgsLogsApi.md#CountOrgAuditLogs) | **Get** /api/v1/orgs/{org_id}/logs/count | countOrgAuditLogs
[**ListOrgAuditLogs**](OrgsLogsApi.md#ListOrgAuditLogs) | **Get** /api/v1/orgs/{org_id}/logs | listOrgAuditLogs

# **CountOrgAuditLogs**
> InlineResponse20016 CountOrgAuditLogs(ctx, orgId, optional)
countOrgAuditLogs

Count by Distinct Attributes of Audit Logs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsLogsApiCountOrgAuditLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsLogsApiCountOrgAuditLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of OrgLogsCountDistinct**](.md)|  | 
 **adminId** | **optional.String**|  | 
 **adminName** | **optional.String**|  | 
 **siteId** | **optional.String**|  | 
 **message** | **optional.String**|  | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgAuditLogs**
> InlineResponse20015 ListOrgAuditLogs(ctx, orgId, optional)
listOrgAuditLogs

Get List of change logs for the current Org

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsLogsApiListOrgAuditLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsLogsApiListOrgAuditLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | **optional.String**| site id | 
 **adminName** | **optional.String**| admin name or email | 
 **message** | **optional.String**| message | 
 **sort** | [**optional.Interface of Sort1**](.md)| sort order | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

