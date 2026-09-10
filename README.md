# CSMcli/LScli

**CSMcli** is a Command Line Interface designed to bridge the gap between the CSM server, the official Riot Games API, and the League Client Update (LCU) API. 

By handling input sanitization, intelligent caching, and direct communication with the active League of Legends client, CSMcli provides users with a terminal-based League of Legends toolset.

---

## Features

- **search command** Query the CSM server and Riot API to search for your user and gain insight into your ingame statistics.
- **show command** Show analytics for one or more matches based on a matchid.
- **ladder command** Show the current challenger ladder of a server of your choice.
- **build command** Show a champions current best build based off CSM's aggregated data.
- **import command** Imports a runepage for your champion based on either your own preset or aggregated data from the CSM server.
- **setspell command** Sets the summoner spell for your role and champion based on aggregated data from the CSM server.
- **help command** Prints a list of supported commands and a description

## Example Commands

```console
foo@bar:~$ search euw foo bar

foo@bar:~$ show EUW1_12344556

foo@bar:~$ ladder euw

foo@bar:~$ build <champion>

foo@bar:~$ build <champion> <patchid>

foo@bar:~$ import <champion>

foo@bar:~$ setspell <champion>

```

## Prerequisites

To use the CSMcli you must ensure you have the following:

- **Go 1.26+**
- **Active League of Legends Client:** The client must be running to use the import and setspell command.
- **Riot API Key or CSM server installed**: If you wish to run CSMcli without the server follow the steps below. Otherwise running the server will allow you to make request without an API key

## Installation


