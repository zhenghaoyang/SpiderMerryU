package parser

import (
	"SpiderMerryU/engine"
	"SpiderMerryU/img"
	"SpiderMerryU/model"
	"regexp"
	"strconv"
)

//<span class="content_top_name">栗子</span>
//<span class="content_top_info">30岁&nbsp;&nbsp;·&nbsp;&nbsp;
//<span class="content_top_info">29岁&nbsp;&nbsp;·&nbsp;&nbsp;165cm&nbsp;&nbsp;·&nbsp;&nbsp;10W以下&nbsp;&nbsp;·&nbsp;&nbsp;天蝎座</span>
//<span class="content_top_abode">山东-青岛</span>
//<p class="content_main_info">民族<span>彝族</span></p>
//<p class="content_main_info_right">户籍<span>云南-昆明</span></p>
//<p class="content_main_info">婚姻状况<span>从未结婚</span></p>
//<p class="content_main_info_right">职业<span>未填写</span></p>
//<p class="content_main_info">信仰<span>其他</span></p>
//<p class="content_main_info_right">饮酒<span>从不</span></p>
//<p class="content_main_info">吸烟<span>从不</span></p>
//<p class="content_main_info_right">家中排行<span>一</span></p>
//<p class="content_main_info">有无子女<span>没有</span></p>
//<p class="content_main_info_right">体重<span>50kg</span></p>
//<p title="双眼皮">
// <p class="content_top_about_me selected" title='点击查看全部'>自我介绍：祖籍福建。在澳門出生長大。居住澳門和珠海。從事餐飲行業。不抽煙不喝酒</p>
//<title>29岁未婚男
var (
	//在外边编译性能好些
	nameRe   = regexp.MustCompile(` <span class="content_top_name">([^>]+)</span>`)
	genderRe = regexp.MustCompile(` <title>([\d]+)岁未婚女`)
	heightRe = regexp.MustCompile(`&nbsp;&nbsp;·&nbsp;&nbsp;([\d]+)cm&nbsp;&nbsp;·&nbsp;&nbsp;`)
	weightRe = regexp.MustCompile(`<p class="content_main_info_right">体重<span>([\d]+)kg</span></p>`)
	incomeRe = regexp.MustCompile(`cm&nbsp;&nbsp;·&nbsp;&nbsp;([^;]+)&nbsp;&nbsp;·&nbsp;&nbsp;`)

	hujiRe       = regexp.MustCompile(`<p class="content_main_info_right">户籍<span>([^>]+)</span></p>`)
	minzuRe      = regexp.MustCompile(`<p class="content_main_info">民族<span>([^>]+)</span></p>`)
	occupationRe = regexp.MustCompile(`<p class="content_main_info_right">职业<span>([^>]+)</span></p>`)

	believeRe  = regexp.MustCompile(`<p class="content_main_info">信仰<span>([^>]+)</span></p>`)
	marriageRe = regexp.MustCompile(`<p class="content_main_info">婚姻状况<span>([^>]+)</span></p>`)
	//ageRe      = regexp.MustCompile(`<span class="content_top_info">([\d]+)岁`)
	ageRe     = regexp.MustCompile(` <title>([\d]+)岁未婚女`)
	aboutmeRe = regexp.MustCompile(`<p class="content_top_about_me selected" title='点击查看全部'>([^>]+)</p>`)
	tgaRe     = regexp.MustCompile(`<p title="([^"]+)">`)
	//https://www.51marryyou.com/user/([0-9a-z]+[^.html]).html
	idUrlRe = regexp.MustCompile(`https://www.51marryyou.com/user/([0-9a-z]+[^.html]).html`)
	//<a href="../user/96ac8504151f6eec019c42226f2452d4.html" target="_blank"
	guessRe = regexp.MustCompile(`<a href="..(/user/[0-9a-z]+[^.html]).html" target="_blank" `)

	//"https://img1.51marryyou.com/2018-06-04/efd36b88ebfbf969222b04afdfd8833e.jpg@450w_450h_1e_1c.jpg"
	//<img class="big_img" src="
	// <img class="big_img" src="https://img1.51marryyou.com/2018-07-23/73d8dc3beb29c3b8d06511415820fd03.jpg@450w_450h_1e_1c" alt="

)

func ParserProfile(contents []byte, url string) engine.ParseResult {

	cityprefix := "https://www.51marryyou.com"

	profile := model.Profile{}
	agemusthave := extractString(contents, ageRe)
	if agemusthave == "" {
		return engine.ParseResult{}
	}

	age, err := strconv.Atoi(agemusthave)
	if age > 25 {
		return engine.ParseResult{}
	}
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Name = extractString(contents, nameRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Huji = extractString(contents, hujiRe)
	profile.Aboutme = extractString(contents, aboutmeRe)
	profile.Tag = extractString(contents, tgaRe)
	idUrl := extractString([]byte(url), idUrlRe)
	// <img class="big_img" src="https://app.51marryyou.com/Public/image/woman.png 异常情况
	imgRe, err := regexp.Compile(`<img class="big_img" src="([^"]*)" alt="`)
	//todo 多张小图片
	//smallRe, err := regexp.Compile(`data-src="([^"]*)" src=`)
	//profile.SmallUrl = extractString(contents, smallRe) + ".jpg"
	//go img.Dowmload(profile.ImgUrl)
	//fmt.Println("profile.SmallUrl = ", profile.SmallUrl)
	//img.Dowmload(profile.SmallUrl)

	//profile.SmallUrl =
	//<img class="small_img" data-src="https://img1.51marryyou.com/2016-10-19/93a1fe4e25d2ab90b8a16a97cab904f6.jpg@450w_450h_1e_1c" src="https://img1.51marryyou.com/2016-10-19/93a1fe4e25d2ab90b8a16a97cab904f6.jpg@140w_140h_1e_1c" alt="">
	if err != nil {
		return engine.ParseResult{}
	}

	if profile.ImgUrl == "https://app.51marryyou.com/Public/image/woman.png" {
		return engine.ParseResult{}
	}
	profile.ImgUrl = extractString(contents, imgRe) + ".jpg"

	//https://img1.51marryyou.com/2018-07-23/73d8dc3beb29c3b8d06511415820fd03.jpg@450w_450h_1e_1c
	if profile.ImgUrl != "" && profile.ImgUrl != "@450w_450h_1e_1c.jpg" {
		//fmt.Println("profile.ImgUrl==", profile.ImgUrl)
		reg, err := regexp.Compile(`(\w|\d|_)*.jpg`)
		name := reg.FindStringSubmatch(profile.ImgUrl)[0]
		//fmt.Printf("图片名字=%s\n", name)
		if err == nil && profile.ImgUrl != ".jpg" {
			profile.ImgName = name
			go func() {
				img.Dowmload(profile.ImgUrl)
			}()
		}
	} else {
		profile.ImgName = "woman.png"
	}
	result := engine.ParseResult{
		Items: []engine.Item{ //这一层是数组
			{
				Url:     url,
				Type:    "user",
				Id:      idUrl,
				Payload: profile,
			},
		},
	}

	//当前页面的相关用户 猜你喜欢
	citymatchs := guessRe.FindAllSubmatch(contents, -1)
	for _, submatch := range citymatchs {
		//下一级Request
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        cityprefix + string(submatch[1]), //城市URL
				ParserFunc: ParserProfile,
			},
		)
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
