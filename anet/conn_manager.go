// @Title
// @Description
// @Author  Wangwengang  2021/8/17 下午5:17
// @Update  Wangwengang  2021/8/17 下午5:17
package anet

/*
	连接管理抽象层
*/
type ConnManager interface {
	Add(conn Connection)                   //添加链接
	Remove(conn Connection)                //删除连接
	Get(connID uint64) (Connection, error) //利用ConnID获取链接
	Len() int                              //获取当前连接
	ClearConn()                            //删除并停止所有链接
}
