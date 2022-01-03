# ipcalc-contains
Golang micro-app, which check whether an IP address belongs to a given network

I use it as an addition to standard **ipcalc** binary distributed
with common Linux distributions.

# Installation

Simply copy it to your ~/bin directory, and you'll all set (assuming your ~/bin directory is in your PATH)

I use it inside of my [Linux desktop deployment automation](https://github.com/docent-net/fedora-desktop-ansible).

# Usage

![alt text](https://github.com/docent-net/ipcalc-contains/blob/main/static/docs-screenshot-usage.png?raw=true)

```bash
$ ./ipcalc-contains 192.168.0.0/24 192.168.0.10
CONTAINS! Network 192.168.0.0/24 contains IP 192.168.0.10

$ ./ipcalc-contains 192.168.0.0/24 192.168.1.10
NOPE! Network 192.168.0.0/24 DOES NOT contain 192.168.1.10
```