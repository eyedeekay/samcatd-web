// +build webface

package samcatwebstyle

func defaultCSS() string {
	r := "body {\n" +
		"  background: #0e1111;\n" +
		"  color: #84aca8;\n" +
		"}\n" +
		".btn {\n" +
		"  text-align: center;\n" +
		"  border: dotted;\n" +
		"  border-color: #9789a7;\n" +
		"  width: 111px;\n" +
		"}\n" +
		".inbound {\n" +
		"  text-align: right;\n" +
		"  border: dotted;\n" +
		"  border-color: #9789a7;\n" +
		"}\n" +
		".outbound {\n" +
		"  text-align: right;\n" +
		"  border: dotted;\n" +
		"  border-color: #9789a7;\n" +
		"}\n" +
		".i2cp {\n" +
		"  text-align: right;\n" +
		"  border: dotted;\n" +
		"  border-color: #9789a7;\n" +
		"}\n" +
		".base32 {\n" +
		"  text-align: right;\n" +
		"  border: dotted;\n" +
		"  border-color: #9789a7;\n" +
		"}\n" +
		".base64 {\n" +
		"  text-align: right;\n" +
		"  border: dotted;\n" +
		"  border-color: #9789a7;\n" +
		"}\n" +
		".type {\n" +
		"  text-align: right;\n" +
		"  border: dotted;\n" +
		"  border-color: #9789a7;\n" +
		"}\n" +
		".dest {\n" +
		"  text-align: right;\n" +
		"  border: dotted;\n" +
		"  border-color: #9789a7;\n" +
		"}\n" +
		".parent {\n" +
		"  margin-left: 60px;\n" +
		"  border: dotted;\n" +
		"  background: #0e1111;\n" +
		"  color: #9789a7;\n" +
		"}\n" +
		"#toolbar {\n" +
		"  margin-left: auto;\n" +
		"  margin-right: auto;\n" +
		"  text-align: center;\n" +
		"  background: #313b3b;\n" +
		"  width: 666px;\n" +
		"}\n" +
		"\n"
	return r
}
