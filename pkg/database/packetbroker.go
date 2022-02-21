package database

func InsertOrUpdateRoutingPolicy(routingPolicy PacketBrokerRoutingPolicy) {
	// Find existing entry or create a new one
	routingPolicyDatabase := PacketBrokerRoutingPolicy{HomeNetworkId: routingPolicy.HomeNetworkId, ForwarderNetworkId: routingPolicy.ForwarderNetworkId}
	Db.FirstOrCreate(&routingPolicyDatabase, &routingPolicyDatabase)

	// Use existing entry's ID
	routingPolicy.ID = routingPolicyDatabase.ID

	// Update in database
	Db.Save(&routingPolicy)
}
