package main

import(
	"encoding/json"
	"fmt"
	//"net/http"
	"strconv"
	//"time"

	"github.com/garyburd/redigo/redis"
)


var (
	  currentPostID int
)

func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	HandleError(err)
	return c
}

func init() {
	CreateData(Post{
		ID: 12345,
		Name : "vaibhav",
		Role :  "abc",
	})
}

func CreateData(p Post) {
	currentPostID = 12345
	p.ID = currentPostID

	//p.Timestamp =time.Now()
	c := RedisConnect()
	defer c.Close()

	b, err := json.Marshal(p)
	HandleError(err)

	reply, err := c.Do("SET", "post:" + strconv.Itoa(p.ID), b)
	HandleError(err)

	fmt.Println("Get",reply)
}


func GetData(id int) Post {
	var person Post

	c := RedisConnect()
	defer c.Close()

	reply, err := c.Do("GET", "post:"+strconv.Itoa(id))
	HandleError(err)

	fmt.Println("GET OK")
	fmt.Println(err)
	if err = json.Unmarshal(reply.([]byte), &person); err != nil {
		panic(err)
	}

	return person
}

