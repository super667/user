package queuex

//
//import (
//	"fmt"
//	"github.com/zeromicro/go-zero/core/stores/redis"
//	"math/rand"
//	"os"
//	"strconv"
//	"time"
//)
//
//const expire = 5
//
//type Queue struct {
//	client *redis.Redis
//	prefix string
//}
//
//func New() *Queue {
//	return &Queue{}
//}
//
//// 待部署，部署中，需要两个队列    lua脚本是不是可以实现
//// 数据中的要使用zset,时间排序，超时剔除
//// 部署完成有别的请求处理，删除
//
//func (q *Queue) Producer(key, value string) {
//
//}
//
//func (q *Queue) Consumer() (value string, err error) {
//	return
//}
//
//const RMQ string = "mqtest"
//
//func producer() {
//	redis_conn, err := redis.Dial("tcp", "127.0.0.1:6379")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	defer redis_conn.Close()
//
//	rand.Seed(time.Now().UnixNano())
//
//	var i = 1
//
//	for {
//		_, err = redis_conn.Do("rpush", RMQ, strconv.Itoa(i))
//		if err != nil {
//			fmt.Println("produce error")
//			continue
//		}
//		fmt.Println("produce element:%d", i)
//		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
//		i++
//	}
//}
//
//func consumer() {
//	redis_conn, err := redis.Dial("tcp", "127.0.0.1:6379")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	defer redis_conn.Close()
//
//	rand.Seed(time.Now().UnixNano())
//
//	for {
//		ele, err := redis.String(redis_conn.Do("lpop", RMQ))
//		if err != nil {
//			fmt.Println("no msg.sleep now")
//			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
//		} else {
//			fmt.Println("cosume element:%s", ele)
//		}
//	}
//}
//
//func main() {
//	list := os.Args
//	fmt.Println(list)
//	if list[1] == "pro" {
//		go producer()
//	} else if list[1] == "con" {
//		go consumer()
//	}
//	for {
//		time.Sleep(time.Duration(10000) * time.Second)
//	}
//}
