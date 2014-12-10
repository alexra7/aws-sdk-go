// Package kinesis provides a client for Amazon Kinesis.
package kinesis

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

// Kinesis is a client for Amazon Kinesis.
type Kinesis struct {
	client *aws.JSONClient
}

// New returns a new Kinesis client.
func New(creds aws.Credentials, region string, client *http.Client) *Kinesis {
	if client == nil {
		client = http.DefaultClient
	}

	service := "kinesis"
	endpoint, service, region := endpoints.Lookup("kinesis", region)

	return &Kinesis{
		client: &aws.JSONClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			}, Client: client,
			Endpoint:     endpoint,
			JSONVersion:  "1.1",
			TargetPrefix: "Kinesis_20131202",
		},
	}
}

// AddTagsToStream adds or updates tags for the specified Amazon Kinesis
// stream. Each stream can have up to 10 tags. If tags have already been
// assigned to the stream, AddTagsToStream overwrites any existing tags
// that correspond to the specified tag keys.
func (c *Kinesis) AddTagsToStream(req AddTagsToStreamInput) (err error) {
	// NRE
	err = c.client.Do("AddTagsToStream", "POST", "/", req, nil)
	return
}

// CreateStream creates a Amazon Kinesis stream. A stream captures and
// transports data records that are continuously emitted from different
// data sources or producers . Scale-out within an Amazon Kinesis stream is
// explicitly supported by means of shards, which are uniquely identified
// groups of data records in an Amazon Kinesis stream. You specify and
// control the number of shards that a stream is composed of. Each open
// shard can support up to 5 read transactions per second, up to a maximum
// total of 2 MB of data read per second. Each shard can support up to 1000
// records written per second, up to a maximum total of 1 MB data written
// per second. You can add shards to a stream if the amount of data input
// increases and you can remove shards if the amount of data input
// decreases. The stream name identifies the stream. The name is scoped to
// the AWS account used by the application. It is also scoped by region.
// That is, two streams in two different accounts can have the same name,
// and two streams in the same account, but in two different regions, can
// have the same name. CreateStream is an asynchronous operation. Upon
// receiving a CreateStream request, Amazon Kinesis immediately returns and
// sets the stream status to . After the stream is created, Amazon Kinesis
// sets the stream status to . You should perform read and write operations
// only on an stream. You receive a LimitExceededException when making a
// CreateStream request if you try to do one of the following: Have more
// than five streams in the state at any point in time. Create more shards
// than are authorized for your account. The default limit for an AWS
// account is 10 shards per stream. If you need to create a stream with
// more than 10 shards, contact AWS Support to increase the limit on your
// account. You can use DescribeStream to check the stream status, which is
// returned in StreamStatus CreateStream has a limit of 5 transactions per
// second per account.
func (c *Kinesis) CreateStream(req CreateStreamInput) (err error) {
	// NRE
	err = c.client.Do("CreateStream", "POST", "/", req, nil)
	return
}

// DeleteStream deletes a stream and all its shards and data. You must shut
// down any applications that are operating on the stream before you delete
// the stream. If an application attempts to operate on a deleted stream,
// it will receive the exception ResourceNotFoundException If the stream is
// in the state, you can delete it. After a DeleteStream request, the
// specified stream is in the state until Amazon Kinesis completes the
// deletion. Note: Amazon Kinesis might continue to accept data read and
// write operations, such as PutRecord , PutRecords , and GetRecords , on a
// stream in the state until the stream deletion is complete. When you
// delete a stream, any shards in that stream are also deleted, and any
// tags are dissociated from the stream. You can use the DescribeStream
// operation to check the state of the stream, which is returned in
// StreamStatus DeleteStream has a limit of 5 transactions per second per
// account.
func (c *Kinesis) DeleteStream(req DeleteStreamInput) (err error) {
	// NRE
	err = c.client.Do("DeleteStream", "POST", "/", req, nil)
	return
}

// DescribeStream describes the specified stream. The information about the
// stream includes its current status, its Amazon Resource Name and an
// array of shard objects. For each shard object, there is information
// about the hash key and sequence number ranges that the shard spans, and
// the IDs of any earlier shards that played in a role in creating the
// shard. A sequence number is the identifier associated with every record
// ingested in the Amazon Kinesis stream. The sequence number is assigned
// when a record is put into the stream. You can limit the number of
// returned shards using the Limit parameter. The number of shards in a
// stream may be too large to return from a single call to DescribeStream .
// You can detect this by using the HasMoreShards flag in the returned
// output. HasMoreShards is set to true when there is more data available.
// DescribeStream is a paginated operation. If there are more shards
// available, you can request them using the shard ID of the last shard
// returned. Specify this ID in the ExclusiveStartShardId parameter in a
// subsequent request to DescribeStream . DescribeStream has a limit of 10
// transactions per second per account.
func (c *Kinesis) DescribeStream(req DescribeStreamInput) (resp *DescribeStreamOutput, err error) {
	resp = &DescribeStreamOutput{}
	err = c.client.Do("DescribeStream", "POST", "/", req, resp)
	return
}

// GetRecords gets data records from a shard. Specify a shard iterator
// using the ShardIterator parameter. The shard iterator specifies the
// position in the shard from which you want to start reading data records
// sequentially. If there are no records available in the portion of the
// shard that the iterator points to, GetRecords returns an empty list.
// Note that it might take multiple calls to get to a portion of the shard
// that contains records. You can scale by provisioning multiple shards.
// Your application should have one thread per shard, each reading
// continuously from its stream. To read from a stream continually, call
// GetRecords in a loop. Use GetShardIterator to get the shard iterator to
// specify in the first GetRecords call. GetRecords returns a new shard
// iterator in NextShardIterator . Specify the shard iterator returned in
// NextShardIterator in subsequent calls to GetRecords . Note that if the
// shard has been closed, the shard iterator can't return more data and
// GetRecords returns null in NextShardIterator . You can terminate the
// loop when the shard is closed, or when the shard iterator reaches the
// record with the sequence number or other attribute that marks it as the
// last record to process. Each data record can be up to 50 KB in size, and
// each shard can read up to 2 MB per second. You can ensure that your
// calls don't exceed the maximum supported size or throughput by using the
// Limit parameter to specify the maximum number of records that GetRecords
// can return. Consider your average record size when determining this
// limit. For example, if your average record size is 40 KB, you can limit
// the data returned to about 1 MB per call by specifying 25 as the limit.
// The size of the data returned by GetRecords will vary depending on the
// utilization of the shard. The maximum size of data that GetRecords can
// return is 10 MB. If a call returns 10 MB of data, subsequent calls made
// within the next 5 seconds throw ProvisionedThroughputExceededException .
// If there is insufficient provisioned throughput on the shard, subsequent
// calls made within the next 1 second throw
// ProvisionedThroughputExceededException . Note that GetRecords won't
// return any data when it throws an exception. For this reason, we
// recommend that you wait one second between calls to GetRecords ;
// however, it's possible that the application will get exceptions for
// longer than 1 second. To detect whether the application is falling
// behind in processing, add a timestamp to your records and note how long
// it takes to process them. You can also monitor how much data is in a
// stream using the CloudWatch metrics for write operations PutRecord and
// PutRecords ). For more information, see Monitoring Amazon Kinesis with
// Amazon CloudWatch in the Amazon Kinesis Developer Guide
func (c *Kinesis) GetRecords(req GetRecordsInput) (resp *GetRecordsOutput, err error) {
	resp = &GetRecordsOutput{}
	err = c.client.Do("GetRecords", "POST", "/", req, resp)
	return
}

// GetShardIterator gets a shard iterator. A shard iterator expires five
// minutes after it is returned to the requester. A shard iterator
// specifies the position in the shard from which to start reading data
// records sequentially. A shard iterator specifies this position using the
// sequence number of a data record in a shard. A sequence number is the
// identifier associated with every record ingested in the Amazon Kinesis
// stream. The sequence number is assigned when a record is put into the
// stream. You must specify the shard iterator type. For example, you can
// set the ShardIteratorType parameter to read exactly from the position
// denoted by a specific sequence number by using the shard iterator type,
// or right after the sequence number by using the shard iterator type,
// using sequence numbers returned by earlier calls to PutRecord ,
// PutRecords , GetRecords , or DescribeStream . You can specify the shard
// iterator type in the request to cause ShardIterator to point to the last
// untrimmed record in the shard in the system, which is the oldest data
// record in the shard. Or you can point to just after the most recent
// record in the shard, by using the shard iterator type , so that you
// always read the most recent data in the shard. When you repeatedly read
// from an Amazon Kinesis stream use a GetShardIterator request to get the
// first shard iterator to to use in your first GetRecords request and then
// use the shard iterator returned by the GetRecords request in
// NextShardIterator for subsequent reads. A new shard iterator is returned
// by every GetRecords request in NextShardIterator , which you use in the
// ShardIterator parameter of the next GetRecords request. If a
// GetShardIterator request is made too often, you receive a
// ProvisionedThroughputExceededException . For more information about
// throughput limits, see GetRecords If the shard is closed, the iterator
// can't return more data, and GetShardIterator returns null for its
// ShardIterator . A shard can be closed using SplitShard or MergeShards
// GetShardIterator has a limit of 5 transactions per second per account
// per open shard.
func (c *Kinesis) GetShardIterator(req GetShardIteratorInput) (resp *GetShardIteratorOutput, err error) {
	resp = &GetShardIteratorOutput{}
	err = c.client.Do("GetShardIterator", "POST", "/", req, resp)
	return
}

// ListStreams lists your streams. The number of streams may be too large
// to return from a single call to ListStreams . You can limit the number
// of returned streams using the Limit parameter. If you do not specify a
// value for the Limit parameter, Amazon Kinesis uses the default limit,
// which is currently 10. You can detect if there are more streams
// available to list by using the HasMoreStreams flag from the returned
// output. If there are more streams available, you can request more
// streams by using the name of the last stream returned by the ListStreams
// request in the ExclusiveStartStreamName parameter in a subsequent
// request to ListStreams . The group of stream names returned by the
// subsequent request is then added to the list. You can continue this
// process until all the stream names have been collected in the list.
// ListStreams has a limit of 5 transactions per second per account.
func (c *Kinesis) ListStreams(req ListStreamsInput) (resp *ListStreamsOutput, err error) {
	resp = &ListStreamsOutput{}
	err = c.client.Do("ListStreams", "POST", "/", req, resp)
	return
}

// ListTagsForStream lists the tags for the specified Amazon Kinesis
// stream.
func (c *Kinesis) ListTagsForStream(req ListTagsForStreamInput) (resp *ListTagsForStreamOutput, err error) {
	resp = &ListTagsForStreamOutput{}
	err = c.client.Do("ListTagsForStream", "POST", "/", req, resp)
	return
}

// MergeShards merges two adjacent shards in a stream and combines them
// into a single shard to reduce the stream's capacity to ingest and
// transport data. Two shards are considered adjacent if the union of the
// hash key ranges for the two shards form a contiguous set with no gaps.
// For example, if you have two shards, one with a hash key range of
// 276...381 and the other with a hash key range of 382...454, then you
// could merge these two shards into a single shard that would have a hash
// key range of 276...454. After the merge, the single child shard receives
// data for all hash key values covered by the two parent shards.
// MergeShards is called when there is a need to reduce the overall
// capacity of a stream because of excess capacity that is not being used.
// You must specify the shard to be merged and the adjacent shard for a
// stream. For more information about merging shards, see Merge Two Shards
// in the Amazon Kinesis Developer Guide If the stream is in the state, you
// can call MergeShards . If a stream is in the , , or state, MergeShards
// returns a ResourceInUseException . If the specified stream does not
// exist, MergeShards returns a ResourceNotFoundException . You can use
// DescribeStream to check the state of the stream, which is returned in
// StreamStatus MergeShards is an asynchronous operation. Upon receiving a
// MergeShards request, Amazon Kinesis immediately returns a response and
// sets the StreamStatus to . After the operation is completed, Amazon
// Kinesis sets the StreamStatus to . Read and write operations continue to
// work while the stream is in the state. You use DescribeStream to
// determine the shard IDs that are specified in the MergeShards request.
// If you try to operate on too many streams in parallel using CreateStream
// , DeleteStream , MergeShards or SplitShard , you will receive a
// LimitExceededException . MergeShards has limit of 5 transactions per
// second per account.
func (c *Kinesis) MergeShards(req MergeShardsInput) (err error) {
	// NRE
	err = c.client.Do("MergeShards", "POST", "/", req, nil)
	return
}

// PutRecord puts (writes) a single data record from a producer into an
// Amazon Kinesis stream. Call PutRecord to send data from the producer
// into the Amazon Kinesis stream for real-time ingestion and subsequent
// processing, one record at a time. Each shard can support up to 1000
// records written per second, up to a maximum total of 1 MB data written
// per second. You must specify the name of the stream that captures,
// stores, and transports the data; a partition key; and the data blob
// itself. The data blob can be any type of data; for example, a segment
// from a log file, geographic/location data, website clickstream data, and
// so on. The partition key is used by Amazon Kinesis to distribute data
// across shards. Amazon Kinesis segregates the data records that belong to
// a data stream into multiple shards, using the partition key associated
// with each data record to determine which shard a given data record
// belongs to. Partition keys are Unicode strings, with a maximum length
// limit of 256 bytes. An MD5 hash function is used to map partition keys
// to 128-bit integer values and to map associated data records to shards
// using the hash key ranges of the shards. You can override hashing the
// partition key to determine the shard by explicitly specifying a hash
// value using the ExplicitHashKey parameter. For more information, see
// Partition Key in the Amazon Kinesis Developer Guide PutRecord returns
// the shard ID of where the data record was placed and the sequence number
// that was assigned to the data record. Sequence numbers generally
// increase over time. To guarantee strictly increasing ordering, use the
// SequenceNumberForOrdering parameter. For more information, see Sequence
// Number in the Amazon Kinesis Developer Guide If a PutRecord request
// cannot be processed because of insufficient provisioned throughput on
// the shard involved in the request, PutRecord throws
// ProvisionedThroughputExceededException . Data records are accessible for
// only 24 hours from the time that they are added to an Amazon Kinesis
// stream.
func (c *Kinesis) PutRecord(req PutRecordInput) (resp *PutRecordOutput, err error) {
	resp = &PutRecordOutput{}
	err = c.client.Do("PutRecord", "POST", "/", req, resp)
	return
}

// PutRecords puts (writes) multiple data records from a producer into an
// Amazon Kinesis stream in a single call (also referred to as a PutRecords
// request). Use this operation to send data from a data producer into the
// Amazon Kinesis stream for real-time ingestion and processing. Each shard
// can support up to 1000 records written per second, up to a maximum total
// of 1 MB data written per second. You must specify the name of the stream
// that captures, stores, and transports the data; and an array of request
// Records , with each record in the array requiring a partition key and
// data blob. The data blob can be any type of data; for example, a segment
// from a log file, geographic/location data, website clickstream data, and
// so on. The partition key is used by Amazon Kinesis as input to a hash
// function that maps the partition key and associated data to a specific
// shard. An MD5 hash function is used to map partition keys to 128-bit
// integer values and to map associated data records to shards. As a result
// of this hashing mechanism, all data records with the same partition key
// map to the same shard within the stream. For more information, see
// Partition Key in the Amazon Kinesis Developer Guide Each record in the
// Records array may include an optional parameter, ExplicitHashKey , which
// overrides the partition key to shard mapping. This parameter allows a
// data producer to determine explicitly the shard where the record is
// stored. For more information, see Adding Multiple Records with
// PutRecords in the Amazon Kinesis Developer Guide The PutRecords response
// includes an array of response Records . Each record in the response
// array directly correlates with a record in the request array using
// natural ordering, from the top to the bottom of the request and
// response. The response Records array always includes the same number of
// records as the request array. The response Records array includes both
// successfully and unsuccessfully processed records. Amazon Kinesis
// attempts to process all records in each PutRecords request. A single
// record failure does not stop the processing of subsequent records. A
// successfully-processed record includes ShardId and SequenceNumber
// values. The ShardId parameter identifies the shard in the stream where
// the record is stored. The SequenceNumber parameter is an identifier
// assigned to the put record, unique to all records in the stream. An
// unsuccessfully-processed record includes ErrorCode and ErrorMessage
// values. ErrorCode reflects the type of error and can be one of the
// following values: ProvisionedThroughputExceededException or
// InternalFailure . ErrorMessage provides more detailed information about
// the ProvisionedThroughputExceededException exception including the
// account ID, stream name, and shard ID of the record that was throttled.
// Data records are accessible for only 24 hours from the time that they
// are added to an Amazon Kinesis stream.
func (c *Kinesis) PutRecords(req PutRecordsInput) (resp *PutRecordsOutput, err error) {
	resp = &PutRecordsOutput{}
	err = c.client.Do("PutRecords", "POST", "/", req, resp)
	return
}

// RemoveTagsFromStream deletes tags from the specified Amazon Kinesis
// stream. If you specify a tag that does not exist, it is ignored.
func (c *Kinesis) RemoveTagsFromStream(req RemoveTagsFromStreamInput) (err error) {
	// NRE
	err = c.client.Do("RemoveTagsFromStream", "POST", "/", req, nil)
	return
}

// SplitShard splits a shard into two new shards in the stream, to increase
// the stream's capacity to ingest and transport data. SplitShard is called
// when there is a need to increase the overall capacity of stream because
// of an expected increase in the volume of data records being ingested.
// You can also use SplitShard when a shard appears to be approaching its
// maximum utilization, for example, when the set of producers sending data
// into the specific shard are suddenly sending more than previously
// anticipated. You can also call SplitShard to increase stream capacity,
// so that more Amazon Kinesis applications can simultaneously read data
// from the stream for real-time processing. You must specify the shard to
// be split and the new hash key, which is the position in the shard where
// the shard gets split in two. In many cases, the new hash key might
// simply be the average of the beginning and ending hash key, but it can
// be any hash key value in the range being mapped into the shard. For more
// information about splitting shards, see Split a Shard in the Amazon
// Kinesis Developer Guide You can use DescribeStream to determine the
// shard ID and hash key values for the ShardToSplit and NewStartingHashKey
// parameters that are specified in the SplitShard request. SplitShard is
// an asynchronous operation. Upon receiving a SplitShard request, Amazon
// Kinesis immediately returns a response and sets the stream status to .
// After the operation is completed, Amazon Kinesis sets the stream status
// to . Read and write operations continue to work while the stream is in
// the state. You can use DescribeStream to check the status of the stream,
// which is returned in StreamStatus . If the stream is in the state, you
// can call SplitShard . If a stream is in or or states, DescribeStream
// returns a ResourceInUseException If the specified stream does not exist,
// DescribeStream returns a ResourceNotFoundException . If you try to
// create more shards than are authorized for your account, you receive a
// LimitExceededException . The default limit for an AWS account is 10
// shards per stream. If you need to create a stream with more than 10
// shards, contact AWS Support to increase the limit on your account. If
// you try to operate on too many streams in parallel using CreateStream ,
// DeleteStream , MergeShards or SplitShard , you receive a
// LimitExceededException . SplitShard has limit of 5 transactions per
// second per account.
func (c *Kinesis) SplitShard(req SplitShardInput) (err error) {
	// NRE
	err = c.client.Do("SplitShard", "POST", "/", req, nil)
	return
}

// AddTagsToStreamInput is undocumented.
type AddTagsToStreamInput struct {
	StreamName string            `json:"StreamName"`
	Tags       map[string]string `json:"Tags"`
}

// CreateStreamInput is undocumented.
type CreateStreamInput struct {
	ShardCount int    `json:"ShardCount"`
	StreamName string `json:"StreamName"`
}

// DeleteStreamInput is undocumented.
type DeleteStreamInput struct {
	StreamName string `json:"StreamName"`
}

// DescribeStreamInput is undocumented.
type DescribeStreamInput struct {
	ExclusiveStartShardID string `json:"ExclusiveStartShardId,omitempty"`
	Limit                 int    `json:"Limit,omitempty"`
	StreamName            string `json:"StreamName"`
}

// DescribeStreamOutput is undocumented.
type DescribeStreamOutput struct {
	StreamDescription StreamDescription `json:"StreamDescription"`
}

// GetRecordsInput is undocumented.
type GetRecordsInput struct {
	Limit         int    `json:"Limit,omitempty"`
	ShardIterator string `json:"ShardIterator"`
}

// GetRecordsOutput is undocumented.
type GetRecordsOutput struct {
	NextShardIterator string   `json:"NextShardIterator,omitempty"`
	Records           []Record `json:"Records"`
}

// GetShardIteratorInput is undocumented.
type GetShardIteratorInput struct {
	ShardID                string `json:"ShardId"`
	ShardIteratorType      string `json:"ShardIteratorType"`
	StartingSequenceNumber string `json:"StartingSequenceNumber,omitempty"`
	StreamName             string `json:"StreamName"`
}

// GetShardIteratorOutput is undocumented.
type GetShardIteratorOutput struct {
	ShardIterator string `json:"ShardIterator,omitempty"`
}

// HashKeyRange is undocumented.
type HashKeyRange struct {
	EndingHashKey   string `json:"EndingHashKey"`
	StartingHashKey string `json:"StartingHashKey"`
}

// ListStreamsInput is undocumented.
type ListStreamsInput struct {
	ExclusiveStartStreamName string `json:"ExclusiveStartStreamName,omitempty"`
	Limit                    int    `json:"Limit,omitempty"`
}

// ListStreamsOutput is undocumented.
type ListStreamsOutput struct {
	HasMoreStreams bool     `json:"HasMoreStreams"`
	StreamNames    []string `json:"StreamNames"`
}

// ListTagsForStreamInput is undocumented.
type ListTagsForStreamInput struct {
	ExclusiveStartTagKey string `json:"ExclusiveStartTagKey,omitempty"`
	Limit                int    `json:"Limit,omitempty"`
	StreamName           string `json:"StreamName"`
}

// ListTagsForStreamOutput is undocumented.
type ListTagsForStreamOutput struct {
	HasMoreTags bool  `json:"HasMoreTags"`
	Tags        []Tag `json:"Tags"`
}

// MergeShardsInput is undocumented.
type MergeShardsInput struct {
	AdjacentShardToMerge string `json:"AdjacentShardToMerge"`
	ShardToMerge         string `json:"ShardToMerge"`
	StreamName           string `json:"StreamName"`
}

// PutRecordInput is undocumented.
type PutRecordInput struct {
	Data                      []byte `json:"Data"`
	ExplicitHashKey           string `json:"ExplicitHashKey,omitempty"`
	PartitionKey              string `json:"PartitionKey"`
	SequenceNumberForOrdering string `json:"SequenceNumberForOrdering,omitempty"`
	StreamName                string `json:"StreamName"`
}

// PutRecordOutput is undocumented.
type PutRecordOutput struct {
	SequenceNumber string `json:"SequenceNumber"`
	ShardID        string `json:"ShardId"`
}

// PutRecordsInput is undocumented.
type PutRecordsInput struct {
	Records    []PutRecordsRequestEntry `json:"Records"`
	StreamName string                   `json:"StreamName"`
}

// PutRecordsOutput is undocumented.
type PutRecordsOutput struct {
	FailedRecordCount int                     `json:"FailedRecordCount,omitempty"`
	Records           []PutRecordsResultEntry `json:"Records"`
}

// PutRecordsRequestEntry is undocumented.
type PutRecordsRequestEntry struct {
	Data            []byte `json:"Data"`
	ExplicitHashKey string `json:"ExplicitHashKey,omitempty"`
	PartitionKey    string `json:"PartitionKey"`
}

// PutRecordsResultEntry is undocumented.
type PutRecordsResultEntry struct {
	ErrorCode      string `json:"ErrorCode,omitempty"`
	ErrorMessage   string `json:"ErrorMessage,omitempty"`
	SequenceNumber string `json:"SequenceNumber,omitempty"`
	ShardID        string `json:"ShardId,omitempty"`
}

// Record is undocumented.
type Record struct {
	Data           []byte `json:"Data"`
	PartitionKey   string `json:"PartitionKey"`
	SequenceNumber string `json:"SequenceNumber"`
}

// RemoveTagsFromStreamInput is undocumented.
type RemoveTagsFromStreamInput struct {
	StreamName string   `json:"StreamName"`
	TagKeys    []string `json:"TagKeys"`
}

// SequenceNumberRange is undocumented.
type SequenceNumberRange struct {
	EndingSequenceNumber   string `json:"EndingSequenceNumber,omitempty"`
	StartingSequenceNumber string `json:"StartingSequenceNumber"`
}

// Shard is undocumented.
type Shard struct {
	AdjacentParentShardID string              `json:"AdjacentParentShardId,omitempty"`
	HashKeyRange          HashKeyRange        `json:"HashKeyRange"`
	ParentShardID         string              `json:"ParentShardId,omitempty"`
	SequenceNumberRange   SequenceNumberRange `json:"SequenceNumberRange"`
	ShardID               string              `json:"ShardId"`
}

// SplitShardInput is undocumented.
type SplitShardInput struct {
	NewStartingHashKey string `json:"NewStartingHashKey"`
	ShardToSplit       string `json:"ShardToSplit"`
	StreamName         string `json:"StreamName"`
}

// StreamDescription is undocumented.
type StreamDescription struct {
	HasMoreShards bool    `json:"HasMoreShards"`
	Shards        []Shard `json:"Shards"`
	StreamARN     string  `json:"StreamARN"`
	StreamName    string  `json:"StreamName"`
	StreamStatus  string  `json:"StreamStatus"`
}

// Tag is undocumented.
type Tag struct {
	Key   string `json:"Key"`
	Value string `json:"Value,omitempty"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name
