package main

import "project/di"

func main() {

	//new时候将对应dao对象传递进去, 执行对应方法 -- 依赖注入
	s, _ := di.NewService()
	s.ProcessBusiness()  // hello dao


	s, _ = di.NewTestService()
	s.ProcessBusiness() // hello test dao


	//service.New(dao.New()).ProcessBusiness()
	//service.New(dao.NewTestDao()).ProcessBusiness()
}
