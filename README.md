# 实现自动签到
`config.json`是配置文件，自行修改



使用`golang`编译，假设生成的可执行文件为`autocheck`

RUN:
```bash
    nohup ./autocheck &>> log &
```

追加到标准输出、错误输出到`log`。