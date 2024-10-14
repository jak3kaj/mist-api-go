# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgWxTunnel**](OrgsWxTunnelsApi.md#CreateOrgWxTunnel) | **Post** /api/v1/orgs/{org_id}/wxtunnels | createOrgWxTunnel
[**DeleteOrgWxTunnel**](OrgsWxTunnelsApi.md#DeleteOrgWxTunnel) | **Delete** /api/v1/orgs/{org_id}/wxtunnels/{wxtunnel_id} | deleteOrgWxTunnel
[**GetOrgWxTunnel**](OrgsWxTunnelsApi.md#GetOrgWxTunnel) | **Get** /api/v1/orgs/{org_id}/wxtunnels/{wxtunnel_id} | getOrgWxTunnel
[**ListOrgWxTunnels**](OrgsWxTunnelsApi.md#ListOrgWxTunnels) | **Get** /api/v1/orgs/{org_id}/wxtunnels | listOrgWxTunnels
[**UpdateOrgWxTunnel**](OrgsWxTunnelsApi.md#UpdateOrgWxTunnel) | **Put** /api/v1/orgs/{org_id}/wxtunnels/{wxtunnel_id} | updateOrgWxTunnel

# **CreateOrgWxTunnel**
> InlineResponse200117 CreateOrgWxTunnel(ctx, orgId, optional)
createOrgWxTunnel

Create Org WxAN Tunnel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsWxTunnelsApiCreateOrgWxTunnelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWxTunnelsApiCreateOrgWxTunnelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdWxtunnelsBody**](OrgIdWxtunnelsBody.md)| Request Body | 

### Return type

[**InlineResponse200117**](inline_response_200_117.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgWxTunnel**
> DeleteOrgWxTunnel(ctx, orgId, wxtunnelId)
deleteOrgWxTunnel

Delete Org WxLAN Tunnel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **wxtunnelId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgWxTunnel**
> InlineResponse200117 GetOrgWxTunnel(ctx, orgId, wxtunnelId)
getOrgWxTunnel

Get Org WxLAN Tunnel Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **wxtunnelId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200117**](inline_response_200_117.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgWxTunnels**
> []WxlanTunnel ListOrgWxTunnels(ctx, orgId, optional)
listOrgWxTunnels

Get List of Org WxLAN Tunnels

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsWxTunnelsApiListOrgWxTunnelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWxTunnelsApiListOrgWxTunnelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]WxlanTunnel**](wxlan_tunnel.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgWxTunnel**
> InlineResponse200117 UpdateOrgWxTunnel(ctx, orgId, wxtunnelId, optional)
updateOrgWxTunnel

Update Org WxLAN Tunnel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **wxtunnelId** | [**string**](.md)|  | 
 **optional** | ***OrgsWxTunnelsApiUpdateOrgWxTunnelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWxTunnelsApiUpdateOrgWxTunnelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of WxtunnelsWxtunnelIdBody**](WxtunnelsWxtunnelIdBody.md)| Request Body | 

### Return type

[**InlineResponse200117**](inline_response_200_117.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

