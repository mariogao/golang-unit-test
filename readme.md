一个单元测试的demo

生成mock文件 
mockgen -source=model/user.go -destination=controller/mocks/user.go -package=mocks Iuser 
- source 指定源文件，该文件必须包含接口，mock是通过实现接口方法的方式来使程序调用mock出的方法，而不是继续调用程序中原来的方法
- destination 生成文件的路径
- package 生成mock文件的包名 不指定的话默认原包名+前缀 mocks_
- Iuser 指定哪个接口生成mock文件

生成测试覆盖率文件 
go test -covermode=count -coverprofile=testcover.cov -coverpkg=./... -run TestGetUser  ./controller
- covermode 覆盖率统计方式
- coverprofile 生成文件路径
- coverpkg 以哪些文件为基础统计覆盖率 ./... 当前路径下所有文件
- run 指定测试用例 省略时指所有的测试用例
- 最后一个参数 测试用例文件路径  

把测试覆盖率文件转换成网页格式
go tool cover -html testcover.cov -o cover.html
