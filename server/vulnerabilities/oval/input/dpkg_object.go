package oval_input

// https://oval.mitre.org/language/version5.10.1/ovaldefinition/documentation/linux-definitions-schema.html#dpkginfo_object

type dpkObjectNameXML struct {
	VarRef string `xml:"var_ref,attr"`
}

type DpkgObjectXML struct {
	Id   string           `xml:"id,attr"`
	Name dpkObjectNameXML `xml:"name"`
}
