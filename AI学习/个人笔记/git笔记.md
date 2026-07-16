### git分区
- **工作区**：就是在电脑上的文件，就是.git之外的文件
- **暂存区**：.git/index这里就是暂存区，如果工作区做出修改之后，先放到暂存区
- **本地仓库**：是目的文件夹同级目录下一个.git的文件，如果确认不修改了，就把暂存区的文件放在本地仓库，
- **远程仓库（GitHub）**：把本地仓库的内容在远程仓库留一个快照，其他人可以从这里下载我的项目到他的电脑

### 本地仓库git的创建
**1.cd /所在盘符（小写）/绝对路径**
先打开git bash，用cd，change directory，切换到目的文件（你想要给这个文件创建一个仓库）目录
**2.git init**
切换到目的目录之后，执行该条指令就会初始化仓库，即在目的文件的同级目录下创建一个以.git为结尾的文件夹（有时候这个文件夹是隐藏的，需要设置一下才能看见）
**3.之后就可以进行各种操作了**

### 远程仓库origin的创建
**1.一般远程仓库命名为origin**
**2.git remote add origin git@github.com:biglightstar/AI-.git**
remote,远的
add，增加
origin,给远程仓库起一个origin的名字<远程仓库在GitHub上的地址>
**3.git@github.com:biglightstar/AI-.git**
以git这个登录名去访问GitHub的服务器，访问GitHub用户名为biglightstar的名为/AI-.git的仓库
git:用git这个登录名去访问
@：去链接到
github.com：git的服务器地址
biglightstar：GitHub用户名
AI-.git：要访问的该用户的这个仓库
**4.注意**
写完3之后，对远程仓库相关操作的时候，就可以直接用仓库名，不用再写它的地址了

### git分区之间的“正向交流”
**1.工作区---add---暂存区**
git add 文件名
add,增加
把本地工作区的指定文件放到暂存区，则该文件的状态由untracked变为changes to be committed
另一种情况，该文件已经追踪了，但是我又修改了，这时候状态由changes to be committed变为changes not staged to be committed，add会把状态在转变为changs to be committed
**2.暂存区---commit---本地仓库**
git commit -m"备注信息"（备注信息，这里自己写内容）
commit,提交 承诺
-m message,消息 信息
把暂存区也就是状态为changes to be committed的文件全部提交到本地仓库，状态变为nothing to commit ,working tree clean
**3.本地仓库---push---远程仓库**
git push origin main
push,推
把本地仓库的内容推送到远程仓库


### git分区之间的“反向交流”
**1.暂存区---reset---工作区**
git reset 文件名
reset，重设
把刚放进暂存区的东西又拿出来，文件状态由changes to be committed退回untracked
**2.远程仓库---pull---本地仓库**
git pull 远程仓库名 远程仓库分支名
pull:把远程仓库对应分支的东西直接更新到工作区了
pull=fetch + merge
fetch 抓取，把远程分支的东西下载到本地仓库git（先看一下这些修改是否是自己想要的）
merge 合并，更新工作区（如果合适的话就更新到本地）
