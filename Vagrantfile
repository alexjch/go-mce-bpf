# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "fedora/38-cloud-base"
  config.vm.box_version = "38.20230413.1"
  config.vm.provider "virtualbox" do |vb|
    vb.memory = 2048
    vb.cpus = 2
  end
  if Vagrant.has_plugin?("vagrant-proxyconf")
    if ENV["http_proxy"]
      config.proxy.http = ENV["http_proxy"]
    end

    if ENV["https_proxy"]
      config.proxy.https = ENV["https_proxy"]
    end

    if ENV["no_proxy"]
      config.proxy.no_proxy = ENV["no_proxy"]
    end
  end
  config.vm.provision "shell", inline: <<-SHELL
    sudo dnf update -y
    sudo dnf install -y git vim bpftool gcc clang libbpf-devel llvm bcc-tools
    # mcelog
    git clone https://git.kernel.org/pub/scm/utils/cpu/mce/mcelog.git
    pushd ./mcelog
    make
    sudo make install
    sudo modprobe mce-inject
    popd
    sudo rm -rf mcelog
    # mce-inject
    git clone https://github.com/andikleen/mce-inject.git
    pushd ./mce-inject
    make
    sudo make install
    popd
    rm -rf ./mce-inject
  SHELL
end

