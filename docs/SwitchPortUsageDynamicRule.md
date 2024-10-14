# SwitchPortUsageDynamicRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Equals** | **string** |  | [optional] [default to null]
**EqualsAny** | **[]string** | use &#x60;equals_any&#x60; to match any item in a list | [optional] [default to null]
**Expression** | **string** | \&quot;[0:3]\&quot;:\&quot;abcdef\&quot; -&gt; \&quot;abc\&quot; \&quot;split(.)[1]\&quot;: \&quot;a.b.c\&quot; -&gt; \&quot;b\&quot; \&quot;split(-)[1][0:3]: \&quot;a1234-b5678-c90\&quot; -&gt; \&quot;b56\&quot; | [optional] [default to null]
**Src** | [***AllOfswitchPortUsageDynamicRuleSrc**](AllOfswitchPortUsageDynamicRuleSrc.md) |  | [default to null]
**Usage** | **string** | &#x60;port_usage&#x60; name | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

