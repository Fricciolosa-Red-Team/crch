while :
do  
    OLD=$(pwd)
    cd ~/nuclei-templates
    git pull
    cd $OLD
    nuclei -t ~/nuclei-templates/takeovers -l subdomains.txt -silent | anew takeovers.txt | notify

    # this every 6 hours
    sleep 21600

done