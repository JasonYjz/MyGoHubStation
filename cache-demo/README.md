groupcache 分布式缓存和缓存填充库，在许多状况下均可以用来替代内存缓存节点池，不支持 expire。html

go-cache　 内存中键值存储/缓存库（相似于Memcached），适用于单机应用程序。 git

freecache  支持 expire，相似 cache2go。github

ristretto　　未作好面向生产环境golang

golang-lru　固定尺寸大小的 线程安全的 LRU 缓存库，基于 Groupcache，比较简陋。api

cache2go   支持 expire，并发安全的缓存库，api 简单。缓存

gcache　　支持 expire，LFU, LRU and ARC 缓存库，Goroutine 安全。安全

fastcache  不支持 expire，据称比 freecache 更快。并发