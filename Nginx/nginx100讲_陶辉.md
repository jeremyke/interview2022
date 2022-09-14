# Nginx核心100讲.md

# 一、初识Nginx

## 1.1 nginx应用场景

## 1.2 ginx优点

## 1.3 Nginx组成

## 1.4 Nginx配置语法

## 1.5 Nginx 命令行

## 1.6 日志

web日志分析工具：goAccess

## 1.7 SSl安全协议

- 对称加密

基于异或算法，明文可以转为密文，密文也可以转为明文，而且性能好，只需要一次便历过程。

- 非对称加密

自己发出去的文本用私钥加密，接收方使用公钥加密；反之，接收方回复消息使用公钥加密，自己使用私钥解密。

在Nginx上可以设置为：

```
ssl_verify_client on;#以便 OCSP 验证工作
ssl_ocsp on;#启用客户端证书链的 OCSP 验证
resolver 192.0.2.1;#解析器应指定为解析 OCSP 响应器主机名

```

浏览器获取到证书后如何生效，需要验证证书链：

站点证书由三部分构成（根证书、二级证书、主证书），操作系统的根证书很难修改，大部分浏览器（除firebox）使用的是操作系统的证书库。所以浏览器在验证证书是否有效时，除了验证nginx发过来的两个证书（二级证书和主站点证书）是否过期外，还要要在根证书是否有效且被认证。

### 1.7.1 TLS的通信过程

第一步：浏览器向服务器发送clinet hello消息，告诉服务器我支持哪些加密算法；
第二步：服务器把最偏向的加密算法发送给客户端，发送server hello消息，告诉服务器最终选择哪个安全套件；
第三步：服务器向客户端发送证书链；
第四步：客户端验证服务器相关证书；
第五步：服务器发送server hello done,且在第五步前向客户的发送加密算法的公共参数；
第六步：浏览器根据公共参数生成自己的私钥，再把公钥发送给服务器；
第七步：服务器生成自己的一对公钥和私钥，用自己的私钥和客户端发来的私钥，生成双方加密的密钥。
第八步：浏览器根据服务器发来的公钥和自己生成的私钥也会生成双方加密的密钥，通过非对称加密，二者生成的密钥是相同的。
第九步：服务器使用生成的密钥加密发送的消息，传给浏览器。

**Nginx对加密算法的优化：**

对于小文件，Nginx需要优化非对称加密算法,适当弱化密码强度；
对于大文件，需要考虑优化对称加密算法（AES）。

### 1.7.2 使用免费SSL证书把Http网站改造为Https网站

```bash
[root]:yum install python2-certbot-nginx
[root]:certbot --nginx --nginx-server-root=conf目录 -d 需要安装证书的server name

```

## 1.8 基于openResty使用lua实现简单服务

### 1.8.1 下载

openresty.org-->下载-->源码发布--下载-->解压

### 1.8.2 分析目录结构

### 1.8.3 编译

编译一个基本的openresty

```bash
.configure
make&make install
```

### 1.8.4 添加lua代码

### 1.8.5 运行

# 二、Nginx架构基础

## 2.1 Nginx的进程结构

Nginx包含master进程和子进程，子进程又分为两大类，Cache相关进程和worker进程，子进程间的通信是通过共享内存来解决的。
master进程：用来管理worker进程，负责监控worker进程正常工作，是否需要重新加载配置文件等。
缓存：是多个worker进程共享的，同时也会被cache manager（缓存的关联）和cache loader（缓存的载入）进程使用。cache manager和cache loader进程使用是用于后端发来的动态请求做缓存的来使用的。
为了保证Nginx的高可用性，nginx被设计为多进程的模式，因为多线程不同线程会共享一块内存空间，线程间相互影响，如果一个第三方模块导致地址越界等问题会使得整个Nginx全部挂掉。

## 2.2 使用信号管理Nginx父子进程

## 2.3 reload流程

## 2.4 热升级流程

## 2.5 优化的关闭worker进程

优化的关闭只针对Http请求，对于websocket,TCP，UDP，Nginx无法得知worker是否在处理请求。

## 2.6 网络收发和Nginx事件间的对应关系

Nginx是一个事件驱动的框架，事件是指网络事件，一个网络连接对应两个事件（读事件和写事件）

### 2.6.1 网络传输

### 2.6.2 TCP流和报文

## 2.7 Nginx事件循环

### 2.7.1 epool优劣已经原理

上图nginx等待服务器内核的事件队列使用epool来处理的。

## 2.8 Nginx的请求切换

nginx是用户态直接切换的，除非操作系统分给worker的时间分片到期了，否则一直工作，所以将worker的优先级调为-19，可以使得操作系统给worker分配更多的时间分片，提高Nginx性能。

## 2.9 同步和异步、阻塞和非阻塞的区别

阻塞和非阻塞是线程在访问某个资源时，数据是否准备就绪的一种处理方式。

阻塞方法：操作系统或者底层C库提供的方法或者是一个系统调用，这个方法可能是我的进程进入sleep状态（当前条件不满足，操作系统把我的进程切换到另外一个进程）。
非阻塞方法：我们调用该方法永远不会在我们时间分片未用完时，切换到另外一个进程。

同步和异步是用户态的调用方式而言。

同步：调用方法时，需要等待返回结果。
异步：调用方法后，无需等待返回结果，被调用方法处理完成后后主动通知给调用方。（支付回调）

阻塞调用：

非阻塞调用：

非阻塞调用下的同步和异步：

openResty使得我们可以通过写同步的的方法，实际上以异步的来执行。

## 2.10 Nginx模块

### 2.10.1 Nginx模块分类

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/c9250cc7-b1ba-4adc-9ce7-e94f43a5ab54/Untitled.png)

## 2.11 Nginx连接池

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/917a7881-b45f-4b35-840d-6efb75164abb/Untitled.png)

每一个连接对应2个事件（读事件和写事件），如上图是通过数组序号来配合使用的。消耗的内存如下：

一个connection(232)+2个event(96*2) = 424字节。设置的连接数越多，消耗的内存就越大。

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e45c3fbe-9f2e-4967-acd4-ef7772d28ace/Untitled.png)

## 2.12 Nginx内存池

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/7f33af90-672c-412d-b4d3-a28d38907a60/Untitled.png)

**connection_pool_size:**

- Default：connection_pool_size:256|512

内存池配置512字节，并不意味着只能分配512字节，当内存超过预分配内存大小时，是可以继续分配的。提前分配内存空间，减小分配的资源消耗。

**request_pool_size:**

- Default：request_pool_size:4k

**为什么请求内存池大于远远大于连接内存池？**
因为连接需要存储的上下文信息很少，只需要帮助后面的请求读取最初一部分字节就行；对于请求而言，需要保存大量上下文信息，比如：url和header。其对性能的影响比较小，极端场景下，url特别长，可以修改配置，增大请求内存池预分配的空间大小；通常情况下，url和header都很小，可以考虑降低请求内存池的预分配空间大小，最大化Nginx的并发量。

内存池对减小内存碎片和第三方模块开发是很有意义的。

## 2.13 Nginx进程间通信方式

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/a11f77b0-c6f0-4150-a691-2c4fd273ddca/Untitled.png)

哪些官方模块使用了共享内存呢？

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/63317730-a95c-4e2d-905d-fde76e3b2de5/Untitled.png)

Nginx_http_lua_api是OpenResty的核心模块：

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/867c507d-635e-4a88-9208-f6bc34319d17/Untitled.png)

在如上代码中，同时使用了rbtree和链表：

```bash
lua_shared_dict dogs 10m;//使用红黑树来保存每一个key-value,每一个节点是它的key,节点值就是value。
//当内存大于10m时，使用lru方式淘汰，最先设置了key-value就会被淘汰，这就说明每个key-value连在了一起形成了一张链表。
```

### 2.13.1 Slab内存分配管理

如何把一整块内存切割成小块分配给每个红黑树节点使用的？Slab内存管理

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/753a52ef-f8f1-4a84-bcd7-a2d2bf9e4c57/Untitled.png)

Slab会把共享内存切割很多页面（4K），每个页面被切分为不同的slot（不同的slot分配内存空间不同128|256|512,乘2方式向上增长)。

Bestfit:比如30字节的内存，会被分配到32字节的slot。

### 2.13.2 查看slot使用状态

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/28f444f0-5363-431b-bfd9-1e8629d88fd7/Untitled.png)

### 2.13.3 在OpenResty上使用tengine的slab_stat模块查看共享内存分配情况

编译安装（在编译安装OpenResty时，使用add-module把tengine的slab_stat模块安装进来）：

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/4e757b3f-72f9-4d93-99ad-601843c018e1/Untitled.png)

案例：

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/13c8f45d-954d-46c6-b2b0-76d4ae925867/Untitled.png)

## 2.14 Nginx数据容器

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/8ad27719-507b-4f0b-9a56-ff1aaf74b340/Untitled.png)

Nginx数组（nginx_array_T)：多块连续内存，每块连续内存中可以存放许多元素。

链表：nginx_list_T

队列：nginx_queue_T

### 2.14.1 哈希表

哈希表：nginx_hash_elt_T

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/b0e1f946-e954-4f29-a97c-37e69703dae4/Untitled.png)

哈希表每个元素占用连续的内存。value是指针，指向用户数据，len长度，name就是hash的key。

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/958a7dfd-ac7b-453b-9dc0-48ee98293f83/Untitled.png)

Hash表应用于静态不变的内容。

Max size:控制了最大的hash表Bucket的个数，并不是实际最大hash表bucket的个数。

Bucket size:每个bucket的大小，向上对齐(cache line)。比如64位操作系统,操作系统每次读取64个字节，但bucket size配置60，这个时候实际是Nginx会设置为64，与操作系统对齐。,bucket size配置尽量不要超过64字节，以免占用内存过高。

### 2.14.2 红黑树

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/7b0d336c-a2b8-42ef-89f7-f721f90d16dc/Untitled.png)

红黑树本身是一个二叉树，包含左节点和右节点。其次是一个查找二叉树，左节点比右节点小。可能会退化成一个链表，如右图。

**红黑树优点：**

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/4e9f25c6-fb40-4f5f-860c-a34183e12f46/Untitled.png)

**使用红黑树的模块：**

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/cc0395dd-32b1-4af1-86d0-6708c0602f7a/Untitled.png)

## 2.15 动态模块

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/b4945b27-c4cd-4958-bdf5-ea2ab7097c07/Untitled.png)

使用案例：

(1)查看哪些模块支持动态模块：./configure —help | more

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/8abc93c4-3bca-47cb-b106-dc3f8ea445ea/Untitled.png)

(2)把一个模块动态编译到nginx

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/a8feba50-dfdc-4a41-bbfd-454fbcf05c01/Untitled.png)

(3)打开动态模块配置

```bash
load_module module/nginx_http_image_filter_module.so
image_filter resize 15 10;
```

## 2.16 配置指令

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/55055274-05a0-4295-bc3f-174dfbf5f427/Untitled.png)

### 2.16.1 什么是指令的context?

该条配置指令能处于的配置块。

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/8defed08-20b3-41d2-8d04-e4f3333b19e0/Untitled.png)

上图表明：log_format能在http模块上配置；access_log能在http,server…等模块配置。

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/3e3ad074-4ccd-4f09-ae30-c94ec0e9405b/Untitled.png)

一个配置存在多个配置块是是可以合并的。值指令可以合并；动作类指令不可以合并。判断依据：看该条配置的生效阶段。

**值指令向上覆盖：**

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/44d71509-0b9e-48b9-a9a9-6d22d9c7dc31/Untitled.png)

http模块指令合并规则：

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/6fdca059-3a54-40b2-8c80-aaee5fb50666/Untitled.png)

# 三、详解HTTP模块

## 3.1 Listen指令

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/3f065581-d6fe-49db-86bc-75c48e35734b/Untitled.png)

## 3.2 接收请求的事件模块

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/4ae1e07e-f401-47b5-a62f-74ae56ffd7ff/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e2089572-2a2e-4f6b-a96e-258cc3e40019/Untitled.png)

## 3.3 正则表达式：

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/ad05b7e8-81b7-4453-87a2-fd6a96a9af27/Untitled.png)

## 3.4 如何找到处理请求的server指令块

### 3.4.1 server_name指令

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/d67f70ab-59bd-4edf-8019-7bcd215ec3eb/Untitled.png)

**server_name多域名：**

例如：

```bash
server{
		server_name [aa.com](http://aa.com) [bb.com](http://bb.com); #其中aa.com是主域名
		server_name_in_redirect off;
		redirect 302 /redirect;
} 
```

如上配置：如果访问bb.com，会直接跳转到bb.com/redirect

如果配置：

```bash
server{
		server_name [aa.com](http://aa.com) [bb.com](http://bb.com); #其中aa.com是主域名
		server_name_in_redirect on;
		redirect 302 /redirect;
} 
```

如上配置：如果访问bb.com,会302到aa.com/redirect

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/07160cf5-a48f-4c70-9728-fb53b48c0312/Untitled.png)

**server_name匹配顺序：**

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/77b2760b-2372-4057-92b7-4679d30dbc10/Untitled.png)

## 3.5 HTTP请求11个阶段

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/75225120-53a8-4857-9ccb-1bacaa9ba76d/Untitled.png)

HTTP请求11个阶段的执行顺序：

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/119e805c-e10d-429a-8e62-1f745076e1b5/Untitled.png)

上图说明：

（1）每一个阶段会有多个模块得到执行。

（2）limit_req和limit_conn都在preaccess阶段得到执行，但是如果limit_req阻止了本次请求，就会直接返回，limit_conn就得不到执行。

（3）access阶段的access执行通过后，后续auth_basic和auth_request两个阶段不会执行，直接跳到precontent阶段，同理content阶段也是如此。

### 3.5.1  postread阶段

（1）realip模块

**把realip模块编译进Nginx模块？**

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/be9557ef-eaba-4933-832f-519151dded21/Untitled.png)

三个指令：

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/fa6d71a6-dea0-497e-b46c-c21c122727c8/Untitled.png)

set_real_ip_from：设置可信赖的ip(如本机IP或者机器某台机器IP)，找到用户真实ip

real_ip_header:ip值从哪个参数中取（X-Real_Ip |X-Forwarded-For）

real_ip_recursive:环回地址，取 X-Forwarded-For最后一个地址如果和客户端地址一样就pass掉，取上一个地址。

(2)**如何拿到真实的用户IP地址？**

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/0a7362a9-54b0-44aa-bc28-bcd01a4fc1ea/Untitled.png)

说明：当网络请求过程中，存在很多反向代理服务器时，Nginx会把用户真实的ip写在X-Real_IP中，而每经过一层反向代理，会吧上一层IP追加在X-Forwarded-For中。

**拿到真实用户IP后如何使用？**

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/4a52b7ce-5406-4dc6-90be-14eb50fec310/Untitled.png)

### 3.5.2  rewrite阶段

（1）rewrite模块return指令

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/cfd3944f-a95c-4f07-a02f-9eb648ba27ff/Untitled.png)

error_page:收到某个特定的返回码时，重定向某个url或者给用户返回特定内容。

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/63c1e107-ccef-4806-a78e-84b0105272e2/Untitled.png)

return示例：

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/abd192ec-6720-4ad8-921b-095d57a34d8c/Untitled.png)

server与location块下的return指令关系？

server块的return先于location块的return执行。

return与error_page指令的关系？

访问不存在的资源，执行error_page指令；如果执行了return指令，error_page得不到执行。

(2) rewrite指令

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/411c71af-9a8e-4b34-910c-13c5d9f7b48c/Untitled.png)

案例：

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/789d3189-8cef-4006-a594-e4cc319742f7/Untitled.png)

问题1：rewrite指令优先级高于return

问题2：依次返回“test3”、“test3”、”third!”

问题3：不携带break，就往下执行到return 200 ‘second’模块

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/ba57cf0e-2d11-4761-8a78-e0c02d349686/Untitled.png)

问题：依次返回301、302、302、301

**rewrite日志记录：rewrite_log:on就可以开启了**

（3）if指令

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/6b854864-67a1-4ba0-a3e5-6e7dbaf770f2/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/a8d9f87f-d0e2-4279-a433-0b8fd9ec3aaa/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/d9bb7673-75ec-4b64-9d7d-294900bc580f/Untitled.png)

### 3.5.3 find_config阶段

（1）location指令

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/c00770b5-7d8b-4014-b5da-b7b79c1819d2/Untitled.png)

匹配规则：

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/31f0c8bd-093a-4850-a163-73afdd980204/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/b52720a0-1301-4bbb-bc3d-0f1ce0e2b919/Untitled.png)

匹配结果：

/Test1:5,6;

/Test1/:1,3,5;

/Test1/Test2:2,4,5

/Test1/Test2/:4,5

/test1/Test2：2

匹配顺序

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/90f50042-e759-40dd-8963-7a5b2d398ed9/Untitled.png)

同上问题，返回什么？

/Test1:exact match!

/Test1/:stop regular expressions match!

/Test1/Test2:lonest regular expressions match!

/Test1/Test2/:lonest regular expressions match!

/test1/Test2:lonest regular expressions match!

### 3.5.4 preaccess阶段

(1)limit_conn指令

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/8d67e221-0d0c-44a3-b141-6a7932858bbe/Untitled.png)

key一般为用户的IP地址(remote_addr变量)。

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/fd083517-a968-4706-bd74-c6173f5ae95a/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/4af09e98-faea-4d17-b564-2c9dc36b218a/Untitled.png)

$binary_remote_addr 关键字，这里是限制同一客户端ip地址；

限制连接数:

要限制连接，必须先有一个容器对连接进行计数，在http段加入如下代码：

"zone=" 给它一个名字，可以随便叫，这个名字要跟下面的 limit_conn 一致

$binary_remote_addr = 用二进制来储存客户端的地址，1m 可以储存 32000 个并发会话

上图配置说明：现在某个ip的并发数为1，超过1就返回500，记warn级别日志，网站返回速率为50字节/s

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/30006aac-4c79-418a-b0c8-b0010c828329/Untitled.png)

(1)limit_req指令

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/a99b6a7a-c9fc-40e2-b811-632ca697b98c/Untitled.png)

```bash

```

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/8c82ef5c-0041-4f2a-ac10-40bc834efdd4/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/2856e383-2543-4923-b99a-325a22651498/Untitled.png)

leaky bucket算法：

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/d586e19d-10b7-4f11-9caf-4e60e330f6ce/Untitled.png)

```bash
limit_req_zone:$binary_remote_addr zone=one:10m rate=2r/m;
location / {
	limit_req zone=one burst=3 nodely;
}
```

**第一段配置参数：**

- $binary_remote_addr ：表示通过remote_addr这个标识来做限制，“binary_”的目的是缩写内存占用量，是限制同一客户端ip地
- zone=one:10m：表示生成一个大小为10M，名字为one的内存区域，用来存储访问的频次信息
- rate=2r/s：表示允许相同标识的客户端的**访问频次**，这里限制的是每秒1次，即每秒只处理一个请求，还可以有比如**30r/m的，即限制每2秒访问一次，即每2秒才处理一个请求。**

**第二段配置参数：**

- zone=one ：设置使用哪个配置区域来做限制，与上面limit_req_zone 里的name对应
- burst=5：重点说明一下这个配置，burst爆发的意思，这个配置的意思是设置一个大小为5的缓冲区当有大量请求（爆发）过来时，超过了**访问频次**限制的**请求可以先放到这个缓冲区内等待，但是这个等待区里的位置只有5个**，超过的请求会直接报503的错误然后返回。
- nodelay：
- **如果设置，会在瞬时提供处理(burst + rate)个请求的能力**，请求超过**（burst + rate）**的时候就会直接返回503，永**远不存在请求需要等待的情况**。（这里的rate的单位是：r/s）
- 如果没有设置，则所有请求会依次等待排队

`说明：limit_conn模块优先级大于limit_req。`

### 3.5.5 access阶段

- access模块（对用户ip校验）

判断请求是否可以继续向下访问。

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/8ff9b5ac-8df7-47cc-9983-6fa9e18f01d2/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/2769b3f1-2fe7-4659-ba12-d39ce6aa615b/Untitled.png)

- auth_basic模块（对用户用户名+密码验证）

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/08cd3670-3857-4ebd-802b-69e7cca2e903/Untitled.png)

示例：

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/906b57e9-d1a2-44ff-88ae-5d387dbc9e2d/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/9af9174c-1c0f-426e-a144-591555dba702/Untitled.png)

- auth_request模块（用于做统一用户鉴权系统）

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/23751c43-6c75-4e89-8441-77095ba4c6c4/Untitled.png)

- satisfy指令

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/49f34f07-072d-449f-8fab-61a400b0b3df/Untitled.png)

satisfy all：access的3和指令均放行这个请求，这个请求才向下执行，任何一个拒绝，400或500返回。

       anny:access的3和指令有一个指令放行这个请求，这个请求就向下执行。

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/22573660-1906-43a6-b961-28c2774c958f/Untitled.png)

1. 不会。
2. 有影响。
3. 可以
4. 可以
5. 没有机会输入

### 3.5.6 precontent阶段

- try_file指令

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/8ce70a6d-0f97-4149-8a6c-41f86c0c68ff/Untitled.png)

- mirror模块（做流量拷贝）

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/7678a002-f487-4441-a9df-1a48472e31d4/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/debe7f54-b641-473a-b9a0-dff3143872ee/Untitled.png)

### 3.5.7 content阶段

- static模块

（1）root和alias指令

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/69f88a38-ab85-4036-9083-9c5601449d8d/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/f03cd8a6-ef57-406c-b157-96d789389579/Untitled.png)

/root :404

/root/1.txt :—>html/first/root/1.txt 也是404

/alias:—>html/index.html 200

/alias/1.txt:—>html/first/1.txt 200

（2）3个nginx变量

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/0d81728e-2b36-4f90-8a2d-6d90a316e025/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/3709bc4b-d37e-425c-80cc-2aae56fa9b67/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/c1492df0-e11b-4241-bd60-0e403fdbe287/Untitled.png)

（3）static模块对url不以斜杠结尾却访问目录的的做法

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/15960ef9-74c0-4eac-8898-ad5003e0a5c7/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/aa6bc411-18b9-4e31-a1ac-b2c7633df2ed/Untitled.png)

- index模块和autoindex模块

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/2a6d5d49-54b5-422a-a400-ee168d9af2c9/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/0e140c72-cb0c-463d-a8e7-16d9521fd49e/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e8bb2e31-3622-4cc6-b629-09e87568d6b9/Untitled.png)

- concat模块（合并小文件，提升网络性能）

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/0e35cfd6-546b-40d8-ab46-7152deb41149/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/a9499eab-315e-4772-ae68-22470a7eb0cb/Untitled.png)

### 3.5.8 log阶段

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/a39c0626-4a38-4e95-a9e4-4492ad013dd5/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/16d0e62d-c6a6-4b17-b3e7-b727406e6f3f/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/51dc9015-5275-4212-87df-9fde19555788/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/f2a17073-54a1-46d0-8928-c178d08d3d41/Untitled.png)

### 3.5.9 HTTP过滤模块

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/80c5ebb5-eb18-4ec6-8336-b84c9755b7dc/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/56d95e84-b652-4082-b97f-8a7c4b87a69c/Untitled.png)

- sub模块

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/1d2ad3bd-268b-4fb2-ae46-ba8823b5a24f/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/fb17f705-a04a-4431-8349-957ac4904925/Untitled.png)

示例：

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/43a56e15-33c1-4575-b7cf-fdda7eef6427/Untitled.png)

- addition模块

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/4aae21ff-df93-4b07-842a-6e20a39b3954/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/5178115f-610a-47a6-acf3-a0b80223b6d1/Untitled.png)

## 3.6 HTTP的变量

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/f9d19adc-ca9c-4939-bb34-70aeba872cc8/Untitled.png)

### 3.6.1 HTTP请求相关变量

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/0a80686d-0255-4bc4-afae-19c257ff9cd8/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/f1f9f9df-1be1-4c6b-a346-691e20ab61ea/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/4707d145-4f52-412b-8609-5d8802b9c276/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/27065018-565a-429d-baa3-97da72a496dc/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/d431fc6a-b6e8-49d9-b663-6300197edf8b/Untitled.png)

### 3.6.2 TCP连接相关的变量

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e124a8f5-b6bb-40a7-bd86-f6b05bb4650b/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/d941885b-05a8-4cb5-bd66-342f2d111fd9/Untitled.png)

### 3.6.3 请求中产生的变量

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/99671b48-2704-457b-9b15-91cf3ae64433/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/6ebc7775-87ec-4286-be6f-00456ea03ffc/Untitled.png)

### 3.6.4 方式HTTP响应时相关的变量

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/d1913a72-331f-4bff-9bce-c21023a4f24a/Untitled.png)

### 3.6.5 系统变量

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/cce900b8-b181-4274-9ed8-c805b53bcd32/Untitled.png)

## 3.7 referer模块

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/8eced5be-d5aa-4e69-b949-ab4c70bfef94/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/37c85a63-c879-4c4c-a3c9-1ca10e447ba8/Untitled.png)

- valid_referers指令

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/dafda2db-3e4d-44d9-966f-bdcb11219b1d/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/562664a3-c625-4aca-9ca7-3c7b6b1944bd/Untitled.png)

结果：403，valid，valid，valid,403,valid,403,valid

## 3.8 map模块

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/43ea7fbb-6468-4bb5-975b-dfaba56af7ae/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/78388ea0-6752-4cf9-9379-7b9dc726d9dd/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/305f4d87-892d-4037-8287-02ac48d84206/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e1456888-2e34-4677-9c16-2ef3fc6251bb/Untitled.png)

## 3.9 split_clients模块

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/dfc087d5-867d-4f4c-87bb-a8329d798113/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/924c4315-31ca-4372-abf8-4e2233f7f527/Untitled.png)

## 3.10 geo模块

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/ea902be1-620e-4d2c-a19a-e61fcf23f329/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/0da74904-3e04-44b7-b56d-e7ad2de823b9/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/70c02f54-cfb2-4152-85a7-6d8ff3ee1fda/Untitled.png)

## 3.11 geoip模块

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/5edb4f8b-4220-43c3-8673-c1f1dc5f9008/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/ef9c35b8-41a8-48ec-9742-8577fd747bfa/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/6bdf1773-7c40-4f5c-85b9-7dc0677dbbda/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/93bfe12b-5ca8-4772-b536-8d278099f214/Untitled.png)

3.12 keepalive

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/098746e8-e44c-4bbd-a0ca-0b2715d4b376/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/cbe22bfb-a391-4bb5-a526-fa3d5b724f36/Untitled.png)

---

## 四、反向代理和负载均衡

## 4.1 基本介绍

### 4.1.1 负载均衡

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/fa687e80-99d2-4613-a0df-b2ec0a780a6c/Untitled.png)

一个服务的扩展方向：

A:水平扩展，加机器；

B:纵向扩展：把业务复杂的服务，拆分为业务小的服务，上层nginx通过location把请求分发到不同的服务中去；

C:Z轴扩展：基于用户的信息进行扩展。例如根据请求ip或者其他信息，把请求分发到特定的服务中去；

### 4.1.2 反向代理

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/2969c4c4-2e8c-4097-b60d-8e788131cf16/Untitled.png)

### 4.1.3 缓存

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/eeae5721-97b8-4daa-af4c-79ad6bdb2f24/Untitled.png)

时间缓存：Nginx从下游服务拿到信息后，一边发给客户端，一边把返回内容缓存在nginx中；

空间缓存：上游服务请求nginx，nginx可以预请求下游服务，把数据缓存到nginx中；

## 4.2 负载均衡策略

### 4.2.1 upstream与server指令

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/afe2a147-b8e4-437d-bde8-1283ab8fffb0/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/7c05cf22-ef58-487b-a1a3-0452d5328510/Untitled.png)

- **加权Round-Robin负载均衡算法：**

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/bbe31a3a-9838-49df-bf49-01691875f353/Untitled.png)

- **对下游服务使用keepalive长链接：**

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/22897f1a-9f92-47bb-92d9-b7f98eb10d50/Untitled.png)

- upstream_keepalive指令：

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e31aa469-9f98-4aee-ba70-5f73a81e15a9/Untitled.png)

keepalive:nginx和下游服务最多保持多少个空闲的http连接

keepalive_request:一个tcp最多跑多少个请求

keepalive_timeout:一个tcp连接空闲多少秒后关闭

- resolver指令

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/41435dcc-2986-47a1-932e-267ddc1970fa/Untitled.png)

官方解释下就是：反向代理的场景，upstream后端用域名时，配置resolver以便于nginx能够解析该域名。

当proxy_pass 后面接变量时，而且设置了resolver,会把变量的负载值通过resolver来解析，其他情况通过本地dns服务，etc或host 来解释域名。

[https://www.jianshu.com/p/5caa48664da5](https://www.jianshu.com/p/5caa48664da5)

## 4.3 负载均衡哈希算法

### 4.3.1 upstream_ip_hash模块

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/ac61f81e-5a41-441a-97ce-5554b76cb2d8/Untitled.png)

### 4.3.2 upstream_hash模块

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e0969e4f-239d-4fa1-80eb-a81b3f43aae7/Untitled.png)

### 4.3.3 演示

```bash
upstream iphashtest { 
	ip_hash;
	hash user_$arg_username;#使用username作为hash算法的关键字
	server 127.0.0.1:8011;
	server 127.0.0.1:8012;
}
```

## 4.4 一致性hash算法

使用hash算法，下游服务器异常会导致大量请求的路由策略失效。一致性hash算法能有效解决该问题。

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/722532b2-d429-4438-af91-4c02efb4fbd7/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/4d2f87c6-e954-4bb5-83cc-3b092bd7d84f/Untitled.png)

把0-2^32围成一个环，4个服务均有的分布在环上，0-2^30请求第一个服务，以此类推；

扩容后：

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/cae04007-32f4-4c3b-9828-898a4c1e3371/Untitled.png)

扩容后只会改变node2—node4前半段的hash点。

### 4.4.1 使用方法

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/e4553567-2430-47db-a108-6861a762e877/Untitled.png)

## 4.5 最小连接数算法

### 4.5.1 upstream_least_conn模块

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/6fb045b6-4086-46b2-85dc-e4c3c73ff344/Untitled.png)

## 4.6 怎么跨worker生效：upstream_zone模块

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/f83da43c-39a4-4a3a-b983-5b081913da0b/Untitled.png)

如上所有算法均可以使用upstream_zone来使得所有worker生效。

## 4.7 upstream模块间的顺序

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/dbdb06c7-80cc-47d0-97a0-cadfa77d44ea/Untitled.png)

## 4.8 http upstream提供的变量

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/22c614dd-1035-4665-abb3-9daeba1d61a7/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/9417c091-11e9-42f9-bc41-e51d841457dc/Untitled.png)

## 4.9 反向代理

### 4.9.1 http反向代理proxy处理请求的流程

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/14f72d61-55cf-4c6d-8191-adc83daaa96f/Untitled.png)

### 4.9.2 proxy模块

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/0934dc17-3f5c-4256-8738-c4b7c4c0822c/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/40eb18d2-dc6e-41a1-b8d2-abf4b8b743c3/Untitled.png)

- 案例

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/98222279-57b0-4fb6-b71c-5940aec8fc59/Untitled.png)

http://proxyups/addurl会被换为：http://proxyups/a

### 4.9.3 修改请求下游服务的的请求

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/ec72a8c9-9d94-4bac-8895-a1d30f0e5d5b/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/3f06c69f-23a5-422e-b5af-fcc1a52ee24d/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/ec8b60a4-eb22-47ad-8d5f-c9842b80cc3f/Untitled.png)

## 4.10 接收用户请求body

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/8e9d90cc-e0db-4492-b783-c6d64cf8819b/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/a00e4e38-a9be-49bd-9759-3e6787ce96fb/Untitled.png)

client_body_buffer_size：Nginx分配给请求数据的Buffer大小，如果请求的数据小于client_body_buffer_size直接将数据先在内存中存储。如果请求的值大于client_body_buffer_size小于client_max_body_size，就会将数据先放到client_body_buffer_size的内存中，再一遍一遍地写到临时文件中。

client_body_in_single_buffer:客户端请求数据的body一律存储到内存buffer中。当然，如果HTTP包体的大小超过了下面client_body_buffer_size设置的值，包体还是会写入到磁盘文件中。

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/12a96f98-9c9d-4f78-88b2-d1ce56b28b54/Untitled.png)

client_max_body_size 默认 1M，表示 客户端请求服务器最大允许大小，在“Content-Length”请求头中指定。如果请求的正文数据大于client_max_body_size，HTTP协议会报错 413 Request Entity Too Large。就是说如果请求的正文大于client_max_body_size，一定是失败的。如果需要上传大文件，一定要修改该值。

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/4f11b2c7-3138-44dd-9aee-7476a8003bf8/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/2de20323-47d8-40e7-90e5-3759ff544759/Untitled.png)

2次读取body时间超过60s后，返回408.

## 4.11 连接下游服务器

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/43f90e53-ac82-40ed-b799-14c7e6a23d34/Untitled.png)

当与下游服务建立连接超时时，再换一台服务器重新连接。

### 4.11.1 TCP keepalive

关闭无用的连接，减少资源浪费。

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/463228bf-f151-48d7-9afe-43b68d80ba19/Untitled.png)

使用操作系统设置的默认keepalive相关配置来控制tcp的keepalive来降低资源使用。

### 4.11.2 HTTP keepalive

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/887e9c16-c05c-4fad-82a8-8168df81600e/Untitled.png)

### 4.11.3 修改TCP连接中的local address

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/5d98cbf3-8873-4c5f-be5a-a1e87d59a056/Untitled.png)

proxy_bind隶属于proxy_module，为向后端建立连接时的local ip，在nginx源码中只支持bind一个ip进行回源，若想使用多个ip进行回源时，可以修改源码支持bind ip数组。在实际应用中我就是这样做的。bind ip数据轮询选择ip进行回源与upstream建立连接，以解决单ip回源连接数限制问题。

proxy_bind：它的用法主要有两类用途，第一类用途就是当我们nginx上有多个ip地址时，可能有多个路由的策略是不同的，比如内网或者外网等，这个时候不要使用系统默认给我们选择的ip地址，而是主动使用一个ip地址。这个时候用proxy_bind。第二种场景，很可能为了传递一个ip地址，就是透传ip地址的策略，比如在stream反向代理中会很常用，在之后还会详述。这里先说下proxy_bind用法。

proxy_bind $remote_addr 也就是客户端的地址绑定到这里。绑定变量的时候呢，如果地址不是本地的地址，linux必须要加transparent。非linux操作系统呢需要保证worker进程有root 权限的，才能人为的修改socket的local address。

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/b25f9cae-ca86-4f18-99de-e20695d41e2e/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/3d3dc40e-181a-4e07-9ce0-ca0f1a264ca7/Untitled.png)

## 4.12 接收下游的响应

### 4.12.1 接收下游服务的响应头部

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/3bfa06bc-157b-4a5b-8fd4-02542586ecfd/Untitled.png)

proxy_buffer_size：限定了接收自上游的http response中header的最大值。所以当上游的server发送了http响应，如果有set cookie这种特别长的header可能就会导致整个全部的response header超出了这个值。超出完之后这个请求就不能够被nginx正确的处理了。我们的error.log中会看到`upstream sent too big header`
就是这样的一个原因。

### 4.12.2 接收下游服务的响应body

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/b64c36d7-6ee3-4b40-9092-1a5221b486b3/Untitled.png)

proxy_buffering：来控制我们是不是先接收完整的包体，接收完了才开始转发。或者说不接收完，而是每接收一部分就同步的向客户端发送我收到的那部分响应。这两种方式各有各的好处。通常情况下默认开启（on）。因为我们认为上游服务和我们nginx走的是内网，网速更快。如果我们边发边接收上游边往客户端发。因为客户端跟nginx之间的网速可能很慢。所以就会导致，对于比较大的body 的时候，nginx长时间与上游建立连接。而上游比如说tomcat、Django等它们的并发能力是很弱的。当然如果我们用了proxy_buffering off 它的优点是能让客户端及时接收到响应，特别是一些大包体的情况下。客户端不用再等待。

在我们接收上游发来的HTTP包体，即使我们开启了proxy_buffering on也并不一定向磁盘中写入包体，因为如果包体非常的小，在内存中就可以放入的话，就没有必要写到磁盘中，因为磁盘io总是比较慢的。所以这个时候就有了proxy_buffers指令。也就是包体大小没有超过这个设定值就不用写入磁盘。否则的话就要写入磁盘了。

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/b53ca84e-93ea-4a96-b511-7f81e996a984/Untitled.png)

proxy_buffering ，默认开启，希望尽快释放上游服务器的连接，当然proxy_buffering 还有一个nginx特定的header。这个header（X-Accel-Buffering头部）只有nginx才会认。当上游的tomcat 如果在response中加入X-Accel-Buffering头部，如果配置为yes，就会强制要求nginx先接收完上游的http body 再向client发送。也就是它会替换指定的内容。

当我们向磁盘中写入包体的时候还有三个指令，proxy_max_temp_file_size、proxy_temp_file_write_size、proxy_temp_path。

proxy_max_temp_file_size：限制写入磁盘中这个文件的最大值。如果上游服务中返回了非常大的文件超出了临时文件大小也会出错的。默认是1G。

proxy_temp_file_write_size：每一次向磁盘文件中写入的字节数。

proxy_temp_path：设定了存放临时文件的目录在哪里，以及目录level层级。

### 4.12.3 及时转发body

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/1ee04064-ef13-4991-90aa-a310e59fa82a/Untitled.png)

proxy_busy_buffers_size。虽然被缓存所有的响应，我们希望更及时的向客户端发送部分响应。比如我们收到1G文件。当我们接收到前8k或前16k（proxy_busy_buffers_size 8k|16k）的时候，就先向客户端转发接收到的这一部分响应的话就可以使用proxy_busy_buffers_size 。

### 4.12.4 接收下游服务时网络速度相关指令

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/926c7d10-3502-46a0-ba7c-a471c7b44e5e/Untitled.png)

proxy_read_timeout：两次读取操作之间最多60秒的超时。两次读取是一个TCP层的一个概念。

proxy_limit_reate：限速，和客户端limit_rate 有些类似，但是它限制的是读取上游的响应，而不是发送给上游服务的网速。设置为0表示不限制读取上游响应的速度。

### 4.12.5 上游body的持久化

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/5d7f950f-2c83-4b64-af38-0ed01c31304d/Untitled.png)

proxy_store_access：配置指定目录权限。

proxy_store：把临时文件改名到root对应的目录下，默认不开启，如果是`string`，可再次指定，使用变量的方式，指定这个文件存放的位置。

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/d60596a3-df8c-4532-b343-7d5b113195cf/Untitled.png)

![Untitled](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/4a2bb6c4-0616-4730-ac72-787155006ea2/Untitled.png)

## 五、Nginx的系统层性能优化

## 六、从源码视角深度使用Nginx和OpenResty