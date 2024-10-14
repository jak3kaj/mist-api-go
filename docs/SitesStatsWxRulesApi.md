# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteWxRulesUsage**](SitesStatsWxRulesApi.md#GetSiteWxRulesUsage) | **Get** /api/v1/sites/{site_id}/stats/wxrules | getSiteWxRulesUsage

# **GetSiteWxRulesUsage**
> []StatsWxrule GetSiteWxRulesUsage(ctx, siteId)
getSiteWxRulesUsage

Get Wxlan Rule usage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

[**[]StatsWxrule**](stats_wxrule.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

