# This repository is under GNU General Public License v3.0.
# see https://github.com/Fricciolosa-Red-Team/crch/blob/main/LICENSE

echo ""
echo " => Continuous Recon Continuous Hacking <= "
echo ""
sudo apt update && sudo apt upgrade && sudo apt autoremove
sudo apt install golang
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
cd utils
go build -o addpro
sudo mv addpro /usr/bin
cd ~
git clone https://github.com/projectdiscovery/nuclei-templates.git
git clone https://github.com/edoardottt/scilla.git
cd scilla
go get
make linux
cd ~
echo ""
echo "[ !!!! ] Remember to: [ !!!! ]"
echo ""
echo "    > set ~/.config/notify/notify.conf !"
echo "    > add the go/bin folder to your PATH system variable !"
echo ""
