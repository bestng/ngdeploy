ngdeploy
========

每一次部署都是痛苦的，设置dns，设置nginx配置

一个go的小程序用来自动生成配置文件

=======

写好自动生成配置文件的设置


nginx在http{。。。}的最后一行添加进include /opt/nginx/vhosts/*;

新建目录，然后把生成的配置放进去就好了


// lib 目录 用于管理工具

