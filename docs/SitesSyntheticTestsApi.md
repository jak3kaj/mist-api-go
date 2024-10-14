# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteDeviceSyntheticTest**](SitesSyntheticTestsApi.md#GetSiteDeviceSyntheticTest) | **Get** /api/v1/sites/{site_id}/devices/{device_id}/synthetic_test | getSiteDeviceSyntheticTest
[**GetSiteSyntheticTestStatus**](SitesSyntheticTestsApi.md#GetSiteSyntheticTestStatus) | **Get** /api/v1/sites/{site_id}/synthetic_test | getSiteSyntheticTestStatus
[**SearchSiteSyntheticTest**](SitesSyntheticTestsApi.md#SearchSiteSyntheticTest) | **Get** /api/v1/sites/{site_id}/synthetic_test/search | searchSiteSyntheticTest
[**StartSiteSwitchRadiusSyntheticTest**](SitesSyntheticTestsApi.md#StartSiteSwitchRadiusSyntheticTest) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/check_radius_server | startSiteSwitchRadiusSyntheticTest
[**TriggerSiteDeviceSyntheticTest**](SitesSyntheticTestsApi.md#TriggerSiteDeviceSyntheticTest) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/synthetic_test | triggerSiteDeviceSyntheticTest
[**TriggerSiteSyntheticTest**](SitesSyntheticTestsApi.md#TriggerSiteSyntheticTest) | **Post** /api/v1/sites/{site_id}/synthetic_test | triggerSiteSyntheticTest

# **GetSiteDeviceSyntheticTest**
> InlineResponse200160 GetSiteDeviceSyntheticTest(ctx, siteId, deviceId, optional)
getSiteDeviceSyntheticTest

Get Device Synthetic Test

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***SitesSyntheticTestsApiGetSiteDeviceSyntheticTestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSyntheticTestsApiGetSiteDeviceSyntheticTestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | [**optional.Interface of Type3**](.md)| synthetic test type | 
 **tenant** | **optional.String**| tenant network in which &#x60;lan_connectivity&#x60; test is run | 
 **node** | [**optional.Interface of Node**](.md)| tenant network in which &#x60;lan_connectivity&#x60; test is run | 

### Return type

[**InlineResponse200160**](inline_response_200_160.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSyntheticTestStatus**
> InlineResponse200160 GetSiteSyntheticTestStatus(ctx, siteId)
getSiteSyntheticTestStatus

Get Synthetic Testing Status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 

### Return type

[**InlineResponse200160**](inline_response_200_160.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSiteSyntheticTest**
> InlineResponse200162 SearchSiteSyntheticTest(ctx, siteId, optional)
searchSiteSyntheticTest

Search Site Synthetic Testing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesSyntheticTestsApiSearchSiteSyntheticTestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSyntheticTestsApiSearchSiteSyntheticTestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **optional.String**| Device MAC Address | 
 **portId** | **optional.String**| port_id used to run the test (for SSR only) | 
 **vlanId** | **optional.String**| VLAN ID | 
 **by** | **optional.String**| entity who triggers the test | 
 **reason** | **optional.String**| test failure reason | 
 **type_** | [**optional.Interface of Type4**](.md)| synthetic test type | 
 **tenant** | **optional.String**| tenant network in which lan_connectivity test was run | 

### Return type

[**InlineResponse200162**](inline_response_200_162.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartSiteSwitchRadiusSyntheticTest**
> WebsocketSession StartSiteSwitchRadiusSyntheticTest(ctx, siteId, deviceId, optional)
startSiteSwitchRadiusSyntheticTest

Ping test from the AP to confirm ‘reachability’ of the Radius server. Utilize Juniper EX switch(to which an AP is connected to) radius test capabilities to get details on the Radius Server ‘availability’ .

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***SitesSyntheticTestsApiStartSiteSwitchRadiusSyntheticTestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSyntheticTestsApiStartSiteSwitchRadiusSyntheticTestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of DeviceIdCheckRadiusServerBody**](DeviceIdCheckRadiusServerBody.md)|  | 

### Return type

[**WebsocketSession**](websocket_session.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TriggerSiteDeviceSyntheticTest**
> TriggerSiteDeviceSyntheticTest(ctx, siteId, deviceId, optional)
triggerSiteDeviceSyntheticTest

Trigger Device Synthetic Test

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **deviceId** | [**string**](.md)|  | 
 **optional** | ***SitesSyntheticTestsApiTriggerSiteDeviceSyntheticTestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSyntheticTestsApiTriggerSiteDeviceSyntheticTestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SynthetictestDevice**](SynthetictestDevice.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TriggerSiteSyntheticTest**
> InlineResponse200161 TriggerSiteSyntheticTest(ctx, siteId, optional)
triggerSiteSyntheticTest

Trigger Synthetic Testing

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***SitesSyntheticTestsApiTriggerSiteSyntheticTestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSyntheticTestsApiTriggerSiteSyntheticTestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of SiteIdSyntheticTestBody**](SiteIdSyntheticTestBody.md)|  | 

### Return type

[**InlineResponse200161**](inline_response_200_161.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

