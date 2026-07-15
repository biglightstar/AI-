# Git 学习进度追踪

**学习启动日期**：2026-07-13
**目标完成时间**：2026-08-24（6 周）
**学习方向**：Git 版本控制系统深入学习
**学习状态**：进行中
**学习节奏**：系统推进，每 1-2 天一个话题，命令行实操驱动
**参考教程**：[菜鸟教程 Git 教程](https://www.runoob.com/git/git-tutorial.html)

本文档追踪所有 Git 学习进度，包括：
- 各领域话题掌握情况
- 知识盲区识别
- 优先学习计划

---

## 快速统计

- **总体进度**：8/18 话题掌握 = **44%**
- **上次学习**：2026-07-15 — GIT3.1 分支基础 + GIT4.1 SSH 远程配置 + git push
- **学习方向**：Git 基础概念 → 本地操作 → 分支管理 → 远程协作 → 自动化 CI/CD

---

## 领域进度汇总

| 领域 | 权重 | 话题数 | 掌握 | 状态 | 优先级 |
|------|------|--------|------|------|--------|
| **GIT1. Git 基础概念与环境** | 20% | 4 | 4 | 🟢 已完成 | — |
| **GIT2. Git 本地核心操作** | 30% | 5 | 4 | 🟡 进行中 | 高 |
| **GIT3. Git 分支管理** | 22% | 3 | 1 | 🟡 进行中 | 高 |
| **GIT4. Git 远程协作** | 18% | 3 | 1 | 🟡 已起步 | 中 |
| **GIT5. DevOps 自动化** | 10% | 3 | 0 | 未开始 | 低 |

---

## GIT1. Git 基础概念与环境（20%，4 话题）

> 🎯 **目标**：理解版本控制概念，完成 Git 环境搭建，理解 Git 核心三区模型

> 📖 参考：[Git 教程](https://www.runoob.com/git/git-tutorial.html) | [安装配置](https://www.runoob.com/git/git-install-setup.html) | [工作流程](https://www.runoob.com/git/git-workflow.html) | [工作区暂存区版本库](https://www.runoob.com/git/git-workspace-index-repo.html)

### 已掌握（4/4）✅

- [x] **GIT1.1** Git 简介与版本控制概念（什么是 Git/集中式 vs 分布式版本控制/Git 诞生背景/Git vs SVN 区别/适用场景）— 2026-07-13
- [x] **GIT1.2** Git 安装与配置（git config 三级配置 system/global/local、.gitconfig 文件位置与格式、git config --list --show-origin 查看配置来源）— 2026-07-14
- [x] **GIT1.3** Git 工作流程（工作区 Working Directory/暂存区 Staging Area/仓库 Repository 三区模型、git status 状态解读、完整流转：edit→add→commit）— 2026-07-14
- [x] **GIT1.4** Git 工作区、暂存区和版本库深入（git add 本质—内容复制到暂存区+开始追踪/文件四种状态 Untracked/Modified/Staged/Committed/三区流转动手实验）— 2026-07-14

---

## GIT2. Git 本地核心操作（30%，5 话题）

> 🎯 **目标**：掌握日常工作最频繁使用的 Git 命令——创建仓库、提交、查看历史、标签、进阶操作

> 📖 参考：[创建仓库](https://www.runoob.com/git/git-create-repository.html) | [基本操作](https://www.runoob.com/git/git-basic-operations.html) | [提交历史](https://www.runoob.com/git/git-commit-history.html) | [标签](https://www.runoob.com/git/git-tag.html) | [进阶操作](https://www.runoob.com/git/git-advance.html)

### 已掌握（5/5）

- [x] **GIT2.1** git init 与 git clone（git init 初始化本地仓库已实操/仓库目录结构/.git 目录/仓库根目录概念/git clone 概念理解待远程实操）— 2026-07-14
- [x] **GIT2.2** Git 基本操作（git add/status/diff/commit/reset/restore 深入掌握/文件状态流转 Untracked→Staged→Committed/— git rm/git mv 待学）— 2026-07-14
- [x] **GIT2.3** git log 与 git blame（git log 核心参数—oneline/-p/--stat/--author/--since/对比两次 commit/git blame 逐行追责）— 2026-07-14
- [x] **GIT2.4** git tag 标签管理（轻量标签 git tag/附注标签 git tag -a -m/查看标签 git tag -l/标签用于版本标记 v1.0.0）— 2026-07-14
- [x] **GIT2.5** Git 进阶操作（git stash 暂存工作现场 stash/pop/list/show/git stash 场景：临时保存半成品修改/git rebase/cherry-pick 待学）— 2026-07-14

### 未学习（0/5）

---

## GIT3. Git 分支管理（22%，3 话题）

> 🎯 **目标**：掌握 Git 最强大的特性——分支，理解合并、冲突解决与团队分支策略

> 📖 参考：[分支管理](https://www.runoob.com/git/git-branch.html) | [Git Flow](https://www.runoob.com/git/git-flow.html)

### 未学习（0/3）

- [ ] **GIT3.1** Git 分支基础（什么是分支——指向 commit 的指针/HEAD 指针概念/创建分支 git branch <name>/切换分支 git checkout git switch/查看分支 git branch -a -r/删除分支 git branch -d -D/合并分支 git merge/Fast-forward 快进合并 vs 三方合并/合并冲突的产生与解决——<<<<<<< ======= >>>>>>> 标记/VS Code 可视化冲突解决/git mergetool）
- [ ] **GIT3.2** Git Flow 工作流（Git Flow 安装 Linux/macOS/Windows/Git Flow 分支模型 5 种分支—master/main/develop/feature/release/hotfix/分支操作原理/初始化 git flow init/创建功能分支 git flow feature start/完成功能分支 git flow feature finish/创建发布分支 git flow release/创建修复分支 git flow hotfix/Git Flow 优缺点——适合固定发布周期的团队，不适合持续部署）
- [ ] **GIT3.3** 分支策略与实战练习（GitHub Flow——main + 短命 feature 分支 + PR/Trunk-Based Development——主干开发频繁提交/分支命名规范 feat/fix/docs/chore/refactor/test/综合实战：创建 feature 分支→开发提交→推送到远程→创建 PR→Code Review→合并到 main→解决冲突→删除分支 完整流程）

---

## GIT4. Git 远程协作（18%，3 话题）

> 🎯 **目标**：掌握 Git 远程仓库操作，能在 GitHub/Gitee 上进行团队协作

> 📖 参考：[远程仓库](https://www.runoob.com/git/git-remote-repo.html) | [服务器搭建](https://www.runoob.com/git/git-server.html) | [SourceTree](https://www.runoob.com/git/source-tree-intro.html)

### 未学习（0/3）

- [ ] **GIT4.1** Git 远程仓库操作（添加远程仓库 git remote add/查看远程 git remote -v/从远程拉取 git fetch 只拉不合并 vs git pull 拉取并合并/git pull --rebase 避免多余 merge commit/推送到远程 git push/git push -u 设置上游追踪/删除远程仓库 git remote remove/HTTP vs SSH 协议对比/CODING/GitHub 免密 Push/Pull 配置）
- [ ] **GIT4.2** Git 服务器搭建（建立裸存储库 bare repository——git init --bare/服务器端安装 Git/创建证书登录 SSH Key/克隆远程仓库 git clone user@server:/path/理解自建 Git 服务器本质：一个可通过 SSH 访问的裸仓库/与 GitHub/Gitee 的关系）
- [ ] **GIT4.3** SourceTree GUI 工具（SourceTree 安装与配置/连接 GitHub 账户/创建本地仓库/图形化理解 Git 操作/SourceTree vs 命令行——何时用 GUI 何时用命令行/不依赖 GUI，以命令行理解为核心）

---

## GIT5. DevOps 自动化（10%，3 话题）

> 🎯 **目标**：掌握 GitHub Actions 实现 CI/CD 自动化流水线

> 📖 参考：[GitHub Actions](https://www.runoob.com/git/github-actions.html)

### 未学习（0/3）

- [ ] **GIT5.1** GitHub Actions 核心概念（什么是 CI/CD/Workflow 工作流/Job 任务/Step 步骤/Action 动作/Event 触发事件/Runner 运行环境/目录结构 .github/workflows/*.yml/第一个 Workflow 编写—Hello World）
- [ ] **GIT5.2** GitHub Actions 核心语法（触发事件 on—push/pull_request/schedule 定时/workflow_dispatch 手动/Jobs 与 Steps 详解/runs-on 运行环境/needs Job 间依赖/if 条件执行/环境变量 env/github 内置上下文变量/Secrets 敏感信息管理/矩阵构建 Matrix—多环境多版本同时测试/常用 Actions—checkout/setup-node/setup-python/cache/upload-artifact）
- [ ] **GIT5.3** GitHub Actions 实战（Node.js 项目 CI—自动测试/Python 项目 CI—pytest 自动运行/构建并推送 Docker 镜像/构建后自动部署到服务器/定时任务 自动备份数据库/调试技巧—Debug 日志/本地测试 Workflow 用 act/常见问题与注意事项）

---

## 已掌握话题

| 话题 | 掌握日期 | 掌握程度 | 关键要点 |
|------|----------|----------|----------|
| GIT1.1 Git 简介与版本控制概念 | 2026-07-13 | 中高 | 版本控制痛点/存档点模型/分支概念/git init→add→commit 三步流程/命令英文原意 |
| GIT1.2 Git 安装与配置 | 2026-07-14 | 中 | 三级配置、.gitconfig 文件位置 |
| GIT1.3 Git 工作流程 | 2026-07-14 | 高 | 三区流转、git status/git diff 理解透彻 |
| GIT1.4 三区模型深入 | 2026-07-14 | 高 | add 双重作用、暂存区与工作区关系 |
| GIT2.1 init/clone | 2026-07-14 | 中 | 仓库根目录概念、.git 位置、init 已实操 clone 待学 |
| GIT2.2 基本操作 | 2026-07-14 | 高 | add/status/diff/commit/reset/restore 核心命令 |
| GIT2.3 log/blame | 2026-07-14 | 中高 | log 核心参数、blame 逐行追责、路径写法 |
| GIT2.4 tag | 2026-07-14 | 高 | 标签创建/查看/v1.0.0 里程碑标记 |
| GIT2.5 stash | 2026-07-14 | 中高 | 暂存工作现场 stash/pop/show/list |

---

## 知识盲区

### 高优先级

| 盲区 | 严重程度 | 说明 |
|------|----------|------|
| GIT3 分支管理全部 | 高 | 分支创建/合并/冲突解决——Git 最核心特性 |
| GIT4 远程协作 | 高 | push/pull/PR 流程——团队协作必备技能 |

### 中优先级

| 盲区 | 严重程度 | 说明 |
|------|----------|------|
| GIT4 远程协作 | 中 | push/pull/PR 流程是团队协作必备技能 |

### 低优先级

| 盲区 | 严重程度 | 说明 |
|------|----------|------|
| GIT5 CI/CD | 低 | GitHub Actions 可在实际项目中边做边学 |

### 最近解决

| 话题 | 解决日期 | 说明 |
|------|----------|------|
| GIT1.1 版本控制概念 | 2026-07-13 | 理解了版本控制本质/完成 git init→add→commit 首个存档/掌握三步口诀 |
| GIT1.2-GIT1.4 三区模型 | 2026-07-14 | 掌握 git config 三级配置/git add 双重作用/三区流转模型/16 题复习正确率 81% |
| GIT2.1-GIT2.5 本地操作 | 2026-07-14 | 掌握 init/clone 概念/log 历史查看/blame 逐行追责/tag 标签管理/stash 暂存工作现场/路径理解 |

---

## 学习计划

### 当前学习状态

**学生背景**：
- 已掌握：HTML 基础
- Python：之前学过但现在已经忘完了，需要重新捡起来
- Go：尚未开始，计划系统学习
- Git：当前电脑上的 Git 是别人帮忙安装的，自己对 Git 完全不懂，从零开始

### 六周学习路径（2026-07-13 → 2026-08-24）

**第一阶段：Git 基础概念（第 1 周，7/13 - 7/19）**
1. GIT1.1 Git 简介与版本控制概念
2. GIT1.2 Git 安装与配置
3. GIT1.3 Git 工作流程
4. GIT1.4 工作区、暂存区和版本库深入

**第二阶段：本地核心操作（第 1 周，7/13 - 7/14）— ✅ 已提前完成**
5. GIT2.1 git init 与 git clone
6. GIT2.2 Git 基本操作——add/commit/status/diff/reset
7. GIT2.3 git log 与 git blame——历史查看与追溯
8. GIT2.4 git tag 标签管理
9. GIT2.5 Git 进阶操作——stash（已完成）/rebase/cherry-pick 待学

**第三阶段：分支管理（第 2 周，7/15 - 7/19）**
10. GIT3.1 Git 分支基础——创建/合并/冲突解决
11. GIT3.2 Git Flow 工作流
12. GIT3.3 分支策略与实战练习

**第四阶段：远程协作（第 3 周，7/20 - 7/26）**
13. GIT4.1 Git 远程仓库操作
14. GIT4.2 Git 服务器搭建
15. GIT4.3 SourceTree GUI 工具

**第五阶段：CI/CD 自动化（第 4-5 周）**
16. GIT5.1-GIT5.3 GitHub Actions

### 下次会话建议

**确定方向**：**GIT3.1 Git 分支基础** — 分支创建/切换/合并/冲突解决，Git 最核心特性（或切换 Go 语言 G1.3）

**教学要求**：
- **命令行实操驱动**：不用 GUI 工具，每个命令在终端手敲
- 一个小知识点一个小知识点的讲
- 每个概念讲完马上动手验证
- 准备一个 Git 练习仓库，实时操作

---

## 学习原则

- **不只学命令，要理解数据模型。** Git 本质是一个内容可寻址的文件系统——理解 blob/tree/commit 对象模型才能驾驭上层命令。
- **不只背参数，要理解工作流。** 每个命令都有它存在的场景——在正确的时间用正确的命令。
- **不只自己练，要模拟协作。** Git 是团队工具——用两个本地目录互相 push/pull，或建 GitHub 仓库实战。
- **不要怕犯错，要学会"后悔药"。** reflog、reset、checkout、stash 是你最好的朋友——Git 几乎不会真正丢失数据。

---

## 仓库结构

```
/Git/                          # Git 练习目录（一个本地仓库即可覆盖多数场景）
/sessions/                     # 学习会话记录（跨学科共用）
  /YYYY-MM-DD/
    session-notes.md
/progress/
  Git-study.md                 # ← 本文件：Git 学习唯一追踪文件
  go-study-tracker.md          # Go 学习追踪文件
  study-tracker.md             # Python 全栈学习追踪文件
```
