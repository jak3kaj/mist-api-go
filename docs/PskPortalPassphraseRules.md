# PskPortalPassphraseRules

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlphabertsEnabled** | **bool** |  | [optional] [default to true]
**Length** | **int32** |  | [optional] [default to null]
**MaxLength** | **int32** | for valid &#x60;max_length&#x60; and &#x60;min_length&#x60;, passphrase size is set randomly from that range. - if &#x60;max_length&#x60; and/or &#x60;min_length&#x60; are invalid, passphrase size is equal to &#x60;length&#x60; parameter - if &#x60;length&#x60; is not set or is invalid, default passphrase size is 8. valid &#x60;max_length&#x60;, &#x60;min_length&#x60;, &#x60;length&#x60; should be an integer between 8 to 63. Also, &#x60;max_length&#x60; &gt; &#x60;min_length&#x60; | [optional] [default to null]
**MinLength** | **int32** | for valid &#x60;max_length&#x60; and &#x60;min_length&#x60;, passphrase size is set randomly from that range. - if &#x60;max_length&#x60; and/or &#x60;min_length&#x60; are invalid, passphrase size is equal to &#x60;length&#x60; parameter - if &#x60;length&#x60; is not set or is invalid, default passphrase size is 8. valid &#x60;max_length&#x60;, &#x60;min_length&#x60;, &#x60;length&#x60; should be an integer between 8 to 63. Also, &#x60;max_length&#x60; &gt; &#x60;min_length&#x60; | [optional] [default to null]
**NumericsEnabled** | **bool** |  | [optional] [default to true]
**Symbols** | **string** |  | [optional] [default to null]
**SymbolsEnabled** | **bool** |  | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

