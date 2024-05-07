

##### Scan Open Ports

```
nmap -p- --open -sS --min-rate 5000 -vvv -Pn <ip> -oG allPosrts
```



##### Scan Ports

```
nmap -p<80,22,443> -sCV <ip>
```

##### WFUZZ

Domain and S

```
wfuzz -c --hc 404 -t 200 -w /usr/share/SecLists/Discovery/Web-Content/directory-list-2.3-medium.txt <ip>
```

```
sudo wfuzz -c -f sub-fighter.txt -Z -w /usr/share/SecLists/Discovery/DNS/subdomains-top1million-5000.txt --sc 200,202,204,301,302,307,403 FUZZ.docugene.eu
```

