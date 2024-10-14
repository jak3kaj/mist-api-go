# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSiteNetworksDerived**](SitesNetworksApi.md#ListSiteNetworksDerived) | **Get** /api/v1/sites/{site_id}/networks/derived | listSiteNetworksDerived

# **ListSiteNetworksDerived**
> []Network ListSiteNetworksDerived(ctx, siteId, optional)
listSiteNetworksDerived

Retrieves the list of Networks available for the Site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesNetworksApiListSiteNetworksDerivedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesNetworksApiListSiteNetworksDerivedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resolve** | **optional.Bool**| whether resolve the site variables | [default to false]

### Return type

[**[]Network**](network.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

