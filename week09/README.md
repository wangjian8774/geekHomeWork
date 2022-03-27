## 1.总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder

​	粘包：由于TCP协议传输数据是以流式数据的形式，数据之间并没有明确的开始和结束，导致TCP无法区分消息的开始和结束。

### fix length

​	这种方式即每次发送固定缓冲区大小，客户端和服务器约定每次发送请求的大小。如客户端发送1024个字节，服务器接受1024个字节。 这样虽然可以解决粘包的问题，但是如果发送的数据小于1024个字节，就会导致数据内存冗余和浪费；且如果发送请求大于1024字节，会出现半包的问题。  



### delimiter based

​	这种是基于定界符来判断是不是一个请求（例如结尾'\n')。客户端发送过来的数据，每次以\n结束，服务器每接受到一个 \n 就以此作为一个请求。这种方式的缺点在于如果数据量过大，查找定界符会消耗一些性能。



### length field based frame decoder


​	这种是在TCP协议头里面写入每次发送请求的长度。 客户端在协议头里面带入数据长度，服务器在接收到请求后，根据协议头里面的数据长度来决定接受多少数据。



## 2.实现一个从 socket connection 中解码出 goim 协议的解码器

​	[链接](https://github.com/wangjian8774/geekHomeWork/tree/main/week09)

