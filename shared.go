// Code generated by "admgencmd"; DO NOT EDIT.
package mdmcommands

// GenericCommandPayload is the "inner" generic payload for Apple MDM commands.
type GenericCommandPayload struct {
	RequestType                  string // supported value: the MDM command name
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

// GenericResponsers can extract a GenericResponse.
type GenericResponser interface {
	GetGenericResponse() *GenericResponse
}

// NewGenericCommand creates a new generic Apple MDM command.
func NewGenericCommand(requestType string) *GenericCommand {
	return &GenericCommand{Command: GenericCommandPayload{RequestType: requestType}}
}

var newCommandFuncs map[string]func(string) interface{} = make(map[string]func(string) interface{})
var newResponseFuncs map[string]func() interface{} = make(map[string]func() interface{})

// NewCommand creates a new command from requestType.
func NewCommand(requestType string, uuid string) interface{} {
	newCmdFn, ok := newCommandFuncs[requestType]
	if !ok || newCmdFn == nil {
		return nil
	}
	return newCmdFn(uuid)
}

// ValidRequestType checks that we are able to create a new command from requestType.
func ValidRequestType(requestType string) bool {
	_, ok := newCommandFuncs[requestType]
	return ok
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

// GenericResponse represents the common MDM command response fields.
type GenericResponse struct {
	CommandUUID  string
	Status       string // supported values: Acknowledged, Error, CommandFormatError, Idle, NotNow
	NotOnConsole bool
	ErrorChain   *[]ErrorChain `plist:",omitempty"`
	Enrollment
}
