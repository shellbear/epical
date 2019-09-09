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
