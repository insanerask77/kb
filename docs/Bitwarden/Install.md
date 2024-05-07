1. Download the Bitwarden installation script (`bitwarden.sh`) to your machine:

   ```
   curl -Lso bitwarden.sh https://go.btwrdn.co/bw-sh && chmod 700 bitwarden.sh
   ```

2. Run the installer script. A `./bwdata` directory will be created relative to the location of `bitwarden.sh`.

   ```
   ./bitwarden.sh install
   ```

3. Complete the prompts in the installer:

   - **Enter the domain name for your Bitwarden instance:**

     Typically, this value should be the configured DNS record.

   - **Do you want to use Let's Encrypt to generate a free SSL certificate? (y/n):**

     Specify `y` to generate a trusted SSL certificate using Let's Encrypt. You will be prompted to enter an email address for expiration reminders from Let's Encrypt. For more information, see [Certificate Options](https://bitwarden.com/help/certificates/).

     Alternatively, specify `n` and use the **Do you have a SSL certificate to use?** option.

   - **Enter your installation id:**

     Retrieve an installation id using a valid email at [https://bitwarden.com/host](https://bitwarden.com/host/). For more information, see [What are my installation id and installation key used for?](https://bitwarden.com/help/hosting-faqs/#general).

   - **Enter your installation key:**

     Retrieve an installation key using a valid email at [https://bitwarden.com/host](https://bitwarden.com/host/). For more information, see [What are my installation id and installation key used for?](https://bitwarden.com/help/hosting-faqs/#general).

   - **Do you have a SSL certificate to use? (y/n):**

     If you already have your own SSL certificate, specify `y` and place the necessary files in the `./bwdata/ssl/your.domain` directory. You will be asked whether it is a trusted SSL certificate (y/n). For more information, see [Certificate Options](https://bitwarden.com/help/certificates/).

     Alternatively, specify `n` and use the **self-signed SSL certificate?** option, which is only recommended for testing purposes.

   - **Do you want to generate a self-signed SSL certificate? (y/n):**

     Specify `y` to have Bitwarden generate a self-signed certificate for you. This option is only recommended for testing. For more information, see [Certificate Options](https://bitwarden.com/help/certificates/).

     If you specify `n`, your instance will not use an SSL certificate and you will be required to front your installation with a HTTPS proxy, or else Bitwarden applications will not function properly.

[](https://bitwarden.com/help/install-on-premise-linux/#post-install-configuration)

### Post-Install Configuration

Configuring your environment can involve making changes to two files; an [environment variables file](https://bitwarden.com/help/install-on-premise-linux/#environment-variables) and an [installation file](https://bitwarden.com/help/install-on-premise-linux/#installation-configuration):

#### Environment Variables (*Required*)

Some features of Bitwarden are not configured by the `bitwarden.sh` script. Configure these settings by editing the environment file, located at `./bwdata/env/global.override.env`. **At a minimum, you should replace the values for:**

```
...
globalSettings__mail__smtp__host=<placeholder>
globalSettings__mail__smtp__port=<placeholder>
globalSettings__mail__smtp__ssl=<placeholder>
globalSettings__mail__smtp__username=<placeholder>
globalSettings__mail__smtp__password=<placeholder>
...
adminSettings__admins=
...
```

Replacing `globalSettings__mail__smtp...=` placeholder will configure the SMTP Mail Server that will be used to send verification emails to new users and invitations to Organizations. Adding an email address to `adminSettings__admins=` will provision access to the Admin Portal.

After editing `global.override.env`, run the following command to apply your changes:

```
./bitwarden.sh restart
```

#### Installation File

The Bitwarden installation script uses settings in `./bwdata/config.yml` to generate the necessary assets for installation. Some installation scenarios (e.g. installations behind a proxy with alternate ports) may require adjustments to `config.yml` that were not provided during standard installation.

Edit `config.yml` as necessary and apply your changes by running:

```
./bitwarden.sh rebuild
```

[](https://bitwarden.com/help/install-on-premise-linux/#start-bitwarden)

### Start Bitwarden

Once you've completed all previous steps, start your Bitwarden instance:

```
./bitwarden.sh start
```