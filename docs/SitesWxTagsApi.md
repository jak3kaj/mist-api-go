# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteWxTag**](SitesWxTagsApi.md#CreateSiteWxTag) | **Post** /api/v1/sites/{site_id}/wxtags | createSiteWxTag
[**DeleteSiteWxTag**](SitesWxTagsApi.md#DeleteSiteWxTag) | **Delete** /api/v1/sites/{site_id}/wxtags/{wxtag_id} | deleteSiteWxTag
[**GetSiteApplicationList**](SitesWxTagsApi.md#GetSiteApplicationList) | **Get** /api/v1/sites/{site_id}/wxtags/apps | getSiteApplicationList
[**GetSiteCurrentMatchingClientsOfAWxTag**](SitesWxTagsApi.md#GetSiteCurrentMatchingClientsOfAWxTag) | **Get** /api/v1/sites/{site_id}/wxtags/{wxtag_id}/clients | getSiteCurrentMatchingClientsOfAWxTag
[**GetSiteWxTag**](SitesWxTagsApi.md#GetSiteWxTag) | **Get** /api/v1/sites/{site_id}/wxtags/{wxtag_id} | getSiteWxTag
[**ListSiteWxTags**](SitesWxTagsApi.md#ListSiteWxTags) | **Get** /api/v1/sites/{site_id}/wxtags | listSiteWxTags
[**UpdateSiteWxTag**](SitesWxTagsApi.md#UpdateSiteWxTag) | **Put** /api/v1/sites/{site_id}/wxtags/{wxtag_id} | updateSiteWxTag

# **CreateSiteWxTag**
> InlineResponse200116 CreateSiteWxTag(ctx, siteId, optional)
createSiteWxTag

Create Site WxTag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesWxTagsApiCreateSiteWxTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWxTagsApiCreateSiteWxTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SiteIdWxtagsBody**](SiteIdWxtagsBody.md)| Request Body | 

### Return type

[**InlineResponse200116**](inline_response_200_116.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteWxTag**
> DeleteSiteWxTag(ctx, siteId, wxtagId)
deleteSiteWxTag

Delete Site WxTag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **wxtagId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteApplicationList**
> []SearchWxtagAppsItem GetSiteApplicationList(ctx, siteId)
getSiteApplicationList

Get Application List

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

[**[]SearchWxtagAppsItem**](search_wxtag_apps_item.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteCurrentMatchingClientsOfAWxTag**
> []WxtagMatching GetSiteCurrentMatchingClientsOfAWxTag(ctx, siteId, wxtagId)
getSiteCurrentMatchingClientsOfAWxTag

Get Current Matching Clients of a WXLAN Tag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **wxtagId** | [**string**](.md)|  | 

### Return type

[**[]WxtagMatching**](wxtag_matching.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteWxTag**
> InlineResponse200116 GetSiteWxTag(ctx, siteId, wxtagId)
getSiteWxTag

Get Site WxTag Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **wxtagId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200116**](inline_response_200_116.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteWxTags**
> []WxlanTag ListSiteWxTags(ctx, siteId, optional)
listSiteWxTags

Get List of Site WxTags

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesWxTagsApiListSiteWxTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWxTagsApiListSiteWxTagsOpts struct
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

# **UpdateSiteWxTag**
> InlineResponse200116 UpdateSiteWxTag(ctx, siteId, wxtagId, optional)
updateSiteWxTag

Update Site WxTag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **wxtagId** | [**string**](.md)|  | 
 **optional** | ***SitesWxTagsApiUpdateSiteWxTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWxTagsApiUpdateSiteWxTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of WxtagsWxtagIdBody1**](WxtagsWxtagIdBody1.md)| Request Body | 

### Return type

[**InlineResponse200116**](inline_response_200_116.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

