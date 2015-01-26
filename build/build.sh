# Update and upgrade via apt
apt-get update
apt-get upgrade

# Install primary tools
apt-get install -y gcc cmake build-essential golang git valgrind
apt-get install -y zlib1g-dev libdb-dev libaio-dev libncurses5-dev
apt-get install -y bison libevent-1.4

# Set up .bashrc
cd /root
git clone https://gist.github.com/f9bc461bbcd365efe089.git
mv f9bc461bbcd365efe089/.bashrc ./
rm -rf ./f9bc461bbcd365efe089
cd /opt

# Set golang variables in bashrc
echo 'export GOROOT="/opt/sphire/go"' >> ~/.bashrc
echo 'export GOROOT_FINAL="${GOROOT}"' >> ~/.bashrc
echo 'export GOOS="linux"' >> ~/.bashrc
echo 'export GOARCH="amd64"' >> ~/.bashrc
echo 'export GOHOSTOS="linux"' >> ~/.bashrc
echo 'export GOHOSTARCH="amd64"' >> ~/.bashrc
echo 'export GOBIN="${GOROOT}/bin"' >> ~/.bashrc
echo 'unset GOROOT'

# Do a bit of cleanup
apt-get autoremove
apt-get autoclean

