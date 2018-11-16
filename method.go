package rejson

import "github.com/go-redis/redis"

//ExtendClient ...
func ExtendClient(client *redis.Client) *Client {
	return &Client{
		client,
		&redisProcessor{
			Process: client.Process,
		},
	}
}

//ExtendPipeline ...
func ExtendPipeline(pipeline redis.Pipeliner) *Pipeline {
	return &Pipeline{
		pipeline,
		&redisProcessor{
			Process: pipeline.Process,
		},
	}
}

func concatWithCmd(cmdName string, args []interface{}) []interface{} {
	result := make([]interface{}, 1)
	result[0] = cmdName
	for _, v := range args {
		if str, ok := v.(string); ok {
			if len(str) == 0 {
				continue
			}
		}
		result = append(result, v)
	}
	return result
}

func jsonSetExecute(c *redisProcessor, args ...interface{}) *redis.StatusCmd {
	cmd := redis.NewStatusCmd(concatWithCmd("JSON.SET", args)...)
	c.Process(cmd)
	return cmd
}

func jsonDelExecute(c *redisProcessor, args ...interface{}) *redis.IntCmd {
	cmd := redis.NewIntCmd(concatWithCmd("JSON.DEL", args)...)
	c.Process(cmd)
	return cmd
}

func jsonGetExecute(c *redisProcessor, args ...interface{}) *redis.StringCmd {
	cmd := redis.NewStringCmd(concatWithCmd("JSON.GET", args)...)
	c.Process(cmd)
	return cmd
}

func jsonMGetExecute(c *redisProcessor, args ...interface{}) *redis.StringSliceCmd {
	cmd := redis.NewStringSliceCmd(concatWithCmd("JSON.MGET", args)...)
	c.Process(cmd)
	return cmd
}
