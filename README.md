
[![Go Report Card](https://goreportcard.com/badge/github.com/zggsong/message?style=flat-square)](https://goreportcard.com/report/github.com/zggsong/message)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.16-61CFDD.svg?style=flat-square)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/zggsong/message)](https://pkg.go.dev/mod/github.com/zggsong/message)


Send notification messages (bark, mail), and support real-time monitoring of configuration file changes

## Install

```shell
go get github.com/zggsong/message
```

## Usage

```shell
touch config.yml
```

write the following contents to the file

```yml
message:
  enabled: true #是否开启发送信息
  type: "bark"  #选择发送信息方式
  bark:
    url: ""
    key: ""
  mail:
    host: ""    #例:smtp.qq.com
    protocol:   #Optional
    port:       #例:587
    username: ""
    password: ""
    from_name: "" #"zgg <zgg@mail.com>"
    to: ""      #分号内空格分开
```

```go
func main() {
	/*配置初始化*/
	conf, err := message.InitConfig()
	if err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}
	message.GLO_CONF = conf

	/*测试发送*/
	var s = message.Service{Body: message.Body{
		Title:   "test title",
		Content: "this is content, time is " + time.Now().Format("2006-01-02 15:04:05"),
	}}
	if err = s.Run(); err != nil {
		log.Fatalf("failed to send message: %v", err)
	} else {
		log.Printf("send message successfully...")
	}

	/*监听配置*/
	for {
		_conf := <-message.GLO_CONF_CH
		log.Printf("config changed: %v", _conf)
		message.GLO_CONF = _conf
	}
}
```

## License

Message © [zggsong](https://github.com/zggsong), Follow [MIT](https://github.com/zggsong/message/blob/master/LICENSE) certificate.