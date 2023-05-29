# Shell script to reproduce a pseudo “production” environment locally:

sudo apt-get update
sudo apt-get install -y hugo make
make build
exit 255
"recipe for target 'build' failed" 2>&1