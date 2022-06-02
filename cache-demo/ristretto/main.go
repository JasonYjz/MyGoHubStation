package main

import (
	"fmt"
	"github.com/dgraph-io/ristretto"
)

type Person struct {
	Name string
	Age int
	Addr string
}

func main() {
	// type Config struct {
	//    // 可以简单理解成key的数量，他是用来保存key被点击次数的，但实际数量并不是这个设置的值，而是最靠近并且大于或等于该值的2的n次方值减1
	//    // 比如
	//    // 设置成 1，2^0=1,2^1=2,这时候2^0次方等于1，所以最终的值就是2^0-1=0
	//    // 设置成 2，2^0=1,2^1=2,这时候2^1次方等于2，所以最终的值就是2^1-1=1
	//    // 设置成 3，2^0=1,2^1=2,2^2=4,这时候2^2次方大于等于3，所以最终的值就是2^2-1=3
	//    // 设置成 6，2^0=1,2^1=2,2^2=4,2^3=8,这时候2^3次方大于等于7，所以最终的值就是2^3-1=7
	//    // 设置成 20，2^0=1,2^1=2,2^2=4,2^3=8,...,2^4=16,2^5=32,这时候2^5次方大于等于7，所以最终的值就是2^5-1=31
	//    // 官方建议设置成你想要设置key数量的10倍，因为这样会具有更高的准确性
	//    // 根据这个值，可以知道计数器要用的内存 NumCounters / 1024 / 1024 * 4 MB
	//    NumCounters int64
	//    // 单位是可以随意的，例如你想限制内存最大为100MB,你可以把MaxCost设置为100,000,000，那么每个key的成本cost就是bytes
	//    MaxCost int64
	//    // BufferItems决定获取缓冲区的大小。除非您有一个罕见的用例，否则使用'64'作为BufferItems值可以
	//    获得良好的性能。
	//    // BufferItems 决定了Get缓冲区的大小,在有损环形缓冲区ringBuffer中，当数据这个值，就会去对这批key进行点击数量统计
	//    BufferItems int64
	//    // 设置为true，就会统计操作类型的次数，设置为true会消耗成本，建议在测试的时候才开启
	//    Metrics bool
	//    // 当cache被清除的时候调用，过期清除 还有 策略清除
	//    OnEvict func(item *Item)
	//    // 设置一个key失败的时候调用，失败的条件一般有，已经存在key，再次add,或者cost不足
	//    OnReject func(item *Item)
	//    // 删除一个值的时候调用，可以用做手动回收内存。
	//    OnExit func(val interface{})
	//    // 计算key的hash函数
	//    KeyToHash func(key interface{}) (uint64, uint64)
	//    // 计算成本的函数，没有设置成本的时候用这个来计算成本
	//    Cost func(value interface{}) int64
	//    //set(k,v,)
	//    // 设置为true的时候 不计内部结构的成本，默认是计算的，用于存储key-value 结构
	//    // type storeItem struct {
	//    //    key        uint64
	//    //    conflict   uint64
	//    //    value      interface{}
	//    //    expiration int64
	//    // }
	//    IgnoreInternalCost bool
	//}
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}

	p := Person{
		Name: "Jason",
		Age:  36,
		Addr: "Chengdu",
	}

	// set a value with a cost of 1
	cache.Set("key", p, 1)

	// wait for value to pass through buffers
	cache.Wait()

	value, found := cache.Get("key")
	if !found {
		panic("missing value")
	}
	person := value.(Person)
	fmt.Printf("Name:%s, Age:%d, Addr:%s\n", person.Name, person.Age, person.Addr)
	//fmt.Println(value)
	cache.Del("key")
}
