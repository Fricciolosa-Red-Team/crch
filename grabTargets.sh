# This repository is under GNU General Public License v3.0.
# see https://github.com/Fricciolosa-Red-Team/crch/blob/main/LICENSE

# ! it takes ~ 6-7 hours to complete !


# grab all bug bounty 2nd level domains
bbtargets > targets.txt

# for every target grab subdomains
for line in $(cat targets.txt)
do

    echo "[ + ] $line"

    scilla subdomain -db -no-check -target $line -o txt

    subfinder -d $line -o subfinder.txt

    findomain -t $line -u findomain-subs.txt

    assetfinder --subs-only $line > assetfinder-subs.txt

    cat output-scilla/$line.subdomain.txt subfinder.txt findomain-subs.txt assetfinder-subs.txt | sort | uniq > "${line}.subs.txt"

    rm -rf subfinder.txt findomain-subs.txt assetfinder-subs.txt output-scilla/$line.subdomain.txt
    
    # and if there are new subs add them to the proper file.
    cat "${line}.subs.txt" | grep -v "*" | addpro | anew subdomains.txt

    rm -rf "${line}.subs.txt"

done
    
rm -rf targets.txt
