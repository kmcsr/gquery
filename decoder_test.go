
package gquery

import (
	io "io"
	bytes "bytes"
	testing "testing"
)

func TestReadToStr(t *testing.T){
	r := bytes.NewReader(([]byte)("abcd文efgo28y</script中hh 00</script>abcdefxx"))
	data, err := readToStr(r, "</script>")
	t.Log("read:", (string)(data), err)
	buf, err := io.ReadAll(r)
	t.Log("lose:", (string)(buf), err)
}

func TestDecodeDoc(t *testing.T){
	str := `<!DOCTYPE html>
<html>
	<head>
		<title>Test page</title>
		<meta name="viewport" content="initial-scale=0.8"/>
		<link rel="stylesheet" href="./index.css" type="text/css" />
	</head>
	<body>
		<div id="body">
			Hello Sir <br/>
			testtext testtext  testtext   testtext    testtext    testtext	testtext		testtext			testtext
			testtext

			testtext
		</div>
		<script type="text/javascript" src="./jquery-3.6.0.min.js"></script>
		<script type="text/javascript" src="./index.js"></script>
		<script>
			var testvar = "";
			function testfunc(){
				testvar += "1234";
				console.log(testvar);
			}
			testfunc();
		</script>
	</body>
</html>
`
	t.Log("decoding:", str)
	doc, err := DecodeDocString(str)
	if err != nil {
		t.Fatal("decode error:", err)
	}
	t.Log("decoded:", doc)
	t.Logf("doctype: '%s'", doc.GetDocType().GetValue())
	t.Logf("content: %s", doc.GetText())
}

