#!/usr/bin/env python3

import os
import sys

GOPKG = "mdmcommands"
ADMDIR = "device-management/mdm/commands"

groups = {
    "media": [
        "media.install.yaml",
        "media.remove.yaml",
        "media.managed.list.yaml",
    ],
    "information": [
        "information.device.yaml",
        "information.security.yaml",
        # "information.contentcaching.yaml", # parse/gen problems
        "certificate.list.yaml",
    ],
    "profile": [
        "profile.install.yaml",
        "profile.remove.yaml",
        "profile.list.yaml",
        "profile.provisioning.install.yaml",
        "profile.provisioning.remove.yaml",
        "profile.provisioning.list.yaml",
    ],
    "remotedesktop": [
        "remotedesktop.enable.yaml",
        "remotedesktop.disable.yaml",
    ],
    "mirroring": ["mirroring.stop.yaml", "mirroring.request.yaml"],
    # "settings": ["settings.yaml"], # parse/gen problems, dup structs
    "device": [
        "device.shutdown.yaml",
        "device.restart.yaml",
        "device.configured.yaml",
        "device.erase.yaml",
        "device.lostmode.enable.yaml",
        "device.lostmode.disable.yaml",
        "device.lostmode.location.yaml",
        "device.lostmode.playsound.yaml",
        # "device.restrictions.list.yaml", # parse/gen problems
        "device.restrictions.clearpassword.yaml",
        "device.lock.yaml",
        "device.esim.yaml",
        "device.activationlock.bypasscode.yaml",
        "device.activationlock.clearbypasscode.yaml",
        "declarativemanagement.yaml",
        "rotate.file.vault.key.yaml",
    ],
    "update": [
        "system.update.schedule.yaml",
        "system.update.available.yaml",
        "system.update.scan.yaml",
        "system.update.status.yaml",
    ],
    "user": [
        "user.delete.yaml",
        "user.list.yaml",
        "user.logout.yaml",
        "user.unlock.yaml",
        "set.auto.admin.password.yaml",
        "account.configuration.yaml",
    ],
    "lom": ["lom.devicerequest.yaml", "lom.setuprequest.yaml"],
    "passcode": [
        "passcode.unlocktoken.yaml",
        "passcode.recovery.set.yaml",
        "passcode.clear.yaml",
        "passcode.firmware.verify.yaml",
        "passcode.recovery.verify.yaml",
        "passcode.firmware.set.yaml",
    ],
    "apps": [
        "application.remove.yaml",
        "application.install.yaml",
        "application.install.enterprise.yaml",
        "application.extensions.mappings.yaml",
        "application.validate.yaml",
        "application.installed.list.yaml",
        # "application.extensions.listactive.yaml", # dup structs
        "application.redemptioncode.yaml",
        "application.invitetoprogram.yaml",
        # "application.managed.list.yaml", # parse/gen problems
        # "managed.application.attributes.yaml", # parse/gen problems
        "managed.application.configuration.yaml",
        "managed.application.feedback.yaml",
    ],
}

allfiles = [f for f in os.listdir(ADMDIR) if f.endswith("yaml")]
rfiles = allfiles.copy()

print("package mdmcommands")
print()
print(f"//go:generate admgencmd -pkg {GOPKG} -o shared.go")

for group, files in groups.items():
    for f in files:
        if f in rfiles:
            rfiles.remove(f)
        else:
            sys.stderr.write(
                f"WARNING: {f} specified, but not found in {ADMDIR}\n"
            )

    dir_files = " ".join([f"{ADMDIR}/{f}" for f in files])
    print(
        f"//go:generate admgencmd -pkg {GOPKG} -no-shared -o cmd_{group}.go {dir_files}"
    )

if len(rfiles) >= 1:
    sys.stderr.write(f"WARNING: did not include {len(rfiles)} files:\n")
for f in rfiles:
    sys.stderr.write(f"WARNING: did not include {f}\n")
