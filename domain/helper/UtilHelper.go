package helper

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/net/idna"
)

func ParsePrimitive[T any](s string) (T, error) {
	var zero T
	switch any(zero).(type) {
	case string:
		return any(s).(T), nil
	case bool:
		v, err := strconv.ParseBool(s)
		return any(v).(T), err
	case int:
		v, err := strconv.Atoi(s)
		return any(v).(T), err
	case int8:
		v, err := strconv.ParseInt(s, 10, 8)
		return any(int8(v)).(T), err
	case int16:
		v, err := strconv.ParseInt(s, 10, 16)
		return any(int16(v)).(T), err
	case int32:
		v, err := strconv.ParseInt(s, 10, 32)
		return any(int32(v)).(T), err
	case int64:
		v, err := strconv.ParseInt(s, 10, 64)
		return any(v).(T), err
	case uint:
		v, err := strconv.ParseUint(s, 10, 0)
		return any(uint(v)).(T), err
	case uint8:
		v, err := strconv.ParseUint(s, 10, 8)
		return any(uint8(v)).(T), err
	case uint16:
		v, err := strconv.ParseUint(s, 10, 16)
		return any(uint16(v)).(T), err
	case uint32:
		v, err := strconv.ParseUint(s, 10, 32)
		return any(uint32(v)).(T), err
	case uint64:
		v, err := strconv.ParseUint(s, 10, 64)
		return any(v).(T), err

	case float32:
		v, err := strconv.ParseFloat(s, 32)
		return any(float32(v)).(T), err
	case float64:
		v, err := strconv.ParseFloat(s, 64)
		return any(v).(T), err
	}
	return zero, fmt.Errorf("unsupported type")
}

func GetTypeName[T any]() string {
	var zero T
	return reflect.TypeOf(zero).String()
}

func ValidateEmail(ctx context.Context, email string, checkMX bool) (bool, error) {
	email = strings.TrimSpace(email)
	if email == "" {
		return false, errors.New("email is empty")
	}
	if len(email) > 254 {
		return false, errors.New("email exceeds maximum length (254)")
	}

	addr, err := mail.ParseAddress(email)
	if err != nil {
		return false, fmt.Errorf("invalid address syntax: %w", err)
	}
	// Chỉ chấp nhận plain email, không chấp nhận "Name <email>" format
	if addr.Address != email {
		return false, errors.New("input must be a plain email address")
	}

	parts := strings.SplitN(addr.Address, "@", 2)
	if len(parts) != 2 {
		return false, errors.New("missing @ separator")
	}
	local, domain := parts[0], parts[1]

	if len(local) == 0 || len(local) > 64 {
		return false, errors.New("local part length invalid (1-64)")
	}

	asciiDomain, err := idna.Lookup.ToASCII(domain)
	if err != nil {
		return false, fmt.Errorf("invalid domain name: %w", err)
	}
	if len(asciiDomain) == 0 || len(asciiDomain) > 255 {
		return false, errors.New("domain part length invalid")
	}

	labels := strings.Split(asciiDomain, ".")
	if len(labels) < 2 {
		return false, errors.New("domain must contain a TLD")
	}
	for _, l := range labels {
		if l == "" {
			return false, errors.New("empty domain label")
		}
		if len(l) > 63 {
			return false, errors.New("domain label exceeds 63 characters")
		}
	}
	// TLD không được toàn số
	tld := labels[len(labels)-1]
	if _, atoiErr := strconv.Atoi(tld); atoiErr == nil {
		return false, errors.New("TLD must not be all numeric")
	}

	if checkMX {
		resolver := &net.Resolver{}
		mxRecords, mxErr := resolver.LookupMX(ctx, asciiDomain)
		if mxErr != nil || len(mxRecords) == 0 {
			hosts, hostErr := resolver.LookupHost(ctx, asciiDomain)
			if hostErr != nil || len(hosts) == 0 {
				var dnsErr *net.DNSError
				if errors.As(hostErr, &dnsErr) && dnsErr.IsTemporary {
					return false, fmt.Errorf("DNS lookup temporarily failed: %w", hostErr)
				}
				return false, errors.New("no MX or A/AAAA records found for domain")
			}
		}
	}
	return true, nil
}
