#!/usr/bin/env python3

import os
import sys

GOPKG = "mdmcommands"
ADMDIR = "device-management/mdm/commands"

groups = {
    "installmedia": ["media.install.yaml"],
    "deviceinfo": ["information.device.yaml", "information.security.yaml"],
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
            sys.stderr.write(f"WARNING: {f} specified, but not found in {ADMDIR}\n")

    dir_files = " ".join([f"{ADMDIR}/{f}" for f in files])
    print(
        f"//go:generate admgencmd -pkg {GOPKG} -no-shared -o cmd_{group}.go {dir_files}"
    )

if len(rfiles) >= 1:
    sys.stderr.write(f"WARNING: did not include {len(rfiles)} files:\n")
for f in rfiles:
    sys.stderr.write(f"WARNING: did not include {f}\n")
