# Python + AI 全栈开发学习仓库

基于 AI 辅助引导式学习方法，系统学习 Python 全栈与 AI 应用开发——从 Web 后端、数据库到 Prompt Engineering，聚焦工程实践。

## 课程定位

本课程面向希望掌握 **AI Native 应用开发**能力的工程师，核心不是训练大模型或研究深度学习理论，而是掌握：

```text
Python 基础 -> Web 后端 -> 数据库 -> Prompt -> AI Product
```

最终产出是一个可演示、可评测、可写进简历的企业级 AI 项目。

## 适合人群

- 希望系统学习 Python 全栈开发的工程师。
- 想转 AI 工程方向的开发者。
- 想系统学习 Prompt Engineering / LLMOps 的学习者。
- 准备 AI 应用工程相关岗位的求职者。

## 知识领域

| 领域 | 权重 | 核心内容 |
|------|------|----------|
| **A. 前端基础** | 13% | HTML/CSS/JS、React、TypeScript、前端工程化 |
| **B. Python 核心基础** | 13% | 数据类型、OOP、并发、类型提示、模块系统 |
| **C. 数据库** | 15% | SQL、PostgreSQL、Redis、ORM |
| **D. Web 后端开发** | 16% | FastAPI、认证授权、API 设计、WebSocket |
| **E. AI 基础与 Prompt Engineering** | 16% | LLM 原理、Prompt 设计、模型选型、结构化输出 |
| **F. 全栈集成** | 7% | 前后端联调、AI UI 组件、模板引擎 |
| **G. 测试与质量保障** | 5% | pytest、TDD、代码质量工具 |
| **H. DevOps 与部署** | 7% | Docker、CI/CD、Nginx、云服务、监控与日志 |
| **I. 安全与性能** | 4% | Web 安全、限流、缓存 |
| **J. 项目实战** | 4% | 完整 AI 项目搭建、工程化、文档与发布 |

**学习优先级**（按权重排序）：
1. **Web 后端（16%）** + **AI 基础与 Prompt（16%）** — 全栈核心
2. **前端（13%）** + **Python 基础（13%）** + **数据库（13%）**
3. **全栈集成（7%）** + **DevOps（7%）** + **测试（5%）**
4. **安全与性能（4%）** + **项目实战（4%）**

## 学习路径

### 第一阶段：前端基础（第 1周）
- HTML/CSS/JavaScript 核心
- React 基础与进阶
- TypeScript 基础
- 前端工程化

### 第二阶段：Python + Web 后端（第 2-5 周）
- Python 核心基础（类型、OOP、异步、模块）
- HTTP 协议与 FastAPI 框架
- 数据库基础（SQL + PostgreSQL + SQLAlchemy）
- Pydantic 数据验证

### 第三阶段：AI 基础与 Prompt（第 6-7 周）
- LLM API 调用（OpenAI / DeepSeek / 国产模型）
- Prompt Engineering 系统方法
- 流式输出、结构化输出
- 模型选型与成本控制

### 第四阶段：全栈集成与部署（第 8-9 周）
- 前后端联调与 AI UI 组件
- Docker 容器化与 CI/CD
- 云服务部署与监控
- 测试与评测体系

### 第五阶段：项目实战（第 10 周）
- 毕业项目
- 评测、文档、演示

## 毕业项目（三选一）

### 方向 A：AI Chat 应用
- 多轮对话 + 流式输出
- Prompt 模板管理 + 上下文控制
- 对话历史、成本统计、用户管理

### 方向 B：AI 内容生成平台
- 多模型接入（OpenAI / DeepSeek）
- 结构化输出 + 模板系统
- 内容审核、使用量计费

### 方向 C：全栈 AI 工具
- FastAPI 后端 + React 前端
- 用户认证 + API 限流
- Docker 部署 + CI/CD 流水线

## 最终交付物

- GitHub 仓库 + 在线演示
- 项目 README + 系统架构图
- 评测报告
- 日志和指标截图
- 演示视频脚本 + 简历项目描述

## 仓库结构

```
/sessions/                    # 每日学习会话记录
  /YYYY-MM-DD/               # 每天一个文件夹
  SESSION-TEMPLATE.md        # 会话记录模板，不会的知识点也要给我记录下来

/progress/                    # 学习进度追踪（唯一权威来源）
  study-tracker.md           # 综合追踪文件

CLAUDE.md                     # AI 导师指令
README.md                     # 本文件
```

## 使用方法

### 每日学习会话

1. 在本仓库中打开 Claude Code
2. 自然地提问 Python / AI 全栈相关问题——就像和导师对话一样
3. 回答 Claude 提出的理解检查问题
4. 每次会话结束后，Claude 会自动记录学习进度

### 复习与练习

随时告诉 Claude：
- "我们来复习一下我遇到困难的话题"
- "今天我应该关注什么？"
- "给我出一道 FastAPI 依赖注入的练习题"
- "帮我梳理一下 Prompt Engineering 的知识脉络"

### 追踪进度

查看 `/progress/study-tracker.md` 了解：
- 整体学习进度
- 各领域掌握情况
- 知识盲区
- 优先学习计划

## 学习原则

- **不只学 API，要学系统设计。**
- **不只看回答效果，要做评测。**
- **不只做 Demo，要有日志、成本、延迟和安全边界。**
- **不只写功能，要能讲清架构、难点、取舍和指标。**

## 推荐学习资源

**Python & Web：**
- [Python 官方文档](https://docs.python.org/3/)
- [FastAPI 官方文档](https://fastapi.tiangolo.com/)
- [SQLAlchemy 官方文档](https://docs.sqlalchemy.org/)
- [Pydantic 官方文档](https://docs.pydantic.dev/)

**AI 应用开发：**
- [OpenAI API 文档](https://platform.openai.com/docs)
- [Anthropic API 文档](https://docs.anthropic.com/)
- [DeepSeek API 文档](https://platform.deepseek.com/)
- [Vercel AI SDK](https://sdk.vercel.ai/)

**工程化：**
- [Real Python](https://realpython.com/)
- [TestDriven.io](https://testdriven.io/)
- [FastAPI 最佳实践](https://github.com/zhanymkanov/fastapi-best-practices)

## 环境变量约定

所有项目统一使用 `.env.example` 描述配置项，不提交真实 `.env`。

```bash
LLM_API_KEY=
LLM_BASE_URL=
LLM_MODEL=
DATABASE_URL=
REDIS_URL=
```

## 快速开始

1. **克隆本仓库**：
   ```bash
   git clone <repo-url>
   cd fullstarck-study
   ```

2. **运行 Claude Code**：
   ```bash
   claude-code
   ```

3. **开始学习！** 提出你的第一个问题，Claude 将引导你学习并自动追踪进度。

`CLAUDE.md` 文件包含了 Claude 如何辅导你的所有指令。

---

## 关于本项目

本项目将 Python 全栈开发与 AI 应用工程相结合，帮助你系统掌握从 Web 后端到 AI 产品的完整能力栈。无论你是编程初学者还是希望转型 AI 工程方向的开发者，这套方法都能帮助你高效学习。

主分支修改的内容

功能分支做的修改

这是我在git hub 上做的修改

我要修改一下，测试git fetch 和 git pull 的区别

main新提交
