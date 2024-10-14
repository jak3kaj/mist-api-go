# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteGuestAuthorizations**](SitesGuestsApi.md#CountSiteGuestAuthorizations) | **Get** /api/v1/sites/{site_id}/guests/count | countSiteGuestAuthorizations
[**DeleteSiteGuestAuthorization**](SitesGuestsApi.md#DeleteSiteGuestAuthorization) | **Delete** /api/v1/sites/{site_id}/guests/{guest_mac} | deleteSiteGuestAuthorization
[**GetSiteGuestAuthorization**](SitesGuestsApi.md#GetSiteGuestAuthorization) | **Get** /api/v1/sites/{site_id}/guests/{guest_mac} | getSiteGuestAuthorization
[**ListSiteAllGuestAuthorizations**](SitesGuestsApi.md#ListSiteAllGuestAuthorizations) | **Get** /api/v1/sites/{site_id}/guests | listSiteAllGuestAuthorizations
[**ListSiteAllGuestAuthorizationsDerived**](SitesGuestsApi.md#ListSiteAllGuestAuthorizationsDerived) | **Get** /api/v1/sites/{site_id}/guests/derived | listSiteAllGuestAuthorizationsDerived
[**SearchSiteGuestAuthorization**](SitesGuestsApi.md#SearchSiteGuestAuthorization) | **Get** /api/v1/sites/{site_id}/guests/search | searchSiteGuestAuthorization
[**UpdateSiteGuestAuthorization**](SitesGuestsApi.md#UpdateSiteGuestAuthorization) | **Put** /api/v1/sites/{site_id}/guests/{guest_mac} | updateSiteGuestAuthorization

# **CountSiteGuestAuthorizations**
> InlineResponse20016 CountSiteGuestAuthorizations(ctx, siteId, optional)
countSiteGuestAuthorizations

Count Authorized Guest

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesGuestsApiCountSiteGuestAuthorizationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesGuestsApiCountSiteGuestAuthorizationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of Distinct7**](.md)|  | 
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

# **DeleteSiteGuestAuthorization**
> DeleteSiteGuestAuthorization(ctx, siteId, guestMac)
deleteSiteGuestAuthorization

Delete Guest Authorization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **guestMac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteGuestAuthorization**
> InlineResponse20060 GetSiteGuestAuthorization(ctx, siteId, guestMac)
getSiteGuestAuthorization

Get Guest Authorization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **guestMac** | **string**|  | 

### Return type

[**InlineResponse20060**](inline_response_200_60.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteAllGuestAuthorizations**
> []Guest ListSiteAllGuestAuthorizations(ctx, siteId, optional)
listSiteAllGuestAuthorizations

Get List of Site Guest Authorizations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesGuestsApiListSiteAllGuestAuthorizationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesGuestsApiListSiteAllGuestAuthorizationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wlanId** | **optional.String**| UUID of single or multiple (Comma separated) WLAN under Site &#x60;site_id&#x60; (to filter by WLAN) | 

### Return type

[**[]Guest**](guest.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteAllGuestAuthorizationsDerived**
> []Guest ListSiteAllGuestAuthorizationsDerived(ctx, siteId, optional)
listSiteAllGuestAuthorizationsDerived

Get List of Site Guest Authorizations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesGuestsApiListSiteAllGuestAuthorizationsDerivedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesGuestsApiListSiteAllGuestAuthorizationsDerivedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wlanId** | **optional.String**| UUID of single or multiple (Comma separated) WLAN under Site &#x60;site_id&#x60; (to filter by WLAN) | 
 **crossSite** | **optional.Bool**| whether to get org level guests, default is false i.e get site level guests | [default to false]

### Return type

[**[]Guest**](guest.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteGuestAuthorization**
> InlineResponse20059 SearchSiteGuestAuthorization(ctx, siteId, optional)
searchSiteGuestAuthorization

Search Authorized Guest

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesGuestsApiSearchSiteGuestAuthorizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesGuestsApiSearchSiteGuestAuthorizationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wlanId** | **optional.String**|  | 
 **authMethod** | **optional.String**|  | 
 **ssid** | **optional.String**|  | 
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

# **UpdateSiteGuestAuthorization**
> InlineResponse20060 UpdateSiteGuestAuthorization(ctx, siteId, guestMac, optional)
updateSiteGuestAuthorization

Update Guest Authorization

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **guestMac** | **string**|  | 
 **optional** | ***SitesGuestsApiUpdateSiteGuestAuthorizationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesGuestsApiUpdateSiteGuestAuthorizationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of GuestsGuestMacBody1**](GuestsGuestMacBody1.md)| Request Body | 

### Return type

[**InlineResponse20060**](inline_response_200_60.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

