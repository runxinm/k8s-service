## 介绍
使用nfs作为集群存储后端，自动生成pvc（和pv）

## 注意事项
- 设置默认存储类
```
kubectl patch sc nfs-client -p '{"metadata": {"annotations": {"storageclass.beta.kubernetes.io/is-default-class": "true"}}}'
```
- sc回收策略
  - archiveOnDelete: "false" 不会做归档而是直接删除文件夹,但使用statefulset时并没有删除
- test_without_pv
  - 可修改不同的访问模式
- test_without_pvc_pc
  - kubectl delete -f statefulset.yaml 后pvc和pv仍然存在
  - 需要手动删除pvc，才能删除pv，因为pv实际上是与pvc绑定的，删除statefulset并未删除pvc，满足有状态服务的需求

## 各yaml文件作用和含义
- 1-ns.yaml
  - 创建命名空间
- 2-patch-*
  - 填写NFS 服务IP地址和共享文件路径
- kustomization.yaml
  - 部署nfs-Subdir External Provisioner
- test_without_pv/*
  - 测试使用nfs动态配置PVC
- test_without_pvc_pv/*
  - 测试使用nfs动态配置pvc和pv

## 参考链接
- [subdir provisioner部署 / 动态配置PV](https://github.com/kubernetes-sigs/nfs-subdir-external-provisioner)
- [test_without_pvc_pv 动态配置pvc和pv](https://blog.51cto.com/mshxuyi/5858663)