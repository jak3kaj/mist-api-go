# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgSites**](OrgsSitesApi.md#CountOrgSites) | **Get** /api/v1/orgs/{org_id}/sites/count | countOrgSites
[**CreateOrgSite**](OrgsSitesApi.md#CreateOrgSite) | **Post** /api/v1/orgs/{org_id}/sites | createOrgSite
[**ListOrgSites**](OrgsSitesApi.md#ListOrgSites) | **Get** /api/v1/orgs/{org_id}/sites | listOrgSites
[**SearchOrgSites**](OrgsSitesApi.md#SearchOrgSites) | **Get** /api/v1/orgs/{org_id}/sites/search | searchOrgSites

# **CountOrgSites**
> InlineResponse20016 CountOrgSites(ctx, orgId, optional)
countOrgSites

Count Sites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSitesApiCountOrgSitesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSitesApiCountOrgSitesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**optional.Interface of OrgSitesCountDistinct**](.md)|  | 
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

# **CreateOrgSite**
> InlineResponse20098 CreateOrgSite(ctx, orgId, optional)
createOrgSite

Create Org Site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSitesApiCreateOrgSiteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSitesApiCreateOrgSiteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OrgIdSitesBody**](OrgIdSitesBody.md)| Request Body | 

### Return type

[**InlineResponse20098**](inline_response_200_98.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgSites**
> []Site ListOrgSites(ctx, orgId, optional)
listOrgSites

Get List of Org Sites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSitesApiListOrgSitesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSitesApiListOrgSitesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Site**](site.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrgSites**
> InlineResponse20099 SearchOrgSites(ctx, orgId, optional)
searchOrgSites

Search Sites

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
 **optional** | ***OrgsSitesApiSearchOrgSitesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrgsSitesApiSearchOrgSitesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **analyticEnabled** | **optional.Bool**| if Advanced Analytic feature is enabled | 
 **appWaking** | **optional.Bool**| if App Waking feature is enabled | 
 **assetEnabled** | **optional.Bool**| if Asset Tracking is enabled | 
 **autoUpgradeEnabled** | **optional.Bool**| if Auto Upgrade feature is enabled | 
 **autoUpgradeVersion** | **optional.String**| if Auto Upgrade feature is enabled | 
 **countryCode** | **optional.String**| site country code | 
 **honeypotEnabled** | **optional.Bool**| if Honeypot detection is enabled | 
 **id** | **optional.String**| site id | 
 **locateUnconnected** | **optional.Bool**| if unconnected client are located | 
 **meshEnabled** | **optional.Bool**| if Mesh feature is enabled | 
 **name** | **optional.String**| site name | 
 **rogueEnabled** | **optional.Bool**| if Rogue detection is enabled | 
 **remoteSyslogEnabled** | **optional.Bool**| if Remote Syslog is enabled | 
 **rtsaEnabled** | **optional.Bool**| if managed mobility feature is enabled | 
 **vnaEnabled** | **optional.Bool**| if Virtual Network Assistant is enabled | 
 **wifiEnabled** | **optional.Bool**| if WIFI feature is enabled | 
 **limit** | **optional.Int32**|  | [default to 100]
 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse20099**](inline_response_200_99.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

