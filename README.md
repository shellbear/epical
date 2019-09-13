# EpiCal
A simple tool to synchronize your Epitech events with Google calendar.

# Requirements

#### Epitech token
To run this script you have to use your Epitech auto login token which can be found [here](https://intra.epitech.eu/admin/autolog).
The URL is of the form `https://intra.epitech.eu/auth-XXX...`
The token used by the script is only the right part of the URL after the `auth-` part.

#### Google calendar credentials

You have to enable the Google Calendar API for your account.
You can follow the [Google developer QuickStart](https://developers.google.com/calendar/quickstart/go) and click **ENABLE THE GOOGLE CALENDAR API** button.
Then in the resulting dialog click **DOWNLOAD CLIENT CONFIGURATION** and save the file credentials.json to your working directory.

## Installation

#### With go CLI
```bash
go get -v -u github.com/ShellBear/epical/...
epical version
```

#### Direct download

You can also download the binaries from the **Release** page: https://github.com/ShellBear/epical/releases.

## Commands

### Sync
A command to synchronize Epitech events with your Google calendar. 
```bash
./epical --token YOUR_EPITECH_AUTOLOGIN_TOKEN sync
```

### List
A command to list Epitech events. 
```bash
./epical --token YOUR_EPITECH_AUTOLOGIN_TOKEN list
```

### Clear
Delete all events previously created by EpiCal.
```bash
./epical clear
```

## Options

#### Credentials

You can specify the Google Calendar API folder path containing `credentials.json` and `token.json` files using the `--credentials` or `-c` option.

```bash
./epical --token YOUR_EPITECH_AUTOLOGIN_TOKEN --credentials /run/secrets/ sync
```

## Usage

The easiest way to run this tool is to launch it with a periodic periodic job.

First, you need to generate the `token.json` file, which is the Google Calendar API token file.
Make sure you have the file `credentials.json` and then run the script for the first time:

If the `credentials.json` file is in your current directory then you can just do:
```bash
epical --token YOUR_EPITECH_AUTOLOGIN_TOKEN
```

Otherwise, you can specify the path of the folder where it can be found:
```bash
epical --token YOUR_EPITECH_AUTOLOGIN_TOKEN --credentials /run/secrets/
```

Then a message will appear:

```
Go to the following link in your browser then type the authorization code: 
https://accounts.google.com/o/oauth2/auth...
```

Just open this link in a browser, sign in with your google account, and paste the code you received.
Then, the `token.json` will be generated automatically.

Finally, make sure that cron is installed and that the `/ etc / cron.hourly` folder exists.
Just create a script in the `/ etc / cron.hourly` directory or any other cron directory of your choice:

```bash
> cat /etc/cron.hourly/epical
#!/bin/sh

EPICAL_BINARY_PATH/epical --token YOUR_EPITECH_AUTOLOGIN_TOKEN -c CREDENTIALS_AND_TOKEN_FOLDER_PATH sync
```

Then give execute permissions to the file and check that everything is working correctly:

```bash
chmod +x /etc/cron.hourly/epical
run-parts -v /etc/cron.hourly/
```