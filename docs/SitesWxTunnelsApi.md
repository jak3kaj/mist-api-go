# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteWxTunnel**](SitesWxTunnelsApi.md#CreateSiteWxTunnel) | **Post** /api/v1/sites/{site_id}/wxtunnels | createSiteWxTunnel
[**DeleteSiteWxTunnel**](SitesWxTunnelsApi.md#DeleteSiteWxTunnel) | **Delete** /api/v1/sites/{site_id}/wxtunnels/{wxtunnel_id} | deleteSiteWxTunnel
[**GetSiteWxTunnel**](SitesWxTunnelsApi.md#GetSiteWxTunnel) | **Get** /api/v1/sites/{site_id}/wxtunnels/{wxtunnel_id} | getSiteWxTunnel
[**ListSiteWxTunnels**](SitesWxTunnelsApi.md#ListSiteWxTunnels) | **Get** /api/v1/sites/{site_id}/wxtunnels | listSiteWxTunnels
[**UpdateSiteWxTunnel**](SitesWxTunnelsApi.md#UpdateSiteWxTunnel) | **Put** /api/v1/sites/{site_id}/wxtunnels/{wxtunnel_id} | updateSiteWxTunnel

# **CreateSiteWxTunnel**
> InlineResponse200117 CreateSiteWxTunnel(ctx, siteId, optional)
createSiteWxTunnel

Create Site WxLan Tunnel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesWxTunnelsApiCreateSiteWxTunnelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWxTunnelsApiCreateSiteWxTunnelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SiteIdWxtunnelsBody**](SiteIdWxtunnelsBody.md)| Request Body | 

### Return type

[**InlineResponse200117**](inline_response_200_117.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteWxTunnel**
> DeleteSiteWxTunnel(ctx, siteId, wxtunnelId)
deleteSiteWxTunnel

Delete Site WxLan Tunnel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **wxtunnelId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteWxTunnel**
> InlineResponse200117 GetSiteWxTunnel(ctx, siteId, wxtunnelId)
getSiteWxTunnel

Get Site WxLan tunnel Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **wxtunnelId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200117**](inline_response_200_117.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteWxTunnels**
> []WxlanTunnel ListSiteWxTunnels(ctx, siteId, optional)
listSiteWxTunnels

Get List of Site WxLan Tunnels

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesWxTunnelsApiListSiteWxTunnelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWxTunnelsApiListSiteWxTunnelsOpts struct
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

# **UpdateSiteWxTunnel**
> InlineResponse200117 UpdateSiteWxTunnel(ctx, siteId, wxtunnelId, optional)
updateSiteWxTunnel

Update Site WxLan Tunnel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **wxtunnelId** | [**string**](.md)|  | 
 **optional** | ***SitesWxTunnelsApiUpdateSiteWxTunnelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWxTunnelsApiUpdateSiteWxTunnelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of WxtunnelsWxtunnelIdBody1**](WxtunnelsWxtunnelIdBody1.md)| Request Body | 

### Return type

[**InlineResponse200117**](inline_response_200_117.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

