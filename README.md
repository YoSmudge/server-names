# Server Naming Thing
## An unfortunately poorly named tool for naming servers

When you're working with many servers at once, having names like prod-web-i-0c9e0a5fe22fca1a8 gets dull fast. It's also hard to communicate in person. What if you could encode the instance ID into something friendly and human readable? Like prod-web-i-disgusted-leonurus?*

So I built this little tool, it generates names by using the original instance ID then encoding that using a list of words. The format is always {adjective or verb}-{noun}. Running the program on an EC2 instance will call the EC2 metadata API to fetch the instance ID and returning the encoded value.

    sam@m-sensu-i-282f2cb4:~$ ./server-names
    intermingle-juno

Or you can provide an ID on the command line

    sam@m-sensu-i-282f2cb4:~$ ./server-names 0c9e0a5fe22fca1a8
    disgusted-leonurus

Pre-compiled binaries are avaliable from the releases tab, or you can build for yourself with the Makefile.

Still somewhat a work in progress, YMMV

*Leonurus, or Leonurus cardiaca, [is an herbaceous perennial plant in the mint family, Lamiaceae. Other common names include throw-wort, lion's ear, and lion's tail](https://en.wikipedia.org/wiki/Leonurus_cardiaca)
