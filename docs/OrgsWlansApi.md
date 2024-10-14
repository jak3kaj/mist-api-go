# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgWlan**](OrgsWlansApi.md#CreateOrgWlan) | **Post** /api/v1/orgs/{org_id}/wlans | createOrgWlan
[**DeleteOrgWlan**](OrgsWlansApi.md#DeleteOrgWlan) | **Delete** /api/v1/orgs/{org_id}/wlans/{wlan_id} | deleteOrgWlan
[**DeleteOrgWlanPortalImage**](OrgsWlansApi.md#DeleteOrgWlanPortalImage) | **Delete** /api/v1/orgs/{org_id}/wlans/{wlan_id}/portal_image | deleteOrgWlanPortalImage
[**GetOrgWLAN**](OrgsWlansApi.md#GetOrgWLAN) | **Get** /api/v1/orgs/{org_id}/wlans/{wlan_id} | getOrgWLAN
[**ListOrgWlans**](OrgsWlansApi.md#ListOrgWlans) | **Get** /api/v1/orgs/{org_id}/wlans | listOrgWlans
[**UpdateOrgWlan**](OrgsWlansApi.md#UpdateOrgWlan) | **Put** /api/v1/orgs/{org_id}/wlans/{wlan_id} | updateOrgWlan
[**UpdateOrgWlanPortalTemplate**](OrgsWlansApi.md#UpdateOrgWlanPortalTemplate) | **Put** /api/v1/orgs/{org_id}/wlans/{wlan_id}/portal_template | updateOrgWlanPortalTemplate
[**UploadOrgWlanPortalImage**](OrgsWlansApi.md#UploadOrgWlanPortalImage) | **Post** /api/v1/orgs/{org_id}/wlans/{wlan_id}/portal_image | uploadOrgWlanPortalImage

# **CreateOrgWlan**
> InlineResponse200113 CreateOrgWlan(ctx, orgId, optional)
createOrgWlan

Create Org Wlan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsWlansApiCreateOrgWlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWlansApiCreateOrgWlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdWlansBody**](OrgIdWlansBody.md)| Request Body | 

### Return type

[**InlineResponse200113**](inline_response_200_113.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgWlan**
> DeleteOrgWlan(ctx, orgId, wlanId)
deleteOrgWlan

Delete Org WLAN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **wlanId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgWlanPortalImage**
> DeleteOrgWlanPortalImage(ctx, orgId, wlanId)
deleteOrgWlanPortalImage

Delete Org WLAN Portal Image

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **wlanId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgWLAN**
> InlineResponse200113 GetOrgWLAN(ctx, orgId, wlanId)
getOrgWLAN

Get Org Wlan Detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **wlanId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200113**](inline_response_200_113.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgWlans**
> []Wlan ListOrgWlans(ctx, orgId, optional)
listOrgWlans

Get List of Org Wlans

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsWlansApiListOrgWlansOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWlansApiListOrgWlansOpts struct
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

# **UpdateOrgWlan**
> InlineResponse200113 UpdateOrgWlan(ctx, orgId, wlanId, optional)
updateOrgWlan

Update Org Wlan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **wlanId** | [**string**](.md)|  | 
 **optional** | ***OrgsWlansApiUpdateOrgWlanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWlansApiUpdateOrgWlanOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of WlansWlanIdBody**](WlansWlanIdBody.md)| Request Body | 

### Return type

[**InlineResponse200113**](inline_response_200_113.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgWlanPortalTemplate**
> InlineResponse200114 UpdateOrgWlanPortalTemplate(ctx, orgId, wlanId, optional)
updateOrgWlanPortalTemplate

Update a Portal Template  #### Sponsor Email Template Sponsor Email Template supports following template variables:  | **Name** | **Description** | | --- | --- | | approve_url | Renders URL to approve the request; optionally &minutes=N query param can be appended to change the Authorization period of the guest, where N is a valid integer denoting number of minutes a guest remains authorized | | deny_url | Renders URL to reject the request | | guest_email | Renders Email ID of the guest | | guest_name | Renders Name of the guest | | field1 | Renders value of the Custom Field 1 | | field2 | Renders value of the Custom Field 2 | | company | Renders value of the Company field | | sponsor_link_validity_duration | Renders validity time of the request (i.e. Approve/Deny URL) | | auth_expire_minutes | Renders Wlan-level configured Guest Authorization Expiration time period (in minutes), If not configured then default (1 day in minutes) |

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **wlanId** | [**string**](.md)|  | 
 **optional** | ***OrgsWlansApiUpdateOrgWlanPortalTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWlansApiUpdateOrgWlanPortalTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of WlanIdPortalTemplateBody**](WlanIdPortalTemplateBody.md)| Request Body | 

### Return type

[**InlineResponse200114**](inline_response_200_114.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadOrgWlanPortalImage**
> UploadOrgWlanPortalImage(ctx, orgId, wlanId, optional)
uploadOrgWlanPortalImage

Upload Org WLAN Portal Image

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **wlanId** | [**string**](.md)|  | 
 **optional** | ***OrgsWlansApiUploadOrgWlanPortalImageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsWlansApiUploadOrgWlanPortalImageOpts struct
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

