package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	text := `
　 　 　 　 　 ／////////////////////////　------＼/////////ヽ
　　　　　　 ////:///////////////＞　　"´　　　　　　　 ∨////////∧
　　　　　 .///////////////＞ ´　　　　　　　　　　　 　 　 ∨////////∧
　　　　　/////////／/＞ ´　　　　　　　　　　　　　　　　 　 ∨////////∧
　　　　　|///////,〃／　　　　　　　　　　　　　　　　　　 　 　 i//////////ﾊ
　　　　　|////////i/　　　　　　　　　　　　_　　--――― ､--}//////////∧
　　　　　|////////|　　　　　　　　　　　 ´　　　　ﾄ.　　　　　ヽ ＼'/////////∧
　　　　　|////////|　　　　　 ..　 イ　　　　　　　 i　＼　　 ノ　 ',　＼////////∧
　　　　　|////////| 　 .　　 ´i　　 i　　 　 　 　 ｰ|-- 斗､.　　　 ',　　＼///////∧
　　　　　|////////| ／　　　 |　　ハ　　　 ',　 ∧!　,ｨf＝ミ､　　　',ー― ///////∧
　　　　　|////////|' 　　.、＿|..斗七ヽﾄ､ 　 ヽ| u 　 {::♥::} ゞ ＼　',i///////////,∧
　　　　　|////////|　　　　　 |/i/,.ィf=ミ ヽト､!|　　 ｀¨¨¨´/　　/＼!.ﾏニ ア//////,∧
　　　　　|////////ﾊ　　　 　/ .ｨ f:::♥ﾉ 　 | / ／／／／/　／　i　　|=イ/////////∧
　　　　　|////////,∧　 　 /{　ゞ ｀¨´ ／, j/　／／／ ./イ /　 .|　　i　 i///////////
　　　　　|/////////|.ﾍ　 /　乂　／／／／／　　　 　 u　/　　 |　　|　 i!///////////
　　　　　|/////////＞j／　　　＼ 　　　　　 　 　＿　　　j/|　　.| ∧!　　!///////////
　　　　　|//////／ ∠ｨ ﾍ 　　　　＞_ 　　 　 ,ｨv´　 _)　　/ |　　|/:::::∧ .|V//////////
　　　　　|//////ゝ---一 ﾍ　　ゝ----' u 　 ゝ-　´　 　 /i .|　　|:::::::::::∧!::∨/////////
　　　　　|//////////|　i　∧　　　　',≧=-　　 ＿＿　.イ　!ﾊ　　|:::::::::::::＿:::乂////////
　　　　　|//////////| ∧　∧　　　ト{ 　　　}　　　　i　　　 {ｰﾍ　!´￣　／ アニ≧= ---
　　　　　|//////////|/::::ﾍ.　 ',　　 |从　__ノi　　　　',　　　ﾉ　 ﾍ|　　／ ／ニニニニニニ
　　　　　|//////////|:::::::::::＼{ﾍ　 .|-＜　　 ゝ　　　 r ､ -､　　　　 /　 /ニニニニニニニ
　　　　　|/////////ノ==-＜　　 ',　|　　　　　　　　 r' 　 く　　　　　　./ニニニニニニﾆﾆ
　　　　 .ﾉ//＞≦ニニニニニ＼　 }/　　　　　　　 　 ｀ ´ ＼＼　　　　/ニニニニニニニニ
　　　 .//／ニニニニニニニニﾆ＼　　　　　　　　　　　　　 ｀´　　　　{ﾆﾆニニニニニニニ
　　　 .イニニニニニニニニニニニヽ　　　　　　　　　　　　　　　　　 /ニニニニニニニﾆﾆ
	`
	if b64, err := EncodeEncodeBase64Base64(text); err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println(b64)
		if aa, err := DecodeBase64Gzip(b64); err != nil {
			fmt.Println("err:", err)
		} else {
			fmt.Println(string(aa))
		}
	}

}

func EncodeEncodeBase64Base64(str string) (string, error) {
	var buf64 bytes.Buffer
	wb64 := base64.NewEncoder(base64.URLEncoding, &buf64)
	wgz := gzip.NewWriter(wb64)
	_, err := wgz.Write([]byte(str))
	if err != nil {
		return "", err
	}
	wgz.Close()
	wb64.Close()
	sRequest64 := string(buf64.Bytes())
	return sRequest64, nil
}

func DecodeBase64Gzip(b64 string) ([]byte, error) {
	if r, err := gzip.NewReader(base64.NewDecoder(base64.URLEncoding, strings.NewReader(b64))); err != nil {
		fmt.Println(err)
		return nil, err
	} else {

		if s, err := ioutil.ReadAll(r); err != nil {
			fmt.Println(err)
			return nil, err
		} else {
			return s, nil
		}
	}
}
