## Docker Macos

| Nov 30, 2022| [Linux](https://christitus.com/categories/linux) [Windows](https://christitus.com/categories/windows) [MacOS](https://christitus.com/categories/macos)

Don’t have a Mac and need MacOS? No problem, run it in a docker container.

Credit goes to this twitter user:

<iframe id="twitter-widget-0" scrolling="no" frameborder="0" allowtransparency="true" allowfullscreen="true" class="" title="Twitter Tweet" src="https://platform.twitter.com/embed/Tweet.html?dnt=false&amp;embedId=twitter-widget-0&amp;features=eyJ0ZndfdGltZWxpbmVfbGlzdCI6eyJidWNrZXQiOltdLCJ2ZXJzaW9uIjpudWxsfSwidGZ3X2ZvbGxvd2VyX2NvdW50X3N1bnNldCI6eyJidWNrZXQiOnRydWUsInZlcnNpb24iOm51bGx9LCJ0ZndfdHdlZXRfZWRpdF9iYWNrZW5kIjp7ImJ1Y2tldCI6Im9uIiwidmVyc2lvbiI6bnVsbH0sInRmd19yZWZzcmNfc2Vzc2lvbiI6eyJidWNrZXQiOiJvbiIsInZlcnNpb24iOm51bGx9LCJ0Zndfc2hvd19idXNpbmVzc192ZXJpZmllZF9iYWRnZSI6eyJidWNrZXQiOiJvbiIsInZlcnNpb24iOm51bGx9LCJ0ZndfbWl4ZWRfbWVkaWFfMTU4OTciOnsiYnVja2V0IjoidHJlYXRtZW50IiwidmVyc2lvbiI6bnVsbH0sInRmd19leHBlcmltZW50c19jb29raWVfZXhwaXJhdGlvbiI6eyJidWNrZXQiOjEyMDk2MDAsInZlcnNpb24iOm51bGx9LCJ0ZndfZHVwbGljYXRlX3NjcmliZXNfdG9fc2V0dGluZ3MiOnsiYnVja2V0Ijoib24iLCJ2ZXJzaW9uIjpudWxsfSwidGZ3X3ZpZGVvX2hsc19keW5hbWljX21hbmlmZXN0c18xNTA4MiI6eyJidWNrZXQiOiJ0cnVlX2JpdHJhdGUiLCJ2ZXJzaW9uIjpudWxsfSwidGZ3X3Nob3dfYmx1ZV92ZXJpZmllZF9iYWRnZSI6eyJidWNrZXQiOiJvbiIsInZlcnNpb24iOm51bGx9LCJ0ZndfbGVnYWN5X3RpbWVsaW5lX3N1bnNldCI6eyJidWNrZXQiOnRydWUsInZlcnNpb24iOm51bGx9LCJ0Zndfc2hvd19nb3ZfdmVyaWZpZWRfYmFkZ2UiOnsiYnVja2V0Ijoib24iLCJ2ZXJzaW9uIjpudWxsfSwidGZ3X3Nob3dfYnVzaW5lc3NfYWZmaWxpYXRlX2JhZGdlIjp7ImJ1Y2tldCI6Im9uIiwidmVyc2lvbiI6bnVsbH0sInRmd190d2VldF9lZGl0X2Zyb250ZW5kIjp7ImJ1Y2tldCI6Im9uIiwidmVyc2lvbiI6bnVsbH19&amp;frame=false&amp;hideCard=false&amp;hideThread=false&amp;id=1591779981390413826&amp;lang=en&amp;origin=https%3A%2F%2Fchristitus.com%2Fdocker-macos%2F&amp;sessionId=7372c3a918a7a8e8a19e0c37e5733a9bbeca1cad&amp;theme=light&amp;widgetsVersion=aaf4084522e3a%3A1674595607486&amp;width=550px" data-tweet-id="1591779981390413826" style="box-sizing: border-box; position: static; visibility: visible; width: 550px; height: 832px; display: block; flex-grow: 1;"></iframe>

## Requirements

Install Docker Ubuntu 22.04

```fallback
sudo apt install apt-transport-https ca-certificates curl software-properties-common -y
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu jammy stable"
sudo apt install docker-ce -y
sudo usermod -aG docker $USER
```

Copy

**Reboot or logout/login**

(Optional) GUI Webpage for Managing Docker - Portainer

```fallback
docker volume create portainer_data
docker run -d -p 8000:8000 -p 9443:9443 --name portainer \
--restart=always \
-v /var/run/docker.sock:/var/run/docker.sock \
-v portainer_data:/data \
portainer/portainer-ce:2.9.3
```

Copy

Open up browser and navigate to [https://localhost:9443](https://localhost:9443/)

*Click Advanced and proceed with any certificate errors*

## MacOS Docker Setup

Setup the MacOS Docker Container with the following docker command

```fallback
docker run -it \
    --device /dev/kvm \
    -p 50922:10022 \
    -v /tmp/.X11-unix:/tmp/.X11-unix \
    -e "DISPLAY=${DISPLAY:-:0.0}" \
    -e GENERATE_UNIQUE=true \
    -e MASTER_PLIST_URL='https://raw.githubusercontent.com/sickcodes/osx-serial-generator/master/config-custom.plist' \
    sickcodes/docker-osx:monterey

# docker build -t docker-osx --build-arg SHORTNAME=monterey .
```

Copy

Use Disk Utility to “erase” the 270GB virtual disk: *Note: This is just virtual and doesn’t erase your drive*

![img](https://christitus.com/images/2022/docker-macos/disk-util.png)

## Start MacOS Docker Container

Before we start the container find the name with:

```fallback
docker ps -a
```

Copy

*Look for the NAMES column and pick the container name.*

![img](https://christitus.com/images/2022/docker-macos/container.png)

Start with the following command NAME = Name from column above

```fallback
docker start NAME
```

Copy

### Portainer Method for Starting

I love portainer because you can easily manage your containers. Start, Stop, and see resource usage… Portainer does it all! Here is what mine looks like:

![img](https://christitus.com/images/2022/docker-macos/portainer.png)

## Optimize the Container

Source: https://github.com/sickcodes/osx-optimizer

Run the following from Root Prompt # `sudo su`

```fallback
defaults write com.apple.loginwindow autoLoginUser -bool true
mdutil -i off -a
nvram boot-args="serverperfmode=1 $(nvram boot-args 2>/dev/null | cut -f 2-)"
defaults write /Library/Preferences/com.apple.loginwindow DesktopPicture ""
defaults write com.apple.Accessibility DifferentiateWithoutColor -int 1
defaults write com.apple.Accessibility ReduceMotionEnabled -int 1
defaults write com.apple.universalaccess reduceMotion -int 1
defaults write com.apple.universalaccess reduceTransparency -int 1
defaults write com.apple.Accessibility ReduceMotionEnabled -int 1
defaults write /Library/Preferences/com.apple.SoftwareUpdate AutomaticDownload -bool false
defaults write com.apple.SoftwareUpdate AutomaticCheckEnabled -bool false
defaults write com.apple.commerce AutoUpdate -bool false
defaults write com.apple.commerce AutoUpdateRestartRequired -bool false
defaults write com.apple.SoftwareUpdate ConfigDataInstall -int 0
defaults write com.apple.SoftwareUpdate CriticalUpdateInstall -int 0
defaults write com.apple.SoftwareUpdate ScheduleFrequency -int 0
defaults write com.apple.SoftwareUpdate AutomaticDownload -int 0
defaults write com.apple.loginwindow DisableScreenLock -bool true
defaults write com.apple.loginwindow TALLogoutSavesState -bool false
```

Copy

## Final Result

![img](https://christitus.com/images/2022/docker-macos/macos-final.png)

## Walkthrough Video

<iframe src="https://www.youtube.com/embed/XWo2gnNbeGQ" allowfullscreen="" title="YouTube Video" style="box-sizing: border-box; position: absolute; top: 0px; left: 0px; width: 650px; height: 365.625px; border: 0px;"></iframe>

[#Docker](https://christitus.com/tags/docker) [#Ubuntu](https://christitus.com/tags/ubuntu)