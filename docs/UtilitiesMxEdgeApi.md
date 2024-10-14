# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreemptSitesMxTunnel**](UtilitiesMxEdgeApi.md#PreemptSitesMxTunnel) | **Post** /api/v1/sites/{site_id}/mxtunnels/{mxtunnel_id}/preempt_aps | preemptSitesMxTunnel

# **PreemptSitesMxTunnel**
> ResponseMxtunnelsPreemptAps PreemptSitesMxTunnel(ctx, siteId, mxtunnelId)
preemptSitesMxTunnel

To preempt AP’s which are not connected to preferred peer to the preferred peer

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **mxtunnelId** | [**string**](.md)|  | 

### Return type

[**ResponseMxtunnelsPreemptAps**](response_mxtunnels_preempt_aps.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

