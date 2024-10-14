# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeauthSiteWirelessClientsConnectedToARogue**](UtilitiesWiFiApi.md#DeauthSiteWirelessClientsConnectedToARogue) | **Post** /api/v1/sites/{site_id}/rogues/{rogue_bssid}/deauth_clients | deauthSiteWirelessClientsConnectedToARogue
[**DisconnectSiteMultipleClients**](UtilitiesWiFiApi.md#DisconnectSiteMultipleClients) | **Post** /api/v1/sites/{site_id}/clients/disconnect | disconnectSiteMultipleClients
[**DisconnectSiteWirelessClient**](UtilitiesWiFiApi.md#DisconnectSiteWirelessClient) | **Post** /api/v1/sites/{site_id}/clients/{client_mac}/disconnect | disconnectSiteWirelessClient
[**OptimizeSiteRrm**](UtilitiesWiFiApi.md#OptimizeSiteRrm) | **Post** /api/v1/sites/{site_id}/rrm/optimize | optimizeSiteRrm
[**ReauthOrgDot1xWirelessClient**](UtilitiesWiFiApi.md#ReauthOrgDot1xWirelessClient) | **Post** /api/v1/orgs/{org_id}/clients/{client_mac}/coa | reauthOrgDot1xWirelessClient
[**ReauthSiteDot1xWirelessClient**](UtilitiesWiFiApi.md#ReauthSiteDot1xWirelessClient) | **Post** /api/v1/sites/{site_id}/clients/{client_mac}/coa | reauthSiteDot1xWirelessClient
[**ReprovisionSiteAllAps**](UtilitiesWiFiApi.md#ReprovisionSiteAllAps) | **Post** /api/v1/sites/{site_id}/devices/reprovision | reprovisionSiteAllAps
[**ResetSiteAllApsToUseRrm**](UtilitiesWiFiApi.md#ResetSiteAllApsToUseRrm) | **Post** /api/v1/sites/{site_id}/devices/reset_radio_config | resetSiteAllApsToUseRrm
[**TestSiteWlanTelstraSetup**](UtilitiesWiFiApi.md#TestSiteWlanTelstraSetup) | **Post** /api/v1/utils/test_telstra | testSiteWlanTelstraSetup
[**TestSiteWlanTwilioSetup**](UtilitiesWiFiApi.md#TestSiteWlanTwilioSetup) | **Post** /api/v1/utils/test_twilio | testSiteWlanTwilioSetup
[**UnauthorizeSiteMultipleClients**](UtilitiesWiFiApi.md#UnauthorizeSiteMultipleClients) | **Post** /api/v1/sites/{site_id}/clients/unauthorize | unauthorizeSiteMultipleClients
[**UnauthorizeSiteWirelessClient**](UtilitiesWiFiApi.md#UnauthorizeSiteWirelessClient) | **Post** /api/v1/sites/{site_id}/clients/{client_mac}/unauthorize | unauthorizeSiteWirelessClient
[**ZeroizeSiteFipsAllAps**](UtilitiesWiFiApi.md#ZeroizeSiteFipsAllAps) | **Post** /api/v1/sites/{site_id}/devices/zeroize | zeroizeSiteFipsAllAps

# **DeauthSiteWirelessClientsConnectedToARogue**
> DeauthSiteWirelessClientsConnectedToARogue(ctx, siteId, rogueBssid)
deauthSiteWirelessClientsConnectedToARogue

Send Deauth frame to clients connected to a Rogue AP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **rogueBssid** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisconnectSiteMultipleClients**
> DisconnectSiteMultipleClients(ctx, siteId, optional)
disconnectSiteMultipleClients

To unauthorize multiple clients

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesWiFiApiDisconnectSiteMultipleClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesWiFiApiDisconnectSiteMultipleClientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ClientsDisconnectBody**](ClientsDisconnectBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisconnectSiteWirelessClient**
> DisconnectSiteWirelessClient(ctx, siteId, clientMac)
disconnectSiteWirelessClient

This disconnect a client (and it’s likely to connect back)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **clientMac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OptimizeSiteRrm**
> OptimizeSiteRrm(ctx, siteId, optional)
optimizeSiteRrm

Optimize Site RRM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesWiFiApiOptimizeSiteRrmOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesWiFiApiOptimizeSiteRrmOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RrmOptimizeBody**](RrmOptimizeBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReauthOrgDot1xWirelessClient**
> ReauthOrgDot1xWirelessClient(ctx, orgId, clientMac)
reauthOrgDot1xWirelessClient

Trigger a CoA (change of authorization) against a client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | [**string**](.md)|  | 
  **clientMac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReauthSiteDot1xWirelessClient**
> ReauthSiteDot1xWirelessClient(ctx, siteId, clientMac)
reauthSiteDot1xWirelessClient

Trigger a CoA (change of authorization) against a Wireless client

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **clientMac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReprovisionSiteAllAps**
> ReprovisionSiteAllAps(ctx, siteId)
reprovisionSiteAllAps

To force all APs to reprovision itself again.

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

# **ResetSiteAllApsToUseRrm**
> ResetSiteAllApsToUseRrm(ctx, siteId, optional)
resetSiteAllApsToUseRrm

Reset all APs in the Site to use RRM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesWiFiApiResetSiteAllApsToUseRrmOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesWiFiApiResetSiteAllApsToUseRrmOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DevicesResetRadioConfigBody**](DevicesResetRadioConfigBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestSiteWlanTelstraSetup**
> TestSiteWlanTelstraSetup(ctx, optional)
testSiteWlanTelstraSetup

Allows validation of Telstra sms gateway credentials.  In case of success, a text message confirming successful setup should be received. In case of error, telstra error message are returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UtilitiesWiFiApiTestSiteWlanTelstraSetupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesWiFiApiTestSiteWlanTelstraSetupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UtilsTestTelstraBody**](UtilsTestTelstraBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestSiteWlanTwilioSetup**
> TestSiteWlanTwilioSetup(ctx, optional)
testSiteWlanTwilioSetup

Allows validation of twilio setup In case of success, a text message confirming successful setup should be received. In case of error, twilio error code and message are returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UtilitiesWiFiApiTestSiteWlanTwilioSetupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesWiFiApiTestSiteWlanTwilioSetupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UtilsTestTwilioBody**](UtilsTestTwilioBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnauthorizeSiteMultipleClients**
> UnauthorizeSiteMultipleClients(ctx, siteId, optional)
unauthorizeSiteMultipleClients

This unauthorize clients (if they are guest) and disconnect them. From the guest’s perspective, they will see the splash page again and go through the flow (e.g. Terms of Use) again.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesWiFiApiUnauthorizeSiteMultipleClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesWiFiApiUnauthorizeSiteMultipleClientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ClientsUnauthorizeBody**](ClientsUnauthorizeBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnauthorizeSiteWirelessClient**
> UnauthorizeSiteWirelessClient(ctx, siteId, clientMac)
unauthorizeSiteWirelessClient

This unauthorize a client (if it’s a guest) and disconnect it. From the guest’s perspective, s/he will see the splash page again and go through the flow (e.g. Terms of Use) again.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **clientMac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ZeroizeSiteFipsAllAps**
> ZeroizeSiteFipsAllAps(ctx, siteId, optional)
zeroizeSiteFipsAllAps

Zeroize all FIPS APs in the Site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
 **optional** | ***UtilitiesWiFiApiZeroizeSiteFipsAllApsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UtilitiesWiFiApiZeroizeSiteFipsAllApsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DevicesZeroizeBody**](DevicesZeroizeBody.md)| Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

