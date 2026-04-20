package config

import (
	"reflect"
	"testing"
)

// TestExtractHeaders 测试 parseHeaderArr
func TestExtractHeaders(t *testing.T) {
	input := `
a: foo
b: bar
c: foo:bar`
	expected := map[string]string{
		"a": "foo",
		"b": "bar",
		"c": "foo:bar",
	}

	parsedHeaders := extractHeaders(input)
	if !reflect.DeepEqual(parsedHeaders, expected) {
		t.Errorf("Expected %v, got %v", expected, parsedHeaders)
	}
}

// TestReplaceParaWithTimestamp 测试 replacePara 在保留原有变量替换行为的同时支持 timestamp。
func TestReplaceParaWithTimestamp(t *testing.T) {
	domains := &Domains{
		Ipv4Addr: "1.2.3.4",
		Ipv4Domains: []*Domain{
			{DomainName: "example.com", SubDomain: "www"},
		},
		Ipv6Addr: "2001:db8::1",
		Ipv6Domains: []*Domain{
			{DomainName: "example.com", SubDomain: "v6"},
		},
	}

	template := "ipv4=#{ipv4Addr};ipv4r=#{ipv4Result};ipv4d=#{ipv4Domains};ipv6=#{ipv6Addr};ipv6r=#{ipv6Result};ipv6d=#{ipv6Domains};ts=#{timestamp}"
	timestamp := "1713571200"
	got := replacePara(domains, template, UpdatedSuccess, UpdatedFailed, timestamp)
	want := "ipv4=1.2.3.4;ipv4r=success;ipv4d=www.example.com;ipv6=2001:db8::1;ipv6r=failed;ipv6d=v6.example.com;ts=1713571200"

	if got != want {
		t.Fatalf("replacePara() = %q, want %q", got, want)
	}
}
