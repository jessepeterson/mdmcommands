package mdmcommands

//go:generate admgencmd -pkg mdmcommands -o shared.go
//go:generate admgencmd -omit-shared -pkg mdmcommands -o cmd_installmedia.go device-management/mdm/commands/media.install.yaml
//go:generate admgencmd -omit-shared -pkg mdmcommands -o cmd_deviceinfo.go device-management/mdm/commands/information.device.yaml
