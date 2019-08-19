# Golang HTTP并发测试

**[Gin](https://github.com/gin-gonic/gin) (use goroutine)  VS  [Ants](https://github.com/panjf2000/ants) (use thread poo)**

对比Gin（使用原生goroutine）和Ants（使用goroutine线程池）



测试程序源码：

本目录下`gin.go` 和 本目录下`panjf2000-ants.go`

为了测试慢执行，测试程序将在sleep一秒后返回。



## 压力测试

### Gin (origin go routine)

#### ab -n 100000 -c 16000

```powershell
C:\Users\XThundering>ab -n 100000 -c 16000 http://127.0.0.1:8080/reverse/gin
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /reverse/gin
Document Length:        2 bytes

Concurrency Level:      16000
Time taken for tests:   33.091 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      12400000 bytes
HTML transferred:       200000 bytes
Requests per second:    3021.95 [#/sec] (mean)
Time per request:       5294.599 [ms] (mean)
Time per request:       0.331 [ms] (mean, across all concurrent requests)
Transfer rate:          365.94 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.2      0       6
Processing:  1681 4818 827.6   5109    5825
Waiting:     1001 3473 1034.7   3447    5272
Total:       1681 4818 827.6   5110    5826

Percentage of the requests served within a certain time (ms)
  50%   5110
  66%   5209
  75%   5252
  80%   5305
  90%   5639
  95%   5682
  98%   5723
  99%   5734
 100%   5826 (longest request)
```



#### 日志文件

输出程序当前goroutine数量：[gin.log](./gin.log)







### Ants (1000 go routine limit)

#### ab -n 100000 -c 16000

```powershell
C:\Users\XThundering>ab -n 100000 -c 16000 http://127.0.0.1:8081/reverse/ants
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8081

Document Path:          /reverse/ants
Document Length:        0 bytes

Concurrency Level:      16000
Time taken for tests:   101.240 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      7500000 bytes
HTML transferred:       0 bytes
Requests per second:    987.75 [#/sec] (mean)
Time per request:       16198.425 [ms] (mean)
Time per request:       1.012 [ms] (mean, across all concurrent requests)
Transfer rate:          72.34 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0      32
Processing:  1496 14644 3279.8  15990   18700
Waiting:     1001 14419 3312.8  15721   18539
Total:       1496 14644 3279.8  15990   18700

Percentage of the requests served within a certain time (ms)
  50%  15990
  66%  16005
  75%  16017
  80%  16021
  90%  16032
  95%  16045
  98%  16077
  99%  16114
 100%  18700 (longest request)
```



#### 日志文件

输出程序当前goroutine数量：[ants_1000.log](./ants_1000.log)







### Ants (8000 go routine limit)

#### ab -n 100000 -c 16000

```powershell
C:\Users\XThundering>ab -n 100000 -c 16000 http://127.0.0.1:8081/reverse/ants
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8081

Document Path:          /reverse/ants
Document Length:        0 bytes

Concurrency Level:      16000
Time taken for tests:   34.690 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      7500000 bytes
HTML transferred:       0 bytes
Requests per second:    2882.69 [#/sec] (mean)
Time per request:       5550.378 [ms] (mean)
Time per request:       0.347 [ms] (mean, across all concurrent requests)
Transfer rate:          211.13 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0       6
Processing:  1520 5087 995.2   5238    6545
Waiting:     1001 3661 1142.0   3634    6115
Total:       1520 5088 995.2   5238    6545

Percentage of the requests served within a certain time (ms)
  50%   5238
  66%   5522
  75%   5653
  80%   5937
  90%   6099
  95%   6418
  98%   6529
  99%   6535
 100%   6545 (longest request)
```



#### 日志文件

输出程序当前goroutine数量：[ants_8000.log](./ants_8000.log)







### Ants (16000 go routine limit)

#### ab -n 100000 -c 16000

```powershell
C:\Users\XThundering>ab -n 100000 -c 16000 http://127.0.0.1:8081/reverse/ants
This is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8081

Document Path:          /reverse/ants
Document Length:        0 bytes

Concurrency Level:      16000
Time taken for tests:   32.276 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      7500000 bytes
HTML transferred:       0 bytes
Requests per second:    3098.29 [#/sec] (mean)
Time per request:       5164.143 [ms] (mean)
Time per request:       0.323 [ms] (mean, across all concurrent requests)
Transfer rate:          226.93 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.3      0       2
Processing:  1618 4693 813.3   4957    5743
Waiting:     1000 3395 1003.0   3375    5178
Total:       1618 4693 813.3   4957    5743

Percentage of the requests served within a certain time (ms)
  50%   4957
  66%   5150
  75%   5163
  80%   5167
  90%   5330
  95%   5550
  98%   5561
  99%   5639
 100%   5743 (longest request)
```



#### 日志文件

输出程序当前goroutine数量：[ants_16000.log](ants_16000.log)







## 数据统计

请求量：100000，并发量：16000

| 库与限制     | 吞吐率(RPS) | 平均响应时间(ms) | 测试耗时(s) | min goroutine | max goroutine |
| ------------ | ----------- | ---------------- | ----------- | ------------- | ------------- |
| Gin          | 3021.95     | 5294.599         | 33.091      | 2             | 12972         |
| Ants - 1000  | 987.75      | 16198.425        | 101.240     | 4             | 33004         |
| Ants - 8000  | 2882.69     | 5550.378         | 34.690      | 4             | 19698         |
| Ants - 16000 | 3098.29     | 5164.143         | 32.276      | 4             | 13997         |







## 参考文章

https://github.com/panjf2000/ants

https://zhuanlan.zhihu.com/p/37754274

