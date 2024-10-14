# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgNacTag**](OrgsNACTagsApi.md#CreateOrgNacTag) | **Post** /api/v1/orgs/{org_id}/nactags | createOrgNacTag
[**DeleteOrgNacTag**](OrgsNACTagsApi.md#DeleteOrgNacTag) | **Delete** /api/v1/orgs/{org_id}/nactags/{nactag_id} | deleteOrgNacTag
[**GetOrgNacTag**](OrgsNACTagsApi.md#GetOrgNacTag) | **Get** /api/v1/orgs/{org_id}/nactags/{nactag_id} | getOrgNacTag
[**ListOrgNacTags**](OrgsNACTagsApi.md#ListOrgNacTags) | **Get** /api/v1/orgs/{org_id}/nactags | listOrgNacTags
[**UpdateOrgNacTag**](OrgsNACTagsApi.md#UpdateOrgNacTag) | **Put** /api/v1/orgs/{org_id}/nactags/{nactag_id} | updateOrgNacTag

# **CreateOrgNacTag**
> InlineResponse20072 CreateOrgNacTag(ctx, orgId, optional)
createOrgNacTag

Create Org NAC Tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsNACTagsApiCreateOrgNacTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNACTagsApiCreateOrgNacTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NacTag**](NacTag.md)|  | 

### Return type

[**InlineResponse20072**](inline_response_200_72.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgNacTag**
> DeleteOrgNacTag(ctx, orgId, nactagId)
deleteOrgNacTag

Delete Org NAC Tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **nactagId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgNacTag**
> InlineResponse20072 GetOrgNacTag(ctx, orgId, nactagId)
getOrgNacTag

Get Org NAC Tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **nactagId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20072**](inline_response_200_72.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgNacTags**
> []NacTag ListOrgNacTags(ctx, orgId, optional)
listOrgNacTags

Get List of Org NAC Tags

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsNACTagsApiListOrgNacTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNACTagsApiListOrgNacTagsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| Type of NAC Tag | 
 **name** | **optional.String**| Name of NAC Tag | 
 **match** | **optional.String**| Type of NAC Tag | 
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]NacTag**](nac_tag.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgNacTag**
> InlineResponse20072 UpdateOrgNacTag(ctx, orgId, nactagId, optional)
updateOrgNacTag

Update Org NAC Tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **nactagId** | [**string**](.md)|  | 
 **optional** | ***OrgsNACTagsApiUpdateOrgNacTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsNACTagsApiUpdateOrgNacTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of NacTag**](NacTag.md)|  | 

### Return type

[**InlineResponse20072**](inline_response_200_72.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

