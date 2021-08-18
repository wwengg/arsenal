// @Title  
// @Description  
// @Author  Wangwengang  2021/8/17 下午5:13
// @Update  Wangwengang  2021/8/17 下午5:13
package anet

/*
	消息管理抽象层
*/
type MsgHandle interface {
	DoMsgHandler(request Request)          //马上以非阻塞方式处理消息
	AddRouter(msgID uint32, router Router) //为消息添加具体的处理逻辑
	StartWorkerPool()                       //启动worker工作池
	SendMsgToTaskQueue(request Request)    //将消息交给TaskQueue,由worker进行处理
}
