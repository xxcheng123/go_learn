package t_runtime

import (
	"sync/atomic"
	"unsafe"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/8 14:53
 */

// 已定义的常量
const (
	// G 状态
	//
	// 除了表示协程的一般状态外，
	// G 状态也充当了对协程栈（因此对执行用户代码的能力）的锁。
	//
	// 如果你添加到这个列表中，请也在 mgcmark.go 中的
	// "垃圾收集期间允许的" 状态列表中添加。
	//
	//
	// TODO（austin）：_Gscan 位可能更轻量。
	// 例如，我们可以选择不运行在运行队列中找到的 _Gscanrunnable
	// 的 goroutines，而不是通过 CAS 循环直到它们变为 _Grunnable。
	// 并且像 _Gscanwaiting -> _Gscanrunnable 这样的转换实际上是可以的，因为
	// 它们不影响栈所有权。

	// _Gidle 表示这个协程刚刚分配并且尚未初始化。
	_Gidle = iota // 0

	// _Grunnable 表示这个协程在运行队列中。它当前没有执行用户代码。栈不属于它。
	_Grunnable // 1

	// _Grunning 表示这个协程可能执行用户代码。栈属于这个协程。它不在运行队列上。
	// 它分配了一个 M 和一个 P（g.m 和 g.m.p 是有效的）。
	_Grunning // 2

	// _Gsyscall 表示这个协程正在执行系统调用。
	// 它没有执行用户代码。栈属于这个协程。它不在运行队列上。
	// 它分配了一个 M。
	_Gsyscall // 3

	// _Gwaiting 表示这个协程在运行时中被阻塞。
	// 它没有执行用户代码。它不在运行队列上，但应该被记录在某处（例如，通道等待队列），
	// 以便在必要时可以被 ready()。
	// 栈不属于它，除非在适当的通道锁定下，通道操作可能读取或写入栈的某些部分。
	// 否则，在协程进入 _Gwaiting 后，访问栈是不安全的（例如，可能会被移动）。
	_Gwaiting // 4

	// _Gmoribund_unused 目前未使用，但在 gdb 脚本中硬编码。
	_Gmoribund_unused // 5

	// _Gdead 表示这个协程目前未使用。它可能刚刚退出，位于空闲列表上，或者刚刚初始化。
	// 它没有执行用户代码。它可能有也可能没有分配栈。
	// G 和其栈（如果有的话）由退出 G 的 M 或从空闲列表中获取 G 的 M 所拥有。
	_Gdead // 6

	// _Genqueue_unused 目前未使用。
	_Genqueue_unused // 7

	// _Gcopystack 表示这个协程的栈正在被移动。
	// 它没有执行用户代码，也不在运行队列上。栈由将其放在 _Gcopystack 中的协程所拥有。
	_Gcopystack // 8

	// _Gpreempted 表示这个协程停止了自己，以进行 suspendG 抢占。
	// 它类似于 _Gwaiting，但是还没有什么负责 ready() 这个 G。
	// 一些 suspendG 必须使用 CAS 操作将状态更改为 _Gwaiting，以负责 ready() 这个 G。
	_Gpreempted // 9

	// _Gscan 与上述除 _Grunning 外的任何状态结合在一起表示 GC 正在扫描栈。
	// 协程没有执行用户代码，栈由设置 _Gscan 位的协程所拥有。
	//
	// _Gscanrunning 不同：它用于在 GC 向 G 信号扫描其自己的栈时暂时阻塞状态转换。
	// 除此之外，它与 _Grunning 一样。
	// atomicstatus&~Gscan 给出了协程在扫描完成时将返回的状态。
	_Gscan          = 0x1000
	_Gscanrunnable  = _Gscan + _Grunnable  // 0x1001
	_Gscanrunning   = _Gscan + _Grunning   // 0x1002
	_Gscansyscall   = _Gscan + _Gsyscall   // 0x1003
	_Gscanwaiting   = _Gscan + _Gwaiting   // 0x1004
	_Gscanpreempted = _Gscan + _Gpreempted // 0x1009
)

const (
	// P 状态

	// _Pidle 表示一个 P 没有用于运行用户代码或调度器。
	// 通常，它在空闲 P 列表上并可供调度器使用，
	// 但它可能正在其他状态之间过渡。
	//
	// P 属于空闲列表或正在转换其状态的任何实体。它的运行队列为空。
	_Pidle = iota

	// _Prunning 表示一个 P 属于一个 M 并正在用于运行用户代码或调度器。
	// 只有拥有此 P 的 M 允许从 _Prunning 更改 P 的状态。
	// M 可以将 P 转换为 _Pidle（如果没有更多工作要做）、
	// _Psyscall（进入系统调用时）或 _Pgcstop（为了停止 GC）。
	// M 还可以直接将 P 的所有权移交给另一个 M
	// （例如，为了调度一个被锁定的 G）。
	_Prunning

	// _Psyscall 表示一个 P 正在运行系统调用，而不是运行用户代码。
	// 它与处于系统调用中的 M 有关联，但不由它拥有，可能被其他 M 抢占。
	// 这类似于 _Pidle，但使用轻量级转换并保持 M 关联。
	//
	// 离开 _Psyscall 必须使用 CAS 操作，以偷取或重新获取 P。
	// 请注意存在 ABA 危险：
	// 即使 M 在系统调用后成功 CAS 回其原始 P 为 _Prunning，
	// 它也必须了解在此期间 P 可能已由其他 M 使用。
	_Psyscall

	// _Pgcstop 表示一个 P 已经停止运行，
	// 用于 Stop The World（STW），
	// 并由停止 World 的 M 拥有。
	// 停止 World 的 M 继续使用其 P，
	// 即使在 _Pgcstop 中也是如此。
	// 从 _Prunning 过渡到 _Pgcstop 会导致 M 释放其 P 并进入休眠状态。

	// P 保留其运行队列，
	// startTheWorld 将在具有非空运行队列的 P 上重新启动调度器。
	_Pgcstop

	// _Pdead 表示一个 P 不再使用（GOMAXPROCS 减小）。
	// 如果 GOMAXPROCS 增加，则重新使用 P。
	// 已死亡的 P 在大多数情况下会被剥夺其资源，
	// 尽管仍然有一些东西保留（例如，跟踪缓冲区）。
	_Pdead
)

// 互斥锁。在无争用的情况下，
// 与自旋锁一样快（只需几个用户级指令），
// 但在争用的情况下，它们会在内核中休眠。
// 零值的 Mutex 是未锁定的（无需初始化每个锁）。
// 对于静态锁排名，初始化是有帮助的，但不是必需的。
type mutex struct {
	// 如果禁用锁排名，则为空结构体，否则包括锁排名。
	lockRankStruct
	// 基于 Futex 的实现将其视为 uint32 键，
	// 而基于信号量的实现将其视为 M* waitm。
	// 以前是一个联合体，但联合体会破坏精确的垃圾收集。
	key uintptr
}

// 用于一次性事件的休眠和唤醒。
// 在调用 notesleep 或 notewakeup 之前，
// 必须调用 noteclear 来初始化 Note。
// 然后，确切地一个线程可以调用 notesleep，
// 确切地一个线程可以调用 notewakeup（一次）。
// 一旦调用了 notewakeup，notesleep 将返回。
// 以后的 notesleep 将立即返回。
// 后续的 noteclear 必须仅在
// 先前的 notesleep 返回后调用，例如，
// 在 notewakeup 后直接调用 noteclear 是不允许的。
//
// notetsleep 类似于 notesleep，但在给定的纳秒数之后唤醒，
// 即使事件尚未发生。如果 goroutine 使用 notetsleep 提前唤醒，
// 它必须等待调用 noteclear，直到可以确保没有其他 goroutine 在调用 notewakeup。
//
// notesleep/notetsleep 通常在 g0 上调用，
// notetsleepg 类似于 notetsleep，但在用户 g 上调用。
type note struct {
	// 基于 Futex 的实现将其视为 uint32 键，
	// 而基于信号量的实现将其视为 M* waitm。
	// 以前是一个联合体，但联合体会破坏精确的垃圾收集。
	key uintptr
}

type funcval struct {
	fn uintptr
	// 变长，fn 特定的数据在这里
}

type iface struct {
	tab  *itab
	data unsafe.Pointer
}

type eface struct {
	_type *_type
	data  unsafe.Pointer
}

func efaceOf(ep *any) *eface {
	return (*eface)(unsafe.Pointer(ep))
}

// guintptr、muintptr 和 puintptr 用于绕过写屏障。
// 特别是在当前 P 已经被释放的情况下，尽量避免写屏障非常重要，
// 因为 GC 认为世界已经停止，而意外的写屏障不会与 GC 同步，
// 这可能导致一个半执行的写屏障已标记对象但未排队。
// 如果 GC 跳过该对象并在排队发生之前完成，它将错误地释放该对象。
//
// 我们尝试使用仅在没有运行中的 P 时调用的特殊赋值函数，
// 但是某些特定内存字的更新通过写屏障，而另一些则没有。
// 这破坏了写屏障阴影检查模式，并且也令人担忧：
// 最好有一个完全被 GC 忽略的字，而不是仅有一些更新被忽略。
//
// G 和 P 始终可以通过 allgs 和 allp 列表或（在达到这些列表之前的分配期间）
// 从栈变量中的真指针访问。
//
// M 始终可以通过 true 指针从 allm 或 freem 访问。
// 与 G 和 P 不同，我们确实会释放 M，因此非常重要的是
// 在安全点之间绝对不要持有 muintptr。

// 一个 guintptr 包含一个 goroutine 指针，但被类型化为 uintptr
// 以绕过写屏障。它用于 Gobuf goroutine 状态
// 和在没有 P 的情况下操纵的调度列表。
//
// Gobuf.g 的 goroutine 指针几乎总是由汇编代码更新的。
// 在仅有一个少数几个 Go 代码更新它的地方 - 函数 save - 它必须被视为 uintptr，
// 以避免在不好的时候生成写屏障。
// 而不是弄清楚如何在汇编操作中缺失的写屏障中发出写屏障，
// 我们将字段的类型更改为 uintptr，
// 以便根本不需要写屏障。
//
// Goroutine 结构体发布在 allg 列表中并且永不释放。
// 这将阻止 goroutine 结构体被 GC 回收。
// 永远不会出现 Gobuf.g 包含对 goroutine 的唯一引用的情况：
// 首先在 allg 中发布 goroutine。
// Goroutine 指针也保留在不对 GC 可见的地方，如 TLS 中，
// 因此我无法想象它们会移动。
// 如果我们确实想要在 GC 中开始移动数据，
// 那么我们需要从替代的 arena 中分配 goroutine 结构体。
// 使用 guintptr 不会使这个问题更糟糕。
// 请注意，pollDesc.rg、pollDesc.wg 也以 uintptr 形式存储 g，
// 因此如果 g 开始移动，它们也需要被更新。
type guintptr uintptr

//go:nosplit
func (gp guintptr) ptr() *g { return (*g)(unsafe.Pointer(gp)) }

//go:nosplit
func (gp *guintptr) set(g *g) { *gp = guintptr(unsafe.Pointer(g)) }

//go:nosplit
func (gp *guintptr) cas(old, new guintptr) bool {
	return atomic.Casuintptr((*uintptr)(unsafe.Pointer(gp)), uintptr(old), uintptr(new))
}

//go:nosplit
func (gp *g) guintptr() guintptr {
	return guintptr(unsafe.Pointer(gp))
}

// setGNoWB 执行 *gp = new 而没有写屏障。
// 对于在实际使用 guintptr 不切实际的情况。
//
//go:nosplit
//go:nowritebarrier
func setGNoWB(gp **g, new *g) {
	(*guintptr)(unsafe.Pointer(gp)).set(new)
}

type puintptr uintptr

//go:nosplit
func (pp puintptr) ptr() *p { return (*p)(unsafe.Pointer(pp)) }

//go:nosplit
func (pp *puintptr) set(p *p) { *pp = puintptr(unsafe.Pointer(p)) }

// muintptr 是一个 *m，不受垃圾收集器跟踪。
//
// 因为我们会释放 M，muintptr 有一些额外的限制：
//
//  1. 永远不要在安全点之间持有本地的 muintptr。
//
//  2. 堆中的任何 muintptr 必须由 M 本身拥有，
//     以便它能够确保在释放最后一个真正的 *m 时它没有在使用。
type muintptr uintptr

//go:nosplit
func (mp muintptr) ptr() *m { return (*m)(unsafe.Pointer(mp)) }

//go:nosplit
func (mp *muintptr) set(m *m) { *mp = muintptr(unsafe.Pointer(m)) }

// setMNoWB 执行 *mp = new 而没有写屏障。
// 对于在实际使用 muintptr 不切实际的情况。
//
//go:nosplit
//go:nowritebarrier
func setMNoWB(mp **m, new *m) {
	(*muintptr)(unsafe.Pointer(mp)).set(new)
}

type gobuf struct {
	// sp、pc 和 g 的偏移量已知（在 libmach 中硬编码）。
	//
	// 与 GC 相关的 ctxt 对于 GC 来说是不同寻常的：它可能是一个堆分配的 funcval，
	// 因此 GC 需要跟踪它，但是它需要在汇编中设置和清除，
	// 在汇编中很难使用写屏障。然而，ctxt 实际上是一个保存的、活动的寄存器，
	// 我们只在真正的寄存器和 gobuf 之间交换它。因此，在堆栈扫描期间，
	// 我们将其视为根，这意味着从保存和恢复它的汇编中不需要写屏障。
	// 它仍然被声明为指针，以便来自 Go 的任何其他写入都带有写屏障。
	sp   uintptr
	pc   uintptr
	g    guintptr
	ctxt unsafe.Pointer
	ret  uintptr
	lr   uintptr
	bp   uintptr // 适用于启用帧指针的架构
}

// sudog 表示在等待列表中的 g，例如在通道上发送/接收时。
//
// sudog 是必需的，因为 g ↔ 同步对象的关系是多对多的。
// 一个 g 可以在许多等待列表中，因此可能有很多 sudog 对应一个 g；
// 同样，许多 g 可能在等待同一个同步对象，因此可能有很多 sudog 对应一个对象。
//
// sudog 从一个特殊的池中分配。使用 acquireSudog 和
// releaseSudog 来分配和释放它们。
type sudog struct {
	// 下面的字段由此 sudog 阻塞的通道的 hchan.lock 保护。
	// shrinkstack 依赖于此 sudog 在通道操作中的情况。

	g *g

	next *sudog
	prev *sudog
	elem unsafe.Pointer // 数据元素（可能指向堆栈）

	// 下面的字段永远不会被并发访问。
	// 对于通道，waitlink 仅由 g 访问。
	// 对于信号量，只有在持有 semaRoot 锁时才会访问所有字段（包括上面的字段）。

	acquiretime int64
	releasetime int64
	ticket      uint32

	// isSelect 表示 g 参与了 select，因此
	// 必须对 g.selectDone 进行 CAS 以赢得唤醒竞争。
	isSelect bool

	// success 表示通信是否成功。
	// 如果 goroutine 被唤醒是因为在通道 c 上传递了值，则为 true；
	// 如果是因为关闭了 c，则为 false。
	success bool

	parent   *sudog // semaRoot 二叉树
	waitlink *sudog // g.waiting 列表或 semaRoot
	waittail *sudog // semaRoot
	c        *hchan // 通道
}

type libcall struct {
	fn   uintptr
	n    uintptr // 参数数量
	args uintptr // 参数
	r1   uintptr // 返回值
	r2   uintptr
	err  uintptr // 错误号
}

// Stack 描述一个 Go 执行栈。
// 栈的边界正好是 [lo, hi)，两边没有隐含的数据结构。
type stack struct {
	lo uintptr
	hi uintptr
}

// heldLockInfo 提供有关持有的锁及该锁的排名的信息
type heldLockInfo struct {
	lockAddr uintptr
	rank     lockRank
}
type g struct {
	// 栈参数。
	// stack 描述实际的堆栈内存：[stack.lo, stack.hi)。
	// stackguard0 是在 Go 堆栈增长序言中比较的堆栈指针。
	// 它通常为 stack.lo+StackGuard，但可以为 StackPreempt 以触发抢占。
	// stackguard1 是在 C 堆栈增长序言中比较的堆栈指针。
	// 在 g0 和 gsignal 堆栈上为 stack.lo+StackGuard。
	// 在其他 goroutine 堆栈上为 ~0，以触发调用 morestackc（并崩溃）。
	stack       stack   // 偏移已知于 runtime/cgo
	stackguard0 uintptr // 偏移已知于 liblink
	stackguard1 uintptr // 偏移已知于 liblink

	_panic    *_panic // 最内层的 panic - 偏移已知于 liblink
	_defer    *_defer // 最内层的延迟
	m         *m      // 当前 m；偏移已知于 arm liblink
	sched     gobuf
	syscallsp uintptr // 如果 status==Gsyscall，则 syscallsp = sched.sp 用于 gc 期间使用
	syscallpc uintptr // 如果 status==Gsyscall，则 syscallpc = sched.pc 用于 gc 期间使用
	stktopsp  uintptr // 期望堆栈顶部的 sp，用于在 traceback 中检查
	// param 是一个通用指针参数字段，用于在其他情境中传递参数，
	// 在那里为参数找到其他存储会很困难。它当前用于三种方式：
	// 1. 当通道操作唤醒了被阻塞的 goroutine 时，它将 param 设置为
	//    指向已完成的阻塞操作的 sudog。
	// 2. 通过 gcAssistAlloc1 向其调用者发出信号，表明 goroutine 完成了
	//    GC 周期。以其他方式这样做是不安全的，因为 goroutine 的堆栈
	//    在此期间可能已经移动了。
	// 3. 通过 debugCallWrap 传递参数给新 goroutine，因为在运行时分配
	//    闭包是被禁止的。
	param        unsafe.Pointer
	atomicstatus atomic.Uint32
	stackLock    uint32 // sigprof/scang 锁；TODO：折叠到 atomicstatus 中
	goid         uint64
	schedlink    guintptr
	waitsince    int64      // g 变为阻塞状态的近似时间
	waitreason   waitReason // 如果 status==Gwaiting

	preempt       bool // 抢占信号，与 stackguard0 = stackpreempt 重复
	preemptStop   bool // 在抢占时转换到 _Gpreempted；否则，只是取消调度
	preemptShrink bool // 在同步安全点收缩堆栈

	// asyncSafePoint 如果 g 在异步安全点被停止，则设置为 true。
	// 这意味着堆栈上有没有精确指针信息的帧。
	asyncSafePoint bool

	paniconfault bool // 在意外故障地址上发生 panic（而不是崩溃）
	gcscandone   bool // g 已扫描堆栈；在状态的 _Gscan 位中受保护
	throwsplit   bool // 不能分割堆栈
	// activeStackChans 表示有未锁定的通道指向该 goroutine 的堆栈。
	// 如果为 true，堆栈复制需要获取通道锁以保护堆栈的这些区域。
	activeStackChans bool
	// parkingOnChan 表示 goroutine 即将在 chansend 或 chanrecv 上停泊。
	// 用于在堆栈收缩的不安全点发出信号。
	parkingOnChan atomic.Bool

	raceignore    int8  // 忽略 race 检测事件
	tracking      bool  // 我们是否正在跟踪此 G 以获取调度延迟统计信息
	trackingSeq   uint8 // 用于决定是否跟踪此 G
	trackingStamp int64 // G 上次开始被跟踪的时间戳
	runnableTime  int64 // 处于可运行状态的时间，运行时清零，仅在跟踪时使用
	lockedm       muintptr
	sig           uint32
	writebuf      []byte
	sigcode0      uintptr
	sigcode1      uintptr
	sigpc         uintptr
	parentGoid    uint64          // goid of goroutine that created this goroutine
	gopc          uintptr         // pc of go statement that created this goroutine
	ancestors     *[]ancestorInfo // ancestor information goroutine(s) that created this goroutine (only used if debug.tracebackancestors)
	startpc       uintptr         // pc of goroutine function
	racectx       uintptr
	waiting       *sudog         // sudog structures this g is waiting on (that have a valid elem ptr); in lock order
	cgoCtxt       []uintptr      // cgo traceback context
	labels        unsafe.Pointer // profiler labels
	timer         *timer         // cached timer for time.Sleep
	selectDone    atomic.Uint32  // are we participating in a select and did someone win the race?

	// goroutineProfiled 表示此 goroutine 的堆栈在当前进行中的 goroutine 配置文件中的状态
	goroutineProfiled goroutineProfileStateHolder

	// Per-G tracer state.
	trace gTraceState

	// Per-G GC state

	// gcAssistBytes 是此 G 的 GC 辅助信用，以字节为单位。
	// 如果这是正数，则 G 具有分配 gcAssistBytes 字节的信用，而无需辅助。
	// 如果这是负数，则 G 必须通过执行扫描工作来纠正。
	// 我们以字节为单位跟踪这一点，以便在 malloc 热路径中快速更新和检查债务。
	// 辅助比率确定这与扫描工作债务的对应关系。
	gcAssistBytes int64
}
