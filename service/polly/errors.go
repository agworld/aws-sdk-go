// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package polly

const (

	// ErrCodeInvalidLexiconException for service response error code
	// "InvalidLexiconException".
	//
	// Amazon Polly can't find the specified lexicon. Verify that the lexicon's
	// name is spelled correctly, and then try again.
	ErrCodeInvalidLexiconException = "InvalidLexiconException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The NextToken is invalid. Verify that it's spelled correctly, and then try
	// again.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidSampleRateException for service response error code
	// "InvalidSampleRateException".
	//
	// The specified sample rate is not valid.
	ErrCodeInvalidSampleRateException = "InvalidSampleRateException"

	// ErrCodeInvalidSsmlException for service response error code
	// "InvalidSsmlException".
	//
	// The SSML you provided is invalid. Verify the SSML syntax, spelling of tags
	// and values, and then try again.
	ErrCodeInvalidSsmlException = "InvalidSsmlException"

	// ErrCodeLexiconNotFoundException for service response error code
	// "LexiconNotFoundException".
	//
	// Amazon Polly can't find the specified lexicon. This could be caused by a
	// lexicon that is missing, its name is misspelled or specifying a lexicon that
	// is in a different region.
	//
	// Verify that the lexicon exists, is in the region (see ListLexicons) and that
	// you spelled its name is spelled correctly. Then try again.
	ErrCodeLexiconNotFoundException = "LexiconNotFoundException"

	// ErrCodeLexiconSizeExceededException for service response error code
	// "LexiconSizeExceededException".
	//
	// The maximum size of the specified lexicon would be exceeded by this operation.
	ErrCodeLexiconSizeExceededException = "LexiconSizeExceededException"

	// ErrCodeMarksNotSupportedForFormatException for service response error code
	// "MarksNotSupportedForFormatException".
	//
	// Speech marks are not supported for the OutputFormat selected. Speech marks
	// are only available for content in json format.
	ErrCodeMarksNotSupportedForFormatException = "MarksNotSupportedForFormatException"

	// ErrCodeMaxLexemeLengthExceededException for service response error code
	// "MaxLexemeLengthExceededException".
	//
	// The maximum size of the lexeme would be exceeded by this operation.
	ErrCodeMaxLexemeLengthExceededException = "MaxLexemeLengthExceededException"

	// ErrCodeMaxLexiconsNumberExceededException for service response error code
	// "MaxLexiconsNumberExceededException".
	//
	// The maximum number of lexicons would be exceeded by this operation.
	ErrCodeMaxLexiconsNumberExceededException = "MaxLexiconsNumberExceededException"

	// ErrCodeServiceFailureException for service response error code
	// "ServiceFailureException".
	//
	// An unknown condition has caused a service failure.
	ErrCodeServiceFailureException = "ServiceFailureException"

	// ErrCodeSsmlMarksNotSupportedForTextTypeException for service response error code
	// "SsmlMarksNotSupportedForTextTypeException".
	//
	// SSML speech marks are not supported for plain text-type input.
	ErrCodeSsmlMarksNotSupportedForTextTypeException = "SsmlMarksNotSupportedForTextTypeException"

	// ErrCodeTextLengthExceededException for service response error code
	// "TextLengthExceededException".
	//
	// The value of the "Text" parameter is longer than the accepted limits. The
	// limit for input text is a maximum of 3000 characters total, of which no more
	// than 1500 can be billed characters. SSML tags are not counted as billed characters.
	ErrCodeTextLengthExceededException = "TextLengthExceededException"

	// ErrCodeUnsupportedPlsAlphabetException for service response error code
	// "UnsupportedPlsAlphabetException".
	//
	// The alphabet specified by the lexicon is not a supported alphabet. Valid
	// values are x-sampa and ipa.
	ErrCodeUnsupportedPlsAlphabetException = "UnsupportedPlsAlphabetException"

	// ErrCodeUnsupportedPlsLanguageException for service response error code
	// "UnsupportedPlsLanguageException".
	//
	// The language specified in the lexicon is unsupported. For a list of supported
	// languages, see Lexicon Attributes (http://docs.aws.amazon.com/polly/latest/dg/API_LexiconAttributes.html).
	ErrCodeUnsupportedPlsLanguageException = "UnsupportedPlsLanguageException"
)
//Added a line for testing
//Adding another line for Git event testing part 2
//Adding another line for Git event testing part 2.1
//Adding another line for Git event testing part 2.2
