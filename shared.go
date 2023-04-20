// Code generated by "admgencmd"; DO NOT EDIT.
package mdmcommands

// GenericCommandPayload is the "inner" generic payload for Apple MDM commands.
type GenericCommandPayload struct {
	RequestType                  string // must be set to MDM command name
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// GenericCommand represents a generic command.
type GenericCommand struct {
	CommandUUID string
	Command     GenericCommandPayload
}

// GenericCommanders can extract a GenericCommand.
type GenericCommander interface {
	GenericCommand() *GenericCommand
}

// NewGenericCommand creates a new generic Apple MDM command.
func NewGenericCommand(requestType string) *GenericCommand {
	return &GenericCommand{Command: GenericCommandPayload{RequestType: requestType}}
}

var newCommandFuncs map[string]func() interface{} = make(map[string]func() interface{})
var newResponseFuncs map[string]func() interface{} = make(map[string]func() interface{})

// NewCommand creates a new command from requestType.
func NewCommand(requestType string) interface{} {
	newCmdFn, ok := newCommandFuncs[requestType]
	if !ok || newCmdFn == nil {
		return nil
	}
	return newCmdFn()
}

// NewResponse creates a new command response from requestType.
func NewResponse(requestType string) interface{} {
	newRespFn, ok := newResponseFuncs[requestType]
	if !ok || newRespFn == nil {
		return nil
	}
	return newRespFn()
}

// ErrorChain represents errors that occured on the client executing an MDM command.
type ErrorChain struct {
	ErrorCode            int
	ErrorDomain          string
	LocalizedDescription string
	USEnglishDescription string
}

// Enrollment represents the various enrollment-related data sent with responses.
type Enrollment struct {
	UDID             *string `plist:",omitempty"`
	UserID           *string `plist:",omitempty"`
	UserShortName    *string `plist:",omitempty"`
	UserLongName     *string `plist:",omitempty"`
	EnrollmentID     *string `plist:",omitempty"`
	EnrollmentUserID *string `plist:",omitempty"`
}
