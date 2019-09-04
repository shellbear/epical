# EpiCal
A simple to synchronize your Epitech events with Google calendar.

# Requirements

To run this script you have to use your Epitech auto login token which can be found [here](https://intra.epitech.eu/admin/autolog).
The URL is of the form `https://intra.epitech.eu/auth-XXX...`
The token used by the script is only the right part of the URL after the `auth-` part.

## Installation
```go
go build .
```

## Commands

### Sync
A command to synchronize Epitech events with your Google calendar. 
```go
./epical --token YOUR_EPITECH_AUTOLOGIN_TOKEN sync
```

### List
A command to list Epitech events. 
```go
./epical --token YOUR_EPITECH_AUTOLOGIN_TOKEN list
```

### Clear
Delete all events previously created by EpiCal.
```go
./epical --token YOUR_EPITECH_AUTOLOGIN_TOKEN clear
```