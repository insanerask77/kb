```
In case our automatic password reset does not work, you need to reset the password manually trough rescue:

1. Goto the Rescue Section of your server, but click on the Activate Rescue & Power Cycle button instead of the (reset root password) button.
2. After that your server will boot into a rescue system.
3. Connect to your server via SSH with the rescue login details
4. Run the following commands in Rescue:
---8<---
mount /dev/sda1 /mnt

chroot-prepare /mnt

chroot /mnt

passwd
---8<---

5. Reboot your server and login with the new password (Also check whether the Quemu Guest Agent is installed).

If you have any further questions, feel free to contact u
```

