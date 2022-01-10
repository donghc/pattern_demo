package factorymethod

// 工厂方法模式使用子类的方式延迟生成对象到子类中实现。
//工厂方法模式属于类的创建型模式又被称为多态工厂模式。
//工厂方法模式的意义在于定义一个创建产品对象的工厂接口，将实际创建工作推迟到子类当中。
//核心工厂类不再负责产品的创建，仅负责声明具体工厂子类必须实现的接口。
//这样进一步抽象化的好处是使得工厂方法模式可以使系统在不修改具体工厂角色的情况下引进新的产品。
// Go中不存在继承 所以使用匿名组合来实现

//Operator 是被封装的实际类接口
type Operator interface {
	i()
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase 是Operator接口的实现基类，封装公用方法
type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) i() {
}

func (o *OperatorBase) SetA(i int) {
	o.a = i
}

func (o *OperatorBase) SetB(i int) {
	o.b = i
}

//PlusOperatorFactory 是 PlusOperator 的工厂类
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

//PlusOperator Operator 的实际加法实现
type PlusOperator struct {
	*OperatorBase
}

//Result 获取结果
func (o PlusOperator) Result() int {
	return o.a + o.b
}

//MinusOperatorFactory 是 MinusOperator 的工厂类
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

//MinusOperator Operator 的实际减法实现
type MinusOperator struct {
	*OperatorBase
}

//Result 获取结果
func (o MinusOperator) Result() int {
	return o.a - o.b
}
