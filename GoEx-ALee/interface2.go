package main

func main() {
    var a interface{} = 1

    i := a       // a와 i 는 dynamic type, 값은 1
    j := a.(int) // j는 int 타입, 값은 1

    println(i)  // 포인터주소 출력
    println(j)  // 1 출력
}
