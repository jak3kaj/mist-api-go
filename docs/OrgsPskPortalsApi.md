# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgPskPortalLogs**](OrgsPskPortalsApi.md#CountOrgPskPortalLogs) | **Get** /api/v1/orgs/{org_id}/pskportals/logs/count | countOrgPskPortalLogs
[**CreateOrgPskPortal**](OrgsPskPortalsApi.md#CreateOrgPskPortal) | **Post** /api/v1/orgs/{org_id}/pskportals | createOrgPskPortal
[**DeleteOrgPskPortal**](OrgsPskPortalsApi.md#DeleteOrgPskPortal) | **Delete** /api/v1/orgs/{org_id}/pskportals/{pskportal_id} | deleteOrgPskPortal
[**DeleteOrgPskPortalImage**](OrgsPskPortalsApi.md#DeleteOrgPskPortalImage) | **Delete** /api/v1/orgs/{org_id}/pskportals/{pskportal_id}/portal_image | deleteOrgPskPortalImage
[**GetOrgPskPortal**](OrgsPskPortalsApi.md#GetOrgPskPortal) | **Get** /api/v1/orgs/{org_id}/pskportals/{pskportal_id} | getOrgPskPortal
[**ListOrgPskPortalLogs**](OrgsPskPortalsApi.md#ListOrgPskPortalLogs) | **Get** /api/v1/orgs/{org_id}/pskportals/logs | listOrgPskPortalLogs
[**ListOrgPskPortals**](OrgsPskPortalsApi.md#ListOrgPskPortals) | **Get** /api/v1/orgs/{org_id}/pskportals | listOrgPskPortals
[**SearchOrgPskPortalLogs**](OrgsPskPortalsApi.md#SearchOrgPskPortalLogs) | **Get** /api/v1/orgs/{org_id}/pskportals/logs/search | searchOrgPskPortalLogs
[**UpdateOrgPskPortal**](OrgsPskPortalsApi.md#UpdateOrgPskPortal) | **Put** /api/v1/orgs/{org_id}/pskportals/{pskportal_id} | updateOrgPskPortal
[**UpdateOrgPskPortalTemplate**](OrgsPskPortalsApi.md#UpdateOrgPskPortalTemplate) | **Put** /api/v1/orgs/{org_id}/pskportals/{pskportal_id}/portal_template | updateOrgPskPortalTemplate
[**UploadOrgPskPortalImage**](OrgsPskPortalsApi.md#UploadOrgPskPortalImage) | **Post** /api/v1/orgs/{org_id}/pskportals/{pskportal_id}/portal_image | uploadOrgPskPortalImage

# **CountOrgPskPortalLogs**
> InlineResponse20016 CountOrgPskPortalLogs(ctx, orgId, optional)
countOrgPskPortalLogs

Count by Distinct Attributes of PskPortal Logs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsPskPortalsApiCountOrgPskPortalLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsPskPortalsApiCountOrgPskPortalLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of OrgPskPortalLogsCountDistinct**](.md)|  | 
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

# **CreateOrgPskPortal**
> PskPortal CreateOrgPskPortal(ctx, orgId, optional)
createOrgPskPortal

Create Org Psk Portal

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsPskPortalsApiCreateOrgPskPortalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsPskPortalsApiCreateOrgPskPortalOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PskPortal**](PskPortal.md)|  | 

### Return type

[**PskPortal**](psk_portal.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgPskPortal**
> DeleteOrgPskPortal(ctx, orgId, pskportalId)
deleteOrgPskPortal

Delete Org Psk Portal

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **pskportalId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgPskPortalImage**
> DeleteOrgPskPortalImage(ctx, orgId, pskportalId)
deleteOrgPskPortalImage

Delete background image for PskPortal   If image is not uploaded or is deleted, PskPortal will use default image.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **pskportalId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgPskPortal**
> PskPortal GetOrgPskPortal(ctx, orgId, pskportalId)
getOrgPskPortal

get Org Psk Portal Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **pskportalId** | [**string**](.md)|  | 

### Return type

[**PskPortal**](psk_portal.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgPskPortalLogs**
> InlineResponse20086 ListOrgPskPortalLogs(ctx, orgId, optional)
listOrgPskPortalLogs

Get the list of PSK Portals Logs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsPskPortalsApiListOrgPskPortalLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsPskPortalsApiListOrgPskPortalLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse20086**](inline_response_200_86.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgPskPortals**
> []PskPortal ListOrgPskPortals(ctx, orgId, optional)
listOrgPskPortals

Get List of Org Psk Portals

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsPskPortalsApiListOrgPskPortalsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsPskPortalsApiListOrgPskPortalsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]PskPortal**](psk_portal.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrgPskPortalLogs**
> InlineResponse20086 SearchOrgPskPortalLogs(ctx, orgId, optional)
searchOrgPskPortalLogs

Search Org PSK Portal Logs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsPskPortalsApiSearchOrgPskPortalLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsPskPortalsApiSearchOrgPskPortalLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]
 **pskName** | **optional.String**|  | 
 **pskId** | **optional.String**|  | 
 **pskportalId** | **optional.String**|  | 
 **id** | [**optional.Interface of string**](.md)| audit_id | 
 **adminName** | **optional.String**|  | 
 **adminId** | **optional.String**|  | 
 **nameId** | [**optional.Interface of string**](.md)| name_id used in SSO | 

### Return type

[**InlineResponse20086**](inline_response_200_86.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgPskPortal**
> PskPortal UpdateOrgPskPortal(ctx, orgId, pskportalId, optional)
updateOrgPskPortal

update Org Psk Portal

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **pskportalId** | [**string**](.md)|  | 
 **optional** | ***OrgsPskPortalsApiUpdateOrgPskPortalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsPskPortalsApiUpdateOrgPskPortalOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PskPortal**](PskPortal.md)|  | 

### Return type

[**PskPortal**](psk_portal.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgPskPortalTemplate**
> UpdateOrgPskPortalTemplate(ctx, orgId, pskportalId, optional)
updateOrgPskPortalTemplate

Update Org Psk Portal Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **pskportalId** | [**string**](.md)|  | 
 **optional** | ***OrgsPskPortalsApiUpdateOrgPskPortalTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsPskPortalsApiUpdateOrgPskPortalTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PskPortalTemplate**](PskPortalTemplate.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadOrgPskPortalImage**
> UploadOrgPskPortalImage(ctx, orgId, pskportalId, optional)
uploadOrgPskPortalImage

Upload background image for PskPortal

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **pskportalId** | [**string**](.md)|  | 
 **optional** | ***OrgsPskPortalsApiUploadOrgPskPortalImageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsPskPortalsApiUploadOrgPskPortalImageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | **optional.Interface of *os.File****optional.**|  | 
 **json** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

