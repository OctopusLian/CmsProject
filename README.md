# Elm后台管理平台  

## 技术栈  

（1）后台：Beego框架。  
（2）前端：Vue框架。  
（3）数据库：MySQL数据库+Redis数据库。  

## VSCode Debug  

```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/main.go",
            "args": ["-c","${workspaceFolder}/conf/app.conf"]
        }
    ]
}
```