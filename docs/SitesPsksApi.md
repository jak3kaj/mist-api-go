# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSitePsk**](SitesPsksApi.md#CreateSitePsk) | **Post** /api/v1/sites/{site_id}/psks | createSitePsk
[**DeleteSitePsk**](SitesPsksApi.md#DeleteSitePsk) | **Delete** /api/v1/sites/{site_id}/psks/{psk_id} | deleteSitePsk
[**GetSitePsk**](SitesPsksApi.md#GetSitePsk) | **Get** /api/v1/sites/{site_id}/psks/{psk_id} | getSitePsk
[**ImportSitePsks**](SitesPsksApi.md#ImportSitePsks) | **Post** /api/v1/sites/{site_id}/psks/import | importSitePsks
[**ListSitePsks**](SitesPsksApi.md#ListSitePsks) | **Get** /api/v1/sites/{site_id}/psks | listSitePsks
[**UpdateSiteMultiplePsks**](SitesPsksApi.md#UpdateSiteMultiplePsks) | **Put** /api/v1/sites/{site_id}/psks | updateSiteMultiplePsks
[**UpdateSitePsk**](SitesPsksApi.md#UpdateSitePsk) | **Put** /api/v1/sites/{site_id}/psks/{psk_id} | updateSitePsk

# **CreateSitePsk**
> InlineResponse20085 CreateSitePsk(ctx, siteId, optional)
createSitePsk

Create Site PSK   When `usage`==`macs`, corresponding \"macs\" field will hold a list consisting of client mac addresses ([\"xx:xx:xx:xx:xx\",...]) or mac patterns([\"xx:xx:*\",\"xx*\",...]) or both ([\"xx:xx:xx:xx:xx:xx\", \"xx:*\", ...]). This list is capped at 5000

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesPsksApiCreateSitePskOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesPsksApiCreateSitePskOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SiteIdPsksBody**](SiteIdPsksBody.md)| Request Body | 

### Return type

[**InlineResponse20085**](inline_response_200_85.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSitePsk**
> DeleteSitePsk(ctx, siteId, pskId)
deleteSitePsk

Delete Site PSK

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **pskId** | [**string**](.md)| PSK ID | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSitePsk**
> InlineResponse20085 GetSitePsk(ctx, siteId, pskId)
getSitePsk

Get Site PSK Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **pskId** | [**string**](.md)| PSK ID | 

### Return type

[**InlineResponse20085**](inline_response_200_85.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportSitePsks**
> []Psk ImportSitePsks(ctx, siteId, optional)
importSitePsks

Import PSK from CSV file or JSON  ## CSV File Format ```csv PSK Import CSV File Format: name,ssid,passphrase,usage,vlan_id,mac Common,warehouse,foryoureyesonly,single,35,a31425f31278 Justin,reception,visible,multi,1002 ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesPsksApiImportSitePsksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesPsksApiImportSitePsksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **optional.Interface of *os.File****optional.**|  | 

### Return type

[**[]Psk**](psk.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSitePsks**
> []Psk ListSitePsks(ctx, siteId, optional)
listSitePsks

Get List of Site PSKs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesPsksApiListSitePsksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesPsksApiListSitePsksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ssid** | **optional.String**|  | 
 **role** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | [default to 100]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**[]Psk**](psk.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSiteMultiplePsks**
> []Psk UpdateSiteMultiplePsks(ctx, siteId, optional)
updateSiteMultiplePsks

Update multiple PSKs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesPsksApiUpdateSiteMultiplePsksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesPsksApiUpdateSiteMultiplePsksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of []Psk**](psk.md)|  | 

### Return type

[**[]Psk**](psk.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSitePsk**
> InlineResponse20085 UpdateSitePsk(ctx, siteId, pskId, optional)
updateSitePsk

Update Site PSK

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **pskId** | [**string**](.md)| PSK ID | 
 **optional** | ***SitesPsksApiUpdateSitePskOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesPsksApiUpdateSitePskOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PsksPskIdBody1**](PsksPskIdBody1.md)| Request Body | 

### Return type

[**InlineResponse20085**](inline_response_200_85.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

