using MongoDB.Bson.Serialization.Attributes;
namespace DataTransferObjects;
public class RequestUserModel
{
    [BsonElement("EmailId")]
    public string? EmailId { get; set; }

    [BsonElement("Age")]
    public int ? Age { get; set; }
}
