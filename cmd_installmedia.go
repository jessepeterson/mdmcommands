// Code generated by "admgencmd"; DO NOT EDIT.
// Source: media.install.yaml
// Options: no-shared=true
package mdmcommands

const InstallMediaRequestType = "InstallMedia"

// InstallMediaPayload is the "inner" command-specific payload for the "InstallMedia" Apple MDM command.
type InstallMediaPayload struct {
	ITunesStoreID                *int    `plist:"iTunesStoreID,omitempty"`
	MediaURL                     *string `plist:",omitempty"`
	MediaType                    string
	PersistentID                 *string `plist:",omitempty"`
	Kind                         *string `plist:",omitempty"`
	Version                      *string `plist:",omitempty"`
	Author                       *string `plist:",omitempty"`
	Title                        *string `plist:",omitempty"`
	RequestType                  string  // must be set to "InstallMedia"
	RequestRequiresNetworkTether *bool   `plist:",omitempty"`
}

// InstallMediaCommand is the top-level structure for the "InstallMedia" Apple MDM command.
type InstallMediaCommand struct {
	Command     InstallMediaPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c
func (c *InstallMediaCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewInstallMediaCommand creates a new "InstallMedia" Apple MDM command.
func NewInstallMediaCommand() *InstallMediaCommand {
	return &InstallMediaCommand{Command: InstallMediaPayload{RequestType: InstallMediaRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[InstallMediaRequestType] = func() interface{} {
		return NewInstallMediaCommand()
	}
}

type InstallMediaResponse struct {
	ITunesStoreID   *int    `plist:"iTunesStoreID,omitempty"`
	MediaURL        *string `plist:",omitempty"`
	PersistentID    *string `plist:",omitempty"`
	MediaType       *string `plist:",omitempty"`
	State           *string `plist:",omitempty"`
	RejectionReason *string `plist:",omitempty"`
	GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[InstallMediaRequestType] = func() interface{} {
		return new(InstallMediaResponse)
	}
}
