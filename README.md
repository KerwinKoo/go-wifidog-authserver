## go-wifidog-authserver: Golang Authentication Server for WiFiDog.

Go-wifidog-authserver is designed to be a simple wifidog authentication-server to communicate with [Apfree WiFiDog](https://github.com/liudf0716/apfree_wifidog). This project is supported by HTTP Web Framework of [ECHO](https://github.com/labstack/echo) developed with **Golang** .

## Quick Start

There is a bin auth-server program in path of `example/gowauth` which you can use to test your WiFiDog. 

```
$ git clone https://github.com/KerwinKoo/gowauth.git
$ cd gowauth 
$ cp example/gowauth gowauth
$ ./gowauth

 ⇛ http server started on :8082

```  

After running the command above, the example program will listen to the `8082` port as you can see.

Modify the WiFiDog's configuration file in your router:

```
config wifidog
        option gateway_interface 'br-lan'
        option auth_server_hostname 'your server IP address'
        option auth_server_port '8082'
        option auth_server_path '/wifidog/'
        option check_interval '60'
...

```

Fill your auth-server's IP into `option auth_server_hostname`, fill '8082' port into `option auth_server_port`, and then run the command `/etc/init.d/wifidog stop && /etc/init.d/wifidog start` to restart the wifidog daemon in your router.


Back to the platform which `gowauth` is running on, the auth-server will receive a `ping` request from WiFiDog like:

```
time=2016-12-23T11:21:04+08:00, method=GET, uri=/wifidog/ping/?gw_id=......
```

It works when you get the message above.
