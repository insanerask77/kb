Scan steps

```
nmap -iL domains.txt -p-
nmap -iL domains.txt -T5 -p- --min-rate 100 --open -oA output
sudo nmap -p- --open --min-rate 5000 -n -sS -vvv -Pn 95.217.210.147 -oG allPorts
sudo nmap -p- --open -T5 -sS -vvv -Pn -n 95.217.210.147 -oG allPorts
nmap --script resolveall --script-args newtargets,resolveall.hosts=eulog.3irobotics.net
```

extractPorts

```
function extractPorts(){
	ports="$(cat $1 | grep -oP '\d{1,5}/open' | awk '{print $1}' FS='/' | xargs | tr ' ' ',')"
	ip_address="$(cat $1 | grep -oP '\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}' | sort -u | head -n 1)"
	echo -e "\n[*] Extracting information...\n" > extractPorts.tmp
	echo -e "\t[*] IP Address: $ip_address"  >> extractPorts.tmp
	echo -e "\t[*] Open ports: $ports\n"  >> extractPorts.tmp
	echo $ports | tr -d '\n' | xclip -sel clip
	echo -e "[*] Ports copied to clipboard\n"  >> extractPorts.tmp
	cat extractPorts.tmp; rm extractPorts.tmp
}
```

Discover

```
sudo nmap -p22,80,443 -sCV bluetrail.software -oN targeted
```

##### Subfinder

```
▶ Enumerate/Collect all subdomains using tools like subfinder, assetfinder, Knockpy and haktrails, etc.

Run : subfinder -d someweb.com -o subf.txt -v

Run : echo “someweb.com” | haktrails subdomains > haksubs.txt

Run : assetfinder -subs-only someweb.com > asset.txt

▶ Add all Enumerated/Collected subdomains from different tools in different files into one file with unique subdomains, that may be subdomains.txt

Run : cat subf.txt haksubs.txt asset.txt | sort -u > subdomains.txt

▶ Now, will check the identified subdomains are active or not -

Run : httpx -l subdomains.txt -o activesubs.txt -threads 200 -status-code -follow-redirects
```

