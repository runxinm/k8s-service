nodeip="" 
podcidr=""
svcidr="" # 10.98.0.0/12
node=""

usage() {                                 # Function: Print a help message.
    echo "Usage: $0 [ -n node name  such as k8sm1 ] " 1>&2
    echo "Usage: $0 [ -i node_ip  such as 10.55.195.151 ]" 1>&2
    echo "Usage: $0 [ -c Pod_cidr such as 10.198.0.0/16 ] " 1>&2
    echo "Usage: $0 [ -s service_cidr such as 10.98.0.0/12 ] " 1>&2
}

exit_abnormal() {                         # Function: Exit with error.
  usage
  exit 1
}

while getopts ":n:i:c:s:" options; do
    case "${options}" in              
    n)
      node=${OPTARG}
      ;;
    i)                                
      nodeip=${OPTARG}
      ;;    
    c)                                
      podcidr=${OPTARG}               
      ;;
    s)
      svcidr=${OPTARG}
      ;;
    :)                                    # If expected argument omitted:
      echo "Error: -${OPTARG} requires an argument."
      exit_abnormal                       # Exit abnormally.
      ;;
    *)                                    # If unknown (any other) option:
      exit_abnormal                       # Exit abnormally.
      ;;
  esac
done


mv /etc/apt/sources.list /etc/apt/sources.list.bak
echo "
deb https://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse
deb-src https://mirrors.aliyun.com/ubuntu/ focal main restricted universe multiverse
deb https://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse
deb-src https://mirrors.aliyun.com/ubuntu/ focal-security main restricted universe multiverse
deb https://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse
deb-src https://mirrors.aliyun.com/ubuntu/ focal-updates main restricted universe multiverse
# deb https://mirrors.aliyun.com/ubuntu/ focal-proposed main restricted universe multiverse
# deb-src https://mirrors.aliyun.com/ubuntu/ focal-proposed main restricted universe multiverse
deb https://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse
deb-src https://mirrors.aliyun.com/ubuntu/ focal-backports main restricted universe multiverse
"  >> /etc/apt/sources.list
echo "apt-get update"
apt-get update
echo "vim /etc/ssh/sshd_config"
echo "PermitRootLogin yes" >> /etc/ssh/sshd_config
echo "service ssh restart:------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ "
service ssh restart

echo "swapoff -a  &&  free -h ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------"
swapoff -a
free -h

echo "ufw disable ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- "
ufw disable

echo "time tb ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- "
sudo timedatectl set-timezone Asia/Shanghai
sudo systemctl restart rsyslog

echo "not suspend: ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------"
sudo systemctl mask sleep.target suspend.target hibernate.target hybrid-sleep.target

echo "br_netfilter load?----------------------------------------------------------------------------------------------------------------------------------------------------------------------------"
modprobe br_netfilter
lsmod | grep br_netfilter

echo "set ip6tables iptables = 1---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- "
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF


echo "sysctl --system ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------"
sysctl --system

cat /etc/sysctl.d/10-network-security.conf
echo "rm /etc/sysctl.d/10-network-security.conf -----------------------------------------------------------------------------------------------------------------------------------------------------------"
rm /etc/sysctl.d/10-network-security.conf

echo "
net.ipv4.conf.default.rp_filter=1
net.ipv4.conf.all.rp_filter=1
" >> /etc/sysctl.d/10-network-security.conf

sysctl --system



## docker install


apt-get update
echo "apt-get -y install apt-transport-https ca-certificates curl software-properties-common Start ******************** "
apt-get -y install apt-transport-https ca-certificates curl software-properties-common
echo "apt-get -y install apt-transport-https ca-certificates curl software-properties-common  Finish *************************"

echo "install GPG "
curl -fsSL https://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | sudo apt-key add -

echo "sudo add-apt-repository deb [arch=amd64] https://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable "
sudo add-apt-repository "deb [arch=amd64] https://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable"


echo "apt-get update "
apt-get -y update

echo "docker install "
apt-get install -y docker-ce
apt-get install -y docker-compose


# vim /etc/docker/daemon.json
mkdir /etc/docker
echo "vim /etc/docker/daemon.json "
rm /etc/docker/daemon.json
# echo "
# {
#     \"registry-mirrors\": [
#         \"https://registry.docker-cn.com\",\"http://hub-mirror.c.163.com\"
#     ],
#     \"live-restore\":true,
#     \"exec-opts\": [\"native.cgroupdriver=systemd\"]
# }
# " >> /etc/docker/daemon.json

cat <<EOF | tee /etc/docker/daemon.json
{
    "registry-mirrors": [
        "https://registry.docker-cn.com","http://hub-mirror.c.163.com"
    ],
    "live-restore":true,
    "exec-opts": ["native.cgroupdriver=systemd"]
}
EOF

echo "systemctl restart docker "
systemctl restart docker

echo "systemctl enable docker "
systemctl enable docker

echo "docker version "
docker version

echo "docker info "
docker info


# k8s install
echo "start install k8s ********************************************************************************************"

sudo apt-get update && sudo apt-get install -y ca-certificates curl software-properties-common apt-transport-https curl

curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add -

cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main
EOF

apt-get update
apt-get install -y kubelet=1.22.0-00 kubeadm=1.22.0-00 kubectl=1.22.0-00
kubeadm config images list --kubernetes-version=v1.22.0

imagelist=(
    kube-apiserver:v1.22.0
    kube-controller-manager:v1.22.0
    kube-scheduler:v1.22.0
    kube-proxy:v1.22.0
    pause:3.5
    etcd:3.5.0-0
    coredns:v1.8.4
)

for imageName in ${imagelist[@]};do
    echo $imageName
    docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName
    docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName k8s.gcr.io/$imageName
    docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName
done

# 注意coredns的变了,在以前的版本中按照上述循环即可,新版本不行
docker tag  k8s.gcr.io/coredns:v1.8.4 k8s.gcr.io/coredns/coredns:v1.8.4

# check
docker images


# k8s cluster

# nodeip = localhost
# podcidr = 10.198.0.0/16
# servicecidr = 10.98.0.0/12

# work node 到此为止 直接kubeadm join即可

kubeadm init --apiserver-advertise-address=$nodeip --pod-network-cidr=$podcidr --service-cidr=$svcidr --kubernetes-version=1.22.0


mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

# git clone http://10.21.5.108:3000/runxinma/k8s_install.git
# cd k8s_install

# calico CNI插件
# wget https://raw.githubusercontent.com/projectcalico/calico/v3.24.1/manifests/tigera-operator.yaml
# kubectl apply -f https://raw.githubusercontent.com/projectcalico/calico/v3.23/manifests/tigera-operator.yaml
# kubectl create -f https://projectcalico.docs.tigera.io/archive/v3.23/manifests/tigera-operator.yaml

kubectl create -f ./service/tigera-operator.yaml

# wget https://raw.githubusercontent.com/projectcalico/calico/v3.24.0/manifests/custom-resources.yaml
# wget http://10.21.5.108:3000/runxinma/k8s_install/src/branch/master/custom-resources.yaml
# echo $podcidr
# kubeadm init --apiserver-advertise-address=10.55.195.139 --pod-network-cidr=$podcidr --service-cidr=10.98.0.0/12 --kubernetes-version=1.22.0


cat <<EOF | tee ./service/custom-resources.yaml
# This section includes base Calico installation configuration.
# For more information, see: https://projectcalico.docs.tigera.io/master/reference/installation/api#operator.tigera.io/v1.Installation
apiVersion: operator.tigera.io/v1
kind: Installation
metadata:
  name: default
spec:
  # Configures Calico networking.
  calicoNetwork:
    # Note: The ipPools section cannot be modified post-install.
    ipPools:
    - blockSize: 26
      cidr: $podcidr
      encapsulation: VXLANCrossSubnet
      natOutgoing: Enabled
      nodeSelector: all()

---

# This section configures the Calico API server.
# For more information, see: https://projectcalico.docs.tigera.io/master/reference/installation/api#operator.tigera.io/v1.APIServer
apiVersion: operator.tigera.io/v1
kind: APIServer 
metadata: 
  name: default 
spec: {}
EOF

kubectl create -f ./service/custom-resources.yaml

# watch kubectl get pods -n calico-system
echo "sleep 600s. make sure pods in namespaces calico-system init "
date
sleep 600

# master 可调度
kubectl taint node $node node-role.kubernetes.io/master:NoSchedule-
# kubectl taint node $node node-role.kubernetes.io/master:NoSchedule-

# Local StorageClass
# wget https://openebs.github.io/charts/openebs-operator.yaml
kubectl create -f ./service/openebs-operator.yaml
echo "sleep 240s"
date
sleep 240

# kubectl apply -f local_sv_test.yaml
# watch kubectl get pods --all-namespaces

# set default sc
kubectl patch sc openebs-hostpath -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'

# install KS
kubectl version
#  要求可用 CPU > 1 核；内存 > 2 G。CPU 必须为 x86_64
free -g
kubectl get sc
# kubectl apply -f kubesphere-installer.yaml
date
# echo "sleep 150s"
# sleep 150
# kubectl apply -f cluster-configuration.yaml
# 检查安装日志
# kubectl logs -n kubesphere-system $(kubectl get pod -n kubesphere-system -l 'app in (ks-install, ks-installer)' -o jsonpath='{.items[0].metadata.name}') -f

#  3.3版本的KS 的 角色权限 有问题,使用下面yaml文件重新创建一遍
# kubectl apply -f https://raw.githubusercontent.com/kubesphere/ks-installer/release-3.3/roles/ks-core/prepare/files/ks-init/role-templates.yaml

# kubectl edit cc ks-installer -n kubesphere-system
# 修改ks configure


# kubeadm token work安装好环境后只需要运行下面命令的输出结果即可加入到集群中
# 查看token  以便其他节点加入集群
# kubeadm token create --print-join-command
