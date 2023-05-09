// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

func (dst *ContactInfo) SetFields(src *ContactInfo, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "contact_type":
			if len(subs) > 0 {
				return fmt.Errorf("'contact_type' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ContactType = src.ContactType
			} else {
				dst.ContactType = 0
			}
		case "contact_method":
			if len(subs) > 0 {
				return fmt.Errorf("'contact_method' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ContactMethod = src.ContactMethod
			} else {
				dst.ContactMethod = 0
			}
		case "value":
			if len(subs) > 0 {
				return fmt.Errorf("'value' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Value = src.Value
			} else {
				var zero string
				dst.Value = zero
			}
		case "public":
			if len(subs) > 0 {
				return fmt.Errorf("'public' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Public = src.Public
			} else {
				var zero bool
				dst.Public = zero
			}
		case "validated_at":
			if len(subs) > 0 {
				return fmt.Errorf("'validated_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ValidatedAt = src.ValidatedAt
			} else {
				dst.ValidatedAt = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ContactInfoValidation) SetFields(src *ContactInfoValidation, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "id":
			if len(subs) > 0 {
				return fmt.Errorf("'id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Id = src.Id
			} else {
				var zero string
				dst.Id = zero
			}
		case "token":
			if len(subs) > 0 {
				return fmt.Errorf("'token' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Token = src.Token
			} else {
				var zero string
				dst.Token = zero
			}
		case "entity":
			if len(subs) > 0 {
				var newDst, newSrc *EntityIdentifiers
				if (src == nil || src.Entity == nil) && dst.Entity == nil {
					continue
				}
				if src != nil {
					newSrc = src.Entity
				}
				if dst.Entity != nil {
					newDst = dst.Entity
				} else {
					newDst = &EntityIdentifiers{}
					dst.Entity = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Entity = src.Entity
				} else {
					dst.Entity = nil
				}
			}
		case "contact_info":
			if len(subs) > 0 {
				return fmt.Errorf("'contact_info' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ContactInfo = src.ContactInfo
			} else {
				dst.ContactInfo = nil
			}
		case "created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreatedAt = src.CreatedAt
			} else {
				dst.CreatedAt = nil
			}
		case "expires_at":
			if len(subs) > 0 {
				return fmt.Errorf("'expires_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ExpiresAt = src.ExpiresAt
			} else {
				dst.ExpiresAt = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}