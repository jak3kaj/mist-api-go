# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgNetwork**](OrgsNetworksApi.md#CreateOrgNetwork) | **Post** /api/v1/orgs/{org_id}/networks | createOrgNetwork
[**DeleteOrgNetwork**](OrgsNetworksApi.md#DeleteOrgNetwork) | **Delete** /api/v1/orgs/{org_id}/networks/{network_id} | deleteOrgNetwork
[**GetOrgNetwork**](OrgsNetworksApi.md#GetOrgNetwork) | **Get** /api/v1/orgs/{org_id}/networks/{network_id} | getOrgNetwork
[**ListOrgNetworks**](OrgsNetworksApi.md#ListOrgNetworks) | **Get** /api/v1/orgs/{org_id}/networks | listOrgNetworks
[**UpdateOrgNetwork**](OrgsNetworksApi.md#UpdateOrgNetwork) | **Put** /api/v1/orgs/{org_id}/networks/{network_id} | updateOrgNetwork

# **CreateOrgNetwork**
> InlineResponse20084 CreateOrgNetwork(ctx, orgId, optional)
createOrgNetwork

Create Organization Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsNetworksApiCreateOrgNetworkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNetworksApiCreateOrgNetworkOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdNetworksBody**](OrgIdNetworksBody.md)|  | 

### Return type

[**InlineResponse20084**](inline_response_200_84.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgNetwork**
> DeleteOrgNetwork(ctx, orgId, networkId)
deleteOrgNetwork

Delete Organization Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **networkId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgNetwork**
> InlineResponse20084 GetOrgNetwork(ctx, orgId, networkId)
getOrgNetwork

Get Organization Network Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **networkId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20084**](inline_response_200_84.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgNetworks**
> []Network ListOrgNetworks(ctx, orgId, optional)
listOrgNetworks

Get List of Org Networks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsNetworksApiListOrgNetworksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNetworksApiListOrgNetworksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Network**](network.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgNetwork**
> InlineResponse20084 UpdateOrgNetwork(ctx, orgId, networkId, optional)
updateOrgNetwork

Update Organization Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **networkId** | [**string**](.md)|  | 
 **optional** | ***OrgsNetworksApiUpdateOrgNetworkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNetworksApiUpdateOrgNetworkOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of NetworksNetworkIdBody**](NetworksNetworkIdBody.md)|  | 

### Return type

[**InlineResponse20084**](inline_response_200_84.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

