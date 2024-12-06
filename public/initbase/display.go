package initbase

import "fmt"

func Display() {
	content := `
欢迎使用Mirage的命令控制服务器！
Mirageのコマンド制御サーバへようこそ！
Welcome to using the Mirage command to control the server!

 _____ ______   ___  ________  ________  ________  _______        
|\   _ \  _   \|\  \|\   __  \|\   __  \|\   ____\|\  ___ \     
\ \  \\\__\ \  \ \  \ \  \|\  \ \  \|\  \ \  \___|\ \   __/|     
 \ \  \\|__| \  \ \  \ \   _  _\ \   __  \ \  \  __\ \  \_|/__      
  \ \  \    \ \  \ \  \ \  \\  \\ \  \ \  \ \  \|\  \ \  \_|\ \     
   \ \__\    \ \__\ \__\ \__\\ _\\ \__\ \__\ \_______\ \_______\    
    \|__|     \|__|\|__|\|__|\|__|\|__|\|__|\|_______|\|_______|

你可以开始输入命令或者输入help查看帮助信息！
コマンドの入力を開始するか、ヘルプ情報を表示するhelpを入力することができます！
You can start typing commands or enter help to view help information！
`
	fmt.Println(content)
}
