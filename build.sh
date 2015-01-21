# Update and upgrade via apt
apt-get update
apt-get upgrade

# Install primary tools
apt-get -y install gcc build-essential golang git

# Create our application directory
mkdir /opt/sphire
cd /opt/sphire

# Initialize our repository and fetch/pull the application
git init
git remote add origin https://github.com/jsanc623/SphireApp.git
git fetch
git pull origin master

# Set up .bashrc
cd /root
git clone https://gist.github.com/f9bc461bbcd365efe089.git
mv f9bc461bbcd365efe089/.bashrc ./
rm -rf ./f9bc461bbcd365efe089

# Set golang variables in bashrc
echo 'export GOROOT="/opt/sphire/go"' >> ~/.bashrc
echo 'export GOROOT_FINAL="${GOROOT}"' >> ~/.bashrc
echo 'export GOOS="linux"' >> ~/.bashrc
echo 'export GOARCH="amd64"' >> ~/.bashrc
echo 'export GOHOSTOS="linux"' >> ~/.bashrc
echo 'export GOHOSTARCH="amd64"' >> ~/.bashrc
echo 'export GOBIN="${GOROOT}/bin"' >> ~/.bashrc
