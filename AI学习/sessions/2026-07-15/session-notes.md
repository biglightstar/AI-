# 2026-07-15 Git 分支基础 + 远程协作入门

## 会话概览

- **日期**：2026-07-15
- **时长**：上午约 1h + 下午约 2h + 晚间约 1h = 总计约 4h
- **学习形式**：对话讲解 + 命令行实操 + 个人笔记整理
- **主要话题**：GIT3.1 分支基础 + GIT4.1 远程仓库操作

---

## 学生初始水平

- 已掌握 GIT1 全部（三区模型、配置）+ GIT2 全部（核心命令、log/blame/tag/stash）
- 进度 33%（6/18）
- 分支概念只听过名词，没实操过
- 未接触过 GitHub 远程操作

---

## 讲解内容

### 早间时段：GIT3.1 分支基础

1. **分支本质**：指向 commit 的指针——通过查看 `.git/refs/heads/` 文件验证
2. **创建分支**：`git branch feature-login`——在最新 commit 上贴一张"新贴纸"
3. **切换分支**：`git switch feature-login`——HEAD 指针移到新分支
4. **分叉实验**：在 feature-login 和 main 上分别做提交，形成分叉
5. **合并分支**：`git merge feature-login`——Fast-forward 合并
6. **合并冲突**：两个分支修改同一行 → 手动编辑冲突标记 → `add` → `commit`
7. **删除分支**：`git branch -d feature-login`——只删指针，不删提交
8. **图形查看**：`git log --oneline --graph --all`——看懂分支图

### 午间时段：GIT4.1 远程仓库操作

1. **SSH 密钥**：`ssh-keygen -t rsa -b 4096` 生成 → 配置到 GitHub Settings
2. **SSH 端口问题**：国内网络 22 端口被拒 → 配置 `~/.ssh/config` 用 443 端口
3. **验证连接**：`ssh -T git@github.com` → `Hi biglightstar!` 成功
4. **创建远程仓库**：GitHub 上 New repository → `AI-`
5. **添加远程地址**：`git remote add origin git@github.com:biglightstar/AI-.git`
6. **查看远程**：`git remote -v`（fetch 和 push 两个方向）
7. **推送代码**：`git push -u origin main` → 第一次推送成功
8. **推送其他分支**：`git push -u origin test-remote` → 远程出现新分支
9. **远程跟踪分支**：`origin/main`、`origin/test-remote` 的含义
10. **`-u` 的作用**：设置上游，以后简写 `git push` / `git pull`
11. **GitHub 网页编辑**：在 GitHub 上直接改 README.md
12. **git pull**：从远程拉取更新到本地 `git pull origin main`
13. **git fetch vs git pull**：
    - `fetch`：只下载，不合并
    - `pull`：下载 + 自动合并（pull = fetch + merge）

### 晚间时段：巩固 + 个人笔记

1. **理解 HEAD 指针原理**：通过查看 `.git/HEAD` 和 `.git/refs/heads/main` 文件
   - `git commit` → 分支文件哈希变，HEAD 不变
   - `git switch` → HEAD 文件内容变
2. **自己写笔记**：用户独立编写了 `个人笔记/git笔记.md`
   - 用"正向交流"和"反向交流"分类三区流转
   - 整理了远程仓库地址含义
   - 记录了 stash/restore/reset 等操作

---

## 学生关键问题

1. "分支名是自己命名的吗？" → 是的，团队有命名规范（feature-/fix-/hotfix-）
2. "合并是将分支插入顺序表吗？" → 不是，是指针追赶（Fast-forward）或三方合并
3. "为什么同一行追加会产生冲突？" → Git 按行比较，不是按操作理解
4. "SSH 地址是怎么来的？" → GitHub 自动生成，格式固定
5. "为什么 remote add 后 fetch 和 push 地址一样？" → 同一个仓库的两个操作方向
6. "origin/main 是什么意思？" → 本地存的"远程 main 分支快照"
7. "HEAD 指针怎么变？" → 通过实际操作 `.git/HEAD` 文件和 `.git/refs/heads/` 文件理解

---

## 知识盲区

### 已解决
- ✅ 分支本质：指针文件（通过 cat 命令亲眼验证）
- ✅ 合并冲突的产生与解决（亲自实操）
- ✅ SSH 配置与 GitHub 连接（成功 push）
- ✅ HEAD 指针原理（通过查看文件内容理解）
- ✅ fetch vs pull 区别

### 仍待巩固
- `git reset --soft HEAD~1`（本地仓库→暂存区反向交流）
- `git restore`（工作区丢弃修改）
- 远程仓库地址的深层理解仍有卡壳

---

## 掌握话题

| 话题 | 掌握程度 | 关键要点 |
|------|----------|----------|
| GIT3.1 分支基础 | 高 | 分支是指针、创建/切换/合并/冲突解决/删除、图形查看 |
| GIT4.1 远程操作 | 中高 | SSH 配置、remote add、push/pull/fetch、origin/main 含义 |
| HEAD 指针原理 | 中高 | HEAD 指向分支、分支指向 commit、文件验证 |

---

## 教学反思

- 学生对分支的理解通过"贴纸"比喻和实际文件查看（cat .git/HEAD）有效突破
- 远程操作部分概念的密度偏大（remote add/push/pull/fetch/upstream），需要复习巩固
- 学生自己写笔记的行为值得鼓励——用"正向/反向交流"分类是很好的框架
- 今天约 4 小时学习，信息量大，但学生坚持下来了
