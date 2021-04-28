# krakend-interceptor
KrakenD HTTP Request Interceptor Plugin

# Install GCC
**Update system package:**

```sh
sudo apt-get update && apt-get upgrade
```

**Install build-essential package:**

```sh
sudo apt install build-essential
```

**Install the manual pages about using GNU/Linux for development:**

```sh
sudo apt-get install manpages-dev
```

**Check GCC version:**

```sh
gcc --version
```

# Install Golang
**Update system package:**

```sh
sudo apt-get update && apt-get upgrade
```

**Get binary source from external link:**

```sh
wget https://dl.google.com/go/go1.13.linux-amd64.tar.gz
```

**Verify the tarball:**

```sh
sha256sum go1.13.linux-amd64.tar.gz
```

**Extract the binary into /usr/local/go directory:**

```sh
sudo tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz
```

**Adjusting the PATH variable:**

```sh
export PATH=$PATH:/usr/local/go/bin
```

**Save to current shell session:**

```sh
source ~/.profile
```

**Check golang version:**

```sh
go version
```

**Install make commnd if does not exist:**

```sh
sudo apt-get -y install make
```

**Run the Makefile:**

```sh
make
```
