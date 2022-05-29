cp cgi-bin/server.cgi /usr/lib/cgi-bin/
chmod +x /usr/lib/cgi-bin/server.cgi
chmod 755 /usr/lib/cgi-bin/server.cgi
mkdir -p /srv/tarot
cp -r txt/*.txt /srv/tarot
chmod 755 /srv/tarot/*
apachectl -k restart
