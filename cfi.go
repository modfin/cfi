package cfi

import (
	"fmt"
	"regexp"
	"strings"
)

// https://en.wikipedia.org/wiki/ISO_10962

// From takes a cfi token, validates it and decorates it named Category, Group and Attributes
func From(cfi string) (CFI, error) {

	if !ValidCG(cfi) {
		return CFI{}, fmt.Errorf("CFI %s is not valid", cfi)
	}

	return CFI{
		tag: cfi,
	}, nil
}

var noncfichars = regexp.MustCompile("[^A-Z]")

// New takes a cfi token, but does not validate it (however it will make it into a tag, eg. supplying "EP" will result in "EPXXXX")
func New(cfi string) CFI {

	padRight := func(str, pad string, lenght int) string {
		for {
			str += pad
			if len(str) > lenght {
				return str[0:lenght]
			}
		}
	}

	return CFI{
		tag: noncfichars.ReplaceAllString(padRight(strings.ToUpper(cfi), "X", 6), "X"),
	}
}

var NA = map[rune]string{
	'_': "N/A",
	'X': "N/A",
}

var Form = map[rune]string{ // Debt
	'_': "Form",
	'B': "Bearer",
	'R': "Registered",
	'N': "Bearer/Registered",
	'Z': "Bearer depository receipt",
	'A': "Registered depository receipt",
	'M': "Misc.",
}

var EquityOwnership = map[rune]string{
	'_': "Ownership",
	'T': "Restrictions",
	'U': "Free",
}
var DebtGuarantee = map[rune]string{ // Debt
	'_': "Guarantee",
	'T': "Government / State Guarantee",
	'G': "Joint Guarantee",
	'S': "Secured",
	'U': "Unsecured / Unguaranteed",
	'P': "Negative Pledge",
	'N': "Senior",
	'O': "Senior Subordinated",
	'Q': "Junior",
	'J': "Junior Subordinated",
	'C': "Supranational",
}

var DebtTypeOfInterest = map[rune]string{ // Debt
	'_': "Type of interest",
	'F': "Fixed rate",
	'Z': "Zero rate / discounted rate",
	'V': "Variable",
	'C': "Cash payment",
	'K': "payment in kind",
}

var DebtRedemptionReimbursement = map[rune]string{ // Debt
	'_': "Redemption / Reimbursement",
	'F': "Fixed Maturity",
	'G': "Fixed Maturity with Call Feature",
	'C': "Fixed Maturity with Put Feature",
	'D': "Fixed Maturity with Put and Call",
	'A': "Amortization Plan",
	'B': "Amortization Plan with Call Feature",
	'T': "Amortization Plan with Put Feature",
	'L': "Amortization Plan with Put and Call",
	'P': "Perpetual",
	'Q': "Perpetual with Call Feature",
	'R': "Perpetual with Put Feature",
	'E': "Extendible",
}

var DebtDistribution = map[rune]string{ // Debt
	'_': "Distribution",
	'F': "Fixed Interest Payments",
	'D': "Dividend Payments",
	'V': "Variable Interest Payments",
	'Y': "No Payments",
	'M': "Misc.",
}

var UnderlyingAsset = map[rune]string{
	'_': "Underlying Asset",
	'B': "Baskets",          // Equity, Debt
	'S': "Equities",         // Equity, Debt
	'D': "Debt Instruments", // Equity, Debt
	'G': "Derivatives",      // Equity
	'T': "Commodities",      // Equity, Debt
	'C': "Currencies",       // Equity, Debt
	'I': "Indices",          // Equity, Debt
	'N': "Interest rates",   // Equity
	'M': "Misc.",            // Equity, Debt
}

var DebtInstrumentDependency = map[rune]string{
	'_': "Instrument Dependency",
	'B': "Bonds",
	'C': "Convertible Bonds",
	'W': "Bonds with Warrants Attached",
	'T': "Medium-Term Notes",
	'Y': "Money Market Instruments",
	'G': "Mortgage-Backed Securities",
	'Q': "Asset-Backed Securities",
	'N': "Municipal Bonds",
	'M': "Misc.",
}

var EquityVotingRight = map[rune]string{
	'_': "Voting Right",
	'V': "Voting",
	'N': "Non-Voting",
	'R': "Restricted",
	'E': "Enhanced voting",
}

var EquityPaymentStatus = map[rune]string{
	'_': "Payment Status",
	'F': "Fully Paid",
	'O': "Nil Paid",
	'P': "Partly Paid",
}

var EquityRedemption = map[rune]string{
	'_': "Redemption",
	'R': "Redeemable",
	'E': "Extendible",
	'T': "Redeemable / Extendible",
	'G': "Exchangeable",
	'A': "Redeemable / Exchangeable / Extendible",
	'C': "Redeemable/Exchangeable",
	'N': "Perpetual",
}

var EquityIncome = map[rune]string{
	'_': "Income",
	'F': "Fixed Rate",
	'C': "Cumulative, Fixed Rate",
	'P': "Participating",
	'Q': "Cumulative, Participating",
	'A': "Adjustable/Variable Rate",
	'N': "Normal Rate",
	'U': "Auction Rate",
	'D': "Dividends",
}

var categories = map[rune]category{
	'C': {
		Name: "Fund", //"Collective Investment Vehicles",
		Groups: map[rune]group{
			'I': {Name: "Standard"},
			'H': {Name: "Hedge fund"},
			'B': {Name: "Real estate investment trusts"},
			'E': {Name: "ETF"},
			'S': {Name: "Pension funds"},
			'F': {Name: "Fund of funds"},
			'P': {Name: "PE Fund"},
			'M': {Name: "Misc."},
		},
	},
	'D': {
		Name: "Debt",
		Groups: map[rune]group{
			'B': {Name: "Bond",
				Attr1: DebtTypeOfInterest,
				Attr2: DebtGuarantee,
				Attr3: DebtRedemptionReimbursement,
				Attr4: Form,
			},
			'C': {Name: "Convertible Bonds",
				Attr1: DebtTypeOfInterest,
				Attr2: DebtGuarantee,
				Attr3: DebtRedemptionReimbursement,
				Attr4: Form},
			'W': {Name: "Bonds with warrants attached",
				Attr1: DebtTypeOfInterest,
				Attr2: DebtGuarantee,
				Attr3: DebtRedemptionReimbursement,
				Attr4: Form},
			'T': {Name: "Medium-term notes",
				Attr1: DebtTypeOfInterest,
				Attr2: DebtGuarantee,
				Attr3: DebtRedemptionReimbursement,
				Attr4: Form},
			'S': {Name: "Structured products (with capital protection)",
				Attr1: map[rune]string{
					'_': "Type",
					'A': "Capital Protection Certificate with Participation",
					'B': "Capital Protection Convertible Certificate",
					'C': "Barrier Capital Protection Certificate",
					'D': "Capital Protection Certificate with Coupons",
					'M': "Misc.",
				},
				Attr2: DebtDistribution,
				Attr3: map[rune]string{
					'_': "Repayment",
					'F': "Fixed Cash Repayment (Only Protected Capital Level)",
					'V': "Variable Cash Repayment",
					'M': "Misc.",
				},
				Attr4: UnderlyingAsset,
			},
			'E': {Name: "Structured products (without capital protection)",
				Attr1: map[rune]string{
					'_': "Type",
					'A': "Discount Certificate",
					'B': "Barrier Discount Certificate",
					'C': "Reverse Convertible",
					'D': "Barrier Reverse Convertible",
					'E': "Express Certificate",
					'M': "Misc.",
				},
				Attr2: DebtDistribution,
				Attr3: map[rune]string{
					'_': "Repayment",
					'R': "Repayment in Cash",
					'S': "Repayment in Assets",
					'C': "Repayment in Assets and Cash",
					'T': "Repayment in Assets or Cash",
					'M': "Misc.",
				},
				Attr4: UnderlyingAsset,
			},
			'G': {Name: "Mortgage-backed securities",
				Attr1: DebtTypeOfInterest,
				Attr2: DebtGuarantee,
				Attr3: DebtRedemptionReimbursement,
				Attr4: Form,
			},
			'A': {Name: "Asset backed securities",
				Attr1: DebtTypeOfInterest,
				Attr2: DebtGuarantee,
				Attr3: DebtRedemptionReimbursement,
				Attr4: Form},
			'N': {Name: "Municipal bonds",
				Attr1: DebtTypeOfInterest,
				Attr2: DebtGuarantee,
				Attr3: DebtRedemptionReimbursement,
				Attr4: Form},
			'D': {Name: "Depository receipts",
				Attr1: DebtInstrumentDependency,
				Attr2: DebtTypeOfInterest,
				Attr3: DebtGuarantee,
				Attr4: DebtRedemptionReimbursement,
			},
			'Y': {Name: "Money market instruments",
				Attr1: DebtTypeOfInterest,
				Attr2: DebtGuarantee,
				Attr3: NA,
				Attr4: Form,
			},
			'M': {Name: "Misc.",
				Attr1: map[rune]string{
					'_': "Type",
					'B': "Bank Loan",
					'P': "Promissory Note",
					'M': "Misc.",
				},
				Attr2: NA,
				Attr3: NA,
				Attr4: Form,
			},
		},
	},
	'E': {
		Name: "Equity",
		Groups: map[rune]group{
			'S': {Name: "Shares",
				Attr1: EquityVotingRight,
				Attr2: EquityOwnership,
				Attr3: EquityPaymentStatus,
				Attr4: Form,
			},
			'P': {Name: "Preferred Shares",
				Attr1: EquityVotingRight,
				Attr2: EquityRedemption,
				Attr3: EquityIncome,
				Attr4: Form,
			},
			'C': {Name: "Conv. shares",
				Attr1: EquityVotingRight,
				Attr2: EquityOwnership,
				Attr3: EquityPaymentStatus,
				Attr4: Form},

			'F': {Name: "Preferred conv. shares",
				Attr1: EquityVotingRight,
				Attr2: EquityRedemption,
				Attr3: EquityIncome,
				Attr4: Form},
			'L': {Name: "Limited partnership units",
				Attr1: EquityVotingRight,
				Attr2: EquityOwnership,
				Attr3: EquityPaymentStatus,
				Attr4: Form},
			'D': {Name: "Depository receipts",
				Attr1: map[rune]string{
					'_': "Instrument dependency",
					'S': "Common/Ordinary Shares",
					'P': "Preferred/Preference Shares",
					'C': "Common/Ordinary Convertible Shares",
					'F': "Preferred/Preference Convertible Shares",
					'L': "Limited Partnership Units",
					'M': "Misc.",
				},
				Attr2: map[rune]string{
					'_': "Redemption",
					'R': "Redeemable",
					'N': "Perpetual",
					'B': "Convertible",
					'D': "Convertible/Redeemable",
					'X': "N/A",
				},
				Attr3: EquityIncome,
				Attr4: Form,
			},
			'Y': {Name: "Structured instruments",
				Attr1: map[rune]string{
					'_': "Type",
					'A': "Tracker Certificate",
					'B': "Outperforming Certificate",
					'C': "Bonus Certificate",
					'D': "Outperformance Bonus Certificate",
					'E': "Twin-Win-Certificate",
					'M': "Misc.",
				},
				Attr2: DebtDistribution, // Subset of the same
				Attr3: map[rune]string{
					'_': "Repayment",
					'F': "Cash Repayment",
					'V': "Physical Repayment",
					'E': "Elect at Settlement",
					'M': "Misc.",
				},
				Attr4: UnderlyingAsset,
			},
			'R': {Name: "Preference shares",
				Attr1: EquityVotingRight,
				Attr2: EquityRedemption,
				Attr3: EquityIncome,
				Attr4: Form},
			'V': {Name: "Preference convertibles shares",
				Attr1: EquityVotingRight,
				Attr2: EquityRedemption,
				Attr3: EquityIncome,
				Attr4: Form,
			},
			'U': {Name: "Unit",
				Attr1: map[rune]string{
					'_': "Closed/open-end",
					'C': "Closed-end",
					'O': "Open-end",
				},
				Attr2: map[rune]string{
					'_': "Distribution policy",
					'I': "Income funds",
					'G': "Growth funds",
					'M': "Mixed funds",
				},
				Attr3: map[rune]string{
					'_': "Assets",
					'R': "Real estate",
					'S': "Securities",
					'M': "Mixed-general",
					'C': "Commodities",
					'D': "Derivatives",
				},
				Attr4: Form,
			},

			'M': {Name: "Misc.",
				Attr1: NA,
				Attr2: NA,
				Attr3: NA,
				Attr4: Form},
		},
	},
	'F': {
		Name: "Future",
		Groups: map[rune]group{
			'F': {Name: "Financial"},
			'C': {Name: "Commodities"},
		},
	},
	'H': {
		Name: "Non-listed and complex listed options",
		Groups: map[rune]group{
			'R': {Name: "Rates"},
			'T': {Name: "Commodities"},
			'E': {Name: "Equity"},
			'C': {Name: "Credit"},
			'F': {Name: "FX"},
			'M': {Name: "Misc"},
		},
	},
	'I': {
		Name: "Spots",
		Groups: map[rune]group{
			'F': {Name: "FX",
				Attr1: NA,
				Attr2: NA,
				Attr3: NA,
				Attr4: map[rune]string{
					'P': "Physical",
				},
			},
			'T': {Name: "Commodities",
				Attr1: map[rune]string{
					'A': "Agriculture",
					'J': "Energy",
					'K': "Metals",
					'N': "Environmental",
					'P': "Polypropylene Products",
					'S': "Fertilizer",
					'T': "Paper",
					'M': "Misc.",
				},
				Attr2: NA,
				Attr3: NA,
				Attr4: NA,
			},
		},
	},
	'J': {
		Name: "Forwards",
		Groups: map[rune]group{
			'E': {Name: "Equity",
				Attr2: NA,
			},
			'F': {Name: "FX",
				Attr2: NA,
			},
			'C': {Name: "Credit",
				Attr2: NA,
			},
			'R': {Name: "Rates",
				Attr2: NA,
			},
			'T': {Name: "Commodities",
				Attr2: NA,
			},
		},
	},
	'K': {
		Name: "Strategies",
		Groups: map[rune]group{
			'E': {Name: "Equity"},
			'F': {Name: "FX"},
			'C': {Name: "Credit"},
			'R': {Name: "Rates"},
			'T': {Name: "Commodities"},
			'Y': {Name: "Mixed"},
			'M': {Name: "Misc."},
		},
	},
	'L': {
		Name: "Financing",
		Groups: map[rune]group{
			'L': {Name: "Loan Lease"},
			'R': {Name: "Repurchase agreements"},
			'S': {Name: "Securities Lending"},
		},
	},
	'M': {
		Name: "Misc.",
		Groups: map[rune]group{
			'C': {Name: "Combined instruments"},
			'M': {Name: "Misc."},
		},
	},
	'O': {
		Name: "Listed Options",
		Groups: map[rune]group{
			'C': {Name: "Call Option"},
			'P': {Name: "Put Option"},
			'M': {Name: "Misc"},
		},
	},
	'R': {
		Name: "Entitlement (Rights)",
		Groups: map[rune]group{
			'A': {Name: "Allotments (Bonus Rights)", Attr1: NA,
				Attr2: NA,
				Attr3: NA,
				Attr4: Form,
			},
			'S': {Name: "Subscription Rights",
				Attr4: Form},
			'P': {Name: "Purchase Rights",
				Attr4: Form},
			'W': {Name: "Warrants",
				Attr4: map[rune]string{
					'A': "American",
					'E': "European",
					'B': "Bermudan",
					'M': "Misc.",
				}},
			'F': {Name: "Mini-future",
				Attr4: map[rune]string{
					'A': "American",
					'E': "European",
					'B': "Bermudan",
					'M': "Misc.",
				}},
			'D': {Name: "Depository receipts",
				Attr4: Form},
			'M': {Name: "Misc.",
				Attr1: NA,
				Attr2: NA,
				Attr3: NA,
				Attr4: NA},
		},
	},
	'S': {
		Name: "Swap",
		Groups: map[rune]group{
			'R': {Name: "Rates"},
			'T': {Name: "Commodities"},
			'E': {Name: "Equity"},
			'C': {Name: "Credit"},
			'F': {Name: "FX"},
			'M': {Name: "Misc."},
		},
	},
	'T': {
		Name: "Referential instruments",
		Groups: map[rune]group{
			'C': {Name: "Currencies"},
			'T': {Name: "Commodities"},
			'R': {Name: "Interest rates"},
			'I': {Name: "Indexes"},
			'B': {Name: "Baskets"},
			'D': {Name: "Stock dividends"},
			'M': {Name: "Misc."},
		},
	},
}

type category struct {
	Name   string
	Groups map[rune]group
}

type group struct {
	Name  string
	Attr1 map[rune]string
	Attr2 map[rune]string
	Attr3 map[rune]string
	Attr4 map[rune]string
}

type CFI struct {
	tag string
}

func (c CFI) String() string {
	return c.tag
}
func full(tag_ string) (category_, group_, attr1, attr2, attr3, attr4 string) {
	tag := []rune(tag_)
	cat := categories[tag[0]]
	gro := group{}
	if cat.Groups != nil {
		gro = cat.Groups[tag[1]]
	}
	if gro.Attr1 == nil {
		gro.Attr1 = map[rune]string{}
	}
	if gro.Attr2 == nil {
		gro.Attr2 = map[rune]string{}
	}
	if gro.Attr3 == nil {
		gro.Attr3 = map[rune]string{}
	}
	if gro.Attr4 == nil {
		gro.Attr4 = map[rune]string{}
	}

	return cat.Name,
		gro.Name,
		gro.Attr1[tag[2]],
		gro.Attr2[tag[3]],
		gro.Attr3[tag[4]],
		gro.Attr4[tag[5]]

}

func (c CFI) Tag() string {
	return c.tag
}

func (c CFI) Category() string {
	return categories[rune(c.tag[0])].Name
}
func (c CFI) Group() string {
	cat := categories[rune(c.tag[0])]
	gro := group{}
	if cat.Groups != nil {
		gro = cat.Groups[rune(c.tag[1])]
	}
	return gro.Name
}

type Fmt int

const (
	Tag = iota
	Short
	Long
)

func (c CFI) Format(f Fmt) string {

	tag := []rune(c.tag)
	cat := categories[tag[0]]
	gro := group{}
	if cat.Groups != nil {
		gro = cat.Groups[tag[1]]
	}
	if gro.Attr1 == nil {
		gro.Attr1 = map[rune]string{}
	}
	if gro.Attr2 == nil {
		gro.Attr2 = map[rune]string{}
	}
	if gro.Attr3 == nil {
		gro.Attr3 = map[rune]string{}
	}
	if gro.Attr4 == nil {
		gro.Attr4 = map[rune]string{}
	}

	switch f {
	case Short:
		return fmt.Sprintf(`%s; %s; %s; %s; %s; %s`,
			cat.Name,
			gro.Name,
			gro.Attr1[tag[1+1]],
			gro.Attr2[tag[1+2]],
			gro.Attr3[tag[1+3]],
			gro.Attr4[tag[1+4]])
	case Long:
		return fmt.Sprintf(`Catagory: %s
 Group: %s
  %s: %s 
  %s: %s 
  %s: %s
  %s: %s`, cat.Name,
			gro.Name,
			gro.Attr1['_'], gro.Attr1[tag[1+1]],
			gro.Attr2['_'], gro.Attr2[tag[1+2]],
			gro.Attr3['_'], gro.Attr3[tag[1+3]],
			gro.Attr4['_'], gro.Attr4[tag[1+4]])
	default: // Tag
		return c.tag

	}
}
