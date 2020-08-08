package request

type WebPurifyRequestMethod string

const(
	Check               WebPurifyRequestMethod = "webpurify.live.check"
	CheckCount          WebPurifyRequestMethod = "webpurify.live.checkcount"
	Replace             WebPurifyRequestMethod = "webpurify.live.replace"
	Return              WebPurifyRequestMethod = "webpurify.live.return"
	AddToBlockList      WebPurifyRequestMethod = "webpurify.live.addtoblocklist"
	AddToAllowList      WebPurifyRequestMethod = "webpurify.live.addtoallowlist"
	RemoveFromBlockList WebPurifyRequestMethod = "webpurify.live.removefromblocklist"
	RemoveFromAllowList WebPurifyRequestMethod = "webpurify.live.removefromallowlist"
	GetBlockList        WebPurifyRequestMethod = "webpurify.live.getblocklist"
	GetAllowList        WebPurifyRequestMethod = "webpurify.live.getallowlist"
)
