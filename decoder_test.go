
package gquery_test

import (
	testing "testing"
	gquery "github.com/kmcsr/gquery"
)

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
			<p>testtext testtext  testtext   testtext</p>    testtext    testtext	<span>testtext		testtext</span>			testtext
			testtext

			<p>testtext</p>
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
	doc, err := gquery.DecodeDocString(str)
	if err != nil {
		t.Fatal("decode error:", err)
	}
	t.Log("decoded:", doc)
	t.Logf("doctype: '%s'", doc.GetDocType().GetValue())
	t.Logf("content: %s", doc.GetText())
}

