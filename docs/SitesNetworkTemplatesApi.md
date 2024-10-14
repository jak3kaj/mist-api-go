# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSiteNetworkTemplateDerived**](SitesNetworkTemplatesApi.md#ListSiteNetworkTemplateDerived) | **Get** /api/v1/sites/{site_id}/networktemplates/derived | listSiteNetworkTemplateDerived

# **ListSiteNetworkTemplateDerived**
> InlineResponse20081 ListSiteNetworkTemplateDerived(ctx, siteId, optional)
listSiteNetworkTemplateDerived

Get derived Network Templates for Site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesNetworkTemplatesApiListSiteNetworkTemplateDerivedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesNetworkTemplatesApiListSiteNetworkTemplateDerivedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resolve** | **optional.Bool**| whether resolve the site variables | 

### Return type

[**InlineResponse20081**](inline_response_200_81.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

