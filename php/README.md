# gRPC - PHP

## Prerequisites

* `php`: version 5.5 or higher, 7.0 or higher
* `pecl`: version 1.9 or higher
* `composer`

### Install PHP and PECL on Debian/Ubuntu

PHP 5.5 or above

```bash
$ [sudo] apt-get install php5 php5-dev php-pear zlib1g-dev
```

PHP 7.0 or above

```bash
$ [sudo] apt-get install php7.0 php7.0-dev php-pear zlib1g-dev
```

### Install PHP and PECL on CentOS/RHEL 7

```bash
$ [sudo] rpm -Uvh https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
$ [sudo] rpm -Uvh https://mirror.webtatic.com/yum/el7/webtatic-release.rpm
$ [sudo] yum install php56w php56w-devel php-pear phpunit gcc zlib-devel
```

### Install PECL on Mac OS
```bash
$ brew install autoconf

$ curl -O http://pear.php.net/go-pear.phar
$ [sudo] php -d detect_unicode=0 go-pear.phar
```

If you are using Mac OS El Capitan (10.11) or above and have not installed
`pecl` before, you might have to temporarily [disable](http://blog.g-design.net/post/137712472685/configuring-apache-and-php-after-updating-to-os-x)
System Integrity Protection, or "SIP" before proceeding.

To disable SIP take the following steps :

 * Reboot into recovery mode by holding down Command+R on reboot
 * Access the Terminal from the dropdown menu
 * Run the command: `csrutil disable`
 * Reboot

To re-enable SIP again boot into recovery mode and run the command
`csrutil enable`.

### Install Composer
```bash
$ curl -sS https://getcomposer.org/installer | php
$ [sudo] mv composer.phar /usr/local/bin/composer
```

### Install gRPC PHP Extension

Install gRPC extension:

```bash
$ [sudo] pecl install grpc
```

After installing the gRPC extension, make sure you add this line to your
`php.ini` file (e.g. `/etc/php5/cli/php.ini`, `/etc/php5/apache2/php.ini`,
or `/usr/local/etc/php/5.6/php.ini`), depending on where your PHP installation
is.

```bash
extension=grpc.so
```

Note: for users on CentOS/RHEL 6, unfortunately this step won't work. Please
follow the instructions [here](https://github.com/grpc/grpc/tree/master/src/php#build-from-source)
to compile the PECL extension from source.

### Install Protobuf

You will need to install the protocol buffer compiler `protoc` and the special
plugin for generating server and client code from `.proto` service definitions.
For the first part of our quickstart example, we've already generated the server
and client stubs from
[helloworld.proto](https://github.com/grpc/grpc/tree/{{site.data.config.grpc_release_branch}}/examples/protos/helloworld.proto),
but you'll need the tools for the rest of our quickstart, as well as later
tutorials and your own projects.

To install `protoc`:

The simplest way to do this is to download pre-compiled binaries for your platform(`protoc-<version>-<platform>.zip`) from here: https://github.com/google/protobuf/releases

  * Unzip this file.
  * Update the environment variable `PATH` to include the path to the protoc binary file.

To compile the gRPC PHP Protoc Plugin:

```bash
$ git clone https://github.com/grpc/grpc
$ cd grpc && git submodule update --init
$ make grpc_php_plugin
```

## Server
PHP can't do gRPC server-y stuff :-( So run the go server instead.
```bash
$ go run server/main.go
```

## Client
Install dependencies:
```bash
$ cd client && composer install
```

Run the client:
```bash
$ php index.php
```
