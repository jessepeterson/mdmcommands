package mdmcommands

//go:generate admgencmd -pkg mdmcommands -o shared.go
//go:generate admgencmd -pkg mdmcommands -no-shared -o cmd_media.go device-management/mdm/commands/media.install.yaml device-management/mdm/commands/media.remove.yaml device-management/mdm/commands/media.managed.list.yaml
//go:generate admgencmd -pkg mdmcommands -no-shared -o cmd_information.go device-management/mdm/commands/information.device.yaml device-management/mdm/commands/information.security.yaml
//go:generate admgencmd -pkg mdmcommands -no-shared -o cmd_profile.go device-management/mdm/commands/profile.install.yaml device-management/mdm/commands/profile.remove.yaml device-management/mdm/commands/profile.list.yaml device-management/mdm/commands/profile.provisioning.install.yaml device-management/mdm/commands/profile.provisioning.remove.yaml device-management/mdm/commands/profile.provisioning.list.yaml
