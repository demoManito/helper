## conversion 转换相关

1. [upperlower.go](./upperlower.go) 大小写转化
    ```bash
    # 小写转大写 lower to upper
    helper tu user_id => UserId
    helper tu user_id user_id => UserId, UserId
   
    # 大写转小写 upper to lower 
    helper tl UserID => user_id
    helper tl UserID userID => user_id, user_id 
    ```
   
2. [json.go](./json.go) 格式化 json
   > 将其输出到剪切板中