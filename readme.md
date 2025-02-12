# 计算器项目

这是一个简单的计算器项目，支持基本的算术运算，包括加法、减法、乘法、除法、乘方和开根号。

## 功能

- 加法 (`+`)
- 减法 (`-`)
- 乘法 (`*`)
- 除法 (`/`)
- 乘方 (`^`)
- 开根号 (`sqrt()`)

## 使用方法

1. 克隆或下载此项目到本地。
2. 使用终端进入项目目录。
3. 运行以下命令编译并运行程序：

    ```sh
    go run main.go
    ```

4. 在提示符下输入表达式并按回车键。例如：

    ```
    输入表达式: 3 + 5 * (2 - 8)
    结果: -25.000000
    ```

5. 输入 `exit` 退出程序。

## 示例

```sh
输入表达式: 2 + 3 * 4
结果: 14.000000

输入表达式: 2 ^ 3
结果: 8.000000

输入表达式: sqrt(16)
结果: 4.000000

输入表达式: exit
```

## 依赖

- Go 语言环境

确保已安装 Go 语言环境，可以通过以下命令检查：

```sh
go version
```

## 许可证

此项目使用 MIT 许可证。有关详细信息，请参阅 LICENSE 文件。
