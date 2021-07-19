while :
do  
    OLD=$(pwd)
    cd ~/nuclei-templates
    git pull
    cd $OLD
    # Be aware of this:
    #    - with `| cut -d " " -f 3-` we remove the timestamp section, this means
    #       we will not have duplicates (duplicate only for the timestamp), but we don't have
    #       a time information 
    #    - without that part of course we will have a lot of duplicates (also in term of notifications),
    #       but we will have a time information
    #
    nuclei -t ~/nuclei-templates/takeovers -l subdomains.txt -silent | cut -d " " -f 3- | anew takeovers.txt | notify

    # this every 6 hours
    sleep 21600

done