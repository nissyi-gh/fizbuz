package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 引数を取得
	flag.Parse()
	args := flag.Args()

	// FizBuzを実行
	o, err := FizBuz(args[0])
	if err != nil {
		// 標準エラーにエラーを表示
		_, _ = fmt.Fprintln(os.Stderr, err)
		// exit status 1 で終了
		os.Exit(1)
	}

	// 結果を表示
	fmt.Println(o)
}

func FizBuz(in string) (string, error) {
	// string型をint型に変換
	i, err := strconv.Atoi(in)
	if err != nil {
		return "", err
	}
	switch {
	case i%3 == 0:
		// 3の倍数ならFizz
		return "Fizz", nil
	case i%5 == 0:
		// 5の倍数ならBuzz
		return "Buzz", nil
	case i%15 == 0:
		// 15の倍数ならFizzBuzz
		return "FizzBuzz", nil
	}
	return "", fmt.Errorf("判定できませんでした: %s", in)
}
