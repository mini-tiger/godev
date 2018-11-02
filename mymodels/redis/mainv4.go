package main

import (
	"fmt"
	"gopkg.in/redis.v4"
	"time"
	"sync"
	"log"
	"encoding/json"
)

type redisStruct struct {
	sync.RWMutex
	Conn *redis.Client
}
var Redis1 redisStruct

type inputjson struct {
	S string
	I int64
}

func initclient()  {
	client1:=createClient()
	//defer client.Close()
	Redis1.Conn=client1
	Redis1.stingJson()

	Redis1.stringOperation()
	fmt.Println("===============================")
	Redis1.stringExists("name")
	fmt.Println("===============================")
	Redis1.listOperation()
	fmt.Println("===============================")
	Redis1.Set()
	fmt.Println("===============================")

	Redis1.pubsub1() // 订阅 消息分发


	//setOperation(client)
	//hashOperation(client)
	//
	//connectPool(client)
	select{}
}
func main() {

	initclient()
}

func (c redisStruct)Set(){
	client:=c.Conn
	fmt.Printf("%T\n",client)

	i:=client.SAdd("1111111111111","0edb4518-1fe2-476b-bd3c-f38e1a81b821")
	fmt.Printf("添加了%d个\n",i.Val())
	b:=client.SIsMember("a","abc")
	fmt.Println("KEY中是否存在mem abc" ,b.Val())
	//client.SRem("a","abc")
}

func (c redisStruct)pubsub1()  {
	redisdb:=c.Conn
	var pubsub *redis.PubSub
	pubsub,_ = redisdb.Subscribe("mychannel1")

	// Wait for confirmation that subscription is created before publishing anything.
	_, err := pubsub.Receive() // 等待发布订阅通道完成
	if err != nil {
		panic(err)
	}

	go func() { //生产者
		for{
		//fmt.Println(pubsub.Ping("mychannel1"))

		// Publish a message.
		err = redisdb.Publish("mychannel1", "hello1").Err()
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Duration(2)*time.Second)
		}
	}()


	//time.AfterFunc(time.Second, func() {
	//	// When pubsub is closed channel is closed too.
	//	_ = pubsub.Close()
	//})

	go func() { //消费者
		// Consume messages.
		for {
			//pubsub,_ := redisdb.Subscribe("mychannel1")
			message,_:=pubsub.ReceiveMessage()
			log.Println(message.Channel,message.Payload)
			log.Println(message.String())
			//time.Sleep(time.Duration(10)*time.Second)
		}

	}()
}

// 创建 redis 客户端
func createClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.43.11:6378",
		Password: "",
		DB:       0,
		PoolSize: 5,
	})

	_, err := client.Ping().Result()
	//fmt.Println(pong)
	if err!=nil{
		log.Println("redis conn fail")
	}

	return client
}



func (c redisStruct)stringExists(key string){
	client:=c.Conn

	bool1 := client.Exists(key)
	fmt.Println(bool1)
}

func (c *redisStruct)stingJson() {
	client:=c.Conn

	var ss inputjson
	ss.S="a"
	ss.I=1
	b,_:=json.Marshal(ss)
	//fmt.Println(b,e)
	//fmt.Println(string(b))

	c.Lock()
	defer c.Unlock()

	err := client.Set("j",b , 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("j").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("key: j value type %T, data:%s\n",val, val)

	var sss inputjson
	json.Unmarshal([]byte(val),&sss)
	fmt.Printf("key: j value type %T, data:%+v\n",sss,sss)



}


// String 操作
func (c *redisStruct)stringOperation() {
	c.Lock()
	defer c.Unlock()
	client:=c.Conn
	// 第三个参数是过期时间, 如果是0, 则表示没有过期时间.

	err := client.Set("name", "xys", time.Duration(10)*time.Second).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name", val)


	// 这里设置过期时间.
	err = client.Set("age", "20", 5 * time.Second).Err()
	if err != nil {
		panic(err)
	}

	client.Incr("age") // 自增
	client.Incr("age") // 自增
	client.Decr("age") // 自减

	val, err = client.Get("age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("age", val) // age 的值为21

	// 因为 key "age" 的过期时间是一秒钟, 因此当一秒后, 此 key 会自动被删除了.
	time.Sleep(1 * time.Second)
	val, err = client.Get("age").Result()
	if err != nil {
		// 因为 key "age" 已经过期了, 因此会有一个 redis: nil 的错误.
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println("age", val)
}

// list 操作
func (c *redisStruct)listOperation() {
	c.Lock()
	defer c.Unlock()
	client :=c.Conn

	client.RPush("fruit", "apple") //在名称为 fruit 的list尾添加一个值为value的元素
	client.LPush("fruit", "banana") //在名称为 fruit 的list头添加一个值为value的 元素
	length, err := client.LLen("fruit").Result() //返回名称为 fruit 的list的长度
	if err != nil {
		panic(err)
	}
	fmt.Println("length: ", length) // 长度为2

	value, err := client.LPop("fruit").Result() //返回并删除名称为 fruit 的list中的首元素
	if err != nil {
		panic(err)
	}
	fmt.Println("fruit: ", value)

	value, err = client.RPop("fruit").Result() // 返回并删除名称为 fruit 的list中的尾元素
	if err != nil {
		panic(err)
	}
	fmt.Println("fruit: ", value)
}

// set 操作
func setOperation(client *redis.Client) {
	client.SAdd("blacklist", "Obama") // 向 blacklist 中添加元素
	client.SAdd("blacklist", "Hillary") // 再次添加
	client.SAdd("blacklist", "the Elder") // 添加新元素

	client.SAdd("whitelist", "the Elder") // 向 whitelist 添加元素

	// 判断元素是否在集合中
	isMember, err := client.SIsMember("blacklist", "Bush").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Is Bush in blacklist: ", isMember)


	// 求交集, 即既在黑名单中, 又在白名单中的元素
	names, err := client.SInter("blacklist", "whitelist").Result()
	if err != nil {
		panic(err)
	}
	// 获取到的元素是 "the Elder"
	fmt.Println("Inter result: ", names)


	// 获取指定集合的所有元素
	all, err := client.SMembers("blacklist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("All member: ", all)
}


// hash 操作
func hashOperation(client *redis.Client) {
	client.HSet("user_xys", "name", "xys"); // 向名称为 user_xys 的 hash 中添加元素 name
	client.HSet("user_xys", "age", "18"); // 向名称为 user_xys 的 hash 中添加元素 age

	// 批量地向名称为 user_test 的 hash 中添加元素 name 和 age
	client.HMSet("user_test", map[string]string{"name": "test", "age":"20"})
	// 批量获取名为 user_test 的 hash 中的指定字段的值.
	fields, err := client.HMGet("user_test", "name", "age").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("fields in user_test: ", fields)


	// 获取名为 user_xys 的 hash 中的字段个数
	length, err := client.HLen("user_xys").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("field count in user_xys: ", length) // 字段个数为2

	// 删除名为 user_test 的 age 字段
	client.HDel("user_test", "age")
	age, err := client.HGet("user_test", "age").Result()
	if err != nil {
		fmt.Printf("Get user_test age error: %v\n", err)
	} else {
		fmt.Println("user_test age is: ", age) // 字段个数为2
	}
}

// redis.v4 的连接池管理
func connectPool(client *redis.Client) {
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 100; j++ {
				client.Set(fmt.Sprintf("name%d", j), fmt.Sprintf("xys%d", j), 0).Err()
				client.Get(fmt.Sprintf("name%d", j)).Result()
			}

			fmt.Printf("PoolStats, TotalConns: %d, FreeConns: %d\n", client.PoolStats().TotalConns, client.PoolStats().FreeConns);
		}()
	}

	wg.Wait()
}