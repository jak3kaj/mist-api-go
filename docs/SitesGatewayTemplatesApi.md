# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSiteGatewayTemplateDerived**](SitesGatewayTemplatesApi.md#ListSiteGatewayTemplateDerived) | **Get** /api/v1/sites/{site_id}/gatewaytemplates/derived | listSiteGatewayTemplateDerived

# **ListSiteGatewayTemplateDerived**
> []GatewayTemplate ListSiteGatewayTemplateDerived(ctx, siteId, optional)
listSiteGatewayTemplateDerived

Get derived Gateway Templates for Site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesGatewayTemplatesApiListSiteGatewayTemplateDerivedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesGatewayTemplatesApiListSiteGatewayTemplateDerivedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resolve** | **optional.Bool**| whether resolve the site variables | 

### Return type

[**[]GatewayTemplate**](gateway_template.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

