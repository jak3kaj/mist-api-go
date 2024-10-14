# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteApps**](SitesStatsAppsApi.md#CountSiteApps) | **Get** /api/v1/sites/{site_id}/stats/apps/count | countSiteApps

# **CountSiteApps**
> InlineResponse20016 CountSiteApps(ctx, siteId, optional)
countSiteApps

Count by Distinct Attributes of Applications

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesStatsAppsApiCountSiteAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesStatsAppsApiCountSiteAppsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of Distinct10**](.md)| Default for wireless devices is &#x60;ap&#x60;. Default for wired devices is &#x60;device_mac&#x60; | 
 **deviceMac** | **optional.String**| MAC of the device | 
 **app** | **optional.String**| Application name | 
 **wired** | **optional.String**| If a device is wired or wireless. Default is False. | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

