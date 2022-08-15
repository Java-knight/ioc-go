# IOC-Go 


## 思路
基本上是有 sdMap<sdID, sd>【basic autowire】，它相当于是一个父级Map，所有的 sd 注册都要写入。
如果normal autowire，它维护的是一个 normalEntryDescriptorMap（但是注册它前也得把这个 sd 注册到 sdMap中）。

