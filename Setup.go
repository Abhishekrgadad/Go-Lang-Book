Install Go (Golang) and Set Up Path

For Linux (Ubuntu/Debian-based)
Download and Install Go

*command:
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz


*set Up Environment Variables

echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
source ~/.bashrc

*Verify Installation

go version

**For Windows

Download the Go installer

Run the installer and follow the on-screen instructions.

Set Path Manually (if not done automatically):

Go to Control Panel → System → Advanced system settings.
Click Environment Variables.
Under System Variables, find Path and edit.
Add: C:\Go\bin and %USERPROFILE%\go\bin
Verify Installation
go version

**For macOS (Using Homebrew)
Install Go:
brew install go
Set Path:
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
echo 'export GOPATH=$HOME/go' >> ~/.zshrc
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.zshrc
source ~/.zshrc
Verify Installation

go version

Now you're all set to use Go! 🚀







