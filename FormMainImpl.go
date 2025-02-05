package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/imroc/req/v3"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/win"
	"govcl/work"
	"io"
	"log"
	"time"
)

// ::private::
type TFormMainFields struct {
	stopChan  chan struct{} // 添加停止信号通道
	isRunning bool
}

var Task *work.Worker

func iniTest() {

	iniFile := vcl.NewIniFile(".\\test.ini")
	defer iniFile.Free()

	iniFile.WriteBool("First", "Bool", true)
	iniFile.WriteString("First", "String", "这是字符串")
	iniFile.WriteDateTime("First", "Time", time.Now())
	iniFile.WriteInteger("First", "Integer", 123123)
	iniFile.WriteFloat("First", "Float", 123123)

	fmt.Println("Bool:", iniFile.ReadBool("First", "Bool", false))
	fmt.Println("String:", iniFile.ReadString("First", "String", ""))
	fmt.Println("Time:", iniFile.ReadDate("First", "Time", time.Now()))
	fmt.Println("Integer:", iniFile.ReadInteger("First", "Integer", 0))
	fmt.Println("Float:", iniFile.ReadFloat("First", "Float", 0.0))
}

func (f *TFormMain) OnFormCreate(sender vcl.IObject) {

	//rtl.SysOpen("http://www.baidu.com") // 打开链接
	//编译命令
	// go build -ldflags="-H=windowsgui" -tags tempdll

	// 异常捕获
	defer func() {
		err := recover() // 捕获panic
		if err != nil {
			fmt.Println("Exception: ", err)      // 打印异常信息
			vcl.ShowMessage(err.(error).Error()) // 显示异常信息对话框
		}
	}()
	iniTest()                       // 测试ini文件
	f.TForm.SetDoubleBuffered(true) // 双缓冲
	// 绑定事件
	f.SetOnClose(f.OnTFormClose) // 绑定窗体关闭事件

	// 托盘
	f.TrayIcon1.SetHint(f.TForm.Caption())
	f.TrayIcon1.SetIcon(vcl.Application.Icon())
	f.TrayIcon1.SetBalloonTitle("友情提示")  // 设置气球提示标题
	f.TrayIcon1.SetBalloonTimeout(10000) // 设置气球提示超时时间
	f.TrayIcon1.SetBalloonHint("欢迎使用")   // 设置气球提示内容
	f.TrayIcon1.ShowBalloonHint()

	// 绑定托盘点击事件
	f.TrayIcon1.SetOnClick(f.OnTrayIconClick)

	// 清空日志框内容
	f.Memo_log.Clear()

	// 拖放文件
	f.SetAllowDropFiles(true)

	//// 窗口大小约束
	//f.TForm.SetOnConstrainedResize(func(sender vcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32) { // 设置窗体大小约束事件
	//	*minWidth = 800 // 设置最小宽度
	//	*minHeight = 600 // 设置最小高度
	//	*maxWidth = 800 // 设置最大宽度
	//	*maxHeight = 600 // 设置最大高度
	//}

	f.TForm.SetOnDestroy(func(sender vcl.IObject) { // 设置窗体销毁事件
		fmt.Println("Form Destroy.") // 打印窗体销毁信息
	})
	// 设置窗体关闭查询事件
	f.TForm.SetOnCloseQuery(f.OnTFormCloseQuery)

	// 设置默认值
	imagePath := "F:\\所有照片\\钢铁侠.jpeg" // 可以考虑从配置文件中读取
	f.Image1.Picture().LoadFromFile(imagePath)
	f.stopChan = make(chan struct{}) // 初始化停止信号通道
}

// 设置窗体关闭查询事件
func (f *TFormMain) OnTFormCloseQuery(Sender vcl.IObject, CanClose *bool) {
	*CanClose = vcl.MessageDlg("是否退出?", types.MtInformation, types.MbYes, types.MbNo) == types.MrYes // 显示确认退出对话框
	fmt.Println("OnCloseQuery")
}

// OnFormDestroy 处理窗体销毁事件。
func (f *TFormMain) OnFormDestroy(sender vcl.IObject) {
	fmt.Println("OnFormDestroy")
}

// OnTFormClose  处理窗体关闭事件。
func (f *TFormMain) OnTFormClose(sender vcl.IObject, action *types.TCloseAction) {
	fmt.Println("OnFormClose")
	// 在这里添加其他关闭时需要执行的逻辑
}

func (f *TFormMain) OnFormDropFiles(sender vcl.IObject, aFileNames []string) {
	fmt.Println("当前拖放文件事件执行，文件数：", len(aFileNames))
	for i, s := range aFileNames {
		fmt.Println("index:", i, ", filename:", s)
	}
}

// OnMenuItem_tuiChuClick 处理退出菜单项的点击事件。
func (f *TFormMain) OnMenuItem_tuiChuClick(sender vcl.IObject) {
	f.Close()
}

func (f *TFormMain) OnButtonClick(sender vcl.IObject) {
	fmt.Println("OnButton1Click")
}

func (f *TFormMain) OnTabSheet1ContextPopup(sender vcl.IObject, mousePos types.TPoint, handled *bool) {
}

func (f *TFormMain) OnButton2Click(sender vcl.IObject) {
	var items []string
	for i := 0; i < 10; i++ {
		items = append(items, gconv.String(i))
	}

	vcl.ThreadSync(func() {
		for _, itemStr := range items {
			item := f.ListView1.Items().Add()
			item.SetCaption(itemStr)
			item.SubItems().Add("Alice")
			item.SubItems().Add("25")
			item.SetChecked(true)

			f.StatusBar1.Panels().Items(1).SetText(itemStr)
			f.Memo_log.Lines().Add(itemStr)
			f.ProgressBar1.SetPosition(gconv.Int32(itemStr))
		}
	})
}

func (f *TFormMain) OnScrollBox1Click(sender vcl.IObject) {
}

func (f *TFormMain) OnAction1Execute(sender vcl.IObject) {
	fmt.Println("OnAction1Execute")
}

func (f *TFormMain) OnMenuItem1Click(sender vcl.IObject) {
}

func (f *TFormMain) OnMenuItem_FuZhiMiMaClick(sender vcl.IObject) {
	selectedItem := f.ListView1.Selected()
	if selectedItem != nil {
		fmt.Println("选中状态:", selectedItem.Checked())
		fmt.Println("索引:", selectedItem.Index())
		firstColumnValue := selectedItem.SubItems().Strings(1)
		fmt.Println("值:", firstColumnValue)
	} else {
		fmt.Println("没有选中的项")
	}
}

func (f *TFormMain) OnMenuItem_FuZhiZhangHaoClick(sender vcl.IObject) {
	selectedItem := f.ListView1.Selected()
	if selectedItem != nil {
		fmt.Println("选中状态:", selectedItem.Checked())
		fmt.Println("索引:", selectedItem.Index())
		firstColumnValue := selectedItem.SubItems().Strings(0)
		fmt.Println("值:", firstColumnValue)
		vcl.Clipboard.SetAsText(firstColumnValue) // 复制到剪切板
	} else {
		fmt.Println("没有选中的项")
	}
}

func (f *TFormMain) OnButton_stopClick(sender vcl.IObject) {
	f.Button_start.SetEnabled(true)
	f.Button_stop.SetEnabled(false)

	if Task != nil {
		Task.T停止()
	}
}

func (f *TFormMain) OnButton_startClick(sender vcl.IObject) {

	go func() {
		// 设置按钮状态
		f.Button_start.SetEnabled(false)
		f.Button_stop.SetEnabled(true)
		defer f.Button_start.SetEnabled(true)
		defer f.Button_stop.SetEnabled(false)
		// 清空列表
		f.ListView1.Items().BeginUpdate()
		f.ListView1.Items().Clear()
		f.ListView1.Items().EndUpdate()
		f.ListView1.Refresh()

		// 停止旧的任务
		if Task != nil {
			Task.T停止()
		}

		// 创建新任务
		count := gconv.Int(f.EdtThreads.Text())
		if count <= 0 {
			vcl.ShowMessage("线程数必须大于0")
			return
		}
		Task = work.New(count)
		defer Task.T停止()
		err := Task.C创建(count)
		if err != nil {
			vcl.ShowMessage(err.Error())
			return
		}

		// 提交任务
		for i := 0; i < 1000; i++ {
			if Task.S是否关闭() {
				fmt.Println("任务已关闭，不再提交新任务")
				break
			}
			err := Task.T提交任务(f.Job(i))
			if err != nil {
				fmt.Println("Failed to submit task:", err)
			}
		}

		// 等待所有任务完成
		go func() {
			for {
				if Task.H获取剩余任务数() == 0 {
					break
				}
				time.Sleep(100 * time.Microsecond)
			}
		}()
	}()
}

func (f *TFormMain) Job(id int) func() {
	return func() {
		vcl.ThreadSync(func() {
			// 插入数据

			//f.ListView1.Items().BeginUpdate()
			//f.ListView1.Items().EndUpdate()
			item := f.ListView1.Items().Add()
			item.SetCaption(gconv.String(id))
			item.SubItems().Add("Alice")
			item.SubItems().Add("25")
			item.SetChecked(true)

		})
	}
}

func (f *TFormMain) OnButton_FangWenBaiDuClick(sender vcl.IObject) {
	// 创建一个自定义的 client,随机TLS
	client := req.C().SetTLSFingerprintRandomized()

	// 设置请求头
	headers := map[string]string{
		"Connection":                "keep-alive",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.95 Safari/537.36",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"Sec-Fetch-Site":            "none",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-User":            "?1",
		"Sec-Fetch-Dest":            "document",
		"sec-ch-ua":                 `"Chromium";v="122", "Not(A:Brand";v="24", "Google Chrome";v="122"`,
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        `"Windows"`,
		"Accept-Encoding":           "gzip, deflate, br",
		"Accept-Language":           "zh-CN,zh;q=0.9,en;q=0.8",
		"Cookie":                    "BD_UPN=12314753; BDUSS_BFESS=mhDV2JWUW5ZekdENHhhVmZNd3RVRTZ6STlhQzY5Um5iSld6aHFPcnZuSGdnRVJuSVFBQUFBJCQAAAAAAAAAAAEAAADNk3A0c3N4cHZpY3AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAODzHGfg8xxne; PSTM=1732688015; BIDUPSID=87BD4ACEB24A19BB0CFBCD387F82A4CA; BAIDUID=C8692A2B8D2D2E26DFCA98A09751CEC5:FG=1; BAIDUID_BFESS=C8692A2B8D2D2E26DFCA98A09751CEC5:FG=1; ZFY=vQBVGQNTqR5JAQ406xk8YqdEiH9A4vchI4E:ASyjCqRo:C; Hm_lvt_aec699bb6442ba076c8981c6dc490771=1734506638; H_WISE_SIDS_BFESS=60276_61027_61217_60851_61368_61390_61421_61427_61462; B64_BOT=1; baikeVisitId=24ed39d5-2394-45a2-acf5-10dc200d0856; COOKIE_SESSION=6637_3_4_8_2_11_1_0_4_3_0_12_362_126875_0_0_1734534205_1734660910_1735027836%7C9%23127096_42_1734660907%7C9; H_PS_PSSID=60276_61027_61217_61390_61427_61462_60853_61430_61508_61524_61521_61568_61361",
	}
	// 发送请求
	resp, err := client.R().
		SetHeaders(headers).
		Get("https://www.baidu.com/")
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	//g.Dump(resp.GetHeader("Content-Encoding"))
	//g.Dump(resp.HeaderToString())
	fmt.Println("响应状态码:", resp.StatusCode)
	//fmt.Println("响应体:", resp.String())

	// 获取原始的 GZIP 压缩数据
	gzipData := resp.Bytes() // 这会返回响应体的原始字节数据（包括 GZIP 压缩内容）

	// 创建一个 bytes.Reader 来读取 GZIP 数据
	gzipReader, err := gzip.NewReader(bytes.NewReader(gzipData))
	if err != nil {
		log.Fatal(err)
	}
	defer gzipReader.Close()

	// 解压数据
	var out bytes.Buffer
	_, err = io.Copy(&out, gzipReader)
	if err != nil {
		log.Fatal(err)
	}

	// 打印解压后的内容
	//fmt.Println(resp.String())

	f.Memo_log.Lines().Add(out.String())
	// 滚动到底部
	f.Memo_log.ScrollBy(0, f.Memo_log.Lines().Count()*f.Memo_log.Font().Height())

	vcl.Application.MessageBox("访问完成", "友情提示", win.MB_YESNO+win.MB_ICONINFORMATION) // 显示消息框
}

func (f *TFormMain) GMM() {
	client := req.C().SetTLSFingerprintRandomized()

	headers := map[string]string{
		"Host":               "www.gmmsj.com",
		"Connection":         "keep-alive",
		"sec-ch-ua":          `"Chromium";v="122", "Not(A:Brand";v="24", "Google Chrome";v="122"`,
		"Accept":             "application/json, text/javascript, */*; q=0.01",
		"X-Requested-With":   "XMLHttpRequest",
		"sec-ch-ua-mobile":   "?0",
		"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko",
		"sec-ch-ua-platform": `"Windows"`,
		"Sec-Fetch-Site":     "same-origin",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Dest":     "empty",
		"Accept-Language":    "zh-CN,zh;q=0.9,en;q=0.8",
		"Cookie":             "Hm_lvt_672bace42218fef7cd167f5a48004182=1709044108; Hm_lvt_d6e45b11541f36a1010c54892aa215c2=1735299477; deviceId=v2_dnAqgBwNp0EY1oPCSVq8CNxBdSXOwyPm; device_id=v2_dnAqgBwNp0EY1oPCSVq8CNxBdSXOwyPm; https_waf_cookie=a1562c73-e030-4384e5e30e328c6d493f3c2a309f5857ad53; autoShowQcode=1; GMMSESSID=ad8fdecf49d94b76b9777ef25e93c478; PHPSESSID=11810e2afade4639a5d89e523dc94597; dropDownGame_his=%5B%7B%22game_id%22%3A4%2C%22game_name%22%3A%22%u70ED%u8840%u4F20%u5947%22%7D%5D; JSESSIONID=node017fmvxr6sj32ovqpywc20hb5u347905.node0; displayAccount=; b_uid=; userMid=; nickName=; gmmpc_loginstatus=0",
	}

	url := "https://www.gmmsj.com/gatew/gmmGoodsGW/goodsListV2?app_version=1.0.0.269236&device_id=v2_dnAqgBwNp0EY1oPCSVq8CNxBdSXOwyPm&system_deviceId=v2_dnAqgBwNp0EY1oPCSVq8CNxBdSXOwyPm&app_channel=chrome&src_code=7&keyword=&order_type=total_price&order_dir=a&searchProperties=%7B%22p_level%22%3A%22%5B%5C%2246%5C%22%2C%5C%22%5C%22%5D%22%2C%22p_job%22%3A%22%5B%5C%222%5C%22%5D%22%7D&goods_types=10&game_id=4&area_id_groups=%5B%7B%22area_id%22%3A213%7D%5D&page=1&limit=15&safe_type="

	// 发送请求
	resp, err := client.R().
		SetHeaders(headers).
		Get(url)
	if err != nil {
		log.Printf("请求失败: %v", err)
		return
	}
	defer resp.Body.Close()

	// 获取响应状态码
	log.Printf("响应状态码: %d", resp.StatusCode)

	// 获取响应体
	responseBody, err := resp.ToString()
	if err != nil {
		log.Printf("读取响应体失败: %v", err)
		return
	}
	log.Printf("响应体: %s", responseBody)

	bodyArray := []GmmStruct{}
	j := gjson.New(responseBody)
	err = j.Get("data.goodsList").Scan(&bodyArray)
	if err != nil {
		log.Printf("gjson失败: %v\n响应体: %s", err, responseBody)
		return
	}
	g.Dump(bodyArray)

	// 清空列表
	go func() {
		vcl.ThreadSync(func() {
			f.updateListViewItems(bodyArray)
		})
	}()
}

type GmmStruct struct {
	ExtraHint          string `json:"extra_hint"`
	TransactionMode    int    `json:"transaction_mode"`
	Price              string `json:"price"`
	SafeTypeData       string `json:"safe_type_data"`
	BookId             string `json:"book_id"`
	ExtGoodsType       int    `json:"ext_goods_type"`
	SellerNickname     string `json:"seller_nickname"`
	ModifyTime         string `json:"modify_time"`
	RaceCampName       string `json:"race_camp_name"`
	Thumbnail          string `json:"thumbnail"`
	GoodsTypeName      string `json:"goods_type_name"`
	DeliverTime        string `json:"deliver_time"`
	GoodsType          int    `json:"goods_type"`
	TradeMode          int    `json:"trade_mode"`
	BestProportion     int    `json:"best_proportion"`
	CsScreenShot       int    `json:"cs_screen_shot"`
	IsAreaTrans        int    `json:"is_area_trans"`
	AvailQty           int    `json:"avail_qty"`
	Quantity           int    `json:"quantity"`
	SellerCreditType   string `json:"seller_credit_type"`
	CheckType          int    `json:"check_type"`
	GoodsListTitle     string `json:"goods_list_title"`
	GoodsListSubTitle  string `json:"goods_list_sub_title"`
	ExtraHint2         string `json:"extra_hint2"`
	AutoscreenshotFlag int    `json:"autoscreenshot_flag"`
	ShopUser           int    `json:"shop_user"`
	IsQuick            int    `json:"is_quick"`
}

func (f *TFormMain) updateListViewItems(bodyArray []GmmStruct) {
	go func() {
		vcl.ThreadSync(func() {
			f.ListView1.Items().BeginUpdate()
			defer f.ListView1.Items().EndUpdate()

			f.ListView1.Items().Clear()
			for i, itemData := range bodyArray {
				item := f.addListViewItem(itemData, i+1)
				f.checkPriceAndNotify(itemData, item)
			}
		})
	}()
}

func (f *TFormMain) addListViewItem(itemData GmmStruct, index int) *vcl.TListItem {
	item := f.ListView1.Items().Add()
	item.SetCaption(gconv.String(index))
	item.SubItems().Add(fmt.Sprintf("https://www.gmmsj.com/dy/4_zh/detail_%s.shtml", itemData.BookId))
	item.SubItems().Add(itemData.Price)
	item.SubItems().Add(itemData.BookId)
	item.SubItems().Add(itemData.GoodsListTitle)
	item.SetChecked(true)
	return item
}

func (f *TFormMain) checkPriceAndNotify(itemData GmmStruct, item *vcl.TListItem) {
	if gconv.Int(itemData.Price) <= 400 && gconv.Int(itemData.Price) >= 150 {
		f.showPriceNotification(itemData.GoodsListTitle)
	}
}

func (f *TFormMain) showPriceNotification(title string) {
	vcl.Application.MessageBox(fmt.Sprintf("有低价了[%s]", title), "友情提示", win.MB_YESNO+win.MB_ICONINFORMATION)
}

func (f *TFormMain) OnButton_GMMClick(sender vcl.IObject) {
	if f.isRunning {
		f.isRunning = false
		f.Button_GMM.SetCaption("刷新G买卖")
		close(f.stopChan) // 关闭停止信号通道
	} else {
		f.isRunning = true
		f.Button_GMM.SetCaption("停止")
		f.stopChan = make(chan struct{}) // 重新初始化停止信号通道
		go func() {
			ticker := time.NewTicker(30 * time.Second)
			defer ticker.Stop()
			go f.GMM() // 先执行刷新操作
			for {
				select {
				case <-ticker.C:
					go f.GMM() // 执行刷新操作，放入goroutine中
				case <-f.stopChan:
					return
				}
			}
		}()
	}
}

// OnTrayIconClick 处理托盘图标点击事件。
func (f *TFormMain) OnTrayIconClick(sender vcl.IObject) {
	if f.TForm.Visible() {
		f.TForm.Hide()
	} else {
		f.TForm.Show()
		f.TForm.BringToFront()
	}
}
