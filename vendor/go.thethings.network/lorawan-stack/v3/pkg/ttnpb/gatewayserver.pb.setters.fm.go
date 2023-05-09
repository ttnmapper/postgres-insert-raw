// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import fmt "fmt"

func (dst *GatewayUp) SetFields(src *GatewayUp, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "uplink_messages":
			if len(subs) > 0 {
				return fmt.Errorf("'uplink_messages' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.UplinkMessages = src.UplinkMessages
			} else {
				dst.UplinkMessages = nil
			}
		case "gateway_status":
			if len(subs) > 0 {
				var newDst, newSrc *GatewayStatus
				if (src == nil || src.GatewayStatus == nil) && dst.GatewayStatus == nil {
					continue
				}
				if src != nil {
					newSrc = src.GatewayStatus
				}
				if dst.GatewayStatus != nil {
					newDst = dst.GatewayStatus
				} else {
					newDst = &GatewayStatus{}
					dst.GatewayStatus = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.GatewayStatus = src.GatewayStatus
				} else {
					dst.GatewayStatus = nil
				}
			}
		case "tx_acknowledgment":
			if len(subs) > 0 {
				var newDst, newSrc *TxAcknowledgment
				if (src == nil || src.TxAcknowledgment == nil) && dst.TxAcknowledgment == nil {
					continue
				}
				if src != nil {
					newSrc = src.TxAcknowledgment
				}
				if dst.TxAcknowledgment != nil {
					newDst = dst.TxAcknowledgment
				} else {
					newDst = &TxAcknowledgment{}
					dst.TxAcknowledgment = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.TxAcknowledgment = src.TxAcknowledgment
				} else {
					dst.TxAcknowledgment = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *GatewayDown) SetFields(src *GatewayDown, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "downlink_message":
			if len(subs) > 0 {
				var newDst, newSrc *DownlinkMessage
				if (src == nil || src.DownlinkMessage == nil) && dst.DownlinkMessage == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkMessage
				}
				if dst.DownlinkMessage != nil {
					newDst = dst.DownlinkMessage
				} else {
					newDst = &DownlinkMessage{}
					dst.DownlinkMessage = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkMessage = src.DownlinkMessage
				} else {
					dst.DownlinkMessage = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ScheduleDownlinkResponse) SetFields(src *ScheduleDownlinkResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "delay":
			if len(subs) > 0 {
				return fmt.Errorf("'delay' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Delay = src.Delay
			} else {
				dst.Delay = nil
			}
		case "downlink_path":
			if len(subs) > 0 {
				var newDst, newSrc *DownlinkPath
				if (src == nil || src.DownlinkPath == nil) && dst.DownlinkPath == nil {
					continue
				}
				if src != nil {
					newSrc = src.DownlinkPath
				}
				if dst.DownlinkPath != nil {
					newDst = dst.DownlinkPath
				} else {
					newDst = &DownlinkPath{}
					dst.DownlinkPath = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.DownlinkPath = src.DownlinkPath
				} else {
					dst.DownlinkPath = nil
				}
			}
		case "rx1":
			if len(subs) > 0 {
				return fmt.Errorf("'rx1' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Rx1 = src.Rx1
			} else {
				var zero bool
				dst.Rx1 = zero
			}
		case "rx2":
			if len(subs) > 0 {
				return fmt.Errorf("'rx2' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Rx2 = src.Rx2
			} else {
				var zero bool
				dst.Rx2 = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *ScheduleDownlinkErrorDetails) SetFields(src *ScheduleDownlinkErrorDetails, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "path_errors":
			if len(subs) > 0 {
				return fmt.Errorf("'path_errors' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.PathErrors = src.PathErrors
			} else {
				dst.PathErrors = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *BatchGetGatewayConnectionStatsRequest) SetFields(src *BatchGetGatewayConnectionStatsRequest, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "gateway_ids":
			if len(subs) > 0 {
				return fmt.Errorf("'gateway_ids' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.GatewayIds = src.GatewayIds
			} else {
				dst.GatewayIds = nil
			}
		case "field_mask":
			if len(subs) > 0 {
				return fmt.Errorf("'field_mask' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.FieldMask = src.FieldMask
			} else {
				dst.FieldMask = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *BatchGetGatewayConnectionStatsResponse) SetFields(src *BatchGetGatewayConnectionStatsResponse, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "entries":
			if len(subs) > 0 {
				return fmt.Errorf("'entries' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Entries = src.Entries
			} else {
				dst.Entries = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}