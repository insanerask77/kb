#### No internet in wsl

```powershell
netsh winsock reset 
netsh int ip reset all
netsh winhttp reset proxy
ipconfig /flushdns
wsl --shutdown
wsl --restart
```

```
Open Powershell or Cmd as Administrator
and run each of these commands:
wsl --shutdown
netsh winsock reset
netsh int ip reset all
netsh winhttp reset proxy
ipconfig /flushdns
Hit the Windows Key,
type Network Reset,
hit enter.
You should see this window.
Click "Reset now".
Restart Windows
```

