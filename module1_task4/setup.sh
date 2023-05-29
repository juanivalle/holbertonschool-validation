# Shell script to reproduce a pseudo “production” environment locally:

# Fix the error by updating the script setup.sh to help your colleagues from the
# production team
# The result of the script execution should be successful (with an exit code 0)
# with the directory ./dist containing the generated website

sudo apt-get update
sudo apt-get install -y hugo make
hugo new site dist