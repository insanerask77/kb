### Tunnel disconnect

Linux :

- create a new bash file, ( **$** `sudo nano /etc/wireguard/wg0retry.bash` ) with this content:



```
#!/bin/bash
while true
do
   sleep 15
   wg-quick up wg0
done
```

**EDIT**: this script (replacement for the bloc of code above - `wg0retry.bash`) is more robust if you don't mind the little extra bandwidth. I added a ping every 15sec to validate that the tunnel is connected, replace "<your_wireguard_server_ip>" for the IP of your Wireguard "server" :

```
#!/bin/bash
while true
do
    sleep 15
    ping -c 1 <your_wireguard_server_ip>
    if [ $? != 0 ]
    then
        wg-quick down wg0
        sleep 4
        wg-quick up wg0
    fi
done
```

it will retry every 15 seconds to start the wg0 tunnel, so when the tunnel is already working it will do nothing and when the tunnel is down it will revive it.

- add a superuser crontab ( **$** `sudo crontab -e` ) to run the previously created script at boot :



```
@reboot bash /etc/wireguard/wg0retry.bash
```

- And you're done! Wireguard is now indestructible on Linux.

Windows:

- Install wireguard-windows
- Copy your configuration file to the Wireguard installation folder C:\Program Files\WireGuard\wg0.conf
- Create a ".bat" file with this content and save it to the Wireguard installation folder `C:\Program Files\WireGuard\wg0retry.bat` :



```
@echo OFF

:loop

set _ServiceName=WireGuardTunnel$wg0

sc query %_ServiceName% | find "does not exist" >nul
if %ERRORLEVEL% EQU 0 "C:\Program Files\WireGuard\WireGuard.exe" /installtunnelservice "C:\Program Files\WireGuard\wg0.conf" & echo "installing service"
if %ERRORLEVEL% EQU 1 echo "Service Exist!"

sc query %_ServiceName% | find "STOPPED" >nul
if %ERRORLEVEL% EQU 0 sc start %_ServiceName% & echo "starting service"
if %ERRORLEVEL% EQU 1 echo "Service Started!"
timeout /t 15

goto loop
```

**EDIT**: this script (replacement for the bloc of code above - `wg0retry.bat`) is more robust if you don't mind the little extra bandwidth. I added a ping every 15sec to validate that the tunnel is connected, replace "<your_wireguard_server_ip>" for the IP of your Wireguard "server" :

```
@echo OFF

:loop

set _ServiceName=WireGuardTunnel$wg0

sc query %_ServiceName% | find "does not exist" >nul
if %ERRORLEVEL% EQU 0 "C:\Program Files\WireGuard\WireGuard.exe" /installtunnelservice "C:\Program Files\WireGuard\wg0.conf" & echo "installing service"
if %ERRORLEVEL% EQU 1 echo "Service Exist!"

sc query %_ServiceName% | find "STOPPED" >nul
if %ERRORLEVEL% EQU 0 sc start %_ServiceName% & echo "starting service"
if %ERRORLEVEL% EQU 1 echo "Service Started!"

timeout /t 15

ping -n 1 <your_wireguard_server_ip> | find "TTL=" >nul
if errorlevel 1 (
    echo "server not responding, restarting service"
    sc stop %_ServiceName%
    timeout /t 5
    sc start %_ServiceName%
    timeout /t 5
)

goto loop
```

this batch script will check in the Windows Service Control (sc) if the service for the tunnel of Wireguard wg0 exist, if it exist, the script does nothing here. If the service does not exist, it will run the command line installation for the Wireguard tunnel service (/installtunnelservice).

Then the script will check if the service for the tunnel of Wireguard wg0 is running. If it's running the script does nothing. If it's not running, it will start the service. The script loop every 15 seconds.

- download NSSM : http://nssm.cc/download
- extract the nssm zip file to `C:\Program Files\NSSM\` since you don't want to delete or move this executable after the next step.
- Run cmd.exe as Administrator, and execute those lines:



```
"C:\Program Files\NSSM\win64\nssm.exe" install wg0retry "C:\Program Files\WireGuard\wg0retry.bat"
"C:\Program Files\NSSM\win64\nssm.exe" start wg0retry
```

It will install a new Windows Service that will start at boot and will execute our script wg0retry.bat

- And you're done! Wireguard is now indestructible on Windows.

\* Please note that this method is independent of the Wireguard tray application and even if you quit Wireguard, or use it to disconnect, the tunnel will still be connected. To be able to disconnect the tunnel, You need to stop the service in (in Administrator cmd.exe) :

```
"C:\Program Files\NSSM\win64\nssm.exe" stop wg0retry
```

or if you want to remove it permanently:

```
"C:\Program Files\NSSM\win64\nssm.exe" remove wg0retry
```

Hope this helps.