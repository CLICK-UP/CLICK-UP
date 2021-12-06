# CLICK-UP

#### Click-UP: Toward the Software Upgrade of Click-Based Modular Network Function,
in Proceedings of *ACM SIGCOMM Conference (demo)*, also published in *IEEE Systems Journal*.

Please find the manuscript from [here](https://ieeexplore.ieee.org/document/9043707).

## How to build

1. Golang configuration

        root@localhost:~$ cd ~
        root@localhost:~$ mkdir go
        root@localhost:~$ cd go
        root@localhost:~/go$ mkdir pkg
        root@localhost:~/go$ mkdir src 
        root@localhost:~/go$ cd src

2. Downloading sourcecode of click-up

        root@localhost:~/go/src$ git clone https://github.com/CLICK-UP/CLICK-UP.git
    
3. Installing dependency for golang and click

        root@localhost:~/go/src$ sudo apt-get update
        root@localhost:~/go/src$ sudo apt-get install golang
        root@localhost:~/go/src$ sudo apt-get install g++
        root@localhost:~/go/src$ sudo apt-get install gcc
        root@localhost:~/go/src$ sudo apt-get install make
        root@localhost:~/go/src$ sudo apt-get install autoconf
    
4. Compiling click

        root@localhost:~/go/src$ git clone https://github.com/kohler/click.git
        root@localhost:~/go/src$ cd click
        root@localhost:~/go/src/click$ sudo ./configure --enable-userlevel --disable-linuxmodule
        root@localhost:~/go/src/click$ sudo make
        root@localhost:~/go/src/click$ sudo make install
        

5. GOPATH configuration

        root@localhost:~/go/src$ mv CLICK-UP/* ./
        root@localhost:~/go/src$ rm -rf CLICK-UP
        root@localhost:~/go/src$ export GOPATH=/home/$HOSTNAME/go/	
        /*please also add it in /etc/profile for permanent force*/
    
6. Compiling and running server of click-up

        root@localhost:~/go/src$ go build httpserver.go
        root@localhost:~/go/src$ ./httpserver

7. Enjoy!
