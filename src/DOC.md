# Node

# Routing Table

# Bucket
All methods will throw return an error if they encouncter anything in the bucket that is not a contact.

#### (b *bucket) AddContact(target contact) error
Attempts to add the contact to buckets content, if the contact already exists in the bucket it is moved to the back. Returns an error if the bucket is full and the contact does not already exist in it.

#### (b *bucket) FindContact(targetID [5]uint32) (contact, error)
Searches the bucket for the target ID, if a contact with matching ID is found it is returned. Returns an error if no match is found.

#### (b *bucket) FindXClosest(x int, targetID [5]uint32) (*list.list, error)
Searches the bucket and returns up to the x closest contacts to the target ID in a list. If less than x contacts are found a "incomplete" error is returned alongside the list containing the found contacts.

# Contact

# ID
Node ID's consist of 160 bits, they are stored as five uint32 stored in an array.
