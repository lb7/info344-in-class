package main

import "time"

type entry struct {
	value     string
	expiresAt time.Time
}

//Cache is a TTL cache that is safe for concurrent use
type Cache struct {
	entries map[string]*entry
	//TODO: protect this for concurrent use!
}

//NewCache constructs a new Cache object
func NewCache() *Cache {
	return &Cache{
		entries: map[string]*entry{},
	}
}

//Set adds a key/value to the cache
func (c *Cache) Set(key string, value string, timeToLive time.Duration) {
	c.entries[key] = &entry{value, time.Now().Add(timeToLive)}
}

//Get gets the value associated with a key
func (c *Cache) Get(key string) (string, bool) {
	//TODO: implement this
}
