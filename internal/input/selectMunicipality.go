package input

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

var MunicipalityMap = map[string]string{
	"ukiha":          "918349529339",
	"dazaifu":        "595704640573",
	"hirao":          "474812254471",
	"hakozaki":       "706066125025",
	"hita":           "279885203651",
	"miyama":         "455202378341",
	"miyako":         "974755246373",
	"paycha":         "908288321433",
	"yame":           "967225779452",
	"itoshima":       "197260688168",
	"arao":           "131270676423",
	"yukuhashi":      "478143842994",
	"asakura":        "943472481426",
	"saijo":          "476406248251",
	"iizuka":         "563812935552",
	"nisshin":        "413760224386",
	"munakata":       "064370235682",
	"tanushimaru":    "995908658351",
	"yanagawa":       "672278509583",
	"ogori":          "793090307814",
	"onga":           "521243538360",
	"takamiya":       "829654954452",
	"omuta":          "627993447010",
	"machinowa":      "646306146919",
	"shimagin":       "971815107672",
	"amakusa":        "628012664431",
	"yamagata":       "017646457068",
	"saga":           "688640212629",
	"yamaguchi":      "172741592625",
	"himeji":         "896445240154",
	"kotake":         "312260539058",
	"chikuzen":       "255832920986",
	"chikugo":        "807563387437",
	"tamana":         "098294295143",
	"matsubara":      "990293378393",
	"yufu":           "093908553536",
	"kashima":        "326485595931",
	"shime":          "346067361561",
	"fukutsu":        "589094778815",
	"osaki":          "383549634189",
	"taketa":         "830381541625",
	"tsuruoka":       "496577377651",
	"koga":           "385306393166",
	"fukuoka":        "525472949674",
	"tosu":           "976477627051",
	"seki":           "412796914019",
	"kinokawa":       "271329952012",
	"oita":           "045039219352",
	"adachi":         "913798776784",
	"hitachiota":     "854606318826",
	"soeda":          "707741866577",
	"kasuga":         "557868907783",
	"sue":            "121592653421",
	"hirokawa":       "715923893570",
	"kurume":         "786685663162",
	"karatsu":        "337181156513",
	"tachiarai":      "880460826592",
	"satsumasendai":  "127724138105",
	"nakagawa":       "652232027595",
	"oonojo":         "263351959420",
	"ootou":          "976564081064",
	"kasuya":         "449974704444",
	"hisayama":       "435978619777",
	"kanda":          "977409549427",
	"kouge":          "957047467321",
	"sasaguri":       "518143405820",
	"inazawa":        "981265391805",
	"imizu":          "603016525483",
	"kurate":         "564720655459",
	"miyazaki":       "413136046724",
	"ookawa":         "087862488451",
	"itoda":          "840875560135",
	"nakama":         "089371766913",
	"buzen":          "505233464539",
	"sunroser":       "363062475663",
	"ooki":           "991260682180",
	"miyawaka":       "885487785656",
	"imari":          "957311760787",
	"chikushino":     "588585261400",
	"kahoku":         "797788541805",
	"nagasaki":       "151878056159",
	"chita":          "279245857502",
	"keisen":         "379858985566",
	"anshincoin":     "629645173701",
	"chkpp1":         "121163660197",
	"furusato":       "595823922675",
	"shizuoka":       "211125665249",
	"sapporo":        "533266982297",
	"shinshiro":      "590183855134",
	"yokatoku":       "381492163811",
	"osaka":          "471112789084",
	"shingu":         "992382440775",
	"kashiihama":     "010438506961",
	"nishijin":       "905418304851",
	"takeo":          "867344451992",
	"corda":          "283372850953",
}

// 自治体名の選択または手動入力
func GetMunicipalityInput(promptMessage string) string {
	// 自治体名リストを取得
	municipalityNames := make([]string, 0, len(MunicipalityMap))
	for name := range MunicipalityMap {
		municipalityNames = append(municipalityNames, name)
	}

	// サーチ関数の定義
	searcher := func(input string, index int) bool {
		municipality := municipalityNames[index]
		return strings.Contains(strings.ToLower(municipality), strings.ToLower(input))
	}

	prompt := promptui.Select{
		Label:             promptMessage,
		Items:             municipalityNames,
		Size:              10,  // 表示する最大件数
		Searcher:          searcher,
		StartInSearchMode: true, // 検索モードをデフォルトでON
	}

	// ユーザーの選択
	index, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("自治体の選択に失敗しました:(インデックス: %d) %s", index, err)
		os.Exit(1)
	}

	return result
}
