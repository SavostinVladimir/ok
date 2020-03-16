package ok

const (
	baseUrlREST  = "https://api.ok.ru/fb.do"
	baseUrlGRAPH = "https://api.ok.ru/graph"
)

type GraphUpdateType string

const (
	MESSAGE_CREATED  GraphUpdateType = "MESSAGE_CREATED"
	MESSAGE_CALLBACK GraphUpdateType = "MESSAGE_CALLBACK"
	CHAT_SYSTEM      GraphUpdateType = "CHAT_SYSTEM"
)

//GroupBeanField - represnets the which filed to fetch on get groups req
type GroupBeanField string

const (
	ABBREVIATION                         GroupBeanField = "ABBREVIATION"
	ACCESS_TYPE                                         = "ACCESS_TYPE"
	ADDRESS                                             = "ADDRESS"
	ADD_CHANNEL_ALLOWED                                 = "ADD_CHANNEL_ALLOWED"
	ADD_PHOTOALBUM_ALLOWED                              = "ADD_PHOTOALBUM_ALLOWED"
	ADD_THEME_ALLOWED                                   = "ADD_THEME_ALLOWED"
	ADD_VIDEO_ALLOWED                                   = "ADD_VIDEO_ALLOWED"
	ADMIN_ID                                            = "ADMIN_ID"
	ADS_MANAGER_ALLOWED                                 = "ADS_MANAGER_ALLOWED"
	ADVANCED_PUBLICATION_ALLOWED                        = "ADVANCED_PUBLICATION_ALLOWED"
	AGE_RESTRICTED                                      = "AGE_RESTRICTED"
	BLOCKED                                             = "BLOCKED"
	BUSINESS                                            = "BUSINESS"
	CALL_ALLOWED                                        = "CALL_ALLOWED"
	CATALOG_CREATE_ALLOWED                              = "CATALOG_CREATE_ALLOWED"
	CATEGORY                                            = "CATEGORY"
	CHANGE_AVATAR_ALLOWED                               = "CHANGE_AVATAR_ALLOWED"
	CHANGE_TYPE_ALLOWED                                 = "CHANGE_TYPE_ALLOWED"
	CITY                                                = "CITY"
	COMMENT_AS_OFFICIAL                                 = "COMMENT_AS_OFFICIAL"
	COMMUNITY                                           = "COMMUNITY"
	CONTENT_AS_OFFICIAL                                 = "CONTENT_AS_OFFICIAL"
	COUNTRY                                             = "COUNTRY"
	COVER                                               = "COVER"
	COVER_SERIES                                        = "COVER_SERIES"
	CREATED_MS                                          = "CREATED_MS"
	CREATE_ADS_ALLOWED                                  = "CREATE_ADS_ALLOWED"
	DELETE_ALLOWED                                      = "DELETE_ALLOWED"
	DESCRIPTION                                         = "DESCRIPTION"
	EDIT_ALLOWED                                        = "EDIT_ALLOWED"
	EDIT_APPS_ALLOWED                                   = "EDIT_APPS_ALLOWED"
	END_DATE                                            = "END_DATE"
	FEED_SUBSCRIPTION                                   = "FEED_SUBSCRIPTION"
	FOLLOWERS_COUNT                                     = "FOLLOWERS_COUNT"
	FOLLOW_ALLOWED                                      = "FOLLOW_ALLOWED"
	FRIENDS_COUNT                                       = "FRIENDS_COUNT"
	GRADUATE_YEAR                                       = "GRADUATE_YEAR"
	HOMEPAGE_NAME                                       = "HOMEPAGE_NAME"
	HOMEPAGE_URL                                        = "HOMEPAGE_URL"
	INVITATION_SENT                                     = "INVITATION_SENT"
	INVITE_ALLOWED                                      = "INVITE_ALLOWED"
	JOIN_ALLOWED                                        = "JOIN_ALLOWED"
	LEAVE_ALLOWED                                       = "LEAVE_ALLOWED"
	LINK_CAROUSEL_ALLOWED                               = "LINK_CAROUSEL_ALLOWED"
	LINK_POSTING_ALLOWED                                = "LINK_POSTING_ALLOWED"
	LOCATION_ID                                         = "LOCATION_ID"
	LOCATION_LATITUDE                                   = "LOCATION_LATITUDE"
	LOCATION_LONGITUDE                                  = "LOCATION_LONGITUDE"
	LOCATION_ZOOM                                       = "LOCATION_ZOOM"
	MAIN_PAGE_TAB                                       = "MAIN_PAGE_TAB"
	MAIN_PHOTO                                          = "MAIN_PHOTO"
	MANAGE_MEMBERS                                      = "MANAGE_MEMBERS"
	MEMBERS_COUNT                                       = "MEMBERS_COUNT"
	MENTIONS_SUBSCRIPTION                               = "MENTIONS_SUBSCRIPTION"
	MENTIONS_SUBSCRIPTION_ALLOWED                       = "MENTIONS_SUBSCRIPTION_ALLOWED"
	MESSAGES_ALLOWED                                    = "MESSAGES_ALLOWED"
	MESSAGING_ALLOWED                                   = "MESSAGING_ALLOWED"
	MIN_AGE                                             = "MIN_AGE"
	MOBILE_COVER                                        = "MOBILE_COVER"
	NAME                                                = "NAME"
	NOTIFICATIONS_SUBSCRIPTION                          = "NOTIFICATIONS_SUBSCRIPTION"
	PARTNER_PROGRAM_ALLOWED                             = "PARTNER_PROGRAM_ALLOWED"
	PARTNER_PROGRAM_STATUS                              = "PARTNER_PROGRAM_STATUS"
	PENALTY_POINTS_ALLOWED                              = "PENALTY_POINTS_ALLOWED"
	PHONE                                               = "PHONE"
	PHOTOS_TAB_HIDDEN                                   = "PHOTOS_TAB_HIDDEN"
	PHOTO_ID                                            = "PHOTO_ID"
	PIC_AVATAR                                          = "PIC_AVATAR"
	POSSIBLE_MEMBERS_COUNT                              = "POSSIBLE_MEMBERS_COUNT"
	PREMIUM                                             = "PREMIUM"
	PRIVATE                                             = "PRIVATE"
	PRODUCTS_TAB_HIDDEN                                 = "PRODUCTS_TAB_HIDDEN"
	PRODUCT_CREATE_ALLOWED                              = "PRODUCT_CREATE_ALLOWED"
	PRODUCT_CREATE_SUGGESTED_ALLOWED                    = "PRODUCT_CREATE_SUGGESTED_ALLOWED"
	PRODUCT_CREATE_ZERO_LIFETIME_ALLOWED                = "PRODUCT_CREATE_ZERO_LIFETIME_ALLOWED"
	PROFILE_BUTTONS                                     = "PROFILE_BUTTONS"
	PROMO_THEME_ALLOWED                                 = "PROMO_THEME_ALLOWED"
	PUBLISH_DELAYED_THEME_ALLOWED                       = "PUBLISH_DELAYED_THEME_ALLOWED"
	REF                                                 = "REF"
	REQUEST_SENT                                        = "REQUEST_SENT"
	RESHARE_ALLOWED                                     = "RESHARE_ALLOWED"
	ROLE                                                = "ROLE"
	SCOPE_ID                                            = "SCOPE_ID"
	SHOP_VISIBLE_ADMIN                                  = "SHOP_VISIBLE_ADMIN"
	SHOP_VISIBLE_PUBLIC                                 = "SHOP_VISIBLE_PUBLIC"
	SHORTNAME                                           = "SHORTNAME"
	START_DATE                                          = "START_DATE"
	STATS_ALLOWED                                       = "STATS_ALLOWED"
	STATUS                                              = "STATUS"
	SUBCATEGORY_ID                                      = "SUBCATEGORY_ID"
	SUGGEST_THEME_ALLOWED                               = "SUGGEST_THEME_ALLOWED"
	TAGS                                                = "TAGS"
	TRANSFERS_ALLOWED                                   = "TRANSFERS_ALLOWED"
	UID                                                 = "UID"
	UNFOLLOW_ALLOWED                                    = "UNFOLLOW_ALLOWED"
	VIDEO_TAB_HIDDEN                                    = "VIDEO_TAB_HIDDEN"
	VIEW_MEMBERS_ALLOWED                                = "VIEW_MEMBERS_ALLOWED"
	VIEW_MODERATORS_ALLOWED                             = "VIEW_MODERATORS_ALLOWED"
	YEAR_FROM                                           = "YEAR_FROM"
	YEAR_TO                                             = "YEAR_TO"
)
