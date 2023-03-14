// Code generated by "admgencmd"; DO NOT EDIT.
package mdmcommands
// DeviceInformationPayload
type DeviceInformationPayload struct {
	Queries                []string `plist:"Queries"` // 82 array values defined in schema
	DeviceAttestationNonce []byte   `plist:"DeviceAttestationNonce,omitempty"`
	RequestType            string   `plist:"RequestType"`
}
// DeviceInformationCommand
type DeviceInformationCommand struct {
	Command     DeviceInformationPayload `plist:"Command"`
	CommandUUID string                   `plist:"CommandUUID"`
}
// OrganizationInfo
type OrganizationInfo struct {
	OrganizationName    string `plist:"OrganizationName"`
	OrganizationAddress string `plist:"OrganizationAddress,omitempty"`
	OrganizationPhone   string `plist:"OrganizationPhone,omitempty"`
	OrganizationEmail   string `plist:"OrganizationEmail,omitempty"`
	OrganizationMagic   string `plist:"OrganizationMagic,omitempty"`
}
// MDMOptions
type MDMOptions struct {
	ActivationLockAllowedWhileSupervised             bool `plist:"ActivationLockAllowedWhileSupervised,omitempty"`
	BootstrapTokenAllowed                            bool `plist:"BootstrapTokenAllowed,omitempty"`
	PromptUserToAllowBootstrapTokenForAuthentication bool `plist:"PromptUserToAllowBootstrapTokenForAuthentication,omitempty"`
}
// OSUpdateSettings
type OSUpdateSettings struct {
	CatalogURL                      string      `plist:"CatalogURL"`
	IsDefaultCatalog                bool        `plist:"IsDefaultCatalog"`
	PreviousScanDate                interface{} `plist:"PreviousScanDate"` // unknown type: <date>
	PreviousScanResult              string      `plist:"PreviousScanResult"`
	PerformPeriodicCheck            bool        `plist:"PerformPeriodicCheck"`
	AutomaticCheckEnabled           bool        `plist:"AutomaticCheckEnabled"`
	BackgroundDownloadEnabled       bool        `plist:"BackgroundDownloadEnabled"`
	AutomaticAppInstallationEnabled bool        `plist:"AutomaticAppInstallationEnabled"`
	AutomaticOSInstallationEnabled  bool        `plist:"AutomaticOSInstallationEnabled"`
	AutomaticSecurityUpdatesEnabled bool        `plist:"AutomaticSecurityUpdatesEnabled"`
}
// AutoSetupAdminAccountsItem
type AutoSetupAdminAccountsItem struct {
	GUID      string `plist:"GUID"`
	ShortName string `plist:"shortName"`
}
// ServiceSubscriptionProperty
type ServiceSubscriptionProperty struct {
	CarrierSettingsVersion   string `plist:"CarrierSettingsVersion"`
	CurrentCarrierNetwork    string `plist:"CurrentCarrierNetwork"`
	CurrentMCC               string `plist:"CurrentMCC"`
	CurrentMNC               string `plist:"CurrentMNC"`
	ICCID                    string `plist:"ICCID"`
	EID                      string `plist:"EID"`
	IMEI                     string `plist:"IMEI"`
	IsDataPreferred          bool   `plist:"IsDataPreferred"`
	IsRoaming                bool   `plist:"IsRoaming"`
	IsVoicePreferred         bool   `plist:"IsVoicePreferred"`
	Label                    string `plist:"Label"`
	LabelID                  string `plist:"LabelID"`
	MEID                     string `plist:"MEID"`
	PhoneNumber              string `plist:"PhoneNumber"`
	Slot                     string `plist:"Slot"`
	SubscriberCarrierNetwork string `plist:"SubscriberCarrierNetwork"`
}
// SoftwareUpdateSettings
type SoftwareUpdateSettings struct {
	RecommendationsCadence int `plist:"RecommendationsCadence"`
}
// AccessibilitySettings
type AccessibilitySettings struct {
	BoldTextEnabled            bool `plist:"BoldTextEnabled"`
	IncreaseContrastEnabled    bool `plist:"IncreaseContrastEnabled"`
	ReduceMotionEnabled        bool `plist:"ReduceMotionEnabled"`
	ReduceTransparencyEnabled  bool `plist:"ReduceTransparencyEnabled"`
	TextSize                   int  `plist:"TextSize"`
	TouchAccommodationsEnabled bool `plist:"TouchAccommodationsEnabled"`
	VoiceOverEnabled           bool `plist:"VoiceOverEnabled"`
	ZoomEnabled                bool `plist:"ZoomEnabled"`
}
// QueryResponses
type QueryResponses struct {
	UDID                                  string                        `plist:"UDID"`
	ProvisioningUDID                      string                        `plist:"ProvisioningUDID"`
	OrganizationInfo                      OrganizationInfo              `plist:"OrganizationInfo"`
	MDMOptions                            MDMOptions                    `plist:"MDMOptions"`
	LastCloudBackupDate                   interface{}                   `plist:"LastCloudBackupDate"` // unknown type: <date>
	AwaitingConfiguration                 bool                          `plist:"AwaitingConfiguration"`
	ITunesStoreAccountIsActive            bool                          `plist:"iTunesStoreAccountIsActive"`
	ITunesStoreAccountHash                string                        `plist:"iTunesStoreAccountHash"`
	DeviceName                            string                        `plist:"DeviceName"`
	OSVersion                             string                        `plist:"OSVersion"`
	SupplementalOSVersionExtra            string                        `plist:"SupplementalOSVersionExtra"`
	BuildVersion                          string                        `plist:"BuildVersion"`
	SupplementalBuildVersion              string                        `plist:"SupplementalBuildVersion"`
	ModelName                             string                        `plist:"ModelName"`
	Model                                 string                        `plist:"Model"`
	IsAppleSilicon                        bool                          `plist:"IsAppleSilicon"`
	ProductName                           string                        `plist:"ProductName"`
	SerialNumber                          string                        `plist:"SerialNumber"`
	DeviceCapacity                        float64                       `plist:"DeviceCapacity"`
	AvailableDeviceCapacity               float64                       `plist:"AvailableDeviceCapacity"`
	IMEI                                  string                        `plist:"IMEI"`
	MEID                                  string                        `plist:"MEID"`
	ModemFirmwareVersion                  string                        `plist:"ModemFirmwareVersion"`
	CellularTechnology                    int                           `plist:"CellularTechnology"`
	BatteryLevel                          float64                       `plist:"BatteryLevel"`
	IsSupervised                          bool                          `plist:"IsSupervised"`
	IsMultiUser                           bool                          `plist:"IsMultiUser"`
	IsDeviceLocatorServiceEnabled         bool                          `plist:"IsDeviceLocatorServiceEnabled"`
	IsActivationLockEnabled               bool                          `plist:"IsActivationLockEnabled"`
	IsActivationLockSupported             bool                          `plist:"IsActivationLockSupported"`
	IsDoNotDisturbInEffect                bool                          `plist:"IsDoNotDisturbInEffect"`
	SupportsLOMDevice                     bool                          `plist:"SupportsLOMDevice"`
	DeviceID                              string                        `plist:"DeviceID"`
	EASDeviceIdentifier                   string                        `plist:"EASDeviceIdentifier"`
	IsCloudBackupEnabled                  bool                          `plist:"IsCloudBackupEnabled"`
	ActiveManagedUsers                    []string                      `plist:"ActiveManagedUsers"`
	OSUpdateSettings                      OSUpdateSettings              `plist:"OSUpdateSettings"`
	LocalHostName                         string                        `plist:"LocalHostName"`
	HostName                              string                        `plist:"HostName"`
	AutoSetupAdminAccounts                []AutoSetupAdminAccountsItem  `plist:"AutoSetupAdminAccounts"`
	SystemIntegrityProtectionEnabled      bool                          `plist:"SystemIntegrityProtectionEnabled"`
	IsMDMLostModeEnabled                  bool                          `plist:"IsMDMLostModeEnabled"`
	MaximumResidentUsers                  int                           `plist:"MaximumResidentUsers"`
	EstimatedResidentUsers                int                           `plist:"EstimatedResidentUsers"`
	QuotaSize                             int                           `plist:"QuotaSize"`
	ResidentUsers                         int                           `plist:"ResidentUsers"`
	UserSessionTimeout                    int                           `plist:"UserSessionTimeout"`
	TemporarySessionTimeout               int                           `plist:"TemporarySessionTimeout"`
	TemporarySessionOnly                  bool                          `plist:"TemporarySessionOnly"`
	ManagedAppleIDDefaultDomains          []string                      `plist:"ManagedAppleIDDefaultDomains"`
	OnlineAuthenticationGracePeriod       int                           `plist:"OnlineAuthenticationGracePeriod"`
	SkipLanguageAndLocaleSetupForNewUsers bool                          `plist:"SkipLanguageAndLocaleSetupForNewUsers"`
	PushToken                             []byte                        `plist:"PushToken"`
	DiagnosticSubmissionEnabled           bool                          `plist:"DiagnosticSubmissionEnabled"`
	AppAnalyticsEnabled                   bool                          `plist:"AppAnalyticsEnabled"`
	TimeZone                              string                        `plist:"TimeZone"`
	ICCID                                 string                        `plist:"ICCID"`
	BluetoothMAC                          string                        `plist:"BluetoothMAC"`
	WiFiMAC                               string                        `plist:"WiFiMAC"`
	EthernetMAC                           string                        `plist:"EthernetMAC"`
	CurrentCarrierNetwork                 string                        `plist:"CurrentCarrierNetwork"`
	SIMCarrierNetwork                     string                        `plist:"SIMCarrierNetwork"`
	SubscriberCarrierNetwork              string                        `plist:"SubscriberCarrierNetwork"`
	CarrierSettingsVersion                string                        `plist:"CarrierSettingsVersion"`
	PhoneNumber                           string                        `plist:"PhoneNumber"`
	DataRoamingEnabled                    bool                          `plist:"DataRoamingEnabled"`
	VoiceRoamingEnabled                   bool                          `plist:"VoiceRoamingEnabled"`
	PersonalHotspotEnabled                bool                          `plist:"PersonalHotspotEnabled"`
	IsNetworkTethered                     bool                          `plist:"IsNetworkTethered"`
	IsRoaming                             bool                          `plist:"IsRoaming"`
	SIMMCC                                string                        `plist:"SIMMCC"`
	SIMMNC                                string                        `plist:"SIMMNC"`
	SubscriberMCC                         string                        `plist:"SubscriberMCC"`
	SubscriberMNC                         string                        `plist:"SubscriberMNC"`
	CurrentMCC                            string                        `plist:"CurrentMCC"`
	CurrentMNC                            string                        `plist:"CurrentMNC"`
	ServiceSubscriptions                  []ServiceSubscriptionProperty `plist:"ServiceSubscriptions"`
	PINRequiredForEraseDevice             bool                          `plist:"PINRequiredForEraseDevice"`
	PINRequiredForDeviceLock              bool                          `plist:"PINRequiredForDeviceLock"`
	SupportsiOSAppInstalls                bool                          `plist:"SupportsiOSAppInstalls"`
	SoftwareUpdateDeviceID                string                        `plist:"SoftwareUpdateDeviceID"`
	SoftwareUpdateSettings                SoftwareUpdateSettings        `plist:"SoftwareUpdateSettings"`
	AccessibilitySettings                 AccessibilitySettings         `plist:"AccessibilitySettings"`
	DevicePropertiesAttestation           [][]byte                      `plist:"DevicePropertiesAttestation"`
}
// DeviceInformationResponse
type DeviceInformationResponse struct {
	QueryResponses QueryResponses `plist:"QueryResponses"`
	CommandUUID    string         `plist:"CommandUUID"`
	Status         string         `plist:"Status"`
}
