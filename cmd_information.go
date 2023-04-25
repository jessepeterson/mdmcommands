// Code generated by "admgencmd"; DO NOT EDIT.
// Sources: information.device.yaml, information.security.yaml, certificate.list.yaml
// Options: no-shared=true
package mdmcommands

import "time"

const DeviceInformationRequestType = "DeviceInformation"

// DeviceInformationPayload is the "inner" command-specific payload for the "DeviceInformation" Apple MDM command.
type DeviceInformationPayload struct {
	Queries                      []string // 85 array value(s) defined in schema
	DeviceAttestationNonce       *[]byte  `plist:",omitempty"`
	RequestType                  string   // must be set to "DeviceInformation"
	RequestRequiresNetworkTether *bool    `plist:",omitempty"`
}

// DeviceInformationCommand is the top-level structure for the "DeviceInformation" Apple MDM command.
type DeviceInformationCommand struct {
	Command     DeviceInformationPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *DeviceInformationCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewDeviceInformationCommand creates a new "DeviceInformation" Apple MDM command.
func NewDeviceInformationCommand() *DeviceInformationCommand {
	return &DeviceInformationCommand{Command: DeviceInformationPayload{RequestType: DeviceInformationRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[DeviceInformationRequestType] = func() interface{} {
		return NewDeviceInformationCommand()
	}
}

type OrganizationInfo struct {
	OrganizationName    string
	OrganizationAddress *string `plist:",omitempty"`
	OrganizationPhone   *string `plist:",omitempty"`
	OrganizationEmail   *string `plist:",omitempty"`
	OrganizationMagic   *string `plist:",omitempty"`
}
type MDMOptions struct {
	ActivationLockAllowedWhileSupervised             *bool `plist:",omitempty"`
	BootstrapTokenAllowed                            *bool `plist:",omitempty"`
	PromptUserToAllowBootstrapTokenForAuthentication *bool `plist:",omitempty"`
}
type OSUpdateSettings struct {
	CatalogURL                      *string
	IsDefaultCatalog                *bool
	PreviousScanDate                *time.Time
	PreviousScanResult              *string
	PerformPeriodicCheck            *bool
	AutoCheckEnabled                *bool
	BackgroundDownloadEnabled       *bool
	AutomaticAppInstallationEnabled *bool
	AutomaticOSInstallationEnabled  *bool
	AutomaticSecurityUpdatesEnabled *bool
}
type AutoSetupAdminAccountsItem struct {
	GUID      *string
	ShortName *string `plist:"shortName"`
}
type ServiceSubscriptionProperty struct {
	CarrierSettingsVersion   *string
	CurrentCarrierNetwork    *string
	CurrentMCC               *string
	CurrentMNC               *string
	ICCID                    *string
	EID                      *string
	IMEI                     *string
	IsDataPreferred          *bool
	IsRoaming                *bool
	IsVoicePreferred         *bool
	Label                    *string
	LabelID                  *string
	MEID                     *string
	PhoneNumber              *string
	Slot                     *string
	SubscriberCarrierNetwork *string
}
type SoftwareUpdateSettings struct {
	RecommendationsCadence *int
}
type AccessibilitySettings struct {
	BoldTextEnabled            *bool
	IncreaseContrastEnabled    *bool
	ReduceMotionEnabled        *bool
	ReduceTransparencyEnabled  *bool
	TextSize                   *int // possible values: -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
	TouchAccommodationsEnabled *bool
	VoiceOverEnabled           *bool
	ZoomEnabled                *bool
}
type QueryResponses struct {
	UDID                                  *string
	ProvisioningUDID                      *string
	OrganizationInfo                      *OrganizationInfo
	MDMOptions                            *MDMOptions
	LastCloudBackupDate                   *time.Time
	AwaitingConfiguration                 *bool
	ITunesStoreAccountIsActive            *bool   `plist:"iTunesStoreAccountIsActive"`
	ITunesStoreAccountHash                *string `plist:"iTunesStoreAccountHash"`
	DeviceName                            *string
	OSVersion                             *string
	SupplementalOSVersionExtra            *string
	BuildVersion                          *string
	SupplementalBuildVersion              *string
	ModelName                             *string
	Model                                 *string
	ModelNumber                           *string
	IsAppleSilicon                        *bool
	ProductName                           *string
	SerialNumber                          *string
	DeviceCapacity                        *float64
	AvailableDeviceCapacity               *float64
	IMEI                                  *string
	MEID                                  *string
	ModemFirmwareVersion                  *string
	CellularTechnology                    *int // possible values: 0, 1, 2, 3
	BatteryLevel                          *float64
	HasBattery                            *bool
	IsSupervised                          *bool
	IsMultiUser                           *bool
	IsDeviceLocatorServiceEnabled         *bool
	IsActivationLockEnabled               *bool
	IsActivationLockSupported             *bool
	IsDoNotDisturbInEffect                *bool
	SupportsLOMDevice                     *bool
	DeviceID                              *string
	EASDeviceIdentifier                   *string
	IsCloudBackupEnabled                  *bool
	ActiveManagedUsers                    *[]string
	OSUpdateSettings                      *OSUpdateSettings
	LocalHostName                         *string
	HostName                              *string
	AutoSetupAdminAccounts                *[]AutoSetupAdminAccountsItem
	SystemIntegrityProtectionEnabled      *bool
	IsMDMLostModeEnabled                  *bool
	MaximumResidentUsers                  *int
	EstimatedResidentUsers                *int
	QuotaSize                             *int
	ResidentUsers                         *int
	UserSessionTimeout                    *int
	TemporarySessionTimeout               *int
	TemporarySessionOnly                  *bool
	ManagedAppleIDDefaultDomains          *[]string
	OnlineAuthenticationGracePeriod       *int
	SkipLanguageAndLocaleSetupForNewUsers *bool
	PushToken                             *[]byte
	DiagnosticSubmissionEnabled           *bool
	AppAnalyticsEnabled                   *bool
	TimeZone                              *string
	ICCID                                 *string
	BluetoothMAC                          *string
	WiFiMAC                               *string
	EthernetMAC                           *string
	CurrentCarrierNetwork                 *string
	SIMCarrierNetwork                     *string
	SubscriberCarrierNetwork              *string
	CarrierSettingsVersion                *string
	PhoneNumber                           *string
	DataRoamingEnabled                    *bool
	VoiceRoamingEnabled                   *bool
	PersonalHotspotEnabled                *bool
	IsNetworkTethered                     *bool
	IsRoaming                             *bool
	SIMMCC                                *string
	SIMMNC                                *string
	SubscriberMCC                         *string
	SubscriberMNC                         *string
	CurrentMCC                            *string
	CurrentMNC                            *string
	ServiceSubscriptions                  *[]ServiceSubscriptionProperty
	PINRequiredForEraseDevice             *bool
	PINRequiredForDeviceLock              *bool
	SupportsiOSAppInstalls                *bool
	SoftwareUpdateDeviceID                *string
	SoftwareUpdateSettings                *SoftwareUpdateSettings
	AccessibilitySettings                 *AccessibilitySettings
	DevicePropertiesAttestation           *[][]byte
	EACSPreflight                         *string
}

// DeviceInformationResponse is the command result report (response) for the "DeviceInformation" Apple MDM command.
type DeviceInformationResponse struct {
	QueryResponses QueryResponses
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *DeviceInformationResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[DeviceInformationRequestType] = func() interface{} {
		return new(DeviceInformationResponse)
	}
}

const SecurityInfoRequestType = "SecurityInfo"

// SecurityInfoCommand is the top-level structure for the "SecurityInfo" Apple MDM command.
type SecurityInfoCommand struct {
	Command     GenericCommandPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *SecurityInfoCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewSecurityInfoCommand creates a new "SecurityInfo" Apple MDM command.
func NewSecurityInfoCommand() *SecurityInfoCommand {
	return &SecurityInfoCommand{Command: GenericCommandPayload{RequestType: SecurityInfoRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[SecurityInfoRequestType] = func() interface{} {
		return NewSecurityInfoCommand()
	}
}

type ApplicationsItem struct {
	Allowed  *bool
	BundleID *string
	Name     *string
}
type FirewallSettings struct {
	FirewallEnabled  *bool
	BlockAllIncoming *bool
	StealthMode      *bool
	Applications     *[]ApplicationsItem
	LoggingEnabled   *bool
	LoggingOption    *string // possible values: throttled, brief, detail
}
type FirmwarePasswordStatus struct {
	PasswordExists *bool
	ChangePending  *bool
	AllowOroms     *bool
}
type ManagementStatus struct {
	EnrolledViaDEP             *bool
	UserApprovedEnrollment     *bool
	IsUserEnrollment           *bool
	IsActivationLockManageable *bool
}
type SecureBoot struct {
	SecureBootLevel   *string   // possible values: off, medium, full, not supported
	ExternalBootLevel *string   // possible values: allowed, disallowed, not supported
	ReducedSecurity   *[]string // 3 array value(s) defined in schema
}
type SecurityInfo struct {
	HardwareEncryptionCaps                           *int
	PasscodePresent                                  *bool
	PasscodeCompliant                                *bool
	PasscodeCompliantWithProfiles                    *bool
	PasscodeLockGracePeriod                          *int
	PasscodeLockGracePeriodEnforced                  *int
	FDEEnabled                                       *bool   `plist:"FDE_Enabled"`
	FDEHasPersonalRecoveryKey                        *bool   `plist:"FDE_HasPersonalRecoveryKey"`
	FDEHasInstitutionalRecoveryKey                   *bool   `plist:"FDE_HasInstitutionalRecoveryKey"`
	FDEPersonalRecoveryKeyCMS                        *[]byte `plist:"FDE_PersonalRecoveryKeyCMS"`
	FDEPersonalRecoveryKeyDeviceKey                  *string `plist:"FDE_PersonalRecoveryKeyDeviceKey"`
	SystemIntegrityProtectionEnabled                 *bool
	FirewallSettings                                 *FirewallSettings
	FirmwarePasswordStatus                           *FirmwarePasswordStatus
	ManagementStatus                                 *ManagementStatus
	SecureBoot                                       *SecureBoot
	RemoteDesktopEnabled                             *bool
	AuthenticatedRootVolumeEnabled                   *bool
	BootstrapTokenAllowedForAuthentication           *string // possible values: allowed, disallowed, not supported
	BootstrapTokenRequiredForSoftwareUpdate          *bool
	BootstrapTokenRequiredForKernelExtensionApproval *bool
	IsRecoveryLockEnabled                            *bool
}

// SecurityInfoResponse is the command result report (response) for the "SecurityInfo" Apple MDM command.
type SecurityInfoResponse struct {
	SecurityInfo SecurityInfo
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *SecurityInfoResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[SecurityInfoRequestType] = func() interface{} {
		return new(SecurityInfoResponse)
	}
}

const CertificateListRequestType = "CertificateList"

// CertificateListPayload is the "inner" command-specific payload for the "CertificateList" Apple MDM command.
type CertificateListPayload struct {
	ManagedOnly                  *bool  `plist:",omitempty"`
	RequestType                  string // must be set to "CertificateList"
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// CertificateListCommand is the top-level structure for the "CertificateList" Apple MDM command.
type CertificateListCommand struct {
	Command     CertificateListPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *CertificateListCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewCertificateListCommand creates a new "CertificateList" Apple MDM command.
func NewCertificateListCommand() *CertificateListCommand {
	return &CertificateListCommand{Command: CertificateListPayload{RequestType: CertificateListRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[CertificateListRequestType] = func() interface{} {
		return NewCertificateListCommand()
	}
}

type CertificateListItem struct {
	CommonName string
	IsIdentity bool
	Data       []byte
}

// CertificateListResponse is the command result report (response) for the "CertificateList" Apple MDM command.
type CertificateListResponse struct {
	CertificateList []CertificateListItem
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *CertificateListResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[CertificateListRequestType] = func() interface{} {
		return new(CertificateListResponse)
	}
}
