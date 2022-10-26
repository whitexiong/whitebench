# whitebench 是什么 ？

> whitebench 构建一个简单的 go 的环境一键启动, 配置简单。无需担心网络问题就可以构建一个 golang + vim 编辑器环境

## 如何使用 ？

> 1. windows 下 需要 wsl2 + docker， 其他系统直接拉取项目后执行 ./builder_image.sh
> 2. 注意需要把 golang 下的 id_rsa 替换成自己的私钥, 这里主要用户 git 拉取项目以及 ssh 服务的配置
>    3. 配置 daemon.json 中的值
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
>    4. 执行 ./run_white_bench.sh

## 环境集成
    
    debian11
    golang
    mysql8.0 // 使用 mysql.test 或者 ip 登录 密码 123456
    nodejs16
    vim
    redis // redis 登录使用自己本机的 ip 地址登录默认是没有密码

## 创建一个 gin 项目

> 进入容器后可拉取一个 golang 项目所以项目在 workspace 目录下

## 其他命令，

> cat /etc/issue 查看系统命令
> docker build . --no-cache // 有缓存的情况下使用


## 配置debian UTF-8 字符集(解决中文乱码的情况)

> dpkg-reconfigure locales
    选择 158 再选择 3 安装 utf8 字符集
  echo "export LC_ALL=en_US.utf8" >> /etc/profile
  source /~/.bashrc


## 配置 vim 的编辑器

  vim ~/.vimrc  // 复制 golang/.vunrc
  :PlugInstall // 安装插件 ctrl+:
  go get golang.org/x/tools/gopls@latest // 安装 go 的语法
  :CocInstall coc-go // 需要 go 环境无需担心已经安装好了

## vim 快捷键 切屏操

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

