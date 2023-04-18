// Code generated by "admgencmd"; DO NOT EDIT.
package mdmcommands

// InstallMediaPayload
type InstallMediaPayload struct {
	ITunesStoreID int    `plist:"iTunesStoreID,omitempty"`
	MediaURL      string `plist:"MediaURL,omitempty"`
	MediaType     string `plist:"MediaType"`
	PersistentID  string `plist:"PersistentID,omitempty"`
	Kind          string `plist:"Kind,omitempty"`
	Version       string `plist:"Version,omitempty"`
	Author        string `plist:"Author,omitempty"`
	Title         string `plist:"Title,omitempty"`
	RequestType   string `plist:"RequestType"`
}

// InstallMediaCommand
type InstallMediaCommand struct {
	Command     InstallMediaPayload `plist:"Command"`
	CommandUUID string              `plist:"CommandUUID"`
}

// InstallMediaResponse
type InstallMediaResponse struct {
	ITunesStoreID   int    `plist:"iTunesStoreID,omitempty"`
	MediaURL        string `plist:"MediaURL,omitempty"`
	PersistentID    string `plist:"PersistentID,omitempty"`
	MediaType       string `plist:"MediaType,omitempty"`
	State           string `plist:"State,omitempty"`
	RejectionReason string `plist:"RejectionReason,omitempty"`
	CommandUUID     string `plist:"CommandUUID"`
	Status          string `plist:"Status"`
}
