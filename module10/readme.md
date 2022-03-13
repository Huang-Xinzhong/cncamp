### 作业要求

1. 为 HTTPServer 添加 0-2 秒的随机延时
2. 为 HTTPServer 项目添加延时 Metric
3. 将 HTTPServer 部署至测试集群，并完成 Prometheus 配置
4. 从 Promethus 界面中查询延时指标数据
5. （可选）创建一个 Grafana Dashboard 展现延时分配情况





在代码中添加随机延迟，添加metrices

代码地址

配置httpserver

```yaml
#添加Prometheus自动发现
	  annotations:
        prometheus.io/scrape: "true"
        prometheus.io/port: "8080"
#声明容器端口
	  ports:
        - containerPort: 8080
```



prometheus查看延迟指标

https://github.com/Huang-Xinzhong/cncamp/blob/main/module10/prometheus.png



grafana dashboard 效果

https://github.com/Huang-Xinzhong/cncamp/blob/main/module10/grafana.png

