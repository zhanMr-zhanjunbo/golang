package model
//供应商：Supplier
//Id
//供应商编码 SupplierCode
//供应商名称 SupplierName
//供应商id   SupplierId    foreign key(supplierId) references Order(id) on delete cascade
//联系人     Contact
//联系电话   ContactNumber
//联系地址   ContactAddress
//传真       Fax
//描述       Describe
type Supplier struct {
	Id int
	SupplierCode string
	SupplierName string
	SupplierId     int
	Contact        string
	ContactNumber  int
	ContactAddress string
	Fax            string
	Describe       string
}
