```
--send-mail-url=smtp://smtp.gmail.com:587/?starttls=when-available
--send-mail-any-operation=true
--send-mail-subject=Duplicati %PARSEDRESULT%, %OPERATIONNAME% report for %backup-name%
--send-mail-to=destination_email_address@whatever.com
--send-mail-username=your_sending_gmail_username@gmail.com
--send-mail-password=your_sending_gmail_password
--send-mail-from=This_computers_name backup <your_sending_gmail_username@gmail.com>
```

```
--send-mail-from=aaa@bluetrailsoft.com
--send-mail-username=aaa@bluetrailsoft.com
--send-mail-level=all
--send-mail-password=
--send-mail-url=smtp://smtp.gmail.com:587/?starttls=when-available
--send-mail-subject=Duplicati %OPERATIONNAME%
--send-mailbody=%backup-name%
--send-mail-to=rafaelm@bluetrailsoft.com
```



# [How to configure automatic email notifications via gmail for every backup job](https://forum.duplicati.com/t/how-to-configure-automatic-email-notifications-via-gmail-for-every-backup-job/869)



This article explains how to configure automated email notifications via gmail at the conclusion of every Duplicati operation (backup, restore, etc.). With this approach, every Duplicati job on the machine will send automated notifications – you don’t need to configure each job individually.

- Open Duplicati

- Click on Settings.

  <details open="" style="display: block; position: relative; outline: none;"><summary style="display: block; outline: none; cursor: pointer;"><span>&nbsp;</span>Screenshot 1</summary><p style="display: block;"></p><div class="lightbox-wrapper" style="display: block; outline: 0px;"><a class="lightbox" href="https://forum.duplicati.com/uploads/default/original/1X/38ee273086c62464bae9889d16b2f6d350fa15ef.png" data-download-href="https://forum.duplicati.com/uploads/default/38ee273086c62464bae9889d16b2f6d350fa15ef" title="image.png" style="background-color: transparent; color: var(--tertiary); text-decoration: none; cursor: pointer; outline: 0px; overflow-wrap: break-word; position: relative; display: inline-block; overflow: hidden; transition: all 0.6s cubic-bezier(0.165, 0.84, 0.44, 1) 0s;"><img src="https://forum.duplicati.com/uploads/default/optimized/1X/38ee273086c62464bae9889d16b2f6d350fa15ef_2_690x222.png" alt="image" width="690" height="222" srcset="https://forum.duplicati.com/uploads/default/optimized/1X/38ee273086c62464bae9889d16b2f6d350fa15ef_2_690x222.png, https://forum.duplicati.com/uploads/default/optimized/1X/38ee273086c62464bae9889d16b2f6d350fa15ef_2_1035x333.png 1.5x, https://forum.duplicati.com/uploads/default/optimized/1X/38ee273086c62464bae9889d16b2f6d350fa15ef_2_1380x444.png 2x" data-small-upload="https://forum.duplicati.com/uploads/default/optimized/1X/38ee273086c62464bae9889d16b2f6d350fa15ef_2_10x10.png" loading="lazy" style="border-style: none; vertical-align: middle; outline: 0px; object-fit: cover; object-position: center top; max-width: 100%; height: auto; aspect-ratio: 690 / 222;"><div class="meta" style="color: var(--secondary); font-weight: bold; outline: 0px; position: absolute; bottom: 0px; width: 652.5px; background: var(--primary); opacity: 0; transition: opacity 0.2s ease 0s; display: flex; align-items: center;"><svg class="fa d-icon d-icon-far-image svg-icon" aria-hidden="true"><use xlink:href="#far-image"></use></svg><span class="filename" style="outline: 0px; margin: 6px 6px 6px 0px; overflow: hidden; white-space: nowrap; text-overflow: ellipsis;">image.png</span><span class="informations" style="outline: 0px; margin: 6px; padding-right: 20px; color: var(--secondary-high); font-size: var(--font-0); flex-shrink: 0; flex-grow: 3;">1439×463 34.3 KB</span><svg class="fa d-icon d-icon-discourse-expand svg-icon" aria-hidden="true"><use xlink:href="#discourse-expand"></use></svg></div></a></div><p style="display: block;"></p></details>

- Scroll down to “default options” and click on “edit as text”.

  <details open="" style="display: block; position: relative; outline: none;"><summary style="display: block; outline: none; cursor: pointer;"><span>&nbsp;</span>Screenshot 2</summary><p style="display: block;"></p><div class="lightbox-wrapper" style="display: block; outline: 0px;"><a class="lightbox" href="https://forum.duplicati.com/uploads/default/original/1X/fa3444a26b828f508b9ebf3fbd8fe0b73c295364.png" data-download-href="https://forum.duplicati.com/uploads/default/fa3444a26b828f508b9ebf3fbd8fe0b73c295364" title="image.png" style="background-color: transparent; color: var(--tertiary); text-decoration: none; cursor: pointer; outline: 0px; overflow-wrap: break-word; position: relative; display: inline-block; overflow: hidden; transition: all 0.6s cubic-bezier(0.165, 0.84, 0.44, 1) 0s;"><img src="https://forum.duplicati.com/uploads/default/optimized/1X/fa3444a26b828f508b9ebf3fbd8fe0b73c295364_2_690x254.png" alt="image" width="690" height="254" srcset="https://forum.duplicati.com/uploads/default/optimized/1X/fa3444a26b828f508b9ebf3fbd8fe0b73c295364_2_690x254.png, https://forum.duplicati.com/uploads/default/original/1X/fa3444a26b828f508b9ebf3fbd8fe0b73c295364.png 1.5x, https://forum.duplicati.com/uploads/default/original/1X/fa3444a26b828f508b9ebf3fbd8fe0b73c295364.png 2x" data-small-upload="https://forum.duplicati.com/uploads/default/optimized/1X/fa3444a26b828f508b9ebf3fbd8fe0b73c295364_2_10x10.png" loading="lazy" style="border-style: none; vertical-align: middle; outline: 0px; object-fit: cover; object-position: center top; max-width: 100%; height: auto; aspect-ratio: 690 / 254;"><div class="meta" style="color: var(--secondary); font-weight: bold; outline: 0px; position: absolute; bottom: 0px; width: 652.5px; background: var(--primary); opacity: 0; transition: opacity 0.2s ease 0s; display: flex; align-items: center;"><svg class="fa d-icon d-icon-far-image svg-icon" aria-hidden="true"><use xlink:href="#far-image"></use></svg><span class="filename" style="outline: 0px; margin: 6px 6px 6px 0px; overflow: hidden; white-space: nowrap; text-overflow: ellipsis;">image.png</span><span class="informations" style="outline: 0px; margin: 6px; padding-right: 20px; color: var(--secondary-high); font-size: var(--font-0); flex-shrink: 0; flex-grow: 3;">959×354 17.8 KB</span><svg class="fa d-icon d-icon-discourse-expand svg-icon" aria-hidden="true"><use xlink:href="#discourse-expand"></use></svg></div></a></div><p style="display: block;"></p></details>

- Copy the text block below, and paste it into the “Options” text box

  <details open="" style="display: block; position: relative; outline: none;"><summary style="display: block; outline: none; cursor: pointer;"><span>&nbsp;</span>Text block you need to copy</summary><p style="display: block;"><code style="font-family: Consolas, Menlo, Monaco, &quot;Lucida Console&quot;, &quot;Liberation Mono&quot;, &quot;DejaVu Sans Mono&quot;, &quot;Bitstream Vera Sans Mono&quot;, &quot;Courier New&quot;, monospace; font-size: 1em; color: var(--primary-very-high); background: var(--hljs-bg);">--</code>send-mail-url=smtp://smtp.gmail.com:587/?starttls=when-available<br><code style="font-family: Consolas, Menlo, Monaco, &quot;Lucida Console&quot;, &quot;Liberation Mono&quot;, &quot;DejaVu Sans Mono&quot;, &quot;Bitstream Vera Sans Mono&quot;, &quot;Courier New&quot;, monospace; font-size: 1em; color: var(--primary-very-high); background: var(--hljs-bg);">--</code>send-mail-any-operation=true<br><code style="font-family: Consolas, Menlo, Monaco, &quot;Lucida Console&quot;, &quot;Liberation Mono&quot;, &quot;DejaVu Sans Mono&quot;, &quot;Bitstream Vera Sans Mono&quot;, &quot;Courier New&quot;, monospace; font-size: 1em; color: var(--primary-very-high); background: var(--hljs-bg);">--</code>send-mail-subject=Duplicati %PARSEDRESULT%, %OPERATIONNAME% report for %backup-name%<br><code style="font-family: Consolas, Menlo, Monaco, &quot;Lucida Console&quot;, &quot;Liberation Mono&quot;, &quot;DejaVu Sans Mono&quot;, &quot;Bitstream Vera Sans Mono&quot;, &quot;Courier New&quot;, monospace; font-size: 1em; color: var(--primary-very-high); background: var(--hljs-bg);">--</code>send-mail-to=<em><a href="mailto:destination_email_address@whatever.com" style="background-color: transparent; color: var(--tertiary); text-decoration: none; cursor: pointer; overflow-wrap: break-word;">destination_email_address@whatever.com</a></em><br><code style="font-family: Consolas, Menlo, Monaco, &quot;Lucida Console&quot;, &quot;Liberation Mono&quot;, &quot;DejaVu Sans Mono&quot;, &quot;Bitstream Vera Sans Mono&quot;, &quot;Courier New&quot;, monospace; font-size: 1em; color: var(--primary-very-high); background: var(--hljs-bg);">--</code>send-mail-username=<em><a href="mailto:your_sending_gmail_username@gmail.com" style="background-color: transparent; color: var(--tertiary); text-decoration: none; cursor: pointer; overflow-wrap: break-word;">your_sending_gmail_username@gmail.com</a></em><br><code style="font-family: Consolas, Menlo, Monaco, &quot;Lucida Console&quot;, &quot;Liberation Mono&quot;, &quot;DejaVu Sans Mono&quot;, &quot;Bitstream Vera Sans Mono&quot;, &quot;Courier New&quot;, monospace; font-size: 1em; color: var(--primary-very-high); background: var(--hljs-bg);">--</code>send-mail-password=<strong style="font-weight: bolder;"><em>your_sending_gmail_password</em></strong><br><code style="font-family: Consolas, Menlo, Monaco, &quot;Lucida Console&quot;, &quot;Liberation Mono&quot;, &quot;DejaVu Sans Mono&quot;, &quot;Bitstream Vera Sans Mono&quot;, &quot;Courier New&quot;, monospace; font-size: 1em; color: var(--primary-very-high); background: var(--hljs-bg);">--</code>send-mail-from=<strong style="font-weight: bolder;"><em>This_computers_name</em></strong><span>&nbsp;</span>backup<span>&nbsp;</span><em>&lt;<a href="mailto:your_sending_gmail_username@gmail.com" style="background-color: transparent; color: var(--tertiary); text-decoration: none; cursor: pointer; overflow-wrap: break-word;">your_sending_gmail_username@gmail.com</a>&gt;</em><br><code style="font-family: Consolas, Menlo, Monaco, &quot;Lucida Console&quot;, &quot;Liberation Mono&quot;, &quot;DejaVu Sans Mono&quot;, &quot;Bitstream Vera Sans Mono&quot;, &quot;Courier New&quot;, monospace; font-size: 1em; color: var(--primary-very-high); background: var(--hljs-bg);"></code><br><code style="font-family: Consolas, Menlo, Monaco, &quot;Lucida Console&quot;, &quot;Liberation Mono&quot;, &quot;DejaVu Sans Mono&quot;, &quot;Bitstream Vera Sans Mono&quot;, &quot;Courier New&quot;, monospace; font-size: 1em; color: var(--primary-very-high); background: var(--hljs-bg);"></code></p></details>

  <details open="" style="display: block; position: relative; outline: none;"><summary style="display: block; outline: none; cursor: pointer;"><span>&nbsp;</span>Screenshot 3, showing location of where to paste the copied text</summary><p style="display: block;"></p><div class="lightbox-wrapper" style="display: block; outline: 0px;"><a class="lightbox" href="https://forum.duplicati.com/uploads/default/original/1X/8a6f3b6f37eeb26bbe890e704352534671b43b61.png" data-download-href="https://forum.duplicati.com/uploads/default/8a6f3b6f37eeb26bbe890e704352534671b43b61" title="image.png" style="background-color: transparent; color: var(--tertiary); text-decoration: none; cursor: pointer; outline: 0px; overflow-wrap: break-word; position: relative; display: inline-block; overflow: hidden; transition: all 0.6s cubic-bezier(0.165, 0.84, 0.44, 1) 0s;"><img src="https://forum.duplicati.com/uploads/default/optimized/1X/8a6f3b6f37eeb26bbe890e704352534671b43b61_2_690x336.png" alt="image" width="690" height="336" srcset="https://forum.duplicati.com/uploads/default/optimized/1X/8a6f3b6f37eeb26bbe890e704352534671b43b61_2_690x336.png, https://forum.duplicati.com/uploads/default/optimized/1X/8a6f3b6f37eeb26bbe890e704352534671b43b61_2_1035x504.png 1.5x, https://forum.duplicati.com/uploads/default/original/1X/8a6f3b6f37eeb26bbe890e704352534671b43b61.png 2x" data-small-upload="https://forum.duplicati.com/uploads/default/optimized/1X/8a6f3b6f37eeb26bbe890e704352534671b43b61_2_10x10.png" loading="lazy" style="border-style: none; vertical-align: middle; outline: 0px; object-fit: cover; object-position: center top; max-width: 100%; height: auto; aspect-ratio: 690 / 336;"><div class="meta" style="color: var(--secondary); font-weight: bold; outline: 0px; position: absolute; bottom: 0px; width: 652.5px; background: var(--primary); opacity: 0; transition: opacity 0.2s ease 0s; display: flex; align-items: center;"><svg class="fa d-icon d-icon-far-image svg-icon" aria-hidden="true"><use xlink:href="#far-image"></use></svg><span class="filename" style="outline: 0px; margin: 6px 6px 6px 0px; overflow: hidden; white-space: nowrap; text-overflow: ellipsis;">image.png</span><span class="informations" style="outline: 0px; margin: 6px; padding-right: 20px; color: var(--secondary-high); font-size: var(--font-0); flex-shrink: 0; flex-grow: 3;">1045×509 18.1 KB</span><svg class="fa d-icon d-icon-discourse-expand svg-icon" aria-hidden="true"><use xlink:href="#discourse-expand"></use></svg></div></a></div><p style="display: block;"></p></details>

- Click on the “Edit as list” button

  <details open="" style="display: block; position: relative; outline: none;"><summary style="display: block; outline: none; cursor: pointer;"><span>&nbsp;</span>Screenshot 4</summary><p style="display: block;"></p><div class="lightbox-wrapper" style="display: block; outline: 0px;"><a class="lightbox" href="https://forum.duplicati.com/uploads/default/original/1X/54c6e7d40a9509e278f1d149455c560a37bed54b.png" data-download-href="https://forum.duplicati.com/uploads/default/54c6e7d40a9509e278f1d149455c560a37bed54b" title="image.png" style="background-color: transparent; color: var(--tertiary); text-decoration: none; cursor: pointer; outline: 0px; overflow-wrap: break-word; position: relative; display: inline-block; overflow: hidden; transition: all 0.6s cubic-bezier(0.165, 0.84, 0.44, 1) 0s;"><img src="https://forum.duplicati.com/uploads/default/optimized/1X/54c6e7d40a9509e278f1d149455c560a37bed54b_2_690x225.png" alt="image" width="690" height="225" srcset="https://forum.duplicati.com/uploads/default/optimized/1X/54c6e7d40a9509e278f1d149455c560a37bed54b_2_690x225.png, https://forum.duplicati.com/uploads/default/original/1X/54c6e7d40a9509e278f1d149455c560a37bed54b.png 1.5x, https://forum.duplicati.com/uploads/default/original/1X/54c6e7d40a9509e278f1d149455c560a37bed54b.png 2x" data-small-upload="https://forum.duplicati.com/uploads/default/optimized/1X/54c6e7d40a9509e278f1d149455c560a37bed54b_2_10x10.png" loading="lazy" style="border-style: none; vertical-align: middle; outline: 0px; object-fit: cover; object-position: center top; max-width: 100%; height: auto; aspect-ratio: 690 / 225;"><div class="meta" style="color: var(--secondary); font-weight: bold; outline: 0px; position: absolute; bottom: 0px; width: 652.5px; background: var(--primary); opacity: 0; transition: opacity 0.2s ease 0s; display: flex; align-items: center;"><svg class="fa d-icon d-icon-far-image svg-icon" aria-hidden="true"><use xlink:href="#far-image"></use></svg><span class="filename" style="outline: 0px; margin: 6px 6px 6px 0px; overflow: hidden; white-space: nowrap; text-overflow: ellipsis;">image.png</span><span class="informations" style="outline: 0px; margin: 6px; padding-right: 20px; color: var(--secondary-high); font-size: var(--font-0); flex-shrink: 0; flex-grow: 3;">968×316 21.3 KB</span><svg class="fa d-icon d-icon-discourse-expand svg-icon" aria-hidden="true"><use xlink:href="#discourse-expand"></use></svg></div></a></div><p style="display: block;"></p></details>

You will now see 7 advanced options.

<details open="" style="display: block; position: relative; outline: none;"><summary style="display: block; outline: none; cursor: pointer;"><span>&nbsp;</span>Screenshot 5</summary><p style="display: block;"></p><div class="lightbox-wrapper" style="display: block; outline: 0px;"><a class="lightbox" href="https://forum.duplicati.com/uploads/default/original/1X/3757dba33ca5f1e99e61c8c3e3f8da00bea8b1e6.png" data-download-href="https://forum.duplicati.com/uploads/default/3757dba33ca5f1e99e61c8c3e3f8da00bea8b1e6" title="image.png" style="background-color: transparent; color: var(--tertiary); text-decoration: none; cursor: pointer; outline: 0px; overflow-wrap: break-word; position: relative; display: inline-block; overflow: hidden; transition: all 0.6s cubic-bezier(0.165, 0.84, 0.44, 1) 0s;"><img src="https://forum.duplicati.com/uploads/default/optimized/1X/3757dba33ca5f1e99e61c8c3e3f8da00bea8b1e6_2_539x500.png" alt="image" width="539" height="500" srcset="https://forum.duplicati.com/uploads/default/optimized/1X/3757dba33ca5f1e99e61c8c3e3f8da00bea8b1e6_2_539x500.png, https://forum.duplicati.com/uploads/default/optimized/1X/3757dba33ca5f1e99e61c8c3e3f8da00bea8b1e6_2_808x750.png 1.5x, https://forum.duplicati.com/uploads/default/original/1X/3757dba33ca5f1e99e61c8c3e3f8da00bea8b1e6.png 2x" data-small-upload="https://forum.duplicati.com/uploads/default/optimized/1X/3757dba33ca5f1e99e61c8c3e3f8da00bea8b1e6_2_10x10.png" loading="lazy" style="border-style: none; vertical-align: middle; outline: 0px; object-fit: cover; object-position: center top; max-width: 100%; height: auto; aspect-ratio: 539 / 500;"><div class="meta" style="color: var(--secondary); font-weight: bold; outline: 0px; position: absolute; bottom: 0px; width: 539px; background: var(--primary); opacity: 0; transition: opacity 0.2s ease 0s; display: flex; align-items: center;"><svg class="fa d-icon d-icon-far-image svg-icon" aria-hidden="true"><use xlink:href="#far-image"></use></svg><span class="filename" style="outline: 0px; margin: 6px 6px 6px 0px; overflow: hidden; white-space: nowrap; text-overflow: ellipsis;">image.png</span><span class="informations" style="outline: 0px; margin: 6px; padding-right: 20px; color: var(--secondary-high); font-size: var(--font-0); flex-shrink: 0; flex-grow: 3;">816×756 65.1 KB</span><svg class="fa d-icon d-icon-discourse-expand svg-icon" aria-hidden="true"><use xlink:href="#discourse-expand"></use></svg></div></a></div><p style="display: block;"></p></details>

The first 3 do not need to be modified. The last 4 need to be modified as follows:

- Replace *[destination_email_address@whatever.com](mailto:destination_email_address@whatever.com)* with the email address you want the status message sent to.
- Replace *[your_sending_gmail_username@gmail.com](mailto:your_sending_gmail_username@gmail.com)* with the email address of the gmail account you are using to send the message. (Note that there are two places where you need to make this change)
- Replace ***your_sending_gmail_password\*** with the password of the gmail account you are using to send the message.
- Replace ***This_computers_name\*** with the name of the computer that the backup job is running on.

Click the “OK” button at the bottom of the screen to save your results.

You will now receive an email message at the conclusion of every Duplicati job that is run on the current computer.

This can generate a lot of email clutter – a report for every backup job run. There are two approaches to reducing the volume of email received.

My recommend approach is using [dupReport.py 346](https://forum.duplicati.com/t/announcing-dupreport-a-duplicati-email-report-summary-generator/1116) to generate summary reports. This excellent tool was developed by [@handyguy](https://forum.duplicati.com/u/handyguy). If you do this, be sure to modify the “subjectregex=” line in dupReport.rc as shown below to ensure the summary reports work for either the default subject line or the subject line used in this writeup.

```
subjectregex = ^Duplicati ([\w ]*, |)Backup report for
```

You also need to add a filter to your email client to automatically move the duplicati-generated email’s out of your inbox and into an email folder which dupReport.py can process.

Alternatively, as suggested by [@sanderson](https://forum.duplicati.com/u/sanderson), below, you can add this setting:
`--send-mail-level=Warning,Error,Fatal`

The one downside to this approach is that it doesn’t differentiate between a backup job that succeeded and one that never ran, as the result of both conditions is that no email is sent. This is why I prefer the dupReport.py approach.