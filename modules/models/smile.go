package models

import (

	bbcode "github.com/frustra/bbcode"
	"strings"
)

var Smiles map[string]string
var Compiler bbcode.Compiler
func init()  {
	Smiles = map[string]string{
	"^:)^" : "https://vozforums.com/images/smilies/Off/lay.gif",
	":hang:" : "https://vozforums.com/images/smilies/Off/hang.gif",
		":phone:" : "https://vozforums.com/images/smilies/Off/phone.gif",
		":lmao:" : "https://vozforums.com/images/smilies/Off/lmao.gif",
		":bye:" : "https://vozforums.com/images/smilies/Off/bye.gif",
		":flame:" : "https://vozforums.com/images/smilies/Off/flame.gif",
		":bang:" : "https://vozforums.com/images/smilies/Off/bang.gif",
		":hug:" : "https://vozforums.com/images/smilies/Off/hug.gif",
		":fix:" : "https://vozforums.com/images/smilies/Off/fix.gif",
		":cheers:" : "https://vozforums.com/images/smilies/Off/cheers.gif",
		"-_-" : "https://vozforums.com/images/smilies/Off/sleep.gif",
		":shitty:" : "https://vozforums.com/images/smilies/Off/shit.gif",
		":rofl:" : "https://vozforums.com/images/smilies/Off/rofl.gif",
		":capture:" : "https://vozforums.com/images/smilies/Off/capture.gif",
		":theft:" : "https://vozforums.com/images/smilies/Off/theft.gif",
		":spam:" : "https://vozforums.com/images/smilies/Off/spam.gif",
		":hehe:" : "https://vozforums.com/images/smilies/Off/hehe.gif",
		":smoke:" : "https://vozforums.com/images/smilies/Off/smoke.gif",
		":sos:" : "https://vozforums.com/images/smilies/Off/sos.gif",
		":bike:" : "https://vozforums.com/images/smilies/Off/bike.gif",
		":mage:" : "https://vozforums.com/images/smilies/Off/mage.gif",
		":shit:" : "https://vozforums.com/images/smilies/emos/shit.gif",
		":runrun:" : "https://vozforums.com/images/smilies/Off/runrun.gif",
		":loveyou:" : "https://vozforums.com/images/smilies/emos/loveyou.gif",
		":lovemachine:" : "https://vozforums.com/images/smilies/emos/lovemachine.gif",
		":stupid:" : "https://vozforums.com/images/smilies/emos/stupid.gif",
		":doublegun:" : "https://vozforums.com/images/smilies/emos/doublegun.gif",
		":ban:" : "https://vozforums.com/images/smilies/Off/bann.gif",
		":please:" : "https://vozforums.com/images/smilies/Off/please.gif",
		":boom:" : "https://vozforums.com/images/smilies/emos/boom.gif",
		":lol:" : "https://vozforums.com/images/smilies/emos/lol.gif",
		":welcome:" : "https://vozforums.com/images/smilies/Off/welcome.gif",
		":puke:" : "https://vozforums.com/images/smilies/emos/puke.gif",
		":shoot1:" : "https://vozforums.com/images/smilies/emos/shoot1.gif",
		":no:" : "https://vozforums.com/images/smilies/emos/no.gif",
		":yes:" : "https://vozforums.com/images/smilies/emos/yes.gif",
		":birthday:" : "https://vozforums.com/images/smilies/emos/Birthday.gif",
		":winner:" : "https://vozforums.com/images/smilies/emos/winner.gif",
		":band:" : "https://vozforums.com/images/smilies/emos/band.gif",




	":((" :"/static/smile/20.gif",
	":))" :"/static/smile/21.gif",
	":)" :"/static/smile/1.gif",
	":(" :"/static/smile/2.gif",
	";)" :"/static/smile/3.gif",
	":D" :"/static/smile/4.gif",
	";;)" :"/static/smile/5.gif",
	//">:D<" :"/static/smile/6.gif",
	":-/" :"/static/smile/7.gif",
	":x" :"/static/smile/8.gif",
	":\">" :"/static/smile/9.gif",
	":P" :"/static/smile/10.gif",
	":-*" :"/static/smile/11.gif",
	"=((" :"/static/smile/12.gif",
	":-O" :"/static/smile/13.gif",
	"X(" :"/static/smile/14.gif",
	":>" :"/static/smile/15.gif",
	"B-)" :"/static/smile/16.gif",
	":-S" :"/static/smile/17.gif",
	"#:-S" :"/static/smile/18.gif",
	">:)" :"/static/smile/19.gif",

	":|" :"/static/smile/22.gif",
	"/:)" :"/static/smile/23.gif",
	"=))" :"/static/smile/24.gif",
	"O:-)" :"/static/smile/25.gif",
	":-B" :"/static/smile/26.gif",
	"=;" :"/static/smile/27.gif",
	":-c" :"/static/smile/28.gif",
	":)]" :"/static/smile/29.gif",
	"~X(" :"/static/smile/30.gif",
	":-h" :"/static/smile/31.gif",
	":-t" :"/static/smile/32.gif",
	"8->" :"/static/smile/33.gif",
	"I-)" :"/static/smile/34.gif",
	"8-|" :"/static/smile/35.gif",
	"L-)" :"/static/smile/36.gif",
	":-&" :"/static/smile/37.gif",
	":-$" :"/static/smile/38.gif",
	"[-(" :"/static/smile/39.gif",
	":O)" :"/static/smile/40.gif",
	"8-}" :"/static/smile/41.gif",
	"<:-P" :"/static/smile/42.gif",
	"(:|" :"/static/smile/43.gif",
	"=P~" :"/static/smile/44.gif",
	":-?" :"/static/smile/45.gif",
	"#-o" :"/static/smile/46.gif",
	"=D>" :"/static/smile/47.gif",
	":-SS" :"/static/smile/48.gif",
	"@-)" :"/static/smile/49.gif",
	":^o" :"/static/smile/50.gif",
	":-w" :"/static/smile/51.gif",
	":-<" :"/static/smile/52.gif",
	">:P" :"/static/smile/53.gif",
	"<):)" :"/static/smile/54.gif",
	"X_X" :"/static/smile/55.gif",
	":!!" :"/static/smile/56.gif",
	"\\m/" :"/static/smile/57.gif",
	":-q" :"/static/smile/58.gif",
	":-bd" :"/static/smile/59.gif",
	"^#(^" :"/static/smile/60.gif",
	":ar!" :"/static/smile/61.gif",
}
	Compiler = bbcode.NewCompiler(true, true) // autoCloseTags, ignoreUnmatchedClosingTags
}

func ReplaceSmile(content string) string {

	content = Compiler.Compile(content)
	for key, value := range Smiles {

		content = strings.Replace(content,key,"<img src='"+value+"'/>",-1);


	}
	return content;
}










































































































































































































































































































































































































































































































































































































































































































































































































































































































































































