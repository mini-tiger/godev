package g

import (
	"godev/tjtools/db/redis"
	"log"
)

var Redis1 redis.RedisStruct

func ReturnRedis() *redis.RedisStruct {
	return &Redis1
}

func InitRedis() {
	// 创建 redis 客户端
	r := Config().Redis
	client1, err := redis.NewRedisClient(r.Dsn, r.Password, 0, r.PoolSize)
	if err != nil {
		log.Fatalf("Redis %+v err:%s\n", r, err)
	}

	Redis1.Conn = client1
	//Redis1.ListOperation()

}

//func (c *RedisS) StringExists(key string) bool {
//	client := c.Conn
//
//	bool1 := client.Exists(key)
//	return bool1.Val()
//}
//
//func (c *RedisS) SetJson(key string, value interface{}, ex time.Duration) {
//	client := c.Conn
//
//	b, err := json.Marshal(value)
//	if err != nil {
//		log.Printf("uuid %s convert json faile data:%+v,err:%s\n", key, value, err)
//	}
//	//fmt.Println(b,e)
//	//fmt.Println("=========================================")
//	log.Printf("存入redis json %s\n", string(b))
//
//	c.Lock()
//	defer c.Unlock()
//
//	err = client.Set(key, b, ex).Err()
//	if err != nil {
//		log.Printf("uuid %s 存入redis json 失败 data:%+v, ERR: %s\n", key, value, err)
//	}
//}

//func (c *RedisS) GetJson(key string) (h interface{}) {
//	client := c.Conn
//
//	c.Lock()
//	defer c.Unlock()
//
//	value, err := client.Get(key).Result()
//	if err != nil {
//		log.Printf("uuid %s output redis faile!\n", key)
//		return
//	}
//	json.Unmarshal([]byte(value), &h)
//	//fmt.Println("==============================")
//	//fmt.Println(value)
//	//fmt.Println(h)
//	//fmt.Println("==============================")
//	return
//
//}

//func (c *RedisS) Expire(key string, time time.Duration) bool {
//	client := c.Conn
//
//	c.Lock()
//	defer c.Unlock()
//
//	b := client.Expire(key, time)
//	return b.Val()
//
//}
//
//func (c *RedisS) SetAdd(key, val string) (bool) {
//
//	client := c.Conn
//
//	c.Lock()
//	defer c.Unlock()
//	i := client.SAdd(key, val)
//	//fmt.Println("=================",i.Val())
//	if i.Val() == 1 { // 1代表成功添加，里面没有此UUID,
//		//log.Printf("uuid:%s, 任务ID:%s 加入redis 任务组合 成功！", val, key)
//		return true
//	} else {
//		//log.Printf("uuid:%s, 任务ID:%s 加入redis 任务组合 失败！", val, key)
//
//		return false
//	}
//
//}
//
//func (c *RedisS) SetInMember(key, val string) (bool) {
//	client := c.Conn
//
//	c.Lock()
//	defer c.Unlock()
//
//	b := client.SIsMember(key, val)
//	return b.Val()
//
//}
//
//func (c *RedisS) SetRemMember(key, appname string, ex time.Duration) bool {
//	client := c.Conn
//
//	c.Lock()
//	defer c.Unlock()
//	//err := client.Set("name", "xys", time.Duration(10)*time.Second).Err()
//
//	err := client.SRem(key, appname).Err()
//	if err != nil {
//		//log.Printf("uuid %s 存入redis app_Set 失败 data:%+v, ERR: %s\n", key, appname, err)
//		return false
//	}
//	return true
//}
