# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMspOrgGroup**](MSPsOrgGroupsApi.md#CreateMspOrgGroup) | **Post** /api/v1/msps/{msp_id}/orggroups | createMspOrgGroup
[**DeleteMspOrgGroup**](MSPsOrgGroupsApi.md#DeleteMspOrgGroup) | **Delete** /api/v1/msps/{msp_id}/orggroups/{orggroup_id} | deleteMspOrgGroup
[**GetMspOrgGroup**](MSPsOrgGroupsApi.md#GetMspOrgGroup) | **Get** /api/v1/msps/{msp_id}/orggroups/{orggroup_id} | getMspOrgGroup
[**ListMspOrgGroups**](MSPsOrgGroupsApi.md#ListMspOrgGroups) | **Get** /api/v1/msps/{msp_id}/orggroups | listMspOrgGroups
[**UpdateMspOrgGroup**](MSPsOrgGroupsApi.md#UpdateMspOrgGroup) | **Put** /api/v1/msps/{msp_id}/orggroups/{orggroup_id} | updateMspOrgGroup

# **CreateMspOrgGroup**
> InlineResponse20020 CreateMspOrgGroup(ctx, mspId, optional)
createMspOrgGroup

Create MSP Org Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
 **optional** | ***MSPsOrgGroupsApiCreateMspOrgGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsOrgGroupsApiCreateMspOrgGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MspIdOrggroupsBody**](MspIdOrggroupsBody.md)| Request Body | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMspOrgGroup**
> DeleteMspOrgGroup(ctx, mspId, orggroupId)
deleteMspOrgGroup

Delete MSP Org Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **orggroupId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMspOrgGroup**
> InlineResponse20020 GetMspOrgGroup(ctx, mspId, orggroupId)
getMspOrgGroup

Get MSP Org Group Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **orggroupId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMspOrgGroups**
> []Orggroup ListMspOrgGroups(ctx, mspId)
listMspOrgGroups

Get List of MSP Org Groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 

### Return type

[**[]Orggroup**](orggroup.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMspOrgGroup**
> InlineResponse20020 UpdateMspOrgGroup(ctx, mspId, orggroupId, optional)
updateMspOrgGroup

Update MSP Org Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **orggroupId** | [**string**](.md)|  | 
 **optional** | ***MSPsOrgGroupsApiUpdateMspOrgGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsOrgGroupsApiUpdateMspOrgGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of OrggroupsOrggroupIdBody**](OrggroupsOrggroupIdBody.md)| Request Body | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

