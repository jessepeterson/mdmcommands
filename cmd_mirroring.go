// Code generated by "admgencmd"; DO NOT EDIT.
// Sources: mirroring.stop.yaml, mirroring.request.yaml
// Options: no-shared=true
package mdmcommands

const StopMirroringRequestType = "StopMirroring"

// StopMirroringCommand is the top-level structure for the "StopMirroring" Apple MDM command.
type StopMirroringCommand struct {
	Command     GenericCommandPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *StopMirroringCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewStopMirroringCommand creates a new "StopMirroring" Apple MDM command.
func NewStopMirroringCommand(uuid string) *StopMirroringCommand {
	return &StopMirroringCommand{
		Command:     GenericCommandPayload{RequestType: StopMirroringRequestType},
		CommandUUID: uuid,
	}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[StopMirroringRequestType] = func(uuid string) interface{} {
		return NewStopMirroringCommand(uuid)
	}
}

// StopMirroringResponse is the command result report (response) for the "StopMirroring" Apple MDM command.
type StopMirroringResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *StopMirroringResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[StopMirroringRequestType] = func() interface{} {
		return new(StopMirroringResponse)
	}
}

const RequestMirroringRequestType = "RequestMirroring"

// RequestMirroringPayload is the "inner" command-specific payload for the "RequestMirroring" Apple MDM command.
type RequestMirroringPayload struct {
	DestinationName              *string `plist:",omitempty"`
	DestinationDeviceID          *string `plist:",omitempty"`
	ScanTime                     *int    `plist:",omitempty"`
	Password                     *string `plist:",omitempty"`
	RequestType                  string  // supported value: RequestMirroring
	RequestRequiresNetworkTether *bool   `plist:",omitempty"`
}

// RequestMirroringCommand is the top-level structure for the "RequestMirroring" Apple MDM command.
type RequestMirroringCommand struct {
	Command     RequestMirroringPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *RequestMirroringCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewRequestMirroringCommand creates a new "RequestMirroring" Apple MDM command.
func NewRequestMirroringCommand(uuid string) *RequestMirroringCommand {
	return &RequestMirroringCommand{
		Command:     RequestMirroringPayload{RequestType: RequestMirroringRequestType},
		CommandUUID: uuid,
	}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[RequestMirroringRequestType] = func(uuid string) interface{} {
		return NewRequestMirroringCommand(uuid)
	}
}

// RequestMirroringResponse is the command result report (response) for the "RequestMirroring" Apple MDM command.
type RequestMirroringResponse struct {
	MirroringResult *string `plist:",omitempty"`
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *RequestMirroringResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[RequestMirroringRequestType] = func() interface{} {
		return new(RequestMirroringResponse)
	}
}
