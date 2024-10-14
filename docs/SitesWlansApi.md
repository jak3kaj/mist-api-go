# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteWlan**](SitesWlansApi.md#CreateSiteWlan) | **Post** /api/v1/sites/{site_id}/wlans | createSiteWlan
[**DeleteSiteWlan**](SitesWlansApi.md#DeleteSiteWlan) | **Delete** /api/v1/sites/{site_id}/wlans/{wlan_id} | deleteSiteWlan
[**DeleteSiteWlanPortalImage**](SitesWlansApi.md#DeleteSiteWlanPortalImage) | **Delete** /api/v1/sites/{site_id}/wlans/{wlan_id}/portal_image | deleteSiteWlanPortalImage
[**GetSiteWlan**](SitesWlansApi.md#GetSiteWlan) | **Get** /api/v1/sites/{site_id}/wlans/{wlan_id} | getSiteWlan
[**ListSiteWlanDerived**](SitesWlansApi.md#ListSiteWlanDerived) | **Get** /api/v1/sites/{site_id}/wlans/derived | listSiteWlanDerived
[**ListSiteWlans**](SitesWlansApi.md#ListSiteWlans) | **Get** /api/v1/sites/{site_id}/wlans | listSiteWlans
[**UpdateSiteWlan**](SitesWlansApi.md#UpdateSiteWlan) | **Put** /api/v1/sites/{site_id}/wlans/{wlan_id} | updateSiteWlan
[**UpdateSiteWlanPortalTemplate**](SitesWlansApi.md#UpdateSiteWlanPortalTemplate) | **Put** /api/v1/sites/{site_id}/wlans/{wlan_id}/portal_template | updateSiteWlanPortalTemplate
[**UploadSiteWlanPortalImage**](SitesWlansApi.md#UploadSiteWlanPortalImage) | **Post** /api/v1/sites/{site_id}/wlans/{wlan_id}/portal_image | uploadSiteWlanPortalImage

# **CreateSiteWlan**
> InlineResponse200113 CreateSiteWlan(ctx, siteId, optional)
createSiteWlan

Create Site WLAN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesWlansApiCreateSiteWlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWlansApiCreateSiteWlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SiteIdWlansBody**](SiteIdWlansBody.md)| Request Body | 

### Return type

[**InlineResponse200113**](inline_response_200_113.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteWlan**
> DeleteSiteWlan(ctx, siteId, wlanId)
deleteSiteWlan

Delete Site WLAN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **wlanId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteWlanPortalImage**
> DeleteSiteWlanPortalImage(ctx, siteId, wlanId)
deleteSiteWlanPortalImage

Delete Site WLAN Portal Image

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **wlanId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteWlan**
> InlineResponse200113 GetSiteWlan(ctx, siteId, wlanId)
getSiteWlan

Get Site WLAN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **wlanId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200113**](inline_response_200_113.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteWlanDerived**
> []Wlan ListSiteWlanDerived(ctx, siteId, optional)
listSiteWlanDerived

Get Wlans Derived

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesWlansApiListSiteWlanDerivedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWlansApiListSiteWlanDerivedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resolve** | **optional.Bool**| whether to resolve SITE_VARS | [default to false]
 **wlanId** | **optional.String**| filter by WLAN ID | 

### Return type

[**[]Wlan**](wlan.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteWlans**
> []Wlan ListSiteWlans(ctx, siteId, optional)
listSiteWlans

Get List of Site WLANs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesWlansApiListSiteWlansOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWlansApiListSiteWlansOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Wlan**](wlan.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSiteWlan**
> InlineResponse200113 UpdateSiteWlan(ctx, siteId, wlanId, optional)
updateSiteWlan

Update Site WLAN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **wlanId** | [**string**](.md)|  | 
 **optional** | ***SitesWlansApiUpdateSiteWlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWlansApiUpdateSiteWlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of WlansWlanIdBody1**](WlansWlanIdBody1.md)| Request Body | 

### Return type

[**InlineResponse200113**](inline_response_200_113.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSiteWlanPortalTemplate**
> InlineResponse200114 UpdateSiteWlanPortalTemplate(ctx, siteId, wlanId, optional)
updateSiteWlanPortalTemplate

Update a Portal Template  #### Sponsor Email Template Sponsor Email Template supports following template variables:  | **Name** | **Description** | | --- | --- | | approve_url | Renders URL to approve the request; optionally &minutes=N query param can be appended to change the Authorization period of the guest, where N is a valid integer denoting number of minutes a guest remains authorized | | deny_url | Renders URL to reject the request | | guest_email | Renders Email ID of the guest | | guest_name | Renders Name of the guest | | field1 | Renders value of the Custom Field 1 | | field2 | Renders value of the Custom Field 2 | | company | Renders value of the Company field | | sponsor_link_validity_duration | Renders validity time of the request (i.e. Approve/Deny URL) | | auth_expire_minutes | Renders Wlan-level configured Guest Authorization Expiration time period (in minutes), If not configured then default (1 day in minutes) |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **wlanId** | [**string**](.md)|  | 
 **optional** | ***SitesWlansApiUpdateSiteWlanPortalTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWlansApiUpdateSiteWlanPortalTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of WlanIdPortalTemplateBody1**](WlanIdPortalTemplateBody1.md)| Request Body | 

### Return type

[**InlineResponse200114**](inline_response_200_114.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadSiteWlanPortalImage**
> UploadSiteWlanPortalImage(ctx, siteId, wlanId, optional)
uploadSiteWlanPortalImage

Wlan Portal Image Upload

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **wlanId** | [**string**](.md)|  | 
 **optional** | ***SitesWlansApiUploadSiteWlanPortalImageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesWlansApiUploadSiteWlanPortalImageOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | **optional.Interface of *os.File****optional.**|  | 
 **json** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

