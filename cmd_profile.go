// Code generated by "admgencmd"; DO NOT EDIT.
// Sources: profile.install.yaml, profile.remove.yaml, profile.list.yaml, profile.provisioning.install.yaml, profile.provisioning.remove.yaml, profile.provisioning.list.yaml
// Options: no-shared=true
package mdmcommands

import "time"

const InstallProfileRequestType = "InstallProfile"

// InstallProfilePayload is the "inner" command-specific payload for the "InstallProfile" Apple MDM command.
type InstallProfilePayload struct {
	Payload                      []byte
	RequestType                  string // must be set to "InstallProfile"
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// InstallProfileCommand is the top-level structure for the "InstallProfile" Apple MDM command.
type InstallProfileCommand struct {
	Command     InstallProfilePayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *InstallProfileCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewInstallProfileCommand creates a new "InstallProfile" Apple MDM command.
func NewInstallProfileCommand() *InstallProfileCommand {
	return &InstallProfileCommand{Command: InstallProfilePayload{RequestType: InstallProfileRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[InstallProfileRequestType] = func() interface{} {
		return NewInstallProfileCommand()
	}
}

// InstallProfileResponse is the command result report (response) for the "InstallProfile" Apple MDM command.
type InstallProfileResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *InstallProfileResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[InstallProfileRequestType] = func() interface{} {
		return new(InstallProfileResponse)
	}
}

const RemoveProfileRequestType = "RemoveProfile"

// RemoveProfilePayload is the "inner" command-specific payload for the "RemoveProfile" Apple MDM command.
type RemoveProfilePayload struct {
	Identifier                   string
	RequestType                  string // must be set to "RemoveProfile"
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// RemoveProfileCommand is the top-level structure for the "RemoveProfile" Apple MDM command.
type RemoveProfileCommand struct {
	Command     RemoveProfilePayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *RemoveProfileCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewRemoveProfileCommand creates a new "RemoveProfile" Apple MDM command.
func NewRemoveProfileCommand() *RemoveProfileCommand {
	return &RemoveProfileCommand{Command: RemoveProfilePayload{RequestType: RemoveProfileRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[RemoveProfileRequestType] = func() interface{} {
		return NewRemoveProfileCommand()
	}
}

// RemoveProfileResponse is the command result report (response) for the "RemoveProfile" Apple MDM command.
type RemoveProfileResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *RemoveProfileResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[RemoveProfileRequestType] = func() interface{} {
		return new(RemoveProfileResponse)
	}
}

const ProfileListRequestType = "ProfileList"

// ProfileListPayload is the "inner" command-specific payload for the "ProfileList" Apple MDM command.
type ProfileListPayload struct {
	ManagedOnly                  *bool  `plist:",omitempty"`
	RequestType                  string // must be set to "ProfileList"
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// ProfileListCommand is the top-level structure for the "ProfileList" Apple MDM command.
type ProfileListCommand struct {
	Command     ProfileListPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *ProfileListCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewProfileListCommand creates a new "ProfileList" Apple MDM command.
func NewProfileListCommand() *ProfileListCommand {
	return &ProfileListCommand{Command: ProfileListPayload{RequestType: ProfileListRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[ProfileListRequestType] = func() interface{} {
		return NewProfileListCommand()
	}
}

type PayloadContentItem struct {
	PayloadType         string
	PayloadVersion      int
	PayloadIdentifier   string
	PayloadDisplayName  *string `plist:",omitempty"`
	PayloadDescription  *string `plist:",omitempty"`
	PayloadOrganization *string `plist:",omitempty"`
}
type ProfileListItem struct {
	PayloadUUID              string
	PayloadIdentifier        string
	PayloadVersion           *int                  `plist:",omitempty"`
	PayloadDisplayName       *string               `plist:",omitempty"`
	PayloadOrganization      *string               `plist:",omitempty"`
	PayloadDescription       *string               `plist:",omitempty"`
	PayloadRemovalDisallowed *bool                 `plist:",omitempty"`
	HasRemovalPasscode       *bool                 `plist:",omitempty"`
	IsEncrypted              *bool                 `plist:",omitempty"`
	SignerCertificates       *[][]byte             `plist:",omitempty"`
	IsManaged                *bool                 `plist:",omitempty"`
	PayloadContent           *[]PayloadContentItem `plist:",omitempty"`
}

// ProfileListResponse is the command result report (response) for the "ProfileList" Apple MDM command.
type ProfileListResponse struct {
	ProfileList []ProfileListItem
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *ProfileListResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[ProfileListRequestType] = func() interface{} {
		return new(ProfileListResponse)
	}
}

const InstallProvisioningProfileRequestType = "InstallProvisioningProfile"

// InstallProvisioningProfilePayload is the "inner" command-specific payload for the "InstallProvisioningProfile" Apple MDM command.
type InstallProvisioningProfilePayload struct {
	ProvisioningProfile          []byte
	RequestType                  string // must be set to "InstallProvisioningProfile"
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// InstallProvisioningProfileCommand is the top-level structure for the "InstallProvisioningProfile" Apple MDM command.
type InstallProvisioningProfileCommand struct {
	Command     InstallProvisioningProfilePayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *InstallProvisioningProfileCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewInstallProvisioningProfileCommand creates a new "InstallProvisioningProfile" Apple MDM command.
func NewInstallProvisioningProfileCommand() *InstallProvisioningProfileCommand {
	return &InstallProvisioningProfileCommand{Command: InstallProvisioningProfilePayload{RequestType: InstallProvisioningProfileRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[InstallProvisioningProfileRequestType] = func() interface{} {
		return NewInstallProvisioningProfileCommand()
	}
}

// InstallProvisioningProfileResponse is the command result report (response) for the "InstallProvisioningProfile" Apple MDM command.
type InstallProvisioningProfileResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *InstallProvisioningProfileResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[InstallProvisioningProfileRequestType] = func() interface{} {
		return new(InstallProvisioningProfileResponse)
	}
}

const RemoveProvisioningProfileRequestType = "RemoveProvisioningProfile"

// RemoveProvisioningProfilePayload is the "inner" command-specific payload for the "RemoveProvisioningProfile" Apple MDM command.
type RemoveProvisioningProfilePayload struct {
	UUID                         string
	RequestType                  string // must be set to "RemoveProvisioningProfile"
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// RemoveProvisioningProfileCommand is the top-level structure for the "RemoveProvisioningProfile" Apple MDM command.
type RemoveProvisioningProfileCommand struct {
	Command     RemoveProvisioningProfilePayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *RemoveProvisioningProfileCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewRemoveProvisioningProfileCommand creates a new "RemoveProvisioningProfile" Apple MDM command.
func NewRemoveProvisioningProfileCommand() *RemoveProvisioningProfileCommand {
	return &RemoveProvisioningProfileCommand{Command: RemoveProvisioningProfilePayload{RequestType: RemoveProvisioningProfileRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[RemoveProvisioningProfileRequestType] = func() interface{} {
		return NewRemoveProvisioningProfileCommand()
	}
}

// RemoveProvisioningProfileResponse is the command result report (response) for the "RemoveProvisioningProfile" Apple MDM command.
type RemoveProvisioningProfileResponse struct {
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *RemoveProvisioningProfileResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[RemoveProvisioningProfileRequestType] = func() interface{} {
		return new(RemoveProvisioningProfileResponse)
	}
}

const ProvisioningProfileListRequestType = "ProvisioningProfileList"

// ProvisioningProfileListPayload is the "inner" command-specific payload for the "ProvisioningProfileList" Apple MDM command.
type ProvisioningProfileListPayload struct {
	ManagedOnly                  *bool  `plist:",omitempty"`
	RequestType                  string // must be set to "ProvisioningProfileList"
	RequestRequiresNetworkTether *bool  `plist:",omitempty"`
}

// ProvisioningProfileListCommand is the top-level structure for the "ProvisioningProfileList" Apple MDM command.
type ProvisioningProfileListCommand struct {
	Command     ProvisioningProfileListPayload
	CommandUUID string
}

// GenericCommand creates a new generic command using the values of c.
func (c *ProvisioningProfileListCommand) GenericCommand() *GenericCommand {
	cmd := NewGenericCommand(c.Command.RequestType)
	cmd.CommandUUID = c.CommandUUID
	cmd.Command.RequestRequiresNetworkTether = c.Command.RequestRequiresNetworkTether
	return cmd
}

// NewProvisioningProfileListCommand creates a new "ProvisioningProfileList" Apple MDM command.
func NewProvisioningProfileListCommand() *ProvisioningProfileListCommand {
	return &ProvisioningProfileListCommand{Command: ProvisioningProfileListPayload{RequestType: ProvisioningProfileListRequestType}}
}

func init() {
	// associate our Request Type to a function for creating a command of that type
	newCommandFuncs[ProvisioningProfileListRequestType] = func() interface{} {
		return NewProvisioningProfileListCommand()
	}
}

type ProvisioningProfileListItem struct {
	Name       string
	UUID       string
	ExpiryDate *time.Time `plist:",omitempty"`
}

// ProvisioningProfileListResponse is the command result report (response) for the "ProvisioningProfileList" Apple MDM command.
type ProvisioningProfileListResponse struct {
	ProvisioningProfileList []ProvisioningProfileListItem
	GenericResponse
}

// GetGenericResponse creates a new generic command response using the values of r.
func (r *ProvisioningProfileListResponse) GetGenericResponse() *GenericResponse {
	return &r.GenericResponse
}

func init() {
	// associate our Request Type to a function for creating a response of that type
	newResponseFuncs[ProvisioningProfileListRequestType] = func() interface{} {
		return new(ProvisioningProfileListResponse)
	}
}
