package main
import "fmt"
func main(){
	var str string
	fmt.Scanf("%s",&str)
	res:=reverse(str)
	fmt.Println(res)
}
func reverse(str string) (res string ){
	for i:=len(str)-1;i>=0;i--{
		res+=string(str[i])
	}
	return

}