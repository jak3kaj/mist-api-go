# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgBgpStats**](OrgsStatsBGPPeersApi.md#CountOrgBgpStats) | **Get** /api/v1/orgs/{org_id}/stats/bgp_peers/count | countOrgBgpStats
[**SearchOrgBgpStats**](OrgsStatsBGPPeersApi.md#SearchOrgBgpStats) | **Get** /api/v1/orgs/{org_id}/stats/bgp_peers/search | searchOrgBgpStats

# **CountOrgBgpStats**
> InlineResponse20016 CountOrgBgpStats(ctx, orgId)
countOrgBgpStats

Count Org BGP Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrgBgpStats**
> InlineResponse20076 SearchOrgBgpStats(ctx, orgId)
searchOrgBgpStats

Search Org BGP Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20076**](inline_response_200_76.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

