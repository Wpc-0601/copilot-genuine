# Note after the 1.24 version of  k8s, container-d replaced docker-engine as the default CRI

# step 1: close the swap
vim /etc/hostname

sudo swapoff -a
sudo sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab

# step 2: load the module of Kernel
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
overlay
br_netfilter
EOF

sudo modprobe overlay
sudo modprobe br_netfilter

# step 3: change the iptables
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward=1
EOF

sudo sysctl --system

# step 4: install the kubeadm from Alibaba Cloud
sudo apt update
sudo apt install -y apt-transport-https ca-certificates curl
curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | sudo apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main
EOF
sudo apt update

# step 5: install kubeadm kubelet kubectl from above config resource
# you can designate the version such as kubeadm=v1.27, if not, default download the latest version
sudo apt install -y kubeadm kubelet kubectl

# step 6: hold the current version and avoid to automatic upgrade
sudo apt-mark hold kubeadm kubelet kubectl

# step 7: change the configuration of container-d
containerd config default | sudo tee /etc/containerd/config.toml >/dev/null 2>&1
sudo sed -i 's/SystemdCgroup \= false/SystemdCgroup \= true/g' /etc/containerd/config.toml
# TBD

# step 8: change the configuration of container-d
# find the sandbox_image = "registry.k8s.io/pause:3.6" and update it to sandbox_image = "registry.aliyuncs.com/google_containers/pause:3.9"
sudo vim /etc/containerd/config.toml

# step 9: restart container-d
sudo systemctl restart containerd

# step 10: to be initial, change the --apiserver-advertise-address and --pod-network-cidr=10.0.0.0/16 to your Ip address
sudo kubeadm init \
    --pod-network-cidr=10.0.0.0/16 \
    --apiserver-advertise-address=10.0.8.17 \
    --kubernetes-version=v1.23.3 \
    --image-repository registry.cn-hangzhou.aliyuncs.com/google_containers # specify the repository
kubeadm token create --print-join-command

echo 'source <(kubectl completion bash)' >> ~/.bashrc
source ~/.bashrc

wget https://github.com/flannel-io/flannel/releases/latest/download/kube-flannel.yml --no-check-certificate
