package report

// Stat 数据报表
type Stat struct {
	// CampaignID 广告计划ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignName 广告计划名称
	CampaignName string `json:"campaign_name,omitempty"`
	// UnitID 广告组ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// UnitName 广告组名称
	UnitName string `json:"unit_name,omitempty"`
	// CreativeID 广告创意ID
	CreativeID uint64 `json:"creative_id,omitempty"`
	// CreativeName 广告创意名称
	CreativeName string `json:"creative_name,omitempty"`
	// PhotoID 视频id
	PhotoID string `json:"photo_id,omitempty"`
	// PhotoUrl 视频链接
	PhotoUrl string `json:"photo_url,omitempty"`
	// ImageToken 封面id
	ImageToken string `json:"image_token,omitempty"`
	// CoverUrl 封面链接
	CoverUrl string `json:"cover_url,omitempty"`
	// Description 作品广告语
	Description string `json:"description,omitempty"`
	// StickyUrl 便利贴url地址（便利贴报表时返回）
	StickyUrl string `json:"sticky_url,omitempty"`
	// ImageUrls 图片url地址（图片报表时返回）
	ImageUrls string `json:"image_urls,omitempty"`
	// AdScene 广告投放场景
	AdScene string `json:"ad_scene,omitempty"`
	// Province 省份
	Province string `json:"province,omitempty"`
	// City 城市
	City string `json:"city,omitempty"`
	// Gender 性别
	Gender string `json:"gender,omitempty"`
	// AgeSegment 年龄
	AgeSegment string `json:"age_segment,omitempty"`
	// Client 系统
	Client string `json:"client,omitempty"`
	// BusinessInterestTags 商业兴趣
	BusinessInterestTags string `json:"business_interest_tags,omitempty"`
	// Status 1 - 投放中；2 - 已暂停；3 - 已删除
	Status int `json:"status,omitempty"`
	// StatDate 数据日期，格式：YYYY-MM-DD
	StatDate string `json:"stat_date,omitempty"`
	// StatHour 数据小时
	StatHour int `json:"stat_hour,omitempty"`
	// Charge 花费（元）
	Charge float64 `json:"charge,omitempty"`
	// Show 封面曝光数
	Show int64 `json:"show,omitempty"`
	// PhotoClick 封面点击数
	PhotoClick int64 `json:"photo_click,omitempty"`
	// Aclick 素材曝光数
	Aclick int64 `json:"aclick,omitempty"`
	// Bclick 行为数
	Bclick int64 `json:"bclick,omitempty"`
	// PhotoClickRatio 封面点击率
	PhotoClickRatio float64 `json:"photo_click_ratio,omitempty"`
	// Play3sRatio 3s播放率
	Play3sRatio float64 `json:"play_3s_ratio,omitempty"`
	// ActionRatio 行为率
	ActionRatio float64 `json:"action_ratio,omitempty"`
	// Impression1kCost 平均千次曝光花费（元）
	Impression1kCost float64 `json:"impression_1k_cost,omitempty"`
	// PhotoClickCost 平均点击单价（元）
	PhotoClickCost float64 `json:"photo_click_cost,omitempty"`
	// ActionCost 平均行为单价（元）
	ActionCost float64 `json:"action_cost,omitempty"`
	// Share 分享数
	Share int64 `json:"share,omitempty"`
	// Comment 评论数
	Comment int64 `json:"comment,omitempty"`
	// Like 点赞数
	Like int64 `json:"like,omitempty"`
	// Follow 新增关注数
	Follow int64 `json:"follow,omitempty"`
	// CancelFollow 取消关注数
	CancelFollow int64 `json:"cancel_follow,omitempty"`
	// Report 举报数
	Report int64 `json:"report,omitempty"`
	// Block 拉黑数
	Block int64 `json:"block,omitempty"`
	// Negative 减少此类作品数
	Negative int64 `json:"negative,omitempty"`
	// Submit 提交按钮点击数（历史字段，同下方“表单提交数”，预计年底删除该字段）
	Submit int64 `json:"submit,omitempty"`
	// DownloadStarted 应用下载数据-安卓下载开始数
	DownloadStarted int64 `json:"download_started,omitempty"`
	// DownloadCompleted 应用下载数据-安卓下载完成数
	DownloadCompleted int64 `json:"download_completed,omitempty"`
	// Activation 应用下载数据-激活数
	Activation int64 `json:"activation,omitempty"`
	// EventPayFirstDay 应用下载数据-首日付费次数
	EventPayFirstDay int64 `json:"event_pay_first_day,omitempty"`
	// EventPayPurchaseAmountFirstDay 应用下载数据-首日付费金额
	EventPayPurchaseAmountFirstDay float64 `json:"event_pay_purchase_amount_first_day,omitempty"`
	// EventPayFirstDayRoi 应用下载数据-首日ROI
	EventPayFirstDayRoi float64 `json:"event_pay_first_day_roi,omitempty"`
	// EventPay 应用下载数据-付费次数
	EventPay int64 `json:"event_pay,omitempty"`
	// EventPayPurchaseAmount 应用下载数据-付费金额
	EventPayPurchaseAmount float64 `json:"event_pay_purchase_amount,omitempty"`
	// EventPayRoi 应用下载数据-ROI
	EventPayRoi float64 `json:"event_pay_roi,omitempty"`
	// EventRegister 应用下载数据-注册数
	EventRegister int64 `json:"event_register,omitempty"`
	// EventRegisterCost 应用下载数据-注册成本
	EventRegisterCost float64 `json:"event_register_cost,omitempty"`
	// EventRegisterRatio 应用下载数据-注册率
	EventRegisterRatio float64 `json:"event_register_ratio,omitempty"`
	// EventJinJianApp 应用下载数据-完件数
	EventJinJianApp int64 `json:"event_jin_jian_app,omitempty"`
	// EventJinJianAppCost 应用下载数据-完件成本
	EventJinJianAppCost float64 `json:"event_jin_jian_app_cost,omitempty"`
	// EventCreditGrantApp 应用下载数据-授信数
	EventCreditGrantApp int64 `json:"event_credit_grant_app,omitempty"`
	// EventCreditGrantAppCost 应用下载数据-授信成本
	EventCreditGrantAppCost float64 `json:"event_credit_grant_app_cost,omitempty"`
	// EventCreditGrantAppRatio 应用下载数据-授信率
	EventCreditGrantAppRatio float64 `json:"event_credit_grant_app_ratio,omitempty"`
	// EventOrderPaid 应用下载数据-付款成功数
	EventOrderPaid int64 `json:"event_order_paid,omitempty"`
	// EventOrderPaidPurchaseAmount 应用下载数据-付款成功金额
	EventOrderPaidPurchaseAmount float64 `json:"event_order_paid_purchase_amount,omitempty"`
	// EventOrderPaidCost 应用下载数据-单次付款成本
	EventOrderPaidCost float64 `json:"event_order_paid_cost,omitempty"`
	// EventNextDayStay 应用下载数据-次留数（仅支持分日查询）
	EventNextDayStay int64 `json:"event_next_day_stay,omitempty"`
	// EventNextDayStayCost 应用下载数据-次留成本（仅支持分日查询）
	EventNextDayStayCost float64 `json:"event_next_day_stay_cost,omitempty"`
	// EventNextDayStayRatio 应用下载数据-次留率（仅支持分日查询）
	EventNextDayStayRatio float64 `json:"event_next_day_stay_ratio,omitempty"`
	// FormCount 落地页数据-表单提交数
	FormCount int64 `json:"form_count,omitempty"`
	// FormCost 落地页数据-表单提交单价
	FormCost float64 `json:"form_cost,omitempty"`
	// FormActionRatio 落地页数据-表单提交点击率
	FormActionRatio float64 `json:"form_action_ratio,omitempty"`
	// EventJinJianLandingPage 落地页数据-落地页完件数
	EventJinJianLandingPage int64 `json:"event_jin_jian_landing_page,omitempty"`
	// EventJinJianLandingPageCost 落地页数据-落地页完件成本
	EventJinJianLandingPageCost float64 `json:"event_jin_jian_landing_page_cost,omitempty"`
	// EventCreditGrantLandingPage 落地页数据-落地页授信数
	EventCreditGrantLandingPage int64 `json:"event_credit_grant_landing_page,omitempty"`
	// EventCreditGrantLandingPageCost 落地页数据-落地页授信成本
	EventCreditGrantLandingPageCost float64 `json:"event_credit_grant_landing_page_cost,omitempty"`
	// EventCreditGrantLandingPageRatio 落地页数据-落地页授信率
	EventCreditGrantLandingPageRatio float64 `json:"event_credit_grant_landing_page_ratio,omitempty"`
	// EventValidClues 落地页数据-有效线索数
	EventValidClues int64 `json:"event_valid_clues,omitempty"`
	// EventValidCluesCost 落地页数据-有效线索成本
	EventValidCluesCost float64 `json:"event_valid_clues_cost,omitempty"`
	// AdProductCnt 商品成交数
	AdProductCnt int64 `json:"ad_product_cnt,omitempty"`
	// EventGoodsView 商品访问数
	EventGoodsView int64 `json:"event_goods_view,omitempty"`
	// MerchantRecoFans 涨粉量
	MerchantRecoFans int64 `json:"merchant_reco_fans,omitempty"`
	// EventOrderAmountRoi 涨粉量
	EventOrderAmountRoi float64 `json:"event_order_amount_roi,omitempty"`
	// EventGoodsViewCost 商品访问成本
	EventGoodsViewCost float64 `json:"event_goods_view_cost,omitempty"`
	// MerchantRecoFansCost 涨粉成本
	MerchantRecoFansCost float64 `json:"merchant_reco_fans_cost,omitempty"`
	// EventNewUserPay 新增付费人数
	EventNewUserPay int64 `json:"event_new_user_pay,omitempty"`
	// EventNewUserPayCost 新增付费人数成本：当日消耗 / 当日新增付费人数
	EventNewUserPayCost float64 `json:"event_new_user_pay_cost,omitempty"`
	// EventNewUserPayRatio 新增付费人数率：当日新增付费人数 / 当日行为数
	EventNewUserPayRatio float64 `json:"event_new_user_pay_ratio,omitempty"`
	// EventButtonClick 按钮点击数
	EventButtonClick int64 `json:"event_button_click,omitempty"`
	// EventButtonClickCost 按钮点击成本：当日消耗 / 按钮点击数
	EventButtonClickCost float64 `json:"event_button_click_cost,omitempty"`
	// EventButtonClickRatio 按钮点击成本：当日消耗 / 按钮点击数
	EventButtonClickRatio float64 `json:"event_button_click_ratio,omitempty"`
	// Play5sRatio 5s播放率
	Play5sRatio float64 `json:"play_5s_ratio,omitempty"`
	// PlanEndRatio 完播率
	PlanEndRatio float64 `json:"play_end_ratio,omitempty"`
	// EventOrderPaidRoi 订单支付率：订单支付数/商品访问数
	EventOrderPaidRoi float64 `json:"event_order_paid_roi,omitempty"`
	// EventNewUserJinjianApp 新增完件人数
	EventNewUserJinjianApp int64 `json:"event_new_user_jinjian_app,omitempty"`
	// EventNewUserJinjianAppCost 新增完件人数成本：消耗/新增完件人数
	EventNewUserJinjianAppCost float64 `json:"event_new_user_jinjian_app_cost,omitempty"`
	// EventNewUserJinjianAppRoi 新增完件人数率：新增完件人数/（转化+表单提交）
	EventNewUserJinjianAppRoi float64 `json:"event_new_user_jinjian_app_roi,omitempty"`
	// EventNewUserCreditGrantApp 新增授信人数
	EventNewUserCreditGrantApp int64 `json:"event_new_user_credit_grant_app,omitempty"`
	// EventNewUserCreditGrantAppCost 新增授信人数成本：消耗/新增授信人数
	EventNewUserCreditGrantAppCost float64 `json:"event_new_user_credit_grant_app_cost,omitempty"`
	// EventNewUserCreditGrantAppRoi 新增授信人数率：新增授信人数/（转化+表单提交）
	EventNewUserCreditGrantAppRoi float64 `json:"event_new_user_credit_grant_app_roi,omitempty"`
	// EventNewUserJinjianPage 落地页新增完件人数
	EventNewUserJinjianPage int64 `json:"event_new_user_jinjian_page,omitempty"`
	// EventNewUserJinjianPageCost 落地页新增完件人数成本
	EventNewUserJinjianPageCost float64 `json:"event_new_user_jinjian_page_cost,omitempty"`
	// EventNewUserJinjianPageRoi 落地页新增完件人数率
	EventNewUserJinjianPageRoi float64 `json:"event_new_user_jinjian_page_roi,omitempty"`
	// EventNewUserCreditGrantPage 落地页新增授信人数
	EventNewUserCreditGrantPage int64 `json:"event_new_user_credit_grant_page,omitempty"`
	// EventNewUserCreditGrantPageCost 落地页新增授信人数成本
	EventNewUserCreditGrantPageCost float64 `json:"event_new_user_credit_grant_page_cost,omitempty"`
	// EventNewUserCreditGrantPageRoi 落地页新增授信人数率
	EventNewUserCreditGrantPageRoi float64 `json:"event_new_user_credit_grant_page_roi,omitempty"`
	// EventAddWechat 微信复制数
	EventAddWechat int64 `json:"event_add_wechat,omitempty"`
	// EventAddWechatCost 微信复制成本
	EventAddWechatCost float64 `json:"event_add_wechat_cost,omitempty"`
	// EventAddWechatRatio 微信复制率
	EventAddWechatRatio float64 `json:"event_add_wechat_ratio,omitempty"`
	// EventGetThrough 智能电话-确认接通数
	EventGetThrough int64 `json:"event_get_through,omitempty"`
	// EventGetThroughCost 智能电话-确认接通成本
	EventGetThroughCost float64 `json:"event_get_through_cost,omitempty"`
	// EventGetThroughRatio 智能电话-确认接通率
	EventGetThroughRatio float64 `json:"event_get_through_ratio,omitempty"`
	// AdScene2 .
	AdScene2 string `json:"adScene,omitempty"`
	// PlacementType 广告范围
	PlacementType string `json:"placement_type,omitempty"`
	// Click1KCost 平均千次素材曝光花费(元)
	Click1KCost                      float64 `json:"click_1k_cost,omitempty"`
	EventCreditGrantLandingRatio     float64 `json:"event_credit_grant_landing_ratio,omitempty"`
	EventAppointForm                 int     `json:"event_appoint_form,omitempty"`
	EventAppointFormCost             float64 `json:"event_appoint_form_cost,omitempty"`
	EventAppointFormRatio            float64 `json:"event_appoint_form_ratio,omitempty"`
	EventAppointJumpClick            int     `json:"event_appoint_jump_click,omitempty"`
	EventAppointJumpClickCost        float64 `json:"event_appoint_jump_click_cost,omitempty"`
	EventAppointJumpClickRatio       float64 `json:"event_appoint_jump_click_ratio,omitempty"`
	UnionEventPayPurchaseAmount7D    float64 `json:"union_event_pay_purchase_amount_7d,omitempty"`
	UnionEventPayPurchaseAmount7DRoi float64 `json:"union_event_pay_purchase_amount_7d_roi,omitempty"`
	EventDspGiftForm                 int     `json:"event_dsp_gift_form,omitempty"`
	EventAppInvoked                  int     `json:"event_app_invoked,omitempty"`
	EventAppInvokedCost              float64 `json:"event_app_invoked_cost,omitempty"`
	EventAppInvokedRatio             float64 `json:"event_app_invoked_ratio,omitempty"`
	EventMultiConversion             int     `json:"event_multi_conversion,omitempty"`
	EventMultiConversionRatio        float64 `json:"event_multi_conversion_ratio,omitempty"`
	EventMultiConversionCost         float64 `json:"event_multi_conversion_cost,omitempty"`
	EventWatchAppAd                  int     `json:"event_watch_app_ad,omitempty"`
	EventAdWatchTimes                int     `json:"event_ad_watch_times,omitempty"`
	EventAdWatchTimesRatio           float64 `json:"event_ad_watch_times_ratio,omitempty"`
	EventAdWatchTimesCost            float64 `json:"event_ad_watch_times_cost,omitempty"`
	EventAddShoppingCart             int     `json:"event_add_shopping_cart,omitempty"`
	EventAddShoppingCartCost         float64 `json:"event_add_shopping_cart_cost,omitempty"`
	AdPhotoPlayed75PercentRatio      float64 `json:"ad_photo_played_75percent_ratio,omitempty"`
	AdPhotoPlayed10SRatio            float64 `json:"ad_photo_played_10s_ratio,omitempty"`
	AdPhotoPlayed2SRatio             float64 `json:"ad_photo_played_2s_ratio,omitempty"`
	EventPhoneGetThrough             int     `json:"event_phone_get_through,omitempty"`
	EventIntentionConfirmed          int     `json:"event_intention_confirmed,omitempty"`
	EventWechatConnected             int     `json:"event_wechat_connected,omitempty"`
	EventOrderSuccessed              int     `json:"event_order_successed,omitempty"`
	EventPhoneCardActivate           int     `json:"event_phone_card_activate,omitempty"`
	EventMeasurementHouse            int     `json:"event_measurement_house,omitempty"`
	//AdShow                           interface{} `json:"ad_show"`
	ActionNewRatio                 float64 `json:"action_new_ratio,omitempty"`
	EventOutboundCall              int     `json:"event_outbound_call,omitempty"`
	EventOutboundCallCost          float64 `json:"event_outbound_call_cost,omitempty"`
	EventOutboundCallRatio         float64 `json:"event_outbound_call_ratio,omitempty"`
	KeyAction                      int     `json:"key_action,omitempty"`
	KeyActionCost                  float64 `json:"key_action_cost,omitempty"`
	KeyActionRatio                 float64 `json:"key_action_ratio,omitempty"`
	EventCreditCardRecheck         int     `json:"event_credit_card_recheck,omitempty"`
	EventCreditCardRecheckFirstDay int     `json:"event_credit_card_recheck_first_day,omitempty"`
	EventNoIntention               int     `json:"event_no_intention,omitempty"`
}
