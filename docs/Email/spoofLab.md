# Install spoof lab

#### Installing Postfix and SendEmail

We need to get the proper tools on our machine before we can do anything. In order to successfully complete an attack, we need to first host our own mail server or use a server which will provide one for us, such as [smtp2go](https://www.smtp2go.com/). For simplicity here, we will host our own server with Postfix. You can install it by running the following command as root:

```
apt install postfix
```



You can ignore any warnings about copying configuration files around. Next, we need to get sendemail, which is what we will use to send the actual fake E-Mail messages. Install it by running the following command as root:

```
apt install sendemail
```



#### Step 2Starting the Mail Server

After getting the proper tools, it's time for us to start our local mail server. Execute as root:

```
systemctl start postfix
```



This will start Postfix E-Mail server and have it listen on port 25 for connections.
You can skip this step, if you are using an existing mail server.

#### Step 3Sending the Email

To craft the custom E-Mail, we will be using SendEmail - a lightweight, command line SMTP client, written by Brandon Zehm. We can do so with the following command:



```
sendemail -f a1ucard0@gmail.com -t dummymail31337@gmail.com -u "A Real Email" -m "This is a very real email"
```

[![img](https://img.wonderhowto.com/img/84/45/63740503344643/0/to-spoof-e-mail-using-sendemail-and-postfix.w1456.jpg)](https://img.wonderhowto.com/img/original/84/45/63740503344643/0/637405033446438445.jpg)



As you can see, the E-Mail was sent successfully and if I check my inbox...

[![img](https://img.wonderhowto.com/img/95/98/63740503410893/0/to-spoof-e-mail-using-sendemail-and-postfix.w1456.jpg)](https://img.wonderhowto.com/img/original/95/98/63740503410893/0/637405034108939598.jpg)



You might have noticed that the E-Mail got delivered to my spam folder. This is because I am using my own mail server and so the message never gets encrypted or validated by Google or my ISP's mail server. This wouldn't be the case if I actually used an external mail server, such as the ones provided by smtp2go.

Let's dissect the above command to learn what each of the arguments mean:
**-f** - used to specify the sender's email address. Here you have to input the address you want to appear as.
**-t** - used to specify the recipient's email address.
**-u** - specifies the subject of the E-Mail. Note that this option is not required.
**-m** - specifies the actual contents of the E-Mail.



Some additional arguments you can use are:
**-a** - followed by a filepath for attaching files.

**-o message-file=FILENAME** - you can specify a file that contains the contents of the E-Mail, instead of manually typing them on the command line.



**-o message-header="Name:Value"** - here you can specify additional E-Mail headers such as the "From:" header which will change the name (but not the E-Mail address, that is done with **-f**) of the sender.

**-s SERVER:PORT** - you can use this to specify the IP Address and port of the server you will be using. Defaults to localhost:25.



#### Step 4Cleaning Up

Don't forget to shut down the Postfix mail server:
**$ systemctl stop postfix**

#### Conclusion

Now you know just how easy it is to spoof E-Mail. You also surely realize what the consequences of such an attack might be!



I have shown you this purely for educational purposes and I am not responsible for anything you do with that knowledge!
Keep learning and see you soon!

```
# reconfigure

dpkg-reconfigure postfix
```

https://blog.spookysec.net//email-spoofing/

https://github.com/pablokbs/peladonerd/issues/53



```
sendemail -f fromuser@gmail.com -t touser@domain.com -u subject -m "message" -s smtp.gmail.com:587 -o tls=yes -xu gmailaccount@gmail.com -xp gmailpassword 

```

