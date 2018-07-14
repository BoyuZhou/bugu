package redis

import (
	"log"
	"bytes"
	"encoding/gob"
	"errors"
	"time"

	"bugu/utils"

	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
)

var MyClient *redis.Client

func RegisterClient() {
	db, _ := beego.AppConfig.Int("REDIS_DB")
	poolSize, _ := beego.AppConfig.Int("REDIS_POOL_SIZE")

	client := redis.NewClient(&redis.Options{
		Addr: beego.AppConfig.String("REDIS_URL"),
		Password: func() string {
			if beego.BConfig.RunMode == "dev" && beego.AppConfig.String("dev::REDIS_PASS") =="" {
				return ""
			}
			return beego.AppConfig.String("REDIS_PASS")
		}(),
		DB:       db,
		PoolSize: poolSize,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Printf("failed to connect redis - %s", err)
		return
	}
	log.Printf("connected to redis - %s", pong)

	MyClient = client
}

func SetCache(key string, value interface{}, timeout int) error  {
	data, err := EnCode(value)
	if err != nil {
		return err
	}
	if MyClient == nil {
		return errors.New("MyClient is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			//LogError(r)
			MyClient = nil
		}
	}()
	timeouts := time.Duration(timeout) * time.Second
	err = MyClient.Set(key, data, timeouts).Err()
	if err != nil {
		 utils.LogError(err)
		 utils.LogError("SetCache失败，key:" + key)
		return err
	} else {
		return nil
	}
}

func GetCache(key string, to interface{}) error {
	if MyClient == nil {
		return errors.New("MyClient is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			utils.LogError(r)
			MyClient = nil
		}
	}()

	data, getErr := MyClient.Get(key).Result()
	if getErr != nil {
		utils.LogError(getErr)
		return errors.New("Cache不存在")
	}

	err := DeCode([]byte(data), to)
	if err != nil {
		utils.LogError(err)
		utils.LogError("GetCache失败，key:" + key)
	}

	return err
}

func DelCache(key string) error {
	if MyClient == nil {
		return errors.New("cc is nil")
	}
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("get cache error caught: %v\n", r)
			MyClient = nil
		}
	}()
	err := MyClient.Del(key)
	if err != nil {
		return errors.New("Cache删除失败")
	} else {
		return nil
	}
}

func EnCode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DeCode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}