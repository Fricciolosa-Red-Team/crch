# crch
Continuous Recon Continuous Hacking

## Description

It grabs all the root domains in BB programs (Hackerone, Intigriti, BugCrowd and YesWeHack), then it performes subdomains  
enumeration and check if there are new possible subdomain takeovers with nuclei and nuclei-templates.  
If there are, it will nofity you using one or more than one among Telegram, Discord and Slack.

## Installation

- `./install.sh`

## Usage

- Edit the file `~/.config/notify/notify.conf` with your keys/tokens.
- `./grabTargets.sh`
- `./monitor.sh`

## License

This repository is under [GNU General Public License v3.0](https://github.com/Fricciolosa-Red-Team/crch/blob/main/LICENSE).
