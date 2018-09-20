// +build webface

package samcatweb

func defaultCSS() string {
	r := "body {\n" +
		"  background: #0e1111;\n" +
		"  color: #84aca8;\n" +
		"}\n" +
		".parent {\n" +
		"  margin-left: auto;\n" +
		"  margin-right: auto;\n" +
		"  border: dotted;\n" +
		"  background: #0e1111;\n" +
		"  width: 95%;\n" +
		"  color: #9789a7;\n" +
		"}\n" +
		"#toolbar {\n" +
		"  margin-left: auto;\n" +
		"  margin-right: auto;\n" +
		"  text-align: center;\n" +
		"  background: #313b3b;\n" +
		"  width: 60%;\n" +
		"}\n" +
		".label {\n" +
		"  float: left;\n" +
		"  display: block;\n" +
		"  text-align: left;\n" +
		"  border: dotted;\n" +
		"  width: 20%\n" +
		"  min-width: 20%\n" +
		"  color: #84aca8;\n" +
		"  background: #313b3b;\n" +
		"}\n" +
		".btn {\n" +
		"  text-align: center;\n" +
		"  border: dotted;\n" +
		"  border-color: #9789a7;\n" +
		"  width: 12%;\n" +
		"  min-width: 12%;\n" +
		"}\n" +
		".content {\n" +
		"  position: relative;\n" +
		"  text-align: right;\n" +
		"  border: dotted;\n" +
		"  width: 75%;\n" +
		"  border-color: #9789a7;\n" +
		"}\n" +
		".inbound {\n" +
		"  background: #313b3b;\n" +
		"  color: #9789a7;\n" +
		"}\n" +
		".outbound {\n" +
		"  background: #9789a7;\n" +
		"  color: #313b3b;\n" +
		"}\n" +
		".i2cp {\n" +
		"  background: #313b3b;\n" +
		"  color: #9789a7;\n" +
		"}\n" +
		".base32 {\n" +
		"  background: #313b3b;\n" +
		"  color: #9789a7;\n" +
		"}\n" +
		".base64 {\n" +
		"  background: #313b3b;\n" +
		"  color: #9789a7;\n" +
		"}\n" +
		".type {\n" +
		"  background: #9789a7;\n" +
		"  color: #313b3b;\n" +
		"}\n" +
		".dest {\n" +
		"  background: #9789a7;\n" +
		"  color: #313b3b;\n" +
		"}\n" +
		"\n"
	return r
}
