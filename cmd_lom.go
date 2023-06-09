// Code generated by "admgencmd"; DO NOT EDIT.
// Sources: lom.devicerequest.yaml, lom.setuprequest.yaml
// Options: no-shared=true
package mdmcommands

import "fmt"

const LOMDeviceRequestRequestType = "LOMDeviceRequest"

type RequestListItem struct {
	DeviceRequestType        string // supported values: PowerON, PowerOFF, Reset
	DeviceRequestUUID        string
	DeviceDNSName            string
	PrimaryIPv6AddressList   []string
	SecondaryIPv6AddressList []string
	LOMProtocolVersion       int
}

// LOMDeviceRequestPayload is the "inner" command-specific payload for the "LOMDeviceRequest" Apple MDM command.
type LOMDeviceRequestPayload struct {
	RequestList                  []RequestListItem
	RequestType                  string // supported value: LOMDeviceRequest
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// LOMDeviceRequestCommand is the top-level structure for the "LOMDeviceRequest" Apple MDM command.
type LOMDeviceRequestCommand struct {
	Command     LOMDeviceRequestPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *LOMDeviceRequestCommand) GenericCommand() *GenericCommand {
	if c == nil {
		return nil
	}
	cmd := NewGenericCommand(c.Command.RequestType, c.CommandUUID)
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewLOMDeviceRequestCommand creates a new "LOMDeviceRequest" Apple MDM command.
func NewLOMDeviceRequestCommand(uuid string) *LOMDeviceRequestCommand {
	return &LOMDeviceRequestCommand{
		Command:     LOMDeviceRequestPayload{RequestType: LOMDeviceRequestRequestType},
		CommandUUID: uuid,
	}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[LOMDeviceRequestRequestType] = func(uuid string) interface{} {
		return NewLOMDeviceRequestCommand(uuid)
	}
}

type ResponseListItem struct {
	DeviceRequestSuccess     bool
	DeviceRequestUUID        string
	DeviceRequestReturnError *string `plist:",omitempty"`
}

// LOMDeviceRequestResponse is the command result report (response) for the "LOMDeviceRequest" Apple MDM command.
type LOMDeviceRequestResponse struct {
	ResponseList []ResponseListItem
	GenericResponse
}

// Validate checks for any command response errors.
func (r *LOMDeviceRequestResponse) Validate() error {
	if r.ErrorChain != nil || (r.Status != "Acknowledged" && r.Status != "Idle" && r.Status != "NotNow") {
		return fmt.Errorf("MDM error for status %s: %w", r.Status, r.ErrorChain)
	}
	return nil
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *LOMDeviceRequestResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[LOMDeviceRequestRequestType] = func() interface{} {
		return new(LOMDeviceRequestResponse)
	}
}

const LOMSetupRequestRequestType = "LOMSetupRequest"

// LOMSetupRequestCommand is the top-level structure for the "LOMSetupRequest" Apple MDM command.
type LOMSetupRequestCommand struct {
	Command     GenericCommandPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *LOMSetupRequestCommand) GenericCommand() *GenericCommand {
	if c == nil {
		return nil
	}
	cmd := NewGenericCommand(c.Command.RequestType, c.CommandUUID)
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewLOMSetupRequestCommand creates a new "LOMSetupRequest" Apple MDM command.
func NewLOMSetupRequestCommand(uuid string) *LOMSetupRequestCommand {
	return &LOMSetupRequestCommand{
		Command:     GenericCommandPayload{RequestType: LOMSetupRequestRequestType},
		CommandUUID: uuid,
	}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[LOMSetupRequestRequestType] = func(uuid string) interface{} {
		return NewLOMSetupRequestCommand(uuid)
	}
}

// LOMSetupRequestResponse is the command result report (response) for the "LOMSetupRequest" Apple MDM command.
type LOMSetupRequestResponse struct {
	PrimaryIPv6AddressList   []string
	SecondaryIPv6AddressList []string
	LOMProtocolVersion       int
	GenericResponse
}

// Validate checks for any command response errors.
func (r *LOMSetupRequestResponse) Validate() error {
	if r.ErrorChain != nil || (r.Status != "Acknowledged" && r.Status != "Idle" && r.Status != "NotNow") {
		return fmt.Errorf("MDM error for status %s: %w", r.Status, r.ErrorChain)
	}
	return nil
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *LOMSetupRequestResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[LOMSetupRequestRequestType] = func() interface{} {
		return new(LOMSetupRequestResponse)
	}
}
