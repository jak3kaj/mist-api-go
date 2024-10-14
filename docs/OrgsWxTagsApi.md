# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgWxTag**](OrgsWxTagsApi.md#CreateOrgWxTag) | **Post** /api/v1/orgs/{org_id}/wxtags | createOrgWxTag
[**DeleteOrgWxTag**](OrgsWxTagsApi.md#DeleteOrgWxTag) | **Delete** /api/v1/orgs/{org_id}/wxtags/{wxtag_id} | deleteOrgWxTag
[**GetOrgApplicationList**](OrgsWxTagsApi.md#GetOrgApplicationList) | **Get** /api/v1/orgs/{org_id}/wxtags/apps | getOrgApplicationList
[**GetOrgCurrentMatchingClientsOfAWxTag**](OrgsWxTagsApi.md#GetOrgCurrentMatchingClientsOfAWxTag) | **Get** /api/v1/orgs/{org_id}/wxtags/{wxtag_id}/clients | getOrgCurrentMatchingClientsOfAWxTag
[**GetOrgWxTag**](OrgsWxTagsApi.md#GetOrgWxTag) | **Get** /api/v1/orgs/{org_id}/wxtags/{wxtag_id} | getOrgWxTag
[**ListOrgWxTags**](OrgsWxTagsApi.md#ListOrgWxTags) | **Get** /api/v1/orgs/{org_id}/wxtags | listOrgWxTags
[**UpdateOrgWxTag**](OrgsWxTagsApi.md#UpdateOrgWxTag) | **Put** /api/v1/orgs/{org_id}/wxtags/{wxtag_id} | updateOrgWxTag

# **CreateOrgWxTag**
> InlineResponse200116 CreateOrgWxTag(ctx, orgId, optional)
createOrgWxTag

Create WxLAN Tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsWxTagsApiCreateOrgWxTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWxTagsApiCreateOrgWxTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdWxtagsBody**](OrgIdWxtagsBody.md)| Request Body | 

### Return type

[**InlineResponse200116**](inline_response_200_116.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgWxTag**
> DeleteOrgWxTag(ctx, orgId, wxtagId)
deleteOrgWxTag

Delete WxLAN Tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **wxtagId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgApplicationList**
> []SearchWxtagAppsItem GetOrgApplicationList(ctx, orgId)
getOrgApplicationList

Get Application List

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**[]SearchWxtagAppsItem**](search_wxtag_apps_item.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgCurrentMatchingClientsOfAWxTag**
> []WxtagClient GetOrgCurrentMatchingClientsOfAWxTag(ctx, orgId, wxtagId)
getOrgCurrentMatchingClientsOfAWxTag

Get Current Matching Clients of a WXLAN Tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **wxtagId** | [**string**](.md)|  | 

### Return type

[**[]WxtagClient**](wxtag_client.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgWxTag**
> InlineResponse200116 GetOrgWxTag(ctx, orgId, wxtagId)
getOrgWxTag

Get WxLAN Tag Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **wxtagId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200116**](inline_response_200_116.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgWxTags**
> []WxlanTag ListOrgWxTags(ctx, orgId, optional)
listOrgWxTags

Get List of Org WxLAN Tags

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsWxTagsApiListOrgWxTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWxTagsApiListOrgWxTagsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]WxlanTag**](wxlan_tag.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgWxTag**
> InlineResponse200116 UpdateOrgWxTag(ctx, orgId, wxtagId, optional)
updateOrgWxTag

Update WxLAN Tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **wxtagId** | [**string**](.md)|  | 
 **optional** | ***OrgsWxTagsApiUpdateOrgWxTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWxTagsApiUpdateOrgWxTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of WxtagsWxtagIdBody**](WxtagsWxtagIdBody.md)| Request Body | 

### Return type

[**InlineResponse200116**](inline_response_200_116.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

