let’s say we have domain called **example.com**

first we have to do subdomain enumeration , I usually use subfinder with editing the API configuration file you can check it on https://securitytrails.com/blog/subfinder

```
subfinder -d example.com -all -cs > main.txt ; cat main.txt | cut -d "," -f 1 > domains.txt ; rm -rf main.txt
```

now we have created a file named **domains.txt** and then we need to check the live subdomains and checking the status code of them

simply we can use httpx tool by typing command

```
cat domains.txt | httpx -title -wc -sc -cl -ct -location -web-server -asn -o alive-subdomains.txt 
```

I found interested subdomain with error code 404 that points to “GOOGLE-CLOUD-PLATFORM”

let’s name it **altice.example.com**

after visited that url I found a defualt NOT Found page for “Leadpages services”

![img](https://miro.medium.com/v2/resize:fit:700/1*ztGCUXKvVP26OfUnxcBHCw.png)

Here we start to retrieving information about DNS name servers by dig command

we can type on terminal

```
dig altice.example.com
```

or visiting https://www.digwebinterface.com/

![img](https://miro.medium.com/v2/resize:fit:700/1*IB6gCh7RW4UQrX2cQbzn6A.png)

in this result we found the CNAME server is **custom-proxy.leadpages.net**

I tried to visit https://github.com/EdOverflow/can-i-take-over-xyz to check if that service is vulnerable or not but unfortunately, I haven’t found it there

But also that doesn’t mean “Leadpages services” isn’t vulnerable

After creating a free trial for 14 days account and you have to put your valid paypal email or valid credit card on https://www.leadpages.com/

![img](https://miro.medium.com/v2/resize:fit:700/1*Frf5ZVqU0aJf_llnLkdeBg.png)

Start modifying the template and change it to my name as a POC

Here’s the most exciting part

![img](https://miro.medium.com/v2/resize:fit:404/1*MCmQ0B-8TGazxC0BKDTTSw.gif)

Click on **upadte** > **site** **publishing options**

put your vulnerable subdomain in our case : **altice.example.com**

![img](https://miro.medium.com/v2/resize:fit:700/1*KVydSRfjsghnbY48AGPBIg.png)

click on done

now the ssl will be connected to our custom domain

![img](https://miro.medium.com/v2/resize:fit:700/1*6PsXWQ7BLBznfkxU_XrRjA.png)

let’s visit the vulnerable site right now

![img](https://miro.medium.com/v2/resize:fit:700/1*GLE7xfDLErT-02EH-fEvUQ.png)

And I Finally Found My name

![img](https://miro.medium.com/v2/resize:fit:489/1*CF10dLLcjbNw7GeV_uUD9A.png)

while I was writing that writeup I got some ideas of google dorks that may help U

```
site:"*.example.com" intext:"PAGE NOT FOUND" | intext:"project not found" | intext:"Repository not found"  | intext:"domain does not exist" | intext:"This page could not be found" | intext:"404 Blog is not found" | intext:"No settings were found for this company" | intext:"domain name is invalid"
```

In case you have any questions you can ask me in the comments section and I will asnwer it happily

and don’t forget to follow me in [twitter ](https://twitter.com/Amr_MustafaAA)to get useful informations & tips

Thank U ❤