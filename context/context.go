package context

import (
	"errors"
	"time"
)

//包上下文定义了上下文类型，该类型在API边界之间以及进程之间传递截止日期，取消信号和其他请求范围的值。

//向服务器的传入请求应创建一个上下文，而对服务器的传出调用应接受一个上下文。它们之间的函数调用链必须传播Context，可以选择将其替换为使用WithCancel，WithDeadline，WithTimeout或WithValue创建的派生Context。取消上下文后，从该上下文派生的所有上下文也会被取消。

//WithCancel，WithDeadline和WithTimeout函数采用Context（父级）并返回派生的Context（子级）和CancelFunc。调用CancelFunc会取消该子代及其子代，删除父代对该子代的引用，并停止所有关联的计时器。未能调用CancelFunc会使子代及其子代泄漏，直到父代被取消或计时器触发。审核工具检查所有控制流路径上是否都使用了CancelFuncs。

//不要将上下文存储在结构类型中；而是将上下文明确传递给需要它的每个函数。 Context应该是第一个参数，通常命名为ctx：
//func DoSomething(ctx context.Context, arg Arg) error {
// ... use ctx ...
//}

//即使函数允许，也不要传递nil Context。如果不确定使用哪个上下文，请传递context.TODO。

//仅将上下文值用于传递过程和API的请求范围的数据，而不用于将可选参数传递给函数。

//可以将相同的上下文传递给在不同goroutine中运行的函数。上下文对于由多个goroutine同时使用是安全的。

var (
	//取消是Context.Err取消上下文时返回的错误。
	Canceled = errors.New("context canceled")
)

//DeadlineExceeded是上下文的最后期限过去时Context.Err返回的错误。
var DeadlineExceeded error = deadlineExceededError{} 

//WithCancel返回具有新的“完成”通道的父级副本。 当调用返回的cancel函数或关闭父上下文的Done通道时（以先发生者为准），关闭返回的上下文的Done通道。

//取消此上下文将释放与其关联的资源，因此在此上下文中运行的操作完成后，代码应立即调用cancel。
//func WithCancel(parent Context) (ctx Context, cancel CancelFunc) GO1.7

//WithDeadline返回父上下文的副本，并将截止日期调整为不迟于d。 如果父母的截止日期早于d，则WithDeadline（parent，d）在语义上等同于父母。 当截止日期到期，调用返回的cancel函数或关闭父上下文的Done通道（以先到者为准）时，将关闭返回的上下文的Done通道。

//取消此上下文将释放与其关联的资源，因此在此上下文中运行的操作完成后，代码应立即调用cancel。
//func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) GO1.7

//WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
//取消此上下文将释放与之关联的资源，因此在此上下文中运行的操作完成后，代码应立即调用cancel：
// func slowOperationWithTimeout(ctx context.Context) (Result, error) {
// 	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
// 	defer cancel()  // releases resources if slowOperation completes before timeout elapses
// 	return slowOperation(ctx)
// }
//func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)

//CancelFunc告诉操作放弃其工作。 CancelFunc不等待工作停止。多个goroutine可以同时调用CancelFunc。在第一个调用之后，随后对CancelFunc的调用什么也不做。
type CancelFunc func()

//上下文在API边界上带有截止日期，取消信号和其他值。 上下文的方法可以同时由多个goroutine调用。

type Context interface {
	//截止日期返回应取消代表该上下文完成的工作的时间。 如果未设置截止日期，则截止日期返回ok == false。 连续调用Deadline会返回相同的结果。
	Deadline() (deadline time.Time, ok bool)

	Done() <-chan struct {}
}
