cp cgi-bin/server.cgi /Library/WebServer/CGI-Executables/
chmod +x /Library/WebServer/CGI-Executables/server.cgi
chmod 755 /Library/WebServer/CGI-Executables/server.cgi
mkdir -p /tmp/tarot
cp -r txt/extracted_txt/* /tmp/tarot
chmod 755 /tmp/tarot
apachectl -k restart
