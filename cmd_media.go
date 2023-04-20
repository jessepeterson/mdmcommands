// Code generated by "admgencmd"; DO NOT EDIT.
// Sources: media.install.yaml, media.remove.yaml, media.managed.list.yaml
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

const RemoveMediaRequestType = "RemoveMedia"

// RemoveMediaPayload is the "inner" command-specific payload for the "RemoveMedia" Apple MDM command.
type RemoveMediaPayload struct {
	MediaType                    string
	ITunesStoreID                *string `plist:"iTunesStoreID,omitempty"`
	PersistentID                 *string `plist:",omitempty"`
	RequestType                  string  // must be set to "RemoveMedia"
	RequestRequiresNetworkTether *bool   `plist:",omitempty"`
}

// RemoveMediaCommand is the top-level structure for the "RemoveMedia" Apple MDM command.
type RemoveMediaCommand struct {
	Command     RemoveMediaPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c
func (c *RemoveMediaCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewRemoveMediaCommand creates a new "RemoveMedia" Apple MDM command.
func NewRemoveMediaCommand() *RemoveMediaCommand {
	return &RemoveMediaCommand{Command: RemoveMediaPayload{RequestType: RemoveMediaRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[RemoveMediaRequestType] = func() interface{} {
		return NewRemoveMediaCommand()
	}
}

type RemoveMediaResponse struct {
	GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[RemoveMediaRequestType] = func() interface{} {
		return new(RemoveMediaResponse)
	}
}

const ManagedMediaListRequestType = "ManagedMediaList"

// ManagedMediaListCommand is the top-level structure for the "ManagedMediaList" Apple MDM command.
type ManagedMediaListCommand struct {
	Command     GenericCommandPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c
func (c *ManagedMediaListCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewManagedMediaListCommand creates a new "ManagedMediaList" Apple MDM command.
func NewManagedMediaListCommand() *ManagedMediaListCommand {
	return &ManagedMediaListCommand{Command: GenericCommandPayload{RequestType: ManagedMediaListRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[ManagedMediaListRequestType] = func() interface{} {
		return NewManagedMediaListCommand()
	}
}

type BooksItem struct {
	ITunesStoreID int     `plist:"iTunesStoreID"`
	State         *string `plist:",omitempty"`
	PersistentID  *string `plist:",omitempty"`
	Kind          *string `plist:",omitempty"`
	Version       *string `plist:",omitempty"`
	Author        *string `plist:",omitempty"`
	Title         *string `plist:",omitempty"`
}
type ManagedMediaListResponse struct {
	Books []BooksItem
	GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[ManagedMediaListRequestType] = func() interface{} {
		return new(ManagedMediaListResponse)
	}
}
