Please refactor the code below to make it more concise, efficient and readable with good logic flow.
Both A & B are doing the same thing with A written in Golang and B in PHP. Please answer for the 2
languages.

### A (Golang Code)
```
func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1 : indexClosingBracketFound-1])
			} else {
				return ""
			}
		} else {
			return ""
		}
	} else {
		return ""
	}
	return ""
}
```

### B (PHP Code)
```
function findFirstStringInBracket($str){
    if(strlen($str) > 0){
        $firstbracket = strstr($str, '(');
        
        if($firstbracket){
            $firstbracket = ltrim($firstbracket, '(');
            return strstr($firstbracket, ')', true);
        }else{
            return '';
        }
    }else{
        return '';
    }
}
```