# IoT Central CLI

IoT Central CLI can be used to manage apps and migrate assets between apps.
It uses IoT Central public API using APP Tokens.

The source code in this tool can be use modified to meet appropriate needs. 

## Build

Follow the instructions to [install Go](https://golang.org/doc/install). Pick the appropriate package to install the latest 1.15.x release of Go. This will give you access to the Go toolchain and compiler.

- If you installed via the tarball, you will need to add a GOROOT environment variable pointing to the folder where you installed Go (typically /usr/local/go on linux-based systems)
- You should also check to make sure that you can access the Go compiler and tools. They are available at $GOROOT/bin (or $GOROOT\bin) and should be added to your path if they are not already. You can verify this by running the following:
  - Max/Linux: which go
  - Windows (CMD): where go

To build this tool, follow the instructions below:

```
To build your repo (__YOUR_REPO__):
$ cd __YOUR_REPO__/src
$ make
```
You can find the executable (iotc for linux or iotc.exe for windows) under __YOUR_REPO__/bin folder

To do a clean build i.e. remove all compiled binaries and dependencies:
```
$ cd __YOUR_REPO__/src
$ make clean
$ make
``` 

## Configuration
This CLI uses a configuration file `$HOME/.iotc.yml`.
The format of the file is as follows:
```
# iotc configuration file
apps:
    # name that is used in the iotc commands 
  - name: myapp1
    # subdomain of the app. For https://app1.azureiotcentral.com, just enter app1
    subdomain: app1
    # domain of the app. Default is azureiotcentral.com
    domain: azureiotcentral.com
    # apitoken can be obtained from your IoT Central application Administration/API tokens
    apiToken: api1Token
  - name: myapp2
    subdomain: app2
    apiToken: api2Token
  - name: myapp3
    subdomain: app3
    apiToken: api3Token

# Maximum number of rows to retrieve for 'list' commands
maxRows: 25

# Default format of the tables
format: pretty
```

## Usage
Once you add the configuration file mentioned above, you can run the tool using `__YOUR_REPO__/bin/iotc` command.
```
D:\dev\iotcgo\bin>iotc

    ____    ______   ______           __             __
   /  _/___/_  __/  / ____/__  ____  / /__________  / /
   / // __ \/ /    / /   / _ \/ __ \/ __/ ___/ __ \/ /
 _/ // /_/ / /    / /___/  __/ / / / /_/ /  / /_/ / /
/___/\____/_/     \____/\___/_/ /_/\__/_/   \__,_/_/

This tool helps you manage IoT Central applications through simple CLI
commands. You can explore (list, upload, download) several entities such
as devicesTemplates, roles, users in an IoT Central applications.
You can copy these entities between different apps. You an operate against
each of these entities below using sub-commands underneath each command.

Usage:
iotc [command]

Available Commands:
actions         Get the actions used in rules
apiTokens       Create, read, delete access tokens used to interact with the IoT Central public APIs
apps            Create, update and delete  IoT Central applications
cde             Manage data exports within your IoT Central application
completion      Generate command line completion script
deviceGroups    Operate against IoT Central device groups
deviceTemplates Create, read, and delete device templates within an IoT Central application
devices         Get information about and manage devices and IoT Edge modules in your IoT Central application.
help            Help about any command
jobs            Operate against IoT Central jobs
roles           List roles within your application
rules           List rules within your application
users           Add, update, and remove users within your application

Flags:
--config string   Command config file (default is $HOME/.iotc.yml)
-h, --help            help for iotc

Use "iotc [command] --help" for more information about a command.
```
