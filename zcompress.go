// *** DO NOT MODIFY ***
// AUTOGENERATED BY go generate from compress_generate.go

package dns

func compressionLenHelperType(c map[string]int, r RR) {
	switch x := r.(type) {
	case *AFSDB:
		compressionLenHelper(c, x.Hostname)
	case *RRSIG:
		compressionLenHelper(c, x.SignerName)
	case *MD:
		compressionLenHelper(c, x.Md)
	case *MX:
		compressionLenHelper(c, x.Mx)
	case *TALINK:
		compressionLenHelper(c, x.PreviousName)
		compressionLenHelper(c, x.NextName)
	case *MF:
		compressionLenHelper(c, x.Mf)
	case *RP:
		compressionLenHelper(c, x.Mbox)
		compressionLenHelper(c, x.Txt)
	case *RT:
		compressionLenHelper(c, x.Host)
	case *SOA:
		compressionLenHelper(c, x.Ns)
		compressionLenHelper(c, x.Mbox)
	case *NSAPPTR:
		compressionLenHelper(c, x.Ptr)
	case *CNAME:
		compressionLenHelper(c, x.Target)
	case *LP:
		compressionLenHelper(c, x.Fqdn)
	case *NSEC:
		compressionLenHelper(c, x.NextDomain)
	case *PTR:
		compressionLenHelper(c, x.Ptr)
	case *TSIG:
		compressionLenHelper(c, x.Algorithm)
	case *MG:
		compressionLenHelper(c, x.Mg)
	case *PX:
		compressionLenHelper(c, x.Map822)
		compressionLenHelper(c, x.Mapx400)
	case *SRV:
		compressionLenHelper(c, x.Target)
	case *TKEY:
		compressionLenHelper(c, x.Algorithm)
	case *MINFO:
		compressionLenHelper(c, x.Rmail)
		compressionLenHelper(c, x.Email)
	case *NAPTR:
		compressionLenHelper(c, x.Replacement)
	case *SIG:
		compressionLenHelper(c, x.SignerName)
	case *DNAME:
		compressionLenHelper(c, x.Target)
	case *HIP:
		for i := range x.RendezvousServers {
			compressionLenHelper(c, x.RendezvousServers[i])
		}
	case *KX:
		compressionLenHelper(c, x.Exchanger)
	case *MB:
		compressionLenHelper(c, x.Mb)
	case *MR:
		compressionLenHelper(c, x.Mr)
	case *NS:
		compressionLenHelper(c, x.Ns)
	}
}

func compressionLenSearchType(c map[string]int, r RR) (int, bool) {
	switch x := r.(type) {
	case *PTR:
		k1, ok1 := compressionLenSearch(c, x.Ptr)
		return k1, ok1
	case *RT:
		k1, ok1 := compressionLenSearch(c, x.Host)
		return k1, ok1
	case *SOA:
		k1, ok1 := compressionLenSearch(c, x.Ns)
		k2, ok2 := compressionLenSearch(c, x.Mbox)
		return k1 + k2, ok1 && ok2
	case *MB:
		k1, ok1 := compressionLenSearch(c, x.Mb)
		return k1, ok1
	case *MG:
		k1, ok1 := compressionLenSearch(c, x.Mg)
		return k1, ok1
	case *MR:
		k1, ok1 := compressionLenSearch(c, x.Mr)
		return k1, ok1
	case *NS:
		k1, ok1 := compressionLenSearch(c, x.Ns)
		return k1, ok1
	case *MINFO:
		k1, ok1 := compressionLenSearch(c, x.Rmail)
		k2, ok2 := compressionLenSearch(c, x.Email)
		return k1 + k2, ok1 && ok2
	case *MX:
		k1, ok1 := compressionLenSearch(c, x.Mx)
		return k1, ok1
	case *AFSDB:
		k1, ok1 := compressionLenSearch(c, x.Hostname)
		return k1, ok1
	case *CNAME:
		k1, ok1 := compressionLenSearch(c, x.Target)
		return k1, ok1
	case *MD:
		k1, ok1 := compressionLenSearch(c, x.Md)
		return k1, ok1
	case *MF:
		k1, ok1 := compressionLenSearch(c, x.Mf)
		return k1, ok1
	}
	return 0, false
}
