# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgMxEdgeCluster**](OrgsMxClustersApi.md#CreateOrgMxEdgeCluster) | **Post** /api/v1/orgs/{org_id}/mxclusters | createOrgMxEdgeCluster
[**DeleteOrgMxEdgeCluster**](OrgsMxClustersApi.md#DeleteOrgMxEdgeCluster) | **Delete** /api/v1/orgs/{org_id}/mxclusters/{mxcluster_id} | deleteOrgMxEdgeCluster
[**GetOrgMxEdgeCluster**](OrgsMxClustersApi.md#GetOrgMxEdgeCluster) | **Get** /api/v1/orgs/{org_id}/mxclusters/{mxcluster_id} | getOrgMxEdgeCluster
[**ListOrgMxEdgeClusters**](OrgsMxClustersApi.md#ListOrgMxEdgeClusters) | **Get** /api/v1/orgs/{org_id}/mxclusters | listOrgMxEdgeClusters
[**UpdateOrgMxEdgeCluster**](OrgsMxClustersApi.md#UpdateOrgMxEdgeCluster) | **Put** /api/v1/orgs/{org_id}/mxclusters/{mxcluster_id} | updateOrgMxEdgeCluster

# **CreateOrgMxEdgeCluster**
> InlineResponse20066 CreateOrgMxEdgeCluster(ctx, orgId, optional)
createOrgMxEdgeCluster

Create MxCluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsMxClustersApiCreateOrgMxEdgeClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMxClustersApiCreateOrgMxEdgeClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdMxclustersBody**](OrgIdMxclustersBody.md)| Request Body | 

### Return type

[**InlineResponse20066**](inline_response_200_66.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgMxEdgeCluster**
> DeleteOrgMxEdgeCluster(ctx, orgId, mxclusterId)
deleteOrgMxEdgeCluster

Delete Org MXEdge Cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **mxclusterId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgMxEdgeCluster**
> InlineResponse20066 GetOrgMxEdgeCluster(ctx, orgId, mxclusterId)
getOrgMxEdgeCluster

Get Org MxEdge Cluster Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **mxclusterId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20066**](inline_response_200_66.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgMxEdgeClusters**
> []Mxcluster ListOrgMxEdgeClusters(ctx, orgId, optional)
listOrgMxEdgeClusters

Get List of Org MxEdge Clusters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsMxClustersApiListOrgMxEdgeClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMxClustersApiListOrgMxEdgeClustersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Mxcluster**](mxcluster.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgMxEdgeCluster**
> InlineResponse20066 UpdateOrgMxEdgeCluster(ctx, orgId, mxclusterId, optional)
updateOrgMxEdgeCluster

Update Org MxEdge Cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **mxclusterId** | [**string**](.md)|  | 
 **optional** | ***OrgsMxClustersApiUpdateOrgMxEdgeClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsMxClustersApiUpdateOrgMxEdgeClusterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of MxclustersMxclusterIdBody**](MxclustersMxclusterIdBody.md)| Request Body | 

### Return type

[**InlineResponse20066**](inline_response_200_66.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

