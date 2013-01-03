ngdeploy
========

每一次部署都是痛苦的，设置dns，设置nginx配置

一个go的小程序用来自动生成配置文件

=======

写好自动生成配置文件的设置


nginx在http{。。。}的最后一行添加进include /opt/nginx/vhosts/*;

新建目录，然后把生成的配置放进去就好了


======
说明，git代码后，请在script目录里吧dev.sh里面的GOPATH的变量改为你的目录，然后运行

