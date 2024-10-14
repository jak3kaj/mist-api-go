# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteSleClassifierDetails**](SitesSLEsApi.md#GetSiteSleClassifierDetails) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/classifier/{classifier}/summary | getSiteSleClassifierDetails
[**GetSiteSleHistogram**](SitesSLEsApi.md#GetSiteSleHistogram) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/histogram | getSiteSleHistogram
[**GetSiteSleImpactSummary**](SitesSLEsApi.md#GetSiteSleImpactSummary) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/impact-summary | getSiteSleImpactSummary
[**GetSiteSleImpactedApplications**](SitesSLEsApi.md#GetSiteSleImpactedApplications) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/impacted-applications | getSiteSleImpactedApplications
[**GetSiteSleImpactedAps**](SitesSLEsApi.md#GetSiteSleImpactedAps) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/impacted-aps | getSiteSleImpactedAps
[**GetSiteSleImpactedChassis**](SitesSLEsApi.md#GetSiteSleImpactedChassis) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/impacted-chassis | getSiteSleImpactedChassis
[**GetSiteSleImpactedGateways**](SitesSLEsApi.md#GetSiteSleImpactedGateways) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/impacted-gateways | getSiteSleImpactedGateways
[**GetSiteSleImpactedInterfaces**](SitesSLEsApi.md#GetSiteSleImpactedInterfaces) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/impacted-interfaces | getSiteSleImpactedInterfaces
[**GetSiteSleImpactedSwitches**](SitesSLEsApi.md#GetSiteSleImpactedSwitches) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/impacted-switches | getSiteSleImpactedSwitches
[**GetSiteSleImpactedWiredClients**](SitesSLEsApi.md#GetSiteSleImpactedWiredClients) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/impacted-clients | getSiteSleImpactedWiredClients
[**GetSiteSleImpactedWirelessClients**](SitesSLEsApi.md#GetSiteSleImpactedWirelessClients) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/impacted_users | getSiteSleImpactedWirelessClients
[**GetSiteSleMetricClassifiers**](SitesSLEsApi.md#GetSiteSleMetricClassifiers) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/classifiers | getSiteSleMetricClassifiers
[**GetSiteSleSummary**](SitesSLEsApi.md#GetSiteSleSummary) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/summary | getSiteSleSummary
[**GetSiteSleThreshold**](SitesSLEsApi.md#GetSiteSleThreshold) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/threshold | getSiteSleThreshold
[**GetSiteSlesMetrics**](SitesSLEsApi.md#GetSiteSlesMetrics) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | getSiteSlesMetrics
[**ReplaceSiteSleThreshold**](SitesSLEsApi.md#ReplaceSiteSleThreshold) | **Post** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/threshold | replaceSiteSleThreshold
[**UpdateSiteSleThreshold**](SitesSLEsApi.md#UpdateSiteSleThreshold) | **Put** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/threshold | updateSiteSleThreshold

# **GetSiteSleClassifierDetails**
> InlineResponse200146 GetSiteSleClassifierDetails(ctx, siteId, scope, scopeId, metric, classifier, optional)
getSiteSleClassifierDetails

Get SLE classifier details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **scope** | [**SleSummaryScope**](.md)|  | 
  **scopeId** | **string**| * site_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;site&#x60; * device_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;ap&#x60;, &#x60;scope&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;scope&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; * mac if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;client&#x60; | 
  **metric** | **string**| values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 
  **classifier** | **string**|  | 
 **optional** | ***SitesSLEsApiGetSiteSleClassifierDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSLEsApiGetSiteSleClassifierDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse200146**](inline_response_200_146.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSleHistogram**
> InlineResponse200147 GetSiteSleHistogram(ctx, siteId, scope, scopeId, metric, optional)
getSiteSleHistogram

Get the histogram for the SLE metric

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **scope** | [**SiteSleHistogramScopeParameters**](.md)|  | 
  **scopeId** | **string**| * site_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;site&#x60; * device_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;ap&#x60;, &#x60;scope&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;scope&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; * mac if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;client&#x60; | 
  **metric** | **string**| values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 
 **optional** | ***SitesSLEsApiGetSiteSleHistogramOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSLEsApiGetSiteSleHistogramOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse200147**](inline_response_200_147.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSleImpactSummary**
> InlineResponse200148 GetSiteSleImpactSummary(ctx, siteId, scope, scopeId, metric, optional)
getSiteSleImpactSummary

Get impact summary counts optionally filtered by classifier and failure type   * Wireless SLE Fields: `wlan`, `device_type`, `device_os` ,`band`, `ap`, `server`, `mxedge` * Wired SLE Fields: `switch`, `client`, `vlan`, `interface`, `chassis` * WAN SLE Fields: `gateway`, `client`, `interface`, `chassis`, `peer_path`, `gateway_zones`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **scope** | [**SiteSleImpactSummaryScopeParameters**](.md)|  | 
  **scopeId** | **string**| * site_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;site&#x60; * device_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;ap&#x60;, &#x60;scope&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;scope&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; * mac if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;client&#x60; | 
  **metric** | **string**| values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 
 **optional** | ***SitesSLEsApiGetSiteSleImpactSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSLEsApiGetSiteSleImpactSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **fields** | [**optional.Interface of SiteSleImpactSummaryFieldsParameter**](.md)|  | 
 **classifier** | **optional.String**|  | 

### Return type

[**InlineResponse200148**](inline_response_200_148.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSleImpactedApplications**
> InlineResponse200149 GetSiteSleImpactedApplications(ctx, siteId, scope, scopeId, metric, optional)
getSiteSleImpactedApplications

For WAN SLEs. Get list of impacted interfaces optionally filtered by classifier and failure type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **scope** | [**SiteSleScope**](.md)|  | 
  **scopeId** | [**string**](.md)|  | 
  **metric** | **string**| values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 
 **optional** | ***SitesSLEsApiGetSiteSleImpactedApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSLEsApiGetSiteSleImpactedApplicationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **classifier** | **optional.String**|  | 

### Return type

[**InlineResponse200149**](inline_response_200_149.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSleImpactedAps**
> InlineResponse200150 GetSiteSleImpactedAps(ctx, siteId, scope, scopeId, metric, optional)
getSiteSleImpactedAps

For Wireless SLEs. Get list of impacted APs optionally filtered by classifier and failure type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **scope** | [**SiteSleImpactedApsScopeParameters**](.md)|  | 
  **scopeId** | [**string**](.md)|  | 
  **metric** | **string**| values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 
 **optional** | ***SitesSLEsApiGetSiteSleImpactedApsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSLEsApiGetSiteSleImpactedApsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **classifier** | **optional.String**|  | 

### Return type

[**InlineResponse200150**](inline_response_200_150.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSleImpactedChassis**
> InlineResponse200151 GetSiteSleImpactedChassis(ctx, siteId, scope, scopeId, metric, optional)
getSiteSleImpactedChassis

For Wired and WAN SLEs. Get list of impacted interfaces optionally filtered by classifier and failure type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **scope** | [**SiteSleImpactedChassisScopeParameters**](.md)|  | 
  **scopeId** | [**string**](.md)|  | 
  **metric** | **string**| values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 
 **optional** | ***SitesSLEsApiGetSiteSleImpactedChassisOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSLEsApiGetSiteSleImpactedChassisOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **classifier** | **optional.String**|  | 

### Return type

[**InlineResponse200151**](inline_response_200_151.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSleImpactedGateways**
> InlineResponse200153 GetSiteSleImpactedGateways(ctx, siteId, scope, scopeId, metric, optional)
getSiteSleImpactedGateways

For WAN SLEs. Get list of impacted interfaces optionally filtered by classifier and failure type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **scope** | [**SiteSleImpactedGatewaysScopeParameters**](.md)|  | 
  **scopeId** | [**string**](.md)|  | 
  **metric** | **string**| values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 
 **optional** | ***SitesSLEsApiGetSiteSleImpactedGatewaysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSLEsApiGetSiteSleImpactedGatewaysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **classifier** | **optional.String**|  | 

### Return type

[**InlineResponse200153**](inline_response_200_153.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSleImpactedInterfaces**
> InlineResponse200154 GetSiteSleImpactedInterfaces(ctx, siteId, scope, scopeId, metric, optional)
getSiteSleImpactedInterfaces

For Wired and WAN SLEs. Get list of impacted interfaces optionally filtered by classifier and failure type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **scope** | [**SiteSleImpactedInterfacesScopeParameters**](.md)|  | 
  **scopeId** | [**string**](.md)|  | 
  **metric** | **string**| values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 
 **optional** | ***SitesSLEsApiGetSiteSleImpactedInterfacesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSLEsApiGetSiteSleImpactedInterfacesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **classifier** | **optional.String**|  | 

### Return type

[**InlineResponse200154**](inline_response_200_154.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSleImpactedSwitches**
> InlineResponse200155 GetSiteSleImpactedSwitches(ctx, siteId, scope, scopeId, metric, optional)
getSiteSleImpactedSwitches

For Wired SLEs. Get list of impacted switches optionally filtered by classifier and failure type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **scope** | [**SiteSleImpactedSwitchesScopeParameters**](.md)|  | 
  **scopeId** | [**string**](.md)|  | 
  **metric** | **string**| values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 
 **optional** | ***SitesSLEsApiGetSiteSleImpactedSwitchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSLEsApiGetSiteSleImpactedSwitchesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **classifier** | **optional.String**|  | 

### Return type

[**InlineResponse200155**](inline_response_200_155.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSleImpactedWiredClients**
> InlineResponse200152 GetSiteSleImpactedWiredClients(ctx, siteId, scope, scopeId, metric, optional)
getSiteSleImpactedWiredClients

For Wired SLEs. Get list of impacted interfaces optionally filtered by classifier and failure type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **scope** | [**SiteSleImpactedClientsScopeParameters**](.md)|  | 
  **scopeId** | [**string**](.md)|  | 
  **metric** | **string**| values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 
 **optional** | ***SitesSLEsApiGetSiteSleImpactedWiredClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSLEsApiGetSiteSleImpactedWiredClientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **classifier** | **optional.String**|  | 

### Return type

[**InlineResponse200152**](inline_response_200_152.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSleImpactedWirelessClients**
> InlineResponse200156 GetSiteSleImpactedWirelessClients(ctx, siteId, scope, scopeId, metric, optional)
getSiteSleImpactedWirelessClients

For Wireless SLEs. Get list of impacted wireless users optionally filtered by classifier and failure type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **scope** | [**SiteSleImpactedUsersScopeParameter**](.md)|  | 
  **scopeId** | [**string**](.md)|  | 
  **metric** | **string**| values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 
 **optional** | ***SitesSLEsApiGetSiteSleImpactedWirelessClientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSLEsApiGetSiteSleImpactedWirelessClientsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]
 **classifier** | **optional.String**|  | 

### Return type

[**InlineResponse200156**](inline_response_200_156.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSleMetricClassifiers**
> []string GetSiteSleMetricClassifiers(ctx, siteId, scope, scopeId, metric)
getSiteSleMetricClassifiers

Get the list of classifiers for a specific metric

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **scope** | [**SiteSleMetricClassifiersScopeParameters**](.md)|  | 
  **scopeId** | **string**| * site_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;site&#x60; * device_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;ap&#x60;, &#x60;scope&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;scope&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; * mac if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;client&#x60; | 
  **metric** | **string**| values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 

### Return type

**[]string**

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSleSummary**
> InlineResponse200157 GetSiteSleSummary(ctx, siteId, scope, scopeId, metric, optional)
getSiteSleSummary

Get the summary for the SLE metric

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **scope** | [**SiteSleMetricSummaryScopeParameters**](.md)|  | 
  **scopeId** | **string**| * site_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;site&#x60; * device_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;ap&#x60;, &#x60;scope&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;scope&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; * mac if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;client&#x60; | 
  **metric** | **string**| values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 
 **optional** | ***SitesSLEsApiGetSiteSleSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSLEsApiGetSiteSleSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **optional.Int32**| start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **optional.Int32**| end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **optional.String**| duration like 7d, 2w | [default to 1d]

### Return type

[**InlineResponse200157**](inline_response_200_157.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSleThreshold**
> InlineResponse200158 GetSiteSleThreshold(ctx, siteId, scope, scopeId, metric)
getSiteSleThreshold

Get the SLE threshold

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **scope** | [**SiteSleThresholdScopeParameter**](.md)|  | 
  **scopeId** | **string**| * site_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;site&#x60; * device_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;ap&#x60;, &#x60;scope&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;scope&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; * mac if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;client&#x60; | 
  **metric** | **string**| values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 

### Return type

[**InlineResponse200158**](inline_response_200_158.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSiteSlesMetrics**
> InlineResponse200159 GetSiteSlesMetrics(ctx, siteId, scope, scopeId)
getSiteSlesMetrics

Get the list of metrics for the given scope

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **scope** | [**SiteSleMetricsScopeParameters**](.md)|  | 
  **scopeId** | **string**| * site_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;site&#x60; * device_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;ap&#x60;, &#x60;scope&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;scope&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; * mac if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;client&#x60; | 

### Return type

[**InlineResponse200159**](inline_response_200_159.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceSiteSleThreshold**
> InlineResponse200158 ReplaceSiteSleThreshold(ctx, siteId, scope, scopeId, metric, optional)
replaceSiteSleThreshold

Replace the SLE threshold

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **scope** | [**SiteSleThresholdScopeParameter**](.md)|  | 
  **scopeId** | **string**| * site_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;site&#x60; * device_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;ap&#x60;, &#x60;scope&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;scope&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; * mac if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;client&#x60; | 
  **metric** | **string**| values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 
 **optional** | ***SitesSLEsApiReplaceSiteSleThresholdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSLEsApiReplaceSiteSleThresholdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of MetricThresholdBody1**](MetricThresholdBody1.md)|  | 

### Return type

[**InlineResponse200158**](inline_response_200_158.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSiteSleThreshold**
> InlineResponse200158 UpdateSiteSleThreshold(ctx, siteId, scope, scopeId, metric, optional)
updateSiteSleThreshold

Update the SLE threshold

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | [**string**](.md)|  | 
  **scope** | [**SiteSleThresholdScopeParameter**](.md)|  | 
  **scopeId** | **string**| * site_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;site&#x60; * device_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;ap&#x60;, &#x60;scope&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;scope&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; * mac if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;client&#x60; | 
  **metric** | **string**| values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 
 **optional** | ***SitesSLEsApiUpdateSiteSleThresholdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SitesSLEsApiUpdateSiteSleThresholdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of MetricThresholdBody**](MetricThresholdBody.md)|  | 

### Return type

[**InlineResponse200158**](inline_response_200_158.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

