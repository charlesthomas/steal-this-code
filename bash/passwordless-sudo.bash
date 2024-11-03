#!/bin/bash
user=${1:-$(whoami)}
echo "${user} ALL=(ALL) NOPASSWD: ALL" | sudo tee -a /etc/sudoers.d/$user
