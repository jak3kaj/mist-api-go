/*
 * Mist API
 *
 * > Version: **2409.1.9** > > Date: **September 27, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates] (https://www.juniper.net/documentation/us/en/software/mist/api/http/getting-started/how-to-get-started)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 
 *
 * API version: 2409.1.9
 * Contact: tmunzer@juniper.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package mist

// portal template wlan settings
type WlanPortalTemplateSetting struct {
	AccessCodeAlternateEmail string `json:"accessCodeAlternateEmail,omitempty"`
	Alignment *AllOfwlanPortalTemplateSettingAlignment `json:"alignment,omitempty"`
	Ar *WlanPortalTemplateSettingLocale `json:"ar,omitempty"`
	// label for Amazon auth button
	AuthButtonAmazon string `json:"authButtonAmazon,omitempty"`
	// label for Azure auth button
	AuthButtonAzure string `json:"authButtonAzure,omitempty"`
	// label for Email auth button
	AuthButtonEmail string `json:"authButtonEmail,omitempty"`
	// label for Facebook auth button
	AuthButtonFacebook string `json:"authButtonFacebook,omitempty"`
	// label for Google auth button
	AuthButtonGoogle string `json:"authButtonGoogle,omitempty"`
	// label for Microsoft auth button
	AuthButtonMicrosoft string `json:"authButtonMicrosoft,omitempty"`
	// label for passphrase auth button
	AuthButtonPassphrase string `json:"authButtonPassphrase,omitempty"`
	// label for SMS auth button
	AuthButtonSms string `json:"authButtonSms,omitempty"`
	// label for Sponsor auth button
	AuthButtonSponsor string `json:"authButtonSponsor,omitempty"`
	AuthLabel string `json:"authLabel,omitempty"`
	// label of the link to go back to /logon
	BackLink string `json:"backLink,omitempty"`
	CaES *WlanPortalTemplateSettingLocale `json:"ca-ES,omitempty"`
	// Portal main color
	Color string `json:"color,omitempty"`
	ColorDark string `json:"colorDark,omitempty"`
	ColorLight string `json:"colorLight,omitempty"`
	// whether company field is required
	Company bool `json:"company,omitempty"`
	// error message when company not provided
	CompanyError string `json:"companyError,omitempty"`
	// label of company field
	CompanyLabel string `json:"companyLabel,omitempty"`
	CsCZ *WlanPortalTemplateSettingLocale `json:"cs-CZ,omitempty"`
	DaDK *WlanPortalTemplateSettingLocale `json:"da-DK,omitempty"`
	DeDE *WlanPortalTemplateSettingLocale `json:"de-DE,omitempty"`
	ElGR *WlanPortalTemplateSettingLocale `json:"el-GR,omitempty"`
	// whether email field is required
	Email bool `json:"email,omitempty"`
	// error message when a user has valid social login but doesn't match specified email domains.
	EmailAccessDomainError string `json:"emailAccessDomainError,omitempty"`
	// Label for cancel confirmation code submission using email auth
	EmailCancel string `json:"emailCancel,omitempty"`
	EmailCodeCancel string `json:"emailCodeCancel,omitempty"`
	EmailCodeError string `json:"emailCodeError,omitempty"`
	EmailCodeFieldLabel string `json:"emailCodeFieldLabel,omitempty"`
	EmailCodeMessage string `json:"emailCodeMessage,omitempty"`
	EmailCodeSubmit string `json:"emailCodeSubmit,omitempty"`
	EmailCodeTitle string `json:"emailCodeTitle,omitempty"`
	// error message when email not provided
	EmailError string `json:"emailError,omitempty"`
	EmailFieldLabel string `json:"emailFieldLabel,omitempty"`
	// label of email field
	EmailLabel string `json:"emailLabel,omitempty"`
	EmailMessage string `json:"emailMessage,omitempty"`
	// Label for confirmation code submit button using email auth
	EmailSubmit string `json:"emailSubmit,omitempty"`
	// Title for the Email registration
	EmailTitle string `json:"emailTitle,omitempty"`
	EnGB *WlanPortalTemplateSettingLocale `json:"en-GB,omitempty"`
	EnUS *WlanPortalTemplateSettingLocale `json:"en-US,omitempty"`
	EsES *WlanPortalTemplateSettingLocale `json:"es-ES,omitempty"`
	FiFI *WlanPortalTemplateSettingLocale `json:"fi-FI,omitempty"`
	// whether to ask field1
	Field1 bool `json:"field1,omitempty"`
	// error message when field1 not provided
	Field1Error string `json:"field1Error,omitempty"`
	// label of field1
	Field1Label string `json:"field1Label,omitempty"`
	// whether field1 is required field
	Field1Required bool `json:"field1Required,omitempty"`
	// whether to ask field2
	Field2 bool `json:"field2,omitempty"`
	// error message when field2 not provided
	Field2Error string `json:"field2Error,omitempty"`
	// label of field2
	Field2Label string `json:"field2Label,omitempty"`
	// whether field2 is required field
	Field2Required bool `json:"field2Required,omitempty"`
	// whether to ask field3
	Field3 bool `json:"field3,omitempty"`
	// error message when field3 not provided
	Field3Error string `json:"field3Error,omitempty"`
	// label of field3
	Field3Label string `json:"field3Label,omitempty"`
	// whether field3 is required field
	Field3Required bool `json:"field3Required,omitempty"`
	// whether to ask field4
	Field4 bool `json:"field4,omitempty"`
	// error message when field4 not provided
	Field4Error string `json:"field4Error,omitempty"`
	// label of field4
	Field4Label string `json:"field4Label,omitempty"`
	// whether field4 is required field
	Field4Required bool `json:"field4Required,omitempty"`
	FrFR *WlanPortalTemplateSettingLocale `json:"fr-FR,omitempty"`
	HeIL *WlanPortalTemplateSettingLocale `json:"he-IL,omitempty"`
	HiIN *WlanPortalTemplateSettingLocale `json:"hi-IN,omitempty"`
	HrHR *WlanPortalTemplateSettingLocale `json:"hr-HR,omitempty"`
	HuHU *WlanPortalTemplateSettingLocale `json:"hu-HU,omitempty"`
	IdID *WlanPortalTemplateSettingLocale `json:"id-ID,omitempty"`
	ItIT *WlanPortalTemplateSettingLocale `json:"it-IT,omitempty"`
	JaJP *WlanPortalTemplateSettingLocale `json:"ja-JP,omitempty"`
	KoKR *WlanPortalTemplateSettingLocale `json:"ko-KR,omitempty"`
	// custom logo with `data:image/png;base64,` format, default null, uses Juniper Mist Logo. File size must be less than 100kB and image dimensions must be less than 500px x 200px (width x height).
	Logo string `json:"logo,omitempty"`
	// height of the logo, in px
	LogoHeight int32 `json:"logoHeight,omitempty"`
	// width of the logo, in px
	LogoWidth int32 `json:"logoWidth,omitempty"`
	Message string `json:"message,omitempty"`
	MsMY *WlanPortalTemplateSettingLocale `json:"ms-MY,omitempty"`
	MultiAuth bool `json:"multiAuth,omitempty"`
	// whether name field is required
	Name bool `json:"name,omitempty"`
	// error message when name not provided
	NameError string `json:"nameError,omitempty"`
	// label of name field
	NameLabel string `json:"nameLabel,omitempty"`
	NbNO *WlanPortalTemplateSettingLocale `json:"nb-NO,omitempty"`
	NlNL *WlanPortalTemplateSettingLocale `json:"nl-NL,omitempty"`
	// Default value for the `Do not store` checkbox
	OptOutDefault bool `json:"optOutDefault,omitempty"`
	// whether to display Do Not Store My Personal Information
	Optout bool `json:"optout,omitempty"`
	// label for Do Not Store My Personal Information
	OptoutLabel string `json:"optoutLabel,omitempty"`
	PageTitle string `json:"pageTitle"`
	// Label for the Passphrase cancel button
	PassphraseCancel string `json:"passphraseCancel,omitempty"`
	// error message when invalid passphrase is provided
	PassphraseError string `json:"passphraseError,omitempty"`
	// Passphrase
	PassphraseLabel string `json:"passphraseLabel,omitempty"`
	PassphraseMessage string `json:"passphraseMessage,omitempty"`
	// Label for the Passphrase submit button
	PassphraseSubmit string `json:"passphraseSubmit,omitempty"`
	// Title for passphrase details page
	PassphraseTitle string `json:"passphraseTitle,omitempty"`
	PlPL *WlanPortalTemplateSettingLocale `json:"pl-PL,omitempty"`
	// whether to show \\\"Powered by Mist\\\"
	PoweredBy bool `json:"poweredBy,omitempty"`
	// wheter to require the Privacy Term acceptance
	Privacy bool `json:"privacy,omitempty"`
	// prefix of the label of the link to go to Privacy Policy
	PrivacyPolicyAcceptLabel string `json:"privacyPolicyAcceptLabel,omitempty"`
	// error message when Privacy Policy not accepted
	PrivacyPolicyError string `json:"privacyPolicyError,omitempty"`
	// label of the link to go to Privacy Policy
	PrivacyPolicyLink string `json:"privacyPolicyLink,omitempty"`
	// text of the Privacy Policy
	PrivacyPolicyText string `json:"privacyPolicyText,omitempty"`
	PtBR *WlanPortalTemplateSettingLocale `json:"pt-BR,omitempty"`
	PtPT *WlanPortalTemplateSettingLocale `json:"pt-PT,omitempty"`
	// label to denote required field
	RequiredFieldLabel string `json:"requiredFieldLabel,omitempty"`
	ResponsiveLayout bool `json:"responsiveLayout,omitempty"`
	RoRO *WlanPortalTemplateSettingLocale `json:"ro-RO,omitempty"`
	RuRU *WlanPortalTemplateSettingLocale `json:"ru-RU,omitempty"`
	// label of the button to /signin
	SignInLabel string `json:"signInLabel,omitempty"`
	SkSK *WlanPortalTemplateSettingLocale `json:"sk-SK,omitempty"`
	SmsCarrierDefault string `json:"smsCarrierDefault,omitempty"`
	SmsCarrierError string `json:"smsCarrierError,omitempty"`
	// label for mobile carrier drop-down list
	SmsCarrierFieldLabel string `json:"smsCarrierFieldLabel,omitempty"`
	// Label for cancel confirmation code submission
	SmsCodeCancel string `json:"smsCodeCancel,omitempty"`
	// error message when confirmation code is invalid
	SmsCodeError string `json:"smsCodeError,omitempty"`
	SmsCodeFieldLabel string `json:"smsCodeFieldLabel,omitempty"`
	SmsCodeMessage string `json:"smsCodeMessage,omitempty"`
	// Label for confirmation code submit button
	SmsCodeSubmit string `json:"smsCodeSubmit,omitempty"`
	SmsCodeTitle string `json:"smsCodeTitle,omitempty"`
	SmsCountryFieldLabel string `json:"smsCountryFieldLabel,omitempty"`
	SmsCountryFormat string `json:"smsCountryFormat,omitempty"`
	// Label for checkbox to specify that the user has access code
	SmsHaveAccessCode string `json:"smsHaveAccessCode,omitempty"`
	SmsIsTwilio bool `json:"smsIsTwilio,omitempty"`
	// format of access code sms message. {{code}} and {{duration}} are place holders and should be retained as is.
	SmsMessageFormat string `json:"smsMessageFormat,omitempty"`
	// label for canceling mobile details for SMS auth
	SmsNumberCancel string `json:"smsNumberCancel,omitempty"`
	SmsNumberError string `json:"smsNumberError,omitempty"`
	// label for field to provide mobile number
	SmsNumberFieldLabel string `json:"smsNumberFieldLabel,omitempty"`
	SmsNumberFormat string `json:"smsNumberFormat,omitempty"`
	SmsNumberMessage string `json:"smsNumberMessage,omitempty"`
	// label for submit button for code generation
	SmsNumberSubmit string `json:"smsNumberSubmit,omitempty"`
	// Title for phone number details
	SmsNumberTitle string `json:"smsNumberTitle,omitempty"`
	SmsUsernameFormat string `json:"smsUsernameFormat,omitempty"`
	// how long confirmation code should be considered valid (in minutes)
	SmsValidityDuration int32 `json:"smsValidityDuration,omitempty"`
	SponsorBackLink string `json:"sponsorBackLink,omitempty"`
	SponsorCancel string `json:"sponsorCancel,omitempty"`
	// label for Sponsor Email
	SponsorEmail string `json:"sponsorEmail,omitempty"`
	SponsorEmailError string `json:"sponsorEmailError,omitempty"`
	// html template to replace/override default sponsor email template  Sponsor Email Template supports following template variables:   * `approve_url`: Renders URL to approve the request; optionally &minutes=N query param can be appended to change the Authorization period of the guest, where N is a valid integer denoting number of minutes a guest remains authorized   * `deny_url`: Renders URL to reject the request   * `guest_email`: Renders Email ID of the guest   * `guest_name`: Renders Name of the guest   * `field1`: Renders value of the Custom Field 1   * `field2`: Renders value of the Custom Field 2   * `sponsor_link_validity_duration`: Renders validity time of the request (i.e. Approve/Deny URL)   * `auth_expire_minutes`: Renders Wlan-level configured Guest Authorization Expiration time period (in minutes), If not configured then default (1 day in minutes)
	SponsorEmailTemplate string `json:"sponsorEmailTemplate,omitempty"`
	SponsorInfoApproved string `json:"sponsorInfoApproved,omitempty"`
	SponsorInfoDenied string `json:"sponsorInfoDenied,omitempty"`
	SponsorInfoPending string `json:"sponsorInfoPending,omitempty"`
	// label for Sponsor Name
	SponsorName string `json:"sponsorName,omitempty"`
	SponsorNameError string `json:"sponsorNameError,omitempty"`
	SponsorNotePending string `json:"sponsorNotePending,omitempty"`
	// submit button label request Wifi Access and notify sponsor about guest request
	SponsorRequestAccess string `json:"sponsorRequestAccess,omitempty"`
	// text to display if sponsor approves request
	SponsorStatusApproved string `json:"sponsorStatusApproved,omitempty"`
	// text to display when sponsor denies request
	SponsorStatusDenied string `json:"sponsorStatusDenied,omitempty"`
	// text to display if request is still pending
	SponsorStatusPending string `json:"sponsorStatusPending,omitempty"`
	// submit button label to notify sponsor about guest request
	SponsorSubmit string `json:"sponsorSubmit,omitempty"`
	SponsorsError string `json:"sponsorsError,omitempty"`
	SponsorsFieldLabel string `json:"sponsorsFieldLabel,omitempty"`
	SvSE *WlanPortalTemplateSettingLocale `json:"sv-SE,omitempty"`
	ThTH *WlanPortalTemplateSettingLocale `json:"th-TH,omitempty"`
	Tos bool `json:"tos,omitempty"`
	// prefix of the label of the link to go to tos
	TosAcceptLabel string `json:"tosAcceptLabel,omitempty"`
	// error message when tos not accepted
	TosError string `json:"tosError,omitempty"`
	// label of the link to go to tos
	TosLink string `json:"tosLink,omitempty"`
	// text of the Terms of Service
	TosText string `json:"tosText,omitempty"`
	TrTR *WlanPortalTemplateSettingLocale `json:"tr-TR,omitempty"`
	UkUA *WlanPortalTemplateSettingLocale `json:"uk-UA,omitempty"`
	ViVN *WlanPortalTemplateSettingLocale `json:"vi-VN,omitempty"`
	ZhHans *WlanPortalTemplateSettingLocale `json:"zh-Hans,omitempty"`
	ZhHant *WlanPortalTemplateSettingLocale `json:"zh-Hant,omitempty"`
}
