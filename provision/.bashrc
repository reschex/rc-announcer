# .bashrc

# Source global definitions
if [ -f /etc/bashrc ]; then
        . /etc/bashrc
fi

# Uncomment the following line if you don't like systemctl's auto-paging feature:
# export SYSTEMD_PAGER=

# User specific aliases and functions

# Start the shell in the /vagrant directory
cd /vagrant

# Source config file if it exists
source config

# Export the IP address of eth0 so we can use it in things like docker-compose
export HOST_IP=$(ip addr show eth0 | grep -Po "(?<=inet )\d+\.\d+\.\d+\.\d+")

# Add Go bin to PATH
export PATH=$PATH:/opt/go-1.12.4/bin

# Aliases
alias dc="docker-compose"
alias dcj="docker-compose down -v && docker-compose build && docker-compose up -d && docker-compose logs -f"
alias k=kubectl

# ksn = Kubernetes Set Namespace
# Set the current kubernetes namespace to the given namespace
#
# Parameters
# $1 = Kubernetes namespace to set as default/current
ksn() {
  kubectl config set-context $(kubectl config current-context) --namespace=$1
}
