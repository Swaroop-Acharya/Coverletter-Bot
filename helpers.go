package main

import "strings"

func parseList(input string) []string {
	parts := strings.Split(input, ";")
	var result []string
	for _, part := range parts {
		v := strings.TrimSpace(part)
		if v != "" {
			result = append(result, v)
		}

	}
	return result
}

func (cld *coverletterData) isEmpty() bool {
	return cld.TargetCompany == "" &&
		cld.TargetRole == "" &&
		cld.UserName == "" &&
		cld.UserEmail == ""
}

func (cld *coverletterData) formatPreview() string {
	b := strings.Builder{}

	b.WriteString("📄 *Cover Letter Preview*\n\n")

	if cld.UserName != "" {
		b.WriteString("👤 *Name:* " + cld.UserName + "\n")
	}
	if cld.UserEmail != "" {
		b.WriteString("📧 *Email:* " + cld.UserEmail + "\n")
	}
	if cld.UserPhoneNumber != "" {
		b.WriteString("📞 *Phone:* " + cld.UserPhoneNumber + "\n")
	}

	b.WriteString("\n")

	if cld.TargetCompany != "" {
		b.WriteString("🏢 *Target Company:* " + cld.TargetCompany + "\n")
	}
	if cld.TargetRole != "" {
		b.WriteString("💼 *Target Role:* " + cld.TargetRole + "\n")
	}
	if cld.CurrentCompany != "" {
		b.WriteString("🏬 *Current Company:* " + cld.CurrentCompany + "\n")
	}
	if cld.CurrentCompany != "" {
		b.WriteString("🏬 *Current Role:* " + cld.CurrentRole + "\n")
	}

	if len(cld.Skills) > 0 {
		b.WriteString("\n🛠 *Skills:*\n")
		for _, s := range cld.Skills {
			b.WriteString("• " + s + "\n")
		}
	}

	if cld.Intro != "" {
		b.WriteString("\n📝 *Intro:*\n" + cld.Intro + "\n")
	}

	if cld.Closing != "" {
		b.WriteString("\n✍️ *Closing:*\n" + cld.Closing + "\n")
	}

	return b.String()
}
