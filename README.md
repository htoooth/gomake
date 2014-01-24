# gomake.go
gomake是完成go语言编写程序的工具.配置文件用json写成,简单易用。主要帮助程序员快速编写程序，省去一个一个敲命令的麻烦。

## 使用
gomake install 将所有的依赖包装完
gomake build	构建应用自己的应用
配置文件使用json来写。

Name:应用名
Depends: 依赖包的名字
Packages:自己的包的名字

## 关于 
gomake是由htoo写的。
