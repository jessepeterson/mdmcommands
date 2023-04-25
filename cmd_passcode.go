// Code generated by "admgencmd"; DO NOT EDIT.
// Sources: passcode.unlocktoken.yaml, passcode.recovery.set.yaml, passcode.clear.yaml, passcode.firmware.verify.yaml, passcode.recovery.verify.yaml, passcode.firmware.set.yaml
// Options: no-shared=true
package mdmcommands

const RequestUnlockTokenRequestType = "RequestUnlockToken"

// RequestUnlockTokenCommand is the top-level structure for the "RequestUnlockToken" Apple MDM command.
type RequestUnlockTokenCommand struct {
	Command     GenericCommandPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *RequestUnlockTokenCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewRequestUnlockTokenCommand creates a new "RequestUnlockToken" Apple MDM command.
func NewRequestUnlockTokenCommand() *RequestUnlockTokenCommand {
	return &RequestUnlockTokenCommand{Command: GenericCommandPayload{RequestType: RequestUnlockTokenRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[RequestUnlockTokenRequestType] = func() interface{} {
		return NewRequestUnlockTokenCommand()
	}
}

// RequestUnlockTokenResponse is the command result report (response) for the "RequestUnlockToken" Apple MDM command.
type RequestUnlockTokenResponse struct {
	UnlockToken []byte
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *RequestUnlockTokenResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[RequestUnlockTokenRequestType] = func() interface{} {
		return new(RequestUnlockTokenResponse)
	}
}

const SetRecoveryLockRequestType = "SetRecoveryLock"

// SetRecoveryLockPayload is the "inner" command-specific payload for the "SetRecoveryLock" Apple MDM command.
type SetRecoveryLockPayload struct {
	CurrentPassword              *string `plist:",omitempty"`
	NewPassword                  string
	RequestType                  string // must be set to "SetRecoveryLock"
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// SetRecoveryLockCommand is the top-level structure for the "SetRecoveryLock" Apple MDM command.
type SetRecoveryLockCommand struct {
	Command     SetRecoveryLockPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *SetRecoveryLockCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewSetRecoveryLockCommand creates a new "SetRecoveryLock" Apple MDM command.
func NewSetRecoveryLockCommand() *SetRecoveryLockCommand {
	return &SetRecoveryLockCommand{Command: SetRecoveryLockPayload{RequestType: SetRecoveryLockRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[SetRecoveryLockRequestType] = func() interface{} {
		return NewSetRecoveryLockCommand()
	}
}

// SetRecoveryLockResponse is the command result report (response) for the "SetRecoveryLock" Apple MDM command.
type SetRecoveryLockResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *SetRecoveryLockResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[SetRecoveryLockRequestType] = func() interface{} {
		return new(SetRecoveryLockResponse)
	}
}

const ClearPasscodeRequestType = "ClearPasscode"

// ClearPasscodePayload is the "inner" command-specific payload for the "ClearPasscode" Apple MDM command.
type ClearPasscodePayload struct {
	UnlockToken                  []byte
	RequestType                  string // must be set to "ClearPasscode"
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// ClearPasscodeCommand is the top-level structure for the "ClearPasscode" Apple MDM command.
type ClearPasscodeCommand struct {
	Command     ClearPasscodePayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *ClearPasscodeCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewClearPasscodeCommand creates a new "ClearPasscode" Apple MDM command.
func NewClearPasscodeCommand() *ClearPasscodeCommand {
	return &ClearPasscodeCommand{Command: ClearPasscodePayload{RequestType: ClearPasscodeRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[ClearPasscodeRequestType] = func() interface{} {
		return NewClearPasscodeCommand()
	}
}

// ClearPasscodeResponse is the command result report (response) for the "ClearPasscode" Apple MDM command.
type ClearPasscodeResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *ClearPasscodeResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[ClearPasscodeRequestType] = func() interface{} {
		return new(ClearPasscodeResponse)
	}
}

const VerifyFirmwarePasswordRequestType = "VerifyFirmwarePassword"

// VerifyFirmwarePasswordPayload is the "inner" command-specific payload for the "VerifyFirmwarePassword" Apple MDM command.
type VerifyFirmwarePasswordPayload struct {
	Password                     string
	RequestType                  string // must be set to "VerifyFirmwarePassword"
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// VerifyFirmwarePasswordCommand is the top-level structure for the "VerifyFirmwarePassword" Apple MDM command.
type VerifyFirmwarePasswordCommand struct {
	Command     VerifyFirmwarePasswordPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *VerifyFirmwarePasswordCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewVerifyFirmwarePasswordCommand creates a new "VerifyFirmwarePassword" Apple MDM command.
func NewVerifyFirmwarePasswordCommand() *VerifyFirmwarePasswordCommand {
	return &VerifyFirmwarePasswordCommand{Command: VerifyFirmwarePasswordPayload{RequestType: VerifyFirmwarePasswordRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[VerifyFirmwarePasswordRequestType] = func() interface{} {
		return NewVerifyFirmwarePasswordCommand()
	}
}

// VerifyFirmwarePasswordResponse is the command result report (response) for the "VerifyFirmwarePassword" Apple MDM command.
type VerifyFirmwarePasswordResponse struct {
	PasswordVerified bool
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *VerifyFirmwarePasswordResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[VerifyFirmwarePasswordRequestType] = func() interface{} {
		return new(VerifyFirmwarePasswordResponse)
	}
}

const VerifyRecoveryLockRequestType = "VerifyRecoveryLock"

// VerifyRecoveryLockPayload is the "inner" command-specific payload for the "VerifyRecoveryLock" Apple MDM command.
type VerifyRecoveryLockPayload struct {
	Password                     string
	RequestType                  string // must be set to "VerifyRecoveryLock"
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// VerifyRecoveryLockCommand is the top-level structure for the "VerifyRecoveryLock" Apple MDM command.
type VerifyRecoveryLockCommand struct {
	Command     VerifyRecoveryLockPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *VerifyRecoveryLockCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewVerifyRecoveryLockCommand creates a new "VerifyRecoveryLock" Apple MDM command.
func NewVerifyRecoveryLockCommand() *VerifyRecoveryLockCommand {
	return &VerifyRecoveryLockCommand{Command: VerifyRecoveryLockPayload{RequestType: VerifyRecoveryLockRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[VerifyRecoveryLockRequestType] = func() interface{} {
		return NewVerifyRecoveryLockCommand()
	}
}

// VerifyRecoveryLockResponse is the command result report (response) for the "VerifyRecoveryLock" Apple MDM command.
type VerifyRecoveryLockResponse struct {
	PasswordVerified bool
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *VerifyRecoveryLockResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[VerifyRecoveryLockRequestType] = func() interface{} {
		return new(VerifyRecoveryLockResponse)
	}
}

const SetFirmwarePasswordRequestType = "SetFirmwarePassword"

// SetFirmwarePasswordPayload is the "inner" command-specific payload for the "SetFirmwarePassword" Apple MDM command.
type SetFirmwarePasswordPayload struct {
	CurrentPassword              *string `plist:",omitempty"`
	NewPassword                  string
	AllowOroms                   *bool  `plist:",omitempty"`
	RequestType                  string // must be set to "SetFirmwarePassword"
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// SetFirmwarePasswordCommand is the top-level structure for the "SetFirmwarePassword" Apple MDM command.
type SetFirmwarePasswordCommand struct {
	Command     SetFirmwarePasswordPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *SetFirmwarePasswordCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewSetFirmwarePasswordCommand creates a new "SetFirmwarePassword" Apple MDM command.
func NewSetFirmwarePasswordCommand() *SetFirmwarePasswordCommand {
	return &SetFirmwarePasswordCommand{Command: SetFirmwarePasswordPayload{RequestType: SetFirmwarePasswordRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[SetFirmwarePasswordRequestType] = func() interface{} {
		return NewSetFirmwarePasswordCommand()
	}
}

// SetFirmwarePasswordResponse is the command result report (response) for the "SetFirmwarePassword" Apple MDM command.
type SetFirmwarePasswordResponse struct {
	PasswordChanged bool
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *SetFirmwarePasswordResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[SetFirmwarePasswordRequestType] = func() interface{} {
		return new(SetFirmwarePasswordResponse)
	}
}
