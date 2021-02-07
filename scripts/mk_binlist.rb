#!/usr/bin/env ruby
#
# muquit@muquit.com Feb-07-2021 
begin
  bins = []
	bins << "| applehealth2csv_TAG_checksums.txt      | SHA256 checksum of binaries|"
  bins << "| applehealth2csv_TAG_linux-64bit.deb    | Debian package Linux|"
	bins << "| applehealth2csv_TAG_linux-64bit.rpm    | RPM package for Linux |"
  bins << "| applehealth2csv_TAG_linux-64bit.tar.gz | Compressed tar archive for Linux|"
  bins << "| applehealth2csv_TAG_linux-ARM.deb      | Debian package for Raspberry pi (32 bit) | "
	bins << "| applehealth2csv_TAG_linux-ARM.rpm      | RPM package for Raspberry pi (32 bit)| "
  bins << "| applehealth2csv_TAG_linux-ARM.tar.gz   | Compressed tar archive for Raspberry pi (32 bit)| "
  bins << "| applehealth2csv_TAG_macOS-64bit.tar.gz | Compressed tar archive for Mac OS X | "
  bins << "| applehealth2csv_TAG_windows-64bit.zip  | zip archive for Windows 64 bit|"
  bins << "| applehealth2csv_TAG_windows-32bit.zip  | zip archive for Windows 32 bit|"
  
  tag=`git describe --abbrev=0 --tags`.chomp
  puts <<EOF

| Files | Platform |
| :-------| :--------|
EOF
	bins.each do |bin|
    bin = bin.gsub("TAG",tag)
    puts "#{bin}"

	end
rescue => e
  puts "ERROR: #{e}"
end
