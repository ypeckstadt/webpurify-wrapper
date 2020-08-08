package request

type WebPurifyRequestParameterType string

const(

	// APIKey is your API application key.
	APIKey WebPurifyRequestParameterType = "api_key"

	// Text is the text you want checked for profanity.
	Text WebPurifyRequestParameterType = "text"

	// Method is the method to execute on the web purify service
	Method WebPurifyRequestParameterType = "method"

	// Format sets the response format: xml or json. Defaults to xml.
	Format WebPurifyRequestParameterType = "format"

	// Language is the 2 letter language code for the text you are submitting.
	Language WebPurifyRequestParameterType = "lang"

	// Email treats email addresses like profanity. set = 1
	Email WebPurifyRequestParameterType = "semail"

	// Phone treats phone numbers like profanity. set = 1
	Phone WebPurifyRequestParameterType = "sphone"

	// Link treats urls like profanity. set = 1
	Link WebPurifyRequestParameterType = "slink"

	// ReplaceSymbol is the symbol you want to replace each letter of the profane word with.
	ReplaceSymbol WebPurifyRequestParameterType = "replacesymbol"

	// CDATA Set equal to 1 to wrap the return text in <![CDATA[ ]]> , Use this if you plan on passing strings that may break XML (ie “<” and “&”).
	CDATA WebPurifyRequestParameterType = "cdata"

	// Word is the word you want added to your block list.
	Word WebPurifyRequestParameterType = "word"

	// DeepSearch is the deep Search: If you want this word to search substrings. set = 1. is the word you want added to your block list.
	DeepSearch WebPurifyRequestParameterType = "ds"
)

type WebPurifyRequestParameter struct {
	Type  WebPurifyRequestParameterType
	Value string
}