#!/bin/bash
# 通用项目初始化脚本
# 用法：bash project-init.sh <project_type>
# 支持类型: go, node, python, doc

set -e

if [ $# -lt 1 ]; then
  echo "用法: bash project-init.sh <project_type>"
  echo "支持类型: go, node, python, doc"
  exit 1
fi

PROJECT_TYPE=$1

# 生成 README.md
if [ ! -f README.md ]; then
  echo "# $(basename $(pwd))" > README.md
fi

# 生成 .gitignore
if [ ! -f .gitignore ]; then
  echo -e ".DS_Store\n*.log\n*.pid\n.env\n.vscode/\nnode_modules/\n__pycache__/\n" > .gitignore
fi

# 初始化 git 并配置用户名邮箱
if [ ! -d .git ]; then
  git init
  git config user.name "francislyj"
  git config user.email "francislyj@gmail.com"
fi

# 按类型初始化
case "$PROJECT_TYPE" in
  go)
    if [ ! -f go.mod ]; then
      go mod init $(basename $(pwd))
      go mod tidy
    fi
    ;;
  node)
    if [ ! -f package.json ]; then
      npm init -y
    fi
    ;;
  python)
    if [ ! -f requirements.txt ]; then
      touch requirements.txt
    fi
    ;;
  doc)
    # 只生成文档和 git，无需额外操作
    ;;
  *)
    echo "不支持的项目类型: $PROJECT_TYPE"
    exit 1
    ;;
esac

echo "项目初始化完成！类型: $PROJECT_TYPE"