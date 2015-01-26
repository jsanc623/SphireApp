# Create our application directory
mkdir /opt/sphire
cd /opt/sphire

# Initialize our repository and fetch/pull the application
git init
git remote add origin https://github.com/jsanc623/SphireApp.git
git fetch
git pull origin master
