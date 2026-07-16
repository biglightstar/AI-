# 2026-07-16 Git 进阶 + CI/CD 入门

## 会话概览

- **日期**：2026-07-16
- **时长**：上午约 1.5h + 下午约 1.5h = 总计约 3h
- **学习形式**：10 问 10 答复习 + 命令行实操 + 个人笔记整理
- **主要话题**：GIT2.5 进阶（rebase/cherry-pick/revert/reflog）+ GIT3.2 Git Flow + GIT3.3 GitHub Flow + GIT5.1 CI/CD

---

## 上午：10 问 10 答复习

学生对之前学的概念存在 10 个模糊点，逐一解答：

1. ✅ 为什么 objects/ 叫 object
2. ✅ 为什么暂存区叫 index
3. ✅ git restore 具体用法
4. ✅ restore vs reset 区别
5. ✅ git commit -a 范围
6. ✅ git diff 四种对比模式
7. ✅ stash 存在哪里、pop/drop 原理
8. ✅ 删除分支原理（只删 refs 文件）
9. ✅ -u 默认路线
10. ✅ origin/main 存在 .git/refs/remotes/origin/

---

## 下午：新内容

### GIT2.5 进阶操作

1. **git rebase** — 变基，把提交移植到最新位置，历史变直线
   - 实操制作分叉 → rebase 合并 → 对比 merge
   - 黄金法则：不对已推送的 commit 做 rebase
2. **git cherry-pick** — 只挑一个 commit 合并
   - 实操从 feature/pay 挑一个提交到 main
3. **git revert** — 安全撤销，新增"反做"提交
   - 实操 revert HEAD → 对比 reset 和 revert
4. **git reflog** — 操作日志，找回"误删"数据
   - 实操查看 HEAD 历史记录

### GIT3.2 Git Flow 实操

- 5 种分支完整流程：develop → feature/login → merge develop → release/1.0 → merge main + develop → tag v1.1.0
- 理解了 Git Flow vs GitHub Flow 的区别

### GIT3.3 分支策略

- GitHub Flow：main + feature 分支 + PR
- 分支命名规范：feat/fix/docs/refactor/test
- PR 本质：请求仓库主人把分支拉进去

### GIT5.1 CI/CD 入门

- 什么是 CI/CD：push 后自动测试 + 自动部署
- GitHub Actions 目录结构：`.github/workflows/*.yml`
- 第一个 workflow 编写和推送
- 排错：workflows 目录必须放在仓库根目录

---

## 掌握话题

| 话题 | 掌握程度 | 关键要点 |
|------|----------|----------|
| GIT2.5 rebase | 高 | 变基/移植提交/黄金法则 |
| GIT2.5 cherry-pick | 高 | 拣选单个 commit/与 merge 对比 |
| GIT2.5 revert | 高 | 安全撤销/history vs reset |
| GIT2.5 reflog/.gitignore | 中高 | 操作日志/reflog 找回数据 |
| GIT5.1 CI/CD | 中 | GitHub Actions 配置/目录要求 |
| GIT3.2 Git Flow | 高 | 5 种分支/完整实操流程 |
| GIT3.3 分支策略 | 高 | GitHub Flow/PR/命名规范 |

---

## 总进度

**Git 完成：16/18 = 89%** 🚀