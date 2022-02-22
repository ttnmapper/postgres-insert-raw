package database

func GetPeeredNetworks(networkId string) []PacketBrokerRoutingPolicy {
	var routingPolicies []PacketBrokerRoutingPolicy
	Db.Where("home_network_id = ?", networkId).
		Where("uplink_application_data = true").
		Find(&routingPolicies)
	return routingPolicies
}
