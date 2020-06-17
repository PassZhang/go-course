// 变量 泛型

/*
泛型有什么用？
定义了一个add函数，对应的接收类型为int型，如果传其他类型，就会报错
func add(num1 int, num2 int)
...
*/

// 这里定义了一个泛型类，如果在实例化的时候传入int型，那么响应的泛型方法的接收参数也会变成int型

```Java // 这是一段java代码
public class Box<T> {
	private T t;
	public void add(T t) {
		this.t = t;
	}
	public T get() {
		return t;
	}
	public static void main(String[] args) {
		Box<Integer> integerBox = new Box<Integer>();
		Box<String> stringBox = new Box<String>();

		integerBox.add(new Integer(10));
		stringBox.add(new String("Hello World!"));

		System.out.printf("整型值为 :%d\n\n", integerBox.get());
		System.out.printf("字符串为 :%s\n", stringBox.get());
	}
}

```
/*
go里面有泛型吗？

目前没有泛型，Go2.0 将会支持
目前Go的版本是1.11, 在Go2之前预计还有4个大版本，Go每半年发布一个大版本，至少两年后才会整理Go2提案。
*/ 