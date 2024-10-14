# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteWatchedStations**](SitesSettingApi.md#CreateSiteWatchedStations) | **Post** /api/v1/sites/{site_id}/setting/watched_station | createSiteWatchedStations
[**CreateSiteWirelessClientsAllowlist**](SitesSettingApi.md#CreateSiteWirelessClientsAllowlist) | **Post** /api/v1/sites/{site_id}/setting/whitelist | createSiteWirelessClientsAllowlist
[**CreateSiteWirelessClientsBlocklist**](SitesSettingApi.md#CreateSiteWirelessClientsBlocklist) | **Post** /api/v1/sites/{site_id}/setting/blacklist | createSiteWirelessClientsBlocklist
[**DeleteSiteWatchedStations**](SitesSettingApi.md#DeleteSiteWatchedStations) | **Delete** /api/v1/sites/{site_id}/setting/watched_station | deleteSiteWatchedStations
[**DeleteSiteWirelessClientsAllowlist**](SitesSettingApi.md#DeleteSiteWirelessClientsAllowlist) | **Delete** /api/v1/sites/{site_id}/setting/whitelist | deleteSiteWirelessClientsAllowlist
[**DeleteSiteWirelessClientsBlocklist**](SitesSettingApi.md#DeleteSiteWirelessClientsBlocklist) | **Delete** /api/v1/sites/{site_id}/setting/blacklist | deleteSiteWirelessClientsBlocklist
[**GetSiteSetting**](SitesSettingApi.md#GetSiteSetting) | **Get** /api/v1/sites/{site_id}/setting | getSiteSetting
[**UpdateSiteSettings**](SitesSettingApi.md#UpdateSiteSettings) | **Put** /api/v1/sites/{site_id}/setting | updateSiteSettings

# **CreateSiteWatchedStations**
> InlineResponse20095 CreateSiteWatchedStations(ctx, siteId, optional)
createSiteWatchedStations

This endpoint is to provide list of client macs for annotation as watched station.  Retrieve the current clients list from `watched_station_url` under Site:Setting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesSettingApiCreateSiteWatchedStationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSettingApiCreateSiteWatchedStationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SettingWatchedStationBody**](SettingWatchedStationBody.md)| Request Body | 

### Return type

[**InlineResponse20095**](inline_response_200_95.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSiteWirelessClientsAllowlist**
> InlineResponse20095 CreateSiteWirelessClientsAllowlist(ctx, siteId, optional)
createSiteWirelessClientsAllowlist

This endpoint is to provide list of client macs for annotation as whitelist.  Retrieve the current clients list from `whitelist_url` under Site:Setting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesSettingApiCreateSiteWirelessClientsAllowlistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSettingApiCreateSiteWirelessClientsAllowlistOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SettingWhitelistBody**](SettingWhitelistBody.md)| Request Body | 

### Return type

[**InlineResponse20095**](inline_response_200_95.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSiteWirelessClientsBlocklist**
> InlineResponse20095 CreateSiteWirelessClientsBlocklist(ctx, siteId, optional)
createSiteWirelessClientsBlocklist

This endpoint is to provide list of client macs for annotation blacklist.  Retrieve the current clients list `blacklist_url` under Site:Setting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesSettingApiCreateSiteWirelessClientsBlocklistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSettingApiCreateSiteWirelessClientsBlocklistOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SettingBlacklistBody1**](SettingBlacklistBody1.md)| Request Body | 

### Return type

[**InlineResponse20095**](inline_response_200_95.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteWatchedStations**
> DeleteSiteWatchedStations(ctx, siteId)
deleteSiteWatchedStations

Delete Site Watched Station Clients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteWirelessClientsAllowlist**
> DeleteSiteWirelessClientsAllowlist(ctx, siteId)
deleteSiteWirelessClientsAllowlist

Delete Site Whitelist Station Clients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteWirelessClientsBlocklist**
> DeleteSiteWirelessClientsBlocklist(ctx, siteId)
deleteSiteWirelessClientsBlocklist

Delete Site Blacklist Station Clients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSetting**
> InlineResponse200144 GetSiteSetting(ctx, siteId)
getSiteSetting

Get Site Settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200144**](inline_response_200_144.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSiteSettings**
> InlineResponse200144 UpdateSiteSettings(ctx, siteId, optional)
updateSiteSettings

Update Site Settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesSettingApiUpdateSiteSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSettingApiUpdateSiteSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SiteIdSettingBody**](SiteIdSettingBody.md)| Request Body | 

### Return type

[**InlineResponse200144**](inline_response_200_144.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

