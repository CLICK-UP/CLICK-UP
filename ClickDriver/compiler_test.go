package ClickDriver

import (
	"ServiceContext"
	"fmt"
	"testing"
)

func TestUDFCompiler(t *testing.T) {
	var user_defined_element []User_defined_element
	click_name := "empty1"
	click_hh := `#ifndef EMPTY1_H
#define EMPTY1_H
#include <click/element.hh>
CLICK_DECLS

class Empty1 : public Element {

public:	
	Empty1();
	~Empty1();
	
	const char* class_name() const { return "Empty1"; }
	const char* port_count() const { return PORTS_1_1; }
	const char* processing() const { return PUSH; }
	const char* flow_code()  const { return "x/y"; }
	
	int initialize(ErrorHandler *errh);
	int configure(Vector<String>& conf, ErrorHandler* errh);
	void push(int i, Packet* p);

private:
	String _prefix;
};

CLICK_ENDDECLS

#endif`
	click_cc := `#include <click/config.h>
#include <click/error.hh>
#include <click/args.hh>

#include <cstring>

#include "empty1.hh"

CLICK_DECLS

Empty1::Empty1()
	: _prefix() {}

Empty1::~Empty1() {}

int Empty1::initialize(ErrorHandler *errh)
{
	errh->message("Successfully initialized.");
	return 0;
}

int Empty1::configure(Vector<String>& conf, ErrorHandler* errh)
{
	Args args = Args(this, errh).bind(conf);
	
	if(args.read_mp("PREFIX", _prefix).execute() < 0)
		return -1;
	
	errh->message("prefix: %s", _prefix.c_str());
	return 0;
}

void Empty1::push(int i, Packet* p)
{
	click_chatter("%s packet on port %i of length %u B", _prefix.c_str(), i, p->length());
	output(0).push(p);
}

CLICK_ENDDECLS
EXPORT_ELEMENT(Empty1)
`
	udf := User_defined_element{click_name, click_hh, click_cc}
	user_defined_element = append(user_defined_element, udf)
	err := UDFCompiler(user_defined_element)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("UDFCompiler success!!!")
	}

}

func TestSCCompiler(t *testing.T) {
	var serviceContext []ServiceContext.ServiceContext
	sc := ServiceContext.ServiceContext{"../elements/tcpudp/iprewriter.cc", "../elements/tcpudp/iprewriter.hh", "IPRewriter"}
	serviceContext = append(serviceContext, sc)
	err := SCCompiler(serviceContext)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("SCCompiler success!!!")
	}
}
