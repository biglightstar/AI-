# Git 学习笔记

> 学习路径：菜鸟教程 Git 教程（16 个链接对应 16 个话题）
> 学习方式：命令行实操驱动，理解英文原意

---

## 一、Git 基础概念

### 1.1 版本控制概念

- **Git 是什么**：分布式版本控制系统，记录文件的每次修改，可随时回到任意版本
- **核心思想**：存档点（commit）模型，取代"最终版1.docx、最终版2.docx"的混乱
- **分支概念**：回到旧版本修改不会覆盖新版本，形成独立分支

### 1.2 安装与配置

**git config** = **conf**ig**uration**（配置）

三级配置（从通用到具体）：

| 级别 | 命令 | 配置文件位置 | 影响范围 |
|------|------|-------------|---------|
| 系统级 | `git config --system` | Git 安装目录 `gitconfig` | 整台电脑 |
| 用户级 | `git config --global` | `C:\Users\用户名\.gitconfig` | 当前用户所有仓库 |
| 仓库级 | `git config --local` | 仓库 `.git/config` | 仅当前仓库 |

- **优先级**：仓库级 > 用户级 > 系统级（具体的覆盖通用的）
- **查看配置来源**：`git config --list --show-origin`
- **.gitconfig 格式**：纯文本，`[section]` 分组，`key = value`

### 1.3 工作流程（三区模型）

```
工作区 ─── git add ───> 暂存区 ─── git commit ───> 仓库
```

| 区域 | git status 提示 | 说明 |
|------|---------------|------|
| **工作区** | `Changes not staged for commit:` | 你正在编辑的地方 |
| **暂存区** | `Changes to be committed:` | 准备提交的队列 |
| **仓库** | （提交后不显示） | 已存档的历史 |

### 1.4 文件四种状态

| 状态 | 含义 | 所在区 |
|------|------|--------|
| **Untracked** | Git 不认识的新文件 | 工作区 |
| **Modified** | 已跟踪但修改了，未暂存 | 工作区 |
| **Staged** | 已 add，等待提交 | 暂存区 |
| **Committed** | 已提交存档 | 仓库 |

**git add 的双重作用**：
1. 对新文件 → **开始追踪**（Git 从"不认识"变成"认识"）
2. 对已修改文件 → **把修改放进暂存区**

---

## 二、Git 本地核心操作

### 2.1 git init 与 git clone

- **git init** = **init**ialize（初始化）— 创建新仓库，生成 `.git` 目录
- **git clone** = clone（克隆、复制）— 把远程仓库完整复制到本地（含全部历史）
- **仓库根目录** = 放着 `.git` 文件夹的那个目录
- 路径写法：根目录下的文件直接写名字，文件夹里的文件写 `文件夹名/文件名`

### 2.2 核心命令

| 命令 | 英文原意 | 功能 |
|------|---------|------|
| `git add <文件>` | add（添加） | 放入暂存区 + 开始追踪 |
| `git status` | status（状态） | 查看仓库当前状态 |
| `git diff` | diff（difference，差异） | 对比工作区和暂存区 |
| `git commit -m "消息"` | commit（提交） + message（消息） | 把暂存区内容提交存档 |
| `git commit -a -m "消息"` | commit + all | 快捷提交（只对已跟踪文件） |
| `git reset <文件>` | reset（重置） | 从暂存区撤回（add 的反操作） |
| `git restore <文件>` | restore（恢复） | 丢弃工作区修改（危险，内容丢失） |
| `git restore --staged <文件>` | restore + staged | 从暂存区撤回（内容不丢） |

### 2.3 git log — 查看提交历史

| 参数 | 作用 |
|------|------|
| `--oneline` | 一行显示一个提交 |
| `-2` / `-3` | 限制显示条数 |
| `-p` | 显示具体改动（patch） |
| `--stat` | 显示改了多少文件、多少行 |
| `--author="名字"` | 按作者过滤 |
| `--since="1 day ago"` | 按时间过滤 |
| `--graph --all` | 图形化显示分支 |
| `--decorate` | 显示标签/分支指向 |

### 2.4 git blame — 逐行追责

- **blame** = 责备、归咎
- 输出格式：`commit哈希 (作者 日期+时区 行号) 文件内容`
- 用于追溯"这行代码谁写的？什么时候改的？"

### 2.5 git diff — 四种对比模式

| 命令 | 对比范围 |
|------|---------|
| `git diff` | 工作区 ↔ 暂存区 |
| `git diff --staged` | 暂存区 ↔ 最新 commit |
| `git diff HEAD` | 工作区 ↔ 最新 commit |
| `git diff <commitA> <commitB>` | 两次 commit 之间 |

- `-` 开头的行 = 删除/旧版本
- `+` 开头的行 = 添加/新版本

### 2.6 git tag — 标签管理

- **tag** = 标签、标记
- 给某次 commit 起一个有意义的别名（如 `v1.0.0`）

**常用命令**：

```bash
git tag v1.0.0              # 创建轻量标签
git tag -a v1.1.0 -m "消息"  # 创建附注标签（推荐）
git tag -l                  # 列出所有标签（-l = list）
git log --oneline --decorate # 查看标签指向
```

### 2.7 git stash — 暂存工作现场

- **stash** = 藏起来、存放
- 场景：改到一半需要切分支，把半成品"藏起来"

```bash
git stash            # 藏起来（工作区变干净）
git stash list       # 查看 stash 列表
git stash show       # 看 stash 改了哪些文件
git stash show -p    # 看 stash 具体改了什么（-p = patch）
git stash pop        # 取回来（pop = 弹出）
```

---

## 三、常用命令速查表

| 场景 | 命令 | 英文 |
|------|------|------|
| 新建仓库 | `git init` | initialize |
| 查看状态 | `git status` | status |
| 添加文件 | `git add 文件名` | add |
| 提交 | `git commit -m "消息"` | commit + message |
| 查看历史 | `git log --oneline` | log |
| 查看差异 | `git diff` | difference |
| 从暂存区撤回 | `git reset 文件名` | reset |
| 丢弃修改 | `git restore 文件名` | restore |
| 取消暂存 | `git restore --staged 文件名` | restore + staged |
| 打标签 | `git tag 标签名` | tag |
| 暂存现场 | `git stash` | stash |
| 取回暂存 | `git stash pop` | stash pop |

---

## 核心口诀

> **改 → add → commit**（三步循环）
> **工作区 → 暂存区 → 仓库**（三区流转）
