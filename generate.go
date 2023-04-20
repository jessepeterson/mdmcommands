package mdmcommands

//go:generate admgencmd -pkg mdmcommands -o shared.go
//go:generate admgencmd -pkg mdmcommands -no-shared -o cmd_installmedia.go device-management/mdm/commands/media.install.yaml
//go:generate admgencmd -pkg mdmcommands -no-shared -o cmd_deviceinfo.go device-management/mdm/commands/information.device.yaml
