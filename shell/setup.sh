echo "umask u=rwx,g=,o=" >> ${HOME}/.bashrc
mkdir -p "${HOME}"/.ssh
mkdir -p /temp/cs5424c
cd "${HOME}" || exit
wget https://downloads.yugabyte.com/releases/2.15.2.0/yugabyte-2.15.2.0-b87-linux-x86_64.tar.gz
tar xvfz yugabyte-2.15.2.0-b87-linux-x86_64.tar.gz && cd yugabyte-2.15.2.0/ || exit
./bin/post_install.sh
