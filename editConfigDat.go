package main

import (
	"fmt"
	"os"
)

// Definindo a sequência de caracteres na ordem que o daybreak usa
var chars = []rune{
	' ', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'@', '+', '-', '+', '/',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', '#', '!', '?', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'<', '>', '=',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', '%', '&', '^', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'(', ')', '_',
	'あ', 'か', 'さ', 'た', 'な', 'は', 'な', 'ま', 'ら', 'わ', 'が', 'ざ', 'だ', 'ば', 'ぱ', 'ぁ', 'い', 'き', 'し', 'ち', 'に', 'ひ', 'み', 'ゃ', 'り', 'ょ', 'ぎ', 'じ', 'ぢ', 'び', 'ぴ', 'ぃ', 'う', 'く', 'す', 'つ', 'ぬ', 'ふ', 'む', 'ゆ', 'る', 'を', 'ぐ', 'ず', 'づ', 'ぶ', 'ぷ', 'ぅ', 'え', 'け', 'せ', 'て', 'ね', 'へ', 'め', 'ゅ', 'れ', 'っ', 'げ', 'ぜ', 'で', 'べ', 'ぺ', 'ぇ', 'お', 'こ', 'ぞ', 'と', 'の', 'ほ', 'も', 'よ', 'ろ', 'ん', 'ご', 'ぞ', 'ど', 'ぼ', 'ぽ', 'ぉ',
	'ア', 'カ', 'サ', 'タ', 'ナ', 'ハ', 'ナ', 'マ', 'ラ', 'ワ', 'ガ', 'ザ', 'ダ', 'バ', 'パ', 'ァ', 'イ', 'キ', 'シ', 'チ', 'ニ', 'ヒ', 'ミ', 'ャ', 'リ', 'ョ', 'ギ', 'ジ', 'ヂ', 'ビ', 'ピ', 'ィ', 'ウ', 'ク', 'ス', 'ツ', 'ヌ', 'フ', 'ム', 'ユ', 'ル', 'ヲ', 'グ', 'ズ', 'ヅ', 'ブ', 'プ', 'ゥ', 'エ', 'ケ', 'セ', 'テ', 'ネ', 'ヘ', 'メ', 'ュ', 'レ', 'ッ', 'ゲ', 'ゼ', 'デ', 'ベ', 'ペ', 'ェ', 'オ', 'コ', 'ゾ', 'ト', 'ノ', 'ホ', 'モ', 'ヨ', 'ロ', 'ン', 'ゴ', 'ゾ', 'ド', 'ボ', 'ポ', 'ォ',
	'ヴ', 'ー', '～', '、', '。', '￥', '$', '"', '\'', '`', ',', '.', ';', ':', '♪', '※', '愛', '哀', '悪', '一', '隠', '右', '雨', '雲', '円', '園', '炎', '王', '屋', '音', '仮', '暇', '家', '夏', '火', '花', '華', '解', '壊', '改', '械', '皆', '外', '楽', '活', '間', '丸', '機', '気', '鬼', '技', '偽', '魚', '京', '凶', '強', '教', '抂', '玉', '業', '金', '銀', '九', '空', '軍', '係', '兄', '形', '型', '撃', '血', '月', '剣', '犬', '幻', '古', '言', '五', '御', '悟', '光', '公', 'ロ', '好', '甲', '皇', '紅', '行', '号', '国', '黒', '魂', '左', '歳', '祭', '罪', '崎', '桜', '殺', '三', '史', '四', '士', '子', '姉', '死', '詩', '寺', '時', '次', '自', '式', '七', '実', '社', '者', '弱', '主', '守', '手', '首', '終', '秋', '十', '銃', '重', '春', '女', '勝', '将', '小', '少', '笑', '上', '城', '情', '心', '新', '深', '真', '神', '身', '人', '水', '世', '性', '正', '生', '精', '聖', '西', '青', '石', '赤', '絶', '先', '千', '川', '戦', '全', '組', '装', '即', '族', '属', '村', '地', '多', '隊', '大', '団', '弾', '断', '男', '池', '中', '虫', '町', '超', '長', '鳥', '直', '帝', '弟', '的', '天', '店', '電', '徒', '土', '怒', '党', '冬', '刀', '東', '堂', '道', '特', '内', '二', '日', '忍', '猫', '波', '派', '破', '白', '爆', '八', '半', '犯', '美', '飛', '必', '姫', '百', '不', '父', '武', '風', '物', '分', '文', '平', '母', '方', '亡', '某', '北', '本', '魔', '妹', '魅', '万', '末', '命', '夢', '無', '娘', '明', '滅', '木', '門', '夜', '友', '遊', '夕', '幼', '様', '雷', '理', '裏', '里', '立', '流', '竜', '良', '力', '麗', '零', '六', '♦', '⋆', '†', 'д', '₃', 'А', '￣', '∀', 'ф', 'د', '°',
}

const (
	outline     = 212
	shadows     = 213
	higerResTex = 214
	//ingameFullScreen = 215 // ligar isso força tela cheia em 800x600
)

// Posição dos dados no arquivo config
const cfgNameStart = 216
const cfgPagesStart = 224
const nameLength = 8

func getIndex(char rune) int {
	for i, v := range chars {
		if v == char {
			return i
		}
	}
	return -1
}

func getChar(index byte, page byte) rune {
	if int(index) >= len(chars) || page > 1 {
		return ' '
	}
	pageSize := len(chars) / 2
	return chars[int(page)*pageSize+int(index)]
}

func getName(config []byte) string {
	pages := config[cfgPagesStart]
	name := make([]rune, nameLength)
	for i := 0; i < nameLength; i++ {
		index := config[cfgNameStart+i]
		page := (pages >> i) & 1
		name[i] = getChar(index, byte(page))
	}
	return string(name)
}


func setName(config []byte, name string) {
	name = name[:nameLength] // Limita o nome a 8 caracteres
	pages := byte(0)
	for i := range nameLength {
		charIndex := getIndex(rune(name[i]))
		config[cfgNameStart+i] = byte(charIndex)
		if charIndex >= len(chars)/2 {
			pages |= (1 << i) // Define o bit da página
		}
	}
	config[cfgPagesStart] = pages
}

func loadConfigFile() ([]byte, string, error) {
	filePath := "./config.dat"
	config, err := os.ReadFile(filePath)
	if err != nil {
		return nil, "", fmt.Errorf("error opening config file: %w", err)
	}
	return config, filePath, nil
}

func getNetworkName() (string, error) {
	config, _, err := loadConfigFile()
	if err != nil {
		return "", err
	}
	name := getName(config)
	return name, nil
}

func setNetworkName(newName string) error {
	// Abrir o arquivo de configuração
	config, filePath, err := loadConfigFile()
	if err != nil {
		return err
	}
	if len(newName) < 8 {
		// Adicionar espaços em branco para completar 8 caracteres
		for i := len(newName); i < 8; i++ {
			newName += " "
		}
	}
	if len(newName) > 8 {
		fmt.Println("name must be 8 characters")
		os.Exit(1)
	}
	// verificar os caracteres em newName, para garantir que estão na lista de caracteres
	for _, c := range newName {
		if getIndex(c) == -1 {
			fmt.Println("invalid character in name:", c)
			os.Exit(1)
		}
	}
	setName(config, newName)

	// Reescrever o arquivo de configuração
	err = os.WriteFile(filePath, config, 0644)
	if err != nil {
		fmt.Println("error writing the config file:", err)
		os.Exit(1)
	}

	fmt.Println("name changed:", newName)
	return nil
}

func getBoolConfig(location int) (bool, error) {
	config, _, err := loadConfigFile()
	if err != nil {
		return false, err
	}
	return config[location] != 0, nil
}

func setBoolConfig(location int, value bool) error {
	config, filePath, err := loadConfigFile()
	if err != nil {
		return err
	}
	if value {
		config[location] = 1
	} else {
		config[location] = 0
	}
	err = os.WriteFile(filePath, config, 0644)
	if err != nil {
		return fmt.Errorf("erro ao escrever no arquivo: %w", err)
	}
	return nil
}
