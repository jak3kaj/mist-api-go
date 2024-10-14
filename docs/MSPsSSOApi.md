# {{classname}}

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMspSso**](MSPsSSOApi.md#CreateMspSso) | **Post** /api/v1/msps/{msp_id}/ssos | createMspSso
[**DeleteMspSso**](MSPsSSOApi.md#DeleteMspSso) | **Delete** /api/v1/msps/{msp_id}/ssos/{sso_id} | deleteMspSso
[**DownloadMspSsoSamlMetadata**](MSPsSSOApi.md#DownloadMspSsoSamlMetadata) | **Get** /api/v1/msps/{msp_id}/ssos/{sso_id}/metadata.xml | downloadMspSsoSamlMetadata
[**GetMspSso**](MSPsSSOApi.md#GetMspSso) | **Get** /api/v1/msps/{msp_id}/ssos/{sso_id} | getMspSso
[**GetMspSsoSamlMetadata**](MSPsSSOApi.md#GetMspSsoSamlMetadata) | **Get** /api/v1/msps/{msp_id}/ssos/{sso_id}/metadata | getMspSsoSamlMetadata
[**ListMspSsoLatestFailures**](MSPsSSOApi.md#ListMspSsoLatestFailures) | **Get** /api/v1/msps/{msp_id}/ssos/{sso_id}/failures | listMspSsoLatestFailures
[**ListMspSsos**](MSPsSSOApi.md#ListMspSsos) | **Get** /api/v1/msps/{msp_id}/ssos | listMspSsos
[**UpdateMspSso**](MSPsSSOApi.md#UpdateMspSso) | **Put** /api/v1/msps/{msp_id}/ssos/{sso_id} | updateMspSso

# **CreateMspSso**
> InlineResponse20025 CreateMspSso(ctx, mspId, optional)
createMspSso

Create MSP SSO profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
 **optional** | ***MSPsSSOApiCreateMspSsoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsSSOApiCreateMspSsoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MspIdSsosBody**](MspIdSsosBody.md)| Request Body | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMspSso**
> DeleteMspSso(ctx, mspId, ssoId)
deleteMspSso

Delete MSP SSO Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **ssoId** | [**string**](.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadMspSsoSamlMetadata**
> *os.File DownloadMspSsoSamlMetadata(ctx, mspId, ssoId)
downloadMspSsoSamlMetadata

Download MSP SSO SAML Metadata  Example of metadata.xml: ```xml <?xml version=\"1.0\" encoding=\"UTF-8\"?><md:EntityDescriptor xmlns:md=\"urn:oasis:names:tc:SAML:2.0:metadata\" entityID=\"https://api.mist.com/api/v1/saml/5hdF5g/login\" validUntil=\"2027-10-12T21:59:01Z\" xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\">   <md:SPSSODescriptor AuthnRequestsSigned=\"false\" WantAssertionsSigned=\"true\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\">       <md:SingleLogoutService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://api.mist.com/api/v1/saml/5hdF5g/logout\" />       <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>       <md:AssertionConsumerService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://api.mist.com/api/v1/saml/5hdF5g/login\" index=\"0\" isDefault=\"true\"/>       <md:AttributeConsumingService index=\"0\">           <md:ServiceName xml:lang=\"en-US\">Mist</md:ServiceName>           <md:RequestedAttribute Name=\"Role\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:basic\" isRequired=\"true\"/>           <md:RequestedAttribute Name=\"FirstName\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:basic\" isRequired=\"false\"/>           <md:RequestedAttribute Name=\"LastName\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:basic\" isRequired=\"false\"/>       </md:AttributeConsumingService>   </md:SPSSODescriptor> </md:EntityDescriptor> ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **ssoId** | [**string**](.md)|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMspSso**
> InlineResponse20025 GetMspSso(ctx, mspId, ssoId)
getMspSso

Get MSP SSO Config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **ssoId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMspSsoSamlMetadata**
> InlineResponse20027 GetMspSsoSamlMetadata(ctx, mspId, ssoId)
getMspSsoSamlMetadata

Get MSP SSO SAML Metadata

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **ssoId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMspSsoLatestFailures**
> InlineResponse20026 ListMspSsoLatestFailures(ctx, mspId, ssoId)
listMspSsoLatestFailures

Get List of MSP SSO Latest Failures

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **ssoId** | [**string**](.md)|  | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMspSsos**
> []Sso ListMspSsos(ctx, mspId)
listMspSsos

List MSP SSO Configs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 

### Return type

[**[]Sso**](sso.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMspSso**
> InlineResponse20025 UpdateMspSso(ctx, mspId, ssoId, optional)
updateMspSso

Update MSP SSO config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mspId** | [**string**](.md)|  | 
  **ssoId** | [**string**](.md)|  | 
 **optional** | ***MSPsSSOApiUpdateMspSsoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MSPsSSOApiUpdateMspSsoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SsosSsoIdBody**](SsosSsoIdBody.md)| Request Body | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth), [csrfToken](../README.md#csrfToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

