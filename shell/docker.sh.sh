# remove the older version
sudo apt-get remove docker \
               docker-engine \
               docker.io
# update the source of apt
sudo apt-get update

# install general tools
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

# download Ali's docker script and automatic installation
curl -fsSL get.docker.com -o get-docker.sh
sudo sh get-docker.sh --mirror Aliyun

# restart the docker service
sudo systemctl enable docker
sudo systemctl start docker

# add docker to group
sudo groupadd docker
sudo usermod -aG docker $USER

# execute the 'hello-world' images to check that installation is correct
sudo docker run --rm hello-world