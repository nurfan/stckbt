<?php

$txt = "Nurfan Ramadhandi Backend Engineer (Golang)";
echo findFirstStringInBracket($txt);


function findFirstStringInBracket($str) {

    $firstbracket = strstr($str, '(');
    if($firstbracket){
        $firstbracket = ltrim($firstbracket, '(');
        return strstr($firstbracket, ')', true);
    }
        
    return "";
}

?>