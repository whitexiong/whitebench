# whitebench 是什么 ？

## 快速搭建一套脚手架 php/go

> whitebench 构建一个简单的 go 的环境一键启动, 配置简单。无需担心网络问题就可以构建一个 golang + vim 编辑器环境。大大的降低学习成本并且无缝迁移非常好用，但是需要额外的学习一些服务端的知识。

## 如何使用 ？

- 1. windows 下 需要 wsl2 + docker，然后 clone 本地址; 其他系统直接拉取项目后执行 ./builder_image.sh

- 2. [注意]需要把 golang/ 下的 id_rsa 替换成自己的私钥, 这里主要用户 git 拉取项目以及 ssh 服务的配置

- 3. 配置 daemon.json 中的值

```json

  {
      "builder": {
        "gc": {
          "defaultKeepStorage": "20GB",
          "enabled": true
        }
      },
      "dns": [  // dns 配置好有可能镜像下载不了
        "223.5.5.5",
        "223.6.6.6"
      ],
      "experimental": false,
      "features": {
        "buildkit": true
      },
      "registry-mirrors": [
          // 这里写自己的地址,登录阿里云
      ]
  }

```

- 4. 执行 ./run_white_bench.sh

## 环境集成
    
- debian11
- golang1.75.5
- mysql8.0 // 使用 mysql.test 或者 ip 登录 密码 123456
- nodejs16
- vim
- redis // redis 登录使用自己本机的 ip 地址登录默认是没有密码

## 创建一个 gin 项目

> 进入容器后可拉取一个 golang 项目所以项目在 workspace 目录下拉取自己的项目

## 其他命令，

> cat /etc/issue 查看系统命令
> docker build . --no-cache // 有缓存的情况下使用清除缓存


## 配置debian UTF-8 字符集(解决中文乱码的情况)

> dpkg-reconfigure locales
    选择 158 再选择 3 安装 utf8 字符集
  echo "export LC_ALL=en_US.utf8" >> /etc/profile // 容器中已经加了
  source /~/.bashrc


## 配置 vim 的编辑器

  vim ~/.vimrc  // 复制 golang/.vunrc
  :PlugInstall // 安装插件 ctrl+:
  go get golang.org/x/tools/gopls@latest // 安装 go 的包在项目下使用
  :CocInstall coc-go // 需要 go 环境无需担心已经安装好了

## vim 编辑器的快捷键使用

> 由于是服务端开发所以需要熟练使用 vi/vim 编辑器,

    ctrl-w > //向右加宽，默认值为1  
    ctrl-w N > //向右加宽宽度N  
    ctrl-w < // 同理 

> 横屏/竖屏分屏打开当前文件
    
    ctrl+w s  
    ctrl+w v 

> 切换分屏

    ctrl+w h,j,k,l  
    ctrl+w 上下左右键  
    crtl+w进行分屏窗口的切换 按完以后再按一个w  
    crtl+w进行分屏窗口的切换 按完以后再按一个r 互换窗口  
    crtl+w进行分屏窗口的切换 按完以后再按一个c 关闭窗口

> 关闭分屏

    ctrl+W c 关闭当前窗口  
    ctrl+w q 关闭当前窗口，若只有一个分屏且退出vim  
    :only 仅保留当前分屏  
    :hide 关闭当前分屏

> 打开目录树

    :NERDTree 打开目录树
    :NERDTreeClose 关闭目录树

> 此外还能打开命令行

  :ter

## 扩展

  如果想扩展自己的容器可以替换 golang 目录下的容器, 具体的 dockerfile 文件可在历史提交记录中查看。

> 编写自己的 docker-compose.yml 加入扩展一键启动

## 下载其他版本的 go

  wget https://go.dev/dl/go1.18.8.linux-arm64.tar.gz