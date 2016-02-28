package pinyingo

import (
  "github.com/yanyiwu/gojieba"
  "os"
  "regexp"
  "strings"
)

const (
  STYLE_NORMAL       = 1
  STYLE_TONE         = 2
  STYLE_INITIALS     = 3
  STYLE_FIRST_LETTER = 4
  USE_SEGMENT        = true
  NO_SEGMENT         = false
  use_hmm            = true
)

var reg *regexp.Regexp
var INITIALS []string = strings.Split("b,p,m,f,d,t,n,l,g,k,h,j,q,x,r,zh,ch,sh,z,c,s", ",")
var keyString string
var jieba *gojieba.Jieba
var sympolMap = map[string]string{
  "ā": "a1",
  "á": "a2",
  "ǎ": "a3",
  "à": "a4",
  "ē": "e1",
  "é": "e2",
  "ě": "e3",
  "è": "e4",
  "ō": "o1",
  "ó": "o2",
  "ǒ": "o3",
  "ò": "o4",
  "ī": "i1",
  "í": "i2",
  "ǐ": "i3",
  "ì": "i4",
  "ū": "u1",
  "ú": "u2",
  "ǔ": "u3",
  "ù": "u4",
  "ü": "v0",
  "ǘ": "v2",
  "ǚ": "v3",
  "ǜ": "v4",
  "ń": "n2",
  "ň": "n3",
  "": "m2",
}

func init() {
  keyString = getMapKeys()
  reg = regexp.MustCompile("([" + keyString + "])")
  dictPath := getDictPath()
  //初始化时将gojieba实例化到内存
  jieba = gojieba.NewJieba(dictPath+"jieba.dict.utf8", dictPath+"hmm_model.utf8", dictPath+"user.dict.utf8")
}

func getMapKeys() string {
  keyString := ""
  for key, _ := range sympolMap {
    keyString += key
  }
  return keyString
}

func normalStr(str string) string {
  findRet := reg.FindString(str)
  return strings.Replace(str, findRet, string([]byte(sympolMap[findRet])[0]), -1)
}

//获取文件所在的根目录
func getDictPath() string {
  currentPath, _ := os.Getwd()
  return currentPath + "/src/github.com/struCoder/Go-pinyin/dict/"
}