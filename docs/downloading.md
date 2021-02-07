# Downloading and Installing

You can download pre-compiled binaries from the [releases](https://github.com/muquit/applehealth2csv/releases)
page.  

Please add an [issue](https://github.com/muquit/applehealth2csv/issues) if you would need binaries for any other platforms.

Before installing, please make sure to verify the checksum.

## Installing using Homebrew on Mac

You will need to install [Homebrew](https://brew.sh/) first.

### Install

First install the custom tap.

```
    $ brew tap muquit/applehealth2csv https://github.com/muquit/applehealth2csv.git
    $ brew install applehealth2csv
```

### Uninstall
```
    $ brew uninstall applehealth2csv
```


## Installing the debian package on Ubuntu or Debian or Raspberry pi

### Inspect the package content
```
    $ dpkg -c applehealth2csv_linux_64-bit.deb
```

### Install

```
    $ sudo dpkg -i applehealth2csv_linux_64-bit.deb 
    $ applehealth2csv -h
```

### Uninstall

```
    $ sudo dpkg -r applehealth2csv
```

## Install the RPM package

### Inspect the package content
```
    $ rpm -qlp applehealth2csv_linux_64-bit.rpm
```
### Install/Upgrade
```
    # rpm -Uvh applehealth2csv_linux_64-bit.rpm
    # applehealth2csv -h
```
### Uninstall
```
    # rpm -ev applehealth2csv
```

## Install from archive

### Inspect the content
```
    $ tar -tvf applehealth2csv_x.x.x_linux_64-bit.tar.gz
```

```
    $ unzip -l applehealth2csv_x.x.x_windows_64-bit.zip
```

### Install Linux
```
    $ tar -xf applehealth2csv_x.x.x_linux_64-bit.tar.gz
    $ sudo cp applehealth2csv-dir/applehealth2csv /usr/local/bin
    $ sudo cp applehealth2csv-dir/doc/applehealth2csv.1 /usr/local/share/man/man1
```

### Install Windows

After [downloading](#downloading-and-installing) the latest .zip file (e.g., applehealth2csv_x.x.x_windows_64-bit.zip), unzip it, and copy `applehealth2csv-dir\applehealth2csv.exe` somewhere in your PATH or run it from the directory.
