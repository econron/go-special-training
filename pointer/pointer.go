package pointer

import "fmt"

func Pointer(){
	// 初期化
	v := 1
	// vのアドレスを取得
	p := &v
	// アスタリスクを使って実体を参照し変更
	*p = 2
	// 2
	fmt.Println(v)
	// アドレス
	fmt.Println(p)
	// 2
	fmt.Println(*p)
}

type Anime struct {
	Name string
	BroadcastTime string
	OriginalAuthor string
}

func SetAnime(name string, broadcastTime string, originalAuthor string) Anime {

	a := Anime {
		Name : name,
		BroadcastTime : broadcastTime,
		OriginalAuthor : originalAuthor,
	}

	return a

}

func PrintAnimeWithStructCopy(anime Anime){
	fmt.Println(anime.Name)
	fmt.Println(anime.BroadcastTime)
	fmt.Println(anime.OriginalAuthor)
}

func PrintAnimeWithoutStructCopy(anime *Anime){
	fmt.Println(anime.Name)
	fmt.Println(anime.BroadcastTime)
	fmt.Println(anime.OriginalAuthor)
}
