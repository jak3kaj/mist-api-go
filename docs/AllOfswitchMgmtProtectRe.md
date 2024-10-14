# AllOfswitchMgmtProtectRe

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedServices** | [**[]ProtectReAllowedService**](protect_re_allowed_service.md) | optionally, services we&#x27;ll allow | [optional] [default to null]
**Custom** | [**[]ProtectReCustom**](protect_re_custom.md) |  | [optional] [default to null]
**Enabled** | **bool** | when enabled, all traffic that is not essential to our operation will be dropped e.g. ntp / dns / traffic to mist will be allowed by default      if dhcpd is enabled, we&#x27;ll make sure it works | [optional] [default to false]
**TrustedHosts** | **[]string** | host/subnets we&#x27;ll allow traffic to/from | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

