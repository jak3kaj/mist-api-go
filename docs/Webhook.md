# Webhook

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | **float64** |  | [optional] [default to null]
**Enabled** | **bool** | whether webhook is enabled | [optional] [default to true]
**ForSite** | **bool** |  | [optional] [default to null]
**Headers** | **map[string]string** | if &#x60;type&#x60;&#x3D;&#x60;http-post&#x60;, additional custom HTTP headers to add the headers name and value must be string, total bytes of headers name and value must be less than 1000 | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**ModifiedTime** | **float64** |  | [optional] [default to null]
**Name** | **string** | name of the webhook | [optional] [default to null]
**Oauth2ClientId** | **string** | required when &#x60;oauth2_grant_type&#x60;&#x3D;&#x3D;&#x60;client_credentials&#x60; | [optional] [default to null]
**Oauth2ClientSecret** | **string** | required when &#x60;oauth2_grant_type&#x60;&#x3D;&#x3D;&#x60;client_credentials&#x60; | [optional] [default to null]
**Oauth2GrantType** | [***AllOfwebhookOauth2GrantType**](AllOfwebhookOauth2GrantType.md) |  | [optional] [default to null]
**Oauth2Password** | **string** | required when &#x60;oauth2_grant_type&#x60;&#x3D;&#x3D;&#x60;password&#x60; | [optional] [default to null]
**Oauth2Scopes** | **[]string** | required when &#x60;type&#x60;&#x3D;&#x3D;&#x60;oauth2&#x60;, if provided, will be used in the token request | [optional] [default to null]
**Oauth2TokenUrl** | **string** | required when &#x60;type&#x60;&#x3D;&#x3D;&#x60;oauth2&#x60; | [optional] [default to null]
**Oauth2Username** | **string** | required when &#x60;oauth2_grant_type&#x60;&#x3D;&#x3D;&#x60;password&#x60; | [optional] [default to null]
**OrgId** | **string** |  | [optional] [default to null]
**Secret** | **string** | only if &#x60;type&#x60;&#x3D;&#x60;http-post&#x60;   when &#x60;secret&#x60; is provided, two  HTTP headers will be added:    * X-Mist-Signature-v2: HMAC_SHA256(secret, body)   * X-Mist-Signature: HMAC_SHA1(secret, body) | [optional] [default to null]
**SiteId** | **string** |  | [optional] [default to null]
**SplunkToken** | **string** | required if &#x60;type&#x60;&#x3D;&#x60;splunk&#x60; If splunk_token is not defined for a type Splunk webhook, it will not send, regardless if the webhook receiver is configured to accept it.&#x27; | [optional] [default to null]
**Topics** | [**[]WebhookTopic**](webhook_topic.md) | N.B. For org webhooks, only alarms/audits/client-info/client-join/client-sessions/device-events/device-updowns/mxedge-events/nac-sessions/nac_events topics are supported. | [optional] [default to null]
**Type_** | [***AllOfwebhookType_**](AllOfwebhookType_.md) |  | [optional] [default to null]
**Url** | **string** |  | [optional] [default to null]
**VerifyCert** | **bool** | when url uses HTTPS, whether to verify the certificate | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

