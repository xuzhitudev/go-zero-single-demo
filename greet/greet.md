### 1. "form信息输入"

1. route definition

- Url: /from
- Method: GET
- Request: `Request`
- Response: `Response`

2. request definition



```golang
type Request struct {
	Name string `json:"name"`
}
```


3. response definition



```golang
type Response struct {
	Message string `json:"message"`
}
```

### 2. "根据id获取用户信息"

1. route definition

- Url: /user/:id
- Method: GET
- Request: `UserInfoRequest`
- Response: `UserInfoResponse`

2. request definition



```golang
type UserInfoRequest struct {
	Id string `path:"id"`
}
```


3. response definition



```golang
type UserInfoResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Address string `json:"address"`
	AddressInfo string `json:"nnn"`
	AaaaddressInfo string `json:"nnnddd"`
}
```

### 3. "用户信息注册"

1. route definition

- Url: /user/save
- Method: POST
- Request: `UserRegisterRequest`
- Response: `UserRegisterRequest`

2. request definition



```golang
type UserRegisterRequest struct {
	Id string `form:"id"`
	Name string `form:"name"`
	Age int `form:"age"`
	Address string `form:"address"`
}
```


3. response definition



```golang
type UserRegisterRequest struct {
	Id string `form:"id"`
	Name string `form:"name"`
	Age int `form:"age"`
	Address string `form:"address"`
}
```

