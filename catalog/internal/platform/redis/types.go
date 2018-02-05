package redis

//  redis中存储的产品结构
type redisProduct struct {
	SKU          string `redis:"sku"`
	Name         string `redis:"name"`
	Description  string `redis:"description"`
	Manufacturer string `redis:"mfr"`
	Model        string `redis:"model"`
	Price        int64  `redis:"price"`
}

// redis中存储的目录结构
type redisCategory struct {
	Name        string `redis:"name"`
	Description string `redis:"description"`
}
