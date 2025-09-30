#!/usr/bin/env bash

# Source: https://rpi4cluster.com/k3s/k3s-docker-install/

# Remove the installation you have now.
sudo apt remove -y docker docker-engine docker.io containerd runc

# Install prerequisites
sudo apt install -y ca-certificates curl gnupg lsb-release

# Install GPG key
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/debian/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

# Install Repository
echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/debian $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Install Docker itself
sudo apt update
sudo apt install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin


# Edit Docker settings
# Add /etc/docker/daemon.json configuration for docker daemon.

sudo mkdir /etc/docker/
echo '{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "insecure-registries": ["registry.cube.local:5000"],
  "experimental": true,
  "log-driver": "json-file",
  "storage-driver": "overlay2",
  "log-opts": {
    "max-size": "100m"
  }
}' | sudo tee -a /etc/docker/daemon.json


# Enable at boot and start docker daemon
sudo systemctl enable docker
sudo systemctl start docker

# Adding support for multi-arch
# Some additional packages are needed if you're going to use OpenFaaS and build on arm64.

#might already be installed
sudo apt install -y binfmt-support qemu-user-static


# Add user to docker group
sudo usermod -aG docker $USER
newgrp docker