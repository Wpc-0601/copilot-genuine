# output base64 encoding of the origin string
echo -n "generic" | base64

# Note This tutorial(installation procedure) only works for version 1.23 of kubernetes and below,
# container-d is recommended for version higher than 1.24, for details, see ubuntu-k8s-new.sh

# step 1: update the hostname and make kubernetes recognized the machine
sudo vi /etc/hostname

# step 2: change the c-group-driver to systemd and append others information among the daemon.json
cat <<EOF | sudo tee /etc/docker/daemon.json
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2"
}
EOF

sudo systemctl enable docker
sudo systemctl daemon-reload
sudo systemctl restart docker


# step 3: change the default iptables configuration
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF

cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward=1
EOF

sudo sysctl --system

# step 4: close the swap
sudo swapoff -a
sudo sed -ri '/\sswap\s/s/^#?/#/' /etc/fstab


# step 5: install the kubeadm from Alibaba Cloud
sudo apt install -y apt-transport-https ca-certificates curl
curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | sudo apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main
EOF
sudo apt update

# step 6: install kubeadm kubelet kubectl from above config source
sudo apt install -y kubeadm kubelet kubectl


# step 7: hold the current version and avoid to automatic upgrade
sudo apt-mark hold kubeadm kubelet kubectl


# step 8: fetch the designated images from Alibaba images repository and change the tag
# latest version specify the arguments with --images-repository while execute kubeadm-init ,so you can skip the step
# see detail step 9
repo=registry.cn-hangzhou.aliyuncs.com/google_containers
for name in $(kubeadm config images list --kubernetes-version v1.27.1); do

    src_name=${name#registry.k8s.io/}
    src_name=${src_name#coredns/}

    docker pull $repo/$src_name

    docker tag $repo/$src_name $name
    docker rmi $repo/$src_name
done

# step 9: to be initial,  change the --apiserver-advertise-address and --pod-network-cidr=10.0.0.0/16 to your Ip address
sudo kubeadm init \
    --pod-network-cidr=10.0.0.0/16 \
    --apiserver-advertise-address=10.0.8.17 \
    --kubernetes-version=v1.23.3 \
    --image-repository registry.cn-hangzhou.aliyuncs.com/google_containers # specific the repository
