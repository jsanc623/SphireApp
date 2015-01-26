# Build MariaDB with TokuDB engine
apt-get install software-properties-common
apt-key adv --recv-keys --keyserver hkp://keyserver.ubuntu.com:80 0xcbcb082a1bb943db
add-apt-repository 'deb http://nyc2.mirrors.digitalocean.com/mariadb/repo/10.0/ubuntu trusty main'
apt-get update
apt-get install mariadb-server
touch /etc/mysql/conf.d/tokudb.cnf
echo 'plugin-load=ha_tokudb' >> /etc/mysql/conf.d/tokudb.cnf
