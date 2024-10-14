# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLicenseTypes**](ConstantsDefinitionsApi.md#GetLicenseTypes) | **Get** /api/v1/const/license_types | getLicenseTypes
[**ListAlarmDefinitions**](ConstantsDefinitionsApi.md#ListAlarmDefinitions) | **Get** /api/v1/const/alarm_defs | listAlarmDefinitions
[**ListApChannels**](ConstantsDefinitionsApi.md#ListApChannels) | **Get** /api/v1/const/ap_channels | listApChannels
[**ListApLedDefinition**](ConstantsDefinitionsApi.md#ListApLedDefinition) | **Get** /api/v1/const/ap_led_status | listApLedDefinition
[**ListAppCategoryDefinitions**](ConstantsDefinitionsApi.md#ListAppCategoryDefinitions) | **Get** /api/v1/const/app_categories | listAppCategoryDefinitions
[**ListAppSubCategoryDefinitions**](ConstantsDefinitionsApi.md#ListAppSubCategoryDefinitions) | **Get** /api/v1/const/app_subcategories | listAppSubCategoryDefinitions
[**ListApplications**](ConstantsDefinitionsApi.md#ListApplications) | **Get** /api/v1/const/applications | listApplications
[**ListCountryCodes**](ConstantsDefinitionsApi.md#ListCountryCodes) | **Get** /api/v1/const/countries | listCountryCodes
[**ListGatewayApplications**](ConstantsDefinitionsApi.md#ListGatewayApplications) | **Get** /api/v1/const/gateway_applications | listGatewayApplications
[**ListInsightMetrics**](ConstantsDefinitionsApi.md#ListInsightMetrics) | **Get** /api/v1/const/insight_metrics | listInsightMetrics
[**ListSiteLanguages**](ConstantsDefinitionsApi.md#ListSiteLanguages) | **Get** /api/v1/const/languages | listSiteLanguages
[**ListStates**](ConstantsDefinitionsApi.md#ListStates) | **Get** /api/v1/const/states | listStates
[**ListTrafficTypes**](ConstantsDefinitionsApi.md#ListTrafficTypes) | **Get** /api/v1/const/traffic_types | listTrafficTypes

# **GetLicenseTypes**
> []ConstLicenseType GetLicenseTypes(ctx, )
getLicenseTypes

Get License Types

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ConstLicenseType**](const_license_type.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAlarmDefinitions**
> []ConstAlarmDefinition ListAlarmDefinitions(ctx, )
listAlarmDefinitions

Get List of brief definitions of all the supported alarm types.  The example field contains an example payload as you would recieve in the alarm webhook output.   HA cluster node names will be specified in the `node` field, if applicable.'

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ConstAlarmDefinition**](const_alarm_definition.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApChannels**
> InlineResponse200180 ListApChannels(ctx, optional)
listApChannels

Get List of List of Available channels per country code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConstantsDefinitionsApiListApChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConstantsDefinitionsApiListApChannelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCode** | **optional.String**| country code, in two-character | 

### Return type

[**InlineResponse200180**](inline_response_200_180.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApLedDefinition**
> []ConstApLed ListApLedDefinition(ctx, )
listApLedDefinition

Get List of AP LED definition

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ConstApLed**](const_ap_led.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAppCategoryDefinitions**
> []ConstAppCategoryDefinition ListAppCategoryDefinitions(ctx, )
listAppCategoryDefinitions

Get List of definitions of all the supported Application Categories. The example field contains an example payload as you would recieve in the alarm webhook output.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ConstAppCategoryDefinition**](const_app_category_definition.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAppSubCategoryDefinitions**
> []ConstAppSubcategoryDefinition ListAppSubCategoryDefinitions(ctx, )
listAppSubCategoryDefinitions

Get List of definitions of all the supported Application sub-categories. The example field contains an example payload as you would recieve in the alarm webhook output.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ConstAppSubcategoryDefinition**](const_app_subcategory_definition.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApplications**
> []ConstApplicationDefinition ListApplications(ctx, )
listApplications

Get List of a list of applications that Juniper-Mist APs recognize

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ConstApplicationDefinition**](const_application_definition.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCountryCodes**
> []ConstCountry ListCountryCodes(ctx, optional)
listCountryCodes

Get List of available Country Codes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConstantsDefinitionsApiListCountryCodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConstantsDefinitionsApiListCountryCodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extend** | **optional.Bool**| will include more country codes if true | [default to false]

### Return type

[**[]ConstCountry**](const_country.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGatewayApplications**
> []ConstGatewayApplicationsDefinition ListGatewayApplications(ctx, )
listGatewayApplications

Get the full list of applications that we recognize

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ConstGatewayApplicationsDefinition**](const_gateway_applications_definition.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInsightMetrics**
> map[string]ConstInsightMetricsProperty ListInsightMetrics(ctx, )
listInsightMetrics

List Insight Metrics

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**map[string]ConstInsightMetricsProperty**](const_insight_metrics_property.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSiteLanguages**
> []ConstLanguage ListSiteLanguages(ctx, )
listSiteLanguages

Get List of Languages

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ConstLanguage**](const_language.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStates**
> []ConstState ListStates(ctx, countryCode)
listStates

Get List of ISO States based on country code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **countryCode** | **string**| country code, in [two-character]($e/Constants%20Definitions/listCountryCodes) | 

### Return type

[**[]ConstState**](const_state.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTrafficTypes**
> []ConstTrafficType ListTrafficTypes(ctx, )
listTrafficTypes

Get List of identified traffic

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ConstTrafficType**](const_traffic_type.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

