- **Install Go:**
```bash
$ sudo rm -rf /usr/local/go
$ go_version=1.23.3
$ cd ~/Downloads
$ sudo apt-get update
$ sudo apt-get install -y build-essential git curl wget
$ wget https://go.dev/dl/go${go_version}.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf go${go_version}.linux-amd64.tar.gz
$ sudo chown -R $(id -u):$(id -g) /usr/local/go
$ rm go${go_version}.linux-amd64.tar.gz
```

- **Add go to your `$PATH` variable:**
```bash
$ mkdir $HOME/go
$ nano ~/.bashrc
$ export GOPATH=$HOME/go
$ export PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
$ source ~/.bashrc
$ go version
```