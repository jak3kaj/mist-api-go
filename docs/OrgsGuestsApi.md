# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgGuestAuthorizations**](OrgsGuestsApi.md#CountOrgGuestAuthorizations) | **Get** /api/v1/orgs/{org_id}/guests/count | countOrgGuestAuthorizations
[**DeleteOrgGuestAuthorization**](OrgsGuestsApi.md#DeleteOrgGuestAuthorization) | **Delete** /api/v1/orgs/{org_id}/guests/{guest_mac} | deleteOrgGuestAuthorization
[**GetOrgGuestAuthorization**](OrgsGuestsApi.md#GetOrgGuestAuthorization) | **Get** /api/v1/orgs/{org_id}/guests/{guest_mac} | getOrgGuestAuthorization
[**ListOrgGuestAuthorizations**](OrgsGuestsApi.md#ListOrgGuestAuthorizations) | **Get** /api/v1/orgs/{org_id}/guests | listOrgGuestAuthorizations
[**SearchOrgGuestAuthorization**](OrgsGuestsApi.md#SearchOrgGuestAuthorization) | **Get** /api/v1/orgs/{org_id}/guests/search | searchOrgGuestAuthorization
[**UpdateOrgGuestAuthorization**](OrgsGuestsApi.md#UpdateOrgGuestAuthorization) | **Put** /api/v1/orgs/{org_id}/guests/{guest_mac} | updateOrgGuestAuthorization

# **CountOrgGuestAuthorizations**
> InlineResponse20016 CountOrgGuestAuthorizations(ctx, orgId, optional)
countOrgGuestAuthorizations

Count Org Authorized Guest

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsGuestsApiCountOrgGuestAuthorizationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsGuestsApiCountOrgGuestAuthorizationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of OrgGuestsCountDistinct**](.md)|  | 
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgGuestAuthorization**
> DeleteOrgGuestAuthorization(ctx, orgId, guestMac)
deleteOrgGuestAuthorization

Delete Guest Authorization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **guestMac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgGuestAuthorization**
> InlineResponse20060 GetOrgGuestAuthorization(ctx, orgId, guestMac)
getOrgGuestAuthorization

Get Guest Authorization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **guestMac** | **string**|  | 

### Return type

[**InlineResponse20060**](inline_response_200_60.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgGuestAuthorizations**
> []Guest ListOrgGuestAuthorizations(ctx, orgId)
listOrgGuestAuthorizations

Get List of Org Guest Authorizations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 

### Return type

[**[]Guest**](guest.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrgGuestAuthorization**
> InlineResponse20059 SearchOrgGuestAuthorization(ctx, orgId, optional)
searchOrgGuestAuthorization

Search Authorized Guest

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsGuestsApiSearchOrgGuestAuthorizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsGuestsApiSearchOrgGuestAuthorizationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wlanId** | **optional.String**| WLAN ID | 
 **authMethod** | **optional.String**| Authentication Methdo | 
 **ssid** | **optional.String**| SSID | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse20059**](inline_response_200_59.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrgGuestAuthorization**
> InlineResponse20060 UpdateOrgGuestAuthorization(ctx, orgId, guestMac, optional)
updateOrgGuestAuthorization

Update Guest Authorization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **guestMac** | **string**|  | 
 **optional** | ***OrgsGuestsApiUpdateOrgGuestAuthorizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsGuestsApiUpdateOrgGuestAuthorizationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of GuestsGuestMacBody**](GuestsGuestMacBody.md)|  | 

### Return type

[**InlineResponse20060**](inline_response_200_60.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

