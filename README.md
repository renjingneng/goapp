# 功能说明
+ 一个简易的mvc框架
+ 中间件路由用的gin
+ 利用反射把网址映射到controller的方法
+ 核心代码放在了core包里面，core包里面引用第三方库代码，业务代码只需用core包，从而隔离业务代码与第三方包
+ 支持多业务线，每个业务线多个应用，分本地、测试线、正式线三个环境
# 框架结构
```text
goapp                           
├─ core    //核心包                 
├─ business_1    //业务1文件夹                 
│  ├─ config  
│  │  ├─ config-dev.yaml                      
│  │  ├─ config-local.yaml                    
│  │  └─ config-prod.yaml
│  ├─ entrance  
│  │  ├─ app_1.go     //业务1下的应用1                        
│  │  ├─ app_2.go     //业务1下的应用2               
│  │  └─ app_3.go     //业务1下的应用3 
│  └─ database
│  └─ controller
│  └─ service                                 
│  └─ template   
│  └─ app_1.exe  // 业务1下的应用1编译后可执行文件，该文件不入代码库git，本地测试可以执行这个                          
│  └─ app_2.exe  // 业务1下的应用2编译后可执行文件，该文件不入代码库git，本地测试可以执行这个                          
│  └─ app_3.exe  // 业务1下的应用3编译后可执行文件，该文件不入代码库git，本地测试可以执行这个                          
│
├─ business_2    //业务2文件夹                 
│  ├─ config  
│  │  ├─ config-dev.yaml                      
│  │  ├─ config-local.yaml                    
│  │  └─ config-prod.yaml
│  ├─ entrance  
│  │  ├─ app_1.go     //业务2下的应用1                        
│  │  ├─ app_2.go     //业务2下的应用2               
│  │  └─ app_3.go     //业务2下的应用3 
│  └─ database
│  └─ controller
│  └─ service                                 
│  └─ template   
│  └─ app_1.exe  // 业务2下的应用1编译后可执行文件，该文件不入代码库git，本地测试可以执行这个                          
│  └─ app_2.exe  // 业务2下的应用2编译后可执行文件，该文件不入代码库git，本地测试可以执行这个                          
│  └─ app_3.exe  // 业务2下的应用3编译后可执行文件，该文件不入代码库git，本地测试可以执行这个  
├─ .gitignore    
├─ .gitlab-ci.yml                      
├─ go.mod                   
├─ go.sum                                   
└─ README.md
```   