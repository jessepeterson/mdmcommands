// Code generated by "admgencmd"; DO NOT EDIT.
// Sources: application.remove.yaml, application.install.yaml, application.install.enterprise.yaml, application.extensions.mappings.yaml, application.validate.yaml, application.installed.list.yaml, application.redemptioncode.yaml, application.invitetoprogram.yaml, managed.application.configuration.yaml, managed.application.feedback.yaml
// Options: no-shared=true
package mdmcommands

const RemoveApplicationRequestType = "RemoveApplication"

// RemoveApplicationPayload is the "inner" command-specific payload for the "RemoveApplication" Apple MDM command.
type RemoveApplicationPayload struct {
	Identifier                   string
	RequestType                  string // supported value: RemoveApplication
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// RemoveApplicationCommand is the top-level structure for the "RemoveApplication" Apple MDM command.
type RemoveApplicationCommand struct {
	Command     RemoveApplicationPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *RemoveApplicationCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewRemoveApplicationCommand creates a new "RemoveApplication" Apple MDM command.
func NewRemoveApplicationCommand() *RemoveApplicationCommand {
	return &RemoveApplicationCommand{Command: RemoveApplicationPayload{RequestType: RemoveApplicationRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[RemoveApplicationRequestType] = func() interface{} {
		return NewRemoveApplicationCommand()
	}
}

// RemoveApplicationResponse is the command result report (response) for the "RemoveApplication" Apple MDM command.
type RemoveApplicationResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *RemoveApplicationResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[RemoveApplicationRequestType] = func() interface{} {
		return new(RemoveApplicationResponse)
	}
}

const InstallApplicationRequestType = "InstallApplication"

type Options struct {
	PurchaseMethod *int `plist:",omitempty"` // supported values: 0, 1
}
type Attributes struct {
	VPNUUID                                *string   `plist:",omitempty"`
	ContentFilterUUID                      *string   `plist:",omitempty"`
	DNSProxyUUID                           *string   `plist:",omitempty"`
	AssociatedDomains                      *[]string `plist:",omitempty"`
	AssociatedDomainsEnableDirectDownloads *bool     `plist:",omitempty"`
	Removable                              *bool     `plist:",omitempty"`
	TapToPayScreenLock                     *bool     `plist:",omitempty"`
}

// InstallApplicationPayload is the "inner" command-specific payload for the "InstallApplication" Apple MDM command.
type InstallApplicationPayload struct {
	ITunesStoreID                *int        `plist:"iTunesStoreID,omitempty"`
	Identifier                   *string     `plist:",omitempty"`
	Options                      *Options    `plist:",omitempty"`
	ManifestURL                  *string     `plist:",omitempty"`
	ManagementFlags              *int        `plist:",omitempty"` // supported values: 1, 4, 5
	Configuration                interface{} `plist:",omitempty"` // <any> type as single dictionary subkey
	Attributes                   *Attributes `plist:",omitempty"`
	ChangeManagementState        *string     `plist:",omitempty"` // supported value: Managed
	InstallAsManaged             *bool       `plist:",omitempty"`
	IOSApp                       *bool       `plist:"iOSApp,omitempty"`
	RequestType                  string      // supported value: InstallApplication
	RequestRequiresNetworkTether *bool       `plist:",omitempty"`
}

// InstallApplicationCommand is the top-level structure for the "InstallApplication" Apple MDM command.
type InstallApplicationCommand struct {
	Command     InstallApplicationPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *InstallApplicationCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewInstallApplicationCommand creates a new "InstallApplication" Apple MDM command.
func NewInstallApplicationCommand() *InstallApplicationCommand {
	return &InstallApplicationCommand{Command: InstallApplicationPayload{RequestType: InstallApplicationRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[InstallApplicationRequestType] = func() interface{} {
		return NewInstallApplicationCommand()
	}
}

// InstallApplicationResponse is the command result report (response) for the "InstallApplication" Apple MDM command.
type InstallApplicationResponse struct {
	Identifier      *string `plist:",omitempty"`
	State           *string `plist:",omitempty"`
	RejectionReason *string `plist:",omitempty"` // supported values: AppAlreadyInstalled, AppAlreadyQueued, AppStoreDisabled, CouldNotVerifyAppID, ManagementChangeNotSupported, NotAnApp, NotSupported, PurchaseMethodNotSupported, PurchaseMethodNotSupportedInMultiUser
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *InstallApplicationResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[InstallApplicationRequestType] = func() interface{} {
		return new(InstallApplicationResponse)
	}
}

const InstallEnterpriseApplicationRequestType = "InstallEnterpriseApplication"

// InstallEnterpriseApplicationPayload is the "inner" command-specific payload for the "InstallEnterpriseApplication" Apple MDM command.
type InstallEnterpriseApplicationPayload struct {
	Manifest                       interface{} `plist:",omitempty"` // <any> type as single dictionary subkey
	ManifestURL                    *string     `plist:",omitempty"`
	ManifestURLPinningCerts        *[][]byte   `plist:",omitempty"`
	PinningRevocationCheckRequired *bool       `plist:",omitempty"`
	InstallAsManaged               *bool       `plist:",omitempty"`
	ManagementFlags                *int        `plist:",omitempty"` // supported value: 1
	Configuration                  interface{} `plist:",omitempty"` // <any> type as single dictionary subkey
	ChangeManagementState          *string     `plist:",omitempty"` // supported value: Managed
	IOSApp                         *bool       `plist:"iOSApp,omitempty"`
	RequestType                    string      // supported value: InstallEnterpriseApplication
	RequestRequiresNetworkTether   *bool       `plist:",omitempty"`
}

// InstallEnterpriseApplicationCommand is the top-level structure for the "InstallEnterpriseApplication" Apple MDM command.
type InstallEnterpriseApplicationCommand struct {
	Command     InstallEnterpriseApplicationPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *InstallEnterpriseApplicationCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewInstallEnterpriseApplicationCommand creates a new "InstallEnterpriseApplication" Apple MDM command.
func NewInstallEnterpriseApplicationCommand() *InstallEnterpriseApplicationCommand {
	return &InstallEnterpriseApplicationCommand{Command: InstallEnterpriseApplicationPayload{RequestType: InstallEnterpriseApplicationRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[InstallEnterpriseApplicationRequestType] = func() interface{} {
		return NewInstallEnterpriseApplicationCommand()
	}
}

// InstallEnterpriseApplicationResponse is the command result report (response) for the "InstallEnterpriseApplication" Apple MDM command.
type InstallEnterpriseApplicationResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *InstallEnterpriseApplicationResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[InstallEnterpriseApplicationRequestType] = func() interface{} {
		return new(InstallEnterpriseApplicationResponse)
	}
}

const NSExtensionMappingsRequestType = "NSExtensionMappings"

// NSExtensionMappingsCommand is the top-level structure for the "NSExtensionMappings" Apple MDM command.
type NSExtensionMappingsCommand struct {
	Command     GenericCommandPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *NSExtensionMappingsCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewNSExtensionMappingsCommand creates a new "NSExtensionMappings" Apple MDM command.
func NewNSExtensionMappingsCommand() *NSExtensionMappingsCommand {
	return &NSExtensionMappingsCommand{Command: GenericCommandPayload{RequestType: NSExtensionMappingsRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[NSExtensionMappingsRequestType] = func() interface{} {
		return NewNSExtensionMappingsCommand()
	}
}

type ExtensionsItem struct {
	Identifier     string
	ExtensionPoint string
	DisplayName    string
}

// NSExtensionMappingsResponse is the command result report (response) for the "NSExtensionMappings" Apple MDM command.
type NSExtensionMappingsResponse struct {
	Extensions []ExtensionsItem
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *NSExtensionMappingsResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[NSExtensionMappingsRequestType] = func() interface{} {
		return new(NSExtensionMappingsResponse)
	}
}

const ValidateApplicationsRequestType = "ValidateApplications"

// ValidateApplicationsPayload is the "inner" command-specific payload for the "ValidateApplications" Apple MDM command.
type ValidateApplicationsPayload struct {
	Identifiers                  *[]string `plist:",omitempty"`
	RequestType                  string    // supported value: ValidateApplications
	RequestRequiresNetworkTether *bool     `plist:",omitempty"`
}

// ValidateApplicationsCommand is the top-level structure for the "ValidateApplications" Apple MDM command.
type ValidateApplicationsCommand struct {
	Command     ValidateApplicationsPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *ValidateApplicationsCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewValidateApplicationsCommand creates a new "ValidateApplications" Apple MDM command.
func NewValidateApplicationsCommand() *ValidateApplicationsCommand {
	return &ValidateApplicationsCommand{Command: ValidateApplicationsPayload{RequestType: ValidateApplicationsRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[ValidateApplicationsRequestType] = func() interface{} {
		return NewValidateApplicationsCommand()
	}
}

// ValidateApplicationsResponse is the command result report (response) for the "ValidateApplications" Apple MDM command.
type ValidateApplicationsResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *ValidateApplicationsResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[ValidateApplicationsRequestType] = func() interface{} {
		return new(ValidateApplicationsResponse)
	}
}

const InstalledApplicationListRequestType = "InstalledApplicationList"

// InstalledApplicationListPayload is the "inner" command-specific payload for the "InstalledApplicationList" Apple MDM command.
type InstalledApplicationListPayload struct {
	Identifiers                  *[]string `plist:",omitempty"`
	ManagedAppsOnly              *bool     `plist:",omitempty"`
	Items                        *[]string `plist:",omitempty"` // supported values: AdHocCodeSigned, AppStoreVendable, BetaApp, BundleSize, DeviceBasedVPP, DynamicSize, ExternalVersionIdentifier, HasUpdateAvailable, Identifier, Installing, IsAppClip, IsValidated, Name, ShortVersion, Version
	RequestType                  string    // supported value: InstalledApplicationList
	RequestRequiresNetworkTether *bool     `plist:",omitempty"`
}

// InstalledApplicationListCommand is the top-level structure for the "InstalledApplicationList" Apple MDM command.
type InstalledApplicationListCommand struct {
	Command     InstalledApplicationListPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *InstalledApplicationListCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewInstalledApplicationListCommand creates a new "InstalledApplicationList" Apple MDM command.
func NewInstalledApplicationListCommand() *InstalledApplicationListCommand {
	return &InstalledApplicationListCommand{Command: InstalledApplicationListPayload{RequestType: InstalledApplicationListRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[InstalledApplicationListRequestType] = func() interface{} {
		return NewInstalledApplicationListCommand()
	}
}

type InstalledApplicationListItem struct {
	Identifier                *string `plist:",omitempty"`
	ExternalVersionIdentifier *int    `plist:",omitempty"`
	Version                   *string `plist:",omitempty"`
	ShortVersion              *string `plist:",omitempty"`
	Name                      *string `plist:",omitempty"`
	BundleSize                *int    `plist:",omitempty"`
	DynamicSize               *int    `plist:",omitempty"`
	IsValidated               *bool   `plist:",omitempty"`
	Installing                *bool   `plist:",omitempty"`
	AppStoreVendable          *bool   `plist:",omitempty"`
	DeviceBasedVPP            *bool   `plist:",omitempty"`
	BetaApp                   *bool   `plist:",omitempty"`
	AdHocCodeSigned           *bool   `plist:",omitempty"`
	HasUpdateAvailable        *bool   `plist:",omitempty"`
	DownloadFailed            *bool   `plist:",omitempty"`
	DownloadWaiting           *bool   `plist:",omitempty"`
	DownloadPaused            *bool   `plist:",omitempty"`
	DownloadCancelled         *bool   `plist:",omitempty"`
	IsAppClip                 *bool   `plist:",omitempty"`
}

// InstalledApplicationListResponse is the command result report (response) for the "InstalledApplicationList" Apple MDM command.
type InstalledApplicationListResponse struct {
	InstalledApplicationList []InstalledApplicationListItem
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *InstalledApplicationListResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[InstalledApplicationListRequestType] = func() interface{} {
		return new(InstalledApplicationListResponse)
	}
}

const ApplyRedemptionCodeRequestType = "ApplyRedemptionCode"

// ApplyRedemptionCodePayload is the "inner" command-specific payload for the "ApplyRedemptionCode" Apple MDM command.
type ApplyRedemptionCodePayload struct {
	Identifier                   string
	RedemptionCode               string
	RequestType                  string // supported value: ApplyRedemptionCode
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// ApplyRedemptionCodeCommand is the top-level structure for the "ApplyRedemptionCode" Apple MDM command.
type ApplyRedemptionCodeCommand struct {
	Command     ApplyRedemptionCodePayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *ApplyRedemptionCodeCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewApplyRedemptionCodeCommand creates a new "ApplyRedemptionCode" Apple MDM command.
func NewApplyRedemptionCodeCommand() *ApplyRedemptionCodeCommand {
	return &ApplyRedemptionCodeCommand{Command: ApplyRedemptionCodePayload{RequestType: ApplyRedemptionCodeRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[ApplyRedemptionCodeRequestType] = func() interface{} {
		return NewApplyRedemptionCodeCommand()
	}
}

// ApplyRedemptionCodeResponse is the command result report (response) for the "ApplyRedemptionCode" Apple MDM command.
type ApplyRedemptionCodeResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *ApplyRedemptionCodeResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[ApplyRedemptionCodeRequestType] = func() interface{} {
		return new(ApplyRedemptionCodeResponse)
	}
}

const InviteToProgramRequestType = "InviteToProgram"

// InviteToProgramPayload is the "inner" command-specific payload for the "InviteToProgram" Apple MDM command.
type InviteToProgramPayload struct {
	ProgramID                    string // supported value: com.apple.cloudvpp
	InvitationURL                string
	RequestType                  string // supported value: InviteToProgram
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// InviteToProgramCommand is the top-level structure for the "InviteToProgram" Apple MDM command.
type InviteToProgramCommand struct {
	Command     InviteToProgramPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *InviteToProgramCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewInviteToProgramCommand creates a new "InviteToProgram" Apple MDM command.
func NewInviteToProgramCommand() *InviteToProgramCommand {
	return &InviteToProgramCommand{Command: InviteToProgramPayload{RequestType: InviteToProgramRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[InviteToProgramRequestType] = func() interface{} {
		return NewInviteToProgramCommand()
	}
}

// InviteToProgramResponse is the command result report (response) for the "InviteToProgram" Apple MDM command.
type InviteToProgramResponse struct {
	InvitationResult string // supported values: Acknowledged, InvalidProgramID, InvalidInvitationURL
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *InviteToProgramResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[InviteToProgramRequestType] = func() interface{} {
		return new(InviteToProgramResponse)
	}
}

const ManagedApplicationConfigurationRequestType = "ManagedApplicationConfiguration"

// ManagedApplicationConfigurationPayload is the "inner" command-specific payload for the "ManagedApplicationConfiguration" Apple MDM command.
type ManagedApplicationConfigurationPayload struct {
	Identifiers                  []string
	RequestType                  string // supported value: ManagedApplicationConfiguration
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// ManagedApplicationConfigurationCommand is the top-level structure for the "ManagedApplicationConfiguration" Apple MDM command.
type ManagedApplicationConfigurationCommand struct {
	Command     ManagedApplicationConfigurationPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *ManagedApplicationConfigurationCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewManagedApplicationConfigurationCommand creates a new "ManagedApplicationConfiguration" Apple MDM command.
func NewManagedApplicationConfigurationCommand() *ManagedApplicationConfigurationCommand {
	return &ManagedApplicationConfigurationCommand{Command: ManagedApplicationConfigurationPayload{RequestType: ManagedApplicationConfigurationRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[ManagedApplicationConfigurationRequestType] = func() interface{} {
		return NewManagedApplicationConfigurationCommand()
	}
}

type ApplicationConfigurationsItem struct {
	Identifier    string
	Configuration interface{} `plist:",omitempty"` // <any> type as single dictionary subkey
}

// ManagedApplicationConfigurationResponse is the command result report (response) for the "ManagedApplicationConfiguration" Apple MDM command.
type ManagedApplicationConfigurationResponse struct {
	ApplicationConfigurations []ApplicationConfigurationsItem
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *ManagedApplicationConfigurationResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[ManagedApplicationConfigurationRequestType] = func() interface{} {
		return new(ManagedApplicationConfigurationResponse)
	}
}

const ManagedApplicationFeedbackRequestType = "ManagedApplicationFeedback"

// ManagedApplicationFeedbackPayload is the "inner" command-specific payload for the "ManagedApplicationFeedback" Apple MDM command.
type ManagedApplicationFeedbackPayload struct {
	Identifiers                  []string
	DeleteFeedback               *bool  `plist:",omitempty"`
	RequestType                  string // supported value: ManagedApplicationFeedback
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// ManagedApplicationFeedbackCommand is the top-level structure for the "ManagedApplicationFeedback" Apple MDM command.
type ManagedApplicationFeedbackCommand struct {
	Command     ManagedApplicationFeedbackPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *ManagedApplicationFeedbackCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewManagedApplicationFeedbackCommand creates a new "ManagedApplicationFeedback" Apple MDM command.
func NewManagedApplicationFeedbackCommand() *ManagedApplicationFeedbackCommand {
	return &ManagedApplicationFeedbackCommand{Command: ManagedApplicationFeedbackPayload{RequestType: ManagedApplicationFeedbackRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[ManagedApplicationFeedbackRequestType] = func() interface{} {
		return NewManagedApplicationFeedbackCommand()
	}
}

type ManagedApplicationFeedbackItem struct {
	Identifier string
	Feedback   interface{} `plist:",omitempty"` // <any> type as single dictionary subkey
}

// ManagedApplicationFeedbackResponse is the command result report (response) for the "ManagedApplicationFeedback" Apple MDM command.
type ManagedApplicationFeedbackResponse struct {
	ManagedApplicationFeedback []ManagedApplicationFeedbackItem
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *ManagedApplicationFeedbackResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[ManagedApplicationFeedbackRequestType] = func() interface{} {
		return new(ManagedApplicationFeedbackResponse)
	}
}
