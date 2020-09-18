# go-dump
Golang imitates PHP function var_dump
<br>
golang仿php函数var_dump

### Install
 ```
    go get github.com/lgbya/go-dump
 ```

 ### Usage
 ```
    
    data := map[string]string{"a":"apple"}
    formatStr := dump.Format(data) //格式化数据为字符串
    dump.Printf(data) //打印格式数据

    dump.CloseDebug() //关闭测试调试环境不会再打印

 ```
 
 ```
 
  data := map[string]string{"a":"apple","b":"bannel",}
  dump.Printf(data) //打印格式数据
  
  //打印格式
  map[string]string[
          a : apple(string)
          b : bannel(string)
  ]
 ```