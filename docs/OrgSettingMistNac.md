# OrgSettingMistNac

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cacerts** | **[]string** | list of PEM-encoded ca certs | [optional] [default to null]
**DefaultIdpId** | **string** | use this IDP when no explicit realm present in the incoming username/CN OR when no IDP is explicitly mapped to the incoming realm. | [optional] [default to null]
**DisableRsaeAlgorithms** | **bool** | to disable RSAE_PSS_SHA256, RSAE_PSS_SHA384, RSAE_PSS_SHA512 from server side. see https://www.openssl.org/docs/man3.0/man1/openssl-ciphers.html | [optional] [default to false]
**EapSslSecurityLevel** | **int32** | eap ssl security level see https://www.openssl.org/docs/man1.1.1/man3/SSL_CTX_set_security_level.html#DEFAULT-CALLBACK-BEHAVIOUR | [optional] [default to 2]
**EuOnly** | **bool** | By default NAC POD failover considers all NAC pods available around the globe, i.e. EU, US, or APAC based, failover happens based on geo IP of the originating site. For strict GDPR compliancy NAC POD failover would only happen between the PODs located within the EU environment, and no authentication would take place outside of EU. This is an org setting that is applicable to WLANs, switch templates, mxedge clusters that have mist_nac enabled | [optional] [default to false]
**IdpMachineCertLookupField** | [***AllOforgSettingMistNacIdpMachineCertLookupField**](AllOforgSettingMistNacIdpMachineCertLookupField.md) |  | [optional] [default to null]
**IdpUserCertLookupField** | [***AllOforgSettingMistNacIdpUserCertLookupField**](AllOforgSettingMistNacIdpUserCertLookupField.md) |  | [optional] [default to null]
**Idps** | [**[]OrgSettingMistNacIdp**](org_setting_mist_nac_idp.md) |  | [optional] [default to null]
**ServerCert** | [***AllOforgSettingMistNacServerCert**](AllOforgSettingMistNacServerCert.md) |  | [optional] [default to null]
**UseIpVersion** | [***AllOforgSettingMistNacUseIpVersion**](AllOforgSettingMistNacUseIpVersion.md) |  | [optional] [default to null]
**UseSslPort** | **bool** | By default NAS devices (switches/aps) and proxies(mxedge) are configured to use port TCP2083(radsec) to reach mist-nac.  Set &#x60;use_ssl_port&#x60;&#x3D;&#x3D;&#x60;true&#x60; to override that port with TCP43 (ssl),  This is a org level setting that is applicable to wlans, switch_templates, and mxedge_clusters that have mist-nac enabled | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

