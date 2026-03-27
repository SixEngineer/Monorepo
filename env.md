#1.配置Linux环境，这里以运行在VMware虚拟机中的Ubuntu2404为例

#2.克隆仓库
  在本地创建一个项目目录，然后使用git指令克隆仓库git clone https://github.com/SixEngineer/Monorepo.git
  
#3.安装cursor
  在cursor官网上下载cursor Linux版本
  https://cursor.com/cn
  
#4.安装go编译器，版本位1.25.8
  wget https://go.dev/dl/go1.25.8.linux-amd64.tar.gz
  sudo rm -rf /usr/local/go
  sudo tar -C /usr/local -xzf go1.25.8.linux-amd64.tar.gz
  添加环境export PATH=$PATH:/usr/local/go/bin
  验证：
  go version

#5.安装openlist实例并允许，挂载本地目录和一个网盘以便于测试

#6。安装go检查器
  ctrl+shift+x打开拓展界面，安装go插件（安装Chinese插件（可选））

#7.SSH到cursor进行开发
  在主机运行cursor，connect via SSH
  安装remote ssh插件
  输入>（运行命令）
  输入ssh，选择Remote-ssh connect to host
  然后输入用户名@远程IP
  等待片刻输入密码登录
  在终端cd到本地仓库目录，然后输入code .
  再次输入密码，可以看到已经进入仓库

