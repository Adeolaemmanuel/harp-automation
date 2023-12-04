package graphql

func (q Queries) GetStoredCamping() string {
	return ``
}

func (q Queries) GetHarpPeople() string {
	return `query getHarpPeople {
		harp_people {
		  id
		  metadata
		  created_at
		  business_id
		  creator
		  email
		  first_name
		  last_name
		  phone
		  type
		  updated_at
		}
	  }`
}

func (q Queries) GetHarpCampaign() string {
	return `query getHarpCampaign {
		harp_campaigns {
		  business_id
		  channel
		  created_at
		  id
		  message
		  status
		  title
		  recipients
		  batches
		  metadata
		}
	  }`
}