# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMspOrg**](MSPsOrgsApi.md#CreateMspOrg) | **Post** /api/v1/msps/{msp_id}/orgs | createMspOrg
[**DeleteMspOrg**](MSPsOrgsApi.md#DeleteMspOrg) | **Delete** /api/v1/msps/{msp_id}/orgs/{org_id} | deleteMspOrg
[**GetMspOrg**](MSPsOrgsApi.md#GetMspOrg) | **Get** /api/v1/msps/{msp_id}/orgs/{org_id} | getMspOrg
[**ListMspOrgStats**](MSPsOrgsApi.md#ListMspOrgStats) | **Get** /api/v1/msps/{msp_id}/stats/orgs | listMspOrgStats
[**ListMspOrgs**](MSPsOrgsApi.md#ListMspOrgs) | **Get** /api/v1/msps/{msp_id}/orgs | listMspOrgs
[**ManageMspOrgs**](MSPsOrgsApi.md#ManageMspOrgs) | **Put** /api/v1/msps/{msp_id}/orgs | manageMspOrgs
[**SearchMspOrgs**](MSPsOrgsApi.md#SearchMspOrgs) | **Get** /api/v1/msps/{msp_id}/orgs/search | searchMspOrgs
[**UpdateMspOrg**](MSPsOrgsApi.md#UpdateMspOrg) | **Put** /api/v1/msps/{msp_id}/orgs/{org_id} | updateMspOrg

# **CreateMspOrg**
> InlineResponse20021 CreateMspOrg(ctx, mspId, optional)
createMspOrg

Create an Org under MSP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
 **optional** | ***MSPsOrgsApiCreateMspOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsOrgsApiCreateMspOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MspIdOrgsBody1**](MspIdOrgsBody1.md)| Request Body | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMspOrg**
> DeleteMspOrg(ctx, mspId, orgId)
deleteMspOrg

delete MSP Org

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **orgId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMspOrg**
> InlineResponse20021 GetMspOrg(ctx, mspId, orgId)
getMspOrg

Get MSP Org Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **orgId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMspOrgStats**
> []StatsOrg ListMspOrgStats(ctx, mspId, optional)
listMspOrgStats

Get List of MSP Orgs Stats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
 **optional** | ***MSPsOrgsApiListMspOrgStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsOrgsApiListMspOrgStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]StatsOrg**](stats_org.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMspOrgs**
> []Org ListMspOrgs(ctx, mspId)
listMspOrgs

Get List of MSP Orgs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 

### Return type

[**[]Org**](org.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ManageMspOrgs**
> ManageMspOrgs(ctx, mspId, optional)
manageMspOrgs

Assign or Unassign Orgs to an MSP account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
 **optional** | ***MSPsOrgsApiManageMspOrgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsOrgsApiManageMspOrgsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MspIdOrgsBody**](MspIdOrgsBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchMspOrgs**
> InlineResponse20022 SearchMspOrgs(ctx, mspId, optional)
searchMspOrgs

Search Org in MSP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
 **optional** | ***MSPsOrgsApiSearchMspOrgsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsOrgsApiSearchMspOrgsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**|  | 
 **orgId** | [**optional.Interface of string**](.md)| org id | 
 **subInsufficient** | **optional.Bool**| if this org has sufficient subscription | 
 **trialEnabled** | **optional.Bool**| if this org is under trial period | 
 **usageTypes** | [**optional.Interface of []string**](string.md)| a list of types that enabled by usage | 
 **limit** | **optional.Int32**|  | [default to 100]

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMspOrg**
> InlineResponse20021 UpdateMspOrg(ctx, mspId, orgId, optional)
updateMspOrg

Update MSP Org

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **orgId** | [**string**](.md)|  | 
 **optional** | ***MSPsOrgsApiUpdateMspOrgOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsOrgsApiUpdateMspOrgOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of Org**](Org.md)|  | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

