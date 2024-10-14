# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteUiSettings**](SitesUISettingsApi.md#CreateSiteUiSettings) | **Post** /api/v1/sites/{site_id}/uisettings | createSiteUiSettings
[**DeleteSiteUiSetting**](SitesUISettingsApi.md#DeleteSiteUiSetting) | **Delete** /api/v1/sites/{site_id}/uisettings/{uisetting_id} | deleteSiteUiSetting
[**GetSiteUiSetting**](SitesUISettingsApi.md#GetSiteUiSetting) | **Get** /api/v1/sites/{site_id}/uisettings/{uisetting_id} | getSiteUiSetting
[**ListSiteUiSettingDerived**](SitesUISettingsApi.md#ListSiteUiSettingDerived) | **Get** /api/v1/sites/{site_id}/uisettings/derived | listSiteUiSettingDerived
[**ListSiteUiSettings**](SitesUISettingsApi.md#ListSiteUiSettings) | **Get** /api/v1/sites/{site_id}/uisettings | listSiteUiSettings
[**UpdateSiteUiSetting**](SitesUISettingsApi.md#UpdateSiteUiSetting) | **Post** /api/v1/sites/{site_id}/uisettings/{uisetting_id} | updateSiteUiSetting

# **CreateSiteUiSettings**
> InlineResponse200163 CreateSiteUiSettings(ctx, siteId, optional)
createSiteUiSettings

Site UI settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesUISettingsApiCreateSiteUiSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesUISettingsApiCreateSiteUiSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SiteIdUisettingsBody**](SiteIdUisettingsBody.md)| Request Body | 

### Return type

[**InlineResponse200163**](inline_response_200_163.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSiteUiSetting**
> DeleteSiteUiSetting(ctx, siteId, uisettingId)
deleteSiteUiSetting

Site UI settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **uisettingId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteUiSetting**
> []UiSettings GetSiteUiSetting(ctx, siteId, uisettingId)
getSiteUiSetting

Site UI settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **uisettingId** | [**string**](.md)|  | 

### Return type

[**[]UiSettings**](ui_settings.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteUiSettingDerived**
> InlineResponse200163 ListSiteUiSettingDerived(ctx, siteId)
listSiteUiSettingDerived

Get both site UI settings(for_site=true) and org UI settings (for_site=false)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200163**](inline_response_200_163.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteUiSettings**
> []UiSettings ListSiteUiSettings(ctx, siteId)
listSiteUiSettings

Site UI settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

[**[]UiSettings**](ui_settings.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSiteUiSetting**
> InlineResponse200163 UpdateSiteUiSetting(ctx, siteId, uisettingId, optional)
updateSiteUiSetting

Site UI settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **uisettingId** | [**string**](.md)|  | 
 **optional** | ***SitesUISettingsApiUpdateSiteUiSettingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesUISettingsApiUpdateSiteUiSettingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UisettingsUisettingIdBody**](UisettingsUisettingIdBody.md)| Request Body | 

### Return type

[**InlineResponse200163**](inline_response_200_163.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

