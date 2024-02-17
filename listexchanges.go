package stockclient

const (
	// RegionAll stock exchanges (bonds excluded)
	RegionAll = RegionStockholm + RegionCopenhagen + RegionHelsinki + RegionIceland + RegionBaltic
	// RegionStockholm stock exchange list (876 assets)
	RegionStockholm = "List=M:INET:XSTO:SEEQ-SHR&List=M:INET:XSTO:SEEQ-SHR-CCP&List=M:INET:XSTO:SEEQ-SHR-NOK&List=M:INET:XSTO:SEEQ-SHR-AO&List=M:INET:FNSE:SEMM-NM&"
	// RegionCopenhagen stock exchange list (377 assets)
	RegionCopenhagen = "List=M:INET:XCSE:DKEQ-SHR&List=M:INET:XCSE:DKEQ-SHR-CCP&List=M:INET:FNDK:FNDK-CPH&List=M:INET:FNSE:SEMM-FN-NOK&"
	// RegionHelsinki stock exchange list (197 assets)
	RegionHelsinki = "List=M:INET:XHEL:FIEQ-SHR&List=M:INET:XHEL:FIEQ-SHR-CCP&List=M:INET:XHEL:FIEQ-SHR-IC&List=M:INET:FNFI:SEMM-FN-HEL&List=M:INET:FNFI:SEMM-FN-HE-ERW&"
	// RegionIceland stock exchange list (31 assets)
	RegionIceland = "List=M:INET:XICE:ISEQ-SHR&List=M:INET:FNIS:ISEC-SHR&"
	// RegionBaltic stock exchange list (72 assets)
	RegionBaltic = "List=M:INET:FNEE:EEMM-SHR&List=M:INET:FNLV:LVMM-SHR&List=M:INET:FNLT:LTMM-SHR&List=M:INET:XTAL:EEEQ-SHR&List=M:INET:XRIS:LVEQ-SHR&List=M:INET:XLIT:LTEQ-SHR&"
)
