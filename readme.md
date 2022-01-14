一个单元测试的demo
#### 方式一：
生成mock文件 
mockgen -source=model/user.go -destination=controller/mocks/user.go -package=mocks Iuser 
- source 指定源文件，该文件必须包含接口，mock是通过实现接口方法的方式来使程序调用mock出的方法，而不是继续调用程序中原来的方法
- destination 生成文件的路径
- package 生成mock文件的包名 不指定的话默认原包名+前缀 mocks_
- Iuser 指定哪个接口生成mock文件

在生成mock文件的基础上进行接口mock

#### 方式二：
gomonkey 
可以模拟方法+函数
详见user_test_gomonkey.go 文件

生成测试覆盖率文件 
go test -covermode=count -coverprofile=testcover.cov -coverpkg=./... -run TestGetUser  ./controller
- covermode 覆盖率统计方式
- coverprofile 生成文件路径
- coverpkg 以哪些文件为基础统计覆盖率 ./... 当前路径下所有文件
- run 指定测试用例 省略时指所有的测试用例
- 最后一个参数 测试用例文件路径  

把测试覆盖率文件转换成网页格式
go tool cover -html testcover.cov -o cover.html


#### 配合搭建的sonarqube统计项目ut覆盖率
 - 搭建好本地sonarqube环境 参考[链接](https://blog.csdn.net/mario08/article/details/115112060)
 - 分别执行以下两行命令
 - go test -json -covermode=atomic -coverpkg=./... -coverprofile coverage.out ./... > report.json
 - sonar-scanner -Dsonar.projectKey=user_demo -Dsonar.inclusions=**/*.go -Dsonar.go.coverage.reportPaths=coverage.out -Dsonar.go.tests.reportPaths=report.json -Dsonar.exclusions=**/*_test.go
