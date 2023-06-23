API 设计
在开始构建我们的 API 项目之前，需要了解如何设计好一套 API。

这里我们谈论的 API，是作为手机 APP 、小程序等的数据后台服务，其他场景的 API 设计也可作为参考。

Github API 是设计优良的一套 RESTful API，业内知名度很高，以下设计方案部分对其做参考和引用。在此基础上，结合国内 API 项目实战中常见问题，总结如下。

一套优秀的 API 设计，需要具备如下特性：

使用 HTTPS
使用域名
版本区分
使用 URL 来定位资源
使用 HTTP 动词来描述操作
支持资源过滤
使用 HTTP 状态码
数据响应的一致性
支持限流
API 文档
自带分页链接
强制 User-Agent

1. 使用 HTTPS
   HTTPS 为接口的安全提供了保障，可以有效防止通信被窃听和篡改。而且现在部署 HTTPS 的成本也越来越低，你可以通过 certbot 等工具，方便快速的制作免费的安全证书，所以生产环境，请务必使用 HTTPS。

注意：非 HTTPS 的 API 调用，不要重定向到 HTTPS，而要直接返回调用错误以禁止不安全的调用。

明文传输是一个很严重的中间人攻击安全漏洞。假如咖啡店的路由器被攻占，未使用 HTTPS 进行传输的话，黑客即可轻易获取到用户请求的 Access Token，利用此 Token 对用户的密码和邮箱等做修改，用户的账号就会被盗取。

2. 使用域名
   应当尽可能的将 API 与其主域名区分开，可以使用专用的域名，访问我们的 API，例如：

https://api.gohub.com
或者可以放在主域名下，例如：

https://www.gohub.com/api
区分域名的好处多多，以下举几个例子。

无 Cookie 共享 —— API 域名不使用主域名的会话机制，免去了无意义的 Cookie 传输，一定程度上提高了响应速度；
CDN 加速 —— 高流量时，我们可以对 GET 的 API 请求做 CDN 加速，单独的域名很容易做映射；
专属的域名也方便我们做 gzip 以及过期标头的设置。
总之，区分域名为 API 打好了灵活的基础，能够适应更复杂多变的业务需求。

3. 版本控制
   随着业务的发展，需求的不断变化，API 的迭代是必然的，很可能当前版本正在使用，而我们就得开发甚至上线一个不兼容的新版本，为了让旧用户可以正常使用，为了保证开发的顺利进行，我们需要控制好 API 的版本。

通常情况下，有两种做法：

将版本号直接加入 URL 中
https://api.gohub.com/v1
https://api.gohub.com/v2
使用 HTTP 请求头的 Accept 字段进行区分
Accept: application/prs.gohub.v1+json
Accept: application/prs.gohub.v2+json
Github Api 虽然默认使用了第一种方法，但是其实是推荐并实现了第二种方法的，我们同样也尽量使用第二种方式。

4. 用 URL 定位资源
   在 RESTful 的架构中，所有的一切都表示资源，每一个 URL 都代表着一种资源，资源应当是一个名词，而且大部分情况下是名词的复数，尽量不要在 URL 中出现动词。

先来看看 github 的 例子：

GET /issues                                      列出所有的 issue
GET /orgs/:org/issues                            列出某个项目的 issue
GET /repos/:owner/:repo/issues/:number           获取某个项目的某个 issue
POST /repos/:owner/:repo/issues                  为某个项目创建 issue
PATCH /repos/:owner/:repo/issues/:number         修改某个 issue
PUT /repos/:owner/:repo/issues/:number/lock      锁住某个 issue
DELETE /repos/:owner/:repo/issues/:number/lock   解锁某个 issue
例子中冒号开始的代表变量，例如 /repos/summerblue/gohub/issues

在 github 的实现中，我们可以总结出：

资源的设计可以嵌套，表明资源与资源之间的关系。
大部分情况下我们访问的是某个资源集合，想得到单个资源可以通过资源的 id 或 number 等唯一标识获取。
某些情况下，资源会是单数形式，例如某个项目某个 issue 的锁，每个 issue 只会有一把锁，所以它是单数。
错误的例子

POST https://api.gohub.com/createTopic
GET https://api.gohub.com/topic/show/1
POST https://api.gohub.com/topics/1/comments/create
POST https://api.gohub.com/topics/1/comments/100/delete
正确的例子

POST https://api.gohub.com/topics
GET https://api.gohub.com/topics/1
POST https://api.gohub.com/topics/1/comments
DELETE https://api.gohub.com/topics/1/comments/100
养成一个习惯，就是阅读 URL 时连带 HTTP 方法，把 HTTP 方法当成动词：

GET —— 获取
POST —— 创建
PUT/PATCH —— 更新
DELETE —— 删除
如：

POST https://api.gohub.com/topics
阅读起来应该是 —— 创建一个新的话题。而：

PUT https://api.gohub.com/topics
应该阅读 —— 更新一个话题。

关于更新操作 PUT/PATCH 的区别，下面会讲解。

5. 用 HTTP 动词描述操作
   HTTP 设计了很多动词，来表示不同的操作，RESTful 很好的利用的这一点，我们需要正确的使用 HTTP 动词，来表明我们要如何操作资源。

先来解释一个概念，幂等性，指一次和多次请求某一个资源应该具有同样的副作用，也就是一次访问与多次访问，对这个资源带来的变化是相同的。

常用的动词及幂等性

动词    描述    是否幂等
GET    获取资源，单个或多个    是
POST    创建资源    否
PUT    更新资源，客户端提供完整的资源数据    是
PATCH    更新资源，客户端提供部分的资源数据    否
DELETE    删除资源    是
为什么 PUT 是幂等的而 PATCH 是非幂等的，因为 PUT 是根据客户端提供了完整的资源数据，客户端提交什么就替换什么，而 PATCH 有可能是根据客户端提供的参数，动态的计算出某个值，例如每次请求后资源的某个参数减 1，所以多次调用，资源会有不同的变化。

另外需要注意的是，GET 请求对于资源来说是安全的，不允许通过 GET 请求改变（更新或创建）资源。但是真实使用中，为了方便统计类的数据，会有一些例外情况，例如帖子详情，记录访问次数，每调用一次，访问次数 +1;

6. 资源过滤
   我们需要提供合理的参数供客户端过滤资源，例如

?page=2&per_page=100：访问第几页数据，每页多少条。
?sort=name&order=asc：指定返回结果按照哪个属性排序，以及排序顺序。
API 端也应当对过滤的参数做验证。例如 per_page 只允许 10~100 的区间，或者 order 的值只能是 asc 或 desc。

7. 正确使用状态码
   HTTP 提供了丰富的状态码供我们使用，正确的使用状态码可以让响应数据更具可读性。

200 OK - 对成功的 GET、PUT、PATCH 或 DELETE 操作进行响应。也可以被用在不创建新资源的 POST 操作上
201 Created - 对创建新资源的 POST 操作进行响应。应该带着指向新资源地址的 Location 头
202 Accepted - 服务器接受了请求，但是还未处理，响应中应该包含相应的指示信息，告诉客户端该去哪里查询关于本次请求的信息
204 No Content - 对不会返回响应体的成功请求进行响应（比如 DELETE 请求）
304 Not Modified - HTTP 缓存 header 生效的时候用
400 Bad Request - 请求异常，比如请求中的 body 无法解析
401 Unauthorized - 没有进行认证或者认证非法
403 Forbidden - 服务器已经理解请求，但是拒绝执行它
404 Not Found - 请求一个不存在的资源
405 Method Not Allowed - 所请求的 HTTP 方法不允许当前认证用户访问
410 Gone - 表示当前请求的资源不再可用。当调用老版本 API 的时候很有用
415 Unsupported Media Type - 如果请求中的内容类型是错误的
422 Unprocessable Entity - 用来表示校验错误
429 Too Many Requests - 由于请求频次达到上限而被拒绝访问
有一些 API 的设计，不论接口的状态成功与否，都会返回 200 ，然后使用自定的状态码，例如说 30404 （数据不存在），这种方法是不可取的。

HTTP 状态码是行业标准，意味着成千上万开发者都在认同和使用这套规则，意味着他们写出来的 HTTP 通讯程序（类库）也在使用这套规则。所以没有必要，也不应该重新发明自己的一套规则。

8. 数据响应格式
   考虑到响应数据的可读性及通用性，我们将使用 JSON 作为主要的数据响应格式。程序也应当在框架上支持不同的响应格式，做好解析请求和构建响应的封装，这样后续如果有其他格式的需求，也可以很快加上。

一般通过解析 Accept 标头来辨认需要的格式：

https://api.gohub.com/
    Accept: application/prs.gohub.v1+json
    Accept: application/prs.gohub.v1+xml
对于错误数据，默认使用如下结构：

'message' => ':message',         // 错误描述，一般情况下，客户端可用于直接展示给用户
'errors' => ':errors',            // 参数的具体错误描述
'error_code' => ':error_code',    // 自定义的错误码，方便开发人员定位问题
例如

{
    "message": "请求验证失败，请检查",
    "errors": {
        "name": [
            "姓名 必须为字符串。"
        ]
    },
    "error_code": 10201
}
9. 支持限流
为了防止服务器被攻击，减少服务器压力，需要对接口进行合适的限流控制。

限流是非常有必要的。现实的生产环境中，没有限流机制的接口，很容易让服务器陷入不可用的状态。一些接口也会被黑客利用，盗取数据，或者做安全渗透测试。虽然无法完全避免此类情况出现，但是至少你多了一个门槛。

限流时，需要在响应头信息中加入合适的信息，告知客户端当前的限流情况：

X-RateLimit-Limit :10000 最大访问次数
X-RateLimit-Remaining :9993 剩余的访问次数
X-RateLimit-Reset :1513784506 到该时间点，访问次数会重置为 X-RateLimit-Limit
以上三个参数是一般接口常见的限流参数，我们会遵循此命名方式。

超过限流次数后，需要返回 429 Too Many Requests 状态码。

10. API 文档
    为了方便用户使用，我们需要提供清晰的文档，尽可能包括以下几点

包括每个接口的请求参数，每个参数的类型限制，是否必填，可选的值等。
响应结果的例子说明，包括响应结果中，每个参数的释义。
对于某一类接口，需要有尽量详细的文字说明，比如针对一些特定场景，接口应该如何调用。
错误码要做详细标记。
11. 自带分页链接
如下的 pager 字段：

{
    "data": [
        {
            "id": 1,
            .
            .
            .
            "updated_at": "2021-12-26T22:54:26.38+08:00"
        },
        {
            "id": 2,
            .
            .
            .
            "updated_at": "2021-12-26T22:54:26.38+08:00"
        }
    ],
    "pager": {
        "HasPages": true,
        "CurrentPage": 2,
        "PerPage": 2,
        "TotalPage": 54,
        "TotalCount": 108,
        "NextPageURL": "https://api.gohub.com/v1/topics?page=3&sort=id&order=asc&per_page=2",
        "PrevPageURL": "https://api.gohub.comv1/topics?page=1&sort=id&order=asc&per_page=2"
    }
}
这种做法把分页逻辑控制在服务端，既免去了客户端的 URL 拼接，方便调用，另一方面增加了灵活度。假如后续产品需求中，我们需要新增默认过滤规则，客户端的调用不需要修改。

现实生产环境中，做 APP、小程序接口开发时，要时刻记住，调用接口的客户端会存在大量的 老客户端 ，这些是没来得及更新，或者关闭了自动更新的用户。作为接口的设计者，需要考虑他们的兼容性。

12. 强制 User-Agent
    强制客户端在请求时，发送 User-Agent 信息。

User-Agent 信息包含两部分，客户端信息 + 版本，使用斜杆分隔：

User-Agent: Gohub iOS/2.1.37
User-Agent: Gohub Android/2.1.22
API 后端接收到 User-Agent 数据后可以暂时不做处理，但是后续有特殊的业务需求时，可以针对某个客户端具体到版本，进行特殊的数据处理。

常见的使用场景，是废弃客户端：例如一个银行 APP，升级了交易时的加密算法，低于 5.0 版本的客户端因为安全原因，必须废弃。针对此情况，可通过后端 API 判断 User-Agent 标头，对低于 5.0 的版本的客户端请求，返回专属的数据，如 APP 首页的第一个 Banner 显示请升级客户端，安全升级无法使用的提示。

现实生产中，有些客户端用户会关闭系统的应用自动更新功能，多版本客户端是无法避免的问题。有了 User-Agent ，我们可以更加灵活的做针对性处理。




