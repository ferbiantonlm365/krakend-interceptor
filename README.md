# krakend-interceptor
KrakenD HTTP Request Interceptor Plugin

# Install GCC
**Update system package:**
sudo apt-get update && apt-get upgrade

**Install build-essential package:**
sudo apt install build-essential

**Install the manual pages about using GNU/Linux for development:**
sudo apt-get install manpages-dev

**Check GCC version:**
gcc --version

# Install Golang
**Update system package:**
sudo apt-get update && apt-get upgrade

**Get binary source from external link:**
wget https://dl.google.com/go/go1.13.linux-amd64.tar.gz

**Verify the tarball:**
sha256sum go1.13.linux-amd64.tar.gz

**Extract the binary into /usr/local/go directory:**
sudo tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz

**Adjusting the PATH variable:**
export PATH=$PATH:/usr/local/go/bin

**Save to current shell session:**
source ~/.profile

**Check golang version:**
go version

**Install make commnd if does not exist:**
sudo apt-get -y install make

**Run the Makefile:**
make
