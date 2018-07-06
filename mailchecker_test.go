package mailchecker

import (
	"testing"
)

func assertIsValidResult(expected bool, email string, t *testing.T) {
	result := IsValid(email)
	if result != expected {
		t.Errorf("Evaluating: %s. Result: %v, Expected: %v", email, result, expected)
	}
}

func isValid(email string, t *testing.T) {
	assertIsValidResult(true, email, t)
}

func isInvalid(email string, t *testing.T) {
	assertIsValidResult(false, email, t)
}

func TestReturnTrueIfValidEmail(t *testing.T) {
	isValid("plop@plop.com", t)
	isValid("my.ok@ok.plop.com", t)
	isValid("my+ok@ok.plop.com", t)
	isValid("my=ok@ok.plop.com", t)
	isValid("ok@gmail.com", t)
	isValid("ok@hotmail.com", t)
}

func TestReturnFalseIfEmailInvalid(t *testing.T) {
	isInvalid("", t)
	isInvalid("  ", t)
	isInvalid("plopplop.com", t)
	isInvalid("my+ok@ok=plop.com", t)
	isInvalid("my,ok@ok.plop.com", t)
	isInvalid("  ok@gmail.com  ", t)
	isInvalid("  ok@gmail.com", t)
	isInvalid("ok@gmail.com  ", t)
	isInvalid("\nok@gmail.com\n", t)
	isInvalid("\nok@gmail.com", t)
	isInvalid("ok@gmail.com\n", t)
}

func TestReturnFalseIfThrowableDomain(t *testing.T) {
	isInvalid("ok@tmail.com", t)
	isInvalid("ok@33mail.com", t)
	// isInvalid('ok@ok.33mail.com');
	isInvalid("ok@guerrillamailblock.com", t)
}

func TestReturnFalseForBlacklistedDomainsAndTheirSubdomains(t *testing.T) {
	for _, blacklistedDomain := range Blacklist() {
		isInvalid("test@"+blacklisted_domain, t)
		isInvalid("test@subdomain."+blacklisted_domain, t)
		// Should not be invalid as a subdomain of a valid domain.
		isValid("test@"+blacklisted_domain+".gmail.com", t)
	}
}
