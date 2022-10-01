# Layzer-server

#### This is a simple API test program that contains two request methods for simply testing whether the service is available,



#### So, let's see how this program is used.



#### The first thing we need to know is that this program is completely developed in Go language



#### Secondly, we need to know that the purpose of this program is to test whether resources such as Kubernetes (Service, Gateway, Ingress...), Consul, API Gateway and other resources are available.



#### So, let's take a look at the main functions of this tool.



#### First：This program has its own API and JWT functions



#### API usage is as follows



```apl
GET /

curl http://10.0.0.1 -H "authorization: 123456"
```

```apl
POST /

curl -X POST -H "authorization: 123456" -d '{"name": "layzer"}' http://10.0.0.1
```

### warn: Whether it is GET / POST, you need to pass "authorization" as the authentication parameter



### help：We can type "-h" for help

