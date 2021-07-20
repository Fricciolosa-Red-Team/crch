# crch
**C**ontinuous **R**econ **C**ontinuous **H**acking

## Description

It grabs all the root domains in BB programs (Hackerone, Intigriti, BugCrowd and YesWeHack), then it performes  
subdomains enumeration and check if there are new possible subdomain takeovers with nuclei and nuclei-templates.  
If there are, it will nofity you using one or more than one among Telegram, Discord and Slack.

## Installation

- `./install.sh` (Run without sudo, then it will ask for password)

Dependencies (installed with the `install.sh` script) :
  - [Go](https://go.dev/learn/)
  - [scilla](https://github.com/edoardottt/scilla)
  - [notify](https://github.com/projectdiscovery/notify/cmd/notify)
  - [intercept](https://github.com/projectdiscovery/notify/cmd/intercept)
  - [anew](https://github.com/tomnomnom/anew)
  - [subfinder](https://github.com/projectdiscovery/subfinder/v2/cmd/subfinder)
  - [findomain](https://github.com/findomain/findomain)
  - [assetfinder](https://github.com/tomnomnom/assetfinder)
  - [nuclei](https://github.com/projectdiscovery/nuclei/v2/cmd/nuclei)
  - [nuclei-templates](https://github.com/projectdiscovery/nuclei-templates)

## Usage

- Edit the file `~/.config/notify/notify.conf` with your keys/tokens.
- `./grabTargets.sh`
- `./monitor.sh`

## License

This repository is under [GNU General Public License v3.0](https://github.com/Fricciolosa-Red-Team/crch/blob/main/LICENSE).
