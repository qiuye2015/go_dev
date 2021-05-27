package main

func main() {
	a := App{}
	a.InitApp()
	a.InitRouter()
	a.Run(":9999")
}
