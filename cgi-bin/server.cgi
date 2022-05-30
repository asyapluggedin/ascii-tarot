#!/usr/bin/perl
use CGI::Log;

$card = int(rand(77));
Log->debug("card: $card");
my @files = glob "/tmp/tarot/${card}_*.txt";

open my $fh, $files[0] or die "$! opening $files[0]";
print "Content-type: text/html\n\n";
while (my $row = <$fh>) {
    chomp $row;
    print "$row\n";
}
1;
