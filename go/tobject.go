/*
 * it,s ar commentt
 * it,s pirate styl
 * mov' yah hedd
 * now i,m a piratee
 */

package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type tobject struct {
	Name          string    `json:"name"`
	Sign          string    `json:"sign"`
	ID            string    `json:"id"`
	Enforcer      string    `json:"enforcer"`
	Baseon        string    `json:"baseon"`
	EffectiveDate string    `json:"effective_date"`
	PassDate      string    `json:"pass_date"`
	ChapterArray  []chapter `json:"chapter_array"`
	ArticleArray  []article `json:"article_array"`
}

/* thi,s a functionn
 * ar functionnn to catch them,
 * arr, i,ll catch ,em now
 */
func catchit(legis []byte) *tobject {
	result := tobject{}

	/* I,ll catchh ,em weak firstr
	 * id, name, base,
	 * date, signnn
	 * all o' 'em
	 */
	re1 := regexp.MustCompile(idEnforcerRegx)
	matchIdEnforce := re1.FindSubmatch(legis)
	if len(matchIdEnforce) != 0 {
		result.ID = string(matchIdEnforce[2])
		result.Enforcer = string(matchIdEnforce[3]) + string(matchIdEnforce[4])
	} else {
		fmt.Println("WARNING: ID_ENFORCER_REGX no match")
	}

	re2 := regexp.MustCompile(nameRegx)
	matchName := re2.FindSubmatch(legis)
	if len(matchName) != 0 {
		result.Name = string(matchName[1]) + " " + string(matchName[2])
	} else {
		fmt.Println("WARNING: ID_NAME_REGX no match")
	}

	re3 := regexp.MustCompile(baseonRegx)
	matchBaseon := re3.FindSubmatch(legis)
	if len(matchBaseon) != 0 {
		result.Baseon += string(matchBaseon[0])
	} else {
		fmt.Println("WARNING: BASE_ON_REGX no match")
	}

	re4 := regexp.MustCompile(effectiveDateRegx)
	matchEffectiveDate := re4.FindAllSubmatch(legis, -1)
	i := len(matchEffectiveDate)
	if i != 0 {
		result.EffectiveDate = string(matchEffectiveDate[i-1][2]) + "-" +
			string(matchEffectiveDate[i-1][3]) + "-" + string(matchEffectiveDate[i-1][4])
	} else {
		fmt.Println("WARNING: EFFECTIVE_DATE_REGX no match")
	}

	re5 := regexp.MustCompile(passDateRegx)
	matchPassDate := re5.FindAllSubmatch(legis, -1)
	j := len(matchPassDate)
	if j != 0 {
		result.PassDate = string(matchPassDate[j-1][2]) + "-" +
			string(matchPassDate[j-1][3]) + "-" + string(matchPassDate[j-1][4])
	} else {
		fmt.Println("WARNING: PASS_DATE_REGX no match")
	}

	re6 := regexp.MustCompile(signRegx)
	matchSign := re6.FindAllSubmatch(legis, -1)
	k := len(matchSign)
	if k != 0 {
		result.Sign = string(matchSign[0][1])
		for t := 1; t < k; t++ {
			result.Sign += " / " + string(matchSign[t][1])
		}
	} else {
		fmt.Println("WARNING: SIGN_REGX no match")
	}

	/* ,em bosss comingg here, mov' your head
	 * i,ll divde ,em by chapter index,
	 * pushh that bastard to catchchap functionn! Arr!
	 * boil it, i say boild it,
	 */
	result.ChapterArray = []chapter{}
	result.ArticleArray = []article{}
	regx := regexp.MustCompile(chapterRegx)
	matchChapter := regx.FindAllSubmatch(legis, -1)
	matchChapterIndex := regx.FindAllSubmatchIndex(legis, -1)

	switch m := len(matchChapter); m {

	/*  i can,t catch anything!
	    so pity, chapter-regexship drowned!
	*/
	case 0:
		fmt.Println("WARNING: END_REGX no match")

		//  continue contribute ArticleArray with the first element [0] contain "regency"
		regencyElement := article{
			Header:  "regency",
			Content: "unidentify",
		}
		result.ArticleArray = append(result.ArticleArray, regencyElement)

		//got ,em all
		pushAllArticleArray(&result, legis)

	/*  only end signal,
	    parfect, no chapter no topic no problemm~
	*/
	case 1:
		//  contribute ArticleArray with the first element [0] contain "regency"
		regency := ""
		if len(matchChapter[0]) < 4 {
			fmt.Println("WARNING: CHAPTER_REGX run into a problem")
		} else {
			regency = string(matchChapter[0][3])
		}
		regencyElement := article{
			Header:  "regency",
			Content: regency,
		}
		result.ArticleArray = append(result.ArticleArray, regencyElement)

		//comee in
		pushAllArticleArray(&result, legis)

	/*  the time has come,
	    we need to fa,e this!
	*/
	default:
		//pick "regency" and put to first jail celll!
		regency := ""
		if len(matchChapter[m-1]) < 4 {
			fmt.Println("WARNING: CHAPTER_REGX run into a problem")
		} else {
			regency = string(matchChapter[m-1][3])
		}
		regencyElement := article{
			Header:  "regency",
			Content: regency,
		}
		result.ArticleArray = append(result.ArticleArray, regencyElement)

		/*  divvide by chapter, cut ,em
		    put pieces to the sausage factory creatChapter
		*/
		for i := 0; i < m-1; i++ {
			chapterContent := legis[matchChapterIndex[i][1] : matchChapterIndex[i+1][0]+1]
			chapterHeader := string(matchChapter[i][2])
			creatChapter(&result, chapterContent, chapterHeader)
		}

		/*	th,s final part, we handle it in our handd"!
			content same same legis[x:]
		*/
		finalChapterContent := legis[matchChapterIndex[m-1][1]:]
		test := tobject{}
		test1, test2 := pushArticleArray(&test, finalChapterContent)
		if test1 != -1 && test2 != -1 {
			finalChapterHeader := string(matchChapter[m-1][2])
			creatChapter(&result, finalChapterContent, finalChapterHeader)
		}
	}

	/*  alright, allllright, i likke to mov'it
	    gotta abandon ship!
	*/
	return &result
}

func creatChapter(target *tobject, contentPie []byte, header string) {
	regx := regexp.MustCompile(topicRegx)
	matchTopic := regx.FindAllSubmatch(contentPie, -1)

	switch n := len(matchTopic); n {

	/*  thank god,
	    no topic
	*/
	case 0:
		first, last := pushArticleArray(target, contentPie)
		temp := chapter{
			Header:       header,
			FirstArticle: first,
			LastArticle:  last,
			TopicArray:   []topic{},
		}
		target.ChapterArray = append(target.ChapterArray, temp)

	/*  include topics... */
	default:
		temp := chapter{
			Header:       header,
			FirstArticle: -1,
			LastArticle:  -1,
			TopicArray:   []topic{},
		}

		//  diwede them by topic, put them to creatTopic
		matchTopicIndex := regx.FindAllSubmatchIndex(contentPie, -1)
		for i := 0; i < n-1; i++ {
			topicContent := contentPie[matchTopicIndex[i][1] : matchTopicIndex[i+1][0]+1]
			topicHeader := string(matchTopic[i][3])
			creatTopic(target, &temp, topicContent, topicHeader)
		}
		finalTopicContent := contentPie[matchTopicIndex[n-1][1]:]
		finalTopicHeader := string(matchTopic[n-1][3])
		creatTopic(target, &temp, finalTopicContent, finalTopicHeader)
		temp.FirstArticle = temp.TopicArray[0].FirstArticle
		temp.LastArticle = temp.TopicArray[n-1].LastArticle

		//  imprison temp chapter!
		target.ChapterArray = append(target.ChapterArray, temp)
	}

}

/*  pushh articles to the jail from a []byteship ~
    bye ,em
    throw out the first and last
*/
func pushArticleArray(target *tobject, contentPie []byte) (int64, int64) {
	regx := regexp.MustCompile(articleRegx)
	match := regx.FindAllSubmatch(contentPie, -1)
	matchIndex := regx.FindAllSubmatchIndex(contentPie, -1)
	n := len(match)

	//  contentPie empty, fal, bak
	if n == 0 {
		fmt.Println("WARNING: ARTICLE_REGX run into trouble, but don't worry until second time :)")
		return -1, -1
	}

	/* (_''_) so cumplikate
	     (__''_) (_'') (''_)
	   ('')(_''_) (''_)
	*/

	/*  divide article, get content between them
	    catch 1, imprison 1
	*/
	for i := 0; i < n-1; i++ {
		content := strings.TrimSpace(string(contentPie[matchIndex[i][1] : matchIndex[i+1][0]+1]))
		temp := article{
			Header:  string(match[i][4]),
			Content: content,
		}
		target.ArticleArray = append(target.ArticleArray, temp)
	}

	//  we gotta handed the last by hand
	finalArticleContent := strings.TrimSpace(string(contentPie[matchIndex[n-1][1]+1:]))
	finalArticleHeader := string(match[n-1][4])
	finalArticle := article{
		Header:  finalArticleHeader,
		Content: finalArticleContent,
	}
	target.ArticleArray = append(target.ArticleArray, finalArticle)

	//  we call the firt and last to 'now their number
	a := string(match[0][3])
	first, err1 := strconv.ParseInt(a, 10, 64)
	if err1 != nil {
		panic(err1)
	}
	b := string(match[n-1][3])
	last, err2 := strconv.ParseInt(b, 10, 64)
	if err2 != nil {
		panic(err2)
	}
	return first, last
}

func pushAllArticleArray(target *tobject, contentPie []byte) {

	//  divide, put at full force!!
	regx := regexp.MustCompile(articleRegx)
	match := regx.FindAllSubmatch(contentPie, -1)
	matchIndex := regx.FindAllSubmatchIndex(contentPie, -1)
	n := len(match)
	if n < 2 {
		fmt.Println("WARNING: ARTICLE_REGX run into a problem on PushAll")
	}
	for i := 0; i < n-1; i++ {
		content := strings.TrimSpace(string(contentPie[matchIndex[i][1] : matchIndex[i+1][0]+1]))
		temp := article{
			Header:  string(match[i][4]),
			Content: content,
		}
		target.ArticleArray = append(target.ArticleArray, temp)
	}
}

func creatTopic(target *tobject, targetChapter *chapter, contentPie []byte, header string) {

	//  like dad, like son~
	first, last := pushArticleArray(target, contentPie)
	temp := topic{
		Header:       header,
		FirstArticle: first,
		LastArticle:  last,
	}
	targetChapter.TopicArray = append(targetChapter.TopicArray, temp)
}
