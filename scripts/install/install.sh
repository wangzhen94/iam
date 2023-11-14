#!/bin/bash

IAM_ROOT=$(dirname "${BASH_SOURCE[0]}")/../..
source "${IAM_ROOT}/scripts/install/common.sh"

source ${IAM_ROOT}/scripts/install/mariadb.sh
source ${IAM_ROOT}/scripts/install/redis.sh
source ${IAM_ROOT}/scripts/install/mongodb.sh
# source ${IAM_ROOT}/scripts/install/iam-apiserver.sh
# source ${IAM_ROOT}/scripts/install/iam-authz-server.sh
# source ${IAM_ROOT}/scripts/install/iam-pump.sh
# source ${IAM_ROOT}/scripts/install/iam-watcher.sh
# source ${IAM_ROOT}/scripts/install/iamctl.sh
# source ${IAM_ROOT}/scripts/install/man.sh
# source ${IAM_ROOT}/scripts/install/test.sh

function iam::install::install_cfssl()
{
  mkdir -p $HOME/bin/
  wget https://github.com/cloudflare/cfssl/releases/download/v1.6.1/cfssl_1.6.1_linux_amd64 -O $HOME/bin/cfssl
  wget https://github.com/cloudflare/cfssl/releases/download/v1.6.1/cfssljson_1.6.1_linux_amd64 -O $HOME/bin/cfssljson
  wget https://github.com/cloudflare/cfssl/releases/download/v1.6.1/cfssl-certinfo_1.6.1_linux_amd64 -O $HOME/bin/cfssl-certinfo
  #wget https://pkg.cfssl.org/R1.2/cfssl_linux-amd64 -O $HOME/bin/cfssl
  #wget https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64 -O $HOME/bin/cfssljson
  #wget https://pkg.cfssl.org/R1.2/cfssl-certinfo_linux-amd64 -O $HOME/bin/cfssl-certinfo
  chmod +x $HOME/bin/{cfssl,cfssljson,cfssl-certinfo}
  iam::log::info "install cfssl tools successfully"
}

eval $*
