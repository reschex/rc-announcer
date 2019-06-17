# puppet manifest
# installs command line tooling
node 'rc-announcer.dev.local' {

  include stdlib
  include archive

  # read settings yaml
  $settings = loadyaml('/vagrant/provision/settings.yaml')
  $packages = $settings[packages]
  $versions = $settings[versions]

  $vagrant_user   = 'vagrant'
  $vagrant_group  = 'vagrant'
  $vagrant_home   = "/home/${vagrant_user}"
  $vagrant_bashrc = "${vagrant_home}/.bashrc"

  # install software packages
  package { $packages:
    ensure => present,
  }

  # install docker
  class { 'docker':
    version      => $versions[docker],
    docker_users => [$vagrant_user],
  }

  # install docker-compose
  class { 'docker::compose':
    version => $versions[compose],
    ensure  => present,
  }

  # install kubectl
  yumrepo { 'kubernetes-yum-repo':
    ensure        => present,
    descr         => 'Kubectl YUM repository',
    baseurl       => 'https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64',
    gpgkey        => 'https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg',
    enabled       => 1,
    gpgcheck      => 1,
    repo_gpgcheck => 1,
  } ~>

  package { 'kubectl':
    ensure => present
  }

  file { '/home/vagrant/.kube':
    ensure => 'directory',
    mode    => '0711',
    owner => $vagrant_user,
    group => $vagrant_group
  }

  # install helm
  $helm_version = $versions[helm]
  $helm_file = "helm-v${helm_version}-linux-amd64.tar.gz"
  $helm_url = "https://storage.googleapis.com/kubernetes-helm/${helm_file}"
  $helm_dir = "/opt/helm-${helm_version}"
  $helm_exe = "${helm_dir}/helm"

  file { $helm_dir:
    ensure  => directory,
    mode    => '0711',
  } ~>

  archive { $helm_file:
    path            => "/tmp/${helm_file}",
    source          => $helm_url,
    extract         => true,
    extract_path    => $helm_dir,
    extract_command => 'tar xfz %s --strip-components=1',
    creates         => $helm_exe,
    cleanup         => true,
  } ~>

  file { '/usr/bin/helm':
    ensure => 'link',
    target => $helm_exe,
  }

  # Copy .bashrc for aliases etc
  file { $vagrant_bashrc:
    ensure => 'file',
    owner => $vagrant_user,
    group => $vagrant_user,
    mode => '0600',
    source => '/vagrant/provision/.bashrc'
  }

  # install Go
  $go_version = $versions[go]
  $go_file = "go${go_version}.linux-amd64.tar.gz"
  $go_url = "https://dl.google.com/go/${go_file}"
  $go_dir = "/opt/go-${go_version}"
  $go_exe = "${go_dir}/bin"

  file { $go_dir:
    ensure  => directory,
    mode    => '0777',
  } ~>

  archive { $go_file:
    path            => "/tmp/${go_file}",
    source          => $go_url,
    extract         => true,
    extract_path    => $go_dir,
    extract_command => 'tar xfz %s --strip-components=1',
    creates         => $go_exe,
    cleanup         => true,
  } ~>

  file { '/usr/bin/go':
    ensure => 'link',
    target => $go_exe,
  }
}
