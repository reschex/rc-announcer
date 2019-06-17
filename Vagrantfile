#!/usr/bin/env ruby
# vagrant development environment

# read settings file
require "yaml"
cd = File.dirname(__FILE__)
settings = YAML.load_file("#{cd}/provision/settings.yaml")

# check for requried plugins
settings["plugins"].each do |plugin|
  unless Vagrant.has_plugin?(plugin)
    raise "Missing plugin! Run: vagrant plugin install #{plugin}"
  end
end unless settings["plugins"].nil? || settings["plugins"].empty?

# check for required env vars
['HTTP_PROXY', 'HTTPS_PROXY', 'NO_PROXY'].each do |var|
  unless ENV[var] || ENV[var] == ""
    raise %{
      Missing #{var} environment variable!
      Export it into your bashrc: echo export #{var}=... >> ~/.bashrc
      Then source the bashrc: source ~/.bashrc
    }
  end
end

Vagrant.configure(2) do |config|

  # virtual box settings
  config.vm.box = settings["box"]
  config.vm.hostname = settings["hostname"]
  config.vm.provider :virtualbox do |vb|
    vb.name = settings["name"]
    vb.memory = settings["memory"]
    vb.cpus = settings["cpus"]
  end

  # configure proxy
  config.proxy.http = ENV["HTTP_PROXY"]
  config.proxy.https = ENV["HTTPS_PROXY"]
  config.proxy.no_proxy = ENV["NO_PROXY"]

  # forward ports
  settings["ports"].each do |port|
    config.vm.network :forwarded_port, host: port, guest: port
  end unless settings["ports"].nil? || settings["ports"].empty?

  # install puppet
  config.vm.provision :shell, inline: <<-SHELL
    if ! type -p puppet >/dev/null 2>&1; then
      rpm -Uvh https://yum.puppetlabs.com/puppet5/puppet5-release-el-7.noarch.rpm
      yum install -y puppet-agent
    fi
    FUNCTIONS=/etc/puppetlabs/code/environments/production/lib/puppet/functions
    if [ ! -d $FUNCTIONS ]; then
      mkdir -p $FUNCTIONS
      cp /vagrant/provision/env.rb $FUNCTIONS
    fi
  SHELL

  # install puppet modules
  settings["modules"].each do |mod|
    config.vm.provision :shell, inline: "puppet module install #{mod}"
  end unless settings["modules"].nil? || settings["modules"].empty?

  # apply puppet configuration
  config.vm.provision :shell, inline: "puppet apply /vagrant/provision/provision.pp"

  # Always sync git folder
  config.vm.synced_folder ".", "/vagrant", type: "virtualbox",  mount_options: ["dmode=775,fmode=775"]

end
