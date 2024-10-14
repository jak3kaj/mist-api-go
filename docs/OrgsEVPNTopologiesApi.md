# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgEvpnTopology**](OrgsEVPNTopologiesApi.md#CreateOrgEvpnTopology) | **Post** /api/v1/orgs/{org_id}/evpn_topologies | createOrgEvpnTopology
[**DeleteOrgEvpnTopology**](OrgsEVPNTopologiesApi.md#DeleteOrgEvpnTopology) | **Delete** /api/v1/orgs/{org_id}/evpn_topologies/{evpn_topology_id} | deleteOrgEvpnTopology
[**GetOrgEvpnTolopogy**](OrgsEVPNTopologiesApi.md#GetOrgEvpnTolopogy) | **Get** /api/v1/orgs/{org_id}/evpn_topologies/{evpn_topology_id} | getOrgEvpnTolopogy
[**ListOrgEvpnTopologies**](OrgsEVPNTopologiesApi.md#ListOrgEvpnTopologies) | **Get** /api/v1/orgs/{org_id}/evpn_topologies | listOrgEvpnTopologies
[**UpdateOrgEvpnTopology**](OrgsEVPNTopologiesApi.md#UpdateOrgEvpnTopology) | **Put** /api/v1/orgs/{org_id}/evpn_topologies/{evpn_topology_id} | updateOrgEvpnTopology

# **CreateOrgEvpnTopology**
> InlineResponse20052 CreateOrgEvpnTopology(ctx, orgId, optional)
createOrgEvpnTopology

While all the `evpn_id` / `downlink_ips` can be specifidd by hand, the easiest way is to call the `build_vpn_topology` API, allowing you to examine the diff, and update it yourself. You can also simply call it with `overwrite=true` which will apply the updates for you.  **Notes:** 1. You can use `core` / `distribution` / `access` to create a CLOS topology 2. You can also use `core` / `distribution` to form a 2-tier EVPN topology where ESI-Lag is configured distribution to connect to access switches 3. In a small/medium campus, `collapsed-core` can be used where core switches are the inter-connected to do EVPN 4. The API uses a few pre-defined parameters and best-practices to generate the configs. It can be customized by using `evpn_options` in Site Setting / Network Template. (e.g. a different subnet for the underlay)  #### Collapsed Core In a small-medium campus, EVPN can also be enabled only at the core switches (up to 4) by assigning all participating switches with `collapsed-core role`. When there are more than 2 switches, a ring-like topology will be formed.  #### ESI-Lag If the access switchess does not have EVPN support, you can take advantage of EVPN by setting up ESI-Lag on distribution switches  #### Leaf / Access / Collapsed-Core For leaf nodes in a EVPN topology, you’d have to configure the IPs for networks that would participate in EVPN. Optionally, VRFs to isolate traffic from one tenant verus another

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsEVPNTopologiesApiCreateOrgEvpnTopologyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsEVPNTopologiesApiCreateOrgEvpnTopologyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdEvpnTopologiesBody**](OrgIdEvpnTopologiesBody.md)|  | 

### Return type

[**InlineResponse20052**](inline_response_200_52.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgEvpnTopology**
> DeleteOrgEvpnTopology(ctx, orgId, evpnTopologyId)
deleteOrgEvpnTopology

Delete the Org EVPN Topology

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **evpnTopologyId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgEvpnTolopogy**
> InlineResponse20052 GetOrgEvpnTolopogy(ctx, orgId, evpnTopologyId)
getOrgEvpnTolopogy

Get One EVPN Topology Detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **evpnTopologyId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20052**](inline_response_200_52.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgEvpnTopologies**
> []EvpnTopology ListOrgEvpnTopologies(ctx, orgId, optional)
listOrgEvpnTopologies

Get List of the existing Org EVPN topologies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsEVPNTopologiesApiListOrgEvpnTopologiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsEVPNTopologiesApiListOrgEvpnTopologiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]EvpnTopology**](evpn_topology.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgEvpnTopology**
> InlineResponse20052 UpdateOrgEvpnTopology(ctx, orgId, evpnTopologyId, optional)
updateOrgEvpnTopology

Update the EVPN Topolgy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **evpnTopologyId** | [**string**](.md)|  | 
 **optional** | ***OrgsEVPNTopologiesApiUpdateOrgEvpnTopologyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsEVPNTopologiesApiUpdateOrgEvpnTopologyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of EvpnTopologiesEvpnTopologyIdBody**](EvpnTopologiesEvpnTopologyIdBody.md)|  | 

### Return type

[**InlineResponse20052**](inline_response_200_52.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

