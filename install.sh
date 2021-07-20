# see https://github.com/Fricciolosa-Red-Team/crch/blob/main/LICENSE

sudo apt update && sudo apt upgrade && sudo apt autoremove
sudo apt install golang python3
GO111MODULE=on go get -v github.com/projectdiscovery/notify/cmd/notify
GO111MODULE=on go get -v github.com/projectdiscovery/notify/cmd/intercept
go get -u github.com/tomnomnom/anew
GO111MODULE=on go get -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder
wget https://github.com/findomain/findomain/releases/latest/download/findomain-linux
chmod +x findomain-linux
mv findomain-linux findomain
sudo mv findomain /usr/bin
go get -u github.com/tomnomnom/assetfinder
GO111MODULE=on go get -v github.com/projectdiscovery/nuclei/v2/cmd/nuclei
go build -o bbtargets
sudo mv bbtargets /usr/bin
cd ~
git clone https://github.com/projectdiscovery/nuclei-templates.git
echo ""
echo "[ !!!! ] Remember to set ~/.config/notify/notify.conf [ !!!! ]"
echo ""
