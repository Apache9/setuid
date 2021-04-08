# # 通过setuid切换用户
  
  ## 说明
  使用方法
  ```console
  ./su-starter <target username> <command> ...
  ```
  程序会通过调用setuid切换到对应的用户，再通过execve把自己切换成需要执行的命令。程序在执行时需要拥有root权限，因此需要将文件的owner设置为root，同时权限设置为4755(增加setuid位)。
  
  ## 测试方法
  ```console
  go build
  sudo chown root:root su-starter
  sudo chmod 4755 su-starter
  ./su-starter <target username> test.sh
  ```
  test.sh会每隔5秒输出一次当前用户名称，可以检查是否真的切换成功