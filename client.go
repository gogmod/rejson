package rejson

import "github.com/go-redis/redis"

type redisProcessor struct {
	Process func(cmd redis.Cmder) error
}

//Client ...
type Client struct {
	*redis.Client
	*redisProcessor
}

//Pipeline ...
type Pipeline struct {
	redis.Pipeliner
	*redisProcessor
}

//Pipeline ...
func (c *Client) Pipeline() *Pipeline {
	pip := c.Client.Pipeline()
	return ExtendPipeline(pip)
}

//TXPipeline ...
func (c *Client) TXPipeline() *Pipeline {
	pip := c.Client.TxPipeline()
	return ExtendPipeline(pip)
}

//Pipeline ...
func (p *Pipeline) Pipeline() *Pipeline {
	pip := p.Pipeliner.Pipeline()
	return ExtendPipeline(pip)
}

//JSONSet ...
func (cl *redisProcessor) JSONSet(key, path, json string, args ...interface{}) *redis.StatusCmd {
	return jsonSetExecute(cl, append([]interface{}{key, path, json}, args...)...)
}

//JSONDel ...
func (cl *redisProcessor) JSONDel(key, path string) *redis.IntCmd {
	return jsonDelExecute(cl, key, path)
}

//JSONGet ...
func (cl *redisProcessor) JSONGet(key string, args ...interface{}) *redis.StringCmd {
	return jsonGetExecute(cl, append([]interface{}{key}, args...)...)
}

//JSONMGet ...
func (cl *redisProcessor) JSONMGet(key string, args ...interface{}) *redis.StringSliceCmd {
	return jsonMGetExecute(cl, append([]interface{}{key}, args...)...)
}
