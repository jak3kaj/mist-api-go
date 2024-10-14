# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSiteSecIntelProfilesDerived**](SitesSecIntelProfilesApi.md#ListSiteSecIntelProfilesDerived) | **Get** /api/v1/sites/{site_id}/secintelprofiles/derived | listSiteSecIntelProfilesDerived

# **ListSiteSecIntelProfilesDerived**
> []SecintelProfile ListSiteSecIntelProfilesDerived(ctx, siteId, optional)
listSiteSecIntelProfilesDerived

Get derived Sky-ATP secintel profiles for Site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesSecIntelProfilesApiListSiteSecIntelProfilesDerivedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSecIntelProfilesApiListSiteSecIntelProfilesDerivedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resolve** | **optional.Bool**| whether resolve the site variables | 

### Return type

[**[]SecintelProfile**](secintel_profile.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

