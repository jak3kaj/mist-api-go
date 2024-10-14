# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgTunnelsStats**](OrgsStatsTunnelsApi.md#CountOrgTunnelsStats) | **Get** /api/v1/orgs/{org_id}/stats/tunnels/count | countOrgTunnelsStats
[**SearchOrgTunnelsStats**](OrgsStatsTunnelsApi.md#SearchOrgTunnelsStats) | **Get** /api/v1/orgs/{org_id}/stats/tunnels/search | searchOrgTunnelsStats

# **CountOrgTunnelsStats**
> InlineResponse20016 CountOrgTunnelsStats(ctx, orgId, optional)
countOrgTunnelsStats

Count Mist Tunnels Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsStatsTunnelsApiCountOrgTunnelsStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsStatsTunnelsApiCountOrgTunnelsStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of Distinct1**](.md)| - If &#x60;type&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60;: wxtunnel_id / ap / remote_ip / remote_port / state / mxedge_id / mxcluster_id / site_id / peer_mxedge_id; default is wxtunnel_id  - If &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;: mac / site_id / node / peer_ip / peer_host/ ip / tunnel_name / protocol / auth_algo / encrypt_algo / ike_version / last_event / up | 
 **type_** | [**optional.Interface of OrgTunnelTypeCount**](.md)|  | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrgTunnelsStats**
> InlineResponse20079 SearchOrgTunnelsStats(ctx, orgId, optional)
searchOrgTunnelsStats

Search Org Tunnels Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsStatsTunnelsApiSearchOrgTunnelsStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsStatsTunnelsApiSearchOrgTunnelsStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mxclusterId** | **optional.String**| if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60; | 
 **siteId** | **optional.String**|  | 
 **wxtunnelId** | **optional.String**| if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60; | 
 **ap** | **optional.String**| if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60; | 
 **mac** | **optional.String**| if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **node** | **optional.String**| if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **peerIp** | **optional.String**| if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **peerHost** | **optional.String**| if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **ip** | **optional.String**| if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **tunnelName** | **optional.String**| if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **protocol** | **optional.String**| if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **authAlgo** | **optional.String**| if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **encryptAlgo** | **optional.String**| if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **ikeVersion** | **optional.String**| if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **up** | **optional.String**| if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **type_** | [**optional.Interface of TunnelType**](.md)|  | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse20079**](inline_response_200_79.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

