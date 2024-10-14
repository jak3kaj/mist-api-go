# OrgAutoRules

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateNewSiteIfNeeded** | **bool** | if &#x60;src&#x60;&#x3D;&#x3D;&#x60;geoip&#x60;. By default, a claimed device only gets assigned if the site exists to auto-create the site, enable this | [optional] [default to false]
**Expression** | **string** | if &#x60;src&#x60;&#x3D;&#x3D;&#x60;name&#x60;, &#x60;src&#x60;&#x3D;&#x3D;&#x60;lldp_system_name&#x60;,  &#x60;src&#x60;&#x3D;&#x3D;&#x60;dns_suffix&#x60;       \&quot;[0:3]\&quot;            // \&quot;abcdef\&quot; -&gt; \&quot;abc\&quot;       \&quot;split(.)[1]\&quot;      // \&quot;a.b.c\&quot; -&gt; \&quot;b\&quot;       \&quot;split(-)[1][0:3]\&quot; // \&quot;a1234-b5678-c90\&quot; -&gt; \&quot;b56\&quot;&#x27; | [optional] [default to null]
**GatewaytemplateId** | **string** | if &#x60;src&#x60;&#x3D;&#x3D;&#x60;geoip&#x60; and &#x60;create_new_site_if_needed&#x60;&#x3D;&#x3D;&#x60;true&#x60;. If a gateway template is desired for this newly created site | [optional] [default to null]
**MatchCountry** | **string** | if &#x60;src&#x60;&#x3D;&#x3D;&#x60;geoip&#x60; | [optional] [default to null]
**MatchDeviceType** | [***AllOforgAutoRulesMatchDeviceType**](AllOforgAutoRulesMatchDeviceType.md) |  | [optional] [default to null]
**MatchModel** | **string** | optional/additional filter | [optional] [default to null]
**Model** | **string** | if &#x60;src&#x60;&#x3D;&#x3D;&#x60;model&#x60; | [optional] [default to null]
**Prefix** | **string** | if &#x60;src&#x60;&#x3D;&#x3D;&#x60;name&#x60; | [optional] [default to null]
**Src** | [***AllOforgAutoRulesSrc**](AllOforgAutoRulesSrc.md) |  | [default to null]
**Subnet** | **string** | if &#x60;src&#x60;&#x3D;&#x3D;&#x60;subnet&#x60; | [optional] [default to null]
**Suffix** | **string** | if &#x60;src&#x60;&#x3D;&#x3D;&#x60;name&#x60; | [optional] [default to null]
**Value** | **string** | if      * &#x60;src&#x60;&#x3D;&#x3D;&#x60;model&#x60;     *  &#x60;src&#x60;&#x3D;&#x3D;&#x60;geoip: site name for the device to be assigned to (\&quot;city\&quot; / \&quot;city+country\&quot; / ...) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

