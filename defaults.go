// +build webface

package samcatweb

func defaultJS() string {
	r := "" +
		"" +
		"" +
		""
	return r
}

func defaultCSS() string {
	r := ".btn {" +
		"  text-align: center;" +
		"  color: #9789a7;" +
		"  border: dotted;" +
		"  width: 70px;" +
		"}" +
		"" +
		".toolbar {" +
		"  margin-left: auto;" +
		"  margin-right: auto;" +
		"  background-color: #0e1111;" +
		"  width: 420px;" +
		"}" +
		"\n"
	return r
}
