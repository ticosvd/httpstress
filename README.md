# httpstress

htptclient - the generator HTTP traffic . It works with only HTTP service.
```
Usage of ./httpclient:
  -s string
        Set source IP   (default "0.0.0.0")
  -t int
        Set threads (default 1)
  -url string
        Set server  (default "http://10.199.100.100:9443/")
```


Example, start httpclient with thread 1 and get "/" from  http server 80 port with IP address 10.199.20.11  
  
./httpclient -t 1 -url http://10.199.20.11/  
2022/12/01 15:03:23 Thread 1 len response 629 and cont 16.00EB and elapsed  1, throughput 16.00EB  

Example, start httpclient with threads 3 and get "/rfc4" from  http server 80 port  with IP address 10.199.20.11    
./httpclient -t 3 -url http://10.199.20.11/rfc4  
2022/12/01 15:03:47 Thread 3 len response 5988060 and cont 5.71MB and elapsed  41, throughput 1.07GB  
2022/12/01 15:03:47 Thread 2 len response 5988060 and cont 5.71MB and elapsed  48, throughput 939.94MB  
2022/12/01 15:03:47 Thread 1 len response 5988060 and cont 5.71MB and elapsed  55, throughput 826.41MB  



Example, start httpclient with threads 4 and get "/rfc4" from  http server 9443  port  with IP address 10.199.20.11 and sorce IP address 10.20.0.5    

go run httpclient.go -url http://10.199.100.100:9443/rfc4 -t 4 -s 10.20.0.5  
2022/12/01 15:34:38 Thread 1 len response 5988060 and cont 5.71MB and elapsed  73, throughput 617.99MB  
2022/12/01 15:34:38 Thread 4 len response 5988060 and cont 5.71MB and elapsed  77, throughput 586.22MB  
2022/12/01 15:34:38 Thread 3 len response 5988060 and cont 5.71MB and elapsed  81, throughput 563.69MB  
2022/12/01 15:34:38 Thread 2 len response 5988060 and cont 5.71MB and elapsed  88, throughput 515.40MB  


