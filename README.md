# tencent-lbs
[腾讯位置服务](https://lbs.qq.com/)

# Author

zenghongru@gmail.com

# Usage
```go
package main

import (
	"context"
	"fmt"
	txlbs "github.com/xiaojiaoyu100/tencent-lbs"
)

func main() {
	// c需要缓存下来，不能频繁的创建client
	c, err := txlbs.New(
		txlbs.WithKey("your_key"),
		txlbs.WithSecretKey("your_secret_key"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := c.DistrictSearch(context.TODO(), "110000,111100")
	if err != nil {
		fmt.Println(err)
		return
	}
	if !resp.Success() {
		fmt.Println(resp.Message)
		return
	}
}
```