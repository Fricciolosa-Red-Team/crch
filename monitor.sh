while 1:
do  
    bbtargets > targets.txt

    for line in $(cat targets.txt)
    do

        scilla subdomain -db -no-check -target $line -o txt

        subfinder -d $line -o subfinder.txt

        findomain -t $line -u findomain-subs.txt

        assetfinder --subs-only $line > assetfinder-subs.txt

        cat output-scilla/$line.subdomain.txt subfinder.txt findomain-subs.txt assetfinder-subs.txt | sort | uniq > "${line}.subs.txt"

        rm -rf subfinder.txt findomain-subs.txt assetfinder-subs.txt output-scilla/$line.subdomain.txt
    
        cat "${line}.subs.txt" | anew subdomains.txt

    done

    nuclei -t ~/nuclei-templates/takeovers -l subdomains.txt
    
    sleep 3600

done