## go-wifidog-authserver: Golang Authentication server for WiFiDog.

Go-wifidog-authserver is designed to be a simple wifidog-auth-server to communicate with [Apfree WiFiDog](https://github.com/liudf0716/apfree_wifidog). This project is supported by HTTP WEB Framework of [echo](https://github.com/labstack/echo) which developed by **Golang**

## Quick Start

A binary auth-server program for an example, you can run it in Linxu os for testing:

```
$ git clone https://github.com/KerwinKoo/gowauth.git
$ cd gowauth 
$ cp example/gowauth gowauth
$ ./gowauth

 ⇛ http server started on :8082

```  

After running the command above, the example auth-server will listen the `8082` port as you see.

Modify the WiFiDog's configuration file in router:

```
config wifidog
        option gateway_interface 'br-lan'
        option auth_server_hostname 'your server IP address'
        option auth_server_port '8082'
        option auth_server_path '/wifidog/'
        option check_interval '60'
...

```

Fill your server IP which running this auth-server into `option auth_server_hostname`, and using '8082' port to `option auth_server_port`, then run the command `/etc/init.d/wifidog stop && /etc/init.d/wifidog start` to restart the wifidog daemon in your router.


Back to the platform which running `gowauth`, the auth-server will receive a `ping` request from WiFiDog with  like:

```
time=2016-12-23T11:21:04+08:00, method=GET, uri=/wifidog/ping/?gw_id=......
```

It works when you get the message above.