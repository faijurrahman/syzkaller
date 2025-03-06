#!/bin/bash

#-----------------------------------------------------------------------
# Setup Source Code
#-----------------------------------------------------------------------
git clone https://github.com/google/syzkaller.git syzkaller
cd syzkaller

#git remote add origin https://github.com/google/syzkaller.git
#git branch --set-upstream-to=origin/master master

git branch -vv
git remote -vv

git remote add dev git@github.com:faijurrahman/syzkaller.git

git branch -vv
git remote -vv


#-----------------------------------------------------------------------
# Prepare Environment for Syzkaller
#-----------------------------------------------------------------------
# Details doc here:
# https://github.com/faijurrahman/syzkaller/blob/master/docs/linux/setup_linux-host_qemu-vm_arm64-kernel.md

# For Kernel please check here:
# https://github.com/faijurrahman/kernel-stable/blob/linux-6.1.y-syzkaller/mywork/syzkaller/prepare-kernel-build-env-for-syzkaller.sh

# For RootFS please check here:
# https://github.com/faijurrahman/buildroot/blob/master/mywork/syzkaller/prepare-buildroot-for-syzkaller.sh

# For Kernel please check here:
# https://github.com/faijurrahman/qemu/blob/master/mywork/syzkaller/prepare-qemu-for-syzkaller.sh
