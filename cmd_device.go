// Code generated by "admgencmd"; DO NOT EDIT.
// Sources: device.shutdown.yaml, device.restart.yaml, device.configured.yaml, device.erase.yaml, device.lostmode.enable.yaml, device.lostmode.disable.yaml, device.lostmode.location.yaml, device.lostmode.playsound.yaml, device.restrictions.clearpassword.yaml, device.lock.yaml, device.esim.yaml, device.activationlock.bypasscode.yaml, device.activationlock.clearbypasscode.yaml, declarativemanagement.yaml, rotate.file.vault.key.yaml
// Options: no-shared=true
package mdmcommands

const ShutDownDeviceRequestType = "ShutDownDevice"

// ShutDownDeviceCommand is the top-level structure for the "ShutDownDevice" Apple MDM command.
type ShutDownDeviceCommand struct {
	Command     GenericCommandPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *ShutDownDeviceCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewShutDownDeviceCommand creates a new "ShutDownDevice" Apple MDM command.
func NewShutDownDeviceCommand() *ShutDownDeviceCommand {
	return &ShutDownDeviceCommand{Command: GenericCommandPayload{RequestType: ShutDownDeviceRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[ShutDownDeviceRequestType] = func() interface{} {
		return NewShutDownDeviceCommand()
	}
}

// ShutDownDeviceResponse is the command result report (response) for the "ShutDownDevice" Apple MDM command.
type ShutDownDeviceResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *ShutDownDeviceResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[ShutDownDeviceRequestType] = func() interface{} {
		return new(ShutDownDeviceResponse)
	}
}

const RestartDeviceRequestType = "RestartDevice"

// RestartDevicePayload is the "inner" command-specific payload for the "RestartDevice" Apple MDM command.
type RestartDevicePayload struct {
	RebuildKernelCache           *bool     `plist:",omitempty"`
	KextPaths                    *[]string `plist:",omitempty"`
	NotifyUser                   *bool     `plist:",omitempty"`
	RequestType                  string    // must be set to "RestartDevice"
	RequestRequiresNetworkTether *bool     `plist:",omitempty"`
}

// RestartDeviceCommand is the top-level structure for the "RestartDevice" Apple MDM command.
type RestartDeviceCommand struct {
	Command     RestartDevicePayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *RestartDeviceCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewRestartDeviceCommand creates a new "RestartDevice" Apple MDM command.
func NewRestartDeviceCommand() *RestartDeviceCommand {
	return &RestartDeviceCommand{Command: RestartDevicePayload{RequestType: RestartDeviceRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[RestartDeviceRequestType] = func() interface{} {
		return NewRestartDeviceCommand()
	}
}

// RestartDeviceResponse is the command result report (response) for the "RestartDevice" Apple MDM command.
type RestartDeviceResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *RestartDeviceResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[RestartDeviceRequestType] = func() interface{} {
		return new(RestartDeviceResponse)
	}
}

const DeviceConfiguredRequestType = "DeviceConfigured"

// DeviceConfiguredCommand is the top-level structure for the "DeviceConfigured" Apple MDM command.
type DeviceConfiguredCommand struct {
	Command     GenericCommandPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *DeviceConfiguredCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewDeviceConfiguredCommand creates a new "DeviceConfigured" Apple MDM command.
func NewDeviceConfiguredCommand() *DeviceConfiguredCommand {
	return &DeviceConfiguredCommand{Command: GenericCommandPayload{RequestType: DeviceConfiguredRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[DeviceConfiguredRequestType] = func() interface{} {
		return NewDeviceConfiguredCommand()
	}
}

// DeviceConfiguredResponse is the command result report (response) for the "DeviceConfigured" Apple MDM command.
type DeviceConfiguredResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *DeviceConfiguredResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[DeviceConfiguredRequestType] = func() interface{} {
		return new(DeviceConfiguredResponse)
	}
}

const EraseDeviceRequestType = "EraseDevice"

// EraseDevicePayload is the "inner" command-specific payload for the "EraseDevice" Apple MDM command.
type EraseDevicePayload struct {
	PreserveDataPlan             *bool   `plist:",omitempty"`
	DisallowProximitySetup       *bool   `plist:",omitempty"`
	PIN                          *string `plist:",omitempty"`
	ObliterationBehavior         *string `plist:",omitempty"` // possible values: Default, DoNotObliterate, ObliterateWithWarning, Always
	RequestType                  string  // must be set to "EraseDevice"
	RequestRequiresNetworkTether *bool   `plist:",omitempty"`
}

// EraseDeviceCommand is the top-level structure for the "EraseDevice" Apple MDM command.
type EraseDeviceCommand struct {
	Command     EraseDevicePayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *EraseDeviceCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewEraseDeviceCommand creates a new "EraseDevice" Apple MDM command.
func NewEraseDeviceCommand() *EraseDeviceCommand {
	return &EraseDeviceCommand{Command: EraseDevicePayload{RequestType: EraseDeviceRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[EraseDeviceRequestType] = func() interface{} {
		return NewEraseDeviceCommand()
	}
}

// EraseDeviceResponse is the command result report (response) for the "EraseDevice" Apple MDM command.
type EraseDeviceResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *EraseDeviceResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[EraseDeviceRequestType] = func() interface{} {
		return new(EraseDeviceResponse)
	}
}

const EnableLostModeRequestType = "EnableLostMode"

// EnableLostModePayload is the "inner" command-specific payload for the "EnableLostMode" Apple MDM command.
type EnableLostModePayload struct {
	Message                      *string `plist:",omitempty"`
	PhoneNumber                  *string `plist:",omitempty"`
	Footnote                     *string `plist:",omitempty"`
	RequestType                  string  // must be set to "EnableLostMode"
	RequestRequiresNetworkTether *bool   `plist:",omitempty"`
}

// EnableLostModeCommand is the top-level structure for the "EnableLostMode" Apple MDM command.
type EnableLostModeCommand struct {
	Command     EnableLostModePayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *EnableLostModeCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewEnableLostModeCommand creates a new "EnableLostMode" Apple MDM command.
func NewEnableLostModeCommand() *EnableLostModeCommand {
	return &EnableLostModeCommand{Command: EnableLostModePayload{RequestType: EnableLostModeRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[EnableLostModeRequestType] = func() interface{} {
		return NewEnableLostModeCommand()
	}
}

// EnableLostModeResponse is the command result report (response) for the "EnableLostMode" Apple MDM command.
type EnableLostModeResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *EnableLostModeResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[EnableLostModeRequestType] = func() interface{} {
		return new(EnableLostModeResponse)
	}
}

const DisableLostModeRequestType = "DisableLostMode"

// DisableLostModeCommand is the top-level structure for the "DisableLostMode" Apple MDM command.
type DisableLostModeCommand struct {
	Command     GenericCommandPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *DisableLostModeCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewDisableLostModeCommand creates a new "DisableLostMode" Apple MDM command.
func NewDisableLostModeCommand() *DisableLostModeCommand {
	return &DisableLostModeCommand{Command: GenericCommandPayload{RequestType: DisableLostModeRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[DisableLostModeRequestType] = func() interface{} {
		return NewDisableLostModeCommand()
	}
}

// DisableLostModeResponse is the command result report (response) for the "DisableLostMode" Apple MDM command.
type DisableLostModeResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *DisableLostModeResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[DisableLostModeRequestType] = func() interface{} {
		return new(DisableLostModeResponse)
	}
}

const DeviceLocationRequestType = "DeviceLocation"

// DeviceLocationCommand is the top-level structure for the "DeviceLocation" Apple MDM command.
type DeviceLocationCommand struct {
	Command     GenericCommandPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *DeviceLocationCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewDeviceLocationCommand creates a new "DeviceLocation" Apple MDM command.
func NewDeviceLocationCommand() *DeviceLocationCommand {
	return &DeviceLocationCommand{Command: GenericCommandPayload{RequestType: DeviceLocationRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[DeviceLocationRequestType] = func() interface{} {
		return NewDeviceLocationCommand()
	}
}

// DeviceLocationResponse is the command result report (response) for the "DeviceLocation" Apple MDM command.
type DeviceLocationResponse struct {
	Latitude           float64
	Longitude          float64
	HorizontalAccuracy float64
	VerticalAccuracy   float64
	Altitude           float64
	Speed              float64
	Course             float64
	Timestamp          string
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *DeviceLocationResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[DeviceLocationRequestType] = func() interface{} {
		return new(DeviceLocationResponse)
	}
}

const PlayLostModeSoundRequestType = "PlayLostModeSound"

// PlayLostModeSoundCommand is the top-level structure for the "PlayLostModeSound" Apple MDM command.
type PlayLostModeSoundCommand struct {
	Command     GenericCommandPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *PlayLostModeSoundCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewPlayLostModeSoundCommand creates a new "PlayLostModeSound" Apple MDM command.
func NewPlayLostModeSoundCommand() *PlayLostModeSoundCommand {
	return &PlayLostModeSoundCommand{Command: GenericCommandPayload{RequestType: PlayLostModeSoundRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[PlayLostModeSoundRequestType] = func() interface{} {
		return NewPlayLostModeSoundCommand()
	}
}

// PlayLostModeSoundResponse is the command result report (response) for the "PlayLostModeSound" Apple MDM command.
type PlayLostModeSoundResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *PlayLostModeSoundResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[PlayLostModeSoundRequestType] = func() interface{} {
		return new(PlayLostModeSoundResponse)
	}
}

const ClearRestrictionsPasswordRequestType = "ClearRestrictionsPassword"

// ClearRestrictionsPasswordCommand is the top-level structure for the "ClearRestrictionsPassword" Apple MDM command.
type ClearRestrictionsPasswordCommand struct {
	Command     GenericCommandPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *ClearRestrictionsPasswordCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewClearRestrictionsPasswordCommand creates a new "ClearRestrictionsPassword" Apple MDM command.
func NewClearRestrictionsPasswordCommand() *ClearRestrictionsPasswordCommand {
	return &ClearRestrictionsPasswordCommand{Command: GenericCommandPayload{RequestType: ClearRestrictionsPasswordRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[ClearRestrictionsPasswordRequestType] = func() interface{} {
		return NewClearRestrictionsPasswordCommand()
	}
}

// ClearRestrictionsPasswordResponse is the command result report (response) for the "ClearRestrictionsPassword" Apple MDM command.
type ClearRestrictionsPasswordResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *ClearRestrictionsPasswordResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[ClearRestrictionsPasswordRequestType] = func() interface{} {
		return new(ClearRestrictionsPasswordResponse)
	}
}

const DeviceLockRequestType = "DeviceLock"

// DeviceLockPayload is the "inner" command-specific payload for the "DeviceLock" Apple MDM command.
type DeviceLockPayload struct {
	Message                      *string `plist:",omitempty"`
	PhoneNumber                  *string `plist:",omitempty"`
	PIN                          *string `plist:",omitempty"`
	RequestType                  string  // must be set to "DeviceLock"
	RequestRequiresNetworkTether *bool   `plist:",omitempty"`
}

// DeviceLockCommand is the top-level structure for the "DeviceLock" Apple MDM command.
type DeviceLockCommand struct {
	Command     DeviceLockPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *DeviceLockCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewDeviceLockCommand creates a new "DeviceLock" Apple MDM command.
func NewDeviceLockCommand() *DeviceLockCommand {
	return &DeviceLockCommand{Command: DeviceLockPayload{RequestType: DeviceLockRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[DeviceLockRequestType] = func() interface{} {
		return NewDeviceLockCommand()
	}
}

// DeviceLockResponse is the command result report (response) for the "DeviceLock" Apple MDM command.
type DeviceLockResponse struct {
	MessageResult *string `plist:",omitempty"`
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *DeviceLockResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[DeviceLockRequestType] = func() interface{} {
		return new(DeviceLockResponse)
	}
}

const RefreshCellularPlansRequestType = "RefreshCellularPlans"

// RefreshCellularPlansPayload is the "inner" command-specific payload for the "RefreshCellularPlans" Apple MDM command.
type RefreshCellularPlansPayload struct {
	ESIMServerURL                string `plist:"eSIMServerURL"`
	RequestType                  string // must be set to "RefreshCellularPlans"
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// RefreshCellularPlansCommand is the top-level structure for the "RefreshCellularPlans" Apple MDM command.
type RefreshCellularPlansCommand struct {
	Command     RefreshCellularPlansPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *RefreshCellularPlansCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewRefreshCellularPlansCommand creates a new "RefreshCellularPlans" Apple MDM command.
func NewRefreshCellularPlansCommand() *RefreshCellularPlansCommand {
	return &RefreshCellularPlansCommand{Command: RefreshCellularPlansPayload{RequestType: RefreshCellularPlansRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[RefreshCellularPlansRequestType] = func() interface{} {
		return NewRefreshCellularPlansCommand()
	}
}

// RefreshCellularPlansResponse is the command result report (response) for the "RefreshCellularPlans" Apple MDM command.
type RefreshCellularPlansResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *RefreshCellularPlansResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[RefreshCellularPlansRequestType] = func() interface{} {
		return new(RefreshCellularPlansResponse)
	}
}

const ActivationLockBypassCodeRequestType = "ActivationLockBypassCode"

// ActivationLockBypassCodeCommand is the top-level structure for the "ActivationLockBypassCode" Apple MDM command.
type ActivationLockBypassCodeCommand struct {
	Command     GenericCommandPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *ActivationLockBypassCodeCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewActivationLockBypassCodeCommand creates a new "ActivationLockBypassCode" Apple MDM command.
func NewActivationLockBypassCodeCommand() *ActivationLockBypassCodeCommand {
	return &ActivationLockBypassCodeCommand{Command: GenericCommandPayload{RequestType: ActivationLockBypassCodeRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[ActivationLockBypassCodeRequestType] = func() interface{} {
		return NewActivationLockBypassCodeCommand()
	}
}

// ActivationLockBypassCodeResponse is the command result report (response) for the "ActivationLockBypassCode" Apple MDM command.
type ActivationLockBypassCodeResponse struct {
	ActivationLockBypassCode string
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *ActivationLockBypassCodeResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[ActivationLockBypassCodeRequestType] = func() interface{} {
		return new(ActivationLockBypassCodeResponse)
	}
}

const ClearActivationLockBypassCodeRequestType = "ClearActivationLockBypassCode"

// ClearActivationLockBypassCodeCommand is the top-level structure for the "ClearActivationLockBypassCode" Apple MDM command.
type ClearActivationLockBypassCodeCommand struct {
	Command     GenericCommandPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *ClearActivationLockBypassCodeCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewClearActivationLockBypassCodeCommand creates a new "ClearActivationLockBypassCode" Apple MDM command.
func NewClearActivationLockBypassCodeCommand() *ClearActivationLockBypassCodeCommand {
	return &ClearActivationLockBypassCodeCommand{Command: GenericCommandPayload{RequestType: ClearActivationLockBypassCodeRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[ClearActivationLockBypassCodeRequestType] = func() interface{} {
		return NewClearActivationLockBypassCodeCommand()
	}
}

// ClearActivationLockBypassCodeResponse is the command result report (response) for the "ClearActivationLockBypassCode" Apple MDM command.
type ClearActivationLockBypassCodeResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *ClearActivationLockBypassCodeResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[ClearActivationLockBypassCodeRequestType] = func() interface{} {
		return new(ClearActivationLockBypassCodeResponse)
	}
}

const DeclarativeManagementRequestType = "DeclarativeManagement"

// DeclarativeManagementPayload is the "inner" command-specific payload for the "DeclarativeManagement" Apple MDM command.
type DeclarativeManagementPayload struct {
	Data                         *[]byte `plist:",omitempty"`
	RequestType                  string  // must be set to "DeclarativeManagement"
	RequestRequiresNetworkTether *bool   `plist:",omitempty"`
}

// DeclarativeManagementCommand is the top-level structure for the "DeclarativeManagement" Apple MDM command.
type DeclarativeManagementCommand struct {
	Command     DeclarativeManagementPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *DeclarativeManagementCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewDeclarativeManagementCommand creates a new "DeclarativeManagement" Apple MDM command.
func NewDeclarativeManagementCommand() *DeclarativeManagementCommand {
	return &DeclarativeManagementCommand{Command: DeclarativeManagementPayload{RequestType: DeclarativeManagementRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[DeclarativeManagementRequestType] = func() interface{} {
		return NewDeclarativeManagementCommand()
	}
}

// DeclarativeManagementResponse is the command result report (response) for the "DeclarativeManagement" Apple MDM command.
type DeclarativeManagementResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *DeclarativeManagementResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[DeclarativeManagementRequestType] = func() interface{} {
		return new(DeclarativeManagementResponse)
	}
}

const RotateFileVaultKeyRequestType = "RotateFileVaultKey"

type FileVaultUnlock struct {
	Password                 *string `plist:",omitempty"`
	PrivateKeyExport         *[]byte `plist:",omitempty"`
	PrivateKeyExportPassword *string `plist:",omitempty"`
}

// RotateFileVaultKeyPayload is the "inner" command-specific payload for the "RotateFileVaultKey" Apple MDM command.
type RotateFileVaultKeyPayload struct {
	KeyType                      string // possible values: personal, institutional
	FileVaultUnlock              FileVaultUnlock
	NewCertificate               *[]byte `plist:",omitempty"`
	ReplyEncryptionCertificate   *[]byte `plist:",omitempty"`
	RequestType                  string  // must be set to "RotateFileVaultKey"
	RequestRequiresNetworkTether *bool   `plist:",omitempty"`
}

// RotateFileVaultKeyCommand is the top-level structure for the "RotateFileVaultKey" Apple MDM command.
type RotateFileVaultKeyCommand struct {
	Command     RotateFileVaultKeyPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *RotateFileVaultKeyCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewRotateFileVaultKeyCommand creates a new "RotateFileVaultKey" Apple MDM command.
func NewRotateFileVaultKeyCommand() *RotateFileVaultKeyCommand {
	return &RotateFileVaultKeyCommand{Command: RotateFileVaultKeyPayload{RequestType: RotateFileVaultKeyRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[RotateFileVaultKeyRequestType] = func() interface{} {
		return NewRotateFileVaultKeyCommand()
	}
}

type RotateResult struct {
	EncryptedNewRecoveryKey *[]byte `plist:",omitempty"`
}

// RotateFileVaultKeyResponse is the command result report (response) for the "RotateFileVaultKey" Apple MDM command.
type RotateFileVaultKeyResponse struct {
	RotateResult *RotateResult `plist:",omitempty"`
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *RotateFileVaultKeyResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[RotateFileVaultKeyRequestType] = func() interface{} {
		return new(RotateFileVaultKeyResponse)
	}
}
