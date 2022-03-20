# 1.测试redis benchmark

### 1k:

[root@VM-12-15-centos redis-6.2.6]# src/redis-benchmark -d 10 -t get,set

====== SET ======                          

 100000 requests completed in 1.36 seconds

 50 parallel clients

 10 bytes payload

 keep alive: 1

 host configuration "save": 3600 1 300 100 60 10000

 host configuration "appendonly": no

 multi-thread: no



Latency by percentile distribution:

0.000% <= 0.127 milliseconds (cumulative count 1)

50.000% <= 0.335 milliseconds (cumulative count 59339)

75.000% <= 0.359 milliseconds (cumulative count 75652)

87.500% <= 0.415 milliseconds (cumulative count 87965)

93.750% <= 0.479 milliseconds (cumulative count 93778)

96.875% <= 0.599 milliseconds (cumulative count 96908)

98.438% <= 0.823 milliseconds (cumulative count 98465)

99.219% <= 1.063 milliseconds (cumulative count 99224)

99.609% <= 1.631 milliseconds (cumulative count 99612)

99.805% <= 2.159 milliseconds (cumulative count 99805)

99.902% <= 2.999 milliseconds (cumulative count 99903)

99.951% <= 4.663 milliseconds (cumulative count 99952)

99.976% <= 16.375 milliseconds (cumulative count 99976)

99.988% <= 16.495 milliseconds (cumulative count 99988)

99.994% <= 16.559 milliseconds (cumulative count 99994)

99.997% <= 16.607 milliseconds (cumulative count 99998)

99.998% <= 16.623 milliseconds (cumulative count 99999)

99.999% <= 16.639 milliseconds (cumulative count 100000)

100.000% <= 16.639 milliseconds (cumulative count 100000)



Cumulative distribution of latencies:

0.000% <= 0.103 milliseconds (cumulative count 0)

0.082% <= 0.207 milliseconds (cumulative count 82)

15.804% <= 0.303 milliseconds (cumulative count 15804)

86.900% <= 0.407 milliseconds (cumulative count 86900)

94.885% <= 0.503 milliseconds (cumulative count 94885)

96.982% <= 0.607 milliseconds (cumulative count 96982)

97.684% <= 0.703 milliseconds (cumulative count 97684)

98.329% <= 0.807 milliseconds (cumulative count 98329)

98.913% <= 0.903 milliseconds (cumulative count 98913)

99.135% <= 1.007 milliseconds (cumulative count 99135)

99.270% <= 1.103 milliseconds (cumulative count 99270)

99.358% <= 1.207 milliseconds (cumulative count 99358)

99.409% <= 1.303 milliseconds (cumulative count 99409)

99.508% <= 1.407 milliseconds (cumulative count 99508)

99.562% <= 1.503 milliseconds (cumulative count 99562)

99.601% <= 1.607 milliseconds (cumulative count 99601)

99.641% <= 1.703 milliseconds (cumulative count 99641)

99.684% <= 1.807 milliseconds (cumulative count 99684)

99.717% <= 1.903 milliseconds (cumulative count 99717)

99.769% <= 2.007 milliseconds (cumulative count 99769)

99.795% <= 2.103 milliseconds (cumulative count 99795)

99.908% <= 3.103 milliseconds (cumulative count 99908)

99.958% <= 5.103 milliseconds (cumulative count 99958)

100.000% <= 17.103 milliseconds (cumulative count 100000)



Summary:

 throughput summary: 73367.57 requests per second

 latency summary (msec):

​     avg    min    p50    p95    p99    max

​    0.369   0.120   0.335   0.511   0.943   16.639

====== GET ======                          

 100000 requests completed in 1.31 seconds

 50 parallel clients

 10 bytes payload

 keep alive: 1

 host configuration "save": 3600 1 300 100 60 10000

 host configuration "appendonly": no

 multi-thread: no



Latency by percentile distribution:

0.000% <= 0.103 milliseconds (cumulative count 1)

50.000% <= 0.327 milliseconds (cumulative count 50493)

75.000% <= 0.351 milliseconds (cumulative count 76961)

87.500% <= 0.383 milliseconds (cumulative count 88273)

93.750% <= 0.439 milliseconds (cumulative count 94322)

96.875% <= 0.503 milliseconds (cumulative count 96906)

98.438% <= 0.695 milliseconds (cumulative count 98461)

99.219% <= 0.855 milliseconds (cumulative count 99231)

99.609% <= 0.983 milliseconds (cumulative count 99614)

99.805% <= 1.311 milliseconds (cumulative count 99810)

99.902% <= 2.823 milliseconds (cumulative count 99903)

99.951% <= 4.031 milliseconds (cumulative count 99952)

99.976% <= 4.311 milliseconds (cumulative count 99976)

99.988% <= 4.543 milliseconds (cumulative count 99988)

99.994% <= 4.655 milliseconds (cumulative count 99994)

99.997% <= 4.687 milliseconds (cumulative count 99997)

99.998% <= 4.711 milliseconds (cumulative count 99999)

99.999% <= 4.727 milliseconds (cumulative count 100000)

100.000% <= 4.727 milliseconds (cumulative count 100000)



Cumulative distribution of latencies:

0.001% <= 0.103 milliseconds (cumulative count 1)

0.100% <= 0.207 milliseconds (cumulative count 100)

14.117% <= 0.303 milliseconds (cumulative count 14117)

91.580% <= 0.407 milliseconds (cumulative count 91580)

96.906% <= 0.503 milliseconds (cumulative count 96906)

97.955% <= 0.607 milliseconds (cumulative count 97955)

98.507% <= 0.703 milliseconds (cumulative count 98507)

98.990% <= 0.807 milliseconds (cumulative count 98990)

99.420% <= 0.903 milliseconds (cumulative count 99420)

99.648% <= 1.007 milliseconds (cumulative count 99648)

99.733% <= 1.103 milliseconds (cumulative count 99733)

99.778% <= 1.207 milliseconds (cumulative count 99778)

99.804% <= 1.303 milliseconds (cumulative count 99804)

99.845% <= 1.407 milliseconds (cumulative count 99845)

99.879% <= 1.503 milliseconds (cumulative count 99879)

99.891% <= 1.607 milliseconds (cumulative count 99891)

99.897% <= 1.703 milliseconds (cumulative count 99897)

99.900% <= 1.807 milliseconds (cumulative count 99900)

99.923% <= 3.103 milliseconds (cumulative count 99923)

99.958% <= 4.103 milliseconds (cumulative count 99958)

100.000% <= 5.103 milliseconds (cumulative count 100000)



Summary:

 throughput summary: 76277.65 requests per second

 latency summary (msec):

​     avg    min    p50    p95    p99    max

​    0.349   0.096   0.327   0.455   0.815   4.727

### 5k:

[root@VM-12-15-centos redis-6.2.6]# src/redis-benchmark -d 5120 -t get,set
====== SET ======                                                   
  100000 requests completed in 1.48 seconds
  50 parallel clients
  5120 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.103 milliseconds (cumulative count 1)
50.000% <= 0.367 milliseconds (cumulative count 52328)
75.000% <= 0.407 milliseconds (cumulative count 76816)
87.500% <= 0.471 milliseconds (cumulative count 87637)
93.750% <= 0.591 milliseconds (cumulative count 93904)
96.875% <= 0.799 milliseconds (cumulative count 96905)
98.438% <= 0.927 milliseconds (cumulative count 98457)
99.219% <= 1.063 milliseconds (cumulative count 99232)
99.609% <= 1.271 milliseconds (cumulative count 99615)
99.805% <= 1.583 milliseconds (cumulative count 99806)
99.902% <= 2.023 milliseconds (cumulative count 99903)
99.951% <= 2.231 milliseconds (cumulative count 99955)
99.976% <= 2.335 milliseconds (cumulative count 99976)
99.988% <= 2.431 milliseconds (cumulative count 99988)
99.994% <= 2.543 milliseconds (cumulative count 99994)
99.997% <= 2.591 milliseconds (cumulative count 99997)
99.998% <= 2.623 milliseconds (cumulative count 99999)
99.999% <= 2.647 milliseconds (cumulative count 100000)
100.000% <= 2.647 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.001% <= 0.103 milliseconds (cumulative count 1)
0.038% <= 0.207 milliseconds (cumulative count 38)
1.693% <= 0.303 milliseconds (cumulative count 1693)
76.816% <= 0.407 milliseconds (cumulative count 76816)
90.325% <= 0.503 milliseconds (cumulative count 90325)
94.259% <= 0.607 milliseconds (cumulative count 94259)
95.773% <= 0.703 milliseconds (cumulative count 95773)
97.011% <= 0.807 milliseconds (cumulative count 97011)
98.146% <= 0.903 milliseconds (cumulative count 98146)
99.033% <= 1.007 milliseconds (cumulative count 99033)
99.353% <= 1.103 milliseconds (cumulative count 99353)
99.540% <= 1.207 milliseconds (cumulative count 99540)
99.645% <= 1.303 milliseconds (cumulative count 99645)
99.714% <= 1.407 milliseconds (cumulative count 99714)
99.765% <= 1.503 milliseconds (cumulative count 99765)
99.816% <= 1.607 milliseconds (cumulative count 99816)
99.840% <= 1.703 milliseconds (cumulative count 99840)
99.861% <= 1.807 milliseconds (cumulative count 99861)
99.875% <= 1.903 milliseconds (cumulative count 99875)
99.899% <= 2.007 milliseconds (cumulative count 99899)
99.919% <= 2.103 milliseconds (cumulative count 99919)
100.000% <= 3.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 67430.88 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.405     0.096     0.367     0.647     0.999     2.647
====== GET ======                                                   
  100000 requests completed in 1.48 seconds
  50 parallel clients
  5120 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.119 milliseconds (cumulative count 1)
50.000% <= 0.367 milliseconds (cumulative count 59234)
75.000% <= 0.391 milliseconds (cumulative count 76121)
87.500% <= 0.439 milliseconds (cumulative count 87912)
93.750% <= 0.519 milliseconds (cumulative count 93930)
96.875% <= 0.735 milliseconds (cumulative count 96895)
98.438% <= 0.911 milliseconds (cumulative count 98481)
99.219% <= 1.055 milliseconds (cumulative count 99233)
99.609% <= 1.399 milliseconds (cumulative count 99615)
99.805% <= 1.975 milliseconds (cumulative count 99805)
99.902% <= 6.679 milliseconds (cumulative count 99903)
99.951% <= 10.047 milliseconds (cumulative count 99952)
99.976% <= 10.863 milliseconds (cumulative count 99976)
99.988% <= 11.071 milliseconds (cumulative count 99988)
99.994% <= 11.175 milliseconds (cumulative count 99994)
99.997% <= 11.239 milliseconds (cumulative count 99997)
99.998% <= 11.279 milliseconds (cumulative count 99999)
99.999% <= 11.295 milliseconds (cumulative count 100000)
100.000% <= 11.295 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.041% <= 0.207 milliseconds (cumulative count 41)
1.048% <= 0.303 milliseconds (cumulative count 1048)
81.747% <= 0.407 milliseconds (cumulative count 81747)
93.282% <= 0.503 milliseconds (cumulative count 93282)
95.695% <= 0.607 milliseconds (cumulative count 95695)
96.628% <= 0.703 milliseconds (cumulative count 96628)
97.518% <= 0.807 milliseconds (cumulative count 97518)
98.395% <= 0.903 milliseconds (cumulative count 98395)
99.086% <= 1.007 milliseconds (cumulative count 99086)
99.330% <= 1.103 milliseconds (cumulative count 99330)
99.483% <= 1.207 milliseconds (cumulative count 99483)
99.558% <= 1.303 milliseconds (cumulative count 99558)
99.619% <= 1.407 milliseconds (cumulative count 99619)
99.670% <= 1.503 milliseconds (cumulative count 99670)
99.726% <= 1.607 milliseconds (cumulative count 99726)
99.758% <= 1.703 milliseconds (cumulative count 99758)
99.782% <= 1.807 milliseconds (cumulative count 99782)
99.795% <= 1.903 milliseconds (cumulative count 99795)
99.809% <= 2.007 milliseconds (cumulative count 99809)
99.828% <= 2.103 milliseconds (cumulative count 99828)
99.900% <= 3.103 milliseconds (cumulative count 99900)
99.935% <= 7.103 milliseconds (cumulative count 99935)
99.950% <= 8.103 milliseconds (cumulative count 99950)
99.956% <= 10.103 milliseconds (cumulative count 99956)
99.990% <= 11.103 milliseconds (cumulative count 99990)
100.000% <= 12.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 67796.61 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.401     0.112     0.367     0.567     0.991    11.295

# 2.写入kv数据，分析内存空间

结果如下：

```
# 10000
(2.69M - 855.02K - 1B * 10000) / 10000 = 194B

# 50000
(5.51M - 1.21M - 1B *  50000) / 50000 = 89B

# 100000
(9.06M - 1.21M - 1B * 100000) / 100000 = 81B

# 200000
(9.04M - 1.21M - 1B * 200000) / 200000 = 40B

# 300000
(7.79M - 1.21M - 1B * 300000) / 300000 = 21B

# 400000
(8.82M - 1.21M - 1B * 400000) / 400000 = 19B

# 500000
(9.52M - 1.21M - 1B * 500000) / 500000 = 16B
```

结论：

相同长度的value在写入数量越多情况下，平均每个value占用内存更多