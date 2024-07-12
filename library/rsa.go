package library

import "github.com/wenzhenxi/gorsa"

var RsaPubkey = `-----BEGIN public key-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEApCxq50km7NGI3ABETINx
wmYmCnPSJJF9In8/XpOhehjz21ZpjRCuRYZcGInYHl6+TaoHOzXuueh4Uxt9/FUR
Oly0dR4kwuZuwtfaX4Ukb+nuWJu83t+C83hFmAxPKcm1aJBO/2EhHQ6xnVpkZ/dc
xMwku+MvonMgPUqyzgeJzXSUt6g4e2n2U+RFLIIT5hfQTsZhCEGwK3Ck8CPlD01I
seTt5+IoDLRCpQ3WmRO/ykYPd77tTPv/KL7EGkPO+u0kdp0mm3jA3eJ8tjU86CHN
6ONPOgT37jkPgFzgnz9Af4RGlYSvKD7uFC24iyC9ywXqb3+mZiadn1yO1m+j6sEL
uwIDAQAB
-----END public key-----
`

var RsaPrvkey = `-----BEGIN private key-----
MIIEpAIBAAKCAQEApCxq50km7NGI3ABETINxwmYmCnPSJJF9In8/XpOhehjz21Zp
jRCuRYZcGInYHl6+TaoHOzXuueh4Uxt9/FUROly0dR4kwuZuwtfaX4Ukb+nuWJu8
3t+C83hFmAxPKcm1aJBO/2EhHQ6xnVpkZ/dcxMwku+MvonMgPUqyzgeJzXSUt6g4
e2n2U+RFLIIT5hfQTsZhCEGwK3Ck8CPlD01IseTt5+IoDLRCpQ3WmRO/ykYPd77t
TPv/KL7EGkPO+u0kdp0mm3jA3eJ8tjU86CHN6ONPOgT37jkPgFzgnz9Af4RGlYSv
KD7uFC24iyC9ywXqb3+mZiadn1yO1m+j6sELuwIDAQABAoIBAFUA/pUkhjgGQfOP
W822HVfHZhjxYTQ39G/BLQF+NRSwCMtfkaVNpr4u1E6MSM9oxOYXjEj3ItU1bd+y
E2hsb1AlofYmFminwV+3PcVJfdVIbXWJpaBMYFRJyNbyQKblXr12QsDt/KQ78HEF
1YkTRgBLpxV1izt0mpPBspUEyDANQXUIfv5bXzRJX4Ng39mDlDkl8NaHs2u1738O
0k3C2Ol94KMkNfE7lRxZSBMtsTu1DrrqvhhrsROmduu0TM7p05+nFpCAOIUEYYrj
4/XtSPmv+uzaJIDDplW45d3qIQXoYQ6ZqBUaGqQR8ktEwZVkNM5Lw22vu44lQq3t
IwIpMEECgYEA0vbrgA047d5+BBUgbICvraOkBUiHfEzZbmbRe4Olsxv1N8U6FgwR
RMfFnnSfvbti66zt0IhkPhlmW3F/jMC69XclMSvzlDeOLGdIRg/J66eI6raGzJd1
XFP+BfMZzXb3kxjKePx8bgs7/2tYPvItJpYfxXOPPuZmoKgn2mdDlyECgYEAxzhl
uDhl6BqP3MrCQ6KcOwvOLAFUDa9O2MTEuHBvCO13eEO9ROSRaTCXMnkhw8bxXXIh
Wm0UrRSCU/yU96C9mQlkOKSmDV6NyCO7O8VgCElEcwHpbHJeKSPxGmZ/Xr0v0PSQ
GCOUko/3oKgMLSQzWQLEBBfQRGLni3ZlxrKG81sCgYEAtfsD2PkO87oqsDDdQFud
r8ZqVZhkLyhHo4GUENzWEfP7CvnaCmysdM1zPiXSKiO7yBrrJiiReDQG5Li6U1tp
qxb1AT6tplxqCwmpAa7a+qoRZI7BJzk7psJZbR5wUpneJIiNF+KNH4wXXFRn0hys
MA3uKDBwD81GGC4V+sN7PaECgYEAxg8gubSAQ+/cZNLDrcBFeEJ1V7R+XxiR19pQ
BVnqWcLUSqF10r82fZCMUOayldP8dD2aQS8/Jg/EvkDXzRRmSnZWyCWQ3KvLY1V6
uKKypapJvdkBuwo86MPN2MRcDtEhb6kkdIszEz7EHlnTuTFLR0wDmloeeU6nCjTM
XG1y0w8CgYAeS5A+DaxKZMr5/7TjUmeS/ZTgGyRg7DK3T0FL9hXyyNZlxwLX5hyH
XWZqGeaMgOr8sXCrvVtmhboO/6Zlne/IyDEQlNEGHi655bBbHOT6WDS2PojZnC9K
4xR6q5jBZZUdMfwKa4I6qOZQlUsaEiTR1WDhI4aekQkUKfoe3vYarg==
-----END private key-----
`

func PublicEncrypt(str string) (string, error) {
	encrypt, err := gorsa.PublicEncrypt(str, RsaPubkey)
	if err != nil {
		return "", err
	}
	return encrypt, nil
}

func PrikeyDecrypt(str string) (string, error) {
	pridecrypt, err := gorsa.PriKeyDecrypt(str, RsaPubkey)
	if err != nil {
		return "", err
	}
	return pridecrypt, nil
}
