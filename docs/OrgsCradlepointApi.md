# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgCradlepointConnection**](OrgsCradlepointApi.md#DeleteOrgCradlepointConnection) | **Delete** /api/v1/orgs/{org_id}/setting/cradlepoint/setup | deleteOrgCradlepointConnection
[**SetupOrgCradlepointConnectionToMist**](OrgsCradlepointApi.md#SetupOrgCradlepointConnectionToMist) | **Post** /api/v1/orgs/{org_id}/setting/cradlepoint/setup | setupOrgCradlepointConnectionToMist
[**SyncOrgCradlepointRouters**](OrgsCradlepointApi.md#SyncOrgCradlepointRouters) | **Post** /api/v1/orgs/{org_id}/setting/cradlepoint/sync | syncOrgCradlepointRouters
[**UpdateOrgCradlepointConnectionToMist**](OrgsCradlepointApi.md#UpdateOrgCradlepointConnectionToMist) | **Put** /api/v1/orgs/{org_id}/setting/cradlepoint/setup | updateOrgCradlepointConnectionToMist

# **DeleteOrgCradlepointConnection**
> DeleteOrgCradlepointConnection(ctx, orgId)
deleteOrgCradlepointConnection

This deletes the Cradlepoint integration in Mist

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetupOrgCradlepointConnectionToMist**
> SetupOrgCradlepointConnectionToMist(ctx, orgId, optional)
setupOrgCradlepointConnectionToMist

This sets up cradlepoint webhooks to send events to Mist

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsCradlepointApiSetupOrgCradlepointConnectionToMistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsCradlepointApiSetupOrgCradlepointConnectionToMistOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AccountCradlepointConfig**](AccountCradlepointConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncOrgCradlepointRouters**
> SyncOrgCradlepointRouters(ctx, orgId)
syncOrgCradlepointRouters

This syncs cradlepoint devices with Mist. We’ll also attempt to use the LLDP data from cradlepoint to identify the linkage against Mist Site / Device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgCradlepointConnectionToMist**
> UpdateOrgCradlepointConnectionToMist(ctx, orgId, optional)
updateOrgCradlepointConnectionToMist

This updates the Cradlepoint integration settings in Mist

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsCradlepointApiUpdateOrgCradlepointConnectionToMistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsCradlepointApiUpdateOrgCradlepointConnectionToMistOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CradlepointSetupBody**](CradlepointSetupBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

