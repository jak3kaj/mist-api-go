# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgMxTunnel**](OrgsMxTunnelsApi.md#CreateOrgMxTunnel) | **Post** /api/v1/orgs/{org_id}/mxtunnels | createOrgMxTunnel
[**DeleteOrgMxTunnel**](OrgsMxTunnelsApi.md#DeleteOrgMxTunnel) | **Delete** /api/v1/orgs/{org_id}/mxtunnels/{mxtunnel_id} | deleteOrgMxTunnel
[**GetOrgMxTunnel**](OrgsMxTunnelsApi.md#GetOrgMxTunnel) | **Get** /api/v1/orgs/{org_id}/mxtunnels/{mxtunnel_id} | getOrgMxTunnel
[**ListOrgMxTunnels**](OrgsMxTunnelsApi.md#ListOrgMxTunnels) | **Get** /api/v1/orgs/{org_id}/mxtunnels | listOrgMxTunnels
[**UpdateOrgMxTunnel**](OrgsMxTunnelsApi.md#UpdateOrgMxTunnel) | **Put** /api/v1/orgs/{org_id}/mxtunnels/{mxtunnel_id} | updateOrgMxTunnel

# **CreateOrgMxTunnel**
> InlineResponse20071 CreateOrgMxTunnel(ctx, orgId, optional)
createOrgMxTunnel

Create MxTunnel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsMxTunnelsApiCreateOrgMxTunnelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMxTunnelsApiCreateOrgMxTunnelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdMxtunnelsBody**](OrgIdMxtunnelsBody.md)| Request Body | 

### Return type

[**InlineResponse20071**](inline_response_200_71.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgMxTunnel**
> DeleteOrgMxTunnel(ctx, orgId, mxtunnelId)
deleteOrgMxTunnel

Delete Org MxTunnel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **mxtunnelId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgMxTunnel**
> InlineResponse20071 GetOrgMxTunnel(ctx, orgId, mxtunnelId)
getOrgMxTunnel

Get Org MxTunnel Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **mxtunnelId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20071**](inline_response_200_71.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgMxTunnels**
> []Mxtunnel ListOrgMxTunnels(ctx, orgId, optional)
listOrgMxTunnels

Get List of Org MxTunnels

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsMxTunnelsApiListOrgMxTunnelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMxTunnelsApiListOrgMxTunnelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Mxtunnel**](mxtunnel.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgMxTunnel**
> InlineResponse20071 UpdateOrgMxTunnel(ctx, orgId, mxtunnelId, optional)
updateOrgMxTunnel

Update Org MxTunnel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **mxtunnelId** | [**string**](.md)|  | 
 **optional** | ***OrgsMxTunnelsApiUpdateOrgMxTunnelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMxTunnelsApiUpdateOrgMxTunnelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of MxtunnelsMxtunnelIdBody**](MxtunnelsMxtunnelIdBody.md)| Request Body | 

### Return type

[**InlineResponse20071**](inline_response_200_71.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

