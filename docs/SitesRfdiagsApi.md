# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSiteRfdiagRecording**](SitesRfdiagsApi.md#DeleteSiteRfdiagRecording) | **Delete** /api/v1/sites/{site_id}/rfdiags/{rfdiag_id} | deleteSiteRfdiagRecording
[**DownloadSiteRfdiagRecording**](SitesRfdiagsApi.md#DownloadSiteRfdiagRecording) | **Get** /api/v1/sites/{site_id}/rfdiags/{rfdiag_id}/download | downloadSiteRfdiagRecording
[**GetSiteRfdiagRecording**](SitesRfdiagsApi.md#GetSiteRfdiagRecording) | **Get** /api/v1/sites/{site_id}/rfdiags/{rfdiag_id} | getSiteRfdiagRecording
[**GetSiteSiteRfdiagRecording**](SitesRfdiagsApi.md#GetSiteSiteRfdiagRecording) | **Get** /api/v1/sites/{site_id}/rfdiags | getSiteSiteRfdiagRecording
[**StartSiteRecording**](SitesRfdiagsApi.md#StartSiteRecording) | **Post** /api/v1/sites/{site_id}/rfdiags | startSiteRecording
[**StopSiteRfdiagRecording**](SitesRfdiagsApi.md#StopSiteRfdiagRecording) | **Post** /api/v1/sites/{site_id}/rfdiags/{rfdiag_id}/stop | stopSiteRfdiagRecording
[**UpdateSiteRfdiagRecording**](SitesRfdiagsApi.md#UpdateSiteRfdiagRecording) | **Put** /api/v1/sites/{site_id}/rfdiags/{rfdiag_id} | updateSiteRfdiagRecording

# **DeleteSiteRfdiagRecording**
> DeleteSiteRfdiagRecording(ctx, siteId, rfdiagId)
deleteSiteRfdiagRecording

Delete Recording

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **rfdiagId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadSiteRfdiagRecording**
> *os.File DownloadSiteRfdiagRecording(ctx, siteId, rfdiagId)
downloadSiteRfdiagRecording

Download Recording Download raw_events blob

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **rfdiagId** | [**string**](.md)|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteRfdiagRecording**
> []RfDiagInfoItem GetSiteRfdiagRecording(ctx, siteId, rfdiagId)
getSiteRfdiagRecording

Get RF Diage Recording Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **rfdiagId** | [**string**](.md)|  | 

### Return type

[**[]RfDiagInfoItem**](rf_diag_info_item.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSiteRfdiagRecording**
> [][]RfDiagInfoItem GetSiteSiteRfdiagRecording(ctx, siteId, optional)
getSiteSiteRfdiagRecording

List RF Glass Recording

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesRfdiagsApiGetSiteSiteRfdiagRecordingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesRfdiagsApiGetSiteSiteRfdiagRecordingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[][]RfDiagInfoItem**](array.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartSiteRecording**
> []RfDiagInfoItem StartSiteRecording(ctx, siteId, optional)
startSiteRecording

Start RF Glass Recording

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesRfdiagsApiStartSiteRecordingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesRfdiagsApiStartSiteRecordingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SiteIdRfdiagsBody**](SiteIdRfdiagsBody.md)| Request Body | 

### Return type

[**[]RfDiagInfoItem**](rf_diag_info_item.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopSiteRfdiagRecording**
> StopSiteRfdiagRecording(ctx, siteId, rfdiagId)
stopSiteRfdiagRecording

If the recording session is active for the given rfdiag_id, it will finish the recording. duration and end_time will be updated to reflect the correct values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **rfdiagId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSiteRfdiagRecording**
> []RfDiagInfoItem UpdateSiteRfdiagRecording(ctx, siteId, rfdiagId, optional)
updateSiteRfdiagRecording

Update Recording

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **rfdiagId** | [**string**](.md)|  | 
 **optional** | ***SitesRfdiagsApiUpdateSiteRfdiagRecordingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesRfdiagsApiUpdateSiteRfdiagRecordingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of RfdiagsRfdiagIdBody**](RfdiagsRfdiagIdBody.md)| Request Body | 

### Return type

[**[]RfDiagInfoItem**](rf_diag_info_item.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

